package main

import "fmt"

type Sms struct {
	Opt
}

func (s *Sms) genRandomOPT(len int) string {
	ramdomOPT := "1234"
	fmt.Printf("SMS: generating ramdom opt %s\n", ramdomOPT)
	return ramdomOPT
}

func (s *Sms) saveOPTCache(opt string) {
	fmt.Printf("SMS: saving opt %s to cache\n", opt)
}

func (s *Sms) getMessage(opt string) string {
	return "SMS OPT for login is " + opt
}

func (s *Sms) sendNotification(msg string) error {
	fmt.Printf("SMS: sending notification %s\n", msg)
	return nil
}
