/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package keyvaluestore

import (
	"github.com/pkg/errors"
	"github.com/wsw365904/fabric-sdk-go/pkg/common/providers/core"
	"github.com/wsw365904/fabric-sdk-go/pkg/util/cache"
)

var orgAdminKeyCertCache = cache.NewCache()

func SetKeyValueToOrgAdminKeyCertCache(key string, value []byte) {
	if !orgAdminKeyCertCache.IsExist(key) {
		orgAdminKeyCertCache.Set(key, value, -1)
	}
}

// CacheKeyValueStore stores each value into a separate file.
// KeySerializer maps a key to a unique file path (raletive to the store path)
// ValueSerializer and ValueDeserializer serializes/de-serializes a value
// to and from a byte array that is stored in the path derived from the key.
type CacheKeyValueStore struct {
	hash          string
	keySerializer KeySerializer
	marshaller    Marshaller
	unmarshaller  Unmarshaller
}

// CacheKeyValueStoreOptions allow overriding store defaults
type CacheKeyValueStoreOptions struct {
	// Store path, mandatory
	Hash string
	// Optional. If not provided, default key serializer is used.
	KeySerializer KeySerializer
	// Optional. If not provided, default Marshaller is used.
	Marshaller Marshaller
	// Optional. If not provided, default Unmarshaller is used.
	Unmarshaller Unmarshaller
}

// GetPath returns the store path
func (ckvs *CacheKeyValueStore) GetHash() string {
	return ckvs.hash
}

// NewCache creates a new instance of CacheKeyValueStore using provided options
func NewCache(opts *CacheKeyValueStoreOptions) (*CacheKeyValueStore, error) {
	if opts == nil {
		return nil, errors.New("CacheKeyValueStoreOptions is nil")
	}
	if opts.Hash == "" {
		return nil, errors.New("CacheKeyValueStoreOptions Hash is empty")
	}
	if opts.KeySerializer == nil {
		// Default key serializer
		opts.KeySerializer = func(key interface{}) (string, error) {
			keyString, ok := key.(string)
			if !ok {
				return "", errors.New("converting key to string failed")
			}
			return keyString, nil
		}
	}
	if opts.Marshaller == nil {
		opts.Marshaller = defaultMarshaller
	}
	if opts.Unmarshaller == nil {
		opts.Unmarshaller = defaultUnmarshaller
	}
	return &CacheKeyValueStore{
		hash:          opts.Hash,
		keySerializer: opts.KeySerializer,
		marshaller:    opts.Marshaller,
		unmarshaller:  opts.Unmarshaller,
	}, nil
}

// Load returns the value stored in the store for a key.
// If a value for the key was not found, returns (nil, ErrNotFound)
func (ckvs *CacheKeyValueStore) Load(key interface{}) (interface{}, error) {
	hash, err := ckvs.keySerializer(key)
	if err != nil {
		return nil, err
	}
	if hash == "" {
		logger.Warn("hash == \"\"")
		return nil, core.ErrKeyValueNotFound // wsw add
	}
	var (
		// errNotFound is the error of key not found.
		errNotFound = errors.New("cachego: key not found")
	)
	bytes, err := orgAdminKeyCertCache.Get(hash) // nolint: gas
	if err != nil {
		if err == errNotFound {
			logger.Warn("orgAdminKeyCertCache.Get", err)
			return nil, core.ErrKeyValueNotFound // wsw add
		}
		return nil, err
	}
	if bytes == nil {
		logger.Warnf("read value (%v) success but content is nil", hash)
		return nil, core.ErrKeyValueNotFound // wsw add
	}
	return ckvs.unmarshaller(bytes.([]byte))
}

// Store sets the value for the key.
func (ckvs *CacheKeyValueStore) Store(key interface{}, value interface{}) error {
	if key == nil {
		return errors.New("key is nil")
	}
	if value == nil {
		return errors.New("value is nil")
	}
	hash, err := ckvs.keySerializer(key)
	if err != nil {
		return err
	}
	valueBytes, err := ckvs.marshaller(value)
	if err != nil {
		return err
	}
	orgAdminKeyCertCache.Set(hash, valueBytes, -1)
	return nil
}

// Delete deletes the value for a key.
func (ckvs *CacheKeyValueStore) Delete(key interface{}) error {
	if key == nil {
		return errors.New("key is nil")
	}
	hash, err := ckvs.keySerializer(key)
	if err != nil {
		return err
	}
	orgAdminKeyCertCache.Del(hash)
	return nil
}
