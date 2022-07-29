package send

type Sender interface {
	BroadCast()
	Run()
}
