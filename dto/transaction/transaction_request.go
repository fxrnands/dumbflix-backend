package dto

import "time"

type TransactionRequest struct {
	StartDate time.Time            		`json:"startDate"`
	DueDate	  time.Time     	   		`json:"dueDate"`
	UserID	 int				 	    `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price	  int						`json:"price"`
}

type UpdateTransactionRequest struct {
	StartDate time.Time            		`json:"startDate" `
	DueDate	  time.Time     	   		`json:"dueDate" `
	UserID	 int				 	    `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	Price	  int						`json:"price"`
}