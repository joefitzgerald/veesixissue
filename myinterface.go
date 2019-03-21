package countint

//go:generate gobin -m -run github.com/maxbrunsfeld/counterfeiter/v6 . MySpecialInterface

type MySpecialInterface interface {
	DoThings(string, uint64) (int, error)
}
