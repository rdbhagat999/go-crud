package helper

import "fmt"

func ErrorPanic(err error) {
	if err != nil {
		fmt.Println("ErrorPanic start")
		panic(err)
	}
}
