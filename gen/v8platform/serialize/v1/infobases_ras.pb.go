// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package serializev1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
)

func (x *InfobaseInfo) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Uuid opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.Uuid); err != nil {
		return err
	}
	// decode x.DateOffset opts: order:2
	if err := codec256.ParseInt(reader, &x.DateOffset); err != nil {
		return err
	}
	// decode x.Dbms opts: order:3
	if err := codec256.ParseString(reader, &x.Dbms); err != nil {
		return err
	}
	// decode x.DbServer opts: order:4
	if err := codec256.ParseString(reader, &x.DbServer); err != nil {
		return err
	}
	// decode x.DbName opts: order:5
	if err := codec256.ParseString(reader, &x.DbName); err != nil {
		return err
	}
	// decode x.DbUser opts: order:6
	if err := codec256.ParseString(reader, &x.DbUser); err != nil {
		return err
	}
	// decode x.DbPwd opts: order:7
	if err := codec256.ParseString(reader, &x.DbPwd); err != nil {
		return err
	}
	// decode x.DeniedFrom opts: encoder:"time"  order:8
	x.DeniedFrom = &timestamppb.Timestamp{}
	if err := codec256.ParseTime(reader, x.DeniedFrom); err != nil {
		return err
	}
	// decode x.DeniedMessage opts: order:9
	if err := codec256.ParseString(reader, &x.DeniedMessage); err != nil {
		return err
	}
	// decode x.DeniedParameter opts: order:10
	if err := codec256.ParseString(reader, &x.DeniedParameter); err != nil {
		return err
	}
	// decode x.DeniedTo opts: encoder:"time"  order:11
	x.DeniedTo = &timestamppb.Timestamp{}
	if err := codec256.ParseTime(reader, x.DeniedTo); err != nil {
		return err
	}
	// decode x.Descr opts: order:12
	if err := codec256.ParseString(reader, &x.Descr); err != nil {
		return err
	}
	// decode x.Locale opts: order:13
	if err := codec256.ParseString(reader, &x.Locale); err != nil {
		return err
	}
	// decode x.Name opts: order:14
	if err := codec256.ParseString(reader, &x.Name); err != nil {
		return err
	}
	// decode x.PermissionCode opts: order:15
	if err := codec256.ParseString(reader, &x.PermissionCode); err != nil {
		return err
	}
	// decode x.ScheduledJobsDeny opts: order:16
	if err := codec256.ParseBool(reader, &x.ScheduledJobsDeny); err != nil {
		return err
	}
	// decode x.SecurityLevel opts: order:17
	if err := codec256.ParseInt(reader, &x.SecurityLevel); err != nil {
		return err
	}
	// decode x.SessionsDeny opts: order:18
	if err := codec256.ParseBool(reader, &x.SessionsDeny); err != nil {
		return err
	}
	// decode x.LicenseDistribution opts: order:19
	if err := codec256.ParseInt(reader, &x.LicenseDistribution); err != nil {
		return err
	}
	// decode x.ExternalSessionManagerConnectionString opts: order:20
	if err := codec256.ParseString(reader, &x.ExternalSessionManagerConnectionString); err != nil {
		return err
	}
	// decode x.ExternalSessionManagerRequired opts: order:21
	if err := codec256.ParseBool(reader, &x.ExternalSessionManagerRequired); err != nil {
		return err
	}
	// decode x.SecurityProfileName opts: order:22
	if err := codec256.ParseString(reader, &x.SecurityProfileName); err != nil {
		return err
	}
	// decode x.SafeModeSecurityProfileName opts: order:23
	if err := codec256.ParseString(reader, &x.SafeModeSecurityProfileName); err != nil {
		return err
	}
	if version >= 9 {
		// decode x.ReserveWorkingProcesses opts: order:24  version:9
		if err := codec256.ParseBool(reader, &x.ReserveWorkingProcesses); err != nil {
			return err
		}
	}
	return nil
}
func (x *InfobaseInfo) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Uuid opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.Uuid); err != nil {
		return err
	}
	// decode x.DateOffset opts: order:2
	if err := codec256.FormatInt(writer, x.DateOffset); err != nil {
		return err
	}
	// decode x.Dbms opts: order:3
	if err := codec256.FormatString(writer, x.Dbms); err != nil {
		return err
	}
	// decode x.DbServer opts: order:4
	if err := codec256.FormatString(writer, x.DbServer); err != nil {
		return err
	}
	// decode x.DbName opts: order:5
	if err := codec256.FormatString(writer, x.DbName); err != nil {
		return err
	}
	// decode x.DbUser opts: order:6
	if err := codec256.FormatString(writer, x.DbUser); err != nil {
		return err
	}
	// decode x.DbPwd opts: order:7
	if err := codec256.FormatString(writer, x.DbPwd); err != nil {
		return err
	}
	// decode x.DeniedFrom opts: encoder:"time"  order:8
	// TODO check nil
	if err := codec256.FormatTime(writer, x.GetDeniedFrom().AsTime()); err != nil {
		return err
	}
	// decode x.DeniedMessage opts: order:9
	if err := codec256.FormatString(writer, x.DeniedMessage); err != nil {
		return err
	}
	// decode x.DeniedParameter opts: order:10
	if err := codec256.FormatString(writer, x.DeniedParameter); err != nil {
		return err
	}
	// decode x.DeniedTo opts: encoder:"time"  order:11
	// TODO check nil
	if err := codec256.FormatTime(writer, x.GetDeniedTo().AsTime()); err != nil {
		return err
	}
	// decode x.Descr opts: order:12
	if err := codec256.FormatString(writer, x.Descr); err != nil {
		return err
	}
	// decode x.Locale opts: order:13
	if err := codec256.FormatString(writer, x.Locale); err != nil {
		return err
	}
	// decode x.Name opts: order:14
	if err := codec256.FormatString(writer, x.Name); err != nil {
		return err
	}
	// decode x.PermissionCode opts: order:15
	if err := codec256.FormatString(writer, x.PermissionCode); err != nil {
		return err
	}
	// decode x.ScheduledJobsDeny opts: order:16
	if err := codec256.FormatBool(writer, x.ScheduledJobsDeny); err != nil {
		return err
	}
	// decode x.SecurityLevel opts: order:17
	if err := codec256.FormatInt(writer, x.SecurityLevel); err != nil {
		return err
	}
	// decode x.SessionsDeny opts: order:18
	if err := codec256.FormatBool(writer, x.SessionsDeny); err != nil {
		return err
	}
	// decode x.LicenseDistribution opts: order:19
	if err := codec256.FormatInt(writer, x.LicenseDistribution); err != nil {
		return err
	}
	// decode x.ExternalSessionManagerConnectionString opts: order:20
	if err := codec256.FormatString(writer, x.ExternalSessionManagerConnectionString); err != nil {
		return err
	}
	// decode x.ExternalSessionManagerRequired opts: order:21
	if err := codec256.FormatBool(writer, x.ExternalSessionManagerRequired); err != nil {
		return err
	}
	// decode x.SecurityProfileName opts: order:22
	if err := codec256.FormatString(writer, x.SecurityProfileName); err != nil {
		return err
	}
	// decode x.SafeModeSecurityProfileName opts: order:23
	if err := codec256.FormatString(writer, x.SafeModeSecurityProfileName); err != nil {
		return err
	}
	if version >= 9 {
		// decode x.ReserveWorkingProcesses opts: order:24  version:9
		if err := codec256.FormatBool(writer, x.ReserveWorkingProcesses); err != nil {
			return err
		}
	}
	return nil
}
func (x *InfobaseSummaryInfo) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Uuid opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.Uuid); err != nil {
		return err
	}
	// decode x.Descr opts: order:2
	if err := codec256.ParseString(reader, &x.Descr); err != nil {
		return err
	}
	// decode x.Name opts: order:3
	if err := codec256.ParseString(reader, &x.Name); err != nil {
		return err
	}
	return nil
}
func (x *InfobaseSummaryInfo) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Uuid opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.Uuid); err != nil {
		return err
	}
	// decode x.Descr opts: order:2
	if err := codec256.FormatString(writer, x.Descr); err != nil {
		return err
	}
	// decode x.Name opts: order:3
	if err := codec256.FormatString(writer, x.Name); err != nil {
		return err
	}
	return nil
}
