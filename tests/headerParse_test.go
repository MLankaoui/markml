package tests

import (
	"testing"

	"marouane.com/markml/utils"
)

func TestHeaderParse(t *testing.T) {
	t.Run("if the symbole is #", func(t *testing.T) {
		got := utils.HeaderParser("# Welcome to our html page")

		want := "<h1>Welcome to our html page</h1>"

		assertCorrectMessage(t, got, want)
	})
	t.Run("if the symbole is ###", func(t *testing.T) {
		got := utils.HeaderParser("## Welcome to our html page")

		want := "<h2>Welcome to our html page</h2>"

		assertCorrectMessage(t, got, want)
	})

	t.Run("if the symbole is ###", func(t *testing.T) {
		got := utils.HeaderParser("### Welcome to our html page")

		want := "<h3>Welcome to our html page</h3>"

		assertCorrectMessage(t, got, want)
	})

	t.Run("if the symbole is ####", func(t *testing.T) {
		got := utils.HeaderParser("#### Welcome to our html page")

		want := "<h4>Welcome to our html page</h4>"

		assertCorrectMessage(t, got, want)
	})

	t.Run("if the symbole is #####", func(t *testing.T) {
		got := utils.HeaderParser("##### Welcome to our html page")

		want := "<h5>Welcome to our html page</h5>"

		assertCorrectMessage(t, got, want)
	})

	t.Run("if the symbole is ######", func(t *testing.T) {
		got := utils.HeaderParser("###### Welcome to our html page")

		want := "<h6>Welcome to our html page</h6>"

		assertCorrectMessage(t, got, want)

	})
}

func BenchmarkHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.HeaderParser("# welcome in")
	}
}
