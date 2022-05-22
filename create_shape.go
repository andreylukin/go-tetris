package main

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

func Makebox(s tcell.Screen) {
	w, h := s.Size()

	if w == 0 || h == 0 {
		return
	}

	lx := rand.Int() % w
	ly := rand.Int() % h
	lw := rand.Int() % (w - lx)
	lh := rand.Int() % (h - ly)
	st := tcell.StyleDefault
	gl := ' '

	st = st.Background(tcell.Color(rand.Int() % s.Colors()))

	for row := 0; row < lh; row++ {
		for col := 0; col < lw; col++ {
			s.SetCell(lx+col, ly+row, st, gl)
		}
	}
	s.Show()
}
