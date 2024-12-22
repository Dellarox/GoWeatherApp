package utils

import (
	"github.com/joho/godotenv"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	CheckError(err)
}
