// MCAPI project main.go
package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
)

func startHandler(w http.ResponseWriter, r *http.Request) {
	if AuthRequest(r){
		go startMinecraft()
		w.Write([]byte("启动指令已送达!\n"))
	}else{
		w.Write([]byte("授权码错误!\n"))
	}
}

func restartHandler(w http.ResponseWriter, r *http.Request) {
	if AuthRequest(r){
		go restartMinecraft()
		w.Write([]byte("重启指令已送达!\n"))
	}else{
		w.Write([]byte("授权码错误!\n"))
	}
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	if AuthRequest(r){
		go stopMinecraft()
		w.Write([]byte("停机指令已送达!\n"))
	}else{
		w.Write([]byte("授权码错误!\n"))
	}
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.AllowRemoteExec{
		if AuthRequest(r){
			go execMinecraft(r.PostFormValue("command"))
			w.Write([]byte("命令已送达!\n"))
		}else{
			w.Write([]byte("授权码错误!\n"))
		}
	}else{
		w.Write([]byte("远程命令执行未启用!\n"))
	}
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	b , err := json.Marshal(GetSystemStatus())
	checkError(err)
	w.Write(b)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	b,err := ioutil.ReadFile("panel.html")  
    checkError(err)
	w.Header().Set("content-type", "text/html")
	w.Write(b)
}

func main() {
	log.Println("Zjyl1994's Minecraft Remote Control API!")
	log.Println("Build 20160627")
	loadConfigure("./config.json")
	go startMinecraft()
	//HTTP处理
	r := mux.NewRouter()
	r.HandleFunc("/start", startHandler)
	r.HandleFunc("/restart", restartHandler)
	r.HandleFunc("/stop", stopHandler)
	r.HandleFunc("/exec", execHandler)
	r.HandleFunc("/info", infoHandler)
	r.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(cfg.ListenAddr, r))
}
