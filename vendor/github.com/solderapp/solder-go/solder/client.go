package solder

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jackspirou/syscerts"
	"golang.org/x/oauth2"
)

const (
	pathProfile        = "%s/api/profile/self"
	pathProfileToken   = "%s/api/profile/token"
	pathForge          = "%s/api/forge"
	pathForgeBuild     = "%s/api/forge/%v/builds"
	pathMinecraft      = "%s/api/minecraft"
	pathMinecraftBuild = "%s/api/minecraft/%v/builds"
	pathPacks          = "%s/api/packs"
	pathPack           = "%s/api/packs/%v"
	pathPackClient     = "%s/api/packs/%v/clients"
	pathPackUser       = "%s/api/packs/%v/users"
	pathBuilds         = "%s/api/packs/%v/builds"
	pathBuild          = "%s/api/packs/%v/builds/%v"
	pathBuildVersion   = "%s/api/packs/%v/builds/%v/versions"
	pathMods           = "%s/api/mods"
	pathMod            = "%s/api/mods/%v"
	pathModUser        = "%s/api/mods/%v/users"
	pathVersions       = "%s/api/mods/%v/versions"
	pathVersion        = "%s/api/mods/%v/versions/%v"
	pathVersionBuild   = "%s/api/mods/%v/versions/%v/builds"
	pathClients        = "%s/api/clients"
	pathClient         = "%s/api/clients/%v"
	pathClientPack     = "%s/api/clients/%v/packs"
	pathUsers          = "%s/api/users"
	pathUser           = "%s/api/users/%v"
	pathUserMod        = "%s/api/users/%v/mods"
	pathUserPack       = "%s/api/users/%v/packs"
	pathKeys           = "%s/api/keys"
	pathKey            = "%s/api/keys/%v"
)

// DefaultClient implements the client interface.
type DefaultClient struct {
	client *http.Client
	base   string
}

// NewClient returns a client for the specified URL.
func NewClient(uri string) ClientAPI {
	return &DefaultClient{
		http.DefaultClient,
		uri,
	}
}

// NewClientToken returns a client that authenticates
// all outbound requests with the given token.
func NewClientToken(uri, token string) ClientAPI {
	config := oauth2.Config{}

	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: token,
		},
	)

	auther.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: syscerts.SystemRootsPool(),
		},
	}

	return &DefaultClient{
		auther,
		uri,
	}
}

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the Solder API.
func (c *DefaultClient) SetClient(client *http.Client) {
	c.client = client
}

