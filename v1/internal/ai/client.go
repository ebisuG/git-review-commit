package ai

type Reviewer interface {
	Review(question string) (string, error)
}
