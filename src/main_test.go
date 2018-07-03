package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var stdin = `1
	2 3
	test`
	var expected = "6 test"
	actual, _ := StubIO(stdin, func() {
		main()
	})

	assert.Equal(t, expected, actual)
}

func TestMain2(t *testing.T) {
	var stdin = `72
	128 256
	myonmyon`
	var expected = "456 myonmyon"
	actual, _ := StubIO(stdin, func() {
		main()
	})

	assert.Equal(t, expected, actual)
}

// StubIO stubs Stdin Stdout Stderr in 'fn'.return Stdout and Stderr
func StubIO(inbuf string, fn func()) (string, string) {
	inr, inw, _ := os.Pipe()
	outr, outw, _ := os.Pipe()
	errr, errw, _ := os.Pipe()
	orgStdin := os.Stdin
	orgStdout := os.Stdout
	orgStderr := os.Stderr
	inw.Write([]byte(inbuf))
	inw.Close()
	os.Stdin = inr
	os.Stdout = outw
	os.Stderr = errw
	fn()
	os.Stdin = orgStdin
	os.Stdout = orgStdout
	os.Stderr = orgStderr
	outw.Close()
	outbuf, _ := ioutil.ReadAll(outr)
	errw.Close()
	errbuf, _ := ioutil.ReadAll(errr)

	return string(outbuf), string(errbuf)

}
