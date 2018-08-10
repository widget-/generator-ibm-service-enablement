package services

import (
	"errors"
	"github.com/ibm-developer/ibm-cloud-env-golang"
	. "github.com/watson-developer-cloud/go-sdk/textToSpeechV1"
)

// InitializeServiceWatsonTextToSpeech uses IBMCloudEnv to find credentials 
// and initialize the Watson service
func InitializeServiceWatsonTextToSpeech() (*TextToSpeechV1, error) {
	url, ok := IBMCloudEnv.GetString("watson_text_to_speech_url")
	if !ok {
		return nil, errors.New("unable to find watson_text_to_speech_url")
	}

	if apikey, ok := IBMCloudEnv.GetString("watson_text_to_speech_apikey"); ok {
		return NewTextToSpeechV1(&ServiceCredentials{
			ServiceURL: url,
			APIkey: apikey,
		})
	}
	username, ok := IBMCloudEnv.GetString("watson_text_to_speech_username")
	if !ok {
		return nil, errors.New("unable to find watson_text_to_speech_username or watson_text_to_speech_apikey")
	}
	password, ok := IBMCloudEnv.GetString("watson_text_to_speech_password")
	if !ok {
		return nil, errors.New("unable to find watson_text_to_speech_password")
	}
	return NewTextToSpeechV1(&ServiceCredentials{
		ServiceURL: url,
		Username: username,
		Password: password,
	}) 
}
