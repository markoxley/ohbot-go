package ohbot

import "time"

func Smile() {
	Move(LidBlink, 5)
	Move(BottomLip, 10)
	Move(TopLip, 0)
}

func Surprise() {
	Move(LidBlink, 10, 10)
	Move(BottomLip, 10, 10)
	Move(TopLip, 10, 10)
}

func Frown() {
	Move(LidBlink, 3)
	Move(BottomLip, 1)
	Move(TopLip, 8)
}

func Sleep() {
	Reset()
	time.Sleep(time.Millisecond * 500)
	Move(HeadNod, 6)
	time.Sleep(time.Millisecond * 200)
	Move(LidBlink, 0, 1)
	Move(EyeTilt, 10, 1)
	Move(HeadNod, 0, 1)
	time.Sleep(time.Millisecond * 2500)
	for i := 0; i < 3; i++ {
		Move(HeadNod, 4, 1)
		time.Sleep(time.Millisecond * 700)
		Move(HeadNod, 0, 1)
		time.Sleep(time.Second * 2)
	}
	Close()
}

func Wakeup() {
	Move(HeadNod, 5, 2)
	Move(HeadTurn, 5.2)
	time.Sleep(time.Second * 2)
	Move(LidBlink, 5, 3)
	Move(EyeTilt, 5, 5)
	time.Sleep(time.Millisecond * 500)
	for i := 0; i < 3; i++ {
		var p float64
		switch i {
		case 0:
			p = 1
		case 1:
			p = 9
		default:
			p = 5
		}
		Move(EyeTurn, p, 5)
		time.Sleep(time.Millisecond * 100)
		Move(HeadTurn, p, 5)
		time.Sleep(time.Millisecond * 500)
	}
	Smile()
	time.Sleep(time.Second * 3)
	Reset()
}