// ProfileToken returns a profile.
func (c *DefaultClient) ProfileToken() (*Token, error) {
	out := &Token{}

	uri := fmt.Sprintf(pathProfileToken, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfileGet returns a profile.
func (c *DefaultClient) ProfileGet() (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfilePatch updates a profile.
func (c *DefaultClient) ProfilePatch(in *Profile) (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.patch(uri, in, out)

	return out, err
}

// ForgeList returns a list of all Forge versions.
func (c *DefaultClient) ForgeList() ([]*Forge, error) {
	var out []*Forge

	uri := fmt.Sprintf(pathForge, c.base)
	err := c.get(uri, &out)

	return out, err
}

// ForgeGet returns a Forge.
func (c *DefaultClient) ForgeGet(id string) (*Forge, error) {
	var out []*Forge

	uri := fmt.Sprintf(pathForge, c.base)
	err := c.get(uri, &out)

	if err != nil {
		return nil, err
	}

	for _, row := range out {
		if row.Slug == id || strconv.FormatInt(row.ID, 10) == id {
			return row, nil
		}
	}

	return nil, fmt.Errorf("Failed to find matching Forge version")
}

// ForgeRefresh refreshs the available Forge versions.
func (c *DefaultClient) ForgeRefresh() error {
	uri := fmt.Sprintf(pathForge, c.base)
	err := c.patch(uri, nil, nil)

	return err
}

// ForgeBuildList returns a list of related builds for a Forge version.
func (c *DefaultClient) ForgeBuildList(opts ForgeBuildParams) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathForgeBuild, c.base, opts.Forge)
	err := c.get(uri, &out)

	return out, err
}

// ForgeBuildAppend appends a Forge version to a build.
func (c *DefaultClient) ForgeBuildAppend(opts ForgeBuildParams) error {
	uri := fmt.Sprintf(pathForgeBuild, c.base, opts.Forge)
	err := c.patch(uri, opts, nil)

	return err
}

// ForgeBuildDelete remove a Forge version from a build.
func (c *DefaultClient) ForgeBuildDelete(opts ForgeBuildParams) error {
	uri := fmt.Sprintf(pathForgeBuild, c.base, opts.Forge)
	err := c.delete(uri, opts)

	return err
}

// MinecraftList returns a list of all Minecraft versions.
func (c *DefaultClient) MinecraftList() ([]*Minecraft, error) {
	var out []*Minecraft

	uri := fmt.Sprintf(pathMinecraft, c.base)
	err := c.get(uri, &out)

	return out, err
}

// MinecraftGet returns a Minecraft.
func (c *DefaultClient) MinecraftGet(id string) (*Minecraft, error) {
	var out []*Minecraft

	uri := fmt.Sprintf(pathMinecraft, c.base)
	err := c.get(uri, &out)

	if err != nil {
		return nil, err
	}

	for _, row := range out {
		if row.Slug == id || strconv.FormatInt(row.ID, 10) == id {
			return row, nil
		}
	}

	return nil, fmt.Errorf("Failed to find matching Minecraft version")
}

// MinecraftRefresh refreshs the available Minecraft versions.
func (c *DefaultClient) MinecraftRefresh() error {
	uri := fmt.Sprintf(pathMinecraft, c.base)
	err := c.patch(uri, nil, nil)

	return err
}

// MinecraftBuildList returns a list of related builds for a Minecraft version.
func (c *DefaultClient) MinecraftBuildList(opts MinecraftBuildParams) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathMinecraftBuild, c.base, opts.Minecraft)
	err := c.get(uri, &out)

	return out, err
}

// MinecraftBuildAppend appends a Minecraft version to a build.
func (c *DefaultClient) MinecraftBuildAppend(opts MinecraftBuildParams) error {
	uri := fmt.Sprintf(pathMinecraftBuild, c.base, opts.Minecraft)
	err := c.patch(uri, opts, nil)

	return err
}

// MinecraftBuildDelete remove a Minecraft version from a build.
func (c *DefaultClient) MinecraftBuildDelete(opts MinecraftBuildParams) error {
	uri := fmt.Sprintf(pathMinecraftBuild, c.base, opts.Minecraft)
	err := c.delete(uri, opts)

	return err
}

// PackList returns a list of all packs.
func (c *DefaultClient) PackList() ([]*Pack, error) {
	var out []*Pack

	uri := fmt.Sprintf(pathPacks, c.base)
	err := c.get(uri, &out)

	return out, err
}

