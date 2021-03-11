package main

type Result struct {
	ts        string
	Maxguests int
}

func maximumNum(ts []string, n int) Result {
	if len(ts) == 0 {
		return Result{"", 0}
	}
	var arrv, leave []string
	for i := 0; i < len(ts); i++ {
		if i%2 == 0 {
			arrv = append(arrv, ts[i])
		} else {
			leave = append(leave, ts[i])
		}
	}
	var rst Result
	guests_in := 1
	rst = Result{arrv[0], 1}
	i, j := 1, 0
	for i < n && j < n {
		// next arrival small than current leave
		if arrv[i] <= leave[j] {
			guests_in++
			if guests_in > rst.Maxguests {
				rst = Result{arrv[i], guests_in}
			}
			i++
		} else {
			guests_in--
			j++
		}
	}
	return rst
}
