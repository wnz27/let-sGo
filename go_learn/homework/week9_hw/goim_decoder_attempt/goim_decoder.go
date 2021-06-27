/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/6/27 15:56 6月
 **/
package goim_decoder_attempt

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

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

const (
	pLen       = 4
	headerLen  = 2
	versionLen = 2
	opLen      = 4
	seqIdLen   = 4

	rawHeaderSize = pLen + headerLen + versionLen + opLen + seqIdLen
	pkgOffset     = 0
	headerOffset  = pkgOffset + pLen
	versionOffset = headerOffset + headerLen
	opOffset      = versionOffset + versionLen
	seqOffset     = opOffset + opLen
)

type GoimProtocol struct {
	Version int32
	Op      int32
	Seq     int32
	Body    []byte
}

var contentString = "Hello im world!"

func mockGoimProtocol() GoimProtocol {
	return GoimProtocol{
		Version: int32(4),
		Op:      int32(7),
		Seq:     int32(4),
		Body:    []byte(contentString),
	}
}

func GoimProtocalEecoder() error {
	return nil
}

func GoimProtocalDecoder(reader *bufio.Reader) GoimProtocol {
	// 读取消息的长度
	var (
		bodyLen   int
		headerLen int16
		packLen   int32
	)

	if buff, err := reader.Peek(rawHeaderSize); err != nil {
		panic(err)
	} else {
		lenPkgBuff := bytes.NewBuffer(buff[pkgOffset: headerOffset])
		err := binary.Read(lenPkgBuff, binary.BigEndian, &packLen)
		if err != nil {
			panic(err)
		}

		lenHeaderBuff := bytes.NewBuffer(buff[headerOffset: versionOffset])
		err1 := binary.Read(lenHeaderBuff, binary.BigEndian, &headerLen)
		if err1 != nil {
			panic(err1)
		}

		var gp = GoimProtocol{}

		lenVerBuff := bytes.NewBuffer(buff[versionOffset: opOffset])
		err2 := binary.Read(lenVerBuff, binary.BigEndian, &gp.Version)
		if err2 != nil {
			panic(err2)
		}

		lenOpBuff := bytes.NewBuffer(buff[opOffset: seqOffset])
		err3 := binary.Read(lenOpBuff, binary.BigEndian, &gp.Op)
		if err3 != nil {
			panic(err3)
		}

		lenSeqBuff := bytes.NewBuffer(buff[seqOffset: ])
		err4 := binary.Read(lenSeqBuff, binary.BigEndian, &gp.Seq)
		if err4 != nil {
			panic(err4)
		}

		if bodyLen = int(packLen - int32(headerLen)); bodyLen > 0 {
			gp.Body, _ = reader.Peek(bodyLen)
		} else {
			gp.Body = nil
		}
		return gp

	}
}
