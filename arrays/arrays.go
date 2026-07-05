package arrays

func sum(num []int ) (val int) {
	for _,v:=range num{
		val+=v
	}
	return val
}

func SumAll(numberOfSlice ...[]int)[]int{
	ans:= make([]int , len(numberOfSlice))
	for i,val:=range numberOfSlice{
		ans[i]=sum(val)
	}
	return ans
}