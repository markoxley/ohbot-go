package ohbot

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type phrase struct {
	set      string
	variable string
	text     string
}

type SpeechConfig struct {
	UntilDone  bool
	LipSync    bool
	HDMIAudo   bool
	SoundDelay float64
}

var (
	phraseList         []*phrase
	speechDatabaseFile string
)

func NewSpeechConfig() *SpeechConfig {
	return &SpeechConfig{
		UntilDone:  true,
		LipSync:    true,
		HDMIAudo:   false,
		SoundDelay: 0,
	}
}
func newPhrase(set, variable, text string) *phrase {
	return &phrase{
		set:      set,
		variable: variable,
		text:     text,
	}
}

func loadSpeechDatabase() {
	cf, err := os.Open(speechDatabaseFile)
	if err != nil {
		log.Fatalf("Unable to open speech database: %s", err.Error())
	}
	defer cf.Close()
	r := csv.NewReader(cf)
	for {
		r, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		phraseList = append(phraseList, newPhrase(r[0], r[1], r[2]))
	}
}

func generateSpeechFile(text string) error {
	// Note: This does not support gTTs
	log.Printf("Using %v for synth", synthesizer)
	re := regexp.MustCompile(`(?m)[^ .a-zA-Z0-9?\']`)
	st := re.ReplaceAllString(text, "")
	var bc string
	var args []string
	if strings.ToUpper(synthesizer) == "FESTIVAL" {
		bc = "festival"
		args = []string{
			"-b",
			fmt.Sprintf("'(set! mytext (Utterance Text \"%s\"))'", st),
			"'(utt.synth mytext)'",
			fmt.Sprintf("'(utt.save.wave mytext \"%s\")'", speechAudioFile),
			fmt.Sprintf("'(utt.save.segs mytext \"%s\")'", phonemesFile),
		}
		//bc = fmt.Sprintf("festival -b '(set! mytext (Utterance text \"%s\"))' '(utt.synth mytext)' '(utt.save.wave mytext \"%s\")' '(utt.save.segs mytext \"ohbotData/phonemes\")'", st, speechAudioFile)
	} else {
		bc = synthesizer
		args = []string{
			"-w",
			speechAudioFile,
			voice,
			st,
		}
		//		bc = fmt.Sprintf("%s -w %s %s \"%s\"", synthesizer, speechAudioFile, voice, st)
	}
	bc = bc + " " + strings.Join(args, " ")
	log.Println(bc)
	cmd := exec.Command(bc)
	if err := cmd.Run(); err != nil {
		log.Printf("Error: %s", err.Error())
		return err
	}
	return nil
}
