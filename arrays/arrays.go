package arrays

func sum(num []int ) (val int) {
	for _,v:=range num{
		val+=v
	}
	return val
}