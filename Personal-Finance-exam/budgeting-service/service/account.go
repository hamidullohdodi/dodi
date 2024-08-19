package service

import (
	pb "budgeting/genproto/account"
	"budgeting/storage"
	"context"
	"fmt"
	"log/slog"
)

type AccountIService struct {
	pb.UnimplementedAccountServiceServer
	Storage storage.AccountI
	logger  *slog.Logger
}

func NewAccountIService(storage storage.AccountI, logger *slog.Logger) *AccountIService {
	return &AccountIService{
		Storage: storage,
		logger:  logger,
	}
}

func (b *AccountIService) CreateAccount(ctx context.Context, req *pb.CreateAccountReq) (*pb.AccountResp, error) {
	resp, err := b.Storage.CreateAccount(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not create account: %v", err))
		return nil, err
	}
	//_, err = b.Redis.CreateAccount(req)
	//if err != nil {
	//	b.logger.Error(fmt.Sprintf("could not create account: %v", err))
	//	return nil, err
	//}
	return resp, nil
}
func (b *AccountIService) GetAccount(ctx context.Context, req *pb.GetAccountReq) (*pb.AccountResp, error) {
	//resp, err := b.Redis.GetAccount(req)
	//if err != nil {
	//	b.logger.Error(fmt.Sprintf("could not get account: %v", err))
	//	return nil, err
	//}
	//if resp == nil {
	resp, err := b.Storage.GetAccount(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not get account: %v", err))
		return nil, err
	}
	return resp, nil
	//}
	//return resp, nil
}
func (b *AccountIService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountReq) (*pb.AccountResp, error) {
	resp, err := b.Storage.UpdateAccount(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not update account: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *AccountIService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountReq) (*pb.Void1, error) {
	resp, err := b.Storage.DeleteAccount(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not delete account: %v", err))
		return nil, err
	}
	return resp, nil
}
func (b *AccountIService) ListAccounts(ctx context.Context, req *pb.ListAccountsReq) (*pb.ListAccountsResp, error) {
	resp, err := b.Storage.ListAccounts(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("could not list accounts: %v", err))
		return nil, err
	}
	return resp, nil
}
