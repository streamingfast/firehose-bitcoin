// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: sf/bitcoin/type/v1/type.proto

package pbbtc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//the block hash
	// Bitcoin core reverses the byte order of the hash when printing it out has Hex, to prevent
	// the end user from making a mistake we store it as a string directly
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// The block size
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// The block size excluding witness data
	StrippedSize int32 `protobuf:"varint,4,opt,name=stripped_size,json=strippedSize,proto3" json:"stripped_size,omitempty"`
	// The block weight as defined in BIP 141
	Weight int32 `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// The block height or index
	Height int64 `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	// The block version
	Version int32 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
	// The block version formatted in hexadecimal
	VersionHex string `protobuf:"bytes,8,opt,name=version_hex,json=versionHex,proto3" json:"version_hex,omitempty"`
	// The merkle root
	MerkleRoot string `protobuf:"bytes,9,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	// Transaction array
	Tx []*Transaction `protobuf:"bytes,10,rep,name=tx,proto3" json:"tx,omitempty"`
	// The block time expressed in UNIX epoch time
	Time int64 `protobuf:"varint,11,opt,name=time,proto3" json:"time,omitempty"`
	// The median block time expressed in UNIX epoch time
	Mediantime int64 `protobuf:"varint,12,opt,name=mediantime,proto3" json:"mediantime,omitempty"`
	// The nonce
	Nonce uint32 `protobuf:"varint,13,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The bits
	Bits string `protobuf:"bytes,14,opt,name=bits,proto3" json:"bits,omitempty"`
	// The difficulty
	Difficulty float64 `protobuf:"fixed64,15,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	// Expected number of hashes required to produce the chain up to this block (in hex)
	Chainwork string `protobuf:"bytes,16,opt,name=chainwork,proto3" json:"chainwork,omitempty"`
	// The number of transactions in the block
	NTx uint32 `protobuf:"varint,17,opt,name=n_tx,json=nTx,proto3" json:"n_tx,omitempty"`
	// The hash of the previous block
	PreviousHash string `protobuf:"bytes,18,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_sf_bitcoin_type_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Block) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Block) GetStrippedSize() int32 {
	if x != nil {
		return x.StrippedSize
	}
	return 0
}

func (x *Block) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Block) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Block) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Block) GetVersionHex() string {
	if x != nil {
		return x.VersionHex
	}
	return ""
}

func (x *Block) GetMerkleRoot() string {
	if x != nil {
		return x.MerkleRoot
	}
	return ""
}

func (x *Block) GetTx() []*Transaction {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *Block) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Block) GetMediantime() int64 {
	if x != nil {
		return x.Mediantime
	}
	return 0
}

func (x *Block) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Block) GetBits() string {
	if x != nil {
		return x.Bits
	}
	return ""
}

func (x *Block) GetDifficulty() float64 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *Block) GetChainwork() string {
	if x != nil {
		return x.Chainwork
	}
	return ""
}

func (x *Block) GetNTx() uint32 {
	if x != nil {
		return x.NTx
	}
	return 0
}

func (x *Block) GetPreviousHash() string {
	if x != nil {
		return x.PreviousHash
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The serialized, hex-encoded data for 'txid'
	Hex string `protobuf:"bytes,1,opt,name=hex,proto3" json:"hex,omitempty"`
	// The transaction id
	Txid string `protobuf:"bytes,2,opt,name=txid,proto3" json:"txid,omitempty"`
	// The transaction hash (differs from txid for witness transactions)
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	// The serialized transaction size
	Size int32 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	// The virtual transaction size (differs from size for witness transactions)
	Vsize int32 `protobuf:"varint,5,opt,name=vsize,proto3" json:"vsize,omitempty"`
	// The transaction's weight (between vsize*4-3 and vsize*4)
	Weight int32 `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	//  The version
	Version uint32 `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
	// The lock time
	Locktime uint32  `protobuf:"varint,8,opt,name=locktime,proto3" json:"locktime,omitempty"`
	Vin      []*Vin  `protobuf:"bytes,9,rep,name=vin,proto3" json:"vin,omitempty"`
	Vout     []*Vout `protobuf:"bytes,10,rep,name=vout,proto3" json:"vout,omitempty"`
	// the block hash
	Blockhash string `protobuf:"bytes,11,opt,name=blockhash,proto3" json:"blockhash,omitempty"`
	// The block time expressed in UNIX epoch time
	Blocktime int64 `protobuf:"varint,12,opt,name=blocktime,proto3" json:"blocktime,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_sf_bitcoin_type_v1_type_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetHex() string {
	if x != nil {
		return x.Hex
	}
	return ""
}

func (x *Transaction) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Transaction) GetVsize() int32 {
	if x != nil {
		return x.Vsize
	}
	return 0
}

func (x *Transaction) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Transaction) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Transaction) GetLocktime() uint32 {
	if x != nil {
		return x.Locktime
	}
	return 0
}

func (x *Transaction) GetVin() []*Vin {
	if x != nil {
		return x.Vin
	}
	return nil
}

func (x *Transaction) GetVout() []*Vout {
	if x != nil {
		return x.Vout
	}
	return nil
}

func (x *Transaction) GetBlockhash() string {
	if x != nil {
		return x.Blockhash
	}
	return ""
}

func (x *Transaction) GetBlocktime() int64 {
	if x != nil {
		return x.Blocktime
	}
	return 0
}

type Vin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The transaction id
	Txid string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	// The output number
	Vout uint32 `protobuf:"varint,2,opt,name=vout,proto3" json:"vout,omitempty"`
	// The script
	ScriptSig *ScriptSig `protobuf:"bytes,3,opt,name=script_sig,json=scriptSig,proto3" json:"script_sig,omitempty"`
	// The script sequence number
	Sequence uint32 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// hex-encoded witness data (if any)
	Txinwitness []string `protobuf:"bytes,5,rep,name=txinwitness,proto3" json:"txinwitness,omitempty"`
	// hex-encoded coinbase
	Coinbase string `protobuf:"bytes,6,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
}

func (x *Vin) Reset() {
	*x = Vin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vin) ProtoMessage() {}

func (x *Vin) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vin.ProtoReflect.Descriptor instead.
func (*Vin) Descriptor() ([]byte, []int) {
	return file_sf_bitcoin_type_v1_type_proto_rawDescGZIP(), []int{2}
}

func (x *Vin) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *Vin) GetVout() uint32 {
	if x != nil {
		return x.Vout
	}
	return 0
}

