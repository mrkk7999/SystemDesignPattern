package main

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}

	windowsMachineAdapter := &windowsAdapter{windowMachine: windowsMachine}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
