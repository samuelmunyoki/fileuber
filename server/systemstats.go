package server

import (
	"context"
	"time"
	"math"

	"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)
type Download struct {
    File  string `json:"file"`
    Speed float64 `json:"speed"`
}

var downloads []Download

func AddDownload(file string, speed float64) {
    for i := range downloads {
        if downloads[i].File == file {
            // Remove the existing entry if the file matches
            downloads = append(downloads[:i], downloads[i+1:]...)
            break
        }
    }
    // Add a new Download instance to the downloads slice
    newDownload := Download{File: file, Speed: speed}
    downloads = append(downloads, newDownload)
}





type System struct {
	ctx context.Context
}
type CPU struct {
	Usage int `json:"usage"`
	Freq float64 `json:"freq"`
	Model string `json:"model"`
	Cores int32 `json:"cores"`
}


func NewSystem() *System {
	return &System{}
}


func (s *System) getCPU () *CPU{
	pCpu, err := cpu.Percent(time.Second, false)
	if err != nil{
		return &CPU{Usage: 0}
	}
	cpu, err := cpu.Info()
	if err != nil{
		
		return &CPU{}
	}
	cpuI :=cpu[0]
	return &CPU{
		Freq: cpuI.Mhz,
		Model: cpuI.ModelName,
		Cores: cpuI.Cores,
		Usage: int(math.Round(pCpu[0])),}}

func (s *System) SystemData(ctx context.Context){
	s.ctx = ctx

	go func() {
		for { 
			runtime.EventsEmit(s.ctx, "cpudata", s.getCPU())
			downloads=nil
			time.Sleep(5*time.Second)
			
		}
	}()
	go func() {
		for { 
			runtime.EventsEmit(s.ctx, "filedownloads", downloads)
			
			time.Sleep(100*time.Millisecond)
			
			
		}
	}()
}

