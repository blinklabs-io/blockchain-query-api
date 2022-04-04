package account

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync"
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/cardano_db_sync/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func ConfigureRoutes(g *gin.RouterGroup) {
	// Addresses
	g.GET("/account_addresses", HandleGetAccountAddresses)
	g.POST("/account_addresses", HandleGetAccountAddresses)
	// Assets
	g.GET("/account_assets", HandleGetAccountAssets)
	g.POST("/account_assets", HandleGetAccountAssets)
	// History
	g.GET("/account_history", HandleGetAccountHistory)
	g.POST("/account_history", HandleGetAccountHistory)
	// Info
	g.GET("/account_info", HandleGetAccountInfo)
	g.POST("/account_info", HandleGetAccountInfo)
	// List
	g.GET("/account_list", HandleGetAccountList)
	// Rewards
	g.GET("/account_rewards", HandleGetAccountRewards)
	g.POST("/account_rewards", HandleGetAccountRewards)
	// Updates
	g.GET("/account_updates", HandleGetAccountUpdates)
	g.POST("/account_updates", HandleGetAccountUpdates)
}

type Account struct {
	Status           string `gorm:"column:status"`
	DelegatedPool    string `gorm:"column:delegated_pool"`
	TotalBalance     uint64 `gorm:"column:total_balance"`
	Utxo             uint64 `gorm:"column:utxo"`
	Rewards          uint64 `gorm:"column:rewards"`
	Withdrawals      uint64 `gorm:"column:withdrawals"`
	RewardsAvailable uint64 `gorm:"column:rewards_available"`
	Reserves         uint64 `gorm:"column:reserves"`
	Treasury         uint64 `gorm:"column:treasury"`
}

type AccountId struct {
	Id string `gorm:"id"`
}

type Address struct {
	Address string `gorm:"column:address"`
}

type Asset struct {
	AssetPolicy []byte `gorm:"column:asset_policy"`
	AssetName   []byte `gorm:"column:asset_name"`
	Quantity    uint64 `gorm:"column:quantity"`
}

type History struct {
	StakeAddress string `gorm:"column:stake_address"`
	PoolId       string `gorm:"column:pool_id"`
	EpochNumber  uint64 `gorm:"column:epoch_no"`
	ActiveStake  uint64 `gorm:"column:active_stake"`
}

type Reward struct {
	EarnedEpoch    uint64 `gorm:"column:earned_epoch"`
	SpendableEpoch uint64 `gorm:"column:spendable_epoch"`
	Amount         uint64 `gorm:"column:amount"`
	Type           string `gorm:"column:type"`
	PoolId         string `gorm:"column:pool_id"`
}

type Update struct {
	ActionType string `gorm:"column:action_type"`
	TxHash     []byte `gorm:"column:tx_hash"`
}

