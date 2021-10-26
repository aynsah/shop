package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func CreateLog(path string, status int, message string, data interface{}) error {
	year, month, day, hour, min, sec := GetTime()

	log_path, _ := filepath.Abs("logs")
	filename := "err-" + year + "-" + month + "-" + day + ".log"

	log := []byte(day + "-" + month + "-" + year + " " + hour + ":" + min + ":" + sec + " => " + path + "\n" + strconv.Itoa(status) + " => " + message + "\n" + fmt.Sprintf("%v", data) + "\n\n")

	f, err := os.OpenFile(log_path+"\\"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write([]byte(log))

	return err
}
