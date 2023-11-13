package adapter

// 适配器
func Do() {
	client := &Client{}
	mac := &Mac{}
	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{windows}

	client.InsertLightingConnectorIntoComputer(mac)
	client.InsertLightingConnectorIntoComputer(windowsAdapter)
}
