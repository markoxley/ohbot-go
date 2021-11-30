package ohbot

type Phrase struct {
	Set      string
	Variable string
	Text     string
}

func NewPhrase(set, variable, text string) *Phrase {
	return &Phrase{
		Set:      set,
		Variable: variable,
		Text:     text,
	}
}
