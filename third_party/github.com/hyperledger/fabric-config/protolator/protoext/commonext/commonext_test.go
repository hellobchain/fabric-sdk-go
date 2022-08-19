/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package commonext_test

import (
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator"
	commonext2 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/commonext"
)

// ensure structs implement expected interfaces
var (
	_ protolator.StaticallyOpaqueFieldProto      = &commonext2.Envelope{}
	_ protolator.DecoratedProto                  = &commonext2.Envelope{}
	_ protolator.VariablyOpaqueFieldProto        = &commonext2.Payload{}
	_ protolator.DecoratedProto                  = &commonext2.Payload{}
	_ protolator.StaticallyOpaqueFieldProto      = &commonext2.Header{}
	_ protolator.DecoratedProto                  = &commonext2.Header{}
	_ protolator.StaticallyOpaqueFieldProto      = &commonext2.SignatureHeader{}
	_ protolator.DecoratedProto                  = &commonext2.SignatureHeader{}
	_ protolator.StaticallyOpaqueSliceFieldProto = &commonext2.BlockData{}
	_ protolator.DecoratedProto                  = &commonext2.BlockData{}

	_ protolator.StaticallyOpaqueFieldProto    = &commonext2.ConfigUpdateEnvelope{}
	_ protolator.DecoratedProto                = &commonext2.ConfigUpdateEnvelope{}
	_ protolator.StaticallyOpaqueFieldProto    = &commonext2.ConfigSignature{}
	_ protolator.DecoratedProto                = &commonext2.ConfigSignature{}
	_ protolator.DynamicFieldProto             = &commonext2.Config{}
	_ protolator.DecoratedProto                = &commonext2.Config{}
	_ protolator.StaticallyOpaqueMapFieldProto = &commonext2.ConfigUpdate{}
	_ protolator.DecoratedProto                = &commonext2.ConfigUpdate{}

	_ protolator.DynamicMapFieldProto       = &commonext2.DynamicChannelGroup{}
	_ protolator.DecoratedProto             = &commonext2.DynamicChannelGroup{}
	_ protolator.StaticallyOpaqueFieldProto = &commonext2.DynamicChannelConfigValue{}
	_ protolator.DecoratedProto             = &commonext2.DynamicChannelConfigValue{}
	_ protolator.DynamicMapFieldProto       = &commonext2.DynamicConsortiumsGroup{}
	_ protolator.DecoratedProto             = &commonext2.DynamicConsortiumsGroup{}
	_ protolator.DynamicMapFieldProto       = &commonext2.DynamicConsortiumGroup{}
	_ protolator.DecoratedProto             = &commonext2.DynamicConsortiumGroup{}
	_ protolator.VariablyOpaqueFieldProto   = &commonext2.DynamicConsortiumConfigValue{}
	_ protolator.DecoratedProto             = &commonext2.DynamicConsortiumConfigValue{}
	_ protolator.DynamicMapFieldProto       = &commonext2.DynamicConsortiumOrgGroup{}
	_ protolator.DecoratedProto             = &commonext2.DynamicConsortiumOrgGroup{}
	_ protolator.StaticallyOpaqueFieldProto = &commonext2.DynamicConsortiumOrgConfigValue{}
	_ protolator.DecoratedProto             = &commonext2.DynamicConsortiumOrgConfigValue{}
)
