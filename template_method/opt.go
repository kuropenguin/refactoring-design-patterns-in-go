package main

type IOpt interface {
	genRandomOPT(int) string
	saveOPTCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Opt struct {
	iOpt IOpt
}

func (o *Opt) genAndSendOPT(optLength int) error {
	opt := o.iOpt.genRandomOPT(optLength)
	o.iOpt.saveOPTCache(opt)
	msg := o.iOpt.getMessage(opt)
	err := o.iOpt.sendNotification(msg)
	if err != nil {
		return err
	}
	return nil
}
