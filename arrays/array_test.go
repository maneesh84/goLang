package arrays

import (
	"slices"
	"testing"
)

func TestArray(t *testing.T){
	t.Run("using variabel size slice", func(t * testing.T)  {
		num := []int{1,2,3,4}

		got:= sum(num)
		want := 10

		if(got!= want){
			t.Errorf("got %d but want %d , %v", got ,want ,num)
		}
	})

	

}

func TestAll(t *testing.T){
	got := SumAll([]int{1,2,3},[]int{2,3})
	want:= []int{6,5}

	if !slices.Equal(got,want){
		t.Errorf("got %v but want %v",got ,want)
	}
}