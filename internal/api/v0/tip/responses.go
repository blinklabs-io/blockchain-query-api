package tip

import (
	"encoding/hex"
	"time"
)

type TipResponse struct {
	Hash            string     `json:"hash"`
	EpochNumber     uint16     `json:"epoch_no"`
	SlotNumber      uint32     `json:"abs_slot"`
	EpochSlotNumber uint32     `json:"epoch_slot"`
	BlockNumber     uint32     `json:"block_no"`
	Time            *time.Time `json:"block_time"`
}

func NewTipResponse(t *Tip) *TipResponse {
	r := &TipResponse{
		Hash:            hex.EncodeToString(t.Hash),
		EpochNumber:     t.EpochNumber,
		SlotNumber:      t.SlotNumber,
		EpochSlotNumber: t.EpochSlotNumber,
		BlockNumber:     t.BlockNumber,
		Time:            t.Time,
	}
	return r
}
