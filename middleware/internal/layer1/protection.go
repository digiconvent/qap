package layer1

type access int

const (
	anyone access = iota
	loggedIn
	loggedInOrIntent
	intentOnly
)

type protection struct {
	guard       *guard
	access      access
	permissions []string
}
