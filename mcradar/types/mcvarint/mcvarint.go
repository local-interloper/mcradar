package mcvarint

import (
	"bufio"

	"github.com/local-interloper/mc-radar/mcradar/consts"
)

type McVarInt struct {
	bytes []byte
	Value int32
}

func New(value int) *McVarInt {
	return FromInt32(int32(value))
}

func FromInt32(value int32) *McVarInt {
	v := value
	b := []byte{}
	uValue := uint32(value)

	for {
		if (uValue & 0xFFFFFF80) == 0 {
			b = append(b, byte(uValue))
			break
		}

		b = append(b, byte((uValue&consts.SegmentBits))|consts.ContinueBit)
		uValue >>= 7
	}

	return &McVarInt{
		bytes: b,
		Value: v,
	}
}

func FromBytes(rawBytes []byte) *McVarInt {
	var value uint32 = 0
	relevantBytes := []byte{}
	byteIndex := 0

	for {
		value |= (uint32(rawBytes[byteIndex] & consts.SegmentBits)) << (7 * byteIndex)
		relevantBytes = append(relevantBytes, rawBytes[byteIndex])

		if relevantBytes[byteIndex]&consts.ContinueBit == 0 {
			break
		}

		byteIndex++

		if byteIndex > 4 {
			break
		}
	}

	return &McVarInt{
		bytes: relevantBytes,
		Value: int32(value),
	}
}

func (m *McVarInt) Bytes() []byte {
	return m.bytes
}

func (m *McVarInt) Length() int {
	return len(m.bytes)
}

func (m *McVarInt) FromStream(br *bufio.Reader) error {
	var value uint32 = 0
	relevantBytes := []byte{}
	byteIndex := 0

	for {
		currentByte, err := br.ReadByte()
		if err != nil {
			return err
		}
		value |= (uint32(currentByte & consts.SegmentBits)) << (7 * byteIndex)
		relevantBytes = append(relevantBytes, currentByte)

		if relevantBytes[byteIndex]&consts.ContinueBit == 0 {
			break
		}

		byteIndex++

		if byteIndex > 4 {
			break
		}
	}

	m.bytes = relevantBytes
	m.Value = int32(value)

	return nil
}

func (m *McVarInt) ToStream(bw *bufio.Writer) error {
	if _, err := bw.Write(m.bytes); err != nil {
		return err
	}

	return nil
}
