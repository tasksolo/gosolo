package gosolo

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/url"
)

type Config struct {
	BaseURL string `json:"baseURL" toml:"baseURL"`
	Token   string `json:"token" toml:"token"`
	Shard   string `json:"shard" toml:"shard"`

	Debug    bool `json:"debug" toml:"debug"`
	Insecure bool `json:"insecure" toml:"insecure"`
}

type GetUserPassFunc func() (string, string, error)

func NewClient(ctx context.Context, cfg *Config, getCreds GetUserPassFunc) (*Client, error) {
	var err error

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.solotask.io/"
	}

	c := NewClientDirect(cfg.BaseURL).
		SetDebug(cfg.Debug)

	if cfg.Insecure {
		c.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //nolint:gosec
	}

	if cfg.Token == "" {
		user, pass, err := getCreds()
		if err != nil {
			return nil, err
		}

		c.SetBasicAuth(user, pass)
	} else {
		c.SetAuthToken(cfg.Token)
	}

	if cfg.Shard == "" {
		user, err := c.GetUser(ctx, "me", nil)
		if err != nil {
			return nil, err
		}

		cfg.Shard = user.Shard
	}

	shardURL, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return nil, err
	}

	shardURL.Host = fmt.Sprintf("%s.%s", cfg.Shard, shardURL.Host)

	c.SetBaseURL(shardURL.String())

	if cfg.Token == "" {
		token, err := c.CreateToken(ctx, &Token{})
		if err != nil {
			return nil, err
		}

		cfg.Token = token.Token

		c.SetAuthToken(cfg.Token)
	}

	return c, err
}

func (c *Client) Auth(ctx context.Context, user, pass string) (string, error) {
	c.SetBasicAuth(user, pass)

	token, err := c.CreateToken(ctx, &Token{})
	if err != nil {
		return "", err
	}

	c.SetAuthToken(token.Token)

	return token.Token, nil
}
