package main

import (
	"github.com/astaxie/beego"
	"syscall"
	"path/filepath"
	"os"
	"os/exec"
	"log"
)

func DiskUsage2(path string) (float64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		beego.Debug("get path error")
		return 0,err
	}
	all := fs.Blocks * uint64(fs.Bsize)
	used := fs.Bfree * uint64(fs.Bsize)
	free := float64(all - used) / float64(all)
	return free, nil
}

func getCurPathDisk3() string {
	file, _ := exec.LookPath(os.Args[0])
	return filepath.VolumeName(file)
}

func main(){
	disk := getCurPathDisk3()
	log.Printf("当前磁盘：%v",disk)
	usage, err := DiskUsage2(disk)
	log.Printf("usage:%v, err:%v",usage,err)
}