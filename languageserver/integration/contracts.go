/*
 * Cadence languageserver - The Cadence language server
 *
 * Copyright Flow Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"fmt"

	"github.com/onflow/cadence/ast"
	"github.com/onflow/cadence/interpreter"
	"github.com/onflow/cadence/parser"
	"github.com/onflow/cadence/runtime"
	"github.com/onflow/cadence/sema"

	"github.com/onflow/cadence-tools/languageserver/conversion"
	"github.com/onflow/cadence-tools/languageserver/protocol"
)

type contractKind uint8

const (
	contractTypeUnknown contractKind = iota
	contractTypeDeclaration
	contractTypeInterface
)

type contractInfo struct {
	uri                   protocol.DocumentURI
	documentVersion       int32
	startPos              *ast.Position
	kind                  contractKind
	name                  string
	parameters            []sema.Parameter
	pragmaArgumentStrings []string
	pragmaArguments       [][]Argument
	pragmaSignersNames    []string
}

func (c *contractInfo) update(uri protocol.DocumentURI, version int32, checker *sema.Checker) {
	if c.documentVersion == version {
		return // if no change in version do nothing
	}

	var docString string
	contractDeclaration := checker.Program.SoleContractDeclaration()
	contractInterfaceDeclaration := checker.Program.SoleContractInterfaceDeclaration()

	if contractDeclaration != nil {
		docString = contractDeclaration.DocString
		c.name = contractDeclaration.Identifier.String()
		c.startPos = &contractDeclaration.StartPos
		c.kind = contractTypeDeclaration
		contractType := checker.Elaboration.CompositeDeclarationType(contractDeclaration)
		c.parameters = contractType.ConstructorParameters
	} else if contractInterfaceDeclaration != nil {
		docString = contractInterfaceDeclaration.DocString
		c.name = contractInterfaceDeclaration.Identifier.String()
		c.startPos = &contractInterfaceDeclaration.StartPos
		c.kind = contractTypeInterface
	}

	c.pragmaArguments = nil
	c.pragmaArgumentStrings = nil
	c.pragmaSignersNames = nil

	if c.startPos != nil {
		parameterTypes := make([]sema.Type, len(c.parameters))

		for i, parameter := range c.parameters {
			parameterTypes[i] = parameter.TypeAnnotation.Type
		}

		if len(c.parameters) > 0 {

			inter, _ := interpreter.NewInterpreter(nil, nil, &interpreter.Config{})

			for _, pragmaArgumentString := range parser.ParseDocstringPragmaArguments(docString) {
				arguments, err := runtime.ParseLiteralArgumentList(pragmaArgumentString, parameterTypes, inter)
				// TODO: record error and show diagnostic
				if err != nil {
					continue
				}

				convertedArguments := make([]Argument, len(arguments))
				for i, arg := range arguments {
					convertedArguments[i] = Argument{arg}
				}

				c.pragmaArgumentStrings = append(c.pragmaArgumentStrings, pragmaArgumentString)
				c.pragmaArguments = append(c.pragmaArguments, convertedArguments)
			}
		}

		for _, pragmaSignerString := range parser.ParseDocstringPragmaSigners(docString) {
			signers := SignersRegexp.FindAllString(pragmaSignerString, -1)
			c.pragmaSignersNames = append(c.pragmaSignersNames, signers...)
		}
	}

	c.uri = uri
	c.documentVersion = version
}

func (c *contractInfo) codelens(client flowClient) []*protocol.CodeLens {
	if c.kind == contractTypeUnknown || c.startPos == nil {
		return nil
	}
	codelensRange := conversion.ASTToProtocolRange(*c.startPos, *c.startPos)

	var signer string
	if len(c.pragmaSignersNames) == 0 {
		signer = client.GetActiveClientAccount().Name // default to active client account
	} else if len(c.pragmaSignersNames) == 1 {
		signer = c.pragmaSignersNames[0]
	} else {
		title := fmt.Sprintf("%s Only possible to deploy to a single account", prefixError)
		return []*protocol.CodeLens{makeActionlessCodelens(title, codelensRange)}
	}

	account := client.GetClientAccount(signer)
	if account == nil {
		title := fmt.Sprintf("%s Specified account %s does not exist", prefixError, signer)
		return []*protocol.CodeLens{makeActionlessCodelens(title, codelensRange)}
	}

	titleBody := "Deploy contract"
	if c.kind == contractTypeInterface {
		titleBody = "Deploy contract interface"
	}

	title := fmt.Sprintf("%s %s %s to %s", prefixOK, titleBody, c.name, signer)
	arguments, _ := encodeJSONArguments(c.uri, c.name, signer)
	return []*protocol.CodeLens{makeCodeLens(CommandDeployContract, title, codelensRange, arguments)}
}
