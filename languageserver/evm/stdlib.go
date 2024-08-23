/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
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

package evm

import (
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/stdlib"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/cadence/runtime/sema"
)

const (
	ContractName = "EVM"

	evmAddressTypeBytesFieldName      = "bytes"
	evmAddressTypeQualifiedIdentifier = "EVM.EVMAddress"
	evmBalanceTypeQualifiedIdentifier = "EVM.Balance"

	evmResultTypeQualifiedIdentifier       = "EVM.Result"
	evmResultTypeStatusFieldName           = "status"
	evmResultTypeErrorCodeFieldName        = "errorCode"
	evmResultTypeErrorMessageFieldName     = "errorMessage"
	evmResultTypeGasUsedFieldName          = "gasUsed"
	evmResultTypeDataFieldName             = "data"
	evmResultTypeDeployedContractFieldName = "deployedContract"

	evmStatusTypeQualifiedIdentifier = "EVM.Status"

	evmBlockTypeQualifiedIdentifier = "EVM.EVMBlock"
	abiEncodingByteSize             = 32
	AddressLength                   = 20
)

var (
	EVMTransactionBytesCadenceType = cadence.NewVariableSizedArrayType(cadence.UInt8Type)

	evmTransactionBytesType       = sema.NewVariableSizedType(nil, sema.UInt8Type)
	evmTransactionsBatchBytesType = sema.NewVariableSizedType(nil, evmTransactionBytesType)
	evmAddressBytesType           = sema.NewConstantSizedType(nil, sema.UInt8Type, AddressLength)

	evmAddressBytesStaticType = interpreter.ConvertSemaArrayTypeToStaticArrayType(nil, evmAddressBytesType)

	EVMAddressBytesCadenceType = cadence.NewConstantSizedArrayType(AddressLength, cadence.UInt8Type)
)

const internalEVMTypeEncodeABIFunctionName = "encodeABI"

var internalEVMTypeEncodeABIFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:      sema.ArgumentLabelNotRequired,
			Identifier: "values",
			TypeAnnotation: sema.NewTypeAnnotation(
				sema.NewVariableSizedType(nil, sema.AnyStructType),
			),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.ByteArrayType),
}

func newInternalEVMTypeEncodeABIFunction(
	gauge common.MemoryGauge,
	location common.AddressLocation,
) *interpreter.HostFunctionValue {
	return nil
}

// EVM.decodeABI

const internalEVMTypeDecodeABIFunctionName = "decodeABI"

var internalEVMTypeDecodeABIFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Identifier: "types",
			TypeAnnotation: sema.NewTypeAnnotation(
				sema.NewVariableSizedType(nil, sema.MetaType),
			),
		},
		{
			Label:          "data",
			TypeAnnotation: sema.NewTypeAnnotation(sema.ByteArrayType),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(
		sema.NewVariableSizedType(nil, sema.AnyStructType),
	),
}

const internalEVMTypeRunFunctionName = "run"

var internalEVMTypeRunFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "tx",
			TypeAnnotation: sema.NewTypeAnnotation(evmTransactionBytesType),
		},
		{
			Label:          "coinbase",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
	},
	// Actually EVM.Result, but cannot refer to it here
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.AnyStructType),
}

// dry run

const internalEVMTypeDryRunFunctionName = "dryRun"

var internalEVMTypeDryRunFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "tx",
			TypeAnnotation: sema.NewTypeAnnotation(evmTransactionBytesType),
		},
		{
			Label:          "from",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
	},
	// Actually EVM.Result, but cannot refer to it here
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.AnyStructType),
}

const internalEVMTypeBatchRunFunctionName = "batchRun"

var internalEVMTypeBatchRunFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "txs",
			TypeAnnotation: sema.NewTypeAnnotation(evmTransactionsBatchBytesType),
		},
		{
			Label:          "coinbase",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
	},
	// Actually [EVM.Result], but cannot refer to it here
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.NewVariableSizedType(nil, sema.AnyStructType)),
}

const internalEVMTypeCallFunctionName = "call"

var internalEVMTypeCallFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "from",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
		{
			Label:          "to",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
		{
			Label:          "data",
			TypeAnnotation: sema.NewTypeAnnotation(sema.ByteArrayType),
		},
		{
			Label:          "gasLimit",
			TypeAnnotation: sema.NewTypeAnnotation(sema.UInt64Type),
		},
		{
			Label:          "value",
			TypeAnnotation: sema.NewTypeAnnotation(sema.UIntType),
		},
	},
	// Actually EVM.Result, but cannot refer to it here
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.AnyStructType),
}

const internalEVMTypeCreateCadenceOwnedAccountFunctionName = "createCadenceOwnedAccount"

var internalEVMTypeCreateCadenceOwnedAccountFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "uuid",
			TypeAnnotation: sema.NewTypeAnnotation(sema.UInt64Type),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
}

