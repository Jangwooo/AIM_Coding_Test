package service

import (
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
	"github.com/google/uuid"
	"log"
	"sort"
)

func AddStock(req *model.CreateStockRequest) error {
	s := model.Stock{
		ID:        req.StockID,
		StockName: req.StockName,
		Price:     req.Price,
	}

	return s.CreateStock()
}

func GetStockByCode(sid string) (*model.Stock, error) {
	s := &model.Stock{
		ID: sid,
	}

	s, err := s.GetStockByStockCode()

	return s, err
}

func UpdateStock(req *model.UpdateStockRequest, sid string) error {
	s := model.Stock{
		ID:        sid,
		StockName: req.StockName,
		Price:     req.Price,
	}

	return s.UpdateStock()
}

func DeleteStock(sid string) error {
	s := model.Stock{
		ID: sid,
	}

	return s.DeleteStock()
}

func GetStockList() ([]model.Stock, error) {
	return model.GetStockList()
}

func OrderStock(uid, pid, aid string, riskType string, account *model.Account) error {
	sl, err := GetStockList()
	if err != nil {
		return err
	}

	var availableMoney uint

	switch riskType {
	case "high":
		availableMoney = account.Balance
	case "low":
		availableMoney = account.Balance / 2
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i].Price > sl[j].Price
	})

	stockQuantity := map[string]uint{}
	stockPrice := map[string]uint{}

	for _, stock := range sl {
		purchaseAmount := availableMoney / stock.Price // Calculate potential purchase quantity
		if purchaseAmount < 1 {                        // Skip if stock price exceeds available money
			continue
		}

		// Attempt to withdraw and purchase the desired quantity
		err := Withdraw(uid, aid, stock.Price*purchaseAmount)
		if err != nil {
			log.Println(err.Error())
			break // Exit on withdrawal error
		}

		// Update information if purchase successful
		stockPrice[stock.ID] = stock.Price
		availableMoney -= stock.Price * purchaseAmount
		stockQuantity[stock.ID] += purchaseAmount
	}

	stockTransactionList := make([]model.StockTransaction, 0)
	portfolioItemList := make([]model.PortfolioItem, 0)

	for k, v := range stockQuantity {
		temp := model.StockTransaction{
			ID:       uuid.New().String(),
			UserID:   uid,
			StockID:  k,
			Quantity: v,
			Price:    stockPrice[k],
		}

		ptemp := model.PortfolioItem{
			ID:          uuid.New().String(),
			PortfolioID: pid,
			StockCode:   k,
			Quantity:    v,
		}
		stockTransactionList = append(stockTransactionList, temp)
		portfolioItemList = append(portfolioItemList, ptemp)
	}

	err = model.CreateStockTransaction(&stockTransactionList)
	if err != nil {
		return err
	}
	err = model.CreatePortfolioItem(&portfolioItemList)
	if err != nil {
		return err
	}

	return nil
}
