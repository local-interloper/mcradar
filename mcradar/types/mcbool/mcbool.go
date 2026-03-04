package mcbool

import (
	"bufio"
)

var boolMap map[bool]byte = map[bool]byte{
	false: 0x00,
	true:  0x01,
}

var boolUnmap map[byte]bool = map[byte]bool{
	0x00: false,
	0x01: true,
}

type McBool struct {
	bytes []byte
	Value bool
}

func New(value bool) *McBool {
	return &McBool{
		bytes: []byte{boolMap[value]},
		Value: value,
	}
}

func (m *McBool) FromStream(br *bufio.Reader) error {
	b, err := br.ReadByte()
	if err != nil {
		return err
	}

	m.bytes = []byte{b}
	m.Value = boolUnmap[b]

	return nil
}

func (m *McBool) ToStream(bw *bufio.Writer) error {
	if _, err := bw.Write(m.bytes); err != nil {
		return err
	}

	return nil
}

func (m *McBool) Bytes() []byte {
	return m.bytes
}

func (m *McBool) Length() int {
	return len(m.bytes)
}
