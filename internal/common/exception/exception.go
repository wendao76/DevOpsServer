package exception

const (
	ClientErr = 400
	ServerErr = 500
	AuthErr = 403
)

type Exception interface {
	error
	Coder() string
}

type Business struct {
    Code string	`json:"code"`
    Msg string	`json:"msg"`
}

func (b *Business) Error() string {
    return b.Msg
}

func (b *Business) Coder() string {
    return b.Code
}
