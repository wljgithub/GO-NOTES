// exec.go
package main

import (
	"fmt"
	"os"
	"os/exec"
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

func ExecuteDiretely() {
	out, err := exec.Command("ps", `aux`).Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The date is %s\n", out)
}
func killProcesByPID(){
		proc, err := os.FindProcess(3799)
	if err != nil {
		panic(err)
	}

	err = proc.Kill()
	if err != nil {
		panic(err)
	}
}
func main() {


}
