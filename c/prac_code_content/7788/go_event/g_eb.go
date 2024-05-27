/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-05-24 01:48:45
 * @LastEditTime: 2024-05-26 14:13:09
 * @FilePath: /let-sGo/c/prac_code_content/7788/go_event/g_eb.go
 * @description: type some description
 */
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/werbenhu/eventbus"
)

type WMSEventBusManager struct {
	GroupMsgPipe   *eventbus.Pipe[string] // group 事件管道
	StationMsgPipe *eventbus.Pipe[string] // station 事件管道
}

func NewWMSEventBusManager() *WMSEventBusManager {
	return &WMSEventBusManager{
		GroupMsgPipe:   eventbus.NewPipe[string](),
		StationMsgPipe: eventbus.NewPipe[string](),
	}
}

// 外部处理该错误,暂定为打日志
func (bus *WMSEventBusManager) PublishToGroupEventBus(msg string) error {
	return bus.GroupMsgPipe.Publish(msg)
}

// 外部处理该错误暂定为打日志
func (bus *WMSEventBusManager) PublishToStationEventBus(msg string) error {
	return bus.StationMsgPipe.Publish(msg)
}

func handler1(val string) {

	fmt.Printf("handler1 val:%s\n", val)
}

func handler2(val string) {
	fmt.Printf("handler2 val:%s\n", val)
}

func example1() {
	pipe := eventbus.NewPipe[string]()
	pipe.Subscribe(handler1)
	pipe.Subscribe(handler2)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			pipe.Publish(strconv.Itoa(i))
		}
		for i := 10; i < 20; i++ {
			pipe.PublishSync(strconv.Itoa(i))
		}
		wg.Done()
	}()
	wg.Wait()

	time.Sleep(time.Millisecond)
	pipe.Unsubscribe(handler1)
	pipe.Unsubscribe(handler2)
	pipe.Close()
}

type Param struct {
	GroupID int
	AType   string
	fn      func(a int) int
	bus     *eventbus.EventBus
}

func handler(topic string, payload Param) error {
	fmt.Printf("topic:%s, payload:%v\n", topic, payload)
	a1 := payload.fn(123)
	fmt.Println("xxxxx", a1)
	e2 := payload.bus.Unsubscribe(topic, handler)
	fmt.Println("-->", e2)
	return nil
}

func example2() {
	bus := eventbus.New()

	// Subscribe() 订阅一个主题，如果handler不是函数则返回错误。
	// handler必须有两个参数：第一个参数必须是字符串类型，
	// handler的第二个参数类型必须与 `Publish()` 中的 payload 类型一致。
	bus.Subscribe("testtopic", handler)
	bus.Subscribe("testtopic2", handler)
	// 异步方式发布消息
	bus.Publish("testtopic", Param{GroupID: 100, AType: "test", fn: func(a int) int {
		return a + 1
	},
		bus: bus,
	})

	// 同步方式发布消息
	bus.PublishSync("testtopic2", Param{GroupID: 200, AType: "prod123", fn: func(a int) int {
		return a + 2
	},
		bus: bus,
	})

	// 订阅者接收消息。为了确保订阅者可以接收完所有消息的异步消息，这里在取消订阅之前给了一点延迟。
	time.Sleep(time.Millisecond)
	e1 := bus.Unsubscribe("testtopic", handler)
	if e1 != nil {
		fmt.Println("=========> ", e1.Error())
	}

	bus.Close()
}

func main() {
	example2()
}
