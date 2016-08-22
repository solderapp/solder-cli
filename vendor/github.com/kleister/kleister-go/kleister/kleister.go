package kleister

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jackspirou/syscerts"
	"golang.org/x/oauth2"
)

//go:generate mockery -all -case=underscore

const (
	pathAuthLogin      = "%s/api/auth/login"
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

// DefaultClient implements the client interface.
type DefaultClient struct {
	client *http.Client
	base   string
	token  string
}

// NewClient returns a client for the specified URL.
func NewClient(uri string) ClientAPI {
	return &DefaultClient{
		client: http.DefaultClient,
		base:   uri,
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

	if trans, ok := auther.Transport.(*oauth2.Transport); ok {
		trans.Base = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				RootCAs: syscerts.SystemRootsPool(),
			},
		}
	}

	return &DefaultClient{
		client: auther,
		base:   uri,
		token:  token,
	}
}

// IsAuthenticated checks if we already provided an authentication
// token for our client requests. If it returns false you can update
// the client after fetching a valid token.
func (c *DefaultClient) IsAuthenticated() bool {
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

	if resp.StatusCode == http.StatusUnauthorized {
		return false
	}

	return true
}

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the Kleister API.
func (c *DefaultClient) SetClient(client *http.Client) {
	c.client = client
}

