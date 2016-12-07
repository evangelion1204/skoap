package oauth

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"github.com/zalando/go-tokens/client"
	"errors"
	"encoding/json"
)

type Client struct {
	authUrlBase string
	provider client.CredentialsProvider
}

func NewClient(authUrlBase string, provider client.CredentialsProvider) Client {
	return Client{authUrlBase: authUrlBase, provider: provider}
}

func (c *Client) GetAccessTokenByCode(code string) (string, error) {
	client := &http.Client{}

	params := url.Values{}

	credentials, _ := c.provider.Get()

	params.Set("client_id", credentials.Id())
	params.Set("realm", "/employees")
	params.Set("grant_type", "authorization_code")
	params.Set("redirect_uri", "https://router.local/callback")
	params.Set("client_secret", credentials.Secret())
	params.Set("code", code)

	request, _ := http.NewRequest(
		"POST",
		c.authUrlBase + "/oauth2/access_token" + "?" + params.Encode(),
		strings.NewReader(""),
	)

	response, err := client.Do(request)

	defer response.Body.Close()

	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", errors.New("An error occured: " + response.Status)
	}

	rawBytes, _ := ioutil.ReadAll(response.Body)

	var buffer map[string]interface{}
	if err := json.Unmarshal(rawBytes, &buffer); err != nil {
		return "", err
	}

	token, has := buffer["access_token"]
	if !has {
		return "", errors.New("Missing token in response")
	}

	accessToken, ok := token.(string)
	if !ok {
		return "", errors.New("Invalid token in response")
	}

	return accessToken, nil
}