package mculong

import (
	"bufio"
	"encoding/binary"
	"io"
)

type McULong struct {
	bytes []byte
	Value uint64
}

func New(value int) *McULong {
	b := make([]byte, 8)

	binary.BigEndian.PutUint64(b, uint64(value))

	return &McULong{
		bytes: b,
		Value: uint64(value),
	}
}

func (m *McULong) FromStream(br *bufio.Reader) error {
	b := make([]byte, 8)
	_, err := io.ReadFull(br, b)
	if err != nil {
		return err
	}

	m.bytes = b
	m.Value = binary.BigEndian.Uint64(b)

	return nil
}

func (m *McULong) ToStream(bw *bufio.Writer) error {
	if _, err := bw.Write(m.bytes); err != nil {
		return err
	}

	return nil
}

func (m *McULong) Bytes() []byte {
	return m.bytes
}

func (m *McULong) Length() int {
	return len(m.bytes)
}