func (x *Vin) GetScriptSig() *ScriptSig {
	if x != nil {
		return x.ScriptSig
	}
	return nil
}

func (x *Vin) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Vin) GetTxinwitness() []string {
	if x != nil {
		return x.Txinwitness
	}
	return nil
}

func (x *Vin) GetCoinbase() string {
	if x != nil {
		return x.Coinbase
	}
	return ""
}

type Vout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The value in BTC
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// index
	N            uint32        `protobuf:"varint,2,opt,name=n,proto3" json:"n,omitempty"`
	ScriptPubKey *ScriptPubKey `protobuf:"bytes,3,opt,name=script_pubKey,json=scriptPubKey,proto3" json:"script_pubKey,omitempty"`
}

func (x *Vout) Reset() {
	*x = Vout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vout) ProtoMessage() {}

func (x *Vout) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vout.ProtoReflect.Descriptor instead.
func (*Vout) Descriptor() ([]byte, []int) {
	return file_sf_bitcoin_type_v1_type_proto_rawDescGZIP(), []int{3}
}

func (x *Vout) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Vout) GetN() uint32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *Vout) GetScriptPubKey() *ScriptPubKey {
	if x != nil {
		return x.ScriptPubKey
	}
	return nil
}

type ScriptSig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The asm
	Asm string `protobuf:"bytes,1,opt,name=asm,proto3" json:"asm,omitempty"`
	// The hex
	Hex string `protobuf:"bytes,2,opt,name=hex,proto3" json:"hex,omitempty"`
}

func (x *ScriptSig) Reset() {
	*x = ScriptSig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScriptSig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScriptSig) ProtoMessage() {}

func (x *ScriptSig) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScriptSig.ProtoReflect.Descriptor instead.
func (*ScriptSig) Descriptor() ([]byte, []int) {
	return file_sf_bitcoin_type_v1_type_proto_rawDescGZIP(), []int{4}
}

func (x *ScriptSig) GetAsm() string {
	if x != nil {
		return x.Asm
	}
	return ""
}

func (x *ScriptSig) GetHex() string {
	if x != nil {
		return x.Hex
	}
	return ""
}

type ScriptPubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the asm
	Asm string `protobuf:"bytes,1,opt,name=asm,proto3" json:"asm,omitempty"`
	// the hex
	Hex string `protobuf:"bytes,2,opt,name=hex,proto3" json:"hex,omitempty"`
	// The required sigs
	ReqSigs int32 `protobuf:"varint,3,opt,name=req_sigs,json=reqSigs,proto3" json:"req_sigs,omitempty"`
	// The type, eg 'pubkeyhash'
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// bitcoin address
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// bitcoin addresses (deprecated, empty when 'address' is set)
	Addresses []string `protobuf:"bytes,6,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *ScriptPubKey) Reset() {
	*x = ScriptPubKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScriptPubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScriptPubKey) ProtoMessage() {}

func (x *ScriptPubKey) ProtoReflect() protoreflect.Message {
	mi := &file_sf_bitcoin_type_v1_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScriptPubKey.ProtoReflect.Descriptor instead.
func (*ScriptPubKey) Descriptor() ([]byte, []int) {
	return file_sf_bitcoin_type_v1_type_proto_rawDescGZIP(), []int{5}
}

func (x *ScriptPubKey) GetAsm() string {
	if x != nil {
		return x.Asm
	}
	return ""
}

func (x *ScriptPubKey) GetHex() string {
	if x != nil {
		return x.Hex
	}
	return ""
}

func (x *ScriptPubKey) GetReqSigs() int32 {
	if x != nil {
		return x.ReqSigs
	}
	return 0
}

func (x *ScriptPubKey) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ScriptPubKey) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ScriptPubKey) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

var File_sf_bitcoin_type_v1_type_proto protoreflect.FileDescriptor

var file_sf_bitcoin_type_v1_type_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x66, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x73, 0x66, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x03, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x70, 0x70,
	0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73,
	0x74, 0x72, 0x69, 0x70, 0x70, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x68, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x48, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72,
	0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x2f, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x74, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x69, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x69, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63,
	0x75, 0x6c, 0x74, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x77,
	0x6f, 0x72, 0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x11, 0x0a, 0x04, 0x6e, 0x5f, 0x74, 0x78, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6e, 0x54, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x48, 0x61, 0x73, 0x68, 0x22, 0xd4, 0x02, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x68, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68, 0x65, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x29,
	0x0a, 0x03, 0x76, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66,
	0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x69, 0x6e, 0x52, 0x03, 0x76, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x04, 0x76, 0x6f, 0x75,
	0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x69, 0x74,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x75,
	0x74, 0x52, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0xc5, 0x01, 0x0a, 0x03, 0x56, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x76,
	0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x73, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x69, 0x74,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x53, 0x69, 0x67, 0x52, 0x09, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69,
	0x67, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x78, 0x69, 0x6e, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x78, 0x69, 0x6e, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x22, 0x71, 0x0a, 0x04, 0x56,
	0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x6e, 0x12, 0x45, 0x0a, 0x0d, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x5f, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x73, 0x66, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x52, 0x0c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x2f,
	0x0a, 0x09, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x73, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x73, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x68, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68, 0x65, 0x78, 0x22,
	0x99, 0x01, 0x0a, 0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x73, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x68, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x5f, 0x73, 0x69, 0x67, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x71, 0x53, 0x69, 0x67, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x42, 0x4d, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f, 0x73, 0x65,
	0x2d, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x73, 0x66, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x62, 0x74, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sf_bitcoin_type_v1_type_proto_rawDescOnce sync.Once
	file_sf_bitcoin_type_v1_type_proto_rawDescData = file_sf_bitcoin_type_v1_type_proto_rawDesc
)

func file_sf_bitcoin_type_v1_type_proto_rawDescGZIP() []byte {
	file_sf_bitcoin_type_v1_type_proto_rawDescOnce.Do(func() {
		file_sf_bitcoin_type_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_bitcoin_type_v1_type_proto_rawDescData)
	})
	return file_sf_bitcoin_type_v1_type_proto_rawDescData
}

var file_sf_bitcoin_type_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sf_bitcoin_type_v1_type_proto_goTypes = []interface{}{
	(*Block)(nil),        // 0: sf.bitcoin.type.v1.Block
	(*Transaction)(nil),  // 1: sf.bitcoin.type.v1.Transaction
	(*Vin)(nil),          // 2: sf.bitcoin.type.v1.Vin
	(*Vout)(nil),         // 3: sf.bitcoin.type.v1.Vout
	(*ScriptSig)(nil),    // 4: sf.bitcoin.type.v1.ScriptSig
	(*ScriptPubKey)(nil), // 5: sf.bitcoin.type.v1.ScriptPubKey
}
var file_sf_bitcoin_type_v1_type_proto_depIdxs = []int32{
	1, // 0: sf.bitcoin.type.v1.Block.tx:type_name -> sf.bitcoin.type.v1.Transaction
	2, // 1: sf.bitcoin.type.v1.Transaction.vin:type_name -> sf.bitcoin.type.v1.Vin
	3, // 2: sf.bitcoin.type.v1.Transaction.vout:type_name -> sf.bitcoin.type.v1.Vout
	4, // 3: sf.bitcoin.type.v1.Vin.script_sig:type_name -> sf.bitcoin.type.v1.ScriptSig
	5, // 4: sf.bitcoin.type.v1.Vout.script_pubKey:type_name -> sf.bitcoin.type.v1.ScriptPubKey
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sf_bitcoin_type_v1_type_proto_init() }
func file_sf_bitcoin_type_v1_type_proto_init() {
	if File_sf_bitcoin_type_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_bitcoin_type_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_bitcoin_type_v1_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_bitcoin_type_v1_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_bitcoin_type_v1_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vout); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_bitcoin_type_v1_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScriptSig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_bitcoin_type_v1_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScriptPubKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sf_bitcoin_type_v1_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_bitcoin_type_v1_type_proto_goTypes,
		DependencyIndexes: file_sf_bitcoin_type_v1_type_proto_depIdxs,
		MessageInfos:      file_sf_bitcoin_type_v1_type_proto_msgTypes,
	}.Build()
	File_sf_bitcoin_type_v1_type_proto = out.File
	file_sf_bitcoin_type_v1_type_proto_rawDesc = nil
	file_sf_bitcoin_type_v1_type_proto_goTypes = nil
	file_sf_bitcoin_type_v1_type_proto_depIdxs = nil
}
