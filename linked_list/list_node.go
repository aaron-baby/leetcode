package linked_list

import (
	"strconv"
	"strings"
)

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}
)

func (l *ListNode) String() string {
	var s string
	var n *ListNode
	// int is visited element index
	visited := make(map[*ListNode]int)
	visited[l] = 1
	n = l.Next
	s = strconv.Itoa(l.Val)
	for n != nil {
		i, ok := visited[n]
		// break circular visit node
		if ok {
			fl := len(s)
			s += "\n" + strings.Repeat(" ", i-1) + "|" + strings.Repeat(" ", fl-i-1) + "|"
			s += "\n" + strings.Repeat(" ", i-1) + strings.Repeat("-", fl-i+1)
			return s
		}
		s += " -> "
		s += strconv.Itoa(n.Val)
		visited[n] = len(s)
		n = n.Next
	}

	return s
}
