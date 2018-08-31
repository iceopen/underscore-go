package underscorego

import "testing"

func TestGetVersion(t *testing.T) {
	var u U
	t.Log(u.GetVersion())
}