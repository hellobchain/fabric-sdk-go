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

// NewCacheCertStore ...
func NewCacheCertStore(certHash string, certByte []byte) (core.KVStore, error) {
	keyValueCache := cache.NewCache()
	opts := &keyvaluestore.CacheKeyValueStoreOptions{
		Hash: certHash,
		KeySerializer: func(key interface{}) (string, error) {
			keyValueCache.Set(certHash, certByte, -1)
			return certHash, nil
		},
		KeyValueCache: keyValueCache,
	}
	return keyvaluestore.NewCache(opts)
}
