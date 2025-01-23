package client

import (
	"context"
	"encoding/base64"
	"fmt"
	"encoding/json"

	"github.com/qimeila/solana-go-sdk/rpc"
	"github.com/qimeila/solana-go-sdk/types"
)

type SimulateTransaction struct {
	Err               any
	Logs              []string
	Accounts          []*AccountInfo
	ReturnData        *ReturnData
	UnitConsumed      *uint64
	InnerInstructions []SimulateTransactionValueInnerInstruction
}

type SimulateTransactionValueInnerInstruction struct {
	Index        uint64
	Instructions []SimulateTransactionValueInstruction
}



type SimulateTransactionValueInstruction struct {
	Parsed      map[string]interface{}
	Accounts    []string
	Data        string
	Program     string
	ProgramId   string
	StackHeight int
}

type SimulateTransactionConfig struct {
	SigVerify              bool
	Commitment             rpc.Commitment
	ReplaceRecentBlockhash bool
	Addresses              []string
	InnerInstructions      bool
}

func (c SimulateTransactionConfig) toRpc() rpc.SimulateTransactionConfig {
	var accounts *rpc.SimulateTransactionConfigAccounts
	if len(c.Addresses) > 0 {
		accounts = &rpc.SimulateTransactionConfigAccounts{
			Encoding:  rpc.AccountEncodingBase64,
			Addresses: c.Addresses,
		}
	}
	return rpc.SimulateTransactionConfig{
		Encoding:               rpc.SimulateTransactionEncodingBase64,
		SigVerify:              c.SigVerify,
		Commitment:             c.Commitment,
		ReplaceRecentBlockhash: c.ReplaceRecentBlockhash,
		Accounts:               accounts,
		InnerInstructions:      c.InnerInstructions,
	}
}

func (c *Client) SimulateTransaction(ctx context.Context, tx types.Transaction) (SimulateTransaction, error) {
	rawTx, err := tx.Serialize()
	if err != nil {
		return SimulateTransaction{}, fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.SimulateTransactionValue]], error) {
			return c.RpcClient.SimulateTransactionWithConfig(
				ctx,
				base64.StdEncoding.EncodeToString(rawTx),
				SimulateTransactionConfig{}.toRpc(),
			)
		},
		convertSimulateTransaction,
	)
}

func (c *Client) SimulateTransactionWithConfig(ctx context.Context, tx types.Transaction, cfg SimulateTransactionConfig) (SimulateTransaction, error) {
	rawTx, err := tx.Serialize()
	if err != nil {
		return SimulateTransaction{}, fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.SimulateTransactionValue]], error) {
			return c.RpcClient.SimulateTransactionWithConfig(
				ctx,
				base64.StdEncoding.EncodeToString(rawTx),
				cfg.toRpc(),
			)
		},
		convertSimulateTransaction,
	)
}

func (c *Client) SimulateTransactionAndContext(ctx context.Context, tx types.Transaction) (rpc.ValueWithContext[SimulateTransaction], error) {
	rawTx, err := tx.Serialize()
	if err != nil {
		return rpc.ValueWithContext[SimulateTransaction]{}, fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.SimulateTransactionValue]], error) {
			return c.RpcClient.SimulateTransactionWithConfig(
				ctx,
				base64.StdEncoding.EncodeToString(rawTx),
				SimulateTransactionConfig{}.toRpc(),
			)
		},
		convertSimulateTransactionAndContext,
	)
}

func (c *Client) SimulateTransactionAndContextWithConfig(ctx context.Context, tx types.Transaction, cfg SimulateTransactionConfig) (rpc.ValueWithContext[SimulateTransaction], error) {
	rawTx, err := tx.Serialize()
	if err != nil {
		return rpc.ValueWithContext[SimulateTransaction]{}, fmt.Errorf("failed to serialize tx, err: %v", err)
	}
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[rpc.SimulateTransactionValue]], error) {
			return c.RpcClient.SimulateTransactionWithConfig(
				ctx,
				base64.StdEncoding.EncodeToString(rawTx),
				cfg.toRpc(),
			)
		},
		convertSimulateTransactionAndContext,
	)
}

func convertSimulateTransaction(v rpc.ValueWithContext[rpc.SimulateTransactionValue]) (SimulateTransaction, error) {
	var accountInfos []*AccountInfo
	if v.Value.Accounts != nil {
		accountInfos = make([]*AccountInfo, 0, len(v.Value.Accounts))
		for _, r := range v.Value.Accounts {
			if r == nil {
				accountInfos = append(accountInfos, nil)
				continue
			}
			accountInfo, err := convertAccountInfo(*r)
			if err != nil {
				return SimulateTransaction{}, err
			}
			accountInfos = append(accountInfos, &accountInfo)
		}
	}

	var returnData *ReturnData
	if v := v.Value.ReturnData; v != nil {
		d, err := convertReturnData(*v)
		if err != nil {
			return SimulateTransaction{}, fmt.Errorf("failed to process return data, err: %v", err)
		}
		returnData = &d
	}

	innerInstructions, err := convertSimulateTransactionValueInnerInstruction(v.Value.SimulateTransactionValueInnerInstructions)
	if err != nil {
		return SimulateTransaction{}, fmt.Errorf("failed to convert inner instructions: %v", err)
	}

	return SimulateTransaction{
		Err:               v.Value.Err,
		Logs:              v.Value.Logs,
		Accounts:          accountInfos,
		ReturnData:        returnData,
		UnitConsumed:      v.Value.UnitConsumed,
		InnerInstructions: innerInstructions,
	}, nil
}

func convertSimulateTransactionValueInnerInstruction(innerInstructions []rpc.SimulateTransactionValueInnerInstruction) ([]SimulateTransactionValueInnerInstruction, error) {
	if innerInstructions == nil {
		return nil, nil
	}

	result := make([]SimulateTransactionValueInnerInstruction, len(innerInstructions))
	for i, innerInstruction := range innerInstructions {
		instructions := make([]SimulateTransactionValueInstruction, len(innerInstruction.Instructions))
		for j, instruction := range innerInstruction.Instructions {

			if len(instruction.Parsed) > 0 {
				var parsedData map[string]interface{}
				err := json.Unmarshal(instruction.Parsed, &parsedData)
				if err != nil {
					fmt.Printf("Error parsing 'parsed' field: %v\n", err)
				}
				instructions[j] = SimulateTransactionValueInstruction{
					Parsed:         parsedData,
					Program:        instruction.Program,
					ProgramId:      instruction.ProgramId,
					StackHeight:    instruction.StackHeight,
				}
			}else{
				instructions[j] = SimulateTransactionValueInstruction{
					Accounts:       instruction.Accounts,
					Data:           instruction.Data,
					ProgramId:      instruction.ProgramId,
					StackHeight:    instruction.StackHeight,
				}
			}
		}

		result[i] = SimulateTransactionValueInnerInstruction{
			Index:        innerInstruction.Index,
			Instructions: instructions,
		}
	}
	return result, nil
}

func convertSimulateTransactionAndContext(v rpc.ValueWithContext[rpc.SimulateTransactionValue]) (rpc.ValueWithContext[SimulateTransaction], error) {
	simulateTrasaction, err := convertSimulateTransaction(v)
	if err != nil {
		return rpc.ValueWithContext[SimulateTransaction]{}, err
	}
	return rpc.ValueWithContext[SimulateTransaction]{
		Context: v.Context,
		Value:   simulateTrasaction,
	}, nil
}
