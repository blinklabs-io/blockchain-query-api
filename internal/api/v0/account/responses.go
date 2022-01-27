package account

import (
	"encoding/hex"
	"strconv"
)

type AccountResponse struct {
	Status           string `json:"status"`
	DelegatedPool    string `json:"delegated_pool"`
	TotalBalance     string `json:"total_balance"`
	Utxo             string `json:"utxo"`
	Rewards          string `json:"rewards"`
	Withdrawals      string `json:"withdrawals"`
	RewardsAvailable string `json:"rewards_available"`
	Reserves         string `json:"reserves"`
	Treasury         string `json:"treasury"`
}

type AddressResponse struct {
	Address string `json:"address"`
}

type AssetResponse struct {
	AssetPolicy string `json:"asset_policy"`
	AssetName   string `json:"asset_name"`
	Quantity    string `json:"quantity"`
}

type HistoryResponse struct {
	StakeAddress string `json:"stake_address"`
	PoolId       string `json:"pool_id"`
	EpochNumber  uint64 `json:"epoch_no"`
	ActiveStake  string `json:"active_stake"`
}

type RewardResponse struct {
	EarnedEpoch    uint64 `json:"earned_epoch"`
	SpendableEpoch uint64 `json:"spendable_epoch"`
	Amount         string `json:"amount"`
	Type           string `json:"type"`
	PoolId         string `json:"pool_id"`
}

type UpdateResponse struct {
	ActionType string `json:"action_type"`
	TxHash     string `json:"tx_hash"`
}

func NewAccountResponse(a *Account) *AccountResponse {
	r := &AccountResponse{
		Status:           a.Status,
		DelegatedPool:    string(a.DelegatedPool),
		TotalBalance:     strconv.FormatUint(a.TotalBalance, 10),
		Utxo:             strconv.FormatUint(a.Utxo, 10),
		Rewards:          strconv.FormatUint(a.Rewards, 10),
		Withdrawals:      strconv.FormatUint(a.Withdrawals, 10),
		RewardsAvailable: strconv.FormatUint(a.RewardsAvailable, 10),
		Reserves:         strconv.FormatUint(a.Reserves, 10),
		Treasury:         strconv.FormatUint(a.Treasury, 10),
	}
	return r
}

func NewAddressResponse(a *Address) *AddressResponse {
	r := &AddressResponse{
		Address: a.Address,
	}
	return r
}

func NewAssetResponse(a *Asset) *AssetResponse {
	r := &AssetResponse{
		AssetPolicy: hex.EncodeToString(a.AssetPolicy),
		AssetName:   string(a.AssetName),
		Quantity:    strconv.FormatUint(a.Quantity, 10),
	}
	return r
}

func NewHistoryResponse(h *History) *HistoryResponse {
	r := &HistoryResponse{
		StakeAddress: h.StakeAddress,
		PoolId:       h.PoolId,
		EpochNumber:  h.EpochNumber,
		ActiveStake:  strconv.FormatUint(h.ActiveStake, 10),
	}
	return r
}

func NewRewardResponse(x *Reward) *RewardResponse {
	r := &RewardResponse{
		EarnedEpoch:    x.EarnedEpoch,
		SpendableEpoch: x.SpendableEpoch,
		Amount:         strconv.FormatUint(x.Amount, 10),
		Type:           x.Type,
		PoolId:         x.PoolId,
	}
	return r
}

func NewUpdateResponse(u *Update) *UpdateResponse {
	r := &UpdateResponse{
		ActionType: u.ActionType,
		TxHash:     hex.EncodeToString(u.TxHash),
	}
	return r
}
