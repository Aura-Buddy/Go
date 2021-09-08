package main

import(
        "fmt"
        "os"
        "os/exec"
)

func main(){
        MQTTExecPath, err := exec.LookPath("mosquitto_sub")
        if err != nil{
                fmt.Println("Error: ", err)
        } else {
                fmt.Println("Mosquitto executable: ", MQTTExecPath)
        }
        cmdMQTTVer := &exec.Cmd{
                Path: MQTTExecPath,
                Args: []string{MQTTExecPath, "-C", "1", "-t", "test"},
                Stdout: os.Stdout,
                Stderr: os.Stdout,
        }

        fmt.Println(cmdMQTTVer.String())

        if err := cmdMQTTVer.Run(); err!= nil{
                fmt.Println("Error",err);
        }
}
