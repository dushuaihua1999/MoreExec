package exec

// Exceuter 执行者对象
type Exceuter interface {
	Init()
	Run() error
	Continue() error
	Stop() error
}

// Grain 标识粒度对象，执行器只识别这个对象,进行进度的控制
type Grain interface {
	Info() string   // 展示当前粒度的信息
	State() int     // 当前粒度的状态 0未执行 1执行成功 2执行失败 3丢弃
	Identify() bool // 验证接口,验证当前粒度是否符合要求
}
