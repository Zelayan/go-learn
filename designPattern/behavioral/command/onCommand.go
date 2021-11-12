package main

// 打开接口
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}
