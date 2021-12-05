package ohbot

import (
	"log"
	"os"
	"testing"
	"time"
)

func motorTest(mn ...MotorName) {

	for i := 0; i < 11; i += 5 {
		for p := 0; p < 2; p++ {
			for _, m := range mn {
				Move(m, float64(p*10), float64(i))
			}
			time.Sleep(time.Second)
		}
	}
}

func endTest() {
	time.Sleep(time.Second * 3)
	Reset()
}

func startTest() {
	Init("")
	Reset()
}
func TestInit(t *testing.T) {
	err := Init("")
	if err != nil {
		t.Fatalf("Fail: %s", err.Error())
	}
}

func TestReset(t *testing.T) {
	err := Init("")
	if err != nil {
		t.Fatalf("Fail: %s", err.Error())
	}
	Reset()
}

// func TestClose(t *testing.T) {
// 	Init("")
// 	Close()
// }
// func TestEyeTurn(t *testing.T) {
// 	startTest()
// 	motorTest(EyeTurn)
// 	endTest()
// }
// func TestEyeTilt(t *testing.T) {
// 	startTest()
// 	motorTest(EyeTilt)
// 	endTest()
// }

// func TestLidBlink(t *testing.T) {
// 	startTest()
// 	motorTest(LidBlink)
// 	endTest()
// }

// func TestHeadTurn(t *testing.T) {
// 	startTest()
// 	motorTest(HeadTurn)
// 	endTest()
// }

// func TestHeadNod(t *testing.T) {
// 	startTest()
// 	motorTest(HeadNod)
// 	endTest()
// }

func TestSmile(t *testing.T) {
	startTest()
	Smile()
	endTest()
}

// func TestSurprise(t *testing.T) {
// 	startTest()
// 	Surprise()
// 	endTest()
// }

// func TestFrown(t *testing.T) {
// 	startTest()
// 	Frown()
// 	endTest()
// }

// func TestSleep(t *testing.T) {
// 	startTest()
// 	Sleep()
// 	endTest()
// }

func TestExpressions(t *testing.T) {
	startTest()
	Smile()
	time.Sleep(time.Second)
	Frown()
	time.Sleep(time.Second)
	Surprise()
	time.Sleep(time.Second)
	Sleep()
	endTest()
}

func TestSpeech(t *testing.T) {
	startTest()
	s := os.Getenv("PATH")
	log.Printf("PATH = %s", s)
	log.Println("TEST")
	Say("Hello, my name is Mark and I am a programmer", nil)
	endTest()
}

func TestWakeup(t *testing.T) {
	startTest()
	Sleep()
	time.Sleep(time.Second)
	Wakeup()
	endTest()
}

func TestSpeak(t *testing.T) {
	startTest()
	Wait(1)
	Say("Hello Daisy May!", nil)
	endTest()
}

func TestLips(t *testing.T) {
	startTest()
	Move(TopLip, 0, 10)
	Move(BottomLip, 0, 10)
	Wait(1)
	endTest()

}
