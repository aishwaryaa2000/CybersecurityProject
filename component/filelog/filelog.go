package filelog

import (
	"os"
	"time"
)

func WriteUserLog(username, action string) {
	f, err := os.OpenFile("userlog.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	currentTime := time.Now()

	defer f.Close()
	text := username + action + " at " + currentTime.Format("2006-01-02 3:4:5 PM") + "\n"
	_, errr := f.WriteString(text)
	if errr != nil {
		panic(err)
	}

	f.Close()
}

func WriteFileLog(username, fileName, action string) {
	f, err := os.OpenFile("filelog.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	currentTime := time.Now()

	defer f.Close()
	text := username + " " + action + " " + fileName + " at " + currentTime.Format("2006-01-02 3:4:5 PM") + "\n"
	_, errr := f.WriteString(text)
	if errr != nil {
		panic(err)
	}

	f.Close()
}
