/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mocks

import (
	"github.com/hellobchain/fabric-sdk-go/pkg/util/algo"
	"hash"

	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/core"
)

// MockCryptoSuite implementation
type MockCryptoSuite struct {
}

// KeyGen mock key gen
func (m *MockCryptoSuite) KeyGen(opts core.KeyGenOpts) (k core.Key, err error) {
	return nil, nil
}

// KeyImport mock key import
func (m *MockCryptoSuite) KeyImport(raw interface{},
	opts core.KeyImportOpts) (k core.Key, err error) {
	return nil, nil
}

// GetKey mock get key
func (m *MockCryptoSuite) GetKey(ski []byte) (k core.Key, err error) {
	return nil, nil
}

// Hash mock hash
func (m *MockCryptoSuite) Hash(msg []byte, opts core.HashOpts) (hash []byte, err error) {
	return nil, nil
}

// GetHash mock get hash
func (m *MockCryptoSuite) GetHash(opts core.HashOpts) (hash.Hash, error) {
	h := algo.GetDefaultHash().New()
	return h, nil
}

// Sign mock signing
func (m *MockCryptoSuite) Sign(k core.Key, digest []byte,
	opts core.SignerOpts) (signature []byte, err error) {
	return []byte("testSignature"), nil
}

//Verify mock verify implementation
func (m *MockCryptoSuite) Verify(k core.Key, signature, digest []byte, opts core.SignerOpts) (valid bool, err error) {
	return true, nil
}
