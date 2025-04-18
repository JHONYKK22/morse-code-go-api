package morseImpl

import (
	"github.com/Jonathan-Guerra22/morse-code-mod/morse"
)

type MorseRunner interface {
	Run() (string, error)
}

type Encoder struct {
	Text string
}

type Decoder struct {
	Text string
}

func (d Decoder) Run() (string, error) {

	value, err := morse.Decode(d.Text)

	if err != nil {
		return "", err
	}

	return value, nil
}

func (e Encoder) Run() (string, error) {

	value, err := morse.Encode(e.Text)

	if err != nil {
		return "", err
	}

	return value, nil
}
