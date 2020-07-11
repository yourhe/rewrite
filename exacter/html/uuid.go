package html

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"

	"github.com/google/uuid"
)

//NewID uuid
func NewID() string {
	return uuid.New().String()
}

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// with the padding stripped off.
func NewId() string {
	// var b bytes.Buffer
	// encoder := base32.NewEncoder(encoding, &b)
	// encoder.Write(uuid.New())
	// encoder.Close()
	// b.Truncate(26) // removes the '==' padding
	// return b.String()
	return NewRandomString(26)
}

func NewRandomString(length int) string {
	var b bytes.Buffer
	str := make([]byte, length+8)
	rand.Read(str)
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(str)
	encoder.Close()
	b.Truncate(length) // removes the '==' padding
	return b.String()
}
