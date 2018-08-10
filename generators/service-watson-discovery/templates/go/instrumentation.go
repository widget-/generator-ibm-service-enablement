package services

import (
	"errors"
	"github.com/ibm-developer/ibm-cloud-env-golang"
	. "github.com/watson-developer-cloud/go-sdk/discoveryV1"
)

// InitializeServiceWatsonDiscovery uses IBMCloudEnv to find credentials 
// and initialize the Watson service
func InitializeServiceWatsonDiscovery() (*DiscoveryV1, error) {
	url, ok := IBMCloudEnv.GetString("watson_discovery_url")
	if !ok {
		return nil, errors.New("unable to find watson_discovery_url")
	}

	if apikey, ok := IBMCloudEnv.GetString("watson_discovery_apikey"); ok {
		return NewDiscoveryV1(&ServiceCredentials{
			ServiceURL: url,
			Version: "2018-03-05",
			APIkey: apikey,
		})
	}
	username, ok := IBMCloudEnv.GetString("watson_discovery_username")
	if !ok {
		return nil, errors.New("unable to find watson_discovery_username or watson_discovery_apikey")
	}
	password, ok := IBMCloudEnv.GetString("watson_discovery_password")
	if !ok {
		return nil, errors.New("unable to find watson_discovery_password")
	}
	return NewDiscoveryV1(&ServiceCredentials{
		ServiceURL: url,
		Version: "2018-03-05",
		Username: username,
		Password: password,
	}) 
}
