package banksError

import "fmt"


type Bank struct{
	balance int
}

func (b *Bank) Deposit(amount int ){
	b.balance+=amount
	fmt.Printf("address of balance in Deposit is %p \n", &b.balance)
}
func (b *Bank) Balance() int {
	return b.balance
}