/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protoext

import (
	commonext2 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/commonext"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/ledger/rwsetext"
	mspext2 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/mspext"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/ordererext"
	peerext2 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/peerext"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric-protos-go/peer"
	. "github.com/onsi/gomega"
)

type GenericProtoMessage struct {
	GenericField string
}

func (g *GenericProtoMessage) Reset() {
	panic("not implemented")
}

func (g *GenericProtoMessage) String() string {
	return "not implemented"
}

func (g *GenericProtoMessage) ProtoMessage() {
	panic("not implemented")
}

func TestDecorate(t *testing.T) {
	tests := []struct {
		testSpec       string
		msg            proto.Message
		expectedReturn proto.Message
	}{
		{
			testSpec: "common.BlockData",
			msg: &common.BlockData{
				Data: [][]byte{
					[]byte("data-bytes"),
				},
			},
			expectedReturn: &commonext2.BlockData{
				BlockData: &common.BlockData{
					Data: [][]byte{
						[]byte("data-bytes"),
					},
				},
			},
		},
		{
			testSpec: "common.Config",
			msg: &common.Config{
				Sequence: 5,
			},
			expectedReturn: &commonext2.Config{
				Config: &common.Config{
					Sequence: 5,
				},
			},
		},
		{
			testSpec: "common.ConfigSignature",
			msg: &common.ConfigSignature{
				SignatureHeader: []byte("signature-header-bytes"),
			},
			expectedReturn: &commonext2.ConfigSignature{
				ConfigSignature: &common.ConfigSignature{
					SignatureHeader: []byte("signature-header-bytes"),
				},
			},
		},
		{
			testSpec: "common.ConfigUpdate",
			msg: &common.ConfigUpdate{
				ChannelId: "testchannel",
			},
			expectedReturn: &commonext2.ConfigUpdate{
				ConfigUpdate: &common.ConfigUpdate{
					ChannelId: "testchannel",
				},
			},
		},
		{
			testSpec: "common.ConfigUpdateEnvelope",
			msg: &common.ConfigUpdateEnvelope{
				ConfigUpdate: []byte("config-update-bytes"),
			},
			expectedReturn: &commonext2.ConfigUpdateEnvelope{
				ConfigUpdateEnvelope: &common.ConfigUpdateEnvelope{
					ConfigUpdate: []byte("config-update-bytes"),
				},
			},
		},
		{
			testSpec: "common.Envelope",
			msg: &common.Envelope{
				Payload: []byte("payload-bytes"),
			},
			expectedReturn: &commonext2.Envelope{
				Envelope: &common.Envelope{
					Payload: []byte("payload-bytes"),
				},
			},
		},
		{
			testSpec: "common.Header",
			msg: &common.Header{
				ChannelHeader: []byte("channel-header-bytes"),
			},
			expectedReturn: &commonext2.Header{
				Header: &common.Header{
					ChannelHeader: []byte("channel-header-bytes"),
				},
			},
		},
		{
			testSpec: "common.ChannelHeader",
			msg: &common.ChannelHeader{
				Type: 5,
			},
			expectedReturn: &commonext2.ChannelHeader{
				ChannelHeader: &common.ChannelHeader{
					Type: 5,
				},
			},
		},
		{
			testSpec: "common.SignatureHeader",
			msg: &common.SignatureHeader{
				Creator: []byte("creator-bytes"),
			},
			expectedReturn: &commonext2.SignatureHeader{
				SignatureHeader: &common.SignatureHeader{
					Creator: []byte("creator-bytes"),
				},
			},
		},
		{
			testSpec: "common.Payload",
			msg: &common.Payload{
				Header: &common.Header{ChannelHeader: []byte("channel-header-bytes")},
			},
			expectedReturn: &commonext2.Payload{
				Payload: &common.Payload{
					Header: &common.Header{ChannelHeader: []byte("channel-header-bytes")},
				},
			},
		},
		{
			testSpec: "common.Policy",
			msg: &common.Policy{
				Type: 5,
			},
			expectedReturn: &commonext2.Policy{
				Policy: &common.Policy{
					Type: 5,
				},
			},
		},
		{
			testSpec: "msp.MSPConfig",
			msg: &msp.MSPConfig{
				Type: 5,
			},
			expectedReturn: &mspext2.MSPConfig{
				MSPConfig: &msp.MSPConfig{
					Type: 5,
				},
			},
		},
		{
			testSpec: "msp.MSPPrincipal",
			msg: &msp.MSPPrincipal{
				Principal: []byte("principal-bytes"),
			},
			expectedReturn: &mspext2.MSPPrincipal{
				MSPPrincipal: &msp.MSPPrincipal{
					Principal: []byte("principal-bytes"),
				},
			},
		},
		{
			testSpec: "orderer.ConsensusType",
			msg: &orderer.ConsensusType{
				Type: "etcdraft",
			},
			expectedReturn: &ordererext.ConsensusType{
				ConsensusType: &orderer.ConsensusType{
					Type: "etcdraft",
				},
			},
		},
		{
			testSpec: "peer.ChaincodeAction",
			msg: &peer.ChaincodeAction{
				Results: []byte("results-bytes"),
			},
			expectedReturn: &peerext2.ChaincodeAction{
				ChaincodeAction: &peer.ChaincodeAction{
					Results: []byte("results-bytes"),
				},
			},
		},
		{
			testSpec: "peer.ChaincodeActionPayload",
			msg: &peer.ChaincodeActionPayload{
				ChaincodeProposalPayload: []byte("chaincode-proposal-payload-bytes"),
			},
			expectedReturn: &peerext2.ChaincodeActionPayload{
				ChaincodeActionPayload: &peer.ChaincodeActionPayload{
					ChaincodeProposalPayload: []byte("chaincode-proposal-payload-bytes"),
				},
			},
		},
		{
			testSpec: "peer.ChaincodeEndorsedAction",
			msg: &peer.ChaincodeEndorsedAction{
				ProposalResponsePayload: []byte("proposal-response-payload-bytes"),
			},
			expectedReturn: &peerext2.ChaincodeEndorsedAction{
				ChaincodeEndorsedAction: &peer.ChaincodeEndorsedAction{
					ProposalResponsePayload: []byte("proposal-response-payload-bytes"),
				},
			},
		},
		{
			testSpec: "peer.ChaincodeProposalPayload",
			msg: &peer.ChaincodeProposalPayload{
				Input: []byte("input-bytes"),
			},
			expectedReturn: &peerext2.ChaincodeProposalPayload{
				ChaincodeProposalPayload: &peer.ChaincodeProposalPayload{
					Input: []byte("input-bytes"),
				},
			},
		},
		{
			testSpec: "peer.ProposalResponsePayload",
			msg: &peer.ProposalResponsePayload{
				ProposalHash: []byte("proposal-hash-bytes"),
			},
			expectedReturn: &peerext2.ProposalResponsePayload{
				ProposalResponsePayload: &peer.ProposalResponsePayload{
					ProposalHash: []byte("proposal-hash-bytes"),
				},
			},
		},
		{
			testSpec: "peer.TransactionAction",
			msg: &peer.TransactionAction{
				Header: []byte("header-bytes"),
			},
			expectedReturn: &peerext2.TransactionAction{
				TransactionAction: &peer.TransactionAction{
					Header: []byte("header-bytes"),
				},
			},
		},
		{
			testSpec: "rwset.TxReadWriteSet",
			msg: &rwset.TxReadWriteSet{
				NsRwset: []*rwset.NsReadWriteSet{
					{
						Namespace: "namespace",
					},
				},
			},
			expectedReturn: &rwsetext.TxReadWriteSet{
				TxReadWriteSet: &rwset.TxReadWriteSet{
					NsRwset: []*rwset.NsReadWriteSet{
						{
							Namespace: "namespace",
						},
					},
				},
			},
		},
		{
			testSpec: "default",
			msg: &GenericProtoMessage{
				GenericField: "test",
			},
			expectedReturn: &GenericProtoMessage{
				GenericField: "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testSpec, func(t *testing.T) {
			gt := NewGomegaWithT(t)
			decoratedMsg := Decorate(tt.msg)
			gt.Expect(proto.Equal(decoratedMsg, tt.expectedReturn)).To(BeTrue())
		})
	}
}
