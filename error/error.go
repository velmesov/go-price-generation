package error

import "log"

// CheckNil Проверка на ошибку
func CheckNil(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
