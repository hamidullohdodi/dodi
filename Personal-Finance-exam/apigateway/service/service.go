package service

import (
	"api/config"
	pbb "api/genproto/account"
	pbB "api/genproto/budget"
	pbC "api/genproto/category"
	pbG "api/genproto/goal"
	"api/genproto/notification"
	pbT "api/genproto/transaction"
	"fmt"

	pb "api/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	UserService() pb.UserServiceClient
	GoalService() pbG.GoalServiceClient
	CategoryService() pbC.CategoryServiceClient
	BudgetService() pbB.BudgetServiceClient
	AccountService() pbb.AccountServiceClient
	TransactionService() pbT.TransactionServiceClient
	NotificationService() notification.NotificationServiceClient
}

type serviceManagerImpl struct {
	userClient         pb.UserServiceClient
	goalClient         pbG.GoalServiceClient
	categoryClient     pbC.CategoryServiceClient
	budgetClient       pbB.BudgetServiceClient
	accountClient      pbb.AccountServiceClient
	transactionClient  pbT.TransactionServiceClient
	notificationClient notification.NotificationServiceClient
}

func (s *serviceManagerImpl) UserService() pb.UserServiceClient {
	return s.userClient
}

func (s *serviceManagerImpl) NotificationService() notification.NotificationServiceClient {
	return s.notificationClient
}

func (s *serviceManagerImpl) GoalService() pbG.GoalServiceClient {
	return s.goalClient
}

func (s *serviceManagerImpl) AccountService() pbb.AccountServiceClient {
	return s.accountClient
}

func (s *serviceManagerImpl) BudgetService() pbB.BudgetServiceClient {
	return s.budgetClient
}

func (s *serviceManagerImpl) TransactionService() pbT.TransactionServiceClient {
	return s.transactionClient
}

func (s *serviceManagerImpl) CategoryService() pbC.CategoryServiceClient {
	return s.categoryClient
}

func NewServiceManager(cfg config.Config) (ServiceManager, error) {
	userTarget := fmt.Sprintf("%s:%s", cfg.UserHost, cfg.UserPort)
	userConn, err := grpc.NewClient(userTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	budgettingTarget := fmt.Sprintf("%s:%s", cfg.BudgetingHost, cfg.BudgetingPort)
	deviceConn, err := grpc.NewClient(budgettingTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &serviceManagerImpl{
		userClient:         pb.NewUserServiceClient(userConn),
		goalClient:         pbG.NewGoalServiceClient(deviceConn),
		categoryClient:     pbC.NewCategoryServiceClient(deviceConn),
		budgetClient:       pbB.NewBudgetServiceClient(deviceConn),
		accountClient:      pbb.NewAccountServiceClient(deviceConn),
		transactionClient:  pbT.NewTransactionServiceClient(deviceConn),
		notificationClient: notification.NewNotificationServiceClient(deviceConn),
	}, nil
}
