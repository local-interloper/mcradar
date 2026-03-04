package mcserializable

import "bufio"

type McSerializable interface {
	// Read values from buffer
	FromStream(*bufio.Reader) error

	// Write values to buffer
	ToStream(*bufio.Writer) error

	// Bytes
	Bytes() []byte

	// Count of undelrying bytes
	Length() int
}
