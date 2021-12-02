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
	Text       string
	UntilDone  bool
	LipSync    bool
	HDMIAudo   bool
	SoundDelay float64
}

var (
	phraseList         []*phrase
	speechDatabaseFile string
)

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
	re := regexp.MustCompile(`(?m)[^ .a-zA-Z0-9?\']`)
	st := re.ReplaceAllString(text, "")
	var bc string
	if strings.ToUpper(synthesizer) == "FESTIVAL" {
		bc = fmt.Sprintf("festival -b '(set! mytext (Utterance text \"%s\"))' '(utt.synth mytext)' '(utt.save.wave mytext \"%s\")' '(utt.save.segs mytext \"ohbotData/phonemes\")'", st, speechAudioFile)
	} else {
		bc = fmt.Sprintf("%s -w %s %s \"%s\"", synthesizer, speechAudioFile, voice, st)
	}
	cmd := exec.Command(bc)
	return cmd.Run()
}
