package stack

import "testing"

func TestPushPop(t *testing.T) {
	c := new(Stack)
	c.Push(5)
	if val, err := c.Pop(); err == nil && val != 5 {
		t.Log("Pop doesn't give 5")
		t.Fail()
	}

	if val, err := c.Pop(); err == nil && val != 5 {
		t.Log("Pop doesn't give 5")
		t.Fail()
	} else {
		t.Log("Stack is empty")
	}

}
