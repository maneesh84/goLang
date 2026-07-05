package arrays

import "testing"

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