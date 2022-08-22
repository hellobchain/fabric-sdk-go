// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/hyperledger/fabric-sdk-go

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/hyperledger/fabric-protos-go v0.0.0-20220816110612-c96c610ca7b4
	github.com/miekg/pkcs11 v1.1.1
	github.com/mitchellh/mapstructure v1.4.1
	github.com/pkg/errors v0.9.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.1
	github.com/wsw365904/newcryptosm v0.0.0-20220822153922-8852ae7b6d34
	github.com/wsw365904/third_party v0.0.0-20220822153856-449e7dccc2bc
	golang.org/x/crypto v0.0.0-20220817201139-bc19a97f63c8
	golang.org/x/net v0.0.0-20220812174116-3211cb980234
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/hyperledger/fabric-protos-go v0.0.0-20220816110612-c96c610ca7b4 => github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23

go 1.14
