package main_test

import (
	"testing"

	main "github.com/launchquickly/go-github-actions-starter"
)

func TestHello(t *testing.T) {
	got := main.Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
