package services

import (
	"errors"
	"github.com/ibm-developer/ibm-cloud-env-golang"
	. "github.com/watson-developer-cloud/go-sdk/visualRecognitionV3"
)

// InitializeServiceWatsonVisualRecognition uses IBMCloudEnv to find credentials 
// and initialize the Watson service
func InitializeServiceWatsonVisualRecognition() (*visualRecognitionV3.VisualRecognitionV3, error) {
	url, ok := IBMCloudEnv.GetString("watson_visual_recognition_url")
	if !ok {
		return nil, errors.New("unable to find watson_visual_recognition_url")
	}

	if apikey, ok := IBMCloudEnv.GetString("watson_visual_recognition_api_key"); ok {
		return visualRecognitionV3.NewVisualRecognitionV3(watson.Credentials{
			ServiceURL: url,
			Version: "2018-03-19",
			APIkey: apikey,
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
	return visualRecognitionV3.NewVisualRecognitionV3(watson.Credentials{
		ServiceURL: url,
		Version: "2018-03-19",
		Username: username,
		Password: password,
	}) 
}

// TODO: THIS ONE IS DIFFERENT. the iam api key vs apikey. wait for the go-sdk to be completed
