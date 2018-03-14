package main

import (
	"os/exec"
	"log"
	"io/ioutil"
	"os"
)

func cmd_start(){
	cmd := exec.Command("ls", "-a", "-l")
	stdout,err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()

	if err := cmd.Start();err !=nil{
		log.Fatal(err)
	}

	rst,err := ioutil.ReadAll(stdout)
	if err != nil{
		log.Fatal(err)
	}

	log.Println(string(rst))
}

func cmd_run(){
	cmd := exec.Command("df","-h1")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil{
		log.Fatal(err)
	}
}

func main(){
	cmd_run()
}
