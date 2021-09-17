// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package messagesv1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	io "io"
)

func (x *GetAgentVersionRequest) GetMessageType() MessageType {
	return MessageType_GET_AGENT_VERSION_REQUEST
}

func (x *GetAgentVersionRequest) Parse(reader io.Reader, version int32) error {
	return nil
}
func (x *GetAgentVersionRequest) Formatter(writer io.Writer, version int32) error {
	return nil
}
func (x *GetAgentVersionResponse) GetMessageType() MessageType {
	return MessageType_GET_AGENT_VERSION_RESPONSE
}

func (x *GetAgentVersionResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Version opts: order:1
	if err := codec256.ParseString(reader, &x.Version); err != nil {
		return err
	}
	return nil
}
func (x *GetAgentVersionResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Version opts: order:1
	if err := codec256.FormatString(writer, x.Version); err != nil {
		return err
	}
	return nil
}
func (x *GetClusterAdminsRequest) GetMessageType() MessageType {
	return MessageType_GET_AGENT_VERSION_REQUEST
}

func (x *GetClusterAdminsRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetClusterAdminsRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetClusterAdminsResponse) GetMessageType() MessageType {
	return MessageType_GET_CLUSTER_ADMINS_RESPONSE
}

func (x *GetClusterAdminsResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Admins opts: order:1
	var size_Admins int
	if err := codec256.ParseSize(reader, &size_Admins); err != nil {
		return err
	}
	for i := 0; i < size_Admins; i++ {
		val := &AdminInfo{}
		if err := val.Parse(reader, version); err != nil {
			return err
		}

		x.Admins = append(x.Admins, val)
	}
	return nil
}
func (x *GetClusterAdminsResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Admins opts: order:1
	if err := codec256.FormatSize(writer, len(x.Admins)); err != nil {
		return err
	}
	for i := 0; i < len(x.Admins); i++ {
		if err := x.Admins[i].Formatter(writer, version); err != nil {
			return err
		}
	}
	return nil
}
func (x *GetAgentAdminsRequest) GetMessageType() MessageType {
	return MessageType_GET_AGENT_ADMINS_REQUEST
}

func (x *GetAgentAdminsRequest) Parse(reader io.Reader, version int32) error {
	return nil
}
func (x *GetAgentAdminsRequest) Formatter(writer io.Writer, version int32) error {
	return nil
}
func (x *GetAgentAdminsResponse) GetMessageType() MessageType {
	return MessageType_GET_AGENT_ADMINS_RESPONSE
}

func (x *GetAgentAdminsResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Admins opts: order:1
	var size_Admins int
	if err := codec256.ParseSize(reader, &size_Admins); err != nil {
		return err
	}
	for i := 0; i < size_Admins; i++ {
		val := &AdminInfo{}
		if err := val.Parse(reader, version); err != nil {
			return err
		}

		x.Admins = append(x.Admins, val)
	}
	return nil
}
func (x *GetAgentAdminsResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Admins opts: order:1
	if err := codec256.FormatSize(writer, len(x.Admins)); err != nil {
		return err
	}
	for i := 0; i < len(x.Admins); i++ {
		if err := x.Admins[i].Formatter(writer, version); err != nil {
			return err
		}
	}
	return nil
}
func (x *CreateClusterAdminRequest) GetMessageType() MessageType {
	return MessageType_REG_CLUSTER_ADMIN_REQUEST
}

func (x *CreateClusterAdminRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.AdminInfo opts: order:1
	x.AdminInfo = &AdminInfo{}
	if err := x.AdminInfo.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *CreateClusterAdminRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.AdminInfo opts: order:1
	if err := x.AdminInfo.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *DeleteClusterAdminRequest) GetMessageType() MessageType {
	return MessageType_UNREG_CLUSTER_ADMIN_REQUEST
}

func (x *DeleteClusterAdminRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.AdminName opts: order:1
	if err := codec256.ParseString(reader, &x.AdminName); err != nil {
		return err
	}
	return nil
}
func (x *DeleteClusterAdminRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.AdminName opts: order:1
	if err := codec256.FormatString(writer, x.AdminName); err != nil {
		return err
	}
	return nil
}
func (x *CreateAgentAdminRequest) GetMessageType() MessageType {
	return MessageType_REG_AGENT_ADMIN_REQUEST
}

func (x *CreateAgentAdminRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.AdminInfo opts: order:1
	x.AdminInfo = &AdminInfo{}
	if err := x.AdminInfo.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *CreateAgentAdminRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.AdminInfo opts: order:1
	if err := x.AdminInfo.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *DeleteAgentAdminRequest) GetMessageType() MessageType {
	return MessageType_UNREG_AGENT_ADMIN_REQUEST
}

func (x *DeleteAgentAdminRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.AdminName opts: order:1
	if err := codec256.ParseString(reader, &x.AdminName); err != nil {
		return err
	}
	return nil
}
func (x *DeleteAgentAdminRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.AdminName opts: order:1
	if err := codec256.FormatString(writer, x.AdminName); err != nil {
		return err
	}
	return nil
}
func (x *AdminInfo) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Name opts: order:1
	if err := codec256.ParseString(reader, &x.Name); err != nil {
		return err
	}
	// decode x.Description opts: order:2
	if err := codec256.ParseString(reader, &x.Description); err != nil {
		return err
	}
	// decode x.Password opts: order:3
	if err := codec256.ParseString(reader, &x.Password); err != nil {
		return err
	}
	// decode x.PasswordAuthAllowed opts: order:4
	if err := codec256.ParseBool(reader, &x.PasswordAuthAllowed); err != nil {
		return err
	}
	// decode x.SysAuthAllowed opts: order:5
	if err := codec256.ParseBool(reader, &x.SysAuthAllowed); err != nil {
		return err
	}
	// decode x.SysUserName opts: order:6
	if err := codec256.ParseString(reader, &x.SysUserName); err != nil {
		return err
	}
	return nil
}
func (x *AdminInfo) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Name opts: order:1
	if err := codec256.FormatString(writer, x.Name); err != nil {
		return err
	}
	// decode x.Description opts: order:2
	if err := codec256.FormatString(writer, x.Description); err != nil {
		return err
	}
	// decode x.Password opts: order:3
	if err := codec256.FormatString(writer, x.Password); err != nil {
		return err
	}
	// decode x.PasswordAuthAllowed opts: order:4
	if err := codec256.FormatBool(writer, x.PasswordAuthAllowed); err != nil {
		return err
	}
	// decode x.SysAuthAllowed opts: order:5
	if err := codec256.FormatBool(writer, x.SysAuthAllowed); err != nil {
		return err
	}
	// decode x.SysUserName opts: order:6
	if err := codec256.FormatString(writer, x.SysUserName); err != nil {
		return err
	}
	return nil
}