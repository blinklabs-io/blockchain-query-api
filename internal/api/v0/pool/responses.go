package pool

import (
	"encoding/hex",
	"strconv"
)

type BlockResponse struct {
	EpochNumber uint64 `json:"epoch_no"`
	EpochSlot   uint64 `json:"epoch_slot"`
	AbsSlot     uint64 `json:"abs_slot"`
	BlockHeight uint64 `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   *time.Time `json:"block_time"`
}

func NewBlockResponse(b *Block) *BlockResponse {
	r := &BlockResponse{
		EpochNumber: strconv.FormatUint(r.EpochNumber, 10),
		EpochSlot:   strconv.FormatUint(r.EpochSlot, 10),
		AbsSlot:     strconv.FormatUint(r.AbsSlot, 10),
		BlockHeight: strconv.FormatUint(r.BlockHeight, 10),
		BlockHash:   hex.EncodeToString(r.BlockHash),
		BlockTime:   r.BlockTime,
	}
	return r
}
