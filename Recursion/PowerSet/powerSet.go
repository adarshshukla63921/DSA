package main
import "fmt"
func generatePowerSet(
	idx int,
	nums *[]int,
	list *[]int,
	ans *[][]int,
	n int){
		if idx == n{
			// make a copy of the current list and append to ans
			// because list will be modified later
			subset := make([]int, len(*list))
			copy(subset,*list)
			*ans = append(*ans, subset)
			return
		}
		*list = append(*list, (*nums)[idx])
		generatePowerSet(idx+1,nums,list,ans,n)
		*list = (*list)[:len(*list)-1] // shrink the slice itself.
		generatePowerSet(idx+1,nums,list,ans,n)
}
func main(){
	 nums := []int{1,2,3}
	 var list []int
	 var ans [][]int
	 generatePowerSet(0,&nums,&list,&ans,len(nums))
	 fmt.Println(ans)
}