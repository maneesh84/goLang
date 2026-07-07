package banksError

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T){
	bank:= Bank{}

	bank.Deposit(10)

	got:= bank.Balance()
	fmt.Printf("address of balance in test is %p \n", &bank.balance)
	want:= 10

	if(got!=want){
		t.Errorf("got %d but want %d",got,want)
	}
}
