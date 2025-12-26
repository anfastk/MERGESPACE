package limiter

type Rule struct {
	Name      string
	KeyPrefix string
	Limit     int
	WindowSec int
	Algo      string
	FailOpen  bool
}
