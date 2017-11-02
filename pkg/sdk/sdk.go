package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jackspirou/syscerts"
	"golang.org/x/oauth2"
)

const (
	pathAuthLogin      = "%s/api/auth/login"
	pathProfile        = "%s/api/profile/self"
	pathProfileToken   = "%s/api/profile/token"
	pathKeys           = "%s/api/keys"
	pathKey            = "%s/api/keys/%v"
	pathForge          = "%s/api/forge"
	pathForgeBuild     = "%s/api/forge/%v/builds"
	pathMinecraft      = "%s/api/minecraft"
	pathMinecraftBuild = "%s/api/minecraft/%v/builds"
	pathPacks          = "%s/api/packs"
	pathPack           = "%s/api/packs/%v"
	pathPackClient     = "%s/api/packs/%v/clients"
	pathPackUser       = "%s/api/packs/%v/users"
	pathPackTeam       = "%s/api/packs/%v/teams"
	pathBuilds         = "%s/api/packs/%v/builds"
	pathBuild          = "%s/api/packs/%v/builds/%v"
	pathBuildVersion   = "%s/api/packs/%v/builds/%v/versions"
	pathMods           = "%s/api/mods"
	pathMod            = "%s/api/mods/%v"
	pathModUser        = "%s/api/mods/%v/users"
	pathModTeam        = "%s/api/mods/%v/teams"
	pathVersions       = "%s/api/mods/%v/versions"
	pathVersion        = "%s/api/mods/%v/versions/%v"
	pathVersionBuild   = "%s/api/mods/%v/versions/%v/builds"
	pathClients        = "%s/api/clients"
	pathClient         = "%s/api/clients/%v"
	pathClientPack     = "%s/api/clients/%v/packs"
	pathUsers          = "%s/api/users"
	pathUser           = "%s/api/users/%v"
	pathUserTeam       = "%s/api/users/%v/teams"
	pathUserMod        = "%s/api/users/%v/mods"
	pathUserPack       = "%s/api/users/%v/packs"
	pathTeams          = "%s/api/teams"
	pathTeam           = "%s/api/teams/%v"
	pathTeamUser       = "%s/api/teams/%v/users"
	pathTeamMod        = "%s/api/teams/%v/mods"
	pathTeamPack       = "%s/api/teams/%v/packs"
)

