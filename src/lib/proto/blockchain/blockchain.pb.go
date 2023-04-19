// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: blockchain.proto

package blockchain

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx                string `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx" form:"tx"`
	From              string `protobuf:"bytes,2,opt,name=from,proto3" json:"from" form:"from"`
	To                string `protobuf:"bytes,3,opt,name=to,proto3" json:"to" form:"to"`
	Amount            string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount" form:"amount"`
	Gas               uint64 `protobuf:"varint,5,opt,name=gas,proto3" json:"gas" form:"gas"`
	GasPrice          string `protobuf:"bytes,6,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price" form:"gas_price"`
	GasTip            string `protobuf:"bytes,7,opt,name=gas_tip,json=gasTip,proto3" json:"gas_tip" form:"gas_tip"`
	GasFee            string `protobuf:"bytes,8,opt,name=gas_fee,json=gasFee,proto3" json:"gas_fee" form:"gas_fee"`
	Nonce             uint64 `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce" form:"nonce"`
	Cost              string `protobuf:"bytes,10,opt,name=cost,proto3" json:"cost" form:"cost"`
	ChainId           string `protobuf:"bytes,11,opt,name=chain_id,json=chainId,proto3" json:"chain_id" form:"chain_id"`
	IsPending         bool   `protobuf:"varint,12,opt,name=is_pending,json=isPending,proto3" json:"is_pending" form:"is_pending"`
	BlockHash         string `protobuf:"bytes,13,opt,name=block_hash,json=blockHash,proto3" json:"block_hash" form:"block_hash"`
	BlockNumber       string `protobuf:"bytes,14,opt,name=block_number,json=blockNumber,proto3" json:"block_number" form:"block_number"`
	TransactionIndex  uint64 `protobuf:"varint,15,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index" form:"transaction_index"`
	Type              uint32 `protobuf:"varint,16,opt,name=type,proto3" json:"type" form:"type"`
	Status            uint64 `protobuf:"varint,17,opt,name=status,proto3" json:"status" form:"status"`
	GasUsed           uint64 `protobuf:"varint,18,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used" form:"gas_used"`
	EffectiveGasPrice string `protobuf:"bytes,19,opt,name=effective_gas_price,json=effectiveGasPrice,proto3" json:"effective_gas_price" form:"effective_gas_price"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[0]
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
	return file_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetTx() string {
	if x != nil {
		return x.Tx
	}
	return ""
}

func (x *Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Transaction) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Transaction) GetGas() uint64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *Transaction) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *Transaction) GetGasTip() string {
	if x != nil {
		return x.GasTip
	}
	return ""
}

func (x *Transaction) GetGasFee() string {
	if x != nil {
		return x.GasFee
	}
	return ""
}

func (x *Transaction) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetCost() string {
	if x != nil {
		return x.Cost
	}
	return ""
}

func (x *Transaction) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *Transaction) GetIsPending() bool {
	if x != nil {
		return x.IsPending
	}
	return false
}

func (x *Transaction) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Transaction) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *Transaction) GetTransactionIndex() uint64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *Transaction) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Transaction) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Transaction) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *Transaction) GetEffectiveGasPrice() string {
	if x != nil {
		return x.EffectiveGasPrice
	}
	return ""
}

// HTTP POST request: /blockchain/transaction/detail
type GetTransactionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx string `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx" form:"tx"`
}

func (x *GetTransactionReq) Reset() {
	*x = GetTransactionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionReq) ProtoMessage() {}

func (x *GetTransactionReq) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionReq.ProtoReflect.Descriptor instead.
func (*GetTransactionReq) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *GetTransactionReq) GetTx() string {
	if x != nil {
		return x.Tx
	}
	return ""
}

type GetTransactionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	// response message
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
	Data    *Transaction `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
}

func (x *GetTransactionRsp) Reset() {
	*x = GetTransactionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionRsp) ProtoMessage() {}

func (x *GetTransactionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionRsp.ProtoReflect.Descriptor instead.
func (*GetTransactionRsp) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransactionRsp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetTransactionRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetTransactionRsp) GetData() *Transaction {
	if x != nil {
		return x.Data
	}
	return nil
}

// HTTP Get request: /blockchain/account/detail
type GetAccountDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address" form:"address"`
}

func (x *GetAccountDetailReq) Reset() {
	*x = GetAccountDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDetailReq) ProtoMessage() {}

func (x *GetAccountDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDetailReq.ProtoReflect.Descriptor instead.
func (*GetAccountDetailReq) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountDetailReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// HTTP Get request: /blockchain/account/detail
type GetAccountDetailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	// response message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
	// 账号
	Data *EthAccount `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
}

func (x *GetAccountDetailRsp) Reset() {
	*x = GetAccountDetailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDetailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDetailRsp) ProtoMessage() {}

func (x *GetAccountDetailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDetailRsp.ProtoReflect.Descriptor instead.
func (*GetAccountDetailRsp) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccountDetailRsp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAccountDetailRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAccountDetailRsp) GetData() *EthAccount {
	if x != nil {
		return x.Data
	}
	return nil
}

type EthAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// account address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address" form:"address"`
	// 余额
	Balance string `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance" form:"balance"`
	// nonce
	Nonce uint64 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce" form:"nonce"`
}

func (x *EthAccount) Reset() {
	*x = EthAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthAccount) ProtoMessage() {}

func (x *EthAccount) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthAccount.ProtoReflect.Descriptor instead.
func (*EthAccount) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{5}
}

func (x *EthAccount) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EthAccount) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *EthAccount) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

// HTTP Get request: /blockchain/gas/suggest
type GetEthSuggestGasReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEthSuggestGasReq) Reset() {
	*x = GetEthSuggestGasReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEthSuggestGasReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEthSuggestGasReq) ProtoMessage() {}

func (x *GetEthSuggestGasReq) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEthSuggestGasReq.ProtoReflect.Descriptor instead.
func (*GetEthSuggestGasReq) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{6}
}

// HTTP Get request: /blockchain/gas/suggest
type GetEthSuggestGasRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	// response message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
	// 账号
	Data *GasDetail `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
}

func (x *GetEthSuggestGasRsp) Reset() {
	*x = GetEthSuggestGasRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEthSuggestGasRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEthSuggestGasRsp) ProtoMessage() {}

func (x *GetEthSuggestGasRsp) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEthSuggestGasRsp.ProtoReflect.Descriptor instead.
func (*GetEthSuggestGasRsp) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{7}
}

func (x *GetEthSuggestGasRsp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetEthSuggestGasRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetEthSuggestGasRsp) GetData() *GasDetail {
	if x != nil {
		return x.Data
	}
	return nil
}

type GasDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gas price suggest
	GasPriceSuggest string `protobuf:"bytes,1,opt,name=gas_price_suggest,json=gasPriceSuggest,proto3" json:"gas_price_suggest" form:"gas_price_suggest"`
	// gas tip cap suggest
	GasTipCapSuggest string `protobuf:"bytes,2,opt,name=gas_tip_cap_suggest,json=gasTipCapSuggest,proto3" json:"gas_tip_cap_suggest" form:"gas_tip_cap_suggest"`
}

func (x *GasDetail) Reset() {
	*x = GasDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GasDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GasDetail) ProtoMessage() {}

func (x *GasDetail) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GasDetail.ProtoReflect.Descriptor instead.
func (*GasDetail) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{8}
}

func (x *GasDetail) GetGasPriceSuggest() string {
	if x != nil {
		return x.GasPriceSuggest
	}
	return ""
}

func (x *GasDetail) GetGasTipCapSuggest() string {
	if x != nil {
		return x.GasTipCapSuggest
	}
	return ""
}

// HTTP Get request: /wallet/key/create
type GetKeyPairRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	// response message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
	// 账号
	Data *KeyPair `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
}

func (x *GetKeyPairRsp) Reset() {
	*x = GetKeyPairRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyPairRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyPairRsp) ProtoMessage() {}

func (x *GetKeyPairRsp) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyPairRsp.ProtoReflect.Descriptor instead.
func (*GetKeyPairRsp) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{9}
}

func (x *GetKeyPairRsp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetKeyPairRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetKeyPairRsp) GetData() *KeyPair {
	if x != nil {
		return x.Data
	}
	return nil
}

// HTTP Get request: /wallet/mnemonic/create
type GetMnemonicRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code" form:"code"`
	// response message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message" form:"message"`
	// 账号
	Data *KeyPair `protobuf:"bytes,3,opt,name=data,proto3" json:"data" form:"data"`
}

func (x *GetMnemonicRsp) Reset() {
	*x = GetMnemonicRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMnemonicRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMnemonicRsp) ProtoMessage() {}

func (x *GetMnemonicRsp) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMnemonicRsp.ProtoReflect.Descriptor instead.
func (*GetMnemonicRsp) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{10}
}

func (x *GetMnemonicRsp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetMnemonicRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMnemonicRsp) GetData() *KeyPair {
	if x != nil {
		return x.Data
	}
	return nil
}

type KeyPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// private_key
	PrivateKey string `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key" form:"private_key"`
	// public_key
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key" form:"public_key"`
	// address
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address" form:"address"`
	// mnemonic
	Mnemonic string `protobuf:"bytes,4,opt,name=mnemonic,proto3" json:"mnemonic" form:"mnemonic"`
}

func (x *KeyPair) Reset() {
	*x = KeyPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyPair) ProtoMessage() {}

func (x *KeyPair) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyPair.ProtoReflect.Descriptor instead.
func (*KeyPair) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{11}
}

func (x *KeyPair) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *KeyPair) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *KeyPair) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *KeyPair) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

var File_blockchain_proto protoreflect.FileDescriptor

var file_blockchain_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x84,
	0x04, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x73,
	0x5f, 0x74, 0x69, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x73, 0x54,
	0x69, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x73, 0x46, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61,
	0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x61,
	0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x73,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x78, 0x22, 0x6e, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6f, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x74, 0x68, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x0a,
	0x45, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x74, 0x68, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x47, 0x61, 0x73, 0x52, 0x65, 0x71, 0x22, 0x6e, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x45, 0x74, 0x68, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x47, 0x61, 0x73, 0x52,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x61, 0x73, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x09, 0x47,
	0x61, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x61, 0x73, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x67, 0x61, 0x73, 0x5f, 0x74, 0x69, 0x70, 0x5f,
	0x63, 0x61, 0x70, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x67, 0x61, 0x73, 0x54, 0x69, 0x70, 0x43, 0x61, 0x70, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69,
	0x72, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x4b, 0x65,
	0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x67, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x7f, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65,
	0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65,
	0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x42, 0x16, 0x5a, 0x14, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_proto_rawDescOnce sync.Once
	file_blockchain_proto_rawDescData = file_blockchain_proto_rawDesc
)

func file_blockchain_proto_rawDescGZIP() []byte {
	file_blockchain_proto_rawDescOnce.Do(func() {
		file_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_proto_rawDescData)
	})
	return file_blockchain_proto_rawDescData
}

var file_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_blockchain_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: blockchain.Transaction
	(*GetTransactionReq)(nil),   // 1: blockchain.GetTransactionReq
	(*GetTransactionRsp)(nil),   // 2: blockchain.GetTransactionRsp
	(*GetAccountDetailReq)(nil), // 3: blockchain.GetAccountDetailReq
	(*GetAccountDetailRsp)(nil), // 4: blockchain.GetAccountDetailRsp
	(*EthAccount)(nil),          // 5: blockchain.EthAccount
	(*GetEthSuggestGasReq)(nil), // 6: blockchain.GetEthSuggestGasReq
	(*GetEthSuggestGasRsp)(nil), // 7: blockchain.GetEthSuggestGasRsp
	(*GasDetail)(nil),           // 8: blockchain.GasDetail
	(*GetKeyPairRsp)(nil),       // 9: blockchain.GetKeyPairRsp
	(*GetMnemonicRsp)(nil),      // 10: blockchain.GetMnemonicRsp
	(*KeyPair)(nil),             // 11: blockchain.KeyPair
}
var file_blockchain_proto_depIdxs = []int32{
	0,  // 0: blockchain.GetTransactionRsp.data:type_name -> blockchain.Transaction
	5,  // 1: blockchain.GetAccountDetailRsp.data:type_name -> blockchain.EthAccount
	8,  // 2: blockchain.GetEthSuggestGasRsp.data:type_name -> blockchain.GasDetail
	11, // 3: blockchain.GetKeyPairRsp.data:type_name -> blockchain.KeyPair
	11, // 4: blockchain.GetMnemonicRsp.data:type_name -> blockchain.KeyPair
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_blockchain_proto_init() }
func file_blockchain_proto_init() {
	if File_blockchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_blockchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionReq); i {
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
		file_blockchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionRsp); i {
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
		file_blockchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountDetailReq); i {
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
		file_blockchain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountDetailRsp); i {
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
		file_blockchain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthAccount); i {
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
		file_blockchain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEthSuggestGasReq); i {
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
		file_blockchain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEthSuggestGasRsp); i {
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
		file_blockchain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GasDetail); i {
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
		file_blockchain_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyPairRsp); i {
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
		file_blockchain_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMnemonicRsp); i {
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
		file_blockchain_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyPair); i {
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
			RawDescriptor: file_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_proto_goTypes,
		DependencyIndexes: file_blockchain_proto_depIdxs,
		MessageInfos:      file_blockchain_proto_msgTypes,
	}.Build()
	File_blockchain_proto = out.File
	file_blockchain_proto_rawDesc = nil
	file_blockchain_proto_goTypes = nil
	file_blockchain_proto_depIdxs = nil
}
