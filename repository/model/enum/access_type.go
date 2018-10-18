package enum

type AccessType int

func (accessType AccessType) String() string {
	names := [...]string{
		"GET",
		"POST",
	}
	return names[accessType]
}

const (
	get  AccessType = 0
	post AccessType = 1
)