// ClientAPI describes a client API.
type ClientAPI interface {
	// SetClient sets the default http client. This should
	// be used in conjunction with golang.org/x/oauth2 to
	// authenticate requests to the Kleister API.
	SetClient(client *http.Client)

	// IsAuthenticated checks if we already provided an authentication
	// token for our client requests. If it returns false you can update
	// the client after fetching a valid token.
	IsAuthenticated() bool

	// AuthLogin signs in based on credentials and returns a token.
	AuthLogin(string, string) (*Token, error)

	// ProfileToken returns a token.
	ProfileToken() (*Token, error)

	// ProfileGet returns a profile.
	ProfileGet() (*Profile, error)

	// ProfilePatch updates a profile.
	ProfilePatch(*Profile) (*Profile, error)

	// KeyList returns a list of all keys.
	KeyList() ([]*Key, error)

	// KeyGet returns a key.
	KeyGet(string) (*Key, error)

	// KeyPost creates a key.
	KeyPost(*Key) (*Key, error)

	// KeyPatch updates a key.
	KeyPatch(*Key) (*Key, error)

	// KeyDelete deletes a key.
	KeyDelete(string) error

	// ForgeList returns a list of all Forge versions.
	ForgeList() ([]*Forge, error)

	// ForgeGet returns a Forge.
	ForgeGet(string) (*Forge, error)

	// ForgeRefresh refreshs the available Forge versions.
	ForgeRefresh() error

	// ForgeBuildList returns a list of related builds for a Forge version.
	ForgeBuildList(ForgeBuildParams) ([]*Build, error)

	// ForgeBuildAppend appends a Forge version to a build.
	ForgeBuildAppend(ForgeBuildParams) error

	// ForgeBuildAppend remove a Forge version from a build.
	ForgeBuildDelete(ForgeBuildParams) error

	// MinecraftList returns a list of all Minecraft versions.
	MinecraftList() ([]*Minecraft, error)

	// MinecraftGet returns a Minecraft.
	MinecraftGet(string) (*Minecraft, error)

	// MinecraftRefresh refreshs the available Minecraft versions.
	MinecraftRefresh() error

	// MinecraftBuildList returns a list of related builds for a Minecraft version.
	MinecraftBuildList(MinecraftBuildParams) ([]*Build, error)

	// MinecraftBuildAppend appends a Minecraft version to a build.
	MinecraftBuildAppend(MinecraftBuildParams) error

	// MinecraftBuildAppend remove a Minecraft version from a build.
	MinecraftBuildDelete(MinecraftBuildParams) error

	// PackList returns a list of all packs.
	PackList() ([]*Pack, error)

	// PackGet returns a pack.
	PackGet(string) (*Pack, error)

	// PackPost creates a pack.
	PackPost(*Pack) (*Pack, error)

	// PackPatch updates a pack.
	PackPatch(*Pack) (*Pack, error)

	// PackDelete deletes a pack.
	PackDelete(string) error

	// PackClientList returns a list of related clients for a pack.
	PackClientList(PackClientParams) ([]*ClientPack, error)

	// PackClientAppend appends a client to a pack.
	PackClientAppend(PackClientParams) error

	// PackClientDelete remove a client from a pack.
	PackClientDelete(PackClientParams) error

	// PackUserList returns a list of related users for a pack.
	PackUserList(PackUserParams) ([]*UserPack, error)

	// PackUserAppend appends a user to a pack.
	PackUserAppend(PackUserParams) error

	// PackUserPerm updates perms for pack user.
	PackUserPerm(PackUserParams) error

	// PackUserDelete remove a user from a pack.
	PackUserDelete(PackUserParams) error

	// PackTeamList returns a list of related teams for a pack.
	PackTeamList(PackTeamParams) ([]*TeamPack, error)

	// PackTeamAppend appends a team to a pack.
	PackTeamAppend(PackTeamParams) error

	// PackTeamPerm updates perms for pack team.
	PackTeamPerm(PackTeamParams) error

	// PackTeamDelete remove a team from a pack.
	PackTeamDelete(PackTeamParams) error

	// BuildList returns a list of all builds for a specific pack.
	BuildList(string) ([]*Build, error)

	// BuildGet returns a build for a specific pack.
	BuildGet(string, string) (*Build, error)

	// BuildPost creates a build for a specific pack.
	BuildPost(string, *Build) (*Build, error)

	// BuildPatch updates a build for a specific pack.
	BuildPatch(string, *Build) (*Build, error)

	// BuildDelete deletes a build for a specific pack.
	BuildDelete(string, string) error

	// BuildVersionList returns a list of related versions for a build.
	BuildVersionList(BuildVersionParams) ([]*BuildVersion, error)

	// BuildVersionAppend appends a version to a build.
	BuildVersionAppend(BuildVersionParams) error

	// BuildVersionDelete remove a version from a build.
	BuildVersionDelete(BuildVersionParams) error

	// ModList returns a list of all mods.
	ModList() ([]*Mod, error)

	// ModGet returns a mod.
	ModGet(string) (*Mod, error)

	// ModPost creates a mod.
	ModPost(*Mod) (*Mod, error)

	// ModPatch updates a mod.
	ModPatch(*Mod) (*Mod, error)

	// ModDelete deletes a mod.
	ModDelete(string) error

	// ModUserList returns a list of related users for a mod.
	ModUserList(ModUserParams) ([]*UserMod, error)

	// ModUserAppend appends a user to a mod.
	ModUserAppend(ModUserParams) error

	// ModUserPerm updates perms for mod user.
	ModUserPerm(ModUserParams) error

	// ModUserDelete remove a user from a mod.
	ModUserDelete(ModUserParams) error

	// ModTeamList returns a list of related teams for a mod.
	ModTeamList(ModTeamParams) ([]*TeamMod, error)

	// ModTeamAppend appends a team to a mod.
	ModTeamAppend(ModTeamParams) error

	// ModTeamPerm updates perms for mod team.
	ModTeamPerm(ModTeamParams) error

	// ModTeamDelete remove a team from a mod.
	ModTeamDelete(ModTeamParams) error

	// VersionList returns a list of all versions for a specific mod.
	VersionList(string) ([]*Version, error)

	// VersionGet returns a version for a specific mod.
	VersionGet(string, string) (*Version, error)

	// VersionPost creates a version for a specific mod.
	VersionPost(string, *Version) (*Version, error)

	// VersionPatch updates a version for a specific mod.
	VersionPatch(string, *Version) (*Version, error)

	// VersionDelete deletes a version for a specific mod.
	VersionDelete(string, string) error

	// VersionBuildList returns a list of related builds for a version.
	VersionBuildList(VersionBuildParams) ([]*BuildVersion, error)

	// VersionBuildAppend appends a build to a version.
	VersionBuildAppend(VersionBuildParams) error

	// VersionBuildDelete remove a build from a version.
	VersionBuildDelete(VersionBuildParams) error

	// ClientList returns a list of all clients.
	ClientList() ([]*Client, error)

	// ClientGet returns a client.
	ClientGet(string) (*Client, error)

	// ClientPost creates a client.
	ClientPost(*Client) (*Client, error)

	// ClientPatch updates a client.
	ClientPatch(*Client) (*Client, error)

	// ClientDelete deletes a client.
	ClientDelete(string) error

	// ClientPackList returns a list of related packs for a client.
	ClientPackList(ClientPackParams) ([]*ClientPack, error)

	// ClientPackAppend appends a pack to a client.
	ClientPackAppend(ClientPackParams) error

	// ClientPackDelete remove a pack from a client.
	ClientPackDelete(ClientPackParams) error

	// UserList returns a list of all users.
	UserList() ([]*User, error)

	// UserGet returns a user.
	UserGet(string) (*User, error)

	// UserPost creates a user.
	UserPost(*User) (*User, error)

	// UserPatch updates a user.
	UserPatch(*User) (*User, error)

	// UserDelete deletes a user.
	UserDelete(string) error

	// UserModList returns a list of related mods for a user.
	UserModList(UserModParams) ([]*UserMod, error)

	// UserModAppend appends a mod to a user.
	UserModAppend(UserModParams) error

	// UserModPerm updates perms for user mod.
	UserModPerm(UserModParams) error

	// UserModDelete remove a mod from a user.
	UserModDelete(UserModParams) error

	// UserPackList returns a list of related packs for a user.
	UserPackList(UserPackParams) ([]*UserPack, error)

	// UserPackAppend appends a pack to a user.
	UserPackAppend(UserPackParams) error

	// UserPackPerm updates perms for user pack.
	UserPackPerm(UserPackParams) error

	// UserPackDelete remove a pack from a user.
	UserPackDelete(UserPackParams) error

	// UserTeamList returns a list of related teams for a user.
	UserTeamList(UserTeamParams) ([]*TeamUser, error)

	// UserTeamAppend appends a team to a user.
	UserTeamAppend(UserTeamParams) error

	// UserTeamPerm updates perms for user team.
	UserTeamPerm(UserTeamParams) error

	// UserTeamDelete remove a team from a user.
	UserTeamDelete(UserTeamParams) error

	// TeamList returns a list of all teams.
	TeamList() ([]*Team, error)

	// TeamGet returns a team.
	TeamGet(string) (*Team, error)

	// TeamPost creates a team.
	TeamPost(*Team) (*Team, error)

	// TeamPatch updates a team.
	TeamPatch(*Team) (*Team, error)

	// TeamDelete deletes a team.
	TeamDelete(string) error

	// TeamUserList returns a list of related users for a team.
	TeamUserList(TeamUserParams) ([]*TeamUser, error)

	// TeamUserAppend appends a user to a team.
	TeamUserAppend(TeamUserParams) error

	// TeamUserPerm updates perms for team user.
	TeamUserPerm(TeamUserParams) error

	// TeamUserDelete remove a user from a team.
	TeamUserDelete(TeamUserParams) error

	// TeamModList returns a list of related mods for a team.
	TeamModList(TeamModParams) ([]*TeamMod, error)

	// TeamModAppend appends a mod to a team.
	TeamModAppend(TeamModParams) error

	// TeamModPerm updates perms for team mod.
	TeamModPerm(TeamModParams) error

	// TeamModDelete remove a mod from a team.
	TeamModDelete(TeamModParams) error

	// TeamPackList returns a list of related packs for a team.
	TeamPackList(TeamPackParams) ([]*TeamPack, error)

	// TeamPackAppend appends a pack to a team.
	TeamPackAppend(TeamPackParams) error

	// TeamPackPerm updates perms for team pack.
	TeamPackPerm(TeamPackParams) error

	// TeamPackDelete remove a pack from a team.
	TeamPackDelete(TeamPackParams) error
}

