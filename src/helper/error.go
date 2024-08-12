package helper

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
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

func HandleError(db_error error) string {
	var msg = "Internal server error"

	if errors.Is(db_error, gorm.ErrDuplicatedKey) {
		msg = gorm.ErrDuplicatedKey.Error()
	} else if errors.Is(db_error, gorm.ErrRecordNotFound) {
		msg = gorm.ErrRecordNotFound.Error()
	} else if errors.Is(db_error, gorm.ErrMissingWhereClause) {
		msg = gorm.ErrMissingWhereClause.Error()
	} else if errors.Is(db_error, gorm.ErrInvalidData) {
		msg = gorm.ErrInvalidData.Error()
	} else if errors.Is(db_error, gorm.ErrInvalidField) {
		msg = gorm.ErrInvalidField.Error()
	} else if errors.Is(db_error, gorm.ErrInvalidValueOfLength) {
		msg = gorm.ErrInvalidValueOfLength.Error()
	} else if errors.Is(db_error, gorm.ErrInvalidValue) {
		msg = gorm.ErrInvalidValue.Error()
	} else if errors.Is(db_error, gorm.ErrForeignKeyViolated) {
		msg = gorm.ErrForeignKeyViolated.Error()
	} else if errors.Is(db_error, gorm.ErrNotImplemented) {
		msg = gorm.ErrNotImplemented.Error()
	} else if strings.Contains(db_error.Error(), "crypto/bcrypt") {
		msg = "Username or password does not match"
	}

	return msg
}
