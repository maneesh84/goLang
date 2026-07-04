package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	get := Add(2,5)
	expect := 7

	if get != expect {
		t.Errorf("got %d expected %d", get, expect)
	}
}

func ExampleAdd(){
	sum:= Add(2,3)
	fmt.Print(sum)
	//Output:
	//5
}