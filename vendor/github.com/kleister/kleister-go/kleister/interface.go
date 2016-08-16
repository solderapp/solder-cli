package kleister

import (
	"net/http"
)

//go:generate mockery -all -case=underscore

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
	PackClientList(PackClientParams) ([]*Client, error)

	// PackClientAppend appends a client to a pack.
	PackClientAppend(PackClientParams) error

	// PackClientDelete remove a client from a pack.
	PackClientDelete(PackClientParams) error

	// PackUserList returns a list of related users for a pack.
	PackUserList(PackUserParams) ([]*User, error)

	// PackUserAppend appends a user to a pack.
	PackUserAppend(PackUserParams) error

	// PackUserDelete remove a user from a pack.
	PackUserDelete(PackUserParams) error

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
	BuildVersionList(BuildVersionParams) ([]*Version, error)

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
	ModUserList(ModUserParams) ([]*User, error)

	// ModUserAppend appends a user to a mod.
	ModUserAppend(ModUserParams) error

	// ModUserDelete remove a user from a mod.
	ModUserDelete(ModUserParams) error

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
	VersionBuildList(VersionBuildParams) ([]*Build, error)

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
	ClientPackList(ClientPackParams) ([]*Pack, error)

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
	UserModList(UserModParams) ([]*Mod, error)

	// UserModAppend appends a mod to a user.
	UserModAppend(UserModParams) error

	// UserModDelete remove a mod from a user.
	UserModDelete(UserModParams) error

	// UserPackList returns a list of related packs for a user.
	UserPackList(UserPackParams) ([]*Pack, error)

	// UserPackAppend appends a pack to a user.
	UserPackAppend(UserPackParams) error

	// UserPackDelete remove a pack from a user.
	UserPackDelete(UserPackParams) error

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
}
