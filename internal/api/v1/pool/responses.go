package pool

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
)

type PoolResponse struct {
	Address        string  `json:"pool_id"`
	CertIndex      int32   `json:"cert_index"`
	VrfKeyHash     []byte  `json:"vrf_key"`
	Pledge         uint64  `json:"pledge"`
	RewardAddr     string  `json:"reward_address"`
	Margin         float32 `json:"variable_fee"`
	FixedCost      uint64  `json:"fixed_fee"`
	RegisteredTxId int64   `json:"registered_tx_id"`
}

// Build response object from DB model
func NewPoolResponse(p *models.PoolUpdate, n string) *PoolResponse {
	r := &PoolResponse{
		Address:        n,
		CertIndex:      p.CertIndex,
		VrfKeyHash:     p.VrfKeyHash,
		Pledge:         p.Pledge,
		RewardAddr:     p.RewardAddr,
		Margin:         p.Margin,
		FixedCost:      p.FixedCost,
		RegisteredTxId: p.RegisteredTxId,
	}
	return r
}
