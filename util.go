package main

import (
	"log"
	"time"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func GetCPUPercent()(float64,error){
	percent,err := cpu.Percent((500 * time.Millisecond), false)
	return percent[0],err
}

func GetMEMPercent()(float64,error){
	v, err := mem.VirtualMemory()
    return v.UsedPercent,err
}

type systemStatus struct{
	CPU int
	MEM int
	Running bool
}

func GetSystemStatus()systemStatus{
	c, err := GetCPUPercent()
	checkError(err)
	v, err := GetMEMPercent()
	checkError(err)
	return systemStatus{CPU:int(c),MEM:int(v),Running:serverRunning}
}

func AuthRequest(r *http.Request)bool{
	return (r.PostFormValue("password") == cfg.SecertKey)
}