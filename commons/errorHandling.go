package commons

import (
	"log"
)

func InError(err error) bool {
	if err != nil {
		log.Println("Error: " + err.Error())
		return true
	}
	return false
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
