package main

import (
	"bufio"
	//"encoding/json"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"log"
	"os"
	"os/exec"
	//"runtime"
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

//var logging []Log
var tcplog []Log

const (
	timestamp int = 0
)



func runTCP(tool string, logchannel chan Log) {

	//quit := make(chan bool)
	/*fmt.Println("checkpoint 1")
	if tool == "tcptracer" {
		cmd := exec.Command("./tcptracer", "-t")
		cmd.Dir = "/usr/share/bcc/tools"
		stdout, err := cmd.StdoutPipe()



		//fmt.Println("checkpoint 3")


		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
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
			// fmt.Println("checkpoint 6")
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

					logchannel <- n
					if num > 40 {
						close(logchannel)
						os.Exit(-1)
					}
					num ++
					//quit <- true
					//tcplog = append(tcplog, n)

				}
			}

			/*if num > 10 {
			          for i := 0; i < 9; i++ {
			                  fmt.Printf("Struct %d  includes: %v\n", i, tcplog[i])
			                  fmt.Printf("Output %d: %v\n PID:%v \t| TimeStamp:%v \t | ProbeName:%v \n", i, tcplog[i].log, tcplog[i].pid, tcplog[i].time, tcplog[i].probe)
			          }
			          //quit <- true
			  }
			  num += 1

		}*/
	//}

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
								logchannel <- n
								if num > 40 {
								close(logchannel)
								os.Exit(-1)
								}
								num ++
	                           //logging = append(logging, n)
	                   }
	                   /*if num > 10 {
	                           for i := 0; i < 9; i++ {
	                                   fmt.Printf("Struct %d  includes: %v\n", i, logging[i])
	                                   fmt.Printf("Output %d: %v\n PID:%v \t| TimeStamp:%v \t | ProbeName:%v \n", i, logging[i].log, logging[i].pid, logging[i].time, logging[i].probe)
	                           }
	                           //quit <- true
	                   }
	                   num += 1*/
	           }
	   }

}
/*
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
func printStats(mem runtime.MemStats) {
        runtime.ReadMemStats(&mem)
        fmt.Println("mem.Alloc:", mem.Alloc)
        fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
        fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
        fmt.Println("mem.NumGC:", mem.NumGC)
        fmt.Println("-----")
}*/

func main() {

	logchannel := make(chan Log,1)

	go runTCP("tcpconnect", logchannel )

	for val := range logchannel {
		fmt.Println(val)
	}
	/*
	data := <-logchannel
	println(data.log, data.pid)*/
	/* fmt.Sprintf("%-8s %-6s %-8s %-10s %-5s %-15s %-15s %-5s %-5s\n", "TIME(s)", "T", "PID", "COMM", "IP", "SADDR", "DADDR", "SPORT","DPORT")

	   var logstring string

	   logstring = fmt.Sprintf("%-8s %-6s %-8d %-5s IPv%-2d %-15s %-15s %-5d %-5d \n ", "PID: %d \n", "ProbeName: %s \n", "TimeStamp: %v", data.log, data.pid, data.probe, data.time)*/

	//fmt.Printf(logstring)

	//      go runTCP("tcpconnect")
	/*
	   //var mem runtime.MemStats
	   /fmt.Println("HAHAHAHAHAHAHHAAHHA")
	   time.Sleep(10 * time.Second)*/
	for {
		time.Sleep(10 * time.Second)
		fmt.Println("INSIDE HAHAHAHAHAHAHHAAHHA")
		//printStats(mem)
		//go NewMonitor(10)
	}


	/*for i := 0; i < 19; i++ {
	        fmt.Printf("Struct %d  includes: %v\n", i, logging[i])
	        fmt.Printf("Output %d: %v\n PID:%v \t| TimeStamp:%v \t | ProbeName:%v \n", i, logging[i].log, logging[i].pid, logging[i].time, logging[i].probe)
	}*/

}
