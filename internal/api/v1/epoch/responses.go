package epoch

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"time"
)

type EpochResponse struct {
	OutSum      uint64     `json:"out_sum"`
	Fees        uint64     `json:"fees"`
	TxCount     uint32     `json:"tx_count"`
	BlockCount  uint32     `json:"blk_count"`
	EpochNumber uint32     `json:"no"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
}

type EpochStakeResponse struct {
	Amount uint64 `json:"amount"`
}

func NewEpochResponse(e *models.Epoch) *EpochResponse {
	r := &EpochResponse{
		OutSum:      e.OutSum,
		Fees:        e.Fees,
		TxCount:     e.TxCount,
		BlockCount:  e.BlockCount,
		EpochNumber: e.EpochNumber,
		StartTime:   e.StartTime,
		EndTime:     e.EndTime,
	}
	return r
}

func NewEpochStakeResponse(e *StakeAmount) *EpochStakeResponse {
	r := &EpochStakeResponse{
		Amount: e.Amount,
	}
	return r
}
