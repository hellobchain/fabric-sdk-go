/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/

package keyutil

import (
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/third_party/smalgo/ecdsa"
	x5092 "github.com/hyperledger/fabric-sdk-go/third_party/smalgo/x509"
)

func derToPrivateKey(der []byte) (key interface{}, err error) {

	if key, err = x5092.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}

	if key, err = x5092.ParsePKCS8PrivateKey(der); err == nil {
		switch key.(type) {
		case *ecdsa.PrivateKey:
			return
		default:
			return nil, errors.New("found unknown private key type in PKCS#8 wrapping")
		}
	}

	if key, err = x5092.ParseECPrivateKey(der); err == nil {
		return
	}

	return nil, errors.New("invalid key type. The DER must contain an ecdsa.PrivateKey")
}

func PEMToPrivateKey(raw []byte, pwd []byte) (interface{}, error) {
	block, _ := pem.Decode(raw)
	if block == nil {
		return nil, fmt.Errorf("failed decoding PEM. Block must be different from nil [% x]", raw)
	}

	// TODO: derive from header the type of the key

	if x5092.IsEncryptedPEMBlock(block) {
		if len(pwd) == 0 {
			return nil, errors.New("encrypted Key. Need a password")
		}

		decrypted, err := x5092.DecryptPEMBlock(block, pwd)
		if err != nil {
			return nil, fmt.Errorf("failed PEM decryption: [%s]", err)
		}

		key, err := derToPrivateKey(decrypted)
		if err != nil {
			return nil, err
		}
		return key, err
	}

	cert, err := derToPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return cert, err
}
