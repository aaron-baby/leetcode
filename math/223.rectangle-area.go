package math

import "gitlab.com/aaw/leetcode/lib"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	// no overlap
	if B >= H || // top
		A >= G || // right
		D <= F || // bottom
		C <= E {
		return (C-A)*(D-B) + (G-E)*(H-F)
	}
	// overlap area bottom left(x1, y1)
	x1, y1, x2, y2 := lib.Max(A, E), lib.Max(B, F), lib.Min(C, G), lib.Min(D, H)
	return (C-A)*(D-B) + (G-E)*(H-F) - (x2-x1)*(y2-y1)
}
