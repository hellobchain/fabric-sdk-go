/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/wsw365904/fabric-sdk-go/pkg/common/providers/core"
	"github.com/wsw365904/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/wsw365904/fabric-sdk-go/pkg/fab/keyvaluestore"
)

// NewFileCertStore ...
func NewFileCertStore(cryptoConfigMSPPath string) (core.KVStore, error) {
	_, orgName := filepath.Split(filepath.Dir(filepath.Dir(filepath.Dir(cryptoConfigMSPPath))))
	opts := &keyvaluestore.FileKeyValueStoreOptions{
		Path: cryptoConfigMSPPath,
		KeySerializer: func(key interface{}) (string, error) {
			ck, ok := key.(*msp.IdentityIdentifier)
			if !ok {
				return "", errors.New("converting key to CertKey failed")
			}
			if ck == nil || ck.MSPID == "" || ck.ID == "" {
				return "", errors.New("invalid key")
			}

			// TODO: refactor to case insensitive or remove eventually.
			r := strings.NewReplacer("{userName}", ck.ID, "{username}", ck.ID)
			certDir := filepath.Join(r.Replace(cryptoConfigMSPPath), "signcerts")
			certPath := filepath.Join(certDir, fmt.Sprintf("%s@%s-cert.pem", ck.ID, orgName))
			_, err := os.Stat(certPath)
			if os.IsNotExist(err) {
				rd, err := ioutil.ReadDir(certDir)
				if err != nil {
					return "", err
				}
				for _, fi := range rd {
					if !fi.IsDir() {
						return filepath.Join(certDir, fi.Name()), nil
					}
				}
			}
			return certPath, nil // wsw add
		},
	}
	return keyvaluestore.New(opts)
}
