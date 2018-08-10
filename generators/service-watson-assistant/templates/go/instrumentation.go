package services

import (
	"errors"
	"github.com/ibm-developer/ibm-cloud-env-golang"
	. "github.com/watson-developer-cloud/go-sdk/assistantV1"
)

// InitializeServiceWatsonAssistant uses IBMCloudEnv to find credentials 
// and initialize the Watson service
func InitializeServiceWatsonAssistant() (*AssistantV1, error) {
	url, ok := IBMCloudEnv.GetString("watson_assistant_url")
	if !ok {
		return nil, errors.New("unable to find watson_assistant_url")
	}

	if apikey, ok := IBMCloudEnv.GetString("watson_assistant_apikey"); ok {
		return NewAssistantV1(&ServiceCredentials{
			ServiceURL: url,
			Version: "2018-07-10",
			APIkey: apikey,
		})
	} 
	username, ok := IBMCloudEnv.GetString("watson_assistant_username")
	if !ok {
		return nil, errors.New("unable to find watson_assistant_username or watson_assistant_apikey")
	}
	password, ok := IBMCloudEnv.GetString("watson_assistant_password")
	if !ok {
		return nil, errors.New("unable to find watson_assistant_password")
	}
	return NewAssistantV1(&ServiceCredentials{
		ServiceURL: url,
		Version: "2018-07-10",
		Username: username,
		Password: password,
	}) 
}
