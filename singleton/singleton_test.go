package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {

	// cannot refer to unexported name singleton.singleton
	// instance := singleton.singleton{}

	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Errorf("Expect instance to equal, but not equal.\n")
	}

}