const internalEVMTypeDepositFunctionName = "deposit"

var internalEVMTypeDepositFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "from",
			TypeAnnotation: sema.NewTypeAnnotation(sema.AnyResourceType),
		},
		{
			Label:          "to",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.VoidType),
}

const fungibleTokenVaultTypeBalanceFieldName = "balance"
const internalEVMTypeBalanceFunctionName = "balance"

var internalEVMTypeBalanceFunctionType = &sema.FunctionType{
	Purity: sema.FunctionPurityView,
	Parameters: []sema.Parameter{
		{
			Label:          "address",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.UIntType),
}

const internalEVMTypeNonceFunctionName = "nonce"

var internalEVMTypeNonceFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "address",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.UInt64Type),
}

const internalEVMTypeCodeFunctionName = "code"

var internalEVMTypeCodeFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "address",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.ByteArrayType),
}

const internalEVMTypeCodeHashFunctionName = "codeHash"

var internalEVMTypeCodeHashFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "address",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.ByteArrayType),
}

const internalEVMTypeWithdrawFunctionName = "withdraw"

var internalEVMTypeWithdrawFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "from",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
		{
			Label:          "amount",
			TypeAnnotation: sema.NewTypeAnnotation(sema.UIntType),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.AnyResourceType),
}

const internalEVMTypeDeployFunctionName = "deploy"

var internalEVMTypeDeployFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Label:          "from",
			TypeAnnotation: sema.NewTypeAnnotation(evmAddressBytesType),
		},
		{
			Label:          "code",
			TypeAnnotation: sema.NewTypeAnnotation(sema.ByteArrayType),
		},
		{
			Label:          "gasLimit",
			TypeAnnotation: sema.NewTypeAnnotation(sema.UInt64Type),
		},
		{
			Label:          "value",
			TypeAnnotation: sema.NewTypeAnnotation(sema.UIntType),
		},
	},
	// Actually EVM.Result, but cannot refer to it here
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.AnyStructType),
}

const internalEVMTypeCastToAttoFLOWFunctionName = "castToAttoFLOW"

var internalEVMTypeCastToAttoFLOWFunctionType = &sema.FunctionType{
	Purity: sema.FunctionPurityView,
	Parameters: []sema.Parameter{
		{
			Label:          "balance",
			TypeAnnotation: sema.NewTypeAnnotation(sema.UFix64Type),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.UIntType),
}

const internalEVMTypeCastToFLOWFunctionName = "castToFLOW"

var internalEVMTypeCastToFLOWFunctionType = &sema.FunctionType{
	Purity: sema.FunctionPurityView,
	Parameters: []sema.Parameter{
		{
			Label:          "balance",
			TypeAnnotation: sema.NewTypeAnnotation(sema.UIntType),
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.UFix64Type),
}

const internalEVMTypeCommitBlockProposalFunctionName = "commitBlockProposal"

var internalEVMTypeCommitBlockProposalFunctionType = &sema.FunctionType{
	Parameters:           []sema.Parameter{},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.VoidType),
}

const internalEVMTypeGetLatestBlockFunctionName = "getLatestBlock"

var internalEVMTypeGetLatestBlockFunctionType = &sema.FunctionType{
	Parameters: []sema.Parameter{},
	// Actually EVM.Block, but cannot refer to it here
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.AnyStructType),
}

func newInternalEVMFunction(
	gauge common.MemoryGauge,
	t *sema.FunctionType,
) *interpreter.HostFunctionValue {
	return interpreter.NewStaticHostFunctionValue(
		gauge,
		t,
		func(invocation interpreter.Invocation) interpreter.Value {
			return nil
		},
	)
}

func NewInternalEVMContractValue(
	gauge common.MemoryGauge,
	location common.AddressLocation,
) *interpreter.SimpleCompositeValue {
	return interpreter.NewSimpleCompositeValue(
		gauge,
		InternalEVMContractType.ID(),
		internalEVMContractStaticType,
		InternalEVMContractType.Fields,
		map[string]interpreter.Value{
			internalEVMTypeRunFunctionName:                       newInternalEVMFunction(gauge, internalEVMTypeRunFunctionType),
			internalEVMTypeBatchRunFunctionName:                  newInternalEVMFunction(gauge, internalEVMTypeBatchRunFunctionType),
			internalEVMTypeCreateCadenceOwnedAccountFunctionName: newInternalEVMFunction(gauge, internalEVMTypeCreateCadenceOwnedAccountFunctionType),
			internalEVMTypeCallFunctionName:                      newInternalEVMFunction(gauge, internalEVMTypeCallFunctionType),
			internalEVMTypeDepositFunctionName:                   newInternalEVMFunction(gauge, internalEVMTypeDepositFunctionType),
			internalEVMTypeWithdrawFunctionName:                  newInternalEVMFunction(gauge, internalEVMTypeWithdrawFunctionType),
			internalEVMTypeDeployFunctionName:                    newInternalEVMFunction(gauge, internalEVMTypeDepositFunctionType),
			internalEVMTypeBalanceFunctionName:                   newInternalEVMFunction(gauge, internalEVMTypeBalanceFunctionType),
			internalEVMTypeNonceFunctionName:                     newInternalEVMFunction(gauge, internalEVMTypeNonceFunctionType),
			internalEVMTypeCodeFunctionName:                      newInternalEVMFunction(gauge, internalEVMTypeCodeFunctionType),
			internalEVMTypeCodeHashFunctionName:                  newInternalEVMFunction(gauge, internalEVMTypeCodeHashFunctionType),
			internalEVMTypeEncodeABIFunctionName:                 newInternalEVMFunction(gauge, internalEVMTypeEncodeABIFunctionType),
			internalEVMTypeDecodeABIFunctionName:                 newInternalEVMFunction(gauge, internalEVMTypeDecodeABIFunctionType),
			internalEVMTypeCastToAttoFLOWFunctionName:            newInternalEVMFunction(gauge, internalEVMTypeCastToAttoFLOWFunctionType),
			internalEVMTypeCastToFLOWFunctionName:                newInternalEVMFunction(gauge, internalEVMTypeCastToAttoFLOWFunctionType),
			internalEVMTypeGetLatestBlockFunctionName:            newInternalEVMFunction(gauge, internalEVMTypeGetLatestBlockFunctionType),
			internalEVMTypeDryRunFunctionName:                    newInternalEVMFunction(gauge, internalEVMTypeDryRunFunctionType),
			internalEVMTypeCommitBlockProposalFunctionName:       newInternalEVMFunction(gauge, internalEVMTypeCommitBlockProposalFunctionType),
		},
		nil,
		nil,
		nil,
	)
}

const InternalEVMContractName = "InternalEVM"

var InternalEVMContractType = func() *sema.CompositeType {
	ty := &sema.CompositeType{
		Identifier: InternalEVMContractName,
		Kind:       common.CompositeKindContract,
	}

	ty.Members = sema.MembersAsMap([]*sema.Member{
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeRunFunctionName,
			internalEVMTypeRunFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeDryRunFunctionName,
			internalEVMTypeDryRunFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeBatchRunFunctionName,
			internalEVMTypeBatchRunFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeCreateCadenceOwnedAccountFunctionName,
			internalEVMTypeCreateCadenceOwnedAccountFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeCallFunctionName,
			internalEVMTypeCallFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeDepositFunctionName,
			internalEVMTypeDepositFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeWithdrawFunctionName,
			internalEVMTypeWithdrawFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeDeployFunctionName,
			internalEVMTypeDeployFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeCastToAttoFLOWFunctionName,
			internalEVMTypeCastToAttoFLOWFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeCastToFLOWFunctionName,
			internalEVMTypeCastToFLOWFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeBalanceFunctionName,
			internalEVMTypeBalanceFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeNonceFunctionName,
			internalEVMTypeNonceFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeCodeFunctionName,
			internalEVMTypeCodeFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeCodeHashFunctionName,
			internalEVMTypeCodeHashFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeEncodeABIFunctionName,
			internalEVMTypeEncodeABIFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeDecodeABIFunctionName,
			internalEVMTypeDecodeABIFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeGetLatestBlockFunctionName,
			internalEVMTypeGetLatestBlockFunctionType,
			"",
		),
		sema.NewUnmeteredPublicFunctionMember(
			ty,
			internalEVMTypeCommitBlockProposalFunctionName,
			internalEVMTypeCommitBlockProposalFunctionType,
			"",
		),
	})
	return ty
}()

var internalEVMContractStaticType = interpreter.ConvertSemaCompositeTypeToStaticCompositeType(
	nil,
	InternalEVMContractType,
)

func newInternalEVMStandardLibraryValue(
	gauge common.MemoryGauge,
	location common.AddressLocation,
) stdlib.StandardLibraryValue {
	return stdlib.StandardLibraryValue{
		Name:  InternalEVMContractName,
		Type:  InternalEVMContractType,
		Value: NewInternalEVMContractValue(gauge, location),
		Kind:  common.DeclarationKindContract,
	}
}

var internalEVMStandardLibraryType = stdlib.StandardLibraryType{
	Name: InternalEVMContractName,
	Type: InternalEVMContractType,
	Kind: common.DeclarationKindContract,
}

// FVMtandardLibraryValues returns the standard library values which are provided by the FVM
// these are not part of the Cadence standard library
func FVMStandardLibraryValues() []stdlib.StandardLibraryValue {
	return []stdlib.StandardLibraryValue{
		// InternalEVM contract
		{
			Name:  InternalEVMContractName,
			Type:  InternalEVMContractType,
			Value: NewInternalEVMContractValue(nil, common.AddressLocation{}),
			Kind:  common.DeclarationKindContract,
		},
	}
}
