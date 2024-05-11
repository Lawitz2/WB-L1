package main

import "fmt"

// old

type OldLogger struct{}

func (ol *OldLogger) oldLoggerLog(msg string) {
	fmt.Println("old log: " + msg)
}

// new

type logging interface {
	log(message string)
}

type NewLogger struct{}

func (nl *NewLogger) log(message string) {
	fmt.Println("new logger log: " + message)
}

type client21 struct{}

func (c *client21) savelog(message string, log logging) {
	log.log(message)
}

// adapter to use old thing through new methods

type OldLogAdapter struct {
	OldLog *OldLogger
}

func (oldlogAd *OldLogAdapter) log(message string) {
	oldlogAd.OldLog.oldLoggerLog(message)
}

func Ex21_2() {
	cl := client21{}
	nl := &NewLogger{}
	cl.savelog("msg", nl)

	ol := &OldLogger{}
	adap := &OldLogAdapter{ol}
	cl.savelog("msg2", adap)
}
