package enum

type ContentType int

func (contetType ContentType) String() string {
	names := [...]string{
		"application/json",
		"application/msgpack",
	}
	return names[contetType]
}

const (
	json ContentType = iota
	msgpack
)
