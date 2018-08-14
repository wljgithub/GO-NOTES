package main

import (
	"log"
	"path"
	"strings"
	"time"
)

func name() {

}
func nextDayFile(fileName string) (string, bool) {
	baseName := path.Base(fileName)
	date := strings.TrimRight(baseName, ".txt")
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", false
	}
	return t.AddDate(0, 0, 1).Format("2006-01-02") + ".txt", true

}
func main() {
	fileName := "2018-01-02.txt"
	log.Println(nextDayFile(fileName))
}
