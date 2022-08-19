/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mspext_test

import (
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator"
	mspext2 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/mspext"
)

// ensure structs implement expected interfaces
var (
	_ protolator.VariablyOpaqueFieldProto = &mspext2.MSPConfig{}
	_ protolator.DecoratedProto           = &mspext2.MSPConfig{}

	_ protolator.VariablyOpaqueFieldProto = &mspext2.MSPPrincipal{}
	_ protolator.DecoratedProto           = &mspext2.MSPPrincipal{}
)
