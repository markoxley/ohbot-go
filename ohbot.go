package ohbot

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/huin/goserial"
)

type MotorName uint8

const (
	HeadNod MotorName = iota
	HeadTurn
	EyeTurn
	LidBlink
	TopLip
	BottomLip
	EyeTilt
	MouthOpen
)

// Version returns the version of the module
//  @return string
func Version() string {
	return version
}

// Init initialises Ohbot on the specified port.
// If an empty string is passed, the function will
// attempt to find the port that is attached to Ohbot
//  @param portName is the name of the port to be used. If unknown, pass an empty string
//  @return error
func Init(portName string) error {
	if err := loadMotorDefs(); err != nil {
		return err
	}
	silenceFile = dirName + "/Silence1.wav"
	// if err := PlaySoundFile(silenceFile); err != nil {
	// 	return err
	// }

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
			return fmt.Errorf("unable to find Ohbot on port %s", portName)
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
	log.Printf("Ohbot found on %v", port)
	text := "Hi"
	if strings.ToLower(synthesizer) == "festival" {
		generateSpeechFile(text)
	}

	loadSpeechDatabase()
	return nil
}

// CheckPort checks the specified port for the existance of Ohbot
//  @param p is the port
//  @return bool
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

// PlaySoundFile makes Ohbot play a sound file
//  @param fp the filepath of the sound file
//  @return error
func PlaySoundFile(fp string) error {
	cmd := exec.Command("aplay", fp)
	return cmd.Run()
}

// Move moves the specified servo
//  @param mn is the name of the servo
//  @param pos is the position to move the servo to 1 - 10
//  @param spd is the speed to move the servo 1 - 10
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
	absPos := m.absPos(pos)
	log.Printf("Absolute Angle: %v\n", absPos)
	spd = float64(250/10) * spd
	msg := fmt.Sprintf("m0%v,%v,%v\n", mn, absPos, spd)
	log.Printf("Msg: %v\n", msg)
	serWrite(msg)

	m.pos = pos
}

// Attach attaches the specified server
//  @param mn is the name of the servo
func Attach(mn MotorName) {
	motors[int(mn)].attach()
}

// Detach detaches the servo
//  @param mn is the name of the servo to detach
func Detach(mn MotorName) {
	motors[int(mn)].detach()
}

// SetLanguage sets the language for the voice
//  @param l is the language
func SetLanguage(l string) {
	if l == "" {
		return
	}
	voice = l
}

// SetSynthesizer sets a synthisizer for speech
//  @param s is the name of the synthesizer
func SetSynthesizer(s string) {
	if s == "" {
		return
	}
	synthesizer = s
}

// SetSpeechSpeed sets the speed of the speech
//  @param sr is the new speech rate
func SetSpeechSpeed(sr float64) {
	if sr <= 0 {
		return
	}
	speechRate = sr
}

// Say
//  @param text
//  @param sc
func Say(text string, sc *SpeechConfig) {
	log.Printf("Text: %s", text)
	text = strings.TrimSpace(text)
	if text == "" {
		return
	}
	text = strings.ReplaceAll(text, "picoh", "peek oh")
	text = strings.ReplaceAll(text, "Picoh", "peek oh")

	// if sc != nil {
	// 	sc = NewSpeechConfig()
	// }
	// soundDelay := sc.SoundDelay
	// if sc.HDMIAudo {
	// 	soundDelay--
	// }

	generateSpeechFile(text)

	// if strings.ToUpper(synthesizer) == "FESTIVAL" {
	// 	f, err := os.Open(phonemesFile)
	// 	if err != nil {
	// 		return
	// 	}
	// 	defer f.Close()
	// 	phonemes := make([]string, 0)
	// 	times := make([]float64, 0)
	// 	vals := make([]string, 0)

	// 	b := bufio.NewReader(f)
	// 	for {
	// 		line, err := b.ReadBytes('\n')
	// 		if err != nil {
	// 			return
	// 		}
	// 		vals = strings.Split(string(line),"")
	// 		if len(vals) >
	// 	}
	// }
}
