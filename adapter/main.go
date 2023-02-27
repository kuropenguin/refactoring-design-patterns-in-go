package main

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLiightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLiightningConnectorIntoComputer(windowsMachineAdapter)
}
