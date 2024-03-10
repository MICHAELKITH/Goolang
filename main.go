package main
import (
	"fmt"
	"sort"
)

func main(){

	ages := []int{20, 49, 30, 38, 89, 43, 34 , 36 , 68 }

	sort.Ints(ages)

	fmt.Println(ages)
index := sort.SearchInts(ages, 36)
fmt.Println(index)

}




