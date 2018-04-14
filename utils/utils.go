package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"unicode"
)

func CompileCommand(data map[string]string) (string, error) {

	var command string = "restic "

	return command, nil

}

func SetPassFile(path string, secret string) (string, error) {
	f, err := os.Create(path)
	Check(err, "")
	defer f.Close()
	_, err = f.Write([]byte(secret))
	f.Sync()
	passFilePath, _ := filepath.Abs(path)
	return passFilePath, err
}

func CheckProgExists(name string) (bool, error) {
	ret := true
	path, err := exec.LookPath(name)
	if err != nil {
		fmt.Printf("didn't find '" + name + "' executable\n")
		ret = false
	} else {
		fmt.Printf("'"+name+"' executable is in '%s'\n", path)
	}
	return ret, err
}

func Check(err error, typ string) error {
	if err != nil {
		log.Println(err)
		if typ == "fatal" {
			panic(err)
		}
		return err
	}

	return nil
}

func CheckFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
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
