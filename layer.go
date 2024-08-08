package mlgo

type layer interface {
	forward(*matrix) *matrix
	backward(*matrix) *matrix
}
