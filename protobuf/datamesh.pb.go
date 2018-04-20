// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/datamesh.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	protobuf/datamesh.proto

It has these top-level messages:
	Transaction
	Operation
	ReadOperation
	UpsertOperation
	DeleteOperation
	GenerateEventOperation
	ApplicationFailure
	Event
	ReadRequest
	Path
	Data
	Status
	Context
	Empty
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A transaction is a sequence of operation triggered by a event.
type Transaction struct {
	// The trigger
	Event *Event `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	// The projection context
	Context *Context `protobuf:"bytes,2,opt,name=context" json:"context,omitempty"`
	// The operations
	Operations []*Operation `protobuf:"bytes,3,rep,name=operations" json:"operations,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Transaction) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Transaction) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Transaction) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

type Operation struct {
	// Types that are valid to be assigned to Kind:
	//	*Operation_Read
	//	*Operation_Upsert
	//	*Operation_Delete
	//	*Operation_Generate
	//	*Operation_Failure
	Kind isOperation_Kind `protobuf_oneof:"kind"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isOperation_Kind interface {
	isOperation_Kind()
}

type Operation_Read struct {
	Read *ReadOperation `protobuf:"bytes,1,opt,name=read,oneof"`
}
type Operation_Upsert struct {
	Upsert *UpsertOperation `protobuf:"bytes,2,opt,name=upsert,oneof"`
}
type Operation_Delete struct {
	Delete *DeleteOperation `protobuf:"bytes,3,opt,name=delete,oneof"`
}
type Operation_Generate struct {
	Generate *GenerateEventOperation `protobuf:"bytes,4,opt,name=generate,oneof"`
}
type Operation_Failure struct {
	Failure *ApplicationFailure `protobuf:"bytes,5,opt,name=failure,oneof"`
}

func (*Operation_Read) isOperation_Kind()     {}
func (*Operation_Upsert) isOperation_Kind()   {}
func (*Operation_Delete) isOperation_Kind()   {}
func (*Operation_Generate) isOperation_Kind() {}
func (*Operation_Failure) isOperation_Kind()  {}

func (m *Operation) GetKind() isOperation_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Operation) GetRead() *ReadOperation {
	if x, ok := m.GetKind().(*Operation_Read); ok {
		return x.Read
	}
	return nil
}

func (m *Operation) GetUpsert() *UpsertOperation {
	if x, ok := m.GetKind().(*Operation_Upsert); ok {
		return x.Upsert
	}
	return nil
}

func (m *Operation) GetDelete() *DeleteOperation {
	if x, ok := m.GetKind().(*Operation_Delete); ok {
		return x.Delete
	}
	return nil
}

func (m *Operation) GetGenerate() *GenerateEventOperation {
	if x, ok := m.GetKind().(*Operation_Generate); ok {
		return x.Generate
	}
	return nil
}

func (m *Operation) GetFailure() *ApplicationFailure {
	if x, ok := m.GetKind().(*Operation_Failure); ok {
		return x.Failure
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_Read)(nil),
		(*Operation_Upsert)(nil),
		(*Operation_Delete)(nil),
		(*Operation_Generate)(nil),
		(*Operation_Failure)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// kind
	switch x := m.Kind.(type) {
	case *Operation_Read:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Read); err != nil {
			return err
		}
	case *Operation_Upsert:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Upsert); err != nil {
			return err
		}
	case *Operation_Delete:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Delete); err != nil {
			return err
		}
	case *Operation_Generate:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Generate); err != nil {
			return err
		}
	case *Operation_Failure:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Failure); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Kind has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 1: // kind.read
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ReadOperation)
		err := b.DecodeMessage(msg)
		m.Kind = &Operation_Read{msg}
		return true, err
	case 2: // kind.upsert
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpsertOperation)
		err := b.DecodeMessage(msg)
		m.Kind = &Operation_Upsert{msg}
		return true, err
	case 3: // kind.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeleteOperation)
		err := b.DecodeMessage(msg)
		m.Kind = &Operation_Delete{msg}
		return true, err
	case 4: // kind.generate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GenerateEventOperation)
		err := b.DecodeMessage(msg)
		m.Kind = &Operation_Generate{msg}
		return true, err
	case 5: // kind.failure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApplicationFailure)
		err := b.DecodeMessage(msg)
		m.Kind = &Operation_Failure{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// kind
	switch x := m.Kind.(type) {
	case *Operation_Read:
		s := proto.Size(x.Read)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Upsert:
		s := proto.Size(x.Upsert)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Delete:
		s := proto.Size(x.Delete)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Generate:
		s := proto.Size(x.Generate)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Failure:
		s := proto.Size(x.Failure)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ReadOperation struct {
	Path *Path `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *ReadOperation) Reset()                    { *m = ReadOperation{} }
func (m *ReadOperation) String() string            { return proto.CompactTextString(m) }
func (*ReadOperation) ProtoMessage()               {}
func (*ReadOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadOperation) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

type UpsertOperation struct {
	Data *Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *UpsertOperation) Reset()                    { *m = UpsertOperation{} }
func (m *UpsertOperation) String() string            { return proto.CompactTextString(m) }
func (*UpsertOperation) ProtoMessage()               {}
func (*UpsertOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpsertOperation) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteOperation struct {
	Path *Path `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *DeleteOperation) Reset()                    { *m = DeleteOperation{} }
func (m *DeleteOperation) String() string            { return proto.CompactTextString(m) }
func (*DeleteOperation) ProtoMessage()               {}
func (*DeleteOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteOperation) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

type GenerateEventOperation struct {
	Event *Event `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
}

func (m *GenerateEventOperation) Reset()                    { *m = GenerateEventOperation{} }
func (m *GenerateEventOperation) String() string            { return proto.CompactTextString(m) }
func (*GenerateEventOperation) ProtoMessage()               {}
func (*GenerateEventOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GenerateEventOperation) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

// To be used automatically in case of runtime errors (usually bugs)
type ApplicationFailure struct {
	Reason string `protobuf:"bytes,1,opt,name=reason" json:"reason,omitempty"`
}

func (m *ApplicationFailure) Reset()                    { *m = ApplicationFailure{} }
func (m *ApplicationFailure) String() string            { return proto.CompactTextString(m) }
func (*ApplicationFailure) ProtoMessage()               {}
func (*ApplicationFailure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ApplicationFailure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

// A Event object may model a command (action to executed) or a proper event (action happened in the past)
type Event struct {
	Group   string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// Client identifier is used to match the logged Event with a Transaction in case of fast-path processing
	ClientIdentifier string `protobuf:"bytes,4,opt,name=client_identifier,json=clientIdentifier" json:"client_identifier,omitempty"`
	// Client version should be made visible to client API
	ClientVersion string `protobuf:"bytes,5,opt,name=client_version,json=clientVersion" json:"client_version,omitempty"`
	// Version is meaningful only when event is stored (0 before)
	Version uint64 `protobuf:"varint,6,opt,name=version" json:"version,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Event) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetClientIdentifier() string {
	if m != nil {
		return m.ClientIdentifier
	}
	return ""
}

func (m *Event) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *Event) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

// A read request
type ReadRequest struct {
	Context *Context `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Path    *Path    `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReadRequest) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *ReadRequest) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

// A path in the projection store
type Path struct {
	Location string `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Version  uint64 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *Path) Reset()                    { *m = Path{} }
func (m *Path) String() string            { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()               {}
func (*Path) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Path) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Path) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

// The object contained in a specific Path
type Data struct {
	Path *Path `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// JSON encoded content
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Data) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Data) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// The status of a connected client
type Status struct {
	// Types that are valid to be assigned to Status:
	//	*Status_Connect
	//	*Status_Disconnect
	Status isStatus_Status `protobuf_oneof:"status"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type isStatus_Status interface {
	isStatus_Status()
}

type Status_Connect struct {
	Connect *Context `protobuf:"bytes,1,opt,name=connect,oneof"`
}
type Status_Disconnect struct {
	Disconnect *Empty `protobuf:"bytes,2,opt,name=disconnect,oneof"`
}

func (*Status_Connect) isStatus_Status()    {}
func (*Status_Disconnect) isStatus_Status() {}

func (m *Status) GetStatus() isStatus_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Status) GetConnect() *Context {
	if x, ok := m.GetStatus().(*Status_Connect); ok {
		return x.Connect
	}
	return nil
}

func (m *Status) GetDisconnect() *Empty {
	if x, ok := m.GetStatus().(*Status_Disconnect); ok {
		return x.Disconnect
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Status) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Status_OneofMarshaler, _Status_OneofUnmarshaler, _Status_OneofSizer, []interface{}{
		(*Status_Connect)(nil),
		(*Status_Disconnect)(nil),
	}
}

func _Status_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Status)
	// status
	switch x := m.Status.(type) {
	case *Status_Connect:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Connect); err != nil {
			return err
		}
	case *Status_Disconnect:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Disconnect); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Status.Status has unexpected type %T", x)
	}
	return nil
}

