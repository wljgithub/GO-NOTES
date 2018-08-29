package main

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"fmt"
	"os"
)

func convertByteToMB(num uint64) uint64 {
	return  num/(1024*1024)

}
func t()  {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	total:= convertByteToMB(v.Total)
	free:=convertByteToMB(v.Free)
	used:=convertByteToMB(v.Used)
	fmt.Printf("Total: %v, Free:%v, Used:%v ,UsedPercent:%f\n",total,free,used, v.UsedPercent)
}
func main() {

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
	fmt.Println("内存占用信息",infos.String())
	fmt.Println(proce.CreateTime())

	fmt.Println()
}