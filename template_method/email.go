package main

import "fmt"

type Email struct {
	Opt
}

func (s *Email) genRandomOPT(len int) string {
	ramdomOPT := "1234"
	fmt.Printf("Email: generating ramdom opt %s\n", ramdomOPT)
	return ramdomOPT
}

func (s *Email) saveOPTCache(opt string) {
	fmt.Printf("Email: saving opt %s to cache\n", opt)
}

func (s *Email) getMessage(opt string) string {
	return "Email OPT for login is " + opt
}

func (s *Email) sendNotification(msg string) error {
	fmt.Printf("Email: sending notification %s\n", msg)
	return nil
}
