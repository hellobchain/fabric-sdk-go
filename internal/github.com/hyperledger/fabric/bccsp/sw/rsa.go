/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sw

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/wsw365904/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp"
	"github.com/wsw365904/newcryptosm"
	x5092 "github.com/wsw365904/newcryptosm/x509"
)

// An rsaPublicKey wraps the standard library implementation of an RSA public
// key with functions that satisfy the bccsp.Key interface.
//
// NOTE: Fabric does not support RSA signing or verification. This code simply
// allows MSPs to include RSA CAs in their certificate chains.
type rsaPublicKey struct{ pubKey *rsa.PublicKey }

func (k *rsaPublicKey) Symmetric() bool               { return false }
func (k *rsaPublicKey) Private() bool                 { return false }
func (k *rsaPublicKey) PublicKey() (bccsp.Key, error) { return k, nil }

// Bytes converts this key to its serialized representation.
func (k *rsaPublicKey) Bytes() (raw []byte, err error) {
	if k.pubKey == nil {
		return nil, errors.New("Failed marshalling key. Key is nil.")
	}
	raw, err = x5092.MarshalPKIXPublicKey(k.pubKey)
	if err != nil {
		return nil, fmt.Errorf("Failed marshalling key [%s]", err)
	}
	return
}

// SKI returns the subject key identifier of this key.
func (k *rsaPublicKey) SKI() []byte {
	if k.pubKey == nil {
		return nil
	}

	// Marshal the public key and hash it
	raw := x5092.MarshalPKCS1PublicKey(k.pubKey)
	hash := newcryptosm.SHA256.New()
	hash.Write(raw)
	return hash.Sum(nil)[:]
}
