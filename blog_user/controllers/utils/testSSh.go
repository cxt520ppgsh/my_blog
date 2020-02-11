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

func TestBuild() {
	cli := New(SSH_IP, SSH_USERNAME, SSH_PASSWORD, SSH_PORT)
	var outbf, errbf bytes.Buffer
	go printTerminalLast3(&outbf)
	cli.RunTerminal("cd "+SSH_PROJECT_PATH+";pwd"+";./allmake.sh -p ac8257_demo_1g_32 -c 2G", &outbf, &errbf)
}

func TestCd() {
	cli := New(SSH_IP, SSH_USERNAME, SSH_PASSWORD, SSH_PORT)
	output, err := cli.Run("cd " + SSH_PROJECT_PATH + ";pwd")
	fmt.Printf("%v\n%v", output, err)
}

func printTerminalLast3(outb *bytes.Buffer) {
	for {
		time.Sleep(3000 * time.Millisecond)
		arr := strings.Split(outb.String(), "\n")
		if len(arr) > 3 {
			line := arr[len(arr)-3]
			fmt.Println(line)
			SendSocketMessage("buildLog" + line)
		}
	}
}
