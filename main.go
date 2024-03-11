package main
import (
	"fmt"
	"sort"
)


func calculateAverageAge(ages []int) float64 {
    sum := 0
    for _, age := range ages {
        sum += age
    }
    return float64(sum) / float64(len(ages))
}
func findMaxAge(ages []int) int {
    maxAge := ages[0]
    for _, age := range ages {
        if age > maxAge {
            maxAge = age
        }
    }
    return maxAge
}
func findMinAge(ages []int) int {
    minAge := ages[0]
    for _, age := range ages {
        if age < minAge {
            minAge = age
        }
    }
    return minAge
}

func main(){
    ages := []int{20, 49, 30, 38, 89, 43, 34 , 36 , 68 }
    sort.Ints(ages)
    fmt.Println(ages)

    index := sort.SearchInts(ages, 36)
    fmt.Println("Index of 36:", index)

    avgAge := calculateAverageAge(ages)
    fmt.Println("Average Age:", avgAge)

    maxAge := findMaxAge(ages)
    fmt.Println("Maximum Age:", maxAge)

    minAge := findMinAge(ages)
    fmt.Println("Minimum Age:", minAge)
}



