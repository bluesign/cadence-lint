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
	"testing"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/cadence-tools/languageserver/protocol"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/parser"
	"github.com/onflow/cadence/runtime/sema"
	"github.com/stretchr/testify/require"
)

func Test_ContractUpdate(t *testing.T) {

	const code = `
      /// pragma signers Alice
      access(all)
      contract HelloWorld {

            access(all)
            let greeting: String

            access(all)
            fun hello(): String {
                return self.greeting
            }

            init(a: String) {
                self.greeting = a
            }
     }
    `

	program, err := parser.ParseProgram(nil, []byte(code), parser.Config{})
	require.NoError(t, err)

	location := common.StringLocation("foo")
	config := &sema.Config{
		AccessCheckMode: sema.AccessCheckModeStrict,
	}
	checker, err := sema.NewChecker(program, location, nil, config)
	require.NoError(t, err)

	err = checker.Check()
	require.NoError(t, err)

	client := &mockFlowClient{}

	t.Run("update contract information", func(t *testing.T) {
		contract := &contractInfo{}
		contract.update("Hello", 1, checker)

		assert.Equal(t, protocol.DocumentURI("Hello"), contract.uri)
		assert.Equal(t, "HelloWorld", contract.name)
		assert.Equal(t, contractTypeDeclaration, contract.kind)
		assert.Equal(t, [][]Argument(nil), contract.pragmaArguments)
		assert.Equal(t, []string(nil), contract.pragmaArgumentStrings)
		assert.Equal(t, []string{"Alice"}, contract.pragmaSignersNames)

		assert.Len(t, contract.parameters, 1)
		assert.Equal(t, "a", contract.parameters[0].Identifier)
		assert.Equal(t, "", contract.parameters[0].Label)
	})

	t.Run("get codelenses", func(t *testing.T) {
		contract := &contractInfo{}
		contract.update("Hello", 1, checker)

		alice := &clientAccount{
			Account: &flow.Account{
				Address: flow.HexToAddress("0x1"),
			},
			Name:   "Alice",
			Active: true,
		}

		client.
			On("GetActiveClientAccount").
			Return(alice)

		client.
			On("GetClientAccount", "Alice").
			Return(alice)

		codelenses := contract.codelens(client)

		require.Len(t, codelenses, 1)
		codeLens := codelenses[0]
		assert.Equal(t, "💡 Deploy contract HelloWorld to Alice", codeLens.Command.Title)
		assert.Equal(t, "cadence.server.flow.deployContract", codeLens.Command.Command)
		assert.Equal(t, nil, codeLens.Data)
		assert.Equal(t,
			protocol.Range{
				Start: protocol.Position{Line: 2, Character: 6},
				End:   protocol.Position{Line: 2, Character: 7},
			},
			codeLens.Range,
		)
	})
}
