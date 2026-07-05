package itrations

import "testing"

func TestRepeat(t *testing.T){
	get:= repeat("a")
	want:= "aaaaa"

	if(get!= want){
		t.Errorf("got %q but wanted %q", get ,want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop(){repeat("a")

	}
	
}