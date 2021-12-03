package ohbot

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	err := Init("")
	if err != nil {
		t.Fatalf("Fail: %s", err.Error())
	}
}

func TestTurn(t *testing.T) {
	err := Init("")
	if err != nil {
		t.Fatalf("Fail: %s", err.Error())
	}
	Move(HeadTurn, 0, 5)
	time.Sleep(time.Second)
	Move(HeadTurn, 10, 0)
	time.Sleep(time.Second)
	Move(HeadTurn, 0, 10)

	time.Sleep(time.Second)

	Detach(HeadTurn)

	Say("hello", nil)
}
