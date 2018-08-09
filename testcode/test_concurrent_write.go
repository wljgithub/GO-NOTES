package main

import (
	// "fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func writeData(lock *sync.RWMutex, f *os.File, str string) {
	lock.Lock()
	defer lock.Unlock()

	f.WriteString(str)
	f.Write([]byte("\n"))
}
func main() {
	DataFile := "./1.txt"
	var lock = &sync.RWMutex{}
	file, err := os.OpenFile(DataFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		go writeData(lock, file, strconv.Itoa(i))

	}
	time.Sleep(1 * time.Second)
	file.Close()
}
