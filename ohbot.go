package ohbot

import (
	"errors"
	"log"
	"os"

	"github.com/beevik/etree"
)

const (
	Version         = "1.0.0"
	dirName         = "ohbotData"
	speechAudioFile = "ohbotData/ohbotspeech.wav"

	soundFolder  = "ohbotData/Sounds"
	settingsFile = "ohbotData/OhbotSettings.xml"
)

// Constants for motors
const (
	HeadNod uint8 = iota
	HeadTurn
	EyeTurn
	LidBlink
	TopLip
	BottomLip
	EyeTilt
)

var (
	sensors            []float64
	motors             []*Motor
	shapeList          []float64
	phraseList         []float64
	port               string
	writing            bool
	connected          bool
	topLipFree         bool
	speechDatabaseFile string
	ohbotMotorDefFile  string
	defaultEyeShape    string
	eyeShapeFile       string
	synthesizer        string
	voice              string
	language           string
	speechRate         float64
	lastfex, lastfey   float64
)

func init() {
	language = "en-GB"
	ohbotMotorDefFile = "ohbotData/MotorDefinitionsv21.omd"
	sensors = []float64{0, 0, 0, 0, 0, 0, 0, 0}
	for i := uint8(0); i <= EyeTilt; i++ {
		motors = append(motors, NewMotor())
	}
	writing = false
	connected = false
	topLipFree = false

	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Unable to create directory. %s", err.Error())
	}

	if err = testFile(settingsFile, xmlDefault); err != nil {
		log.Fatalf("Unable to create default XML file. %s", err.Error())
	}

	if err = loadSettings(); err != nil {
		log.Fatalf("Unable to load settings. %s", err.Error())
	}

	if err = testFile(speechDatabaseFile, speechDef); err != nil {
		log.Fatalf("Unable to create speech database file. %s", err.Error())
	}

	if err = testFile(eyeShapeFile, eyeDef); err != nil {
		log.Fatalf("Unable to create eye shape file. %s", err.Error())
	}

	if err = testFile(ohbotMotorDefFile, motorDef); err != nil {
		log.Fatalf("Unable to create ohbot motor definition file. %s", err.Error())
	}

	// Maybe add sound option here

	if synthesizer == "" {
		synthesizer = "festival"
	}

	speechRate = 170

	lastfex = 5
	lastfey = 5
}

func loadSettings() error {
	tree := etree.NewDocument()
	if err := tree.ReadFromFile(settingsFile); err != nil {
		return errors.New("unable to read settings file.")
	}

	root := tree.SelectElement("SettingList")
	for _, element := range root.SelectElements("Setting") {
		value := element.SelectAttrValue("Value", "")
		switch element.SelectAttrValue("Name", "") {
		case "DefaultEyeShape":
			defaultEyeShape = value
		case "DefaultSpeechSynth":
			synthesizer = value
		case "DefaultVoice":
			voice = value
		case "DefaultLang":
			language = value
		case "SpeechDBFile":
			speechDatabaseFile = value
		case "EyeShapeList":
			eyeShapeFile = value
		case "MotorDefFile":
			ohbotMotorDefFile = value
		}
	}
	return nil
}

func testFile(fp, cnt string) error {
	_, err := os.Stat(fp)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil {
		return nil
	}

	file, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(cnt)
	return nil

}
