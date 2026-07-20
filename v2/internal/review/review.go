package review

type Reviewer interface {
	Review(question string) (string, error)
}
