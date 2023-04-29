package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "exitsim"
	app.Commands = []*cli.Command{
		&cli.Command{
			Name:        "exit",
			Description: "exit the program with a simple status code",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:     "code",
					Usage:    "0-255",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				code := c.Int("code")
				if code >= 0 && code <= 255 {
					return cli.Exit(fmt.Sprintf("exiting with code: %d", code), code)
				} else {
					return cli.Exit(fmt.Sprintf("code out of range (%d), exiting with code 1", code), 1)
				}
			},
		},
		&cli.Command{
			Name:        "signal",
			Description: "exit the program with a signal",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Usage:    "sigterm",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				signame := c.String("name")
				var killWith syscall.Signal
				switch signame {
				case "sigabrt":
					killWith = syscall.SIGABRT
				case "sigalrm":
					killWith = syscall.SIGALRM
				case "sigbus":
					killWith = syscall.SIGBUS
				case "sigchld":
					killWith = syscall.SIGCHLD
				case "sigcld":
					killWith = syscall.SIGCLD
				case "sigcont":
					killWith = syscall.SIGCONT
				case "sigfpe":
					killWith = syscall.SIGFPE
				case "sighup":
					killWith = syscall.SIGHUP
				case "sigill":
					killWith = syscall.SIGILL
				case "sigint":
					killWith = syscall.SIGINT
				case "sigio":
					killWith = syscall.SIGIO
				case "sigiot":
					killWith = syscall.SIGIOT
				case "sigkill":
					killWith = syscall.SIGKILL
				case "sigpipe":
					killWith = syscall.SIGPIPE
				case "sigpoll":
					killWith = syscall.SIGPOLL
				case "sigprof":
					killWith = syscall.SIGPROF
				case "sigpwr":
					killWith = syscall.SIGPWR
				case "sigquit":
					killWith = syscall.SIGQUIT
				case "sigsegv":
					killWith = syscall.SIGSEGV
				case "sigstkflt":
					killWith = syscall.SIGSTKFLT
				case "sigstop":
					killWith = syscall.SIGSTOP
				case "sigsys":
					killWith = syscall.SIGSYS
				case "sigterm":
					killWith = syscall.SIGTERM
				case "sigtrap":
					killWith = syscall.SIGTRAP
				case "sigtstp":
					killWith = syscall.SIGTSTP
				case "sigttin":
					killWith = syscall.SIGTTIN
				case "sigttou":
					killWith = syscall.SIGTTOU
				case "sigunused":
					killWith = syscall.SIGUNUSED
				case "sigurg":
					killWith = syscall.SIGURG
				case "sigusr1":
					killWith = syscall.SIGUSR1
				case "sigusr2":
					killWith = syscall.SIGUSR2
				case "sigvtalrm":
					killWith = syscall.SIGVTALRM
				case "sigwinch":
					killWith = syscall.SIGWINCH
				case "sigxcpu":
					killWith = syscall.SIGXCPU
				case "sigxfsz":
					killWith = syscall.SIGXFSZ
				default:
					return cli.Exit(fmt.Sprintf("unrecognized signal name (%s), exiting with code 1", signame), 1)
				}
				syscall.Kill(syscall.Getpid(), killWith)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
