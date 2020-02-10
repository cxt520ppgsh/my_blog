package controllers

import (
	"bytes"
	_ "bytes"
	"fmt"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/ssh"
	_ "log"
	_ "os/exec"
	_ "strconv"
	_ "strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func runCmd(session *ssh.Session, cmdStr string) {
	var stdOut, stdErr bytes.Buffer
	session.Stdout = &stdOut
	session.Stderr = &stdErr
	session.Run(cmdStr)
	fmt.Printf("cmd ok " + stdOut.String())
	fmt.Printf("cmd error " + stdErr.String())
}

