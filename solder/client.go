package solder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

//go:generate mockery -all

const (
	pathForge     = "%s/api/forge"
	pathMinecraft = "%s/api/minecraft"
	pathProfile   = "%s/api/profile"
	pathUsers     = "%s/api/users"
	pathUser      = "%s/api/users/%s"
	pathClients   = "%s/api/clients"
	pathClient    = "%s/api/clients/%s"
	pathKeys      = "%s/api/keys"
	pathKey       = "%s/api/keys/%s"
	pathPacks     = "%s/api/packs"
	pathPack      = "%s/api/packs/%s"
	pathBuilds    = "%s/api/builds"
	pathBuild     = "%s/api/builds/%s"
	pathMods      = "%s/api/mods"
	pathMod       = "%s/api/mods/%s"
	pathVersions  = "%s/api/mods/%s/versions"
	pathVersion   = "%s/api/mods/%s/versions/%s"
)

type client struct {
	client *http.Client
	base   string
}

// NewClient returns a client for the specified URL.
func NewClient(uri string) Client {
	return &client{
		http.DefaultClient,
		uri,
	}
}

// NewClientToken returns a client that authenticates
// all outbound requests with the given token.
func NewClientToken(uri, token string) Client {
	config := oauth2.Config{}

	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: token,
		},
	)

	return &client{
		auther,
		uri,
	}
}

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the Solder API.
func (c *client) SetClient(client *http.Client) {
	c.client = client
}

// Helper function for making an GET request.
func (c *client) get(rawurl string, out interface{}) error {
	return c.do(rawurl, "GET", nil, out)
}

// Helper function for making an POST request.
func (c *client) post(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "POST", in, out)
}

// Helper function for making an PUT request.
func (c *client) put(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PUT", in, out)
}

// Helper function for making an PATCH request.
func (c *client) patch(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PATCH", in, out)
}

// Helper function for making an DELETE request.
func (c *client) delete(rawurl string) error {
	return c.do(rawurl, "DELETE", nil, nil)
}

// Helper function to make an HTTP request
func (c *client) do(rawurl, method string, in, out interface{}) error {
	body, err := c.stream(
		rawurl,
		method,
		in,
		out,
	)

	if err != nil {
		return err
	}

	defer body.Close()

	if out != nil {
		return json.NewDecoder(body).Decode(out)
	}

	return nil
}

// Helper function to stream an HTTP request
func (c *client) stream(rawurl, method string, in, out interface{}) (io.ReadCloser, error) {
	uri, err := url.Parse(rawurl)

	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter

	if in != nil {
		buf := bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(in)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri.String(), buf)

	if err != nil {
		return nil, err
	}

	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode > http.StatusPartialContent {
		defer resp.Body.Close()
		out, _ := ioutil.ReadAll(resp.Body)

		return nil, fmt.Errorf(string(out))
	}

	return resp.Body, nil
}
