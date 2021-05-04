package main

import(
	"fmt"
	"os"
	"os/exec"
)

func main(){
	goExecPath, err := exec.LookPath("docker")
	if err != nil{
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Go executable: ", goExecPath)
	}
	cmdGoVer := &exec.Cmd{
		Path: goExecPath,
		Args: []string{goExecPath, "exec", "mqtt", "mosquitto_pub", "-u", "fen536", "-P", "Buddy", "-t", "Buddy", "-m", "check"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	fmt.Println(cmdGoVer.String())

	if err := cmdGoVer.Run(); err!= nil{
		fmt.Println("Error",err);
	}
}
