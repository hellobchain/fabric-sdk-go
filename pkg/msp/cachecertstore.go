/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"github.com/wsw365904/fabric-sdk-go/pkg/common/providers/core"
	"github.com/wsw365904/fabric-sdk-go/pkg/fab/keyvaluestore"
)

// NewCacheCertStore ...
func NewCacheCertStore(certHash string, certByte []byte) (core.KVStore, error) {
	opts := &keyvaluestore.CacheKeyValueStoreOptions{
		Hash: certHash,
		KeySerializer: func(key interface{}) (string, error) {
			if !keyvaluestore.IsExist(certHash) {
				keyvaluestore.SetGlobalCache(certHash, certByte)
			}
			return certHash, nil
		},
	}
	return keyvaluestore.NewCache(opts)
}
