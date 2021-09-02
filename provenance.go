// Copyright (c) 2014-2017 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bitmarkdClient

import (
	"encoding/json"
)

type FullProvenanceArguments struct {
	Id string `json:"bitmarkId"`
}

// GetBitmarkFullProvenance returns the provenance for a bitmark.
func (bc *BitmarkdRPCClient) GetBitmarkFullProvenance(bitmarkID string) (json.RawMessage, error) {
	args := FullProvenanceArguments{Id: bitmarkID}
	var reply json.RawMessage
	err := bc.call("Bitmark.FullProvenance", &args, &reply)
	return reply, err
}
