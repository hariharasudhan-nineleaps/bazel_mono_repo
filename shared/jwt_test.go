package shared

import "testing"

func TestEncode(t *testing.T) {
	expected := "encoded"

	received := Encode()
	if expected != received {
		t.Errorf("got %q, wanted %q", received, expected)
	}
}

func TestDecode(t *testing.T) {
	expected := "decoded"
	got := Decode()
	if expected != got {
		t.Errorf("got %q, wanted %q", got, expected)
	}
}
