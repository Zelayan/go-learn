package main

// 关闭接口
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}
