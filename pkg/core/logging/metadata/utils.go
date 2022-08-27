/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package metadata

import (
	"github.com/wsw365904/fabric-sdk-go/pkg/core/logging/api"
)

//Log level names in string
var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"INFO",
	"DEBUG",
}

//ParseString returns String repressentation of given log level
func ParseString(level api.Level) string {
	return levelNames[level]
}
