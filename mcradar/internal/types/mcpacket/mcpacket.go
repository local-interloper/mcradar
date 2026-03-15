package mcpacket

import (
	"bufio"

	"github.com/local-interloper/mcradar/mcradar/internal/types/mcserializable"
	"github.com/local-interloper/mcradar/mcradar/internal/types/mcvarint"
)

type McPacket struct {
	Length   *mcvarint.McVarInt
	Protocol *mcvarint.McVarInt
	Payload  []mcserializable.McSerializable
}

func New(protocol int, payload ...mcserializable.McSerializable) *McPacket {
	p := mcvarint.New(protocol)

	packet := &McPacket{
		Protocol: p,
		Payload:  payload,
	}

	length := p.Length()
	for _, item := range payload {
		length += item.Length()
	}

	packet.Length = mcvarint.New(length)

	return packet
}

func WithPayload(payload ...mcserializable.McSerializable) *McPacket {
	return &McPacket{
		Length:   mcvarint.New(0),
		Protocol: mcvarint.New(0),
		Payload:  payload,
	}
}

func (m *McPacket) FromStream(br *bufio.Reader) error {
	if err := m.Length.FromStream(br); err != nil {
		return err
	}

	if err := m.Protocol.FromStream(br); err != nil {
		return err
	}

	for _, item := range m.Payload {
		if err := item.FromStream(br); err != nil {
			return err
		}
	}

	return nil
}

func (m *McPacket) ToStream(bw *bufio.Writer) error {
	if err := m.Length.ToStream(bw); err != nil {
		return err
	}

	if err := m.Protocol.ToStream(bw); err != nil {
		return err
	}

	for _, item := range m.Payload {
		if err := item.ToStream(bw); err != nil {
			return err
		}
	}

	return nil
}
