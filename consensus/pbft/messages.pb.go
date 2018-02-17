// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package pbft is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Message
	Request
	PrePrepare
	Prepare
	Commit
	BlockInfo
	Checkpoint
	ViewChange
	PQset
	NewView
	FetchRequestBatch
	RequestBatch
	BatchMessage
	Metadata
*/
package pbft

import "github.com/golang/protobuf/proto"
import "fmt"
import "math"
import (
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/abchain/fabric/debugger"
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


type Message struct {
	// Types that are valid to be assigned to Payload:
	//	*Message_RequestBatch
	//	*Message_PrePrepare
	//	*Message_Prepare
	//	*Message_Commit
	//	*Message_Checkpoint
	//	*Message_ViewChange
	//	*Message_NewView
	//	*Message_FetchRequestBatch
	//	*Message_ReturnRequestBatch
	Payload isMessage_Payload `protobuf_oneof:"payload"`
	PayloadType int32
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_RequestBatch struct {
	RequestBatch *RequestBatch `protobuf:"bytes,1,opt,name=request_batch,json=requestBatch,oneof"`
}
type Message_PrePrepare struct {
	PrePrepare *PrePrepare `protobuf:"bytes,2,opt,name=pre_prepare,json=prePrepare,oneof"`
}
type Message_Prepare struct {
	Prepare *Prepare `protobuf:"bytes,3,opt,name=prepare,oneof"`
}
type Message_Commit struct {
	Commit *Commit `protobuf:"bytes,4,opt,name=commit,oneof"`
}
type Message_Checkpoint struct {
	Checkpoint *Checkpoint `protobuf:"bytes,5,opt,name=checkpoint,oneof"`
}
type Message_ViewChange struct {
	ViewChange *ViewChange `protobuf:"bytes,6,opt,name=view_change,json=viewChange,oneof"`
}
type Message_NewView struct {
	NewView *NewView `protobuf:"bytes,7,opt,name=new_view,json=newView,oneof"`
}
type Message_FetchRequestBatch struct {
	FetchRequestBatch *FetchRequestBatch `protobuf:"bytes,8,opt,name=fetch_request_batch,json=fetchRequestBatch,oneof"`
}
type Message_ReturnRequestBatch struct {
	ReturnRequestBatch *RequestBatch `protobuf:"bytes,9,opt,name=return_request_batch,json=returnRequestBatch,oneof"`
}

func (*Message_RequestBatch) isMessage_Payload()       {}
func (*Message_PrePrepare) isMessage_Payload()         {}
func (*Message_Prepare) isMessage_Payload()            {}
func (*Message_Commit) isMessage_Payload()             {}
func (*Message_Checkpoint) isMessage_Payload()         {}
func (*Message_ViewChange) isMessage_Payload()         {}
func (*Message_NewView) isMessage_Payload()            {}
func (*Message_FetchRequestBatch) isMessage_Payload()  {}
func (*Message_ReturnRequestBatch) isMessage_Payload() {}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRequestBatch() *RequestBatch {
	if x, ok := m.GetPayload().(*Message_RequestBatch); ok {
		return x.RequestBatch
	}
	return nil
}

func (m *Message) GetPrePrepare() *PrePrepare {
	if x, ok := m.GetPayload().(*Message_PrePrepare); ok {
		return x.PrePrepare
	}
	return nil
}

func (m *Message) GetPrepare() *Prepare {
	if x, ok := m.GetPayload().(*Message_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *Message) GetCommit() *Commit {
	if x, ok := m.GetPayload().(*Message_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Message) GetCheckpoint() *Checkpoint {
	if x, ok := m.GetPayload().(*Message_Checkpoint); ok {
		return x.Checkpoint
	}
	return nil
}

func (m *Message) GetViewChange() *ViewChange {
	if x, ok := m.GetPayload().(*Message_ViewChange); ok {
		return x.ViewChange
	}
	return nil
}

func (m *Message) GetNewView() *NewView {
	if x, ok := m.GetPayload().(*Message_NewView); ok {
		return x.NewView
	}
	return nil
}

func (m *Message) GetFetchRequestBatch() *FetchRequestBatch {
	if x, ok := m.GetPayload().(*Message_FetchRequestBatch); ok {
		return x.FetchRequestBatch
	}
	return nil
}

func (m *Message) GetReturnRequestBatch() *RequestBatch {
	if x, ok := m.GetPayload().(*Message_ReturnRequestBatch); ok {
		return x.ReturnRequestBatch
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (
	func(msg proto.Message, b *proto.Buffer) error,
	func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error),
	func(msg proto.Message) (n int),
	[]interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_RequestBatch)(nil),
		(*Message_PrePrepare)(nil),
		(*Message_Prepare)(nil),
		(*Message_Commit)(nil),
		(*Message_Checkpoint)(nil),
		(*Message_ViewChange)(nil),
		(*Message_NewView)(nil),
		(*Message_FetchRequestBatch)(nil),
		(*Message_ReturnRequestBatch)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_RequestBatch:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestBatch); err != nil {
			return err
		}
	case *Message_PrePrepare:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PrePrepare); err != nil {
			return err
		}
	case *Message_Prepare:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *Message_Commit:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Commit); err != nil {
			return err
		}
	case *Message_Checkpoint:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Checkpoint); err != nil {
			return err
		}
	case *Message_ViewChange:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ViewChange); err != nil {
			return err
		}
	case *Message_NewView:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewView); err != nil {
			return err
		}
	case *Message_FetchRequestBatch:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FetchRequestBatch); err != nil {
			return err
		}
	case *Message_ReturnRequestBatch:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReturnRequestBatch); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Payload has unexpected type %T", x)
	}

	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {

	m := msg.(*Message)

	switch tag {
	case 1: // payload.request_batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestBatch)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_RequestBatch{msg}
		m.PayloadType = int32(tag)
		return true, err
	case 2: // payload.pre_prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrePrepare)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_PrePrepare{msg}
		m.PayloadType = int32(tag)
		return true, err
	case 3: // payload.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Prepare)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Prepare{msg}
		m.PayloadType = int32(tag)
		return true, err
	case 4: // payload.commit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Commit)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Commit{msg}
		m.PayloadType = int32(tag)
		return true, err
	case 5: // payload.checkpoint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Checkpoint)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Checkpoint{msg}
		m.PayloadType = int32(tag)
		return true, err
	case 6: // payload.view_change
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ViewChange)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_ViewChange{msg}
		m.PayloadType = int32(tag)
		return true, err
	case 7: // payload.new_view
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NewView)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_NewView{msg}
		m.PayloadType = int32(tag)
		return true, err
	case 8: // payload.fetch_request_batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FetchRequestBatch)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_FetchRequestBatch{msg}
		m.PayloadType = int32(tag)
		return true, err
	case 9: // payload.return_request_batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestBatch)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_ReturnRequestBatch{msg}
		m.PayloadType = int32(tag)
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_RequestBatch:
		s := proto.Size(x.RequestBatch)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_PrePrepare:
		s := proto.Size(x.PrePrepare)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Prepare:
		s := proto.Size(x.Prepare)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Commit:
		s := proto.Size(x.Commit)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Checkpoint:
		s := proto.Size(x.Checkpoint)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_ViewChange:
		s := proto.Size(x.ViewChange)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_NewView:
		s := proto.Size(x.NewView)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_FetchRequestBatch:
		s := proto.Size(x.FetchRequestBatch)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_ReturnRequestBatch:
		s := proto.Size(x.ReturnRequestBatch)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Request struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	ReplicaId uint64                     `protobuf:"varint,3,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	Signature []byte                     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }
func (m *Request) Dump(context string)  { debugger.Log(2, "%s [1]Message_RequestBatch: from<vp%d>, <%+v>", context, m.ReplicaId, m) }

func (m *Request) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type PrePrepare struct {
	View           uint64        `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	SequenceNumber uint64        `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	BatchDigest    string        `protobuf:"bytes,3,opt,name=batch_digest,json=batchDigest" json:"batch_digest,omitempty"`
	RequestBatch   *RequestBatch `protobuf:"bytes,4,opt,name=request_batch,json=requestBatch" json:"request_batch,omitempty"`
	ReplicaId      uint64        `protobuf:"varint,5,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
}

