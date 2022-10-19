/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hellobchain/fabric-sdk-go/pkg/fab/keyvaluestore"
)

// NewCacheCertStore ...
func NewCacheCertStore(certHash string, certByte []byte) (core.KVStore, error) {
	keyValueCache := make(map[string]interface{})
	opts := &keyvaluestore.CacheKeyValueStoreOptions{
		Hash: certHash,
		KeySerializer: func(key interface{}) (string, error) {
			keyValueCache[certHash] = certByte
			return certHash, nil
		},
		KeyValueCache: keyValueCache,
	}
	return keyvaluestore.NewCache(opts)
}
