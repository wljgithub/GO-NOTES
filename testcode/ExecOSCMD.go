// exec.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func m() {
	// 1) os.StartProcess //
	/*********************/
	/* Linux: */
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
		fmt.Printf("The process id is %v", pid)
	}
}

func ExecuteCMD() {
	out, err := exec.Command("ps", `aux`).Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The date is %s\n", out)
}
func ExecuteCMDWithPipe() {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

func PS() {
	c1 := exec.Command("ls")
	c2 := exec.Command("wc", "-l")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	c2.Stderr = os.Stderr
	c2.Start()
	c1.Run()
	c2.Wait()

}
func killProcesByPID() {
	proc, err := os.FindProcess(3799)
	if err != nil {
		panic(err)
	}

	err = proc.Kill()
	if err != nil {
		panic(err)
	}

}

func getCurrentPid() {
	fmt.Println("主进程的PID：", os.Getpid())
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("当前线程程的PID：", os.Getpid())
		}()
		time.Sleep(1e9)
	}
	fmt.Scanln()
}
func main() {
	PS()
}
