// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2021-Present The Zarf Authors

// Package helpers provides generic helper functions with no external imports
package helpers

import (
	"crypto"
	"encoding/hex"
	"hash/crc32"
	"io"
)

// GetCryptoHash returns the computed SHA256 Sum of a given file
func GetCryptoHash(data io.ReadCloser, hashName crypto.Hash) (string, error) {
	hash := hashName.New()
	if _, err := io.Copy(hash, data); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// GetCRCHash returns the computed CRC32 Sum of a given string
func GetCRCHash(text string) uint32 {
	table := crc32.MakeTable(crc32.IEEE)
	return crc32.Checksum([]byte(text), table)
}
