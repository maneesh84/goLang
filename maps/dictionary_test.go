package maps

import "testing"

func TestDictionary(t *testing.T){
	t.Run("search", func(t *testing.T) {
		dict := Dictionary{"test":"this is first test"}
		got,_ := dict.Search("test")
		want:="this is first test"

		assertkey(t,got,want)
	})
	t.Run("unknown word", func(t *testing.T)  {
		dict := Dictionary{"test":"this is first test"}
		_,err := dict.Search("word")
		want:= "word is not in dictionary"
		if(err==nil){
			t.Fatalf("expect error but got nothing")
		}

		assertkey(t,err.Error(),want)

	})
}

func assertkey(t testing.TB, got string , want string){
	if(got != want){
		t.Errorf("got %q but want %q ", got ,want )
	}
}