package pool

import (
	"encoding/hex"
	"strconv"
	"time"
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

type HistoryResponse struct {
	EpochNumber        uint64  `json:"epoch_no"`
	ActiveStake        string  `json:"active_stake"`
	ActiveStakePercent float32 `json:"active_stake_pct"`
	SaturationPercent  float32 `json:"saturation_pct"`
	BlockCount         uint64  `json:"block_cnt"`
	DelegatorCount     uint64  `json:"delegator_cnt"`
	Margin             float32 `json:"margin"`
	FixedCost          string  `json:"fixed_cost"`
	PoolFees           string  `json:"pool_fees"`
	DelegRewards       string  `json:"deleg_rewards"`
	EpochRos           float32 `json:"epoch_ros"`
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
		Amount:       strconv.FormatFloat(d.TotalBalance, 'f', 2, 64),
		EpochNumber:  d.EpochNumber,
	}
	return r
}

func NewHistoryResponse(h *History) *HistoryResponse {
	r := &HistoryResponse{
		EpochNumber:        h.EpochNumber,
		ActiveStake:        strconv.FormatUint(h.ActiveStake, 10),
		ActiveStakePercent: h.ActiveStakePercent,
		SaturationPercent:  h.SaturationPercent,
		BlockCount:         h.BlockCount,
		DelegatorCount:     h.DelegatorCount,
		Margin:             h.Margin,
		FixedCost:          strconv.FormatUint(h.FixedCost, 10),
		PoolFees:           strconv.FormatFloat(h.PoolFees, 'f', 2, 32),
		DelegRewards:       strconv.FormatFloat(h.DelegRewards, 'f', 2, 64),
		EpochRos:           h.EpochRos,
	}
	return r
}
