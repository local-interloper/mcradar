package mcbytearray

import (
	"bufio"
	"io"

	"github.com/local-interloper/mc-radar/mcradar/types/mcvarint"
)

type McByteArray struct {
	bytes []byte
	Value []byte
}

func New(value []byte) *McByteArray {
	length := mcvarint.New(len(value))

	fullBytes := append(length.Bytes(), value...)

	return &McByteArray{
		bytes: fullBytes,
		Value: value,
	}
}

func (m *McByteArray) FromStream(br *bufio.Reader) error {
	length := mcvarint.New(0)

	if err := length.FromStream(br); err != nil {
		return err
	}

	buf := make([]byte, length.Value)
	if _, err := io.ReadFull(br, buf); err != nil {
		return err
	}

	m.bytes = append(length.Bytes(), buf...)
	m.Value = buf

	return nil
}

func (m *McByteArray) ToStream(bw *bufio.Writer) error {
	if _, err := bw.Write(m.bytes); err != nil {
		return err
	}

	return nil
}

func (m *McByteArray) Bytes() []byte {
	return m.bytes
}

func (m *McByteArray) Length() int {
	return len(m.bytes)
}
