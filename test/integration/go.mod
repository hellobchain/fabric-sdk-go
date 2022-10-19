// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/hellobchain/fabric-sdk-go/test/integration

replace github.com/hellobchain/fabric-sdk-go => ../../

require (
	github.com/golang/protobuf v1.5.2
	github.com/hyperledger/fabric-protos-go v0.0.0-20220816110612-c96c610ca7b4
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.1
	github.com/hellobchain/fabric-sdk-go v0.0.0-20220827110349-d15b2aaae095
	github.com/hellobchain/newcryptosm v0.0.0-20220822153922-8852ae7b6d34
	github.com/hellobchain/third_party v0.0.0-20220822153856-449e7dccc2bc
	google.golang.org/grpc v1.48.0
)

go 1.14
