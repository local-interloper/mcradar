package mcushort

import (
	"bufio"
	"encoding/binary"
	"io"
)

type McUShort struct {
	bytes []byte
	Value uint16
}

func New(value int) *McUShort {
	b := make([]byte, 2)

	binary.BigEndian.PutUint16(b, uint16(value))

	return &McUShort{
		bytes: b,
		Value: uint16(value),
	}
}

func (m *McUShort) FromStream(br *bufio.Reader) error {
	b := make([]byte, 2)
	_, err := io.ReadFull(br, b)
	if err != nil {
		return err
	}

	m.bytes = b
	m.Value = binary.BigEndian.Uint16(b)

	return nil
}

func (m *McUShort) ToStream(bw *bufio.Writer) error {
	if _, err := bw.Write(m.bytes); err != nil {
		return err
	}

	return nil
}

func (m *McUShort) Bytes() []byte {
	return m.bytes
}

func (m *McUShort) Length() int {
	return len(m.bytes)
}
