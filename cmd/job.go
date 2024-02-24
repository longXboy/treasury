package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/longxboy/treasure/ent"
	"github.com/longxboy/treasure/ent/request"

	"github.com/longxboy/treasure/pkg/eth"
)

func runJob() {
	// recover from the last program crash
	// assume that only this single-threaded job can operate on company's eth account,so the pending request's nonce is the latest
	req, err := dbClient.Request.Query().Where(request.Status(RequestStatusApproved), request.Executed(false), request.NonceGTE(0)).First(context.Background())
	if err != nil {
		if !ent.IsNotFound(err) {
			fmt.Println("Query error:", err)
			log.Fatal(err)
		}
	}
	if req != nil {
		TranferTokenTillSuccess(context.Background(), req)
		err = dbClient.Request.Update().Where(request.ID(req.ID)).SetExecuted(true).Exec(context.Background())
		if err != nil {
			fmt.Println("Update error:", err)
			log.Fatal(err)
		}
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		req, err = dbClient.Request.Query().Where(request.Status(RequestStatusApproved), request.Executed(false), request.NonceEQ(-1)).Order(ent.Asc(request.FieldID)).First(context.Background())
		if err != nil {
			if !ent.IsNotFound(err) {
				fmt.Println("Query error:", err)
			}
			continue
		}
		nonce, err := eth.GetAccountNonce(context.Background())
		if err != nil {
			fmt.Println("GetAccountNonce error:", err)
			continue
		}
		req.Nonce = int64(nonce)
		TranferTokenTillSuccess(context.Background(), req)
		for {
			err = dbClient.Request.Update().Where(request.ID(req.ID)).SetExecuted(true).Exec(context.Background())
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
		}
	}
}

func TranferTokenTillSuccess(ctx context.Context, req *ent.Request) {
	ticker := time.NewTicker(time.Millisecond * 200)
	defer ticker.Stop()
	for {
		<-ticker.C
		if req.TxHash != "" {
			success, notfound := waitTxDone(req.TxHash)
			if success {
				return
			}
			if !notfound {
				// if the original tx failed, we need to get new nonce
				// if the original tx not found, we use the original nonce
				var err error
				nonce, err := eth.GetAccountNonce(ctx)
				if err != nil {
					continue
				}
				req.Nonce = int64(nonce)
			}
		}

		tx, err := eth.SignTx(ctx, req.Recipient, req.Amount, uint64(req.Nonce))
		if err != nil {
			fmt.Println("eth.SignTx error:", err)
			continue
		}
		// set new tx hash and nonce for request
		err = dbClient.Request.Update().Where(request.ID(req.ID)).SetTxHash(tx.Hash().Hex()).SetNonce(req.Nonce).Exec(ctx)
		if err != nil {
			fmt.Println("Update request error:", err)
			continue
		}
		req.TxHash = tx.Hash().Hex()
		_, err = eth.SendTransaction(ctx, tx)
		if err != nil {
			fmt.Println("SendTransaction error:", err)
			continue
		}
	}
}

func waitTxDone(tx string) (success bool, notfound bool) {
	ticker := time.NewTicker(time.Millisecond * 200)
	defer ticker.Stop()
	for {
		<-ticker.C
		success, pending, err := eth.GetTransactionStatus(context.Background(), tx)
		if err != nil {
			if err.Error() == "not found" {
				return false, true
			}
			fmt.Println("GetTransactionStatus error:", err)
			continue
		}
		if pending {
			continue
		}
		return success, false
	}
}
