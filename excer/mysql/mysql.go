package mysql

import (
	"MoreExec/exec"
	"MoreExec/send"
)

// BaseAuth 基础授权
type BaseAuth struct {
	User     string
	Password string
	IP       string
	Port     string
}

// MySQLExcer mysql命令执行者
type MySQLExcer struct {
	reciver chan exec.Grain
	auth    BaseAuth
	sender  send.Sender // 管理Excer的sender指针
}

func (m *MySQLExcer) Init(auth BaseAuth, opts ...Option) {
	my := &MySQLExcer{}
	my.auth = auth
	for _, op := range opts {
		op(my)
	}
	if m.reciver == nil {
		m.reciver = make(chan exec.Grain)
	}
}

func (m *MySQLExcer) Run() error {

}

func (m *MySQLExcer) Continue() error {

}

func (m *MySQLExcer) Stop() error {

}
