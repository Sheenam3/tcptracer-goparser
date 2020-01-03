package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Log struct {
	log   string
	pid   int64
	time  float64
	probe string
}

type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs uint64

	NumGC        uint32
	NumGoroutine int
}

var logging []Log
var tcplog []Log

const (
	timestamp int = 0
)

func NewMonitor(duration int) {
	var m Monitor
	var rtm runtime.MemStats
	var interval = time.Duration(duration) * time.Second
	for {
		<-time.After(interval)

		// Read full mem stats
		runtime.ReadMemStats(&rtm)

		// Number of goroutines
		m.NumGoroutine = runtime.NumGoroutine()

		// Misc memory stats
		m.Alloc = rtm.Alloc
		m.TotalAlloc = rtm.TotalAlloc
		m.Sys = rtm.Sys
		m.Mallocs = rtm.Mallocs
		m.Frees = rtm.Frees

		// Live objects = Mallocs - Frees
		m.LiveObjects = m.Mallocs - m.Frees

		// GC Stats
		m.PauseTotalNs = rtm.PauseTotalNs
		m.NumGC = rtm.NumGC

		// Just encode to json and print
		b, _ := json.Marshal(m)
		fmt.Println(string(b))
	}
}

func runTCP(tool string) {

	//quit := make(chan bool)

	if tool == "tcptracer" {
		cmd := exec.Command("./tcptracer", "-t")
		cmd.Dir = "/usr/share/bcc/tools"
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()
		probePid := cmd.Process.Pid
		log.Printf("pid: %d", cmd.Process.Pid)
		probeName, err := ps.FindProcess(probePid)
		if err != nil {
			fmt.Println("Error : ", err)
			os.Exit(-1)
		}
		log.Printf("Probe Name: %v", probeName.Executable())
		buf := bufio.NewReader(stdout)
		num := 1
		for {

			line, _, _ := buf.ReadLine()
			parsedLine := strings.Fields(string(line))
			//println("TCP TRACER", parsedLine[0])
			if parsedLine[0] != "Tracing" {
				if parsedLine[0] != "TIME(s)" {
					ppid, err := strconv.ParseInt(parsedLine[2], 10, 64)
					if err != nil {
						println("Tcptracer PID Error")
					}
					timest, err := strconv.ParseFloat(parsedLine[timestamp], 64)
					if err != nil {
						println(" Tcptracer Timestamp Error")
					}

					pn := probeName.Executable()
					n := Log{log: string(line), pid: ppid, time: timest, probe: pn}
					tcplog = append(tcplog, n)

				}
			}

			if num > 10 {
				for i := 0; i < 9; i++ {
					fmt.Printf("Struct %d  includes: %v\n", i, tcplog[i])
					fmt.Printf("Output %d: %v\n PID:%v \t| TimeStamp:%v \t | ProbeName:%v \n", i, tcplog[i].log, tcplog[i].pid, tcplog[i].time, tcplog[i].probe)
				}
				//quit <- true
			}
			num += 1

		}
	}

	if tool == "tcpconnect" {
		cmd := exec.Command("./tcpconnect", "-t")
		cmd.Dir = "/usr/share/bcc/tools"
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()
		probePid := cmd.Process.Pid
		log.Printf("pid: %d", cmd.Process.Pid)
		probeName, err := ps.FindProcess(probePid)
		if err != nil {
			fmt.Println("Error : ", err)
			os.Exit(-1)
		}
		log.Printf("Probe Name: %v", probeName.Executable())

		buf := bufio.NewReader(stdout)
		num := 1

		for {
			line, _, _ := buf.ReadLine()
			parsedLine := strings.Fields(string(line))
			//println(parsedLine[0])
			if parsedLine[0] != "TIME(s)" {
				ppid, err := strconv.ParseInt(parsedLine[1], 10, 64)
				if err != nil {
					println("TCPConnect PID Error")
				}
				timest, err := strconv.ParseFloat(parsedLine[timestamp], 64)
				if err != nil {
					println(" TCPConnect Timestamp Error")
				}

				pn := probeName.Executable()
				n := Log{log: string(line), pid: ppid, time: timest, probe: pn}
				logging = append(logging, n)

			}

			if num > 10 {
				for i := 0; i < 9; i++ {
					fmt.Printf("Struct %d  includes: %v\n", i, logging[i])
					fmt.Printf("Output %d: %v\n PID:%v \t| TimeStamp:%v \t | ProbeName:%v \n", i, logging[i].log, logging[i].pid, logging[i].time, logging[i].probe)
				}
				//quit <- true
			}
			num += 1

		}
	}

}
func main() {

	go runTCP("tcptracer")
	go runTCP("tcpconnect")
	for {
		time.Sleep(10 * time.Second)
		//Displays stats after 10 seconds
		go NewMonitor(10)
	}
	
}
