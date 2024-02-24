package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	enndpoint = "http://127.0.0.1:8545"
	pk        = "564a0460e14e984e1f2a150b424fcb2385b1e6865e5a230474572a83001f9afd"
	account   = "0x9D1C5b4Ff5E29823d686a9360B3508CBD8a26b80"
)

var client *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial(enndpoint)
	if err != nil {
		log.Fatal(err)
	}
}

/*func main() {
	TransferToken("0xd3ac7FDB5e12DDa9cf0a4a3128bf4950c632c36A", 100000000000000000, 4)
	TransferToken("0xd3ac7FDB5e12DDa9cf0a4a3128bf4950c632c36A", 100000000000000000, 4)

	GetTransactionStatus("0xee6e2c97f244d641eb1209fb9ebd13bb97234caa7f6fe207b8f7b604ab2669a3")
}*/

func GetTransactionStatus(ctx context.Context, hash string) (success bool, pending bool, err error) {
	txHash := common.HexToHash(hash)
	_, pending, err = client.TransactionByHash(ctx, txHash)
	if err != nil {
		return
	}
	if pending {
		return
	}
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return
	}
	success = receipt.Status == types.ReceiptStatusSuccessful
	return
}

func GetAccountNonce(ctx context.Context) (uint64, error) {
	address := common.HexToAddress(account)

	return client.PendingNonceAt(ctx, address)
}

func SignTx(ctx context.Context, recipient string, amount int64, nonce uint64) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	value := big.NewInt(amount) // in wei
	gasLimit := uint64(21000)   // in units
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(recipient)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func SendTransaction(ctx context.Context, signedTx *types.Transaction) (string, error) {
	err := client.SendTransaction(ctx, signedTx)
	// possible errors:
	// nonce too low: next nonce
	// already known

	fmt.Printf("tx sent: %s  err:%v\n", signedTx.Hash().Hex(), err)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}
