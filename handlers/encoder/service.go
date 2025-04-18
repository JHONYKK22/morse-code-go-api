package encoder

import (
	"encoding/json"
	"errors"

	"net/http"

	morseimpl "github.com/Jonathan-Guerra22/morse-code-api/handlers/morseImpl"
	"github.com/Jonathan-Guerra22/morse-code-api/handlers/dto"
	Encodertype "github.com/Jonathan-Guerra22/morse-code-api/handlers/encoderType"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var wx http.ResponseWriter

func EncodeBody(w http.ResponseWriter, r *http.Request) {
	wx = w
	
	fromBody(r, Encodertype.ENCODE)
}

func DecodeBody(w http.ResponseWriter, r *http.Request) {
	wx = w
	fromBody(r, Encodertype.DECODE)
}

func EncodePath(w http.ResponseWriter, r *http.Request) {
	wx = w
	fromPath(r, Encodertype.ENCODE)
}

func DecodePath(w http.ResponseWriter, r *http.Request) {
	wx = w
	fromPath(r, Encodertype.DECODE)
}

func fromBody(r *http.Request, encodeType Encodertype.EncodeType) {

	decode := json.NewDecoder(r.Body)
	var body dto.Request
	err := decode.Decode(&body)

	if err != nil {
		errorHandler("Invalid request body")
		return
	}

	err = validator.New().Struct(body)

	if err != nil {

		errorHandler("Invalid request body format")
		return
	}

	morseRunnerI, err := morseEncoderFactory(body.Text, encodeType)

	if err != nil {
		errorHandler(err.Error())
		return
	}

	value, err := morseRunnerI.Run()

	if err != nil {
		errorHandler(err.Error())
		return
	}

	responseHandler(value)

}

func fromPath(r *http.Request, encodeType Encodertype.EncodeType) {

	params := mux.Vars(r)

	text := params["text"]

	morseRunnerI, err := morseEncoderFactory(text, encodeType)

	if err != nil {
		errorHandler(err.Error())
		return
	}

	value, err := morseRunnerI.Run()

	if err != nil {
		errorHandler(err.Error())
		return
	}

	responseHandler(value)
}

func morseEncoderFactory(text string, encodeType Encodertype.EncodeType) (morseimpl.MorseRunner, error) {

	var morseRunnerI morseimpl.MorseRunner

	if encodeType == Encodertype.DECODE {
		morseRunnerI = morseimpl.Decoder{
			Text: text,
		}
	} else if encodeType == Encodertype.ENCODE {
		morseRunnerI = morseimpl.Encoder{
			Text: text,
		}
	} else {
		return nil, errors.New("internal error")
	}

	return morseRunnerI, nil
}

func errorHandler(error string) {

	errorMessage := dto.ErrorResponse{
		Message: error,
	}

	errJSON, _ := json.MarshalIndent(errorMessage, "", "    ")

	http.Error(wx, string(errJSON), 200)
}

func responseHandler(value string) {

	wx.WriteHeader(http.StatusOK)

	res := dto.Response{
		Data: value,
	}

	json.NewEncoder(wx).Encode(res)
}
