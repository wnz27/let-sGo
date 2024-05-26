/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-05-24 02:56:11
 * @LastEditTime: 2024-05-25 08:15:15
 * @FilePath: /let-sGo/c/prac_code_content/7788/go_event/g_eb2.go
 * @description: type some description
 */
package main

import (
	"github.com/werbenhu/eventbus"
)

type WMSEventMsg struct {
	GroupID   int64
	StationID string
	BatchID   string
}

func (msg WMSEventMsg) GroupAlreadyCreated(batchID string) bool {
	return msg.GroupID > 0 && msg.BatchID == batchID
}

func (msg WMSEventMsg) StationAlreadyCreated(batchID string) bool {
	return msg.StationID != "" && msg.BatchID == batchID
}

// 后期可能可以挂 redis client 利用消息上的 uniqID做消息幂等消费
type WMSEventBusManager2 struct {
	EventBus *eventbus.EventBus
}

func NewWMSEventBusManager2() *WMSEventBusManager2 {
	return &WMSEventBusManager2{
		EventBus: eventbus.New(),
	}
}

// 外部处理该错误,暂定为打日志
func (bus *WMSEventBusManager2) PublishToEventBus(topic string, msg WMSEventMsg) error {
	return bus.EventBus.Publish(topic, msg)
}
