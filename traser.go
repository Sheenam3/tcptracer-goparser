package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"os/exec"
)


type Log struct {
                      log string
                      pid [10]int64
                      name [10]string
                      sport [10]int64
                      dport [10]int64
}

const (
                      pid int = 1
                      name int = 2
                      sport int = 6
                      dport int = 7
)

func main() {

	var parse Log
	args := "10 /usr/share/bcc/tools/tcptracer"
	cmd, _ := exec.Command("timeout", strings.Split(args, " ")...).Output()

	storeLog(string(cmd), &parse)

	
	if len(os.Args) > 1 {
		if os.Args[1] == "-pid" {
			//logViaPID(&parse, os.Args[2])
			parseLines(&parse, os.Args[2], pid)
		} else if os.Args[1] == "-name" {
			parseLines(&parse, os.Args[2], name)
		} else if os.Args[1] == "-sport" {
			parseLines(&parse, os.Args[2], sport)
		} else if os.Args[1] == "-dport" {
			parseLines(&parse, os.Args[2], dport)
		} else if os.Args[1] == "-h" {
			log.Print("Usage: \n 1. -pid -> ./traser -pid 27515 \n 2. -name app_name -> ./traser -name kube-apiserver \n 3. -sport port_name -> ./traser -sport 23453 \n 4. dport port_name -> ./traser -dport 8080 \n 5. -id (Only display PIds used) -> ./traser -id \n 6. -app (Only display app names) -> ./traser -app \n 7. -sp (Source Port in use) -> ./traser -sp \n 8. -dp (Destination Port in use) -> ./traser -dp ")
		} else if os.Args[1] == "-id" {
			storeValues(&parse, pid)
			log.Print("\n", parse.pid)
		} else if os.Args[1] == "-app" {
			storeValues(&parse, name)
			log.Print("\n", parse.name)
		} else if os.Args[1] == "-sp" {
			storeValues(&parse, sport)
			log.Print("\n", parse.sport)
		} else if os.Args[1] == "-dp" {
			storeValues(&parse, dport)
			log.Print("\n", parse.dport)
		}
	}else {
		fmt.Printf("\n%s\n", parse.log)
	}
}

func storeLog(cmd string, parse *Log) {
	parse.log = cmd
}


func parseLines(parse *Log, field string, display int) {

	outlines := strings.Split(parse.log, "\n")
	l := len(outlines)
	fmt.Printf("\n%s\n", outlines[1])
	for _, line := range outlines[2 : l-1] {
		parsedLine := strings.Fields(line)
		if parsedLine[display] == field {
			fmt.Printf("\n%s\n", line)
		}

	}

}

func storeValues(parse *Log,display int) {
	outlines := strings.Split(parse.log, "\n")
	l := len(outlines)
	i := 0
	title := strings.Fields(outlines[1])
	fmt.Printf("%s in Use\n", title[display])
	for _, line := range outlines[2 : l-1] {
		parsedLine := strings.Fields(line)

		if i < 10 {
			if  display == 1 {
				pid, err := strconv.ParseInt(parsedLine[display], 10, 64)
				if err != nil {
					println("Error")
				}
				parse.pid[i] = pid
			}else if display == 2{
				name := parsedLine[display]
				parse.name[i] = name
			}else if display ==  6{
				sport, err := strconv.ParseInt(parsedLine[display], 10, 64)
				if err != nil {
					println("Error")
				}
				parse.sport[i] = sport

			} else {
				dport, err := strconv.ParseInt(parsedLine[display], 10, 64)
				if err != nil {
					println("Error")
				}
				parse.dport[i] = dport
			}
			i++

		}
		
	}
}
	
