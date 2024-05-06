package model

type SingUpRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type LoginRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Pwd    string `json:"pwd" binding:"required"`
}

type DepositRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	Amount    uint   `json:"amount" binding:"required"`
}

type WithdrawRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	Amount    uint   `json:"amount" binding:"required"`
}

type CreateStockRequest struct {
	StockID   string `json:"stock_id" binding:"required"`
	StockName string `json:"stock_name" binding:"required"`
	Price     uint   `json:"stock_price" binding:"required"`
}

type UpdateStockRequest struct {
	StockName string `json:"stock_name" binding:"required"`
	Price     uint   `json:"price" binding:"required"`
}

type CreatePortfolioRequest struct {
	RiskType  string `json:"risk_type" binding:"required"`
	AccountID string `json:"account_id" binding:"required"`
}
