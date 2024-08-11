package tests

import (
	"testing"

	"marouane.com/markml/utils"
)

func TestDecoration(t *testing.T) {
	t.Run("checking for italic", func(t *testing.T) {
		got := utils.Decoration("*word*")

		want := "<em>word</em>"

		assertCorrectMessage(t, got, want)
	})

	t.Run("checking for bold", func(t *testing.T) {
		got := utils.Decoration("**word**")

		want := "<bold>word</bold>"

		assertCorrectMessage(t, got, want)

	})

	t.Run("checking for your mom", func(t *testing.T) {
		got := utils.Decoration("word")

		want := "word"

		assertCorrectMessage(t, got, want)

	})

	t.Run("checking for stritrough", func(t *testing.T) {
		got := utils.Decoration("~~word~~")

		want := "<s>word</s>"

		assertCorrectMessage(t, got, want)

	})
}
