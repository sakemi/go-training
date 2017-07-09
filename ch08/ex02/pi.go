package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cmd struct {
	pat    *regexp.Regexp
	handle func(*Ctrl, *Info, ...string)
}
type Info struct {
	addr string
	dir  string
}
type Reply struct {
	code int
	mes  string
}
type Ctrl struct {
	net.Conn
}

var (
	newUser          = Reply{220, "Service ready for new user."}
	password         = Reply{331, "User name okay, need password."}
	login            = Reply{230, "User logged in, proceed."}
	notImpl          = Reply{502, "Command not implemented."}
	illegalArg       = Reply{501, "Syntax error in parameters or arguments."}
	okay             = Reply{200, "Command okay."}
	notFound         = Reply{550, "Requested action not taken."}
	complete         = Reply{250, "Requested file action okay, completed."}
	quit             = Reply{221, "Service closing control connection."}
	openConn         = Reply{150, "File status okay; about to open data connection."}
	failToConn       = Reply{425, "Can't open data connection."}
	failToOpenFile   = Reply{450, "Requested file action not taken."}
	transferAbort    = Reply{426, "Connection closed; transfer aborted."}
	closeConn        = Reply{226, "Closing data connection."}
	transferComplete = Reply{250, "Requested file action okay, completed."}
)

var commands = map[string]Cmd{
	"USER": Cmd{regexp.MustCompile(`(?i)user (.*)$`), usr},
	"PASS": Cmd{regexp.MustCompile(`(?i)pass (.*)$`), pas},
	"CWD":  Cmd{regexp.MustCompile(`(?i)cwd (.*)$`), cwd},
	"QUIT": Cmd{regexp.MustCompile(`(?i)quit$`), qui},
	"PORT": Cmd{regexp.MustCompile(`(?i)port (.*)$`), por},
	"TYPE": Cmd{regexp.MustCompile(`(?i)type (.*)$`), typ},
	"STRU": Cmd{regexp.MustCompile(`(?i)stru (.*)$`), str},
	"RETR": Cmd{regexp.MustCompile(`(?i)retr (.*)$`), ret},
	"STOR": Cmd{regexp.MustCompile(`(?i)stor (.*)$`), sto},
	"NOOP": Cmd{regexp.MustCompile(`(?i)noop$`), noo},
}

func (c *Ctrl) reply(r Reply) {
	io.WriteString(c, fmt.Sprintf("%d %s\r\n", r.code, r.mes))
}

func interpret(c *Ctrl, info *Info, command string) bool {
	for _, com := range commands {
		if com.pat.MatchString(command) {
			log.Println(command)
			if strings.Contains(command, " ") {
				args := com.pat.FindStringSubmatch(command)[1]
				com.handle(c, info, strings.Split(args, " ")...)
				return true
			}
			com.handle(c, info)
			return true
		}
	}
	log.Printf("Unknown command: %s", command)
	c.reply(notImpl)
	return false
}

func usr(ctrl *Ctrl, _ *Info, arg ...string) {
	ctrl.reply(password)
}

func pas(ctrl *Ctrl, _ *Info, arg ...string) {
	ctrl.reply(login)
}

func cwd(ctrl *Ctrl, info *Info, arg ...string) {
	d := "." + arg[0]

	i, err := os.Stat(d)
	if err != nil || !i.IsDir() {
		ctrl.reply(notFound)
		return
	}
	info.dir = d
	ctrl.reply(complete)
}

func qui(ctrl *Ctrl, _ *Info, arg ...string) {
	ctrl.reply(quit)
}

func por(ctrl *Ctrl, info *Info, arg ...string) {
	val := strings.Split(arg[0], ",")
	if len(val) != 6 {
		ctrl.reply(illegalArg)
		return
	}
	p1, err := strconv.Atoi(val[4])
	if err != nil {
		log.Println(err)
		ctrl.reply(illegalArg)
		return
	}
	p2, err := strconv.Atoi(val[5])
	if err != nil {
		log.Println(err)
		ctrl.reply(illegalArg)
		return
	}

	ip := strings.Join(val[0:4], ".")
	port := strconv.Itoa(p1*256 + p2)
	info.addr = ip + ":" + port
	ctrl.reply(okay)
}

func typ(ctrl *Ctrl, _ *Info, arg ...string) {
	if arg[0] == "A" || arg[0] == "I" {
		ctrl.reply(okay)
		return
	}
	ctrl.reply(notImpl)
	return
}

func str(ctrl *Ctrl, _ *Info, arg ...string) {
	ctrl.reply(okay)
}

func ret(ctrl *Ctrl, info *Info, arg ...string) {
	path := info.dir + arg[0]
	i, err := os.Stat(path)
	if err != nil || i.IsDir() {
		log.Println(err)
		ctrl.reply(notFound)
		return
	}

	ctrl.reply(openConn)
	retrData(ctrl, info.addr, path)
}

func sto(ctrl *Ctrl, info *Info, arg ...string) {
	path := info.dir + arg[0]
	ctrl.reply(openConn)
	storData(ctrl, info.addr, path)
}

func noo(ctrl *Ctrl, _ *Info, arg ...string) {
	ctrl.reply(okay)
}
