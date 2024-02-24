package main

import (
	"context"
	"fmt"

	pb "github.com/longxboy/treasure/api/server"
	"github.com/longxboy/treasure/ent"
	"github.com/longxboy/treasure/ent/confirm"
	"github.com/longxboy/treasure/ent/request"
)

const RequestStatusApproved = "approved"
const RequestStatusPending = "pending"
const RequestStatusRejected = "rejected"

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

type server struct {
	pb.UnimplementedTreasuryServer
}

func (s *server) CreateWithdraw(ctx context.Context, in *pb.CreateWithdrawRequest) (*pb.CreateWithdrawReply, error) {
	request, err := dbClient.Request.Create().SetAmount(in.Amount).SetRecipient(in.Recipient).SetStatus(RequestStatusPending).SetNonce(-1).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.CreateWithdrawReply{RequestId: request.ID}, nil
}

func (s *server) GetWithdrawStatus(ctx context.Context, in *pb.GetWithdrawStatusRequest) (*pb.GetWithdrawStatusReply, error) {
	request, err := dbClient.Request.Get(ctx, in.RequestId)
	if err != nil {
		return nil, err
	}
	return &pb.GetWithdrawStatusReply{
		Status:    request.Status,
		Amount:    request.Amount,
		Recipient: request.Recipient,
		TxHash:    request.TxHash,
		Nonce:     request.Nonce,
		Executed:  request.Executed,
	}, nil
}

func (s *server) ConfirmWithdraw(ctx context.Context, in *pb.ConfirmWithdrawRequest) (*pb.ConfirmWithdrawReply, error) {
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return nil, err
	}
	// mysql unique index prevent duplicate confirm
	_, err = tx.Confirm.Create().SetRequestID(in.RequestId).SetManagerID(in.ManagerId).SetApproved(true).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	nums, err := tx.Confirm.Query().Where(confirm.RequestID(in.RequestId)).Count(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}
	if nums >= 2 {
		err = tx.Request.Update().Where(request.ID(in.RequestId)).SetStatus(RequestStatusApproved).Exec(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &pb.ConfirmWithdrawReply{
		Success: true,
	}, nil
}
