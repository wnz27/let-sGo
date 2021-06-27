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
	version int32
	op      int32
	seq     int32
	body    []byte
}

var contentString = "Hello im world!"

func mockGoimProtocol() GoimProtocol {
	return GoimProtocol{
		version: int32(4),
		op:      int32(7),
		seq:     int32(4),
		body:    []byte(contentString),
	}
}

func GoimProtocalEecoder() error {
	return nil
}

func GoimProtocalDecoder(reader *bufio.Reader) (*GoimProtocol, error) {
	// 读取消息的长度
	var (
		bodyLen   int
		headerLen int16
		packLen   int32
		buf       []byte
	)
	if buf, err := reader.Peek(rawHeaderSize); err != nil {
		return
	}
	packLen = binary.BigEndian.Int32(buf[_packOffset:_headerOffset])
	headerLen = binary.BigEndian.Int16(buf[_headerOffset:_verOffset])
	p.Ver = int32(binary.BigEndian.Int16(buf[_verOffset:_opOffset]))
	p.Op = binary.BigEndian.Int32(buf[_opOffset:_seqOffset])
	p.Seq = binary.BigEndian.Int32(buf[_seqOffset:])
	if packLen > _maxPackSize {
		return ErrProtoPackLen
	}
	if headerLen != _rawHeaderSize {
		return ErrProtoHeaderLen
	}
	if bodyLen = int(packLen - int32(headerLen)); bodyLen > 0 {
		p.Body, err = rr.Pop(bodyLen)
	} else {
		p.Body = nil
	}
	return
}
