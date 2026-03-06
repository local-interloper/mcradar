package mcconnection

import (
	"bufio"
	"encoding/json"
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/local-interloper/mc-radar/mcradar/consts"
	"github.com/local-interloper/mc-radar/mcradar/types/mcpacket"
	"github.com/local-interloper/mc-radar/mcradar/types/mcstatus"
	"github.com/local-interloper/mc-radar/mcradar/types/mcstring"
	"github.com/local-interloper/mc-radar/mcradar/types/mculong"
	"github.com/local-interloper/mc-radar/mcradar/types/mcushort"
	"github.com/local-interloper/mc-radar/mcradar/types/mcvarint"
	"github.com/local-interloper/mc-radar/mcradar/types/servertype"
)

type Params struct {
	Address         string
	Port            int
	ProtocolVersion int
}

type McConnection struct {
	params     Params
	connection net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer
}

func Connect(params Params) (*McConnection, error) {
	d := net.Dialer{
		Timeout: consts.TimeoutTime,
	}

	con, err := d.Dial("tcp", net.JoinHostPort(params.Address, strconv.FormatUint(uint64(params.Port), 10)))
	if err != nil {
		return nil, err
	}

	return &McConnection{
		params:     params,
		connection: con,
		reader:     bufio.NewReader(con),
		writer:     bufio.NewWriter(con),
	}, nil
}

func (m *McConnection) Handshake(intent int) {
	protocolVersion := m.params.ProtocolVersion

	if m.params.ProtocolVersion == 0 {
		protocolVersion = 774
	}

	m.SendPacket(mcpacket.New(0x00,
		mcvarint.New(protocolVersion),
		mcstring.New(m.params.Address, 255),
		mcushort.New(m.params.Port),
		mcvarint.New(intent),
	))
}

func (m *McConnection) Status() mcstatus.McStatus {
	m.Handshake(1)

	m.SendPacket(mcpacket.New(
		0x00,
	))

	statusJson := new(mcstring.McString)
	response := mcpacket.WithPayload(
		statusJson,
	)

	m.ReadPacket(response)

	var status mcstatus.McStatus
	json.Unmarshal([]byte(statusJson.Value), &status)

	return status
}

func (m *McConnection) GetServerType() (servertype.ServerType, error) {
	m.Handshake(2)

	m.SendPacket(mcpacket.New(
		0x00,
		mcstring.New("JohnDoe", 16),
		mculong.New(0),
		mculong.New(0),
	))

	serverResponse := mcpacket.New(0x00)

	m.ReadPacket(serverResponse)

	if serverResponse.Protocol.Value == 0x00 {
		return servertype.Unknown, errors.New("Failed to connect")
	}

	if serverResponse.Protocol.Value == 0x01 {
		return servertype.Legit, nil
	}

	return servertype.Cracked, nil
}

func (m *McConnection) ReadPacket(packet *mcpacket.McPacket) error {
	m.connection.SetDeadline(time.Now().Add(consts.TimeoutTime))
	if err := packet.FromStream(m.reader); err != nil {
		return err
	}

	return nil
}

func (m *McConnection) SendPacket(packet *mcpacket.McPacket) error {
	m.connection.SetDeadline(time.Now().Add(consts.TimeoutTime))
	if err := packet.ToStream(m.writer); err != nil {
		return err
	}
	m.writer.Flush()

	return nil
}

func (m *McConnection) Close() {
	m.connection.Close()
}
