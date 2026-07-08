package banksError

import (

	"testing"
)

func TestBank(t *testing.T){
	t.Run("deposit",func(t *testing.T) {
		bank:= Bank{}

		bank.Deposit(Bitcoin(10))
		bank.Withdraw(Bitcoin(5))
		got:= bank.Balance()
		// fmt.Printf("address of balance in test is %p \n", &bank.balance)
		want:=Bitcoin(5)

		if(got!=want){
			t.Errorf("got %s but want %s",got,want)
		}
	})

	t.Run("widthdraw", func(t *testing.T) {
		bank:=Bank{balance: 20}
		err:=bank.Withdraw(30)
		want:= "insufficient balance"

		if(err==nil){
			t.Errorf("expected error but got nil")
		}

		if(err.Error()!=want){
			t.Errorf("got %q but want %q",err.Error(),want)
		}

	})
}
