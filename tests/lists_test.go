package tests

import "testing"

func TestLists(t *testing.T) {
	got := OrderedList("1. test")

	want := "<ol><li>test</li></ol>"

	assertCorrectMessage(t, got, want)
}
