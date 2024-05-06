package service

import (
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
	"github.com/google/uuid"
	"log"
)

func CreatePortfolio(uid, aid string, riskType string, account *model.Account) error {
	p := &model.Portfolio{
		ID:       uuid.New().String(),
		UserID:   uid,
		RiskType: riskType,
	}

	err := p.CreatePortfolio()
	if err != nil {
		return err
	}

	err = OrderStock(uid, p.ID, aid, riskType, account)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func GetPortfolioList(uid string) ([]model.Portfolio, error) {
	p := &model.Portfolio{
		UserID: uid,
	}
	return p.GetPortfolioList()
}

func GetPortfolioByID(pid string) (*model.Portfolio, error) {
	p := &model.Portfolio{
		ID:     pid,
		UserID: pid,
	}

	return p.GetPortfolioByID()
}
