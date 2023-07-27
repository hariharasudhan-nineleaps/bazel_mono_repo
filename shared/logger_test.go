package shared

import "testing"

func TestNoPanicInLog(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Shouldn't panic.. %v", r)
		}
	}()
	Log("test")
}
