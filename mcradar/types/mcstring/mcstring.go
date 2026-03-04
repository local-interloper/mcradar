package mcstring

import (
	"bufio"

	"github.com/local-interloper/mc-radar/mcradar/types/mcvarint"
)

type McString struct {
	bytes []byte
	Value string
}

func New(value string, length int) *McString {
	return FromStringWithLength(value, length)
}

func FromString(value string) *McString {
	b := []byte{}

	strLength := len(value)
	strLengthVarInt := mcvarint.FromInt32(int32(strLength))

	b = append(b, strLengthVarInt.Bytes()...)
	b = append(b, []byte(value)...)

	return &McString{
		bytes: b,
		Value: value,
	}
}

func FromStringWithLength(value string, length int) *McString {
	strLength := len(value)
	strLengthVarInt := mcvarint.FromInt32(int32(strLength))

	b := make([]byte, strLengthVarInt.Length()+strLength)

	copy(b, strLengthVarInt.Bytes())
	copy(b[strLengthVarInt.Length():], []byte(value))

	return &McString{
		bytes: b,
		Value: value,
	}
}

func FromBytes(rawBytes []byte) *McString {
	stringLength := mcvarint.FromBytes(rawBytes)
	relevantBytes := rawBytes[:stringLength.Length()+int(stringLength.Value)]
	stringBytes := rawBytes[stringLength.Length()-1 : stringLength.Length()+int(stringLength.Value)]

	s := string(stringBytes)

	return &McString{
		bytes: relevantBytes,
		Value: s,
	}
}

func (m *McString) Bytes() []byte {
	return m.bytes
}

func (m *McString) Length() int {
	return len(m.bytes)
}

func (m *McString) FromStream(br *bufio.Reader) error {
	stringBytes := []byte{}

	length := mcvarint.New(0)
	if err := length.FromStream(br); err != nil {
		return err
	}

	for i := 0; i < int(length.Value); i++ {
		b, err := br.ReadByte()

		if err != nil {
			return err
		}

		stringBytes = append(stringBytes, b)
	}

	m.bytes = append(length.Bytes(), stringBytes...)
	m.Value = string(stringBytes)

	return nil
}

func (m *McString) ToStream(bw *bufio.Writer) error {
	if _, err := bw.Write(m.bytes); err != nil {
		return err
	}

	return nil
}
