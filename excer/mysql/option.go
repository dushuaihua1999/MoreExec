package mysql

import "MoreExec/exec"

type Option func(excer *MySQLExcer)

func WithReciverSize(n int) Option {
	return func(excer *MySQLExcer) {
		excer.reciver = make(chan exec.Grain, n)
	}
}
