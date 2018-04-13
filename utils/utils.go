package utils

import (
	"log"
	"os"
	"unicode"
)

func CheckErr(err error, typ string) (error) {
	if err != nil {
		log.Println(err)
		if typ == "fatal" {
			panic(err)
		}
		return err
	}

	return nil
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
