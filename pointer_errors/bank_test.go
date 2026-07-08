package banksError

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T){
	bank:= Bank{}

	bank.Deposit(Bitcoin(10))

	got:= bank.Balance()
	fmt.Printf("address of balance in test is %p \n", &bank.balance)
	want:=Bitcoin(10)

	if(got!=want){
		t.Errorf("got %s but want %s",got,want)
	}
}
