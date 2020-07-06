package main

import (
	"fmt"
	"gitlab.com/aaw/leetcode/ds"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ds.ListNode{Val: 2}
	l1.Next = &ds.ListNode{
		Val:  4,
		Next: &ds.ListNode{Val: 3},
	}

	l2 := &ds.ListNode{Val: 5}
	l2.Next = &ds.ListNode{
		Val:  6,
		Next: &ds.ListNode{Val: 4},
	}

	got := addTwoNumbers(l1, l2)
	var s string
	for got != nil {
		s += fmt.Sprint(got.Val)
		got = got.Next
	}
	want := "708"
	if s != want {
		t.Errorf("got %v want %v given", s, want)
	}
}