func HandleGetAccountAddresses(c *gin.Context) {
	address := c.DefaultPostForm("_address", "")

	db := cardano_db_sync.GetHandle()
	var addrIdQuery *gorm.DB
	var addresses []*Address
	r := []*AddressResponse{}
	if address != "" {
		addrIdQuery = AddressIdQuery(db, address)
		result := db.Distinct("address").Where("stake_address_id = (?)", addrIdQuery).Table("tx_out").Find(&addresses)
		if result.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(result.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
		// Create response from returned item
		for _, v := range addresses {
			r = append(r, NewAddressResponse(v))
		}
	}
	c.JSON(200, r)
}

func HandleGetAccountAssets(c *gin.Context) {
	address := c.DefaultPostForm("_address", "")

	db := cardano_db_sync.GetHandle()
	var assets []*Asset
	var addrId uint64
	r := []*AssetResponse{}
	if address != "" {
		// Get database ID of our address
		addrIdResult := AddressIdQuery(db, address).Find(&addrId)
		if addrIdResult.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(addrIdResult.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
		result := db.
			Table("ma_tx_out mtx").
			Select("ma.policy AS asset_policy, ma.name AS asset_name, SUM(mtx.quantity) as quantity").
			Joins("INNER JOIN multi_asset ma ON ma.id = mtx.ident").
			Joins("INNER JOIN tx_out txo ON txo.id = mtx.tx_out_id AND txo.stake_address_id = ?", strconv.FormatUint(addrId, 10)).
			Where("NOT EXISTS (?)", TxOutQuery(db)).
			Group("ma.policy, ma.name").
			Find(&assets)
		if result.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(result.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
		// Create response from returned item
		for _, v := range assets {
			r = append(r, NewAssetResponse(v))
		}
	}
	c.JSON(200, r)
}

func HandleGetAccountHistory(c *gin.Context) {
	address := c.DefaultPostForm("_address", "")
	epoch := c.DefaultPostForm("_epoch_no", "")

	db := cardano_db_sync.GetHandle()
	var history History
	if address != "" {
		// Check if we were given a payment address and lookup stake
		if !strings.HasPrefix(address, "stake") {
			addressResult := db.Table("public.tx_out").
				Select("stake_address.view").
				Limit(1).
				Joins("INNER JOIR public.stake_address ON stake_address.id = tx_out.stake_address_id").
				Where("tx_out.address = ?", address).Find(&address)
			if addressResult.Error != nil {
				// Not found
				if cardano_db_sync.IsRecordNotFoundError(addressResult.Error) {
					c.JSON(404, gin.H{"msg": "address not found"})
					return
				}
				// Some other database error
				// TODO: log this failure
				c.JSON(500, gin.H{"msg": "unexpected error"})
				return
			}
		}
		// Epoch is provided
		if epoch != "" {
			result := db.
				Table("grest.account_active_stake_cache").
				Select("account_active_stake_cache.stake_address, account_active_stake_cache.pool_id, account_active_stake_cache.epoch_no, account_active_stake_cache.amount AS active_stake").
				Where("account_active_stake_cache.epoch_no = ? AND account_active_stake_cache.stake_address = ?",
					epoch,
					address).
				Find(&history)
			if result.Error != nil {
				// Not found
				if cardano_db_sync.IsRecordNotFoundError(result.Error) {
					c.JSON(404, gin.H{"msg": "history not found"})
					return
				}
				// Some other database error
				// TODO: log this failure
				c.JSON(500, gin.H{"msg": "unexpected error"})
				return
			}
			// No epoch provided, get full history
		} else {
			result := db.
				Table("grest.account_active_stake_cache").
				Select("account_active_stake_cache.stake_address, account_active_stake_cache.pool_id, account_active_stake_cache.epoch_no, account_active_stake_cache.amount AS active_stake").
				Where("account_active_stake_cache.stake_address", address).
				Find(&history)
			if result.Error != nil {
				// Not found
				if cardano_db_sync.IsRecordNotFoundError(result.Error) {
					c.JSON(404, gin.H{"msg": "history not found"})
					return
				}
				// Some other database error
				// TODO: log this failure
				c.JSON(500, gin.H{"msg": "unexpected error"})
				return
			}
		}
	}
	r := NewHistoryResponse(&history)
	c.JSON(200, r)
}

func HandleGetAccountInfo(c *gin.Context) {
	address := c.DefaultPostForm("_address", "")

	db := cardano_db_sync.GetHandle()
	var accounts []*Account
	var addrId uint64
	var latestWithdrawalEpoch uint64
	var latestWithdrawalTx uint64
	r := []*AccountResponse{}
	if address != "" {
		// Get database ID of our address
		addrIdResult := AddressIdQuery(db, address).Find(&addrId)
		if addrIdResult.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(addrIdResult.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
		// Get database ID of last withdrawal transaction
		latestWithdrawalTxResult := LatestWithdrawalTxQuery(db, addrId).Find(&latestWithdrawalTx)
		if latestWithdrawalTxResult.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(latestWithdrawalTxResult.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
		// Get database ID of last withdrawal transaction's epoch
		latestWithdrawalEpochResult := LatestWithdrawalEpochQuery(db, latestWithdrawalTx).Find(&latestWithdrawalEpoch)
		if latestWithdrawalEpochResult.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(latestWithdrawalEpochResult.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}

		rewardWithdrawalQuery := db.Raw("COALESCE(rewards_t.rewards, 0) - COALESCE(withdrawals_t.withdrawals, 0)")
		query1 := db.Raw("COALESCE(utxo_t.utxo, 0) + ? + COALESCE(reserves_t.reserves, 0) + COALESCE(treasury_t.treasury, 0) - (?)", rewardWithdrawalQuery, rewardWithdrawalQuery)
		query2 := db.Raw("COALESCE(utxo_t.utxo, 0) + ? + COALESCE(reserves_t.reserves, 0) + COALESCE(treasury_t.treasury, 0)", rewardWithdrawalQuery)
		totalBalanceQuery := db.Raw("CASE WHEN (?) < 0 THEN (?) ELSE (?) END", rewardWithdrawalQuery, query1, query2)
		rewardsAvailableQuery := db.Raw("CASE WHEN (?) <= 0 THEN 0 ELSE (?) END", rewardWithdrawalQuery, rewardWithdrawalQuery)
		result := db.
			Select("? AS status, pool_t.delegated_pool, ? AS total_balance, ? AS utxo, ? AS rewards, ? AS withdrawals, ? AS rewards_available, ? AS reserves, ? AS treasury",
				db.Raw("CASE WHEN status_t.registered = TRUE THEN 'registered' ELSE 'not registered' END"),
				totalBalanceQuery,
				db.Raw("COALESCE(utxo_t.utxo, 0)"),
				db.Raw("COALESCE(rewards_t.rewards, 0)"),
				db.Raw("COALESCE(withdrawals_t.withdrawals, 0)"),
				rewardsAvailableQuery,
				db.Raw("COALESCE(reserves_t.reserves, 0)"),
				db.Raw("COALESCE(treasury_t.treasury, 0)")).
			Table("(?) status_t",
				db.Raw("SELECT EXISTS (SELECT 1 FROM stake_registration WHERE stake_registration.addr_id = ? AND NOT EXISTS (SELECT TRUE FROM stake_deregistration WHERE stake_deregistration.addr_id = stake_registration.addr_id AND stake_deregistration.tx_id > stake_registration.tx_id)) AS registered", addrId)).
			Joins("LEFT JOIN (?) pool_t ON TRUE",
				db.Table("delegation").
					Select("pool_hash.view AS delegated_pool").
					Joins("INNER JOIN pool_hash ON pool_hash.id = delegation.pool_hash_id").
					Where("delegation.addr_id = ? AND NOT EXISTS (?) AND NOT EXISTS (?)",
						addrId,
						db.Table("delegation d").
							Select("TRUE").
							Where("d.addr_id = delegation.addr_id AND d.id > delegation.id"),
						db.Table("stake_deregistration").
							Select("TRUE").
							Where("stake_deregistration.addr_id = delegation.addr_id AND stake_deregistration.tx_id > delegation.tx_id"))).
			Joins("LEFT JOIN (?) utxo_t ON TRUE",
				db.Table("tx_out").
					Select("COALESCE(SUM(value), 0) AS utxo").
					Joins("LEFT JOIN tx_in ON tx_out.tx_id = tx_in.tx_out_id AND tx_out.index = tx_in.tx_out_index").
					Where("tx_out.stake_address_id = ? AND tx_in.tx_in_id IS NULL", addrId)).
			Joins("LEFT JOIN (?) rewards_t ON TRUE",
				db.Table("reward").
					Select("COALESCE(SUM(reward.amount), 0) AS rewards").
					Where("reward.addr_id = ? AND reward.spendable_epoch <= (?)",
						addrId,
						db.Table("epoch").Select("MAX(no)"))).
			Joins("LEFT JOIN (?) withdrawals_t ON TRUE",
				db.Table("withdrawal").
					Select("COALESCE(SUM(withdrawal.amount), 0) AS withdrawals").
					Where("withdrawal.addr_id = ?", addrId)).
			Joins("LEFT JOIN (?) reserves_t ON TRUE",
				db.Table("reserve").
					Select("COALESCE(SUM(reserve.amount), 0) AS reserves").
					Joins("INNER JOIN tx ON tx.id = reserve.tx_id").
					Joins("INNER JOIN block ON block.id = tx.block_id").
					Where("reserve.addr_id = ? AND block.epoch_no >= ?",
						addrId,
						latestWithdrawalEpoch)).
			Joins("LEFT JOIN (?) treasury_t ON TRUE",
				db.Table("treasury").
					Select("COALESCE(SUM(treasury.amount), 0) AS treasury").
					Joins("INNER JOIN tx ON tx.id = treasury.tx_id").
					Joins("INNER JOIN block ON block.id = tx.block_id").
					Where("treasury.addr_id = ? AND block.epoch_no >= ?",
						addrId,
						latestWithdrawalEpoch)).
			Find(&accounts)
		if result.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(result.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
		// Create response from returned item
		for _, v := range accounts {
			r = append(r, NewAccountResponse(v))
		}
	}
	c.JSON(200, r)
}

func HandleGetAccountList(c *gin.Context) {
	var accounts []*AccountId
	var stakeAddress *models.StakeAddress
	db := cardano_db_sync.GetHandle()
	result := db.Debug().
		Model(&stakeAddress).
		Select("view AS id").
		Find(&accounts)
	if result.Error != nil {
		// Not found
		if cardano_db_sync.IsRecordNotFoundError(result.Error) {
			c.JSON(404, gin.H{"msg": "accounts not found"})
			return
		}
		// Some other database error
		// TODO: log this failure
		c.JSON(500, gin.H{"msg": "unexpected error"})
		return
	}
	// Create response from returned item
	r := []*AccountIdResponse{}
	for _, v := range accounts {
		r = append(r, NewAccountListResponse(v))
	}
	c.JSON(200, r)
}

func HandleGetAccountRewards(c *gin.Context) {
	address := c.DefaultPostForm("_stake_address", "")
	epoch := c.DefaultPostForm("_epoch_no", "")

	db := cardano_db_sync.GetHandle()
	var addrId uint64
	var rewards []*Reward
	r := []*RewardResponse{}
	if address != "" {
		// Ensure we're a stake address
		if strings.HasPrefix(address, "stake") {
			addressResult := db.Table("stake_address").
				Select("stake_address.id").
				Where("stake_address.view", address).Find(&addrId)
			if addressResult.Error != nil {
				// Not found
				if cardano_db_sync.IsRecordNotFoundError(addressResult.Error) {
					c.JSON(200, r)
					return
				}
				// Some other database error
				// TODO: log this failure
				c.JSON(500, gin.H{"msg": "unexpected error"})
				return
			}
		}
		// Epoch is provided
		if epoch != "" {
			result := db.
				Table("reward r"). // no AS in upstream
				Select("r.earned_epoch, r.spendable_epoch, r.amount, r.type, ph.view AS pool_id").
				Joins("LEFT JOIN pool_hash AS ph ON r.pool_id = ph.id").
				Where("r.addr_id = ? AND r.earned_epoch = ?", addrId, epoch).
				Find(&rewards)
			if result.Error != nil {
				// Not found
				if cardano_db_sync.IsRecordNotFoundError(result.Error) {
					c.JSON(200, r)
					return
				}
				// Some other database error
				// TODO: log this failure
				c.JSON(500, gin.H{"msg": "unexpected error"})
				return
			}
			// No epoch provided, get full rewards
		} else {
			result := db.
				Table("reward AS r"). // AS is in upstream
				Select("r.earned_epoch, r.spendable_epoch, r.amount, r.type, ph.view AS pool_id").
				Joins("LEFT JOIN pool_hash AS ph ON r.pool_id = ph.id").
				Where("r.addr_id = ?", addrId).
				Order("r.spendable_epoch DESC").
				Find(&rewards)
			if result.Error != nil {
				// Not found
				if cardano_db_sync.IsRecordNotFoundError(result.Error) {
					c.JSON(200, r)
					return
				}
				// Some other database error
				// TODO: log this failure
				c.JSON(500, gin.H{"msg": "unexpected error"})
				return
			}
		}
		// Create response from returned item
		for _, v := range rewards {
			r = append(r, NewRewardResponse(v))
		}
	}
	c.JSON(200, r)
}

func HandleGetAccountUpdates(c *gin.Context) {
	address := c.DefaultPostForm("_stake_address", "")

	db := cardano_db_sync.GetHandle()
	var addrId uint64
	var updates []*Update
	r := []*UpdateResponse{}
	if address != "" {
		// Ensure we're a stake address
		if strings.HasPrefix(address, "stake") {
			addressResult := db.Table("stake_address").
				Select("stake_address.id").
				Where("stake_address.view", address).Find(&addrId)
			if addressResult.Error != nil {
				// Not found
				if cardano_db_sync.IsRecordNotFoundError(addressResult.Error) {
					c.JSON(200, r)
					return
				}
				// Some other database error
				// TODO: log this failure
				c.JSON(500, gin.H{"msg": "unexpected error"})
				return
			}
		}
		result := db.
			Table("(?) actions",
				db.Raw("(?) UNION (?) UNION (?) UNION (?)",
					db.Table("stake_registration").
						Select("'registration' AS action_type, tx_id").
						Where("addr_id = ?", addrId),
					db.Table("stake_deregistration").
						Select("'deregistration' AS action_type, tx_id").
						Where("addr_id = ?", addrId),
					db.Table("delegation").
						Select("'delegation' AS action_type, tx_id").
						Where("addr_id = ?", addrId),
					db.Table("withdrawal").
						Select("'withdrawal' AS action_type, tx_id").
						Where("addr_id = ?", addrId))).
			Select("actions.action_type, tx.hash AS tx_hash").
			Joins("INNER JOIN tx ON tx.id = actions.tx_id").
			Order("tx.id ASC, actions.action_type DESC").
			Find(&updates)
		if result.Error != nil {
			// Not found
			if cardano_db_sync.IsRecordNotFoundError(result.Error) {
				c.JSON(200, r)
				return
			}
			// Some other database error
			// TODO: log this failure
			c.JSON(500, gin.H{"msg": "unexpected error"})
			return
		}
		for _, v := range updates {
			r = append(r, NewUpdateResponse(v))
		}
	}
	c.JSON(200, r)
}
