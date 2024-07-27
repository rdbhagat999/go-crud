package helper

import (
	"fmt"
	"os"
)

func ErrorPanic(err error) {
	if err != nil {
		fmt.Println("ErrorPanic start")
		panic(err)
	}
}

func GetEnvVariable(envVarName string) string {
	return os.Getenv(envVarName)
}
