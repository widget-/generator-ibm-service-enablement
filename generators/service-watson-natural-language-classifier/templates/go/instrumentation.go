package services

import (
	"errors"
	"github.com/ibm-developer/ibm-cloud-env-golang"
	. "github.com/watson-developer-cloud/go-sdk/naturalLanguageClassifierV1"
)

// InitializeServiceWatsonNaturalLanguageClassifier uses IBMCloudEnv to find credentials 
// and initialize the Watson service
func InitializeServiceWatsonNaturalLanguageClassifier() (*NaturalLanguageClassifierV1, error) {
	url, ok := IBMCloudEnv.GetString("watson_natural_language_classifier_url")
	if !ok {
		return nil, errors.New("unable to find watson_natural_language_classifier_url")
	}

	if apikey, ok := IBMCloudEnv.GetString("watson_natural_language_classifier_apikey"); ok {
		return NewNaturalLanguageClassifierV1(&ServiceCredentials{
			ServiceURL: url,
			APIkey: apikey,
		})
	}
	username, ok := IBMCloudEnv.GetString("watson_natural_language_classifier_username")
	if !ok {
		return nil, errors.New("unable to find watson_natural_language_classifier_username or watson_natural_language_classifier_apikey")
	}
	password, ok := IBMCloudEnv.GetString("watson_natural_language_classifier_password")
	if !ok {
		return nil, errors.New("unable to find watson_natural_language_classifier_password")
	}
	return NewNaturalLanguageClassifierV1(&ServiceCredentials{
		ServiceURL: url,
		Username: username,
		Password: password,
	}) 
}

