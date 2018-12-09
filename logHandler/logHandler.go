package logHandler

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const maxFileSizeKb = 100

func WriteLogFile(err error, source, note string, level ...string) {
	var lvl string
	if len(level) > 0 {
		lvl = level[0]
	} else {
		lvl = "ERROR"
	}
	prefix := getPrefix(lvl)
	str := fmt.Sprintln(prefix, err, " In ", source, " || Additional Info: ", note)
	fileName := getLogFile(source)
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()
	//todo: add checking error
	_, _ = file.WriteString(str)
}

func getPrefix(level string) string {
	currentTime := time.Now()
	str := currentTime.Format("2006-01-02\t15:04:05") + " [" + level + "]: "
	return str
}

func getLogFile(source string) string {
	folderName := "./log/"
	fileName := strings.Replace(folderName+source+".log", " ", "", -1)
	checkFolder(folderName)
	if stat, err := os.Stat(fileName); !os.IsNotExist(err) {
		size := stat.Size()
		if size > maxFileSizeKb*1024 {
			fileName = newLogFile(fileName)
		}
	} else {
		//todo: add checking error
		_, _ = os.Create(fileName)
	}
	return fileName
}

func newLogFile(fileName string) string {
	count := 1
	nextFile := fmt.Sprintf("%s.%d", fileName, count)
	for _, err := os.Stat(nextFile); !os.IsNotExist(err); _, err = os.Stat(nextFile) {
		count++
		nextFile = fmt.Sprintf("%s.%d", fileName, count)
	}

	//todo: add checking error
	_ = os.Rename(fileName, nextFile)
	//todo: add checking error
	_, _ = os.Create(fileName)
	return fileName
}

func checkFolder(folderName string) {
	if _, err := os.Stat(folderName); !os.IsNotExist(err) {
		//todo: add checking error
		_ = os.MkdirAll(folderName, os.ModePerm)
	}
}
