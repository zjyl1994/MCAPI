package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var minecraftProc *exec.Cmd
var stdin io.WriteCloser
var stdout io.ReadCloser
var stderr io.ReadCloser
var serverRunning bool

func startMinecraft() {
	if(serverRunning){
		return
	}
	log.Println("Server Starting!")
	minecraftProc = exec.Command("java", fmt.Sprintf("-Xmx%dM", cfg.MaxMemory), "-jar", cfg.JarPath)
	var err error
	stdin, err = minecraftProc.StdinPipe()
	checkError(err)
	stdout, err = minecraftProc.StdoutPipe()
	checkError(err)
	stderr, err = minecraftProc.StderrPipe()
	checkError(err)
	minecraftProc.Dir = cfg.WorkDir // 设置工作目录
	err = minecraftProc.Start()     // 启动程序
	checkError(err)
	// 管道对接
	go io.Copy(stdin, os.Stdin)
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)
	serverRunning = true
}

func stopMinecraft() {
	if(!serverRunning){
		return
	}
	execMinecraft(cfg.StopCommand)
	minecraftProc.Wait()
	log.Println("Server STOPed!")
	serverRunning = false
}

func restartMinecraft() {
	stopMinecraft()
	startMinecraft()
}

func execMinecraft(command string) string {
	if !strings.HasSuffix(command, "\n") {
		command += "\n"
	}
	lenCommand, err := stdin.Write([]byte(command))
	checkError(err)
	if lenCommand != len(command) {
		log.Fatalf("Error: write command not correct", err)
	}
	return "ok"
}
