package ohbot

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/beevik/etree"
)

const (
	version = "1.0.0"
	dirName = "ohbotData"
)

var (
	sensors          []float64
	shapeList        []float64
	port             string
	writing          bool
	connected        bool
	topLipFree       bool
	silenceFile      string
	synthesizer      string
	voice            string
	language         string
	speechRate       float64
	lastfex, lastfey float64
	ser              io.ReadWriteCloser
	workingDir       string
	speechAudioPath  string
	settingsPath     string
	phonemesPath     string
	pathSep          string
	speechAudioFile  string
	soundFolder      string
	settingsFile     string
	phonemesFile     string
)

func init() {
	language = "en-GB"
	ohbotMotorDefFile = "ohbotData/MotorDefinitionsv21.omd"
	sensors = []float64{0, 0, 0, 0, 0, 0, 0, 0}
	for i := uint8(0); i <= uint8(MouthOpen); i++ {
		motors = append(motors, newMotor())
	}
	writing = false
	connected = false
	topLipFree = false
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to determine working directory: %s", err.Error())
	}
	pathSep = string(os.PathSeparator)
	if !strings.HasSuffix(workingDir, pathSep) {
		workingDir += pathSep
	}
	//workingDir += dirName
	err = os.Mkdir(workingDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Unable to create directory. %s", err.Error())
	}

	speechAudioFile = workingDir + "ohbotspeech.wav"
	soundFolder = workingDir + "Sounds"
	settingsFile = workingDir + "OhbotSettings.xml"
	phonemesFile = workingDir + "phonemes"
	//settingsFile = workingDir + settingsFile
	if err = testFile(settingsFile, settingsDef); err != nil {
		log.Fatalf("Unable to create default XML file. %s", err.Error())
	}

	if err = loadSettings(); err != nil {
		log.Fatalf("Unable to load settings. %s", err.Error())
	}

	if err = testFile(speechDatabaseFile, speechDef); err != nil {
		log.Fatalf("Unable to create speech database file. %s", err.Error())
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

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil
}

func loadSettings() error {
	tree := etree.NewDocument()
	if err := tree.ReadFromFile(settingsFile); err != nil {
		return fmt.Errorf("unable to read settings file: %s", err.Error())
	}

	root := tree.SelectElement("SettingList")
	for _, element := range root.SelectElements("Setting") {
		value := element.SelectAttrValue("Value", "")
		switch element.SelectAttrValue("Name", "") {
		case "DefaultSpeechSynth":
			synthesizer = value
		case "DefaultVoice":
			voice = value
		case "DefaultLang":
			language = value
		case "SpeechDBFile":
			speechDatabaseFile = workingDir + value
		case "MotorDefFile":
			ohbotMotorDefFile = workingDir + value
		}
	}
	return nil
}

func listSerialPorts() ([]string, error) {

	dir, err := os.ReadDir("/dev")
	if err != nil {
		return nil, err
	}
	res := make([]string, 0, len(dir))
	for _, d := range dir {
		if d.IsDir() {
			continue
		}
		if len(d.Name()) < 6 {
			continue
		}
		if d.Name()[:4] == "ttyA" {
			res = append(res, "/dev/"+d.Name())
		}
	}
	return res, nil
}

func limit(v float64) float64 {
	if v > 10 {
		return 10
	}
	if v < 0 {
		return 0
	}
	return v
}

func serWrite(s string) {
	if !connected {
		return
	}
	writing = true
	ser.Write([]byte(s))
	writing = false

}
