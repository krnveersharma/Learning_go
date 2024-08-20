package main

import "fmt"

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	ans, _ := divide(10, 0)    //_ is ignoring another output
	ans2, err := divide(10, 0) //error is handling in this case
	fmt.Println("Divison of 10 and 4 is", ans)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Divison of 10 and 0 is", ans2)
	}

}
