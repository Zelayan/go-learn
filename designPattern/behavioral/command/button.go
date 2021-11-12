package main

// 请求者
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}