func _Status_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Status)
	switch tag {
	case 1: // status.connect
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Context)
		err := b.DecodeMessage(msg)
		m.Status = &Status_Connect{msg}
		return true, err
	case 2: // status.disconnect
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Empty)
		err := b.DecodeMessage(msg)
		m.Status = &Status_Disconnect{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Status_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Status)
	// status
	switch x := m.Status.(type) {
	case *Status_Connect:
		s := proto.Size(x.Connect)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Status_Disconnect:
		s := proto.Size(x.Disconnect)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A context representing the projection to which a client is connected
type Context struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Revision uint64 `protobuf:"varint,2,opt,name=revision" json:"revision,omitempty"`
}

func (m *Context) Reset()                    { *m = Context{} }
func (m *Context) String() string            { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()               {}
func (*Context) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Context) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Context) GetRevision() uint64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

// A placeholder for empty message
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*Transaction)(nil), "protobuf.Transaction")
	proto.RegisterType((*Operation)(nil), "protobuf.Operation")
	proto.RegisterType((*ReadOperation)(nil), "protobuf.ReadOperation")
	proto.RegisterType((*UpsertOperation)(nil), "protobuf.UpsertOperation")
	proto.RegisterType((*DeleteOperation)(nil), "protobuf.DeleteOperation")
	proto.RegisterType((*GenerateEventOperation)(nil), "protobuf.GenerateEventOperation")
	proto.RegisterType((*ApplicationFailure)(nil), "protobuf.ApplicationFailure")
	proto.RegisterType((*Event)(nil), "protobuf.Event")
	proto.RegisterType((*ReadRequest)(nil), "protobuf.ReadRequest")
	proto.RegisterType((*Path)(nil), "protobuf.Path")
	proto.RegisterType((*Data)(nil), "protobuf.Data")
	proto.RegisterType((*Status)(nil), "protobuf.Status")
	proto.RegisterType((*Context)(nil), "protobuf.Context")
	proto.RegisterType((*Empty)(nil), "protobuf.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DataMesh service

type DataMeshClient interface {
	// Used to push a event that will be stored on the event log.
	Push(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error)
	// Allows to pass a transaction related to a event before it's explicitly requested.
	Process(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error)
	// The server sends events that need to be processed. The client will reply with the corresponding transactions asynchronously.
	Connect(ctx context.Context, opts ...grpc.CallOption) (DataMesh_ConnectClient, error)
	// Used by the client to query the projections.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*Data, error)
}

type dataMeshClient struct {
	cc *grpc.ClientConn
}

func NewDataMeshClient(cc *grpc.ClientConn) DataMeshClient {
	return &dataMeshClient{cc}
}

func (c *dataMeshClient) Push(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/protobuf.DataMesh/Push", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMeshClient) Process(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/protobuf.DataMesh/Process", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataMeshClient) Connect(ctx context.Context, opts ...grpc.CallOption) (DataMesh_ConnectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DataMesh_serviceDesc.Streams[0], c.cc, "/protobuf.DataMesh/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataMeshConnectClient{stream}
	return x, nil
}

type DataMesh_ConnectClient interface {
	Send(*Status) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type dataMeshConnectClient struct {
	grpc.ClientStream
}

func (x *dataMeshConnectClient) Send(m *Status) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataMeshConnectClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataMeshClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := grpc.Invoke(ctx, "/protobuf.DataMesh/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataMesh service

type DataMeshServer interface {
	// Used to push a event that will be stored on the event log.
	Push(context.Context, *Event) (*Empty, error)
	// Allows to pass a transaction related to a event before it's explicitly requested.
	Process(context.Context, *Transaction) (*Empty, error)
	// The server sends events that need to be processed. The client will reply with the corresponding transactions asynchronously.
	Connect(DataMesh_ConnectServer) error
	// Used by the client to query the projections.
	Read(context.Context, *ReadRequest) (*Data, error)
}

func RegisterDataMeshServer(s *grpc.Server, srv DataMeshServer) {
	s.RegisterService(&_DataMesh_serviceDesc, srv)
}

func _DataMesh_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMeshServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.DataMesh/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMeshServer).Push(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataMesh_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMeshServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.DataMesh/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMeshServer).Process(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataMesh_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataMeshServer).Connect(&dataMeshConnectServer{stream})
}

type DataMesh_ConnectServer interface {
	Send(*Event) error
	Recv() (*Status, error)
	grpc.ServerStream
}

type dataMeshConnectServer struct {
	grpc.ServerStream
}

func (x *dataMeshConnectServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataMeshConnectServer) Recv() (*Status, error) {
	m := new(Status)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataMesh_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataMeshServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.DataMesh/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataMeshServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataMesh_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.DataMesh",
	HandlerType: (*DataMeshServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _DataMesh_Push_Handler,
		},
		{
			MethodName: "Process",
			Handler:    _DataMesh_Process_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _DataMesh_Read_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _DataMesh_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/datamesh.proto",
}

func init() { proto.RegisterFile("protobuf/datamesh.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdb, 0x6e, 0x13, 0x3d,
	0x10, 0xce, 0xa6, 0x9b, 0xd3, 0xa4, 0x47, 0xff, 0x3f, 0x6d, 0x88, 0xb8, 0x88, 0x2c, 0x2a, 0x45,
	0x94, 0xa6, 0x90, 0x08, 0x09, 0x24, 0x04, 0xa2, 0x14, 0x28, 0x17, 0x88, 0xca, 0x1c, 0x2e, 0x41,
	0x6e, 0x76, 0xd2, 0xac, 0xd8, 0x78, 0x17, 0xdb, 0x5b, 0xd1, 0x17, 0xe1, 0x11, 0x78, 0x07, 0xde,
	0x82, 0x47, 0x42, 0xf6, 0x1e, 0xb3, 0x2d, 0xd0, 0xbb, 0x9d, 0xf9, 0xbe, 0xcf, 0x33, 0x9e, 0xf9,
	0x9c, 0xc0, 0x4e, 0x24, 0x43, 0x1d, 0x9e, 0xc6, 0xb3, 0x03, 0x8f, 0x6b, 0xbe, 0x40, 0x35, 0x1f,
	0xd9, 0x0c, 0x69, 0x67, 0x00, 0xfd, 0xee, 0x40, 0xf7, 0xbd, 0xe4, 0x42, 0xf1, 0xa9, 0xf6, 0x43,
	0x41, 0x76, 0xa1, 0x81, 0xe7, 0x28, 0x74, 0xcf, 0x19, 0x38, 0xc3, 0xee, 0x78, 0x63, 0x94, 0x31,
	0x47, 0x2f, 0x4c, 0x9a, 0x25, 0x28, 0xd9, 0x83, 0xd6, 0x34, 0x14, 0x1a, 0xbf, 0xe9, 0x5e, 0xdd,
	0x12, 0xb7, 0x0a, 0xe2, 0xf3, 0x04, 0x60, 0x19, 0x83, 0x4c, 0x00, 0xc2, 0x08, 0x25, 0x37, 0x05,
	0x54, 0x6f, 0x65, 0xb0, 0x32, 0xec, 0x8e, 0xff, 0x2b, 0xf8, 0x6f, 0x33, 0x8c, 0x95, 0x68, 0xf4,
	0x47, 0x1d, 0x3a, 0x39, 0x42, 0xf6, 0xc1, 0x95, 0xc8, 0xbd, 0xb4, 0xab, 0x9d, 0x42, 0xcc, 0x90,
	0x7b, 0x39, 0xed, 0xb8, 0xc6, 0x2c, 0x8d, 0x4c, 0xa0, 0x19, 0x47, 0x0a, 0x65, 0xd6, 0xdd, 0xcd,
	0x42, 0xf0, 0xc1, 0xe6, 0xcb, 0x92, 0x94, 0x6a, 0x44, 0x1e, 0x06, 0xa8, 0xb1, 0xb7, 0x52, 0x15,
	0x1d, 0xd9, 0xfc, 0x92, 0x28, 0xa1, 0x92, 0x27, 0xd0, 0x3e, 0x43, 0x61, 0xf2, 0xd8, 0x73, 0xad,
	0x6c, 0x50, 0xc8, 0x5e, 0xa5, 0x88, 0x1d, 0x5d, 0x59, 0x9d, 0x6b, 0xc8, 0x43, 0x68, 0xcd, 0xb8,
	0x1f, 0xc4, 0x12, 0x7b, 0x0d, 0x2b, 0xbf, 0x55, 0xc8, 0x9f, 0x45, 0x51, 0xe0, 0x4f, 0xad, 0xe6,
	0x65, 0xc2, 0x39, 0xae, 0xb1, 0x8c, 0x7e, 0xd8, 0x04, 0xf7, 0x8b, 0x2f, 0x3c, 0x3a, 0x81, 0xb5,
	0xa5, 0x21, 0x10, 0x0a, 0x6e, 0xc4, 0xf5, 0x3c, 0x9d, 0xd5, 0x7a, 0x71, 0xde, 0x09, 0xd7, 0x73,
	0x66, 0x31, 0xfa, 0x00, 0x36, 0x2a, 0x83, 0x30, 0x32, 0xe3, 0x92, 0xcb, 0xb2, 0x23, 0xae, 0x39,
	0xb3, 0x98, 0x91, 0x55, 0x46, 0x71, 0xad, 0x6a, 0x4f, 0x61, 0xfb, 0xea, 0x51, 0x5c, 0xd3, 0x6e,
	0xf4, 0x2e, 0x90, 0xcb, 0xc3, 0x20, 0xdb, 0xd0, 0x94, 0xc8, 0x55, 0x28, 0xac, 0xba, 0xc3, 0xd2,
	0x88, 0xfe, 0x74, 0xa0, 0x61, 0xe5, 0xe4, 0x7f, 0x68, 0x9c, 0xc9, 0x30, 0x8e, 0x52, 0x42, 0x12,
	0x10, 0x02, 0xae, 0xe0, 0x0b, 0xb4, 0xde, 0xe8, 0x30, 0xfb, 0x4d, 0x7a, 0xd0, 0x8a, 0xf8, 0x45,
	0x10, 0x72, 0xcf, 0x6e, 0x7f, 0x95, 0x65, 0x21, 0xd9, 0x83, 0xad, 0x69, 0xe0, 0xa3, 0xd0, 0x9f,
	0x7d, 0x0f, 0x85, 0xf6, 0x67, 0x3e, 0x4a, 0xbb, 0xea, 0x0e, 0xdb, 0x4c, 0x80, 0xd7, 0x79, 0x9e,
	0xec, 0xc2, 0x7a, 0x4a, 0x3e, 0x47, 0xa9, 0xfc, 0x50, 0xd8, 0xad, 0x76, 0xd8, 0x5a, 0x92, 0xfd,
	0x98, 0x24, 0x4d, 0xb5, 0x0c, 0x6f, 0x0e, 0x9c, 0xa1, 0xcb, 0xb2, 0x90, 0x7e, 0x82, 0xae, 0xd9,
	0x26, 0xc3, 0xaf, 0x31, 0xaa, 0xa5, 0x77, 0xe6, 0xfc, 0xf3, 0x9d, 0x65, 0xab, 0xa8, 0xff, 0x65,
	0x15, 0x8f, 0xc1, 0x35, 0x11, 0xe9, 0x43, 0x3b, 0x08, 0x93, 0x71, 0xa6, 0xc3, 0xc9, 0xe3, 0x72,
	0x77, 0xf5, 0xe5, 0xee, 0x8e, 0xc0, 0x35, 0x6e, 0xb8, 0xce, 0xd2, 0xcd, 0x29, 0xb6, 0x31, 0x91,
	0x3c, 0xc2, 0x55, 0x96, 0x85, 0x54, 0x43, 0xf3, 0x9d, 0xe6, 0x3a, 0x56, 0x64, 0xdf, 0x72, 0x04,
	0x4e, 0xff, 0x7c, 0x3d, 0x63, 0xf9, 0x94, 0x43, 0xee, 0x03, 0x78, 0xbe, 0xca, 0x14, 0xf5, 0x4b,
	0x96, 0x59, 0x44, 0xfa, 0xe2, 0xb8, 0xc6, 0x4a, 0xa4, 0xc3, 0x36, 0x34, 0x95, 0xad, 0x45, 0x1f,
	0x41, 0x2b, 0x3d, 0x32, 0x37, 0x80, 0x53, 0x32, 0x40, 0x1f, 0xda, 0x12, 0xcf, 0xfd, 0xd2, 0xad,
	0xf3, 0x98, 0xb6, 0xa0, 0x61, 0xcf, 0x1e, 0xff, 0x72, 0xa0, 0x6d, 0x06, 0xf0, 0x06, 0xd5, 0x9c,
	0xdc, 0x01, 0xf7, 0x24, 0x56, 0x73, 0x52, 0x35, 0x6d, 0xbf, 0xda, 0x12, 0xad, 0x91, 0x09, 0xb4,
	0x4e, 0x64, 0x38, 0x45, 0xa5, 0xc8, 0x8d, 0x02, 0x2d, 0xfd, 0xf0, 0x5e, 0x25, 0x1a, 0xdb, 0x8e,
	0xed, 0xcd, 0x37, 0x0b, 0x34, 0x19, 0x5d, 0xbf, 0x5a, 0x95, 0xd6, 0x86, 0xce, 0x3d, 0x87, 0x1c,
	0x80, 0x6b, 0xfc, 0x53, 0xae, 0x52, 0xf2, 0x53, 0xbf, 0xf2, 0xac, 0x69, 0xed, 0xf0, 0x36, 0xd0,
	0x05, 0x8e, 0x84, 0x3f, 0x0d, 0x03, 0x3e, 0x43, 0x29, 0xb9, 0x0c, 0x47, 0xcb, 0xff, 0x16, 0xa7,
	0xf1, 0xec, 0xb4, 0x69, 0xbf, 0x26, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x31, 0x76, 0x9b, 0x89,
	0x52, 0x06, 0x00, 0x00,
}
