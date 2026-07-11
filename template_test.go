package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	want := assert.New(t)

	greeting, err := Greet("gomatic")
	want.NoError(err)
	want.Equal("Hello, gomatic!", greeting)
}

func TestGreetRejectsEmptyName(t *testing.T) {
	want := assert.New(t)

	greeting, err := Greet("")
	want.ErrorIs(err, ErrEmptyName)
	want.Empty(greeting)
}

func FuzzGreet(f *testing.F) {
	f.Add("gomatic")
	f.Add("")
	f.Fuzz(func(t *testing.T, name string) {
		greeting, err := Greet(Name(name))
		if name == "" {
			assert.ErrorIs(t, err, ErrEmptyName)
			return
		}
		assert.NoError(t, err)
		assert.Equal(t, "Hello, "+name+"!", greeting)
	})
}

func TestErrorMessage(t *testing.T) {
	assert.Equal(t, "name is empty", ErrEmptyName.Error())
}
