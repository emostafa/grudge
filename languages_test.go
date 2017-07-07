package grudge

import (
	"fmt"
	"testing"
)

func TestPython(t *testing.T) {
	out, err := Python("print('Hello Python')", []string{})
	if err != nil {
		t.Error(err.Error())
		fmt.Println(err)
		t.Fail()
	}
	t.Log(out)
}

func TestJavascript(t *testing.T) {
	out, err := Javascript("console.log('Hello Javascript');", []string{})
	if err != nil {
		t.Fatal(err.Error())
		t.Fail()
	}
	t.Log(out)
}

func TestRuby(t *testing.T) {
	out, err := Ruby("puts 'Hello Ruby'", []string{})
	if err != nil {
		t.Fatal(err.Error())
		t.Fail()
	}
	t.Log(out)
}
