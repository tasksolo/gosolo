package gosolo

import (
	"context"
	"fmt"
	"net/url"
)

type Config struct {
	BaseURL *url.URL
	Token   string
	Shard   string
}

type GetUserPassFunc func() (string, string, error)

func NewClient(ctx context.Context, cfg *Config, getCreds GetUserPassFunc) (*Client, error) {
	var err error

	if cfg.BaseURL == nil {
		cfg.BaseURL, err = url.Parse("https://api.solotask.io/")
		if err != nil {
			return nil, err
		}
	}

	c := NewClientDirect(cfg.BaseURL.String()).SetDebug(true)

	if cfg.Token == "" {
		user, pass, err := getCreds()
		if err != nil {
			return nil, err
		}

		fmt.Printf("USER=%s PASS=%s\n", user, pass)
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

	shardURL := *cfg.BaseURL
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
