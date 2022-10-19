/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/

package middleware

import (
	http2 "github.com/hellobchain/newcryptosm/http"
)

type requireCert struct {
	next http2.Handler
}

// RequireCert is used to ensure that a verified TLS client certificate was
// used for authentication.
func RequireCert() Middleware {
	return func(next http2.Handler) http2.Handler {
		return &requireCert{next: next}
	}
}

func (r *requireCert) ServeHTTP(w http2.ResponseWriter, req *http2.Request) {
	switch {
	case req.TLS == nil:
		fallthrough
	case len(req.TLS.VerifiedChains) == 0:
		fallthrough
	case len(req.TLS.VerifiedChains[0]) == 0:
		w.WriteHeader(http2.StatusUnauthorized)
	default:
		r.next.ServeHTTP(w, req)
	}
}
