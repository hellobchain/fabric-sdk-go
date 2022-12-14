/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sw

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/hellobchain/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp"
	"github.com/hellobchain/newcryptosm/ecdsa"
)

type ecdsaKeyGenerator struct {
	curve elliptic.Curve
}

func (kg *ecdsaKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	privKey, err := ecdsa.GenerateKey(kg.curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("Failed generating ECDSA key for [%v]: [%s]", kg.curve, err)
	}

	return &ecdsaPrivateKey{privKey}, nil
}

type sm2KeyGenerator struct {
	curve elliptic.Curve
}

func (kg *sm2KeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (k bccsp.Key, err error) {
	privKey, err := ecdsa.GenerateKey(kg.curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("Failed generating SM2 key for [%v]: [%s]", kg.curve, err)
	}

	return &ecdsaPrivateKey{privKey}, nil
}

type symmetryKeyGenerator struct {
	length int
}

func (kg *symmetryKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	lowLevelKey, err := GetRandomBytes(int(kg.length))
	if err != nil {
		return nil, fmt.Errorf("Failed generating Symmetry %d key [%s]", kg.length, err)
	}

	return &symmetryPrivateKey{lowLevelKey, false}, nil
}
