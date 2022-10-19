/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"encoding/hex"
	"github.com/hellobchain/fabric-sdk-go/pkg/util/algo"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hellobchain/fabric-sdk-go/pkg/common/providers/msp"
)

// IdentityManager implements fab/IdentityManager
type IdentityManager struct {
	orgName         string
	orgMSPID        string
	config          fab.EndpointConfig
	cryptoSuite     core.CryptoSuite
	embeddedUsers   map[string]fab.CertKeyPair
	mspPrivKeyStore core.KVStore
	mspCertStore    core.KVStore
	userStore       msp.UserStore
}

// NewIdentityManager creates a new instance of IdentityManager
func NewIdentityManager(orgName string, userStore msp.UserStore, cryptoSuite core.CryptoSuite, endpointConfig fab.EndpointConfig) (*IdentityManager, error) {

	netConfig := endpointConfig.NetworkConfig()
	// viper keys are case insensitive
	orgConfig, ok := netConfig.Organizations[strings.ToLower(orgName)]
	if !ok {
		return nil, errors.New("org config retrieval failed")
	}

	if orgConfig.CryptoPath == "" && len(orgConfig.Users) == 0 && (len(orgConfig.SignedCert) == 0 || len(orgConfig.AdminPrivateKey) == 0) {
		return nil, errors.New("Either a cryptopath or an embedded list of users or admin key and admin sign cert is required")
	}

	var mspPrivKeyStore core.KVStore
	var mspCertStore core.KVStore

	orgCryptoPathTemplate := orgConfig.CryptoPath
	signedCert := orgConfig.SignedCert
	adminPrivateKey := orgConfig.AdminPrivateKey
	if orgCryptoPathTemplate != "" {
		var err error
		if !filepath.IsAbs(orgCryptoPathTemplate) {
			orgCryptoPathTemplate = filepath.Join(endpointConfig.CryptoConfigPath(), orgCryptoPathTemplate)
		}
		mspPrivKeyStore, err = NewFileKeyStore(orgCryptoPathTemplate)
		if err != nil {
			return nil, errors.Wrap(err, "creating a private key store failed")
		}
		mspCertStore, err = NewFileCertStore(orgCryptoPathTemplate)
		if err != nil {
			return nil, errors.Wrap(err, "creating a cert store failed")
		}
	} else if len(signedCert) != 0 && len(adminPrivateKey) != 0 {
		var err error
		hasher := algo.GetDefaultHash().New()
		hasher.Write(adminPrivateKey)
		keyHash := hex.EncodeToString(hasher.Sum(nil))
		mspPrivKeyStore, err = NewCacheKeyStore(keyHash, adminPrivateKey)
		if err != nil {
			return nil, errors.Wrap(err, "creating a admin private key store failed")
		}
		logger.Debug("admin key", "store key", keyHash, "value", string(adminPrivateKey))
		hasher = algo.GetDefaultHash().New()
		hasher.Write(signedCert)
		keyHash = hex.EncodeToString(hasher.Sum(nil))
		mspCertStore, err = NewCacheCertStore(keyHash, signedCert)
		if err != nil {
			return nil, errors.Wrap(err, "creating a admin signed cert store failed")
		}
		logger.Debug("admin signed cert", "store key", keyHash, "value", string(signedCert))
	} else {
		logger.Warnf("Cryptopath not provided for organization [%s], admin not provided for organization [%s] MSP stores not created", orgName, orgName)
	}

	mgr := &IdentityManager{
		orgName:         orgName,
		orgMSPID:        orgConfig.MSPID,
		config:          endpointConfig,
		cryptoSuite:     cryptoSuite,
		mspPrivKeyStore: mspPrivKeyStore,
		mspCertStore:    mspCertStore,
		embeddedUsers:   orgConfig.Users,
		userStore:       userStore,
		// CA Client state is created lazily, when (if) needed
	}
	return mgr, nil
}
