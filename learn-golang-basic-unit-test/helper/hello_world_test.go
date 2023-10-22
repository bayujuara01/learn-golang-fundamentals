package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// Per Packages
func TestMain(m *testing.M) {
	// Before Test
	fmt.Println("Before Unit Test")

	m.Run()

	// After Test
	fmt.Println("After Unit Test")
}

func TestSubTestHelloWorld(t *testing.T) {
	t.Run("Sub Test Hello 1", func(t *testing.T) {
		require.Equal(t, "Hello, Bayu", HelloWorld("Bayu"), "Parameter must be Bayu")
	})

	t.Run("Sub Test Hello 2", func(t *testing.T) {
		require.Equal(t, "Hello, Bayu", HelloWorld("Bayu"), "Parameter must be Bayu")
	})
}

type TableHelloWorld struct {
	name     string
	request  string
	expected string
}

func TestTableTestHelloWorld(t *testing.T) {
	tableTest := []TableHelloWorld{
		{name: "Table Test 1", request: "LALA", expected: "Hello, LALA"},
		{name: "Table Test 2", request: "LALAX", expected: "Hello, LALAX"},
		{name: "Table Test 3", request: "LALAZ", expected: "Hello, LALAZ"},
	}

	for _, test := range tableTest {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, HelloWorld(test.request))
		})
	}
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin/amd64" {
		t.Skip("Can't run on Mac OS")
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bayu")
	assert.Equal(t, "Hello, Bayu", result, "Result must be 'Hello, Bayu")

	// Selanjutnya akan tetap dieksekusi
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Bayu")
	require.Equal(t, "Hello, Bayu", result, "Result must be 'Hello, Bayu")

	// Selanjutnya akan tidak akan dieksekusi
}

func TestHelloWorldBayu(t *testing.T) {
	//result := HelloWorld("Bay")
	result := HelloWorld("Bayu")
	if result != "Hello, Bayu" {
		t.FailNow()
	}

	fmt.Println("TestHelloWorldBayu Finish") // Tidak akan dieksekusi
}
