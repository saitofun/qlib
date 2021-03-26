package qsock_test

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"time"

	"git.querycap.com/ss/lib/encoding/qjson"
	"git.querycap.com/ss/lib/net/qsock"
	"git.querycap.com/ss/lib/net/qsock/qbuf"
	"git.querycap.com/ss/lib/net/qsock/qmsg"

	"github.com/google/uuid"
)

type parser struct {
}

func (p *parser) Marshal(buf qbuf.Buffer, msg qmsg.Message) error {
	dat, err := qjson.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = buf.Write(dat)
	if err != nil {
		return err
	}
	return nil
}

func (p *parser) Unmarshal(buf qbuf.Buffer) (qmsg.Message, error) {
	var (
		dat    = buf.Bytes()
		header = &Header{}
		msg    qmsg.Message
	)

	err := qjson.Unmarshal(dat, header)
	if err != nil {
		return nil, err
	}
	switch header.Typ {
	case "greeting_req":
		msg = &GreetingReq{}
	case "greeting_rsp":
		msg = &GreetingRsp{}
	default:
		return nil, errors.New("unknown message")
	}

	err = qjson.Unmarshal(dat, msg)
	if err != nil {
		return nil, err
	}
	buf.Shift(len(dat))

	return msg, nil
}

var Parser = &parser{}

var (
	_ qmsg.Message = (*GreetingReq)(nil)
	_ qmsg.Message = (*GreetingRsp)(nil)
)

type ID string

func (v ID) String() string { return string(v) }

type Type string

func (v Type) String() string { return string(v) }

type Header struct {
	Id        string `json:"id"`
	Typ       string `json:"typ"`
	Timestamp int64  `json:"ts"`
}

type GreetingReq struct {
	Header
	Content string `json:"content"`
	Pid     int    `json:"pid"`
}

func NewGreetingReq(content string) *GreetingReq {
	return &GreetingReq{
		Header: Header{
			Id:        uuid.New().String(),
			Typ:       "greeting_req",
			Timestamp: time.Now().UnixNano(),
		},
		Content: content,
		Pid:     os.Getpid(),
	}
}

func (r *GreetingReq) ID() qmsg.ID            { return ID(r.Id) }
func (r *GreetingReq) Type() qmsg.Type        { return Type(r.Typ) }
func (r *GreetingReq) GetTimestamp() int64    { return r.Timestamp }
func (r *GreetingReq) Renew()                 { r.Id, r.Timestamp = uuid.New().String(), time.Now().UnixNano() }
func (r *GreetingReq) RenewWithID(id qmsg.ID) { r.Id, r.Timestamp = id.String(), time.Now().UnixNano() }

type GreetingRsp GreetingReq

func NewGreetingRsp(content string) *GreetingRsp {
	return &GreetingRsp{
		Header: Header{
			Id:        uuid.New().String(),
			Typ:       "greeting_rsp",
			Timestamp: time.Now().UnixNano(),
		},
		Content: content,
		Pid:     os.Getpid(),
	}
}

func (r *GreetingRsp) ID() qmsg.ID            { return ID(r.Id) }
func (r *GreetingRsp) Type() qmsg.Type        { return Type(r.Typ) }
func (r *GreetingRsp) GetTimestamp() int64    { return r.Timestamp }
func (r *GreetingRsp) Renew()                 { r.Id, r.Timestamp = uuid.New().String(), time.Now().UnixNano() }
func (r *GreetingRsp) RenewWithID(id qmsg.ID) { r.Id, r.Timestamp = id.String(), time.Now().UnixNano() }

const (
	MsgTypeGreetingReq uint32 = 1
	MsgTypeGreetingRsp uint32 = 2
)

type tcpParser struct {
}

var TCPParser = &tcpParser{}

func (r *tcpParser) Marshal(buf qbuf.Buffer, msg qmsg.Message) error {
	tmp := make([]byte, 8)
	dat := qjson.UnsafeMarshal(msg)

	switch msg.Type().String() {
	case "greeting_req":
		binary.BigEndian.PutUint32(tmp[4:8], MsgTypeGreetingReq)
	case "greeting_rsp":
		binary.BigEndian.PutUint32(tmp[4:8], MsgTypeGreetingRsp)
	}

	binary.BigEndian.PutUint32(tmp[0:4], uint32(len(dat)))
	buf.Write(tmp)
	buf.Write(dat)
	return nil
}

func (r *tcpParser) Unmarshal(buf qbuf.Buffer) (msg qmsg.Message, err error) {
	var tmp []byte
	tmp, err = buf.Probe(8)
	if err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(tmp[0:4])
	typ := binary.BigEndian.Uint32(tmp[4:8])

	tmp, err = buf.Probe(int(8 + length))
	if err != nil {
		return nil, err
	}

	switch typ {
	case MsgTypeGreetingReq:
		msg = new(GreetingReq)
	case MsgTypeGreetingRsp:
		msg = new(GreetingRsp)
	default:
		return nil, errors.New("unknown message")
	}
	err = qjson.Unmarshal(tmp[8:], msg)
	if err != nil {
		return nil, err
	}
	buf.Shift(int(8 + length))
	return
}

var rsp = &GreetingRsp{
	Header:  Header{Typ: "greeting_rsp"},
	Content: "hello",
	Pid:     os.Getpid(),
}

var typ = Type("greeting_req")

func HandleGreeting(ev *qsock.Event) {
	if ev == nil || ev.Payload() == nil {
		return
	}
	ep := ev.Endpoint()
	pl := ev.Payload()
	rsp.RenewWithID(pl.ID())
	fmt.Printf("%v -> %s\n", ep.ID(), qjson.UnsafeMarshal(pl))
	fmt.Printf("%v <- %s\n", ep.ID(), qjson.UnsafeMarshal(rsp))
	ev.Response(rsp)
}
