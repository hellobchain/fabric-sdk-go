/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"github.com/wsw365904/fabric-sdk-go/pkg/common/providers/core"
	"github.com/wsw365904/fabric-sdk-go/pkg/fab/keyvaluestore"
	"github.com/wsw365904/fabric-sdk-go/pkg/util/cache"
)

// NewCacheKeyStore loads keys stored in the cryptoconfig directory layout.
// This function will detect if private keys are stored in v1 or v2 format.
func NewCacheKeyStore(keyHash string, keyBytes []byte) (core.KVStore, error) {
	keyValueCache := cache.NewCache()
	opts := &keyvaluestore.CacheKeyValueStoreOptions{
		Hash: keyHash,
		KeySerializer: func(key interface{}) (string, error) {
			keyValueCache.Set(keyHash, keyBytes, -1)
			return keyHash, nil
		},
		KeyValueCache: keyValueCache,
	}
	return keyvaluestore.NewCache(opts)
}
