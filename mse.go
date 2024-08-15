package mlgo

// functions here return the first derivative of the respective cost functions

func MSE(actual *matrix, pred *matrix) *matrix {
	res := Copy(pred)
	res.Subtract(actual)
	return res
}
