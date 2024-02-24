package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	enndpoint = "http://localhost:8545"
	pk        = "564a0460e14e984e1f2a150b424fcb2385b1e6865e5a230474572a83001f9afd"
	account   = "0x9D1C5b4Ff5E29823d686a9360B3508CBD8a26b80"
)

func main() {
	TransferToken("0xd3ac7FDB5e12DDa9cf0a4a3128bf4950c632c36A", 100000000000000000, 4)
	TransferToken("0xd3ac7FDB5e12DDa9cf0a4a3128bf4950c632c36A", 100000000000000000, 4)

	GetTransactionStatus("0xee6e2c97f244d641eb1209fb9ebd13bb97234caa7f6fe207b8f7b604ab2669a3")
}

func GetTransactionStatus(hash string) (bool, error) {
	txHash := common.HexToHash(hash)
	client, err := ethclient.Dial(enndpoint)
	if err != nil {
		return false, err
	}
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return false, err
	}
	return receipt.Status == types.ReceiptStatusSuccessful, nil
}

func GetAccountNonce() (uint64, error) {
	address := common.HexToAddress(account)
	client, err := ethclient.Dial(enndpoint)
	if err != nil {
		return 0, err
	}
	return client.PendingNonceAt(context.Background(), address)
}

func TransferToken(recipient string, amount int64, nonce uint64) (string, error) {
	client, err := ethclient.Dial(enndpoint)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return "", err
	}

	value := big.NewInt(amount) // in wei (0.1 eth)
	gasLimit := uint64(21000)   // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(recipient)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	// nonce too low: next nonce
	// already known
	fmt.Printf("tx sent: %s at nonce %d err:%v\n", signedTx.Hash().Hex(), nonce, err.Error())
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}
