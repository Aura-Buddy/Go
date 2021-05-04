package main

import(
	"fmt"
	"os"
	"os/exec"
)

func main(){
	for i:=0; i < 10; i++{
	goExecPath, err := exec.LookPath("go")
	if err != nil{
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Go executable: ", goExecPath)
	}
	cmdGoVer := &exec.Cmd{
		Path: goExecPath,
		Args: []string{goExecPath, "run", "automatedScript.go"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	fmt.Println(cmdGoVer.String())

	if err := cmdGoVer.Run(); err!= nil{
		fmt.Println("Error",err);
	}
}
}
