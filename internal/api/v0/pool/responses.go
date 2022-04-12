package pool

import (
	"encoding/hex"
	"strconv"
)

type BlockResponse struct {
	EpochNumber uint64     `json:"epoch_no"`
	EpochSlot   uint64     `json:"epoch_slot"`
	AbsSlot     uint64     `json:"abs_slot"`
	BlockHeight uint64     `json:"block_height"`
	BlockHash   string     `json:"block_hash"`
	BlockTime   *time.Time `json:"block_time"`
}

type DelegatorResponse struct {
	StakeAddress string `json:"stake_address"`
	Amount       string `json:"amount"`
	EpochNumber  uint64 `json:"epoch_no"`
}

func NewBlockResponse(b *Block) *BlockResponse {
	r := &BlockResponse{
		EpochNumber: b.EpochNumber,
		EpochSlot:   b.EpochSlot,
		AbsSlot:     b.AbsSlot,
		BlockHeight: b.BlockHeight,
		BlockHash:   hex.EncodeToString(b.BlockHash),
		BlockTime:   b.BlockTime,
	}
	return r
}

func NewDelegatorResponse(d *Delegator) *DelegatorResponse {
	r := &DelegatorResponse{
		StakeAddress: d.StakeAddress,
		Amount:       strconv.FormatFloat(d.Amount, 'f', 2, 32),
		EpochNumber:  d.EpochNumber,
	}
	return r
}
