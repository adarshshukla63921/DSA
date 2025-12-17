package main

import "fmt"

func generateParanthesis(
	ind int,
	s string,
	opening int,
	closing int,
	list *[]string,
	n int) {
	if opening > n {return}

	if opening==closing && opening+closing == 2*n {
		*list = append(*list, s)
		return
	}
	generateParanthesis(ind+1, s+"(",opening+1,closing, list, n)
	if(opening>closing){
		generateParanthesis(ind+1, s+")",opening,closing+1, list, n)
	}
}
func main() {
	n := 3
	var list []string
	generateParanthesis(0, "", 0, 0, &list, n)
	fmt.Println(list)
}