/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protoext

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric-protos-go/peer"
	commonext2 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/commonext"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/ledger/rwsetext"
	mspext2 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/mspext"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/ordererext"
	peerext2 "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric-config/protolator/protoext/peerext"
)

// Docorate will add additional capabilities to some protobuf messages that
// enable proper JSON marshalling and unmarshalling in protolator.
func Decorate(msg proto.Message) proto.Message {
	switch m := msg.(type) {
	case *common.BlockData:
		return &commonext2.BlockData{BlockData: m}
	case *common.Config:
		return &commonext2.Config{Config: m}
	case *common.ConfigSignature:
		return &commonext2.ConfigSignature{ConfigSignature: m}
	case *common.ConfigUpdate:
		return &commonext2.ConfigUpdate{ConfigUpdate: m}
	case *common.ConfigUpdateEnvelope:
		return &commonext2.ConfigUpdateEnvelope{ConfigUpdateEnvelope: m}
	case *common.Envelope:
		return &commonext2.Envelope{Envelope: m}
	case *common.Header:
		return &commonext2.Header{Header: m}
	case *common.ChannelHeader:
		return &commonext2.ChannelHeader{ChannelHeader: m}
	case *common.SignatureHeader:
		return &commonext2.SignatureHeader{SignatureHeader: m}
	case *common.Payload:
		return &commonext2.Payload{Payload: m}
	case *common.Policy:
		return &commonext2.Policy{Policy: m}

	case *msp.MSPConfig:
		return &mspext2.MSPConfig{MSPConfig: m}
	case *msp.MSPPrincipal:
		return &mspext2.MSPPrincipal{MSPPrincipal: m}

	case *orderer.ConsensusType:
		return &ordererext.ConsensusType{ConsensusType: m}

	case *peer.ChaincodeAction:
		return &peerext2.ChaincodeAction{ChaincodeAction: m}
	case *peer.ChaincodeActionPayload:
		return &peerext2.ChaincodeActionPayload{ChaincodeActionPayload: m}
	case *peer.ChaincodeEndorsedAction:
		return &peerext2.ChaincodeEndorsedAction{ChaincodeEndorsedAction: m}
	case *peer.ChaincodeProposalPayload:
		return &peerext2.ChaincodeProposalPayload{ChaincodeProposalPayload: m}
	case *peer.ProposalResponsePayload:
		return &peerext2.ProposalResponsePayload{ProposalResponsePayload: m}
	case *peer.TransactionAction:
		return &peerext2.TransactionAction{TransactionAction: m}

	case *rwset.TxReadWriteSet:
		return &rwsetext.TxReadWriteSet{TxReadWriteSet: m}

	default:
		return msg
	}
}
