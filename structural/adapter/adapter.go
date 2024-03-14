package main

import "fmt"

type (
	Client   struct{}
	Computer interface {
		InsertIntoLightningPort()
	}
	Mac            struct{}
	Windows        struct{}
	WindowsAdapter struct {
		windowsMachie *Windows
	}
)

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine")
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lighting signal to USB.")
	w.windowsMachie.InsertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowsMachie: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
