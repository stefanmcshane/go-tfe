package tfe

import (
	"errors"
	"fmt"
	"time"
)

// OAuthClients handles communication with the oAuth client related methods
// of the Terraform Enterprise API.
//
// TFE API docs:
// https://www.terraform.io/docs/enterprise/api/oauth-clients.html
type OAuthClients struct {
	client *Client
}

// ServiceProviderType represents a VCS type.
type ServiceProviderType string

// List of available VCS types.
const (
	ServiceProviderBitbucket       ServiceProviderType = "bitbucket_hosted"
	ServiceProviderBitbucketServer ServiceProviderType = "bitbucket_server"
	ServiceProviderGithub          ServiceProviderType = "github"
	ServiceProviderGithubEE        ServiceProviderType = "github_enterprise"
	ServiceProviderGitlab          ServiceProviderType = "gitlab_hosted"
	ServiceProviderGitlabCE        ServiceProviderType = "gitlab_community_edition"
	ServiceProviderGitlabEE        ServiceProviderType = "gitlab_enterprise_edition"
)

// OAuthClient represents a connection between an organization and a VCS
// provider.
type OAuthClient struct {
	ID                  string               `jsonapi:"primary,oauth-clients"`
	APIURL              string               `jsonapi:"attr,api-url"`
	CallbackURL         string               `jsonapi:"attr,callback-url"`
	ConnectPath         string               `jsonapi:"attr,connect-path"`
	CreatedAt           time.Time            `jsonapi:"attr,created-at,iso8601"`
	HTTPURL             string               `jsonapi:"attr,http-url"`
	Key                 string               `jsonapi:"attr,key"`
	RSAPublicKey        string               `jsonapi:"attr,rsa-public-key"`
	ServiceProvider     *ServiceProviderType `jsonapi:"attr,service-provider"`
	ServiceProviderName string               `jsonapi:"attr,service-provider-display-name"`

	// Relations
	Organization *Organization `jsonapi:"relation,organization"`
	OAuthToken   *OAuthToken   `jsonapi:"relation,oauth-token"`
}

// CreateOAuthClientOptions represents the options for creating an oauth account.
type CreateOAuthClientOptions struct {
	// For internal use only!
	ID string `jsonapi:"primary,organizations"`

	// The base URL of your VCS provider's API.
	APIURL *string `jsonapi:"attr,api-url,omitempty"`

	// The homepage of your VCS provider.
	HTTPURL *string `jsonapi:"attr,http-url,omitempty"`

	// The token string you were given by your VCS provider.
	OAuthToken *string `jsonapi:"attr,oauth-token-string,omitempty"`

	// The VCS provider being connected with.
	ServiceProvider *ServiceProviderType `jsonapi:"attr,service-provider,omitempty"`
}

func (o CreateOAuthClientOptions) valid() error {
	if !validString(o.APIURL) {
		return errors.New("APIURL is required")
	}
	if !validString(o.HTTPURL) {
		return errors.New("HTTPURL is required")
	}
	if !validString(o.OAuthToken) {
		return errors.New("OAuthToken is required")
	}
	if o.ServiceProvider == nil {
		return errors.New("ServiceProvider is required")
	}
	return nil
}

// Create a VCS connection between an organization and a VCS provider.
func (s *OAuthClients) Create(organization string, options CreateOAuthClientOptions) (*OAuthClient, error) {
	if !validStringID(&organization) {
		return nil, errors.New("Invalid value for organization")
	}
	if err := options.valid(); err != nil {
		return nil, err
	}

	// Make sure we don't send a user provided ID.
	options.ID = ""

	u := fmt.Sprintf("organizations/%s/oauth-clients", organization)
	req, err := s.client.newRequest("POST", u, &options)
	if err != nil {
		return nil, err
	}

	o, err := s.client.do(req, &OAuthClient{})
	if err != nil {
		return nil, err
	}

	return o.(*OAuthClient), nil
}
