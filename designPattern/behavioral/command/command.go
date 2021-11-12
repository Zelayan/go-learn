package main

// 命令接口
type command interface {
	execute()
}