package services

import (
	"errors"
	"github.com/ibm-developer/ibm-cloud-env-golang"
	. "github.com/watson-developer-cloud/go-sdk/visualRecognitionV3"
)

// InitializeServiceWatsonVisualRecognition uses IBMCloudEnv to find credentials 
// and initialize the Watson service
func InitializeServiceWatsonVisualRecognition() (*VisualRecognitionV3, error) {
	url, ok := IBMCloudEnv.GetString("watson_visual_recognition_url")
	if !ok {
		return nil, errors.New("unable to find watson_visual_recognition_url")
	}

	if apikey, ok := IBMCloudEnv.GetString("watson_visual_recognition_apikey"); ok {
		return NewVisualRecognitionV3(&ServiceCredentials{
			ServiceURL: url,
			Version: "2018-03-19",
			APIkey: apikey,
		})
	}

	if api_key, ok := IBMCloudEnv.GetString("watson_visual_recognition_api_key"); ok {
		return NewVisualRecognitionV3(&ServiceCredentials{
			ServiceURL: url,
			Version: "2018-03-19",
			APIkey: api_key,
		})
	}

	username, ok := IBMCloudEnv.GetString("watson_visual_recognition_username")
	if !ok {
		return nil, errors.New("unable to find watson_visual_recognition_username or watson_visual_recognition_api_key")
	}
	password, ok := IBMCloudEnv.GetString("watson_visual_recognition_username")
	if !ok {
		return nil, errors.New("unable to find watson_visual_recognition_username")
	}
	return NewVisualRecognitionV3(&ServiceCredentials{
		ServiceURL: url,
		Version: "2018-03-19",
		Username: username,
		Password: password,
	}) 
}
