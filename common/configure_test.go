package common

import "testing"

func TestConfig(t *testing.T) {
	c := DefaultConfig()
	t.Logf("config : %v", c.DumpJson())
}
