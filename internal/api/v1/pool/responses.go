package pools

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
)

type PoolsResponse struct {







	Hash            []byte `json:"hash"`
	EpochNumber     uint16 `json:"epoch_no"`
	SlotNumber      uint32 `json:"slot_no"`
	EpochSlotNumber uint32 `json:"epoch_slot_no"`
	PoolsNumber     uint32 `json:"block_no"`
	//PreviousID      int64      `json:"previous_id"`
	//SlotLeaderID    int64      `json:"slot_leader_id"`
	Size       uint32     `json:"size"`
	Time       *time.Time `json:"time"`
	TxCount    int64      `json:"tx_count"`
	ProtoMajor uint16     `json:"proto_major"`
	ProtoMinor uint16     `json:"proto_minor"`
	VrfKey     string     `json:"vrf_key"`
	//OpCert          []byte     `json:"op_cert"`
	//OpCertCounter   uint32     `json:"op_cert_counter"`
}

// Build response object from DB model
func NewPoolsResponse(b *models.Pools) *PoolsResponse {
	r := &PoolsResponse{
		Hash:            b.Hash,
		EpochNumber:     b.EpochNumber,
		SlotNumber:      b.SlotNumber,
		EpochSlotNumber: b.EpochSlotNumber,
		PoolsNumber:     b.PoolsNumber,
		Size:            b.Size,
		Time:            b.Time,
		TxCount:         b.TxCount,
		ProtoMajor:      b.ProtoMajor,
		ProtoMinor:      b.ProtoMinor,
		VrfKey:          b.VrfKey,
	}
	return r
}
