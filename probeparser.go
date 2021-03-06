package main

import (
"bufio"
//"encoding/json"
//"fmt"
"log"
//"os"
"os/exec"
//"runtime"
"strconv"
"strings"
"time"
)

type Log struct {
	Fulllog   string
	Pid   int64
	Time  float64
	Probe string
}

/*type Monitor struct {
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
*/

//var IsTracerDoneSig = make(chan bool, 1)
//var IsTCPTracerDoneSig = make(chan bool, 1)
const (
	timestamp int = 0
)

/*func NewMonitor(duration int) {
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
}*/

func RunTcptracer(tool string, logtcptracer chan Log, pid [][]string) {

		var pd string
		for j := range pid {

			for k := range pid[j] {
			        pd = pid[j][k]
			println("pid in array:",pd)

			}
		}
		cmd := exec.Command("./tcptracer", "-t", "-p" + pd)
		cmd.Dir = "/usr/share/bcc/tools"
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()
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
					n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
					logtcptracer <- n
					if num > 1000 {
                                                close(logtcptracer)
                                                log.Println("Tracer has been Stopped")
//                                                IsTCPTracerDoneSig <- true

                        	        }
                            	num++
				//Tcplog = append(Tcplog, n)
				}
			}
		}
}




func RunTcpconnect(tool string, logtcpconnect chan Log) {


	
/*			if num > 10 {
				for i := 0; i < 9; i++ {
					return Tcplog
					/*fmt.Printf("Struct %d  includes: %v\n", i, tcplog[i])
					fmt.Printf("Output %d: %v\n PID:%v \t| TimeStamp:%v \t | ProbeName:%v \n", i, tcplog[i].log, tcplog[i].pid, tcplog[i].time, tcplog[i].probe)
				}
				//quit <- true
			}
			num += 1

		}
	}*/

	//if tool == "tcpconnect" {
		cmd := exec.Command("./tcpconnect", "-t")
		cmd.Dir = "/usr/share/bcc/tools"
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()
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


				n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
				logtcpconnect <- n
				if num > 300 {
						close(logtcpconnect)
						log.Println("Tracerconnect has been Stopped")					
//						IsTracerDoneSig <- true
						
				}
				num++
				//Logging = append(Logging, n)

			}
/*
			if num > 10 {
				for i := 0; i < 9; i++ {
					return Logging
					/*fmt.Printf("Struct %d  includes: %v\n", i, Logging[i])
					fmt.Printf("Output %d: %v\n PID:%v \t| TimeStamp:%v \t | ProbeName:%v \n", i, Logging[i].Log, Logging[i].Pid, Logging[i].Time, Logging[i].Probe)
				}

			}
			num += 1*/

		}
	//}
	
}



func getPID() [][]string{


		var matrix [][]string
		var row []string

		
		cmd := exec.Command("./tcptracer", "-t")
		cmd.Dir = "/usr/share/bcc/tools"
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()
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
					s := strconv.FormatInt(ppid, 10)
					row = append(row,s)
					matrix = append(matrix, row) 
					//timest, err := strconv.ParseFloat(parsedLine[timestamp], 64)
					//if err != nil {
					//	println(" Tcptracer Timestamp Error")
					//}
				//	n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
				println(matrix)
				}
			}
				for  num > 5 {

					return matrix					
				}

		}

}


func main() {
	//go RunTCP("tcptracer")
	logtcptracer :=  make(chan Log, 1)
//	logtcpconnect := make(chan Log, 1)

	var pid [][]string
	pid = getPID()	
	go RunTcptracer("tcptracer", logtcptracer,pid)
	for val := range logtcptracer {
	log.Printf("%v Probe: %s, Pid: %d", val.Fulllog, val.Probe, val.Pid)

	}
	/*
	go RunProbe("tcpconnect", logtcpconnect)
	for val := range logtcpconnect {
	log.Printf("%v Probe: %s, Pid: %d", val.Fulllog, val.Probe, val.Pid)

	}*/

	for 
	{

		time.Sleep(10 * time.Second)
	}
}
