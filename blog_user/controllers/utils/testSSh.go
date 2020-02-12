package utils

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const SSH_USERNAME = "chenxutang"
const SSH_PASSWORD = "fly098321"
const SSH_PROJECT_PATH = "project/010"
const SSH_IP = "172.168.1.187"
const SSH_PORT = 22

//const SSH_USERNAME = "chenxutang"
//const SSH_PASSWORD = "T373683458"
//const SSH_PROJECT_PATH = "project"
//const SSH_IP = "39.100.119.126"
//const SSH_PORT = 22

var cli = New(SSH_IP, SSH_USERNAME, SSH_PASSWORD, SSH_PORT)
var building = false

func TestBuild() {
	var outbf, errbf bytes.Buffer
	go printTerminalLast3(&outbf)
	go cli.RunTerminal("cd "+SSH_PROJECT_PATH+";pwd"+";./allmake.sh -p ac8257_demo_1g_32 -c 2G", &outbf, &errbf)
	building = true
}

func printTerminalLast3(outb *bytes.Buffer) {
	for {
		time.Sleep(3000 * time.Millisecond)
		if building {
			arr := strings.Split(outb.String(), "\n")
			if len(arr) > 3 {
				line := arr[len(arr)-3]
				fmt.Println(line)
				SendSocketMessage(line + "buildLog")
			}
		}
	}
}

func StopBuild() {
	cli.StopSession()
	building = false
}

func TestCd() {
	cli := New(SSH_IP, SSH_USERNAME, SSH_PASSWORD, SSH_PORT)
	output, err := cli.Run("cd " + SSH_PROJECT_PATH + ";pwd")
	fmt.Printf("%v\n%v", output, err)
}
