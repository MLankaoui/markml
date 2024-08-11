package tests

import (
	"testing"

	"marouane.com/markml/utils"
)

func TestImageLinks(t *testing.T) {
	t.Run("checking for links", func(t *testing.T) {
		got := utils.ImageLinks("[A link](https://markdowntohtml.com)")

		want := "<a href='https://markdowntohtml.com'>"

		assertCorrectMessage(t, got, want)

	})

	t.Run("checking for images", func(t *testing.T) {
		got := utils.ImageLinks("![bears](http://placebear.com/200/200)")

		want := "<img src='http://placebear.com/200/200' alt='bears'>"

		assertCorrectMessage(t, got, want)
	})
}
