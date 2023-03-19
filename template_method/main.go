package main

import "fmt"

func main() {
	smsOPT := Sms{}
	o := Opt{iOpt: &smsOPT}
	o.genAndSendOPT(4)
	fmt.Println("")

	emailOPT := Email{}

	o = Opt{iOpt: &emailOPT}
	o.genAndSendOPT(4)
}
