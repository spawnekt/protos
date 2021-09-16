// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package serializev1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
)

func (x *LockInfo) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ConnectionId opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.ConnectionId); err != nil {
		return err
	}
	// decode x.Description opts: order:2
	if err := codec256.ParseString(reader, &x.Description); err != nil {
		return err
	}
	// decode x.LockedAt opts: encoder:"time"  order:3
	x.LockedAt = &timestamppb.Timestamp{}
	if err := codec256.ParseTime(reader, x.LockedAt); err != nil {
		return err
	}
	// decode x.ObjectId opts: encoder:"uuid"  order:4
	if err := codec256.ParseUUID(reader, &x.ObjectId); err != nil {
		return err
	}
	// decode x.SessionId opts: encoder:"uuid"  order:5
	if err := codec256.ParseUUID(reader, &x.SessionId); err != nil {
		return err
	}
	return nil
}
func (x *LockInfo) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ConnectionId opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.ConnectionId); err != nil {
		return err
	}
	// decode x.Description opts: order:2
	if err := codec256.FormatString(writer, x.Description); err != nil {
		return err
	}
	// decode x.LockedAt opts: encoder:"time"  order:3
	// TODO check nil
	if err := codec256.FormatTime(writer, x.GetLockedAt().AsTime()); err != nil {
		return err
	}
	// decode x.ObjectId opts: encoder:"uuid"  order:4
	if err := codec256.FormatUuid(writer, x.ObjectId); err != nil {
		return err
	}
	// decode x.SessionId opts: encoder:"uuid"  order:5
	if err := codec256.FormatUuid(writer, x.SessionId); err != nil {
		return err
	}
	return nil
}