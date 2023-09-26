package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	time := GetTime()

	fmt.Println(time.Format("2006-01-02 15:04:05"))
}

//GetTime возвращает текущее время, полученное с сервера 0.beevik-ntp.pool.ntp.org.

func GetTime() time.Time {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
		os.Exit(1) // При ошибке выводит текст ошибки в stderr и завершает работу программы с кодом 1.

	}

	return time
}
