

package main 

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	/*
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contents)
	}
	*/
	// 简写 err只在判断内有效
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n",contents)
	}

	fmt.Println(
		grade(10),
		grade(50),
		grade(70),
		grade(90),
		grade(100),
		// grade(200)
	)
}

func grade(score int) string {
	result := ""
	switch {
		case score < 0 || score > 100:
			panic(fmt.Sprintf("Wrong score: %d", score))
		case score < 60:
			result = "F"
		case score < 70:
			result = "C"
		case score < 90:
			result = "B"
		case score <= 100:
			result = "A"
	}
	return result
}