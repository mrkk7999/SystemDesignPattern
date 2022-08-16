package main

import "fmt"

type windowsAdapter struct {
	windowMachine *Windows
}

func (w *windowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightnig signal to USB")
	w.windowMachine.insertIntoUSBPort()
}
