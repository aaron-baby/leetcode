package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{
		Val:  4,
		Next: &ListNode{Val: 3},
	}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{
		Val:  6,
		Next: &ListNode{Val: 4},
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
