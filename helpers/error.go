package helpers

import "log"

func LogError(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
	}
}

func FatalOutError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
