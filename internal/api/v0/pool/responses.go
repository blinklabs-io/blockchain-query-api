package pool

import (
	"encoding/hex"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/postgres/types"
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
	ActiveStakePercent float64 `json:"active_stake_pct"`
	SaturationPercent  float64 `json:"saturation_pct"`
	BlockCount         uint64  `json:"block_cnt"`
	DelegatorCount     uint64  `json:"delegator_cnt"`
	Margin             float32 `json:"margin"`
	FixedCost          string  `json:"fixed_cost"`
	PoolFees           string  `json:"pool_fees"`
	DelegRewards       string  `json:"deleg_rewards"`
	EpochRos           float64 `json:"epoch_ros"`
}

type InfoResponse struct {
	PoolIdBech32      string      `json:"pool_id_bech32"`
	PoolIdHex         string      `json:"pool_id_hex"`
	ActiveEpochNumber uint64      `json:"active_epoch_no"`
	VrfHashKey        string      `json:"vrf_hash_key"`
	Margin            float32     `json:"margin"`
	FixedCost         string      `json:"fixed_cost"`
	Pledge            string      `json:"pledge"`
	RewardAddress     string      `json:"reward_address"`
	Owners            string      `json:"owners"`
	Relays            types.Jsonb `json:"relays"`
	MetaUrl           string      `json:"meta_url"`
	MetaHash          string      `json:"meta_hash"`
	MetaJson          types.Jsonb `json:"meta_json"`
	PoolStatus        string      `json:"pool_status"`
	RetiringEpoch     uint64      `json:"retiring_epoch"`
	OpCert            string      `json:"op_cert"`
	OpCertCounter     uint32      `json:"op_cert_counter"`
	ActiveStake       string      `json:"active_stake"`
	BlockCount        float64     `json:"block_count"`
	LivePledge        string      `json:"live_pledge"`
	LiveStake         string      `json:"live_stake"`
	LiveDelegators    uint64      `json:"live_delegators"`
	LiveSaturation    float64     `json:"live_saturation"`
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
		PoolFees:           strconv.FormatFloat(float64(h.PoolFees), 'f', 2, 32),
		DelegRewards:       strconv.FormatFloat(h.DelegRewards, 'f', 2, 64),
		EpochRos:           h.EpochRos,
	}
	return r
}

func NewInfoResponse(p *Pool) *InfoResponse {
	r := &InfoResponse{
		PoolIdBech32:      p.PoolIdBech32,
		PoolIdHex:         p.PoolIdHex,
		ActiveEpochNumber: p.ActiveEpochNumber,
		VrfHashKey:        p.VrfHashKey,
		Margin:            p.Margin,
		FixedCost:         strconv.FormatUint(p.FixedCost, 10),
		Pledge:            strconv.FormatUint(p.Pledge, 10),
		RewardAddress:     p.RewardAddress,
		Owners:            p.Owners,
		Relays:            p.Relays,
		MetaUrl:           p.MetaUrl,
		MetaHash:          p.MetaHash,
		MetaJson:          p.MetaJson,
		PoolStatus:        p.PoolStatus,
		RetiringEpoch:     p.RetiringEpoch,
		OpCert:            p.OpCert,
		OpCertCounter:     p.OpCertCounter,
		ActiveStake:       strconv.FormatUint(p.ActiveStake, 10),
		BlockCount:        p.BlockCount,
		LivePledge:        strconv.FormatFloat(float64(p.LivePledge), 'f', 2, 32),
		LiveStake:         strconv.FormatUint(p.LiveStake, 10),
		LiveDelegators:    p.LiveDelegators,
		LiveSaturation:    p.LiveSaturation,
	}
	return r
}
