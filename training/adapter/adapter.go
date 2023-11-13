package adapter

import "fmt"

// 适配器模式
// 客户端代码
type Client struct {
}

func (c *Client) InsertLightingConnectorIntoComputer(com Computer) {
	fmt.Printf("Inserting lightning connector into computer\n")
	com.InsertIntoLightningPort()
}

// 客户端接口
type Computer interface {
	InsertIntoLightningPort()
}

// mac 服务
type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Printf("Inserting lightning connector into mac\n")
}

// windows 服务
type Windows struct {
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Printf("Inserting lightning connector into windows\n")
}

// 适配器
type WindowsAdapter struct {
	w *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	w.w.InsertIntoUSBPort()
}
