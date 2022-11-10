package gerrors

var (
	ErrDB   = New(100001, "operation DB error.")
	ErrAuth = New(100002, "authentication error.")
	ErrCall = New(100003, "call third party error.")
)
