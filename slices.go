package main

import "fmt"

var (
	sliceList = []string{"$9.50", "$8.00", "$10.20", "$7.40"}
	// sliceList = make([]string, 5, 5)
)

func main() {
	//Insert Code here
	// sliceList[0] = "$9.50"
	// sliceList[1] = "$8.00"
	// sliceList[2] = "$10.20"
	// sliceList[3] = "$7.40"

	fmt.Println(len(sliceList))
	fmt.Println(cap(sliceList))

	fmt.Println(sliceList[3-1])
	sliceList[3-1] = "9.80"
	fmt.Println(sliceList[3-1])

	sliceList = append(sliceList, "$8.40")
	sliceList = append(sliceList, "9.40")
	sliceList = append(sliceList, "$7.20")

	fmt.Println(sliceList)
	fmt.Println(len(sliceList))
	fmt.Println(cap(sliceList))

	sliceNew := sliceList[3-1:]
	fmt.Println(len(sliceNew))
	fmt.Println(cap(sliceNew))

}
