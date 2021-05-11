package qipc_test

import (
	"errors"
	"os"

	"git.querycap.com/ss/lib/encoding/qjson"
	"git.querycap.com/ss/lib/net/qsock/qbuf"
	"git.querycap.com/ss/lib/net/qsock/qmsg"
	"git.querycap.com/ss/lib/os/qipc"
	"github.com/google/uuid"
)

type Type string

func (t Type) String() string {
	return string(t)
}

type ID string

func (v ID) String() string {
	return string(v)
}

type req struct {
	Id     ID     `json:"id"`
	Typ    Type   `json:"typ"`
	ProcId int    `json:"pid"`
	Msg    string `json:"msg"`
	Src    int    `json:"src"`
	Dst    int    `json:"dst"`
}

func (v *req) Pid() int        { return v.ProcId }
func (v *req) ID() qmsg.ID     { return v.Id }
func (v *req) Type() qmsg.Type { return Type("req") }
func (v *req) SrcPid() int     { return v.Src }
func (v *req) DstPid() int     { return v.Dst }

type rsp struct{ req }

func (v *rsp) Type() qmsg.Type { return Type("rsp") }

type parser struct{}

var Parser = &parser{}

func (p *parser) Marshal(buf qbuf.Buffer, msg qmsg.Message) error {
	dat, err := qjson.Marshal(msg)
	if err != nil {
		return err
	}
	buf.Write(dat)
	return nil
}

func (p *parser) Unmarshal(buf qbuf.Buffer) (qmsg.Message, error) {
	var (
		dat = buf.Bytes()
		msg qmsg.Message
	)
	var prob = &struct {
		Typ Type `json:"typ"`
	}{}
	var err = qjson.Unmarshal(dat, prob)
	if err != nil {
		return nil, err
	}
	switch prob.Typ {
	case "req":
		msg = &req{}
	case "rsp":
		msg = &rsp{}
	default:
		return nil, errors.New("unknown")
	}

	if err = qjson.Unmarshal(dat, msg); err != nil {
		return nil, err
	}
	return msg, nil
}

var _ qipc.IPCMessage = (*req)(nil)
var _ qipc.IPCMessage = (*rsp)(nil)
var _ qmsg.Parser = (*parser)(nil)

func NewReq() *req {
	return &req{
		Id:     ID(uuid.New().String()),
		Typ:    "req",
		ProcId: os.Getpid(),
		Msg:    "",
	}
}

func NewRsp() *rsp {
	req := NewReq()
	req.Typ = "rsp"
	return &rsp{*req}
}
