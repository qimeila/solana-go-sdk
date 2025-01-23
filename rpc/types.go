package rpc


import (
	"encoding/json"
)

// Commitment describes how finalized a block is at that point in time
type Commitment string

const (
	CommitmentFinalized Commitment = "finalized"
	CommitmentConfirmed Commitment = "confirmed"
	CommitmentProcessed Commitment = "processed"
)

type Context struct {
	Slot       uint64 `json:"slot"`
	ApiVersion string `json:"apiVersion,omitempty"`
}

type AccountEncoding string

const (
	AccountEncodingBase58     AccountEncoding = "base58"
	AccountEncodingJsonParsed AccountEncoding = "jsonParsed"
	AccountEncodingBase64     AccountEncoding = "base64"
	AccountEncodingBase64Zstd AccountEncoding = "base64+zstd"
)

type TransactionEncoding string

const (
	TransactionEncodingBinary     TransactionEncoding = "binary"
	TransactionEncodingBase64     TransactionEncoding = "base64"
	TransactionEncodingBase58     TransactionEncoding = "base58"
	TransactionEncodingJson       TransactionEncoding = "json"
	TransactionEncodingJsonParsed TransactionEncoding = "jsonParsed"
)

type DataSlice struct {
	Offset uint64 `json:"offset"`
	Length uint64 `json:"length"`
}

type AccountInfo struct {
	Lamports   uint64 `json:"lamports"`
	Owner      string `json:"owner"`
	RentEpoch  uint64 `json:"rentEpoch"`
	Data       any    `json:"data"`
	Executable bool   `json:"executable"`
}

type TokenAccountBalance struct {
	Amount         string `json:"amount"`
	Decimals       uint8  `json:"decimals"`
	UIAmountString string `json:"uiAmountString"`
}

type ReturnData struct {
	ProgramId string `json:"programId"`
	Data      any    `json:"data"`
}

type RewardType string

const (
	RewardTypeFee     RewardType = "Fee"
	RewardTypeRent    RewardType = "Rent"
	RewardTypeVoting  RewardType = "Voting"
	RewardTypeStaking RewardType = "Staking"
)

type Reward struct {
	Pubkey       string      `json:"pubkey"`
	Lamports     int64       `json:"lamports"`
	PostBalances uint64      `json:"postBalance"`
	RewardType   *RewardType `json:"rewardType"`
	Commission   *uint8      `json:"commission"`
}


type SimulateTransactionValueInnerInstruction struct {
	Index        uint64       `json:"index"`
	Instructions []SimulateTransactionValueInstruction        `json:"instructions"`
}



type SimulateTransactionValueInstruction struct {
	Parsed      json.RawMessage  `json:"parsed,omitempty"` // 用于存储动态 JSON 数据
	Accounts    []string         `json:"accounts,omitempty"`
	Data        string           `json:"data,omitempty"`
	Program     string           `json:"program"`
	ProgramId   string           `json:"programId"`
	StackHeight int              `json:"stackHeight"`
}


