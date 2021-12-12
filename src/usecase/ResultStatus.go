package usecase

type ResultStatus struct {
	StatusCode int
	Error      error
}

func NewResultStatus(code int, err error) *ResultStatus {
	return &ResultStatus{StatusCode: code, Error: err}
}
