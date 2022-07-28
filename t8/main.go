package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		command := scanner.Text()
		if strings.Contains(command, "|") {
			for _, v := range strings.Split(command, "|") {
				processCommands(v)
			}
		}
		processCommands(command)
	}
}

func processCommands(command string) {
	splited := strings.Split(command, " ")
	switch splited[0] {
	case "cd":
		changeDirectory(splited[1])
	case "pwd":
		pwd()
	case "echo":
		echo(splited[1])
	case "kill":
		killProcess(splited[1])
	case "ps":
		getProcesses()
	case "quit":
		fmt.Println("exit")
		os.Exit(0)
	case "fp":
		findProcess(splited[1])
	default:
		fmt.Println("invalid command")
	}
}

func changeDirectory(directory string) {
	err := os.Chdir(directory)
	if err != nil {
		fmt.Println(err)
	}
}

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)

}

func echo(data string) {
	fmt.Println(data)
}

func killProcess(process string) {
	pid, err := strconv.Atoi(process)
	if err != nil {
		fmt.Println(err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
	}
	if err := proc.Kill(); err != nil {
		fmt.Println(err)
	}
}

func getProcesses() {
	processes, _ := ps.Processes()
	for _, process := range processes {
		fmt.Println("process: ", process.Executable(), " pid: ", process.Pid())
	}
}

func findProcess(name string) {
	processes, _ := ps.Processes()
	for _, process := range processes {
		if process.Executable() == name {
			fmt.Println("process: ", process.Executable(), " pid: ", process.Pid())
		}
	}
}
