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
	"context"
	http2 "github.com/hyperledger/fabric-sdk-go/third_party/smalgo/gmhttp"
)

var requestIDKey = requestIDKeyType{}

type requestIDKeyType struct{}

func RequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return reqID
	}
	return "unknown"
}

type GenerateIDFunc func() string

type requestID struct {
	generateID GenerateIDFunc
	next       http2.Handler
}

func WithRequestID(generator GenerateIDFunc) Middleware {
	return func(next http2.Handler) http2.Handler {
		return &requestID{next: next, generateID: generator}
	}
}

func (r *requestID) ServeHTTP(w http2.ResponseWriter, req *http2.Request) {
	reqID := req.Header.Get("X-Request-Id")
	if reqID == "" {
		reqID = r.generateID()
		req.Header.Set("X-Request-Id", reqID)
	}

	ctx := context.WithValue(req.Context(), requestIDKey, reqID)
	req = req.WithContext(ctx)

	w.Header().Add("X-Request-Id", reqID)

	r.next.ServeHTTP(w, req)
}