// Default implements the client interface.
type Default struct {
	client *http.Client
	base   string
	token  string
}

// NewClient returns a client for the specified URL.
func NewClient(uri string) ClientAPI {
	return &Default{
		client: http.DefaultClient,
		base:   uri,
	}
}

// NewClientToken returns a client that authenticates
// all outbound requests with the given token.
func NewClientToken(uri, token string) ClientAPI {
	config := oauth2.Config{}

	client := config.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: token,
		},
	)

	if trans, ok := client.Transport.(*oauth2.Transport); ok {
		trans.Base = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				RootCAs: syscerts.SystemRootsPool(),
			},
		}
	}

	return &Default{
		client: client,
		base:   uri,
		token:  token,
	}
}

// IsAuthenticated checks if we already provided an authentication
// token for our client requests. If it returns false you can update
// the client after fetching a valid token.
func (c *Default) IsAuthenticated() bool {
	if c.token == "" {
		return false
	}

	uri, err := url.Parse(fmt.Sprintf(pathProfileToken, c.base))

	if err != nil {
		return false
	}

	req, err := http.NewRequest("GET", uri.String(), nil)

	if err != nil {
		return false
	}

	req.Header.Set(
		"User-Agent",
		"Kleister CLI",
	)

	resp, err := c.client.Do(req)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusUnauthorized
}

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the Kleister API.
func (c *Default) SetClient(client *http.Client) {
	c.client = client
}

