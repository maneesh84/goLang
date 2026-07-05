package arrays

import "testing"

func TestArray(t *testing.T){
	num := [5]int{1,2,3,4,5}

	got:= sum(num)
	want := 15

	if(got!= want){
		t.Errorf("got %d but want %d , %v", got ,want ,num)
	}

}