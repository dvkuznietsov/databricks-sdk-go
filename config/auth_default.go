package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

type HeaderFactory interface {
	Token(ctx context.Context, cfg *Config) (*oauth2.Token, error)
}

var (
	authProviders = []CredentialsStrategy{
		PatCredentials{},
		BasicCredentials{},
		M2mCredentials{},
		DatabricksCliCredentials{},
		MetadataServiceCredentials{},

		// Attempt to configure auth from most specific to most generic (the Azure CLI).
		AzureMsiCredentials{},
		AzureClientSecretCredentials{},
		AzureCliCredentials{},

		// Attempt to configure auth from most specific to most generic (Google Application Default Credentials).
		GoogleCredentials{},
		GoogleDefaultCredentials{},
	}
)

type DefaultCredentials struct {
	name string
}

func (c *DefaultCredentials) Name() string {
	if c.name == "" {
		return "default"
	}
	return c.name
}

var authFlowUrl = "https://docs.databricks.com/en/dev-tools/auth.html#databricks-client-unified-authentication"
var errorMessage = fmt.Sprintf("cannot configure default credentials, please check %s to configure credentials for your preferred authentication method", authFlowUrl)

// ErrCannotConfigureAuth (experimental) is returned when no auth is configured
var ErrCannotConfigureAuth = errors.New(errorMessage)

func (c *DefaultCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	for _, p := range authProviders {
		if cfg.AuthType != "" && p.Name() != cfg.AuthType {
			// ignore other auth types if one is explicitly enforced
			logger.Infof(ctx, "Ignoring %s auth, because %s is preferred", p.Name(), cfg.AuthType)
			continue
		}
		logger.Tracef(ctx, "Attempting to configure auth: %s", p.Name())
		visitor, err := p.Configure(ctx, cfg)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p.Name(), err)
		}
		if visitor == nil {
			continue
		}
		c.name = p.Name()
		return visitor, nil
	}
	return nil, ErrCannotConfigureAuth
}

func (c *DefaultCredentials) Token(ctx context.Context, cfg *Config) (*oauth2.Token, error) {
	// Configure to select credentials provider
	c.Configure(ctx, cfg)
	for _, p := range authProviders {
		if c.name != p.Name() {
			continue
		}
		if provider, ok := p.(HeaderFactory); ok {
			return provider.Token(ctx, cfg)
		} else {
			return nil, fmt.Errorf("cannot get token for %s", p.Name())
		}
	}
	return nil, ErrCannotConfigureAuth
}
