package main

import (
	"bufio"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"log"
	"os"
	"os/exec"
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

var logging []Log
var tcplog []Log

const (
	timestamp int = 0
)

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
	}
	
}
