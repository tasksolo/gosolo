package gosolo

import (
	"context"
)

func (c *Client) Auth(ctx context.Context, user, pass string) (string, error) {
	c.SetBasicAuth(user, pass)

	token, err := c.CreateToken(ctx, &TokenRequest{})
	if err != nil {
		return "", err
	}

	c.SetAuthToken(*token.Token)

	return *token.Token, nil
}