// AuthLogin signs in based on credentials and returns a token.
func (c *Default) AuthLogin(username, password string) (*Token, error) {
	out := &Token{}

	in := struct {
		Username string
		Password string
	}{
		username,
		password,
	}

	uri := fmt.Sprintf(pathAuthLogin, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// ProfileToken returns a profile.
func (c *Default) ProfileToken() (*Token, error) {
	out := &Token{}

	uri := fmt.Sprintf(pathProfileToken, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfileGet returns a profile.
func (c *Default) ProfileGet() (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfilePatch updates a profile.
func (c *Default) ProfilePatch(in *Profile) (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.put(uri, in, out)

	return out, err
}

// KeyList returns a list of all keys.
func (c *Default) KeyList() ([]*Key, error) {
	var out []*Key

	uri := fmt.Sprintf(pathKeys, c.base)
	err := c.get(uri, &out)

	return out, err
}

// KeyGet returns a key.
func (c *Default) KeyGet(id string) (*Key, error) {
	out := &Key{}

	uri := fmt.Sprintf(pathKey, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// KeyPost creates a key.
func (c *Default) KeyPost(in *Key) (*Key, error) {
	out := &Key{}

	uri := fmt.Sprintf(pathKeys, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// KeyPatch updates a key.
func (c *Default) KeyPatch(in *Key) (*Key, error) {
	out := &Key{}

	uri := fmt.Sprintf(pathKey, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// KeyDelete deletes a key.
func (c *Default) KeyDelete(id string) error {
	uri := fmt.Sprintf(pathKey, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// ForgeList returns a list of all Forge versions.
func (c *Default) ForgeList() ([]*Forge, error) {
	var out []*Forge

	uri := fmt.Sprintf(pathForge, c.base)
	err := c.get(uri, &out)

	return out, err
}

// ForgeGet returns a Forge.
func (c *Default) ForgeGet(id string) (*Forge, error) {
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
func (c *Default) ForgeRefresh() error {
	uri := fmt.Sprintf(pathForge, c.base)
	err := c.put(uri, nil, nil)

	return err
}

// ForgeBuildList returns a list of related builds for a Forge version.
func (c *Default) ForgeBuildList(opts ForgeBuildParams) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathForgeBuild, c.base, opts.Forge)
	err := c.get(uri, &out)

	return out, err
}

// ForgeBuildAppend appends a Forge version to a build.
func (c *Default) ForgeBuildAppend(opts ForgeBuildParams) error {
	uri := fmt.Sprintf(pathForgeBuild, c.base, opts.Forge)
	err := c.post(uri, opts, nil)

	return err
}

// ForgeBuildDelete remove a Forge version from a build.
func (c *Default) ForgeBuildDelete(opts ForgeBuildParams) error {
	uri := fmt.Sprintf(pathForgeBuild, c.base, opts.Forge)
	err := c.delete(uri, opts)

	return err
}

// MinecraftList returns a list of all Minecraft versions.
func (c *Default) MinecraftList() ([]*Minecraft, error) {
	var out []*Minecraft

	uri := fmt.Sprintf(pathMinecraft, c.base)
	err := c.get(uri, &out)

	return out, err
}

// MinecraftGet returns a Minecraft.
func (c *Default) MinecraftGet(id string) (*Minecraft, error) {
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
func (c *Default) MinecraftRefresh() error {
	uri := fmt.Sprintf(pathMinecraft, c.base)
	err := c.put(uri, nil, nil)

	return err
}

// MinecraftBuildList returns a list of related builds for a Minecraft version.
func (c *Default) MinecraftBuildList(opts MinecraftBuildParams) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathMinecraftBuild, c.base, opts.Minecraft)
	err := c.get(uri, &out)

	return out, err
}

// MinecraftBuildAppend appends a Minecraft version to a build.
func (c *Default) MinecraftBuildAppend(opts MinecraftBuildParams) error {
	uri := fmt.Sprintf(pathMinecraftBuild, c.base, opts.Minecraft)
	err := c.post(uri, opts, nil)

	return err
}

// MinecraftBuildDelete remove a Minecraft version from a build.
func (c *Default) MinecraftBuildDelete(opts MinecraftBuildParams) error {
	uri := fmt.Sprintf(pathMinecraftBuild, c.base, opts.Minecraft)
	err := c.delete(uri, opts)

	return err
}

// PackList returns a list of all packs.
func (c *Default) PackList() ([]*Pack, error) {
	var out []*Pack

	uri := fmt.Sprintf(pathPacks, c.base)
	err := c.get(uri, &out)

	return out, err
}

// PackGet returns a pack.
func (c *Default) PackGet(id string) (*Pack, error) {
	out := &Pack{}

	uri := fmt.Sprintf(pathPack, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// PackPost creates a pack.
func (c *Default) PackPost(in *Pack) (*Pack, error) {
	out := &Pack{}

	uri := fmt.Sprintf(pathPacks, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// PackPatch updates a pack.
func (c *Default) PackPatch(in *Pack) (*Pack, error) {
	out := &Pack{}

	uri := fmt.Sprintf(pathPack, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// PackDelete deletes a pack.
func (c *Default) PackDelete(id string) error {
	uri := fmt.Sprintf(pathPack, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// PackClientList returns a list of related clients for a pack.
func (c *Default) PackClientList(opts PackClientParams) ([]*ClientPack, error) {
	var out []*ClientPack

	uri := fmt.Sprintf(pathPackClient, c.base, opts.Pack)
	err := c.get(uri, &out)

	return out, err
}

// PackClientAppend appends a client to a pack.
func (c *Default) PackClientAppend(opts PackClientParams) error {
	uri := fmt.Sprintf(pathPackClient, c.base, opts.Pack)
	err := c.post(uri, opts, nil)

	return err
}

// PackClientDelete remove a client from a pack.
func (c *Default) PackClientDelete(opts PackClientParams) error {
	uri := fmt.Sprintf(pathPackClient, c.base, opts.Pack)
	err := c.delete(uri, opts)

	return err
}

// PackUserList returns a list of related users for a pack.
func (c *Default) PackUserList(opts PackUserParams) ([]*UserPack, error) {
	var out []*UserPack

	uri := fmt.Sprintf(pathPackUser, c.base, opts.Pack)
	err := c.get(uri, &out)

	return out, err
}

// PackUserAppend appends a user to a pack.
func (c *Default) PackUserAppend(opts PackUserParams) error {
	uri := fmt.Sprintf(pathPackUser, c.base, opts.Pack)
	err := c.post(uri, opts, nil)

	return err
}

// PackUserPerm updates perms for pack team.
func (c *Default) PackUserPerm(opts PackUserParams) error {
	uri := fmt.Sprintf(pathPackUser, c.base, opts.Pack)
	err := c.put(uri, opts, nil)

	return err
}

// PackUserDelete remove a user from a pack.
func (c *Default) PackUserDelete(opts PackUserParams) error {
	uri := fmt.Sprintf(pathPackUser, c.base, opts.Pack)
	err := c.delete(uri, opts)

	return err
}

// PackTeamList returns a list of related teams for a pack.
func (c *Default) PackTeamList(opts PackTeamParams) ([]*TeamPack, error) {
	var out []*TeamPack

	uri := fmt.Sprintf(pathPackTeam, c.base, opts.Pack)
	err := c.get(uri, &out)

	return out, err
}

// PackTeamAppend appends a team to a pack.
func (c *Default) PackTeamAppend(opts PackTeamParams) error {
	uri := fmt.Sprintf(pathPackTeam, c.base, opts.Pack)
	err := c.post(uri, opts, nil)

	return err
}

// PackTeamPerm updates perms for pack team.
func (c *Default) PackTeamPerm(opts PackTeamParams) error {
	uri := fmt.Sprintf(pathPackTeam, c.base, opts.Pack)
	err := c.put(uri, opts, nil)

	return err
}

// PackTeamDelete remove a team from a pack.
func (c *Default) PackTeamDelete(opts PackTeamParams) error {
	uri := fmt.Sprintf(pathPackTeam, c.base, opts.Pack)
	err := c.delete(uri, opts)

	return err
}

// BuildList returns a list of all builds for a specific pack.
func (c *Default) BuildList(pack string) ([]*Build, error) {
	var out []*Build

	uri := fmt.Sprintf(pathBuilds, c.base, pack)
	err := c.get(uri, &out)

	return out, err
}

// BuildGet returns a build for a specific pack.
func (c *Default) BuildGet(pack, id string) (*Build, error) {
	out := &Build{}

	uri := fmt.Sprintf(pathBuild, c.base, pack, id)
	err := c.get(uri, out)

	return out, err
}

// BuildPost creates a build for a specific pack.
func (c *Default) BuildPost(pack string, in *Build) (*Build, error) {
	out := &Build{}

	uri := fmt.Sprintf(pathBuilds, c.base, pack)
	err := c.post(uri, in, out)

	return out, err
}

// BuildPatch updates a build for a specific pack.
func (c *Default) BuildPatch(pack string, in *Build) (*Build, error) {
	out := &Build{}

	uri := fmt.Sprintf(pathBuild, c.base, pack, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// BuildDelete deletes a build for a specific pack.
func (c *Default) BuildDelete(pack, id string) error {
	uri := fmt.Sprintf(pathBuild, c.base, pack, id)
	err := c.delete(uri, nil)

	return err
}

// BuildVersionList returns a list of related versions for a build.
func (c *Default) BuildVersionList(opts BuildVersionParams) ([]*BuildVersion, error) {
	var out []*BuildVersion

	uri := fmt.Sprintf(pathBuildVersion, c.base, opts.Pack, opts.Build)
	err := c.get(uri, &out)

	return out, err
}

// BuildVersionAppend appends a version to a build.
func (c *Default) BuildVersionAppend(opts BuildVersionParams) error {
	uri := fmt.Sprintf(pathBuildVersion, c.base, opts.Pack, opts.Build)
	err := c.post(uri, opts, nil)

	return err
}

// BuildVersionDelete remove a version from a build.
func (c *Default) BuildVersionDelete(opts BuildVersionParams) error {
	uri := fmt.Sprintf(pathBuildVersion, c.base, opts.Pack, opts.Build)
	err := c.delete(uri, opts)

	return err
}

// ModList returns a list of all mods.
func (c *Default) ModList() ([]*Mod, error) {
	var out []*Mod

	uri := fmt.Sprintf(pathMods, c.base)
	err := c.get(uri, &out)

	return out, err
}

// ModGet returns a mod.
func (c *Default) ModGet(id string) (*Mod, error) {
	out := &Mod{}

	uri := fmt.Sprintf(pathMod, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// ModPost creates a mod.
func (c *Default) ModPost(in *Mod) (*Mod, error) {
	out := &Mod{}

	uri := fmt.Sprintf(pathMods, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// ModPatch updates a mod.
func (c *Default) ModPatch(in *Mod) (*Mod, error) {
	out := &Mod{}

	uri := fmt.Sprintf(pathMod, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// ModDelete deletes a mod.
func (c *Default) ModDelete(id string) error {
	uri := fmt.Sprintf(pathMod, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// ModUserList returns a list of related users for a mod.
func (c *Default) ModUserList(opts ModUserParams) ([]*UserMod, error) {
	var out []*UserMod

	uri := fmt.Sprintf(pathModUser, c.base, opts.Mod)
	err := c.get(uri, &out)

	return out, err
}

// ModUserAppend appends a user to a mod.
func (c *Default) ModUserAppend(opts ModUserParams) error {
	uri := fmt.Sprintf(pathModUser, c.base, opts.Mod)
	err := c.post(uri, opts, nil)

	return err
}

// ModUserPerm updates perms for mod user.
func (c *Default) ModUserPerm(opts ModUserParams) error {
	uri := fmt.Sprintf(pathModUser, c.base, opts.Mod)
	err := c.put(uri, opts, nil)

	return err
}

// ModUserDelete remove a user from a mod.
func (c *Default) ModUserDelete(opts ModUserParams) error {
	uri := fmt.Sprintf(pathModUser, c.base, opts.Mod)
	err := c.delete(uri, opts)

	return err
}

// ModTeamList returns a list of related teams for a mod.
func (c *Default) ModTeamList(opts ModTeamParams) ([]*TeamMod, error) {
	var out []*TeamMod

	uri := fmt.Sprintf(pathModTeam, c.base, opts.Mod)
	err := c.get(uri, &out)

	return out, err
}

// ModTeamAppend appends a team to a mod.
func (c *Default) ModTeamAppend(opts ModTeamParams) error {
	uri := fmt.Sprintf(pathModTeam, c.base, opts.Mod)
	err := c.post(uri, opts, nil)

	return err
}

// ModTeamPerm updates perms for mod team.
func (c *Default) ModTeamPerm(opts ModTeamParams) error {
	uri := fmt.Sprintf(pathModTeam, c.base, opts.Mod)
	err := c.put(uri, opts, nil)

	return err
}

// ModTeamDelete remove a team from a mod.
func (c *Default) ModTeamDelete(opts ModTeamParams) error {
	uri := fmt.Sprintf(pathModTeam, c.base, opts.Mod)
	err := c.delete(uri, opts)

	return err
}

// VersionList returns a list of all versions for a specific mod.
func (c *Default) VersionList(mod string) ([]*Version, error) {
	var out []*Version

	uri := fmt.Sprintf(pathVersions, c.base, mod)
	err := c.get(uri, &out)

	return out, err
}

// VersionGet returns a version for a specific mod.
func (c *Default) VersionGet(mod, id string) (*Version, error) {
	out := &Version{}

	uri := fmt.Sprintf(pathVersion, c.base, mod, id)
	err := c.get(uri, out)

	return out, err
}

// VersionPost creates a version for a specific mod.
func (c *Default) VersionPost(mod string, in *Version) (*Version, error) {
	out := &Version{}

	uri := fmt.Sprintf(pathVersions, c.base, mod)
	err := c.post(uri, in, out)

	return out, err
}

// VersionPatch updates a version for a specific mod.
func (c *Default) VersionPatch(mod string, in *Version) (*Version, error) {
	out := &Version{}

	uri := fmt.Sprintf(pathVersion, c.base, mod, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// VersionDelete deletes a version for a specific mod.
func (c *Default) VersionDelete(mod, id string) error {
	uri := fmt.Sprintf(pathVersion, c.base, mod, id)
	err := c.delete(uri, nil)

	return err
}

// VersionBuildList returns a list of related builds for a version.
func (c *Default) VersionBuildList(opts VersionBuildParams) ([]*BuildVersion, error) {
	var out []*BuildVersion

	uri := fmt.Sprintf(pathVersionBuild, c.base, opts.Mod, opts.Version)
	err := c.get(uri, &out)

	return out, err
}

// VersionBuildAppend appends a build to a version.
func (c *Default) VersionBuildAppend(opts VersionBuildParams) error {
	uri := fmt.Sprintf(pathVersionBuild, c.base, opts.Mod, opts.Version)
	err := c.post(uri, opts, nil)

	return err
}

// VersionBuildDelete remove a build from a version.
func (c *Default) VersionBuildDelete(opts VersionBuildParams) error {
	uri := fmt.Sprintf(pathVersionBuild, c.base, opts.Mod, opts.Version)
	err := c.delete(uri, opts)

	return err
}

// ClientList returns a list of all clients.
func (c *Default) ClientList() ([]*Client, error) {
	var out []*Client

	uri := fmt.Sprintf(pathClients, c.base)
	err := c.get(uri, &out)

	return out, err
}

// ClientGet returns a client.
func (c *Default) ClientGet(id string) (*Client, error) {
	out := &Client{}

	uri := fmt.Sprintf(pathClient, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// ClientPost creates a client.
func (c *Default) ClientPost(in *Client) (*Client, error) {
	out := &Client{}

	uri := fmt.Sprintf(pathClients, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// ClientPatch updates a client.
func (c *Default) ClientPatch(in *Client) (*Client, error) {
	out := &Client{}

	uri := fmt.Sprintf(pathClient, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// ClientDelete deletes a client.
func (c *Default) ClientDelete(id string) error {
	uri := fmt.Sprintf(pathClient, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// ClientPackList returns a list of related packs for a client.
func (c *Default) ClientPackList(opts ClientPackParams) ([]*ClientPack, error) {
	var out []*ClientPack

	uri := fmt.Sprintf(pathClientPack, c.base, opts.Client)
	err := c.get(uri, &out)

	return out, err
}

// ClientPackAppend appends a pack to a client.
func (c *Default) ClientPackAppend(opts ClientPackParams) error {
	uri := fmt.Sprintf(pathClientPack, c.base, opts.Client)
	err := c.post(uri, opts, nil)

	return err
}

// ClientPackDelete remove a pack from a client.
func (c *Default) ClientPackDelete(opts ClientPackParams) error {
	uri := fmt.Sprintf(pathClientPack, c.base, opts.Client)
	err := c.delete(uri, opts)

	return err
}

// UserList returns a list of all users.
func (c *Default) UserList() ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.get(uri, &out)

	return out, err
}

// UserGet returns a user.
func (c *Default) UserGet(id string) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// UserPost creates a user.
func (c *Default) UserPost(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// UserPatch updates a user.
func (c *Default) UserPatch(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// UserDelete deletes a user.
func (c *Default) UserDelete(id string) error {
	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// UserTeamList returns a list of related teams for a user.
func (c *Default) UserTeamList(opts UserTeamParams) ([]*TeamUser, error) {
	var out []*TeamUser

	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserTeamAppend appends a team to a user.
func (c *Default) UserTeamAppend(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.post(uri, opts, nil)

	return err
}

// UserTeamPerm updates perms for user team.
func (c *Default) UserTeamPerm(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.put(uri, opts, nil)

	return err
}

// UserTeamDelete remove a team from a user.
func (c *Default) UserTeamDelete(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// UserModList returns a list of related mods for a user.
func (c *Default) UserModList(opts UserModParams) ([]*UserMod, error) {
	var out []*UserMod

	uri := fmt.Sprintf(pathUserMod, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserModAppend appends a mod to a user.
func (c *Default) UserModAppend(opts UserModParams) error {
	uri := fmt.Sprintf(pathUserMod, c.base, opts.User)
	err := c.post(uri, opts, nil)

	return err
}

// UserModPerm updates perms for user mod.
func (c *Default) UserModPerm(opts UserModParams) error {
	uri := fmt.Sprintf(pathUserMod, c.base, opts.User)
	err := c.put(uri, opts, nil)

	return err
}

// UserModDelete remove a mod from a user.
func (c *Default) UserModDelete(opts UserModParams) error {
	uri := fmt.Sprintf(pathUserMod, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// UserPackList returns a list of related packs for a user.
func (c *Default) UserPackList(opts UserPackParams) ([]*UserPack, error) {
	var out []*UserPack

	uri := fmt.Sprintf(pathUserPack, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserPackAppend appends a pack to a user.
func (c *Default) UserPackAppend(opts UserPackParams) error {
	uri := fmt.Sprintf(pathUserPack, c.base, opts.User)
	err := c.post(uri, opts, nil)

	return err
}

// UserPackPerm updates perms for user pack.
func (c *Default) UserPackPerm(opts UserPackParams) error {
	uri := fmt.Sprintf(pathUserPack, c.base, opts.User)
	err := c.put(uri, opts, nil)

	return err
}

// UserPackDelete remove a pack from a user.
func (c *Default) UserPackDelete(opts UserPackParams) error {
	uri := fmt.Sprintf(pathUserPack, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// TeamList returns a list of all teams.
func (c *Default) TeamList() ([]*Team, error) {
	var out []*Team

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.get(uri, &out)

	return out, err
}

// TeamGet returns a team.
func (c *Default) TeamGet(id string) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// TeamPost creates a team.
func (c *Default) TeamPost(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// TeamPatch updates a team.
func (c *Default) TeamPatch(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// TeamDelete deletes a team.
func (c *Default) TeamDelete(id string) error {
	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// TeamUserList returns a list of related users for a team.
func (c *Default) TeamUserList(opts TeamUserParams) ([]*TeamUser, error) {
	var out []*TeamUser

	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamUserAppend appends a user to a team.
func (c *Default) TeamUserAppend(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.post(uri, opts, nil)

	return err
}

// TeamUserPerm updates perms for team user.
func (c *Default) TeamUserPerm(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.put(uri, opts, nil)

	return err
}

// TeamUserDelete remove a user from a team.
func (c *Default) TeamUserDelete(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}

// TeamModList returns a list of related mods for a team.
func (c *Default) TeamModList(opts TeamModParams) ([]*TeamMod, error) {
	var out []*TeamMod

	uri := fmt.Sprintf(pathTeamMod, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamModAppend appends a mod to a team.
func (c *Default) TeamModAppend(opts TeamModParams) error {
	uri := fmt.Sprintf(pathTeamMod, c.base, opts.Team)
	err := c.post(uri, opts, nil)

	return err
}

// TeamModPerm updates perms for team mod.
func (c *Default) TeamModPerm(opts TeamModParams) error {
	uri := fmt.Sprintf(pathTeamMod, c.base, opts.Team)
	err := c.put(uri, opts, nil)

	return err
}

// TeamModDelete remove a mod from a team.
func (c *Default) TeamModDelete(opts TeamModParams) error {
	uri := fmt.Sprintf(pathTeamMod, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}

// TeamPackList returns a list of related packs for a team.
func (c *Default) TeamPackList(opts TeamPackParams) ([]*TeamPack, error) {
	var out []*TeamPack

	uri := fmt.Sprintf(pathTeamPack, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamPackAppend appends a pack to a team.
func (c *Default) TeamPackAppend(opts TeamPackParams) error {
	uri := fmt.Sprintf(pathTeamPack, c.base, opts.Team)
	err := c.post(uri, opts, nil)

	return err
}

// TeamPackPerm updates perms for team pack.
func (c *Default) TeamPackPerm(opts TeamPackParams) error {
	uri := fmt.Sprintf(pathTeamPack, c.base, opts.Team)
	err := c.put(uri, opts, nil)

	return err
}

// TeamPackDelete remove a pack from a team.
func (c *Default) TeamPackDelete(opts TeamPackParams) error {
	uri := fmt.Sprintf(pathTeamPack, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}
