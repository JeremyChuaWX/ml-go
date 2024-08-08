package mlgo

import (
	"log/slog"
	"os"
)

type linear struct {
	w  *matrix
	b  *matrix
	x  *matrix
	dw *matrix
	db *matrix
}

func Linear(in int, out int) *linear {
	return &linear{
		w: RandomMatrix(out, in),
		b: RandomMatrix(out, 1),
	}
}

func (l *linear) forward(x *matrix) *matrix {
	l.x = Copy(x)
	res, err := Dot(l.w, x)
	if err != nil {
		slog.Error("forward error", "error", err)
		os.Exit(1)
	}
	res.Add(l.b)
	return res
}

func (l *linear) backward(y *matrix) *matrix {
	dw, err := Dot(y, l.x)
	if err != nil {
		slog.Error("backward error calculating dw", "error", err)
		os.Exit(1)
	}
	dx, err := Dot(l.w, y)
	if err != nil {
		slog.Error("backward error calculating dx", "error", err)
		os.Exit(1)
	}
	l.dw = dw
	l.db = RowSum(y)
	return dx
}
