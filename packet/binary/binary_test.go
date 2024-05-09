// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binary implements simple translation between numbers and byte
// sequences and encoding and decoding of varints.
//
// Numbers are translated by reading and writing fixed-size values.
// A fixed-size value is either a fixed-size arithmetic
// type (bool, int8, uint8, int16, float32, complex64, ...)
// or an array or struct containing only fixed-size values.
//
// The varint functions encode and decode single integer values using
// a variable-length encoding; smaller values require fewer bytes.
// For a specification, see
// https://developers.google.com/protocol-buffers/docs/encoding.
//
// This package favors simplicity over efficiency. Clients that require
// high-performance serialization, especially for large data structures,
// should look at more advanced solutions such as the [encoding/gob]
// package or [google.golang.org/protobuf] for protocol buffers.
package binary

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	type teststruct struct {
		magic   [3]byte // "bsr"
		version byte    // protocol version
		Number  uint16
	}

	ts := teststruct{
		magic:   [3]byte{0x45, 0x46, 0x47},
		version: 0x0a,
		Number:  0xaabb,
	}

	bytesBuffer := bytes.NewBuffer([]byte{})

	err := Write(bytesBuffer, BigEndian, ts)
	assert.NoError(t, err)

	res := make([]byte, 7) // one more just in case we could read junk
	bytesRead, err := bytesBuffer.Read(res)

	assert.NoError(t, err)
	assert.Equal(t, 6, bytesRead)
	assert.Equal(t, []byte{0x45, 0x46, 0x47, 0x0a, 0xaa, 0xbb, 0x00}, res)
}
