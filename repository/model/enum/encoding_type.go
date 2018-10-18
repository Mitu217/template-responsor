package enum

type EncodingType int

func (encodingType EncodingType) String() string {
	names := [...]string{
		"application/json",
		"application/msgpack",
	}
	return names[encodingType]
}

const (
	json    EncodingType = 0
	msgpack EncodingType = 1
)
