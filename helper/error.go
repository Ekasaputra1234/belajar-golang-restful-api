package helper

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Println("======= PANIC =======")
		fmt.Println(err)
		panic(err)
	}
}
