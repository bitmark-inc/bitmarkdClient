// Copyright (c) 2014-2017 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bitmarkdClient

import (
	"encoding/json"
)

// GetNodeInfo will fetch node info from bitmarkd rpc server
func (bc *BitmarkdRPCClient) GetNodeInfo() (json.RawMessage, error) {
	args := RPCEmptyArguments{}
	var reply json.RawMessage
	err := bc.call("Node.Info", &args, &reply)
	return reply, err
}

type BlockDumpRangeArguments struct {
	Height uint64 `json:"height,string"`
	Count  int    `json:"count"`
	Txs    bool   `json:"txs"`
}

func (bc *BitmarkdRPCClient) BlockDumpRange(height uint64, count int, txs bool) (json.RawMessage, error) {
	args := BlockDumpRangeArguments{Height: height, Count: count, Txs: txs}
	var reply json.RawMessage
	err := bc.call("Node.BlockDumpRange", &args, &reply)
	return reply, err
}

type BlockDumpArguments struct {
	Height uint64 `json:"height,string"`
	Binary bool   `json:"binary"`
}

func (bc *BitmarkdRPCClient) BlockDump(height uint64, binary bool) (json.RawMessage, error) {
	args := BlockDumpArguments{Height: height, Binary: binary}
	var reply json.RawMessage
	err := bc.call("Node.BlockDump", &args, &reply)
	return reply, err
}
