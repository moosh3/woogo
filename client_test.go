package woogo

import (
	"context"
	"net/http"
	"testing"
)

type M map[string]interface{}

type L []interface{}

/* Table based testing in Go

type fibTest struct {
        n        int
        expected int
}

var fibTests = []fibTest {
        {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Error("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
*/

var c = Client{}

func TestNewClient(t *testing.T) {
	want := new(http.Client)
	got := NewClient(want) // pointer
	if want != got {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestNewRequest(t *testing.T) {
	method := "POST"
	path := "http://test.com/"
	var body interface{}

	want, err := new(Response)
	got := c.NewRequest(method, path, body)
	if want != got {
		t.Fatalf("Expected %s, got %s", want, got)
	}
	assert.Equal(err, nil)
}

func TestDo(t *testing.T) {
	assert := assert.New()
	ctx := context.Context()
	var v interface{}

	want := new(Response)
	got, err := c.Do(ctx, *want, v)
	if want != got {
		t.Fatalf("Expected %s, got %s", want, got)
	}
	assert.Equal(err, nil)
}

func Test_newResponse(t *testing.T) {
	want := new(Response)
	got := c.newResponse(*want)
	if want != got {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestCheckResponse(t *testing.T) {
	assert := assert.New()
	want := ""
	r := new(Response)
	got, err := c.CheckResponse(r)
	if want != got {
		t.Fatalf("Expected %s, got %s", want, got)
	}
	assert.Equal(err, nil)
}

func TestWithContext(t *testing.T) {
	return nil
}

func TestBool(t *testing.T) {
	want := false
	got := Bool(want)
	// wants pointer should point to the same block as got
	if &want != got {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestInt(t *testing.T) {
	want := 1
	got := Int(want)
	// wants pointer should point to the same block as got
	if &want != got {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}

func TestString(t *testing.T) {
	want := "Test string"
	got := String(want)
	// wants pointer should point to the same block as got
	if &want != got {
		t.Fatalf("Expected %s, got %s", want, got)
	}
}
