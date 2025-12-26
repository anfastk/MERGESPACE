package limiter

var SignupIPRule = Rule{
	Name:      "signup_ip",
	KeyPrefix: "signup:ip",
	Limit:     5,
	WindowSec: 600,
	Algo:      "token_bucket",
	FailOpen:  false,
}

var SignupEmailRule = Rule{
	Name:      "signup_email",
	KeyPrefix: "signup:email",
	Limit:     3,
	WindowSec: 3600,
	Algo:      "token_bucket",
	FailOpen:  false,
}
