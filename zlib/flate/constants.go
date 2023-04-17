// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package flate implements the DEFLATE compressed data format, described in
// RFC 1951.  The gzip and zlib packages implement access to DEFLATE-based file
// formats.
package flate

const (
	maxNumLit = 286
)

// An InternalError reports an error in the flate code itself.
type InternalError string

func (e InternalError) Error() string { return "flate: internal error: " + string(e) }
