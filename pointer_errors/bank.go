package banksError

import "fmt"

// import "fmt"


type Bank struct{
	balance Bitcoin
}

func (b *Bank) Deposit(amount Bitcoin ){
	b.balance+=amount
	// fmt.Printf("address of balance in Deposit is %p \n", &b.balance)
}
func (b *Bank) Balance() Bitcoin {
	return b.balance
}

type Bitcoin int 

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC",b)
}