/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package peerext_test

import (
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator"
	peerext2 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/peerext"
)

// ensure structs implement expected interfaces
var (
	_ protolator.DynamicMapFieldProto       = &peerext2.DynamicApplicationGroup{}
	_ protolator.DecoratedProto             = &peerext2.DynamicApplicationGroup{}
	_ protolator.DynamicMapFieldProto       = &peerext2.DynamicApplicationOrgGroup{}
	_ protolator.DecoratedProto             = &peerext2.DynamicApplicationOrgGroup{}
	_ protolator.StaticallyOpaqueFieldProto = &peerext2.DynamicApplicationConfigValue{}
	_ protolator.DecoratedProto             = &peerext2.DynamicApplicationConfigValue{}
	_ protolator.StaticallyOpaqueFieldProto = &peerext2.DynamicApplicationOrgConfigValue{}
	_ protolator.DecoratedProto             = &peerext2.DynamicApplicationOrgConfigValue{}

	_ protolator.StaticallyOpaqueFieldProto = &peerext2.ChaincodeProposalPayload{}
	_ protolator.DecoratedProto             = &peerext2.ChaincodeProposalPayload{}
	_ protolator.StaticallyOpaqueFieldProto = &peerext2.ChaincodeAction{}
	_ protolator.DecoratedProto             = &peerext2.ChaincodeAction{}

	_ protolator.StaticallyOpaqueFieldProto = &peerext2.ProposalResponsePayload{}
	_ protolator.DecoratedProto             = &peerext2.ProposalResponsePayload{}

	_ protolator.StaticallyOpaqueFieldProto = &peerext2.TransactionAction{}
	_ protolator.DecoratedProto             = &peerext2.TransactionAction{}
	_ protolator.StaticallyOpaqueFieldProto = &peerext2.ChaincodeActionPayload{}
	_ protolator.DecoratedProto             = &peerext2.ChaincodeActionPayload{}
	_ protolator.StaticallyOpaqueFieldProto = &peerext2.ChaincodeEndorsedAction{}
	_ protolator.DecoratedProto             = &peerext2.ChaincodeEndorsedAction{}
)
