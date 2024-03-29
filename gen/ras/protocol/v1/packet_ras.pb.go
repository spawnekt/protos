// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package protocolv1

import (
	bytes "bytes"
	fmt "fmt"
	codec256 "github.com/v8platform/encoder/ras/codec256"
	io "io"
)

// Helpers generated by github.com/v8platform/protoc-gen-go-ras. DO NOT EDIT
type PacketMessage interface {
	PacketMessageFormatter
	PacketMessageParser
}

type PacketMessageFormatter interface {
	GetPacketType() PacketType
	Formatter(writer io.Writer, version int32) error
}

type PacketMessageParser interface {
	GetPacketType() PacketType
	Parse(reader io.Reader, version int32) error
}

func NewPacket(message interface{}) (*Packet, error) {
	var packet Packet
	switch typed := message.(type) {
	case io.Reader:
		if err := packet.Parse(typed, 0); err != nil {
			return nil, err
		}
	case PacketMessageFormatter:
		buf := &bytes.Buffer{}
		if err := typed.Formatter(buf, 0); err != nil {
			return nil, err
		}
		packet.Type = typed.GetPacketType()
		packet.Data = buf.Bytes()
		packet.Size = int32(len(packet.Data))
	default:
		return nil, fmt.Errorf("unknown type <%T> to get new packet", typed)
	}
	return &packet, nil
}

func (x *Packet) Unpack(into PacketMessageParser) error {
	switch x.GetType() {
	case into.GetPacketType():
		buf := bytes.NewBuffer(x.Data)
		return into.Parse(buf, 0)
	default:
		if _, err := x.UnpackNew(); err != nil {
			return err
		}
		return fmt.Errorf("unpack type no equal packet type. Has %s want %s", x.GetType(), into.GetPacketType())
	}
}

func (x *Packet) UnpackNew() (interface{}, error) {
	var into interface{}
	switch x.GetType() {
	// type PacketType_PACKET_TYPE_CONNECT_ACK cast ConnectMessageAck
	case PacketType_PACKET_TYPE_CONNECT_ACK:
		into = &ConnectMessageAck{}
	// type PacketType_PACKET_TYPE_ENDPOINT_OPEN cast EndpointOpen
	case PacketType_PACKET_TYPE_ENDPOINT_OPEN:
		into = &EndpointOpen{}
	// type PacketType_PACKET_TYPE_NEGOTIATE cast NegotiateMessage
	case PacketType_PACKET_TYPE_NEGOTIATE:
		into = &NegotiateMessage{}
	// type PacketType_PACKET_TYPE_ENDPOINT_MESSAGE cast EndpointMessage
	case PacketType_PACKET_TYPE_ENDPOINT_MESSAGE:
		into = &EndpointMessage{}
	// type PacketType_PACKET_TYPE_ENDPOINT_OPEN_ACK cast EndpointOpenAck
	case PacketType_PACKET_TYPE_ENDPOINT_OPEN_ACK:
		into = &EndpointOpenAck{}
	// type PacketType_PACKET_TYPE_DISCONNECT cast DisconnectMessage
	case PacketType_PACKET_TYPE_DISCONNECT:
		into = &DisconnectMessage{}
	// type PacketType_PACKET_TYPE_ENDPOINT_FAILURE cast EndpointFailureAck
	case PacketType_PACKET_TYPE_ENDPOINT_FAILURE:
		into = &EndpointFailureAck{}
	// type PacketType_PACKET_TYPE_CONNECT cast ConnectMessage
	case PacketType_PACKET_TYPE_CONNECT:
		into = &ConnectMessage{}
	default:
		return nil, fmt.Errorf("unknown unpack type %s", x.GetType())
	}
	buf := bytes.NewBuffer(x.Data)
	parser := into.(PacketMessageParser)
	if err := parser.Parse(buf, 0); err != nil {
		return nil, err
	}
	if err, ok := into.(error); ok {
		return nil, err
	}
	return into, nil
}

func (x *Packet) WriteTo(w io.Writer) (int64, error) {
	buf := &bytes.Buffer{}
	if err := x.Formatter(buf, 0); err != nil {
		return 0, err
	}
	n, err := w.Write(buf.Bytes())
	return int64(n), err
}
func (x *Packet) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Type opts: encoder:"byte"  order:1
	var val_Type int32
	if err := codec256.ParseByte(reader, &val_Type); err != nil {
		return err
	}
	x.Type = PacketType(val_Type)
	// decode x.Size opts: encoder:"size"  order:2
	if err := codec256.ParseSize(reader, &x.Size); err != nil {
		return err
	}
	// decode x.Data opts: encoder:"bytes"  order:3  type_field:1  size_field:2
	x.Data = make([]byte, x.GetSize())
	if err := codec256.ParseBytes(reader, x.Data); err != nil {
		return err
	}
	return nil
}
func (x *Packet) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Type opts: encoder:"byte"  order:1
	if err := codec256.FormatByte(writer, int32(x.Type)); err != nil {
		return err
	}
	// decode x.Size opts: encoder:"size"  order:2
	if err := codec256.FormatSize(writer, x.Size); err != nil {
		return err
	}
	// decode x.Data opts: encoder:"bytes"  order:3  type_field:1  size_field:2
	if err := codec256.FormatBytes(writer, x.Data); err != nil {
		return err
	}
	return nil
}
