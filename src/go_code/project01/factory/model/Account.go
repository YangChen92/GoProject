package model

import (
	"fmt"
)

type account struct{
	id string
	pwd string
	balance float64
}

func NewAccount(id string,pwd string,balance float64) *account{
	if len(id) <6 || len(id)10{
		fmt.Println("账号为6~10位")
		return nil
	}
	if len(pwd) != 6{
		fmt.Println("密码为6位")
		return nil
	}

	if balance<20{
		fmt.Println("余额不对")
		return nil
	}

	return &account{
		id : id,
		pwd : pwd,
		balance : balance,
	}
	
}

