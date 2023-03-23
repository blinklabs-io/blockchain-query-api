package meta

import (
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"time"
)

type MetaResponse struct {
	StartTime *time.Time `json:"start_time"`
	Network   string     `json:"network_name"`
}

func NewMetaResponse(e *models.Meta) *MetaResponse {
	r := &MetaResponse{
		StartTime: e.StartTime,
		Network:   e.Network,
	}
	return r
}
