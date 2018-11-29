package auth

type AccessLevel int

const (
	BasicUser AccessLevel = iota
	Admin
)