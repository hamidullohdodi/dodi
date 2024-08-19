package models

type UserID struct {
	Id string `json:"id"`
}

type Params struct {
	From     string `json:"from"`
	Password string `json:"password"`
	To       string `json:"to"`
	Massage  string `json:"massage"`
	Code     string `json:"code"`
}

type Voidd struct{}

type UserResp struct {
	Id          string `json:"id"`
	FirstName    string `json:"first_name"`
	Email       string `json:"email"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	CreatedAt   string `json:"created_at"`
}

type EditProfileReq struct {
	Id          string `json:"id"`
	FirstName    string `json:"first_name"`
	Email       string `json:"email"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
}

type PrimaryKey struct {
	UserID string `json:"user_id"`
}

type BudgetTotalItem struct {
	CategoryID  string `json:"category_id"`
	TotalAmount string `json:"total_amount"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Period      string `json:"period"`
}

type GetUserBudgetResponse struct {
	Results []BudgetTotalItem `json:"results"`
}

type GetUserMoneyRequest struct {
	UserID    string `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type GetUserMoneyResponse struct {
	CategoryID  string `json:"category_id"`
	TotalAmount string `json:"total_amount"`
	Time        string `json:"time"`
}

type GoalProgressRequest struct {
	UserID    string `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type GoalProgressItem struct {
	Status           string `json:"status"`
	TargetAmountSum  string `json:"target_amount_sum"`
	CurrentAmountSum string `json:"current_amount_sum"`
	TotalAmount      string `json:"total_amount"`
}

type GoalProgressResponse struct {
	Results []GoalProgressItem `json:"results"`
}

type UpdateProfileReq struct {
	Id          string `json:"id"`
	FirstName    string `json:"first_name"`
	Email       string `json:"email"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
}

type ProfileResp struct {
	Id          string `json:"id"`
	FirstName    string `json:"first_name"`
	Email       string `json:"email"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	UpdatedAt   string `json:"updated_at"`
}

type ChangePasswordReq struct {
	Id              string `json:"id"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type ChangePasswordReqBody struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type Setting struct {
	PrivacyLevel string `json:"privacy_level"`
	Notification string `json:"notification"`
	Language     string `json:"language"`
	Theme        string `json:"theme"`
}

type Word struct {
	Word            string `json:"word"`
	Translation     string `json:"translation"`
	PartOfSpeech    string `json:"part_of_speech"`
	ExampleSentence string `json:"example_sentence"`
}

type Void struct{}

type AccountResp struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type CreateAccountReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetAccountReq struct {
	Id string `json:"id"`
}

type UpdateAccountReq struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteAccountReq struct {
	Id string `json:"id"`
}

type ListAccountsReq struct{}

type ListAccountsResp struct {
	Accounts []AccountResp `json:"accounts"`
}

type TransactionResp struct {
	Id          string `json:"id"`
	AccountId   string `json:"account_id"`
	CategoryId  string `json:"category_id"`
	Amount      string `json:"amount"`
	Date        string `json:"date"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type CreateTransactionReq struct {
	AccountId   string `json:"account_id"`
	CategoryId  string `json:"category_id"`
	Amount      string `json:"amount"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type GetTransactionReq struct {
	Id string `json:"id"`
}

type UpdateTransactionReq struct {
	Id          string `json:"id"`
	AccountId   string `json:"account_id"`
	CategoryId  string `json:"category_id"`
	Amount      string `json:"amount"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type DeleteTransactionReq struct {
	Id string `json:"id"`
}

type ListTransactionsReq struct {
	AccountId  string `json:"account_id"`
	CategoryId string `json:"category_id"`
}

type ListTransactionsResp struct {
	Transactions []TransactionResp `json:"transactions"`
}

type CategoryResp struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type CreateCategoryReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryReq struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteCategoryReq struct {
	Id string `json:"id"`
}

type ListCategoriesReq struct{}

type ListCategoriesResp struct {
	Categories []CategoryResp `json:"categories"`
}

type BudgetResp struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	CreatedAt string `json:"created_at"`
}

type CreateBudgetReq struct {
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetBudgetReq struct {
	Id string `json:"id"`
}

type UpdateBudgetReq struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type DeleteBudgetReq struct {
	Id string `json:"id"`
}

type ListBudgetsReq struct{}

type ListBudgetsResp struct {
	Budgets []BudgetResp `json:"budgets"`
}

type GoalResp struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	TargetAmount  string `json:"target_amount"`
	CurrentAmount string `json:"current_amount"`
	Deadline      string `json:"deadline"`
	CreatedAt     string `json:"created_at"`
}

type CreateGoalReq struct {
	Name         string `json:"name"`
	TargetAmount string `json:"target_amount"`
	Deadline     string `json:"deadline"`
}

type GetGoalReq struct {
	Id string `json:"id"`
}

type UpdateGoalReq struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	TargetAmount  string `json:"target_amount"`
	CurrentAmount string `json:"current_amount"`
	Deadline      string `json:"deadline"`
}

type DeleteGoalReq struct {
	Id string `json:"id"`
}

type ListGoalsReq struct{}

type ListGoalsResp struct {
	Goals []GoalResp `json:"goals"`
}

type GenerateReportReq struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ReportResp struct {
	Report      string `json:"report"`
	GeneratedAt string `json:"generated_at"`
}
