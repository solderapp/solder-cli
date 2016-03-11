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

//go:generate mockery -all -case=underscore

const (
	pathProfile         = "%s/api/profile"
	pathForge           = "%s/api/forge"
	pathForgeBuilds     = "%s/api/forge/%s/builds"
	pathForgeBuild      = "%s/api/forge/%s/builds/%s"
	pathMinecraft       = "%s/api/minecraft"
	pathMinecraftBuilds = "%s/api/minecraft/%s/builds"
	pathMinecraftBuild  = "%s/api/minecraft/%s/builds/%s"
	pathPacks           = "%s/api/packs"
	pathPack            = "%s/api/packs/%v"
	pathPackClients     = "%s/api/packs/%s/clients"
	pathPackClient      = "%s/api/packs/%s/clients/%s"
	pathBuilds          = "%s/api/builds"
	pathBuild           = "%s/api/builds/%v"
	pathBuildVersions   = "%s/api/builds/%s/versions"
	pathBuildVersion    = "%s/api/builds/%s/versions/%s"
	pathMods            = "%s/api/mods"
	pathMod             = "%s/api/mods/%v"
	pathVersions        = "%s/api/versions"
	pathVersion         = "%s/api/versions/%v"
	pathVersionBuilds   = "%s/api/versions/%s/builds"
	pathVersionBuild    = "%s/api/versions/%s/builds/%s"
	pathClients         = "%s/api/clients"
	pathClient          = "%s/api/clients/%v"
	pathClientPacks     = "%s/api/clients/%s/packs"
	pathClientPack      = "%s/api/clients/%s/packs/%s"
	pathKeys            = "%s/api/keys"
	pathKey             = "%s/api/keys/%v"
	pathUsers           = "%s/api/users"
	pathUser            = "%s/api/users/%v"
)

type defaultClient struct {
	client *http.Client
	base   string
}

// NewClient returns a client for the specified URL.
func NewClient(uri string) API {
	return &defaultClient{
		http.DefaultClient,
		uri,
	}
}

// NewClientToken returns a client that authenticates
// all outbound requests with the given token.
func NewClientToken(uri, token string) API {
	config := oauth2.Config{}

	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: token,
		},
	)

	return &defaultClient{
		auther,
		uri,
	}
}

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the Solder API.
func (c *defaultClient) SetClient(client *http.Client) {
	c.client = client
}

// ProfileGet returns a profile.
func (c *defaultClient) ProfileGet() (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfilePatch updates a profile.
func (c *defaultClient) ProfilePatch(in *Profile) (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.patch(uri, in, out)

	return out, err
}

// ForgeList returns a list of all Forge versions.
func (c *defaultClient) ForgeList() ([]*Forge, error) {
	var out []*Forge

	uri := fmt.Sprintf(pathForge, c.base)
	err := c.get(uri, &out)

	return out, err
}

// ForgeRefresh refreshs the available Minecraft versions.
func (c *defaultClient) ForgeRefresh() error {
	uri := fmt.Sprintf(pathForge, c.base)
	err := c.patch(uri, nil, nil)

	return err
}

// ForgeBuildList returns a list of related builds for a Forge version.
func (c *defaultClient) ForgeBuildList(id string) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathForgeBuilds, c.base, id)
	err := c.get(uri, &out)

	return out, err
}

// ForgeBuildAppend appends a Forge version to a build.
func (c *defaultClient) ForgeBuildAppend(id, append string) error {
	uri := fmt.Sprintf(pathForgeBuild, c.base, id, append)
	err := c.patch(uri, nil, nil)

	return err
}

// ForgeBuildAppend remove a Forge version from a build.
func (c *defaultClient) ForgeBuildDelete(id, delete string) error {
	uri := fmt.Sprintf(pathForgeBuild, c.base, id, delete)
	err := c.delete(uri)

	return err
}

// MinecraftList returns a list of all Minecraft versions.
func (c *defaultClient) MinecraftList() ([]*Minecraft, error) {
	var out []*Minecraft

	uri := fmt.Sprintf(pathMinecraft, c.base)
	err := c.get(uri, &out)

	return out, err
}