// PackGet returns a pack.
func (c *DefaultClient) PackGet(id string) (*Pack, error) {
	out := &Pack{}

	uri := fmt.Sprintf(pathPack, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// PackPost creates a pack.
func (c *DefaultClient) PackPost(in *Pack) (*Pack, error) {
	out := &Pack{}

	uri := fmt.Sprintf(pathPacks, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// PackPatch updates a pack.
func (c *DefaultClient) PackPatch(in *Pack) (*Pack, error) {
	out := &Pack{}

	uri := fmt.Sprintf(pathPack, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// PackDelete deletes a pack.
func (c *DefaultClient) PackDelete(id string) error {
	uri := fmt.Sprintf(pathPack, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// PackClientList returns a list of related clients for a pack.
func (c *DefaultClient) PackClientList(opts PackClientParams) ([]*Client, error) {
	var out []*Client

	uri := fmt.Sprintf(pathPackClient, c.base, opts.Pack)
	err := c.get(uri, &out)

	return out, err
}

// PackClientAppend appends a client to a pack.
func (c *DefaultClient) PackClientAppend(opts PackClientParams) error {
	uri := fmt.Sprintf(pathPackClient, c.base, opts.Pack)
	err := c.patch(uri, opts, nil)

	return err
}

// PackClientDelete remove a client from a pack.
func (c *DefaultClient) PackClientDelete(opts PackClientParams) error {
	uri := fmt.Sprintf(pathPackClient, c.base, opts.Pack)
	err := c.delete(uri, opts)

	return err
}

// PackUserList returns a list of related users for a pack.
func (c *DefaultClient) PackUserList(opts PackUserParams) ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathPackUser, c.base, opts.Pack)
	err := c.get(uri, &out)

	return out, err
}

// PackUserAppend appends a user to a pack.
func (c *DefaultClient) PackUserAppend(opts PackUserParams) error {
	uri := fmt.Sprintf(pathPackUser, c.base, opts.Pack)
	err := c.patch(uri, opts, nil)

	return err
}

// PackUserDelete remove a user from a pack.
func (c *DefaultClient) PackUserDelete(opts PackUserParams) error {
	uri := fmt.Sprintf(pathPackUser, c.base, opts.Pack)
	err := c.delete(uri, opts)

	return err
}

// BuildList returns a list of all builds for a specific pack.
func (c *DefaultClient) BuildList(pack string) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathBuilds, c.base, pack)
	err := c.get(uri, &out)

	return out, err
}

// BuildGet returns a build for a specific pack.
func (c *DefaultClient) BuildGet(pack, id string) (*Build, error) {
	out := &Build{}

	uri := fmt.Sprintf(pathBuild, c.base, pack, id)
	err := c.get(uri, out)

	return out, err
}

// BuildPost creates a build for a specific pack.
func (c *DefaultClient) BuildPost(pack string, in *Build) (*Build, error) {
	out := &Build{}

	uri := fmt.Sprintf(pathBuilds, c.base, pack)
	err := c.post(uri, in, out)

	return out, err
}

// BuildPatch updates a build for a specific pack.
func (c *DefaultClient) BuildPatch(pack string, in *Build) (*Build, error) {
	out := &Build{}

	uri := fmt.Sprintf(pathBuild, c.base, pack, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// BuildDelete deletes a build for a specific pack.
func (c *DefaultClient) BuildDelete(pack, id string) error {
	uri := fmt.Sprintf(pathBuild, c.base, pack, id)
	err := c.delete(uri, nil)

	return err
}

// BuildVersionList returns a list of related versions for a build.
func (c *DefaultClient) BuildVersionList(opts BuildVersionParams) ([]*Version, error) {
	var out []*Version

	uri := fmt.Sprintf(pathBuildVersion, c.base, opts.Pack, opts.Build)
	err := c.get(uri, &out)

	return out, err
}

// BuildVersionAppend appends a version to a build.
func (c *DefaultClient) BuildVersionAppend(opts BuildVersionParams) error {
	uri := fmt.Sprintf(pathBuildVersion, c.base, opts.Pack, opts.Build)
	err := c.patch(uri, opts, nil)

	return err
}

// BuildVersionDelete remove a version from a build.
func (c *DefaultClient) BuildVersionDelete(opts BuildVersionParams) error {
	uri := fmt.Sprintf(pathBuildVersion, c.base, opts.Pack, opts.Build)
	err := c.delete(uri, opts)

	return err
}

// ModList returns a list of all mods.
func (c *DefaultClient) ModList() ([]*Mod, error) {
	var out []*Mod

	uri := fmt.Sprintf(pathMods, c.base)
	err := c.get(uri, &out)

	return out, err
}

// ModGet returns a mod.
func (c *DefaultClient) ModGet(id string) (*Mod, error) {
	out := &Mod{}

	uri := fmt.Sprintf(pathMod, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// ModPost creates a mod.
func (c *DefaultClient) ModPost(in *Mod) (*Mod, error) {
	out := &Mod{}

	uri := fmt.Sprintf(pathMods, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// ModPatch updates a mod.
func (c *DefaultClient) ModPatch(in *Mod) (*Mod, error) {
	out := &Mod{}

	uri := fmt.Sprintf(pathMod, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// ModDelete deletes a mod.
func (c *DefaultClient) ModDelete(id string) error {
	uri := fmt.Sprintf(pathMod, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// ModUserList returns a list of related users for a mod.
func (c *DefaultClient) ModUserList(opts ModUserParams) ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathModUser, c.base, opts.Mod)
	err := c.get(uri, &out)

	return out, err
}

// ModUserAppend appends a user to a mod.
func (c *DefaultClient) ModUserAppend(opts ModUserParams) error {
	uri := fmt.Sprintf(pathModUser, c.base, opts.Mod)
	err := c.patch(uri, opts, nil)

	return err
}

// ModUserDelete remove a user from a mod.
func (c *DefaultClient) ModUserDelete(opts ModUserParams) error {
	uri := fmt.Sprintf(pathModUser, c.base, opts.Mod)
	err := c.delete(uri, opts)

	return err
}

// VersionList returns a list of all versions for a specific mod.
func (c *DefaultClient) VersionList(mod string) ([]*Version, error) {
	var out []*Version

	uri := fmt.Sprintf(pathVersions, c.base, mod)
	err := c.get(uri, &out)

	return out, err
}

// VersionGet returns a version for a specific mod.
func (c *DefaultClient) VersionGet(mod, id string) (*Version, error) {
	out := &Version{}

	uri := fmt.Sprintf(pathVersion, c.base, mod, id)
	err := c.get(uri, out)

	return out, err
}

// VersionPost creates a version for a specific mod.
func (c *DefaultClient) VersionPost(mod string, in *Version) (*Version, error) {
	out := &Version{}

	uri := fmt.Sprintf(pathVersions, c.base, mod)
	err := c.post(uri, in, out)

	return out, err
}

// VersionPatch updates a version for a specific mod.
func (c *DefaultClient) VersionPatch(mod string, in *Version) (*Version, error) {
	out := &Version{}

	uri := fmt.Sprintf(pathVersion, c.base, mod, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// VersionDelete deletes a version for a specific mod.
func (c *DefaultClient) VersionDelete(mod, id string) error {
	uri := fmt.Sprintf(pathVersion, c.base, mod, id)
	err := c.delete(uri, nil)

	return err
}

// VersionBuildList returns a list of related builds for a version.
func (c *DefaultClient) VersionBuildList(opts VersionBuildParams) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathVersionBuild, c.base, opts.Mod, opts.Version)
	err := c.get(uri, &out)

	return out, err
}

// VersionBuildAppend appends a build to a version.
func (c *DefaultClient) VersionBuildAppend(opts VersionBuildParams) error {
	uri := fmt.Sprintf(pathVersionBuild, c.base, opts.Mod, opts.Version)
	err := c.patch(uri, opts, nil)

	return err
}

// VersionBuildDelete remove a build from a version.
func (c *DefaultClient) VersionBuildDelete(opts VersionBuildParams) error {
	uri := fmt.Sprintf(pathVersionBuild, c.base, opts.Mod, opts.Version)
	err := c.delete(uri, opts)

	return err
}

// ClientList returns a list of all clients.
func (c *DefaultClient) ClientList() ([]*Client, error) {
	var out []*Client

	uri := fmt.Sprintf(pathClients, c.base)
	err := c.get(uri, &out)

	return out, err
}

// ClientGet returns a client.
func (c *DefaultClient) ClientGet(id string) (*Client, error) {
	out := &Client{}

	uri := fmt.Sprintf(pathClient, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// ClientPost creates a client.
func (c *DefaultClient) ClientPost(in *Client) (*Client, error) {
	out := &Client{}

	uri := fmt.Sprintf(pathClients, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// ClientPatch updates a client.
func (c *DefaultClient) ClientPatch(in *Client) (*Client, error) {
	out := &Client{}

	uri := fmt.Sprintf(pathClient, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// ClientDelete deletes a client.
func (c *DefaultClient) ClientDelete(id string) error {
	uri := fmt.Sprintf(pathClient, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// ClientPackList returns a list of related packs for a client.
func (c *DefaultClient) ClientPackList(opts ClientPackParams) ([]*Pack, error) {
	var out []*Pack

	uri := fmt.Sprintf(pathClientPack, c.base, opts.Client)
	err := c.get(uri, &out)

	return out, err
}

// ClientPackAppend appends a pack to a client.
func (c *DefaultClient) ClientPackAppend(opts ClientPackParams) error {
	uri := fmt.Sprintf(pathClientPack, c.base, opts.Client)
	err := c.patch(uri, opts, nil)

	return err
}

// ClientPackDelete remove a pack from a client.
func (c *DefaultClient) ClientPackDelete(opts ClientPackParams) error {
	uri := fmt.Sprintf(pathClientPack, c.base, opts.Client)
	err := c.delete(uri, opts)

	return err
}

// UserList returns a list of all users.
func (c *DefaultClient) UserList() ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.get(uri, &out)

	return out, err
}

// UserGet returns a user.
func (c *DefaultClient) UserGet(id string) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// UserPost creates a user.
func (c *DefaultClient) UserPost(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// UserPatch updates a user.
func (c *DefaultClient) UserPatch(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// UserDelete deletes a user.
func (c *DefaultClient) UserDelete(id string) error {
	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// UserModList returns a list of related mods for a user.
func (c *DefaultClient) UserModList(opts UserModParams) ([]*Mod, error) {
	var out []*Mod

	uri := fmt.Sprintf(pathUserMod, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserModAppend appends a mod to a user.
func (c *DefaultClient) UserModAppend(opts UserModParams) error {
	uri := fmt.Sprintf(pathUserMod, c.base, opts.User)
	err := c.patch(uri, opts, nil)

	return err
}

// UserModDelete remove a mod from a user.
func (c *DefaultClient) UserModDelete(opts UserModParams) error {
	uri := fmt.Sprintf(pathUserMod, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// UserPackList returns a list of related packs for a user.
func (c *DefaultClient) UserPackList(opts UserPackParams) ([]*Pack, error) {
	var out []*Pack

	uri := fmt.Sprintf(pathUserPack, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserPackAppend appends a pack to a user.
func (c *DefaultClient) UserPackAppend(opts UserPackParams) error {
	uri := fmt.Sprintf(pathUserPack, c.base, opts.User)
	err := c.patch(uri, opts, nil)

	return err
}

// UserPackDelete remove a pack from a user.
func (c *DefaultClient) UserPackDelete(opts UserPackParams) error {
	uri := fmt.Sprintf(pathUserPack, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// KeyList returns a list of all keys.
func (c *DefaultClient) KeyList() ([]*Key, error) {
	var out []*Key

	uri := fmt.Sprintf(pathKeys, c.base)
	err := c.get(uri, &out)

	return out, err
}

// KeyGet returns a key.
func (c *DefaultClient) KeyGet(id string) (*Key, error) {
	out := &Key{}

	uri := fmt.Sprintf(pathKey, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// KeyPost creates a key.
func (c *DefaultClient) KeyPost(in *Key) (*Key, error) {
	out := &Key{}

	uri := fmt.Sprintf(pathKeys, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// KeyPatch updates a key.
func (c *DefaultClient) KeyPatch(in *Key) (*Key, error) {
	out := &Key{}

	uri := fmt.Sprintf(pathKey, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// KeyDelete deletes a key.
func (c *DefaultClient) KeyDelete(id string) error {
	uri := fmt.Sprintf(pathKey, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// Helper function for making an GET request.
func (c *DefaultClient) get(rawurl string, out interface{}) error {
	return c.do(rawurl, "GET", nil, out)
}

// Helper function for making an POST request.
func (c *DefaultClient) post(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "POST", in, out)
}

// Helper function for making an PUT request.
func (c *DefaultClient) put(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PUT", in, out)
}

// Helper function for making an PATCH request.
func (c *DefaultClient) patch(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PATCH", in, out)
}

// Helper function for making an DELETE request.
func (c *DefaultClient) delete(rawurl string, in interface{}) error {
	return c.do(rawurl, "DELETE", in, nil)
}

// Helper function to make an HTTP request
func (c *DefaultClient) do(rawurl, method string, in, out interface{}) error {
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
func (c *DefaultClient) stream(rawurl, method string, in, out interface{}) (io.ReadCloser, error) {
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

	req.Header.Set(
		"User-Agent",
		"Solder CLI",
	)

	if in != nil {
		req.Header.Set(
			"Content-Type",
			"application/json",
		)
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
