package storage

import (
	"budgeting/genproto/account"
	"budgeting/genproto/budget"
	"budgeting/genproto/category"
	"budgeting/genproto/goal"
	"budgeting/genproto/notification"
	pb "budgeting/genproto/transaction"
)

type StorageI interface {
	AccountI() AccountI
	TransactionI() TransactionI
	CategoryI() CategoryI
	Budget() Budget
	Goal() Goal
}

type Redis interface {
	CreateNotification(req *notification.CreateNotificationReq) (*notification.NotificationResp, error)
	GetNotification(req *notification.GetNotificationReq) (*notification.NotificationResp, error)
	UpdateNotification(req *notification.UpdateNotificationReq) (*notification.NotificationResp, error)
	DeleteNotification(req *notification.DeleteNotificationReq) (*notification.DeleteNotificationResp, error)
}
	
type AccountI interface {
	CreateAccount(req *account.CreateAccountReq) (*account.AccountResp, error)
	GetAccount(req *account.GetAccountReq) (*account.AccountResp, error)
	UpdateAccount(req *account.UpdateAccountReq) (*account.AccountResp, error)
	DeleteAccount(req *account.DeleteAccountReq) (*account.Void1, error)
	ListAccounts(req *account.ListAccountsReq) (*account.ListAccountsResp, error)
}

type TransactionI interface {
	CreateTransaction(req *pb.CreateTransactionReq) (*pb.TransactionResp, error)
	GetTransaction(req *pb.GetTransactionReq) (*pb.TransactionResp, error)
	UpdateTransaction(req *pb.UpdateTransactionReq) (*pb.TransactionResp, error)
	DeleteTransaction(req *pb.DeleteTransactionReq) (*pb.Void4, error)
	ListTransactions(req *pb.ListTransactionsReq) (*pb.ListTransactionsResp, error)
}

type CategoryI interface {
	CreateCategory(req *category.CreateCategoryReq) (*category.CategoryResp, error)
	UpdateCategory(req *category.UpdateCategoryReq) (*category.CategoryResp, error)
	DeleteCategory(req *category.DeleteCategoryReq) (*category.Void2, error)
	ListCategories(req *category.ListCategoriesReq) (*category.ListCategoriesResp, error)
}

type Budget interface {
	CreateBudget(req *budget.CreateBudgetReq) (*budget.BudgetResp, error)
	GetBudget(req *budget.GetBudgetReq) (*budget.BudgetResp, error)
	UpdateBudget(req *budget.UpdateBudgetReq) (*budget.BudgetResp, error)
	DeleteBudget(req *budget.DeleteBudgetReq) (*budget.Void, error)
	ListBudgets(req *budget.ListBudgetsReq) (*budget.ListBudgetsResp, error)
}

type Goal interface {
	CreateGoal(req *goal.CreateGoalReq) (*goal.GoalResp, error)
	GetGoal(req *goal.GetGoalReq) (*goal.GoalResp, error)
	UpdateGoal(req *goal.UpdateGoalReq) (*goal.GoalResp, error)
	DeleteGoal(req *goal.DeleteGoalReq) (*goal.Void3, error)
	ListGoals(req *goal.ListGoalsReq) (*goal.ListGoalsResp, error)

	GetUserSpending(rep *goal.GetUserMoneyRequest) (*goal.GetUserMoneyResponse, error)
	GetUserIncome(rep *goal.GetUserMoneyRequest) (*goal.GetUserMoneyResponse, error)
	GetGoalReportProgress(rep *goal.GoalProgressRequest) (*goal.GoalProgressResponse, error)
	GetBudgetSummary(rep *goal.UserId) (*goal.GetUserBudgetResponse, error)
}
