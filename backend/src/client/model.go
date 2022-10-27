package client

type ChangeStateResult struct {
	Body       ChangeStateBodyResult
	StatusCode int
}

type ChangeStateBodyResult struct {
	Message string
}
