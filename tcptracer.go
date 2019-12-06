package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	args := "5 /usr/share/bcc/tools/tcpconnect"
	out, _ := exec.Command("timeout", strings.Split(args, " ")...).Output()
	

	
	if len(os.Args) > 1{
		if  os.Args[1] == "-id" {
			parsePIDUsed(string(out))
		} else if os.Args[1] == "-name" {
			parseNameUsed(string(out))
		} else if os.Args[1] == "-port"   {
			parseViaPort(string(out), os.Args[2])
		} else if os.Args[1] == "-pid" {
			parseViaPID(string(out), os.Args[2])
		} else if os.Args[1] == "-h" {
			println("Usage: \n 1. -pid -> ./parse -pid 27515 \n 2. -port port_name -> ./parse -port 443 \n 3. -id (Only display PIds used) -> ./parse -id \n 4. -name (Only display app names) -> ./parse -name")
		}
	} else {
		fmt.Printf("\n%s\n", string(out))
		}
	}



func parseViaPID(out string, p string) {
	outlines := strings.Split(out, "\n")
	//fmt.Println(outlines)
	l := len(outlines)
	fmt.Printf("\n%s\n", outlines[0])
	for _, line := range outlines[1:l-1] {
		parsedLine := strings.Fields(line)
		if parsedLine[0] == p {
			fmt.Printf("\n%s\n", line)
		}
	}
}

func parseViaPort(out string, p string) {
	outlines := strings.Split(out, "\n")
	//fmt.Println(outlines)
	l := len(outlines)
	fmt.Printf("\n%s\n", outlines[0])
	for _, line := range outlines[1:l-1] {
		parsedLine := strings.Fields(line)
		if parsedLine[5] == p {
			fmt.Printf("\n%s\n", line)
		}
	}
}

func parseNameUsed(out string)  {
	outlines := strings.Split(out, "\n")
	//fmt.Println(outlines)
	l := len(outlines)
	title := strings.Fields(outlines[0])
	fmt.Printf("\n%s\n", title[1])
	for _, line := range outlines[1:l-1] {
		parsedLine := strings.Fields(line)
		name := parsedLine[1]
		fmt.Println("\n", name)
	}
}

func parsePIDUsed(out string)  {
	outlines := strings.Split(out, "\n")
	l := len(outlines)
	title := strings.Fields(outlines[0])
	fmt.Printf("PIDs in Use\n%s\n", title[0])
	for _, line := range outlines[1:l-1] {
		parsedLine := strings.Fields(line)
		pid, err := strconv.ParseInt(parsedLine[0],10, 64)
		if err != nil {
			println("Error")
		}
		fmt.Println("\n",pid)
	}
}
