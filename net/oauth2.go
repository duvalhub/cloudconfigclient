package net

import (
	"github.com/Piszmog/cfservices"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// CreateOAuth2Client creates an OAuth2 http.Client from the provided credentials.
func CreateOAuth2Client(cred *cfservices.Credentials) (*http.Client, error) {
	config, err := CreateOAuth2Config(cred)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create oauth2 config")
	}
	return config.Client(context.Background()), nil
}

// CreateOAuth2Config creates an OAuth2 config from the provided credentials.
func CreateOAuth2Config(cred *cfservices.Credentials) (*clientcredentials.Config, error) {
	if cred == nil {
		return nil, errors.New("cannot create oauth2 config when credentials are nil")
	}
	return &clientcredentials.Config{
		ClientID:     cred.ClientId,
		ClientSecret: cred.ClientSecret,
		TokenURL:     cred.AccessTokenUri,
	}, nil
}