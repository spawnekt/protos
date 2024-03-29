// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package messagesv1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	v1 "github.com/v8platform/protos/gen/v8platform/serialize/v1"
	io "io"
)

func (x *GetClustersRequest) GetMessageType() MessageType {
	return MessageType_GET_CLUSTERS_REQUEST
}

func (x *GetClustersRequest) Parse(reader io.Reader, version int32) error {
	return nil
}
func (x *GetClustersRequest) Formatter(writer io.Writer, version int32) error {
	return nil
}
func (x *GetClustersResponse) GetMessageType() MessageType {
	return MessageType_GET_CLUSTERS_RESPONSE
}

func (x *GetClustersResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Clusters opts: order:1
	var size_Clusters int
	if err := codec256.ParseSize(reader, &size_Clusters); err != nil {
		return err
	}
	for i := 0; i < size_Clusters; i++ {
		val := &v1.ClusterInfo{}
		if err := val.Parse(reader, version); err != nil {
			return err
		}

		x.Clusters = append(x.Clusters, val)
	}
	return nil
}
func (x *GetClustersResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Clusters opts: order:1
	if err := codec256.FormatSize(writer, len(x.Clusters)); err != nil {
		return err
	}
	for i := 0; i < len(x.Clusters); i++ {
		if err := x.Clusters[i].Formatter(writer, version); err != nil {
			return err
		}
	}
	return nil
}
func (x *GetClusterInfoRequest) GetMessageType() MessageType {
	return MessageType_GET_CLUSTER_INFO_REQUEST
}

func (x *GetClusterInfoRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetClusterInfoRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetClusterInfoResponse) GetMessageType() MessageType {
	return MessageType_GET_CLUSTER_INFO_RESPONSE
}

func (x *GetClusterInfoResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterInfo opts: order:1
	x.ClusterInfo = &v1.ClusterInfo{}
	if err := x.ClusterInfo.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *GetClusterInfoResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterInfo opts: order:1
	if err := x.ClusterInfo.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *RegClusterRequest) GetMessageType() MessageType {
	return MessageType_REG_CLUSTER_REQUEST
}

func (x *RegClusterRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterInfo opts: order:1
	x.ClusterInfo = &v1.ClusterInfo{}
	if err := x.ClusterInfo.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *RegClusterRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterInfo opts: order:1
	if err := x.ClusterInfo.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *RegClusterResponse) GetMessageType() MessageType {
	return MessageType_REG_CLUSTER_RESPONSE
}

func (x *RegClusterResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *RegClusterResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *UnregClusterRequest) GetMessageType() MessageType {
	return MessageType_UNREG_CLUSTER_REQUEST
}

func (x *UnregClusterRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *UnregClusterRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	return nil
}
