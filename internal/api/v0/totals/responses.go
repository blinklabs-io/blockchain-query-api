package totals

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"strconv"
)

type TotalsResponse struct {
	EpochNumber uint32 `json:"epoch_no"`
	Circulation string `json:"circulation"`
	Treasury    string `json:"treasury"`
	Rewards     string `json:"reward"`
	Supply      string `json:"supply"`
	Reserves    string `json:"reserves"`
}

func NewTotalsResponse(t *models.AdaPots) *TotalsResponse {
	r := &TotalsResponse{
		EpochNumber: t.EpochNumber,
		Circulation: strconv.FormatUint(t.Utxo, 10),
		Treasury:    strconv.FormatUint(t.Treasury, 10),
		Rewards:     strconv.FormatUint(t.Rewards, 10),
		Supply:      strconv.FormatUint((t.Treasury + t.Rewards + t.Utxo + t.Deposits + t.Fees), 10),
		Reserves:    strconv.FormatUint(t.Reserves, 10),
	}
	return r
}