// MinecraftRefresh refreshs the available Minecraft versions.
func (c *defaultClient) MinecraftRefresh() error {
	uri := fmt.Sprintf(pathMinecraft, c.base)
	err := c.patch(uri, nil, nil)

	return err
}

// MinecraftBuildList returns a list of related builds for a Minecraft version.
func (c *defaultClient) MinecraftBuildList(id string) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathMinecraftBuilds, c.base, id)
	err := c.get(uri, &out)

	return out, err
}

// MinecraftBuildAppend appends a Minecraft version to a build.
func (c *defaultClient) MinecraftBuildAppend(id, append string) error {
	uri := fmt.Sprintf(pathMinecraftBuild, c.base, id, append)
	err := c.patch(uri, nil, nil)

	return err
}

// MinecraftBuildAppend remove a Minecraft version from a build.
func (c *defaultClient) MinecraftBuildDelete(id, delete string) error {
	uri := fmt.Sprintf(pathMinecraftBuild, c.base, id, delete)
	err := c.delete(uri)

	return err
}

// PackList returns a list of all packs.
func (c *defaultClient) PackList() ([]*Pack, error) {
	var out []*Pack

	uri := fmt.Sprintf(pathPacks, c.base)
	err := c.get(uri, &out)

	return out, err
}

// PackGet returns a pack.
func (c *defaultClient) PackGet(id string) (*Pack, error) {
	out := &Pack{}

	uri := fmt.Sprintf(pathPack, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// PackPost creates a pack.
func (c *defaultClient) PackPost(in *Pack) (*Pack, error) {
	out := &Pack{}

	uri := fmt.Sprintf(pathPacks, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// PackPatch updates a pack.
func (c *defaultClient) PackPatch(in *Pack) (*Pack, error) {
	out := &Pack{}

	uri := fmt.Sprintf(pathPack, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// PackDelete deletes a pack.
func (c *defaultClient) PackDelete(id string) error {
	uri := fmt.Sprintf(pathPack, c.base, id)
	err := c.delete(uri)

	return err
}

// PackClientList returns a list of related clients for a pack.
func (c *defaultClient) PackClientList(id string) ([]*Client, error) {
	var out []*Client

	uri := fmt.Sprintf(pathPackClients, c.base, id)
	err := c.get(uri, &out)

	return out, err
}

// PackClientAppend appends a client to a pack.
func (c *defaultClient) PackClientAppend(id, append string) error {
	uri := fmt.Sprintf(pathPackClient, c.base, id, append)
	err := c.patch(uri, nil, nil)

	return err
}

// PackClientDelete remove a client from a pack.
func (c *defaultClient) PackClientDelete(id, delete string) error {
	uri := fmt.Sprintf(pathPackClient, c.base, id, delete)
	err := c.delete(uri)

	return err
}

// BuildList returns a list of all builds for a specific pack.
func (c *defaultClient) BuildList(pack string) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathBuilds, c.base)

	if pack != "" {
		uri = fmt.Sprintf("%s?pack=%s", uri, pack)
	}

	err := c.get(uri, &out)

	return out, err
}

// BuildGet returns a build for a specific pack.
func (c *defaultClient) BuildGet(id string) (*Build, error) {
	out := &Build{}

	uri := fmt.Sprintf(pathBuild, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// BuildPost creates a build for a specific pack.
func (c *defaultClient) BuildPost(in *Build) (*Build, error) {
	out := &Build{}

	uri := fmt.Sprintf(pathBuilds, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// BuildPatch updates a build for a specific pack.
func (c *defaultClient) BuildPatch(in *Build) (*Build, error) {
	out := &Build{}

	uri := fmt.Sprintf(pathBuild, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// BuildDelete deletes a build for a specific pack.
func (c *defaultClient) BuildDelete(id string) error {
	uri := fmt.Sprintf(pathBuild, c.base, id)
	err := c.delete(uri)

	return err
}

// BuildVersionList returns a list of related versions for a build.
func (c *defaultClient) BuildVersionList(id string) ([]*Version, error) {
	var out []*Version

	uri := fmt.Sprintf(pathBuildVersions, c.base, id)
	err := c.get(uri, &out)

	return out, err
}

// BuildVersionAppend appends a version to a build.
func (c *defaultClient) BuildVersionAppend(id, append string) error {
	uri := fmt.Sprintf(pathBuildVersion, c.base, id, append)
	err := c.patch(uri, nil, nil)

	return err
}

// BuildVersionDelete remove a version from a build.
func (c *defaultClient) BuildVersionDelete(id, delete string) error {
	uri := fmt.Sprintf(pathBuildVersion, c.base, id, delete)
	err := c.delete(uri)

	return err
}

// ModList returns a list of all mods.
func (c *defaultClient) ModList() ([]*Mod, error) {
	var out []*Mod

	uri := fmt.Sprintf(pathMods, c.base)
	err := c.get(uri, &out)

	return out, err
}

// ModGet returns a mod.
func (c *defaultClient) ModGet(id string) (*Mod, error) {
	out := &Mod{}

	uri := fmt.Sprintf(pathMod, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// ModPost creates a mod.
func (c *defaultClient) ModPost(in *Mod) (*Mod, error) {
	out := &Mod{}

	uri := fmt.Sprintf(pathMods, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// ModPatch updates a mod.
func (c *defaultClient) ModPatch(in *Mod) (*Mod, error) {
	out := &Mod{}

	uri := fmt.Sprintf(pathMod, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// ModDelete deletes a mod.
func (c *defaultClient) ModDelete(id string) error {
	uri := fmt.Sprintf(pathMod, c.base, id)
	err := c.delete(uri)

	return err
}

// VersionList returns a list of all versions for a specific mod.
func (c *defaultClient) VersionList(mod string) ([]*Version, error) {
	var out []*Version

	uri := fmt.Sprintf(pathVersions, c.base)

	if mod != "" {
		uri = fmt.Sprintf("%s?mod=%s", uri, mod)
	}

	err := c.get(uri, &out)

	return out, err
}

// VersionGet returns a version for a specific mod.
func (c *defaultClient) VersionGet(id string) (*Version, error) {
	out := &Version{}

	uri := fmt.Sprintf(pathVersion, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// VersionPost creates a version for a specific mod.
func (c *defaultClient) VersionPost(in *Version) (*Version, error) {
	out := &Version{}

	uri := fmt.Sprintf(pathVersions, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// VersionPatch updates a version for a specific mod.
func (c *defaultClient) VersionPatch(in *Version) (*Version, error) {
	out := &Version{}

	uri := fmt.Sprintf(pathVersion, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// VersionDelete deletes a version for a specific mod.
func (c *defaultClient) VersionDelete(id string) error {
	uri := fmt.Sprintf(pathVersion, c.base, id)
	err := c.delete(uri)

	return err
}

// VersionBuildList returns a list of related builds for a version.
func (c *defaultClient) VersionBuildList(id string) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathVersionBuilds, c.base, id)
	err := c.get(uri, &out)

	return out, err
}

// VersionBuildAppend appends a build to a version.
func (c *defaultClient) VersionBuildAppend(id, append string) error {
	uri := fmt.Sprintf(pathVersionBuild, c.base, id, append)
	err := c.patch(uri, nil, nil)

	return err
}

// VersionBuildDelete remove a build from a version.
func (c *defaultClient) VersionBuildDelete(id, delete string) error {
	uri := fmt.Sprintf(pathVersionBuild, c.base, id, delete)
	err := c.delete(uri)

	return err
}

// ClientList returns a list of all clients.
func (c *defaultClient) ClientList() ([]*Client, error) {
	var out []*Client

	uri := fmt.Sprintf(pathClients, c.base)
	err := c.get(uri, &out)

	return out, err
}

// ClientGet returns a client.
func (c *defaultClient) ClientGet(id string) (*Client, error) {
	out := &Client{}

	uri := fmt.Sprintf(pathClient, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// ClientPost creates a client.
func (c *defaultClient) ClientPost(in *Client) (*Client, error) {
	out := &Client{}

	uri := fmt.Sprintf(pathClients, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// ClientPatch updates a client.
func (c *defaultClient) ClientPatch(in *Client) (*Client, error) {
	out := &Client{}

	uri := fmt.Sprintf(pathClient, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// ClientDelete deletes a client.
func (c *defaultClient) ClientDelete(id string) error {
	uri := fmt.Sprintf(pathClient, c.base, id)
	err := c.delete(uri)

	return err
}

// ClientPackList returns a list of related packs for a client.
func (c *defaultClient) ClientPackList(id string) ([]*Pack, error) {
	var out []*Pack

	uri := fmt.Sprintf(pathClientPacks, c.base, id)
	err := c.get(uri, &out)

	return out, err
}

// ClientPackAppend appends a pack to a client.
func (c *defaultClient) ClientPackAppend(id, append string) error {
	uri := fmt.Sprintf(pathClientPack, c.base, id, append)
	err := c.patch(uri, nil, nil)

	return err
}

// ClientPackDelete remove a pack from a client.
func (c *defaultClient) ClientPackDelete(id, delete string) error {
	uri := fmt.Sprintf(pathClientPack, c.base, id, delete)
	err := c.delete(uri)

	return err
}

// KeyList returns a list of all keys.
func (c *defaultClient) KeyList() ([]*Key, error) {
	var out []*Key

	uri := fmt.Sprintf(pathKeys, c.base)
	err := c.get(uri, &out)

	return out, err
}

// KeyGet returns a key.
func (c *defaultClient) KeyGet(id string) (*Key, error) {
	out := &Key{}

	uri := fmt.Sprintf(pathKey, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// KeyPost creates a key.
func (c *defaultClient) KeyPost(in *Key) (*Key, error) {
	out := &Key{}

	uri := fmt.Sprintf(pathKeys, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// KeyPatch updates a key.
func (c *defaultClient) KeyPatch(in *Key) (*Key, error) {
	out := &Key{}

	uri := fmt.Sprintf(pathKey, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// KeyDelete deletes a key.
func (c *defaultClient) KeyDelete(id string) error {
	uri := fmt.Sprintf(pathKey, c.base, id)
	err := c.delete(uri)

	return err
}

// UserList returns a list of all users.
func (c *defaultClient) UserList() ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.get(uri, &out)

	return out, err
}

// UserGet returns a user.
func (c *defaultClient) UserGet(id string) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// UserPost creates a user.
func (c *defaultClient) UserPost(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// UserPatch updates a user.
func (c *defaultClient) UserPatch(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// UserDelete deletes a user.
func (c *defaultClient) UserDelete(id string) error {
	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.delete(uri)

	return err
}

// Helper function for making an GET request.
func (c *defaultClient) get(rawurl string, out interface{}) error {
	return c.do(rawurl, "GET", nil, out)
}

// Helper function for making an POST request.
func (c *defaultClient) post(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "POST", in, out)
}

// Helper function for making an PUT request.
func (c *defaultClient) put(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PUT", in, out)
}

// Helper function for making an PATCH request.
func (c *defaultClient) patch(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PATCH", in, out)
}

// Helper function for making an DELETE request.
func (c *defaultClient) delete(rawurl string) error {
	return c.do(rawurl, "DELETE", nil, nil)
}

// Helper function to make an HTTP request
func (c *defaultClient) do(rawurl, method string, in, out interface{}) error {
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
func (c *defaultClient) stream(rawurl, method string, in, out interface{}) (io.ReadCloser, error) {
	uri, err := url.Parse(rawurl)

	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter

	if in != nil {
		buf = new(bytes.Buffer)
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

		msg := &Message{}
		parse := json.Unmarshal(out, msg)

		if parse != nil {
			return nil, fmt.Errorf(string(out))
		}

		return nil, fmt.Errorf(msg.Message)
	}

	return resp.Body, nil
}