// AuthLogin signs in based on credentials and returns a token.
func (c *DefaultClient) AuthLogin(username, password string) (*Token, error) {
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
func (c *DefaultClient) PackClientList(opts PackClientParams) ([]*ClientPack, error) {
	var out []*ClientPack

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
func (c *DefaultClient) PackUserList(opts PackUserParams) ([]*UserPack, error) {
	var out []*UserPack

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

// PackUserPerm updates perms for pack team.
func (c *DefaultClient) PackUserPerm(opts PackUserParams) error {
	uri := fmt.Sprintf(pathPackUser, c.base, opts.Pack)
	err := c.post(uri, opts, nil)

	return err
}

// PackUserDelete remove a user from a pack.
func (c *DefaultClient) PackUserDelete(opts PackUserParams) error {
	uri := fmt.Sprintf(pathPackUser, c.base, opts.Pack)
	err := c.delete(uri, opts)

	return err
}

// PackTeamList returns a list of related teams for a pack.
func (c *DefaultClient) PackTeamList(opts PackTeamParams) ([]*TeamPack, error) {
	var out []*TeamPack

	uri := fmt.Sprintf(pathPackTeam, c.base, opts.Pack)
	err := c.get(uri, &out)

	return out, err
}

// PackTeamAppend appends a team to a pack.
func (c *DefaultClient) PackTeamAppend(opts PackTeamParams) error {
	uri := fmt.Sprintf(pathPackTeam, c.base, opts.Pack)
	err := c.patch(uri, opts, nil)

	return err
}

// PackTeamPerm updates perms for pack team.
func (c *DefaultClient) PackTeamPerm(opts PackTeamParams) error {
	uri := fmt.Sprintf(pathPackTeam, c.base, opts.Pack)
	err := c.post(uri, opts, nil)

	return err
}

// PackTeamDelete remove a team from a pack.
func (c *DefaultClient) PackTeamDelete(opts PackTeamParams) error {
	uri := fmt.Sprintf(pathPackTeam, c.base, opts.Pack)
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
func (c *DefaultClient) BuildVersionList(opts BuildVersionParams) ([]*BuildVersion, error) {
	var out []*BuildVersion

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
func (c *DefaultClient) ModUserList(opts ModUserParams) ([]*UserMod, error) {
	var out []*UserMod

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

// ModUserPerm updates perms for mod user.
func (c *DefaultClient) ModUserPerm(opts ModUserParams) error {
	uri := fmt.Sprintf(pathModUser, c.base, opts.Mod)
	err := c.post(uri, opts, nil)

	return err
}

// ModUserDelete remove a user from a mod.
func (c *DefaultClient) ModUserDelete(opts ModUserParams) error {
	uri := fmt.Sprintf(pathModUser, c.base, opts.Mod)
	err := c.delete(uri, opts)

	return err
}

// ModTeamList returns a list of related teams for a mod.
func (c *DefaultClient) ModTeamList(opts ModTeamParams) ([]*TeamMod, error) {
	var out []*TeamMod

	uri := fmt.Sprintf(pathModTeam, c.base, opts.Mod)
	err := c.get(uri, &out)

	return out, err
}

// ModTeamAppend appends a team to a mod.
func (c *DefaultClient) ModTeamAppend(opts ModTeamParams) error {
	uri := fmt.Sprintf(pathModTeam, c.base, opts.Mod)
	err := c.patch(uri, opts, nil)

	return err
}

// ModTeamPerm updates perms for mod team.
func (c *DefaultClient) ModTeamPerm(opts ModTeamParams) error {
	uri := fmt.Sprintf(pathModTeam, c.base, opts.Mod)
	err := c.post(uri, opts, nil)

	return err
}

// ModTeamDelete remove a team from a mod.
func (c *DefaultClient) ModTeamDelete(opts ModTeamParams) error {
	uri := fmt.Sprintf(pathModTeam, c.base, opts.Mod)
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
func (c *DefaultClient) VersionBuildList(opts VersionBuildParams) ([]*BuildVersion, error) {
	var out []*BuildVersion

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
func (c *DefaultClient) ClientPackList(opts ClientPackParams) ([]*ClientPack, error) {
	var out []*ClientPack

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

// UserTeamList returns a list of related teams for a user.
func (c *DefaultClient) UserTeamList(opts UserTeamParams) ([]*TeamUser, error) {
	var out []*TeamUser

	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserTeamAppend appends a team to a user.
func (c *DefaultClient) UserTeamAppend(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.patch(uri, opts, nil)

	return err
}

// UserTeamPerm updates perms for user team.
func (c *DefaultClient) UserTeamPerm(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.post(uri, opts, nil)

	return err
}

// UserTeamDelete remove a team from a user.
func (c *DefaultClient) UserTeamDelete(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// UserModList returns a list of related mods for a user.
func (c *DefaultClient) UserModList(opts UserModParams) ([]*UserMod, error) {
	var out []*UserMod

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

// UserModPerm updates perms for user mod.
func (c *DefaultClient) UserModPerm(opts UserModParams) error {
	uri := fmt.Sprintf(pathUserMod, c.base, opts.User)
	err := c.post(uri, opts, nil)

	return err
}

// UserModDelete remove a mod from a user.
func (c *DefaultClient) UserModDelete(opts UserModParams) error {
	uri := fmt.Sprintf(pathUserMod, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// UserPackList returns a list of related packs for a user.
func (c *DefaultClient) UserPackList(opts UserPackParams) ([]*UserPack, error) {
	var out []*UserPack

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

// UserPackPerm updates perms for user pack.
func (c *DefaultClient) UserPackPerm(opts UserPackParams) error {
	uri := fmt.Sprintf(pathUserPack, c.base, opts.User)
	err := c.post(uri, opts, nil)

	return err
}

// UserPackDelete remove a pack from a user.
func (c *DefaultClient) UserPackDelete(opts UserPackParams) error {
	uri := fmt.Sprintf(pathUserPack, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// TeamList returns a list of all teams.
func (c *DefaultClient) TeamList() ([]*Team, error) {
	var out []*Team

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.get(uri, &out)

	return out, err
}

// TeamGet returns a team.
func (c *DefaultClient) TeamGet(id string) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// TeamPost creates a team.
func (c *DefaultClient) TeamPost(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// TeamPatch updates a team.
func (c *DefaultClient) TeamPatch(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// TeamDelete deletes a team.
func (c *DefaultClient) TeamDelete(id string) error {
	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// TeamUserList returns a list of related users for a team.
func (c *DefaultClient) TeamUserList(opts TeamUserParams) ([]*TeamUser, error) {
	var out []*TeamUser

	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamUserAppend appends a user to a team.
func (c *DefaultClient) TeamUserAppend(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.patch(uri, opts, nil)

	return err
}

// TeamUserPerm updates perms for team user.
func (c *DefaultClient) TeamUserPerm(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.post(uri, opts, nil)

	return err
}

// TeamUserDelete remove a user from a team.
func (c *DefaultClient) TeamUserDelete(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}

// TeamModList returns a list of related mods for a team.
func (c *DefaultClient) TeamModList(opts TeamModParams) ([]*TeamMod, error) {
	var out []*TeamMod

	uri := fmt.Sprintf(pathTeamMod, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamModAppend appends a mod to a team.
func (c *DefaultClient) TeamModAppend(opts TeamModParams) error {
	uri := fmt.Sprintf(pathTeamMod, c.base, opts.Team)
	err := c.patch(uri, opts, nil)

	return err
}

// TeamModPerm updates perms for team mod.
func (c *DefaultClient) TeamModPerm(opts TeamModParams) error {
	uri := fmt.Sprintf(pathTeamMod, c.base, opts.Team)
	err := c.post(uri, opts, nil)

	return err
}

// TeamModDelete remove a mod from a team.
func (c *DefaultClient) TeamModDelete(opts TeamModParams) error {
	uri := fmt.Sprintf(pathTeamMod, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}

// TeamPackList returns a list of related packs for a team.
func (c *DefaultClient) TeamPackList(opts TeamPackParams) ([]*TeamPack, error) {
	var out []*TeamPack

	uri := fmt.Sprintf(pathTeamPack, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamPackAppend appends a pack to a team.
func (c *DefaultClient) TeamPackAppend(opts TeamPackParams) error {
	uri := fmt.Sprintf(pathTeamPack, c.base, opts.Team)
	err := c.patch(uri, opts, nil)

	return err
}

// TeamPackPerm updates perms for team pack.
func (c *DefaultClient) TeamPackPerm(opts TeamPackParams) error {
	uri := fmt.Sprintf(pathTeamPack, c.base, opts.Team)
	err := c.post(uri, opts, nil)

	return err
}

// TeamPackDelete remove a pack from a team.
func (c *DefaultClient) TeamPackDelete(opts TeamPackParams) error {
	uri := fmt.Sprintf(pathTeamPack, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}
