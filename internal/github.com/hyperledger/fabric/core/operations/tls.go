/*
Copyright IBM Corp All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/

package operations

import (
	"github.com/hyperledger/fabric-sdk-go/third_party/smalgo/gmtls"
	"github.com/hyperledger/fabric-sdk-go/third_party/smalgo/x509"
	"io/ioutil"

	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/sdkinternal/pkg/comm"
)

type TLS struct {
	Enabled            bool
	CertFile           string
	KeyFile            string
	ClientCertRequired bool
	ClientCACertFiles  []string
}

func (t TLS) Config() (*gmtls.Config, error) {
	var tlsConfig *gmtls.Config

	if t.Enabled {
		cert, err := gmtls.LoadX509KeyPair(t.CertFile, t.KeyFile)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		for _, caPath := range t.ClientCACertFiles {
			caPem, err := ioutil.ReadFile(caPath)
			if err != nil {
				return nil, err
			}
			caCertPool.AppendCertsFromPEM(caPem)
		}
		tlsConfig = &gmtls.Config{
			Certificates: []gmtls.Certificate{cert},
			CipherSuites: comm.DefaultTLSCipherSuites,
			ClientCAs:    caCertPool,
		}
		if t.ClientCertRequired {
			tlsConfig.ClientAuth = gmtls.RequireAndVerifyClientCert
		} else {
			tlsConfig.ClientAuth = gmtls.VerifyClientCertIfGiven
		}
	}

	return tlsConfig, nil
}
