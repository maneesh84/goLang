package maps

import "testing"

func TestDictionary(t *testing.T){
	t.Run("search", func(t *testing.T) {
		dict := Dictionary{"test":"this is first test"}
		got := dict.Search("test")
		want:="this is first test"

		assertkey(t,got,want)
	})
}

func assertkey(t testing.TB, got string , want string){
	if(got != want){
		t.Errorf("got %q but want %q ", got ,want )
	}
}