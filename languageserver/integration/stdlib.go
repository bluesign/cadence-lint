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

package integration

import (
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/cadence/runtime/sema"
	"github.com/onflow/cadence/runtime/stdlib"
	evmstdlib "github.com/onflow/flow-go/fvm/evm/stdlib"
)

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
		evmstdlib.InternalEVMContractType.ID(),
		interpreter.ConvertSemaCompositeTypeToStaticCompositeType(
			nil,
			evmstdlib.InternalEVMContractType,
		),
		evmstdlib.InternalEVMContractType.Fields,
		map[string]interpreter.Value{
			evmstdlib.InternalEVMTypeRunFunctionName:                       newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeRunFunctionType),
			evmstdlib.InternalEVMTypeBatchRunFunctionName:                  newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeBatchRunFunctionType),
			evmstdlib.InternalEVMTypeCreateCadenceOwnedAccountFunctionName: newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeCreateCadenceOwnedAccountFunctionType),
			evmstdlib.InternalEVMTypeCallFunctionName:                      newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeCallFunctionType),
			evmstdlib.InternalEVMTypeDepositFunctionName:                   newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeDepositFunctionType),
			evmstdlib.InternalEVMTypeWithdrawFunctionName:                  newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeWithdrawFunctionType),
			evmstdlib.InternalEVMTypeDeployFunctionName:                    newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeDepositFunctionType),
			evmstdlib.InternalEVMTypeBalanceFunctionName:                   newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeBalanceFunctionType),
			evmstdlib.InternalEVMTypeNonceFunctionName:                     newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeNonceFunctionType),
			evmstdlib.InternalEVMTypeCodeFunctionName:                      newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeCodeFunctionType),
			evmstdlib.InternalEVMTypeCodeHashFunctionName:                  newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeCodeHashFunctionType),
			evmstdlib.InternalEVMTypeEncodeABIFunctionName:                 newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeEncodeABIFunctionType),
			evmstdlib.InternalEVMTypeDecodeABIFunctionName:                 newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeDecodeABIFunctionType),
			evmstdlib.InternalEVMTypeCastToAttoFLOWFunctionName:            newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeCastToAttoFLOWFunctionType),
			evmstdlib.InternalEVMTypeCastToFLOWFunctionName:                newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeCastToAttoFLOWFunctionType),
			evmstdlib.InternalEVMTypeGetLatestBlockFunctionName:            newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeGetLatestBlockFunctionType),
			evmstdlib.InternalEVMTypeDryRunFunctionName:                    newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeDryRunFunctionType),
			evmstdlib.InternalEVMTypeCommitBlockProposalFunctionName:       newInternalEVMFunction(gauge, evmstdlib.InternalEVMTypeCommitBlockProposalFunctionType),
		},
		nil,
		nil,
		nil,
	)
}

// FVMtandardLibraryValues returns the standard library values which are provided by the FVM
// these are not part of the Cadence standard library
func FVMStandardLibraryValues() []stdlib.StandardLibraryValue {
	return []stdlib.StandardLibraryValue{
		{
			Name:  evmstdlib.InternalEVMContractName,
			Type:  evmstdlib.InternalEVMContractType,
			Value: NewInternalEVMContractValue(nil, common.AddressLocation{}),
			Kind:  common.DeclarationKindContract,
		},
	}
}
