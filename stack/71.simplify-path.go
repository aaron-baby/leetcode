package stack

import (
	"strings"
)

func simplifyPath(path string) string {
	s := strings.Split(path, "/")
	var res []string
	for _, v := range s {
		if v == "" || v == "." {
			continue
		}
		if v == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
			continue
		}
		res = append(res, v)
	}
	return "/" + strings.Join(res, "/")
}
