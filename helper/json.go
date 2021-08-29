package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{})  {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		log.Println(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{})  {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		log.Println(err)
	}
}