func (m *PrePrepare) Reset()                    { *m = PrePrepare{} }
func (m *PrePrepare) String() string            { return proto.CompactTextString(m) }
func (*PrePrepare) ProtoMessage()               {}
func (*PrePrepare) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }
func (m *PrePrepare) Dump(context string)  { debugger.Log(2, "%s [2]Message_PrePrepare: from<vp%d>, <%+v>", context, m.ReplicaId, m) }


func (m *PrePrepare) GetRequestBatch() *RequestBatch {
	if m != nil {
		return m.RequestBatch
	}
	return nil
}

type Prepare struct {
	View           uint64 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	SequenceNumber uint64 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	BatchDigest    string `protobuf:"bytes,3,opt,name=batch_digest,json=batchDigest" json:"batch_digest,omitempty"`
	ReplicaId      uint64 `protobuf:"varint,4,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
}

func (m *Prepare) Reset()                    { *m = Prepare{} }
func (m *Prepare) String() string            { return proto.CompactTextString(m) }
func (*Prepare) ProtoMessage()               {}
func (*Prepare) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }
func (m *Prepare) Dump(context string)  { debugger.Log(2, "%s [3]Message_Prepare: from<vp%d>, <%+v>", context, m.ReplicaId, m) }

type Commit struct {
	View           uint64 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	SequenceNumber uint64 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	BatchDigest    string `protobuf:"bytes,3,opt,name=batch_digest,json=batchDigest" json:"batch_digest,omitempty"`
	ReplicaId      uint64 `protobuf:"varint,4,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	BlockHeigth    uint64 `protobuf:"varint,5,opt,name=block_heigth,json=blockHeigth" json:"block_heigth,omitempty"`
	CurBlockHash   []byte `protobuf:"bytes,6,opt,name=curblockhash,proto3" json:"curblockhash,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }
func (m *Commit) Dump(context string)  { debugger.Log(2, "%s [4]Message_Commit: from<vp%d>, bh<%d>",
	context, m.ReplicaId, m.BlockHeigth) }

type BlockInfo struct {
	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber" json:"block_number,omitempty"`
	BlockHash   []byte `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (m *BlockInfo) Reset()                    { *m = BlockInfo{} }
func (m *BlockInfo) String() string            { return proto.CompactTextString(m) }
func (*BlockInfo) ProtoMessage()               {}
func (*BlockInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Checkpoint struct {
	SequenceNumber uint64 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	ReplicaId      uint64 `protobuf:"varint,2,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	Id             string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *Checkpoint) Reset()                    { *m = Checkpoint{} }
func (m *Checkpoint) String() string            { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()               {}
func (*Checkpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }
func (m *Checkpoint) Dump(context string)  { debugger.Log(2, "%s [5]Message_Checkpoint: from<vp%d>, <%+v>", context, m.ReplicaId, m) }

type ViewChange struct {
	View      uint64           `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	H         uint64           `protobuf:"varint,2,opt,name=h" json:"h,omitempty"`
	Cset      []*ViewChange_C  `protobuf:"bytes,3,rep,name=cset" json:"cset,omitempty"`
	Pset      []*ViewChange_PQ `protobuf:"bytes,4,rep,name=pset" json:"pset,omitempty"`
	Qset      []*ViewChange_PQ `protobuf:"bytes,5,rep,name=qset" json:"qset,omitempty"`
	ReplicaId uint64           `protobuf:"varint,6,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	Signature []byte           `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ViewChange) Reset()                    { *m = ViewChange{} }
func (m *ViewChange) String() string            { return proto.CompactTextString(m) }
func (*ViewChange) ProtoMessage()               {}
func (*ViewChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }
func (m *ViewChange) Dump(context string)  { debugger.Log(2, "%s [6]Message_ViewChange: from<vp%d>, <%+v>", context, m.ReplicaId, m) }

func (m *ViewChange) GetCset() []*ViewChange_C {
	if m != nil {
		return m.Cset
	}
	return nil
}

func (m *ViewChange) GetPset() []*ViewChange_PQ {
	if m != nil {
		return m.Pset
	}
	return nil
}

func (m *ViewChange) GetQset() []*ViewChange_PQ {
	if m != nil {
		return m.Qset
	}
	return nil
}

// This message should go away and become a checkpoint once replica_id is removed
type ViewChange_C struct {
	SequenceNumber uint64 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	Id             string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *ViewChange_C) Reset()                    { *m = ViewChange_C{} }
func (m *ViewChange_C) String() string            { return proto.CompactTextString(m) }
func (*ViewChange_C) ProtoMessage()               {}
func (*ViewChange_C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type ViewChange_PQ struct {
	SequenceNumber uint64 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	BatchDigest    string `protobuf:"bytes,2,opt,name=batch_digest,json=batchDigest" json:"batch_digest,omitempty"`
	View           uint64 `protobuf:"varint,3,opt,name=view" json:"view,omitempty"`
}

func (m *ViewChange_PQ) Reset()                    { *m = ViewChange_PQ{} }
func (m *ViewChange_PQ) String() string            { return proto.CompactTextString(m) }
func (*ViewChange_PQ) ProtoMessage()               {}
func (*ViewChange_PQ) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 1} }

type PQset struct {
	Set []*ViewChange_PQ `protobuf:"bytes,1,rep,name=set" json:"set,omitempty"`
}

func (m *PQset) Reset()                    { *m = PQset{} }
func (m *PQset) String() string            { return proto.CompactTextString(m) }
func (*PQset) ProtoMessage()               {}
func (*PQset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PQset) GetSet() []*ViewChange_PQ {
	if m != nil {
		return m.Set
	}
	return nil
}

type NewView struct {
	View      uint64            `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Vset      []*ViewChange     `protobuf:"bytes,2,rep,name=vset" json:"vset,omitempty"`
	Xset      map[uint64]string `protobuf:"bytes,3,rep,name=xset" json:"xset,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ReplicaId uint64            `protobuf:"varint,4,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
}

func (m *NewView) Reset()                    { *m = NewView{} }
func (m *NewView) String() string            { return proto.CompactTextString(m) }
func (*NewView) ProtoMessage()               {}
func (*NewView) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }
func (m *NewView) Dump(context string)  { debugger.Log(2, "%s [7]Message_NewView: from<vp%d>, <%+v>", context, m.ReplicaId, m) }

func (m *NewView) GetVset() []*ViewChange {
	if m != nil {
		return m.Vset
	}
	return nil
}

func (m *NewView) GetXset() map[uint64]string {
	if m != nil {
		return m.Xset
	}
	return nil
}

type FetchRequestBatch struct {
	BatchDigest string `protobuf:"bytes,1,opt,name=batch_digest,json=batchDigest" json:"batch_digest,omitempty"`
	ReplicaId   uint64 `protobuf:"varint,2,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
}

func (m *FetchRequestBatch) Reset()                    { *m = FetchRequestBatch{} }
func (m *FetchRequestBatch) String() string            { return proto.CompactTextString(m) }
func (*FetchRequestBatch) ProtoMessage()               {}
func (*FetchRequestBatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }
func (m *FetchRequestBatch) Dump(context string)  { debugger.Log(2, "%s [8]Message_FetchRequestBatch: from<vp%d>, <%+v>", context, m.ReplicaId, m) }

type RequestBatch struct {
	Batch []*Request `protobuf:"bytes,1,rep,name=batch" json:"batch,omitempty"`
}

func (m *RequestBatch) Reset()                    { *m = RequestBatch{} }
func (m *RequestBatch) String() string            { return proto.CompactTextString(m) }
func (*RequestBatch) ProtoMessage()               {}
func (*RequestBatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }
func (m *RequestBatch) Dump(context string)  { debugger.Log(2, "%s [9]Message_ReturnRequestBatch: from<vp%d>, <%+v>", context, -1, m) }

func (m *RequestBatch) GetBatch() []*Request {
	if m != nil {
		return m.Batch
	}
	return nil
}

type BatchMessage struct {
	// Types that are valid to be assigned to Payload:
	//	*BatchMessage_Request
	//	*BatchMessage_RequestBatch
	//	*BatchMessage_PbftMessage
	//	*BatchMessage_Complaint
	Payload isBatchMessage_Payload `protobuf_oneof:"payload"`
}

func (m *BatchMessage) Reset()                    { *m = BatchMessage{} }
func (m *BatchMessage) String() string            { return proto.CompactTextString(m) }
func (*BatchMessage) ProtoMessage()               {}
func (*BatchMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type isBatchMessage_Payload interface {
	isBatchMessage_Payload()
}

type BatchMessage_Request struct {
	Request *Request `protobuf:"bytes,1,opt,name=request,oneof"`
}
type BatchMessage_RequestBatch struct {
	RequestBatch *RequestBatch `protobuf:"bytes,2,opt,name=request_batch,json=requestBatch,oneof"`
}
type BatchMessage_PbftMessage struct {
	PbftMessage []byte `protobuf:"bytes,3,opt,name=pbft_message,json=pbftMessage,proto3,oneof"`
}
type BatchMessage_Complaint struct {
	Complaint *Request `protobuf:"bytes,4,opt,name=complaint,oneof"`
}

func (*BatchMessage_Request) isBatchMessage_Payload()      {}
func (*BatchMessage_RequestBatch) isBatchMessage_Payload() {}
func (*BatchMessage_PbftMessage) isBatchMessage_Payload()  {}
func (*BatchMessage_Complaint) isBatchMessage_Payload()    {}

func (m *BatchMessage) GetPayload() isBatchMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *BatchMessage) GetRequest() *Request {
	if x, ok := m.GetPayload().(*BatchMessage_Request); ok {
		return x.Request
	}
	return nil
}

func (m *BatchMessage) GetRequestBatch() *RequestBatch {
	if x, ok := m.GetPayload().(*BatchMessage_RequestBatch); ok {
		return x.RequestBatch
	}
	return nil
}

func (m *BatchMessage) GetPbftMessage() []byte {
	if x, ok := m.GetPayload().(*BatchMessage_PbftMessage); ok {
		return x.PbftMessage
	}
	return nil
}

func (m *BatchMessage) GetComplaint() *Request {
	if x, ok := m.GetPayload().(*BatchMessage_Complaint); ok {
		return x.Complaint
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BatchMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BatchMessage_OneofMarshaler, _BatchMessage_OneofUnmarshaler, _BatchMessage_OneofSizer, []interface{}{
		(*BatchMessage_Request)(nil),
		(*BatchMessage_RequestBatch)(nil),
		(*BatchMessage_PbftMessage)(nil),
		(*BatchMessage_Complaint)(nil),
	}
}

func _BatchMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BatchMessage)
	// payload
	switch x := m.Payload.(type) {
	case *BatchMessage_Request:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Request); err != nil {
			return err
		}
	case *BatchMessage_RequestBatch:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestBatch); err != nil {
			return err
		}
	case *BatchMessage_PbftMessage:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.PbftMessage)
	case *BatchMessage_Complaint:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Complaint); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BatchMessage.Payload has unexpected type %T", x)
	}
	return nil
}

func _BatchMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BatchMessage)
	switch tag {
	case 1: // payload.request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Request)
		err := b.DecodeMessage(msg)
		m.Payload = &BatchMessage_Request{msg}
		return true, err
	case 2: // payload.request_batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestBatch)
		err := b.DecodeMessage(msg)
		m.Payload = &BatchMessage_RequestBatch{msg}
		return true, err
	case 3: // payload.pbft_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &BatchMessage_PbftMessage{x}
		return true, err
	case 4: // payload.complaint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Request)
		err := b.DecodeMessage(msg)
		m.Payload = &BatchMessage_Complaint{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BatchMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BatchMessage)
	// payload
	switch x := m.Payload.(type) {
	case *BatchMessage_Request:
		s := proto.Size(x.Request)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BatchMessage_RequestBatch:
		s := proto.Size(x.RequestBatch)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BatchMessage_PbftMessage:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PbftMessage)))
		n += len(x.PbftMessage)
	case *BatchMessage_Complaint:
		s := proto.Size(x.Complaint)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Metadata struct {
	SeqNo uint64 `protobuf:"varint,1,opt,name=seqNo" json:"seqNo,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*Message)(nil), "pbft.message")
	proto.RegisterType((*Request)(nil), "pbft.request")
	proto.RegisterType((*PrePrepare)(nil), "pbft.pre_prepare")
	proto.RegisterType((*Prepare)(nil), "pbft.prepare")
	proto.RegisterType((*Commit)(nil), "pbft.commit")
	proto.RegisterType((*BlockInfo)(nil), "pbft.block_info")
	proto.RegisterType((*Checkpoint)(nil), "pbft.checkpoint")
	proto.RegisterType((*ViewChange)(nil), "pbft.view_change")
	proto.RegisterType((*ViewChange_C)(nil), "pbft.view_change.C")
	proto.RegisterType((*ViewChange_PQ)(nil), "pbft.view_change.PQ")
	proto.RegisterType((*PQset)(nil), "pbft.PQset")
	proto.RegisterType((*NewView)(nil), "pbft.new_view")
	proto.RegisterType((*FetchRequestBatch)(nil), "pbft.fetch_request_batch")
	proto.RegisterType((*RequestBatch)(nil), "pbft.request_batch")
	proto.RegisterType((*BatchMessage)(nil), "pbft.batch_message")
	proto.RegisterType((*Metadata)(nil), "pbft.metadata")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0xee, 0x52, 0x94, 0x64, 0x8e, 0x68, 0xd5, 0x5e, 0xfb, 0xc0, 0x0a, 0x35, 0xea, 0xd2, 0xa8,
	0x6b, 0xa3, 0x2d, 0x0d, 0xb8, 0x06, 0x6a, 0x18, 0x3d, 0xd9, 0x2d, 0xea, 0xa2, 0xa8, 0x21, 0x13,
	0x41, 0x92, 0x1b, 0x41, 0x51, 0x2b, 0x89, 0xb0, 0x44, 0xd2, 0x14, 0xe5, 0xc4, 0x2f, 0x90, 0xe4,
	0x92, 0x17, 0xc8, 0x03, 0xe5, 0x92, 0x73, 0xde, 0x27, 0xbb, 0xc3, 0xa5, 0xf8, 0x23, 0x45, 0xf6,
	0x25, 0xc8, 0x81, 0xc0, 0xee, 0xcc, 0xf7, 0x0d, 0x67, 0xbf, 0x9d, 0xdd, 0x59, 0x68, 0x4f, 0xd8,
	0x74, 0xea, 0x0e, 0xd9, 0xd4, 0x8a, 0xe2, 0x30, 0x09, 0xa9, 0x1a, 0xf5, 0x06, 0x49, 0xe7, 0x87,
	0x61, 0x18, 0x0e, 0xc7, 0xec, 0x08, 0x6d, 0xbd, 0xd9, 0xe0, 0x28, 0xf1, 0x39, 0x2e, 0x71, 0x27,
	0x51, 0x0a, 0x33, 0x5f, 0xa9, 0xd0, 0x94, 0x4c, 0x7a, 0x06, 0xeb, 0x31, 0xbb, 0x9d, 0x71, 0xbf,
	0xd3, 0x73, 0x13, 0x6f, 0x64, 0x90, 0x5d, 0x72, 0xd0, 0x3a, 0xde, 0xb2, 0x44, 0x28, 0xab, 0xe4,
	0xba, 0xfc, 0xc6, 0xd6, 0xa5, 0xe1, 0x5c, 0xcc, 0xe9, 0x09, 0xb4, 0xa2, 0x98, 0x39, 0xfc, 0x8b,
	0xdc, 0x98, 0x19, 0x0a, 0x32, 0x37, 0x53, 0x66, 0xc1, 0xc1, 0x79, 0xc0, 0x87, 0xdd, 0x74, 0x46,
	0x0f, 0xa1, 0x99, 0x31, 0x6a, 0xc8, 0x58, 0x9f, 0x33, 0x24, 0x3a, 0xf3, 0xd3, 0x7d, 0x68, 0x78,
	0xe1, 0x64, 0xe2, 0x27, 0x86, 0x8a, 0x48, 0x3d, 0x45, 0xa6, 0x36, 0x0e, 0x94, 0x5e, 0x7a, 0x0c,
	0xe0, 0x8d, 0x98, 0x77, 0x13, 0x85, 0x7e, 0x90, 0x18, 0x75, 0xc4, 0x6e, 0x48, 0xec, 0xdc, 0x2e,
	0xd2, 0xc8, 0x67, 0x22, 0xf9, 0x3b, 0x9f, 0xbd, 0x70, 0xbc, 0x91, 0x1b, 0x0c, 0x99, 0xd1, 0x28,
	0x26, 0x5f, 0x70, 0x08, 0x96, 0x98, 0x5e, 0xe0, 0x8c, 0xfe, 0x02, 0x6b, 0x01, 0xf7, 0x09, 0x8b,
	0xd1, 0x44, 0x4a, 0x3b, 0xa5, 0x64, 0x56, 0x91, 0x3e, 0x1f, 0x3f, 0xe5, 0x43, 0xfa, 0x1f, 0x6c,
	0x0d, 0x18, 0x17, 0xca, 0x29, 0x2b, 0xbc, 0x86, 0xbc, 0xef, 0x52, 0xde, 0x12, 0x00, 0x0f, 0xb1,
	0x89, 0x66, 0xbb, 0x28, 0xf6, 0x3f, 0xb0, 0x1d, 0xb3, 0x64, 0x16, 0x07, 0x95, 0x68, 0xda, 0xaa,
	0xfd, 0xa2, 0x29, 0xa5, 0x18, 0xe8, 0x5c, 0xe3, 0xfa, 0xbb, 0xf7, 0xe3, 0xd0, 0xed, 0x9b, 0xef,
	0x08, 0x34, 0x25, 0x85, 0x9e, 0x82, 0x36, 0xaf, 0x13, 0x59, 0x04, 0x1d, 0x2b, 0xad, 0x24, 0x2b,
	0xab, 0x24, 0xeb, 0x49, 0x86, 0xb0, 0x73, 0x30, 0x35, 0xe6, 0x01, 0xb1, 0x04, 0x74, 0x3b, 0x9b,
	0xd2, 0x1d, 0x00, 0xbe, 0x93, 0x63, 0xdf, 0x73, 0x1d, 0xbf, 0x8f, 0xbb, 0xad, 0xda, 0x9a, 0xb4,
	0xfc, 0xdb, 0xa7, 0xdf, 0x83, 0x36, 0xf5, 0x87, 0x81, 0xcb, 0x53, 0x64, 0xb8, 0xc3, 0xba, 0x9d,
	0x1b, 0xcc, 0xf7, 0xa4, 0x54, 0x5e, 0x94, 0x82, 0x8a, 0xb2, 0x13, 0x0c, 0x83, 0x63, 0xfa, 0x33,
	0x7c, 0x3b, 0x15, 0xf9, 0x07, 0x1e, 0x73, 0x82, 0xd9, 0xa4, 0xc7, 0x62, 0x4c, 0x41, 0xb5, 0xdb,
	0x99, 0xf9, 0x0a, 0xad, 0xf4, 0x47, 0xd0, 0x51, 0x13, 0xa7, 0xef, 0xf3, 0xe3, 0x92, 0x60, 0x2e,
	0x9a, 0xdd, 0x42, 0xdb, 0x5f, 0x68, 0xe2, 0x02, 0x54, 0x4e, 0x82, 0xfa, 0x59, 0x65, 0x2b, 0xe7,
	0xa0, 0xbc, 0xcc, 0x7a, 0x65, 0x99, 0xe6, 0x1b, 0x32, 0xaf, 0xf8, 0x2f, 0xbe, 0x88, 0x72, 0x2a,
	0x6a, 0x35, 0x95, 0xd7, 0x24, 0x3b, 0x51, 0x5f, 0x3b, 0x93, 0x2b, 0x80, 0xde, 0x38, 0xf4, 0x6e,
	0x1c, 0x3f, 0x18, 0x84, 0x18, 0x0f, 0x67, 0xf2, 0xaf, 0x69, 0x52, 0x2d, 0xb4, 0xc9, 0x5f, 0xee,
	0x64, 0x84, 0x91, 0x3b, 0x1d, 0xc9, 0x42, 0xd3, 0xd0, 0x72, 0xc9, 0x0d, 0x66, 0xbf, 0x78, 0x05,
	0x2c, 0x5b, 0x08, 0x59, 0xba, 0x90, 0x72, 0x96, 0x4a, 0xb5, 0x42, 0xdb, 0xa0, 0xc8, 0xc2, 0xd5,
	0x6c, 0x3e, 0x32, 0xdf, 0xd6, 0x4a, 0xb7, 0xc6, 0x52, 0x11, 0x75, 0x20, 0x23, 0x19, 0x89, 0x8c,
	0x78, 0x26, 0xaa, 0x37, 0x65, 0x42, 0xa1, 0x5a, 0x5e, 0x4c, 0x85, 0x10, 0xd6, 0x85, 0x8d, 0x00,
	0x7a, 0x00, 0x6a, 0x24, 0x80, 0x2a, 0x02, 0xb7, 0x17, 0x81, 0xdd, 0x6b, 0x1b, 0x11, 0x02, 0x79,
	0x2b, 0x90, 0xf5, 0x55, 0x48, 0x81, 0xa8, 0xac, 0xae, 0xb1, 0xf2, 0xfc, 0x35, 0x2b, 0xe7, 0xaf,
	0xf3, 0x27, 0x90, 0x8b, 0xc7, 0x0b, 0x59, 0x51, 0xaa, 0xd3, 0x07, 0xa5, 0x7b, 0xfd, 0x78, 0x7a,
	0xb5, 0xa0, 0x94, 0xc5, 0x82, 0xca, 0xb4, 0xae, 0xe5, 0x5a, 0x9b, 0x47, 0x50, 0xef, 0x5e, 0x8b,
	0x95, 0xee, 0x43, 0x4d, 0x48, 0x42, 0x56, 0x48, 0x22, 0x00, 0xe6, 0x07, 0x92, 0x5f, 0xe0, 0x4b,
	0x77, 0xef, 0x27, 0x6e, 0x13, 0x91, 0x14, 0x8c, 0xb4, 0xd8, 0x0f, 0x6c, 0x74, 0xd3, 0x5f, 0x41,
	0x7d, 0x99, 0x6f, 0xab, 0x51, 0xee, 0x01, 0xd6, 0x73, 0xee, 0xfa, 0x3b, 0x48, 0xe2, 0x7b, 0x1b,
	0x51, 0x0f, 0x9c, 0x85, 0xce, 0x1f, 0xa0, 0xcd, 0x19, 0x74, 0x03, 0x6a, 0x37, 0xec, 0x5e, 0xe6,
	0x24, 0x86, 0x74, 0x1b, 0xea, 0x77, 0xee, 0x78, 0xc6, 0xa4, 0x28, 0xe9, 0xe4, 0x4c, 0x39, 0x25,
	0xe6, 0xb3, 0xa5, 0x0d, 0x66, 0x41, 0x4c, 0xf2, 0xd0, 0xe9, 0xac, 0xd6, 0xbd, 0x79, 0x52, 0xb9,
	0x0b, 0xe9, 0x1e, 0xd4, 0xb3, 0xe7, 0x41, 0x2d, 0x6f, 0xd9, 0x12, 0x63, 0xa7, 0x3e, 0xf3, 0x23,
	0x81, 0xf5, 0xf4, 0xc7, 0xd9, 0xeb, 0xe2, 0x70, 0xde, 0x5f, 0x64, 0x4b, 0x29, 0x13, 0x45, 0xb3,
	0xcc, 0xfa, 0xcf, 0xc2, 0x43, 0x44, 0x79, 0xfc, 0x43, 0x64, 0x0f, 0x74, 0x81, 0xca, 0x7e, 0x8b,
	0x25, 0xa2, 0x73, 0x54, 0x4b, 0x58, 0xff, 0x97, 0xb9, 0xfc, 0x06, 0x1a, 0xbf, 0xfa, 0xa2, 0xb1,
	0x2b, 0xde, 0x08, 0xea, 0xf2, 0x6c, 0x72, 0x44, 0xb1, 0x4d, 0xee, 0xc2, 0xda, 0x84, 0x25, 0x6e,
	0xdf, 0x4d, 0x5c, 0xb1, 0x19, 0xbc, 0x74, 0xaf, 0x42, 0xb9, 0x41, 0xe9, 0xa4, 0xd7, 0xc0, 0x0e,
	0xf9, 0xfb, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x21, 0x64, 0x8c, 0x91, 0x09, 0x00, 0x00,
}
