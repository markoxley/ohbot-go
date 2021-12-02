package ohbot

import (
	"errors"
	"fmt"
	"github.com/huin/goserial"
	"os/exec"
	"strings"
)

type MotorName uint8

// Constants for motors
const (
	HeadNod MotorName = iota
	HeadTurn
	EyeTurn
	LidBlink
	TopLip
	BottomLip
	EyeTilt
)

func Version() string {
	return version
}

// Init initialises Ohbot on the specified port.
// If an empty string is passed, the function will
// attempt to find the port that is attached to Ohbot
func Init(portName string) error {
	if err := loadMotorDefs(); err != nil {
		return err
	}
	silenceFile = dirName + "/Silence1.wav"
	if err := PlaySoundFile(silenceFile); err != nil {
		return err
	}

	ports, err := listSerialPorts()
	if err != nil {
		return err
	}

	if portName == "" {
		for _, p := range ports {
			if CheckPort(p) {
				port = p
				connected = true
				break
			}
		}
		if port == "" {
			return errors.New("unable to find Ohbot port")
		}
	} else {
		if !CheckPort(portName) {
			return errors.New(fmt.Sprintf("unable to find Ohbot on port %s", portName))
		}
		connected = true
		port = portName
	}

	c := &goserial.Config{
		Name: port,
		Baud: 19200,
	}
	ser, err = goserial.OpenPort(c)
	if err != nil {
		return err
	}

	text := "Hi"
	if strings.ToLower(synthesizer) == "festival" {
		generateSpeechFile(text)
	}

	loadSpeechDatabase()
	return nil
}

func CheckPort(p string) bool {
	c := &goserial.Config{
		Name: p,
		Baud: 19200,
	}
	ser, err := goserial.OpenPort(c)
	if err != nil {
		return false
	}
	defer ser.Close()

	msg := "v\n"
	_, err = ser.Write([]byte(msg))
	if err != nil {
		return false
	}
	buf := make([]byte, 256)
	n, err := ser.Read(buf)
	if err != nil {
		return false
	}
	line := string(buf[:n])
	return strings.Contains(line, "v1")
}

func PlaySoundFile(fp string) error {
	cmd := exec.Command("aplay", fp)
	return cmd.Run()
}

func Move(mn MotorName, pos float64, spd float64) {
	pos = limit(pos)
	spd = limit(spd)

	if pos > 9 && mn == BottomLip {
		topLipFree = true
	}

	if pos <= 5 && mn == BottomLip {
		topLipFree = false
	}

	if pos < 5 && mn == BottomLip {
		pos = 5 - ((5 - pos) / 2)
	}

	m := motors[int(mn)]
	if m.rev {
		pos = 10 - pos
	}

	m.attach()
	absPos := m.absPos(m.pos)
	spd = (250 / 10) * spd
	msg := fmt.Sprintf("m0%v,%v,%v\n", mn, absPos, spd)
	serWrite(msg)

	m.pos = pos
}

func Attach(mn MotorName) {
	motors[int(mn)].attach()
}

func Detach(mn MotorName) {
	motors[int(mn)].detach()
}

func SetLanguage(l string) {
	if l == "" {
		return
	}
	voice = l
}

func SetSynthesizer(s string) {
	if s == "" {
		return
	}
	synthesizer = s
}

func SetSpeechSpeed(sr float64) {
	if sr <= 0 {
		return
	}
	speechRate = sr
}

func SayWithConfig(sc *SpeechConfig) {
	text := strings.TrimSpace(sc.Text)
	if text == "" {
		return
	}
	text = strings.ReplaceAll(text, "picoh", "peek oh")
	text = strings.ReplaceAll(text, "Picoh", "peek oh")

	soundDelay := sc.SoundDelay
	if sc.HDMIAudo {
		soundDelay--
	}

	generateSpeechFile(text)
}

func Say(t string) {
	SayWithConfig(&SpeechConfig{
		Text:       t,
		UntilDone:  true,
		LipSync:    true,
		HDMIAudo:   false,
		SoundDelay: 0,
	})
}
