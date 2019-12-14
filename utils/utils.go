package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"unicode"
)

func GetDataPath() string {
	path := os.Getenv("HOME")

	return path + "/.restic-gui/"
}

func Init() (bool, error) {
	ret := true
	err := os.MkdirAll(GetDataPath(), os.ModePerm)
	if err != nil {
		fmt.Println("Cannot create hidden directory.") 
		ret = false
	}

	return ret, err
}

func CompileCommand(data map[string]string) (string, error) {
	var command string = "restic "
	
	return command, nil
}

func SetFile(path string, content string) (string, error) {
	f, err := os.Create(path)
	Check(err, "")
	defer f.Close()
	_, err = f.Write([]byte(content))
	f.Sync()
	filePath, _ := filepath.Abs(path)
	
	return filePath, err
}

func CheckProgExists(name string) (bool, error) {
	ret := true
	_, err := exec.LookPath(name)
	if err != nil {
		fmt.Printf("didn't find '" + name + "' executable\n")
		ret = false
	}
	
	return ret, err
}

func getPath() (string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	return pwd
}

func Check(err error, typ string) error {
	if err != nil {
		if err != nil {
			logFile, err := os.OpenFile(GetDataPath() + "error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			log.Fatalf("error opening file: %v", err)
			defer logFile.Close()
			log.SetOutput(logFile)
			log.Println(err)
		}
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

func SetEnvVars(c map[string]string) {
	os.Setenv("RESTIC_PASSWORD", c["passwd"])
	os.Setenv("RESTIC_REPOSITORY", c["destination"])
}

func SliceIndex(f []string, t string) int {
	for i, val := range f {
		if val == t {
			return i
		}
	}
	
	return -1
}

func OpenBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Run()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Run()
	case "darwin":
		err = exec.Command("open", url).Run()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
