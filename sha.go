package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"regexp"
)

const sha1short = 7

var sha1reg = regexp.MustCompile(`\b[0-9a-f]{40}\b`)

func newConversationID() string {
	b := make([]byte, 4096)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", sha1.Sum(b)) //nolint: gosec
}