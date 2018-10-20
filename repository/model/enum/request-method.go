package enum

type RequestMethod int

func (requestMethod RequestMethod) String() string {
	names := [...]string{
		"GET",
		"POST",
	}
	return names[requestMethod]
}

const (
	GET RequestMethod = iota
	POST
)
