package main

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"fmt"
	"os"
	"runtime"
	"time"
)

func convertByteToMB(num uint64) uint64 {
	return  num>>20

}
func t()  {
	v, _ := mem.VirtualMemory()
	// almost every return value is a struct
	total:= convertByteToMB(v.Total)
	free:=convertByteToMB(v.Free)
	used:=convertByteToMB(v.Used)
	fmt.Printf("Total: %v, Free:%v, Used:%v ,UsedPercent:%f\n",total,free,used, v.UsedPercent)
}

func findDeath() {
	stats := new(runtime.MemStats)
	var bigSlice []string
	for {
		bigSlice := append(bigSlice, "abcdefg")
		fmt.Println(len(bigSlice), cap(bigSlice))

		runtime.ReadMemStats(stats)
		fmt.Printf("当前进程内存占用 %v MB", stats.TotalAlloc/(1024*1024))

		time.Sleep(1e2)
	}
}

func temp()  {
	proce,err:=process.NewProcess(int32(os.Getpid()))
	if err!=nil {
		panic(err)
	}

	percent,_:=proce.MemoryPercent()
	fmt.Println("进程内存占用比",percent)

	infos,err:=proce.MemoryInfo()
	if err!=nil {
		panic(err)
	}
	fmt.Printf("内存占用%v (字节)，进程号：%v，父进程为：%v\n",infos.RSS,os.Getpid(),os.Getppid())
	fmt.Scanln()
}
func touchDeath(num int)  {
	for i:=0;i<num ;i++  {
		go func() {
			s := make([]string, 0)
			for {
				s = append(s, "abcdefg")
				time.Sleep(1e3)
			}
		}()
	}
}
func main() {
	currePID:=os.Getpid()
	proce,err:=process.NewProcess(int32(currePID))
	if err!=nil {
		panic(err)
	}
	// 找死
	touchDeath(100)

	// 检测如果内存占用超过1G，关闭此进程
	for {
		infos,err:=proce.MemoryInfo()
		if err!=nil {
			panic(err)
		}
		percent,err:=proce.MemoryPercent()
		if err!=nil {
			panic(err)
		}
		fmt.Printf("进程号：%v,内存占用比%v ，内存占用:%v(M) ,，\n", currePID,percent, infos.RSS>>20)
		if infos.RSS>>20 > 1024 {
			fmt.Println("进程",currePID,"内存占用超过1G，关闭")
			err := proce.Kill()
			fmt.Println(err)
		}
		time.Sleep(1e8)
	}

}