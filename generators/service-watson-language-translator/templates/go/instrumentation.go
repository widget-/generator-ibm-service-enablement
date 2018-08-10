package services

import (
	"errors"
	"github.com/ibm-developer/ibm-cloud-env-golang"
	. "github.com/watson-developer-cloud/go-sdk/languageTranslatorV3"
)

// InitializeServiceWatsonLanguageTranslator uses IBMCloudEnv to find credentials 
// and initialize the Watson service
func InitializeServiceWatsonLanguageTranslator() (*LanguageTranslatorV3, error) {
	url, ok := IBMCloudEnv.GetString("watson_language_translator_url")
	if !ok {
		return nil, errors.New("unable to find watson_language_translator_url")
	}

	if apikey, ok := IBMCloudEnv.GetString("watson_language_translator_apikey"); ok {
		return NewLanguageTranslatorV3(&ServiceCredentials{
			ServiceURL: url,
			Version: "2018-05-01",
			APIkey: apikey,
		})
	}
	username, ok := IBMCloudEnv.GetString("watson_language_translator_username")
	if !ok {
		return nil, errors.New("unable to find watson_language_translator_username or watson_language_translator_apikey")
	}
	password, ok := IBMCloudEnv.GetString("watson_language_translator_password")
	if !ok {
		return nil, errors.New("unable to find watson_language_translator_password")
	}
	return NewLanguageTranslatorV3(&ServiceCredentials{
		ServiceURL: url,
		Version: "2018-05-01",
		Username: username,
		Password: password,
	}) 
}
