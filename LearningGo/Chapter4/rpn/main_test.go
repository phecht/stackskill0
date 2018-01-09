package main

import "testing"
import "github.com/stackskill0/LearningGo/Chapter4/stack"

// This is not a great test.  But it is a test
func TestMain(t *testing.T) {
	st := new(stack.Stack)
	st.Push(50)
	st.Push(10)

	if val, err := executeIt(*st, "+"); err == nil && val != 60 {
		t.Log("Doesn't add up")
		t.Log(val)
		t.Fail()
	}
	if val, err := executeIt(*st, "-"); err == nil && val != 40 {
		t.Log("Doesn't minus up")
		t.Log(val)
		t.Fail()
	}
	if val, err := executeIt(*st, "*"); err == nil && val != 500 {
		t.Log("Doesn't mult up")
		t.Log(val)

		t.Fail()
	}
	if val, err := executeIt(*st, "/"); err == nil && val != 5 {
		t.Log("Doesn't / up")
		t.Log(val)
		t.Fail()
	}
}
