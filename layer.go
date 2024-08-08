package mlgo

type layer interface {
	forward(*matrix) *matrix
	backward(*matrix) *matrix
	update(lr float64)
}
