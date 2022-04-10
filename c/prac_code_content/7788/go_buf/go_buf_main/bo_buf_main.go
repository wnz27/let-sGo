/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/6/27 19:26 6月
 **/
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	//"reflect"
	"unsafe"
)

type T struct {
	A int64
	B float64
}

func bStructToBytes(buffer bytes.Buffer, s interface{}) []byte {
	err := binary.Write(&buffer, binary.BigEndian, s)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func bBytesToSomeStruct(buffer bytes.Buffer) (interface{}, error) {
	// Read into an empty struct.
	t := complexData{}
	err := binary.Read(&buffer, binary.BigEndian, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// binary 包处理二进制
func binarySolve() {
	// Create a struct and write it.
	t := T{A: 0xEEFFEEFF, B: 3.14}
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, t)
	if err != nil {
		panic(err)
	}
	fmt.Println("buf content: ", buf.Bytes())

	// Read into an empty struct.
	t = T{}
	err = binary.Read(buf, binary.BigEndian, &t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x %f", t.A, t.B)
}

type TestStructTobytes struct {
	data int64
}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

// unsafe 包处理二进制
// struct 2 byte
func unsafeSolveStruct2Byte() {

	var testStruct = &TestStructTobytes{100}
	Len := unsafe.Sizeof(*testStruct)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(testStruct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))
	fmt.Println("[]byte is : ", data)
}

// byte 2 struct
func unsafeSolve() {
	var testStruct = &TestStructTobytes{100}
	Len := unsafe.Sizeof(*testStruct)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(testStruct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))
	fmt.Println("[]byte is : ", data)
	var ptestStruct *TestStructTobytes = *(**TestStructTobytes)(unsafe.Pointer(&data))
	fmt.Println("ptestStruct.data is : ", ptestStruct.data)
}

/*
gob 处理二进制
只使用于客户端服务端都使用gob包进行编码和解码的情况。也就是客户端服务端都是go写的，不试用于多种语言。
Gob流不支持函数和通道。试图在最顶层编码这些类型的值会导致失败。
结构体中包含函数或者通道类型的字段的话，会视作非导出字段（忽略）处理。

Gob可以编码任意实现了GobEncoder接口或者encoding.BinaryMarshaler接口的类型的值（通过调用对应的方法），
GobEncoder接口优先。

Gob可以解码任意实现了GobDecoder接口或者encoding.BinaryUnmarshaler接口的类型的值（通过调用对应的方法），
同样GobDecoder接口优先。
 */

type complexData struct {
	N int
	S string
	M map[string]int
	P []byte
	C *complexData
	E Addr
}

type Addr struct {
	Comment string
}

func gobSolve() {

	testStruct := complexData{
		N: 23,
		S: "string data",
		M: map[string]int{"one": 1, "two": 2, "three": 3},
		P: []byte("abc"),
		C: &complexData{
			N: 256,
			S: "Recursive structs? Piece of cake!",
			M: map[string]int{"01": 1, "10": 2, "11": 3},
			E: Addr{
				Comment: "InnerTest123123123123",
			},
		},
		E: Addr{
			Comment: "Test123123123",
		},
	}

	fmt.Println("Outer complexData struct: ", testStruct)
	fmt.Println("Inner complexData struct: ", testStruct.C)
	fmt.Println("Inner complexData struct: ", testStruct.E)
	fmt.Println("===========================")

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(testStruct)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("curr:", enc)

	dec := gob.NewDecoder(&b)
	var data complexData
	fmt.Println("origin", data)
	err = dec.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding GOB data:", err)
		return
	}

	fmt.Println("Outer complexData struct: ", data)
	fmt.Println("Inner complexData struct: ", data.C)
	fmt.Println("Inner complexData struct: ", testStruct.E)

}

// todo 有问题
func StructToBuffer(s interface{}, b bytes.Buffer) error {
	enc := gob.NewEncoder(&b)
	err := enc.Encode(s)
	if err != nil {
		return err
	}
	return nil
}

// todo 有问题
func BufferToStruct(s interface{}, b bytes.Buffer) (interface{}, error) {
	dec := gob.NewDecoder(&b)
	err := dec.Decode(&s)
	if err != nil {
		return nil, err
	}
	//sType := reflect.TypeOf(s)
	//currStruct := (sType).(s)
	return s, err
}


const (
	OpHandshake       = iota // handshake
	OpHandshakeReply  = 1    // handshake reply
	OpHeartbeat              // heartbeat
	OpHeartbeatReply         // heartbeat reply
	OpSendMsg                // send message
	OpSendMsgReply           // send message reply
	OpDisconnectReply        // connection disconnect reply
	OpAuth                   // auth connnect
	OpAuthReply              // auth connect reply
	OpRawBatch               // batch message for websocket
)


type GoimProtocol struct {
	packageLen      int32
	headerLen       int16
	protocolVerSion int16
	operation       int32
	seqId           int32
	body			[]byte
}


func mockGoimProtocol() GoimProtocol {
	return GoimProtocol{
		packageLen: int32(4),
		headerLen: int16(2),
		protocolVerSion: int16(2),
		operation: int32(4),
		seqId: int32(4),
		body: []byte("asdfsfasdfasdfsaf"),
	}
}

func main() {
	//binarySolve()
	//unsafeSolveStruct2Byte()
	//unsafeSolve()

	//gobSolve()
	//testStruct := complexData{
	//	N: 23,
	//	S: "string data",
	//	M: map[string]int{"one": 1, "two": 2, "three": 3},
	//	P: []byte("abc"),
	//	C: &complexData{
	//		N: 256,
	//		S: "Recursive structs? Piece of cake!",
	//		M: map[string]int{"01": 1, "10": 2, "11": 3},
	//		E: Addr{
	//			Comment: "InnerTest123123123123",
	//		},
	//	},
	//	E: Addr{
	//		Comment: "Test123123123",
	//	},
	//}

	//data, _ := json.Marshal(&testStruct)
	//fmt.Println(data)

	//m := make(map[string]interface{})
	//fmt.Println("before: ", m)
	//json.Unmarshal(data, &m)
	//fmt.Println("after: ----> ", m)

	//ttt := complexData{}
	//_ = mapstructure.Decode(m, &ttt)
	//fmt.Println(ttt)

	protocol := mockGoimProtocol()
	fmt.Println(protocol)

	//var b bytes.Buffer
	//bs := bStructToBytes(b, testStruct)
	//fmt.Println(bs)

	//tt, _ := bBytesToSomeStruct(b)
	//fmt.Println(tt)
}
