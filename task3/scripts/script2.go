package scripts

import (
	"errors"
	"gorm.io/gorm"
)

type Account struct {
	ID      int
	Balance int
}

func (account *Account) GetBalance() int {
	return account.Balance
}

type Transaction struct {
	ID            int
	FromAccountId int
	ToAccountId   int
	Amount        int
}

func RunTransaction(db *gorm.DB) {
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	account := &[]Account{{ID: 1, Balance: 300}, {ID: 2, Balance: 200}}
	db.Debug().Create(&account)

	err := db.Transaction(func(tx *gorm.DB) error {
		//查询第一个账户
		A, B := Account{ID: 1}, Account{ID: 2}
		db.Debug().First(&A)
		db.Debug().First(&B)
		if A.GetBalance() >= 100 {
			db.Debug().Model(&Account{}).Where("id = ?", A.ID).Update("balance", A.GetBalance()-100)
			db.Debug().Model(&Account{}).Where("id = ?", B.ID).Update("balance", B.GetBalance()+100)
			record := Transaction{FromAccountId: A.ID, ToAccountId: B.ID, Amount: 100}
			db.Debug().Create(&record)
			return nil
		}
		return errors.New("余额不足")
	})

	if err != nil {
		panic(err)
	}
}
