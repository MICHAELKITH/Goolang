package main
import (
	"fmt"
	"sort"
)

func main(){

	// variable
	// var nameOne string = "mike"//main.go

	// nameOne := "mikee"

	// var nameTwo = "Mike"

	// fmt.Println(nameOne, nameTwo)


	ages := []int{20, 49, 30, 38, 89, 43}

	sort.Ints(ages)

	fmt.Println(ages)
index := sort.SearchInts(ages, 30)
fmt.Println(index)



}




