package main

import (
	"errors"
	"fmt"
)

type RongyunSendMessage interface {
	Send(fromUser string, toUser string, Content string) error
}
type RongyunSendImgMessage struct{}

func (r RongyunSendImgMessage) Send(fromUser string, toUser string, Content string) error {
	// 判断用户状态
	// 判断用户权限
	// 整理发送数据
	// 发送消息
	// 消息记录入库
	fmt.Printf("fromUser [%v] => toUser [%s] : content [%v]\n", fromUser, toUser, Content)
	return nil
}

type RongyunSendImgsMessage struct{}

func (r RongyunSendImgsMessage) Send(fromUser string, toUser string, Content string) error {
	// 判断用户状态
	// 判断用户权限
	// 整理发送数据
	// 发送消息
	// 消息记录入库
	fmt.Printf("fromUser [%v] => toUser [%s] : content [%v]\n", fromUser, toUser, Content)
	return nil
}

type RongyunSendVideoMessage struct{}

func (r RongyunSendVideoMessage) Send(fromUser string, toUser string, Content string) error {
	// 判断用户状态
	// 判断用户权限
	// 整理发送数据
	// 发送消息
	// 消息记录入库
	fmt.Printf("fromUser [%v] => toUser [%s] : content [%v]\n", fromUser, toUser, Content)
	return nil
}

type MessageParams struct {
	Type     string // 消息类型：img=单图，imgs=多图，video=视频 ...
	Content  string // 消息内容
	FromUser string // 发送方
	ToUser   string // 接收方
}

func main() {
	// 发送单图
	sendImg := MessageParams{
		Type:     "img",
		Content:  "发送单图消息",
		FromUser: "A",
		ToUser:   "B",
	}
	// 发送多图
	sendImgs := MessageParams{
		Type:     "imgs",
		Content:  "发送多图",
		FromUser: "A",
		ToUser:   "B",
	}
	// 发送视频
	sendVideo := MessageParams{
		Type:     "video",
		Content:  "发送视频",
		FromUser: "A",
		ToUser:   "B",
	}
	SendMessage(sendImg)   // 单图
	SendMessage(sendImgs)  // 多图
	SendMessage(sendVideo) // 视频
}

var Template = map[string]RongyunSendMessage{
	"img":   new(RongyunSendImgMessage),
	"imgs":  new(RongyunSendImgsMessage),
	"video": new(RongyunSendVideoMessage),
}

func SendMessage(params MessageParams) error {
	if _, ok := Template[params.Type]; !ok {
		return errors.New("tagID invalid")
	}
	return Template[params.Type].Send(params.FromUser, params.ToUser, params.Content)
}
