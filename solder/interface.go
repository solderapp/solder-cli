package solder

// API describes a Solder API client.
type API interface {
	// ProfileGet returns a profile.
	ProfileGet() (*Profile, error)

	// ProfilePatch updates a profile.
	ProfilePatch(*Profile) (*Profile, error)

	// ForgeList returns a list of all Forge versions.
	ForgeList() ([]*Forge, error)

	// ForgeRefresh refreshs the available Forge versions.
	ForgeRefresh() error

	// ForgeBuildList returns a list of related builds for a Forge version.
	ForgeBuildList(string) ([]*Build, error)

	// ForgeBuildAppend appends a Forge version to a build.
	ForgeBuildAppend(string, string) error

	// ForgeBuildAppend removde a Forge version from a build.
	ForgeBuildDelete(string, string) error

	// MinecraftList returns a list of all Minecraft versions.
	MinecraftList() ([]*Minecraft, error)

	// MinecraftRefresh refreshs the available Minecraft versions.
	MinecraftRefresh() error

	// MinecraftBuildList returns a list of related builds for a Minecraft version.
	MinecraftBuildList(string) ([]*Build, error)

	// MinecraftBuildAppend appends a Minecraft version to a build.
	MinecraftBuildAppend(string, string) error

	// MinecraftBuildAppend removde a Minecraft version from a build.
	MinecraftBuildDelete(string, string) error

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
	PackClientList(string) ([]*Client, error)

	// PackClientAppend appends a client to a pack.
	PackClientAppend(string, string) error

	// PackClientDelete removde a client from a pack.
	PackClientDelete(string, string) error

	// BuildList returns a list of all builds for a specific pack.
	BuildList(string) ([]*Build, error)

	// BuildGet returns a build for a specific pack.
	BuildGet(string) (*Build, error)

	// BuildPost creates a build for a specific pack.
	BuildPost(*Build) (*Build, error)

	// BuildPatch updates a build for a specific pack.
	BuildPatch(*Build) (*Build, error)

	// BuildDelete deletes a build for a specific pack.
	BuildDelete(string) error

	// BuildVersionList returns a list of related versions for a build.
	BuildVersionList(string) ([]*Version, error)

	// BuildVersionAppend appends a version to a build.
	BuildVersionAppend(string, string) error

	// BuildVersionDelete remove a version from a build.
	BuildVersionDelete(string, string) error

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

	// VersionList returns a list of all versions for a specific mod.
	VersionList(string) ([]*Version, error)

	// VersionGet returns a version for a specific mod.
	VersionGet(string) (*Version, error)

	// VersionPost creates a version for a specific mod.
	VersionPost(*Version) (*Version, error)

	// VersionPatch updates a version for a specific mod.
	VersionPatch(*Version) (*Version, error)

	// VersionDelete deletes a version for a specific mod.
	VersionDelete(string) error

	// VersionBuildList returns a list of related builds for a version.
	VersionBuildList(string) ([]*Build, error)

	// VersionBuildAppend appends a build to a version.
	VersionBuildAppend(string, string) error

	// VersionBuildDelete remove a build from a version.
	VersionBuildDelete(string, string) error

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
	ClientPackList(string) ([]*Pack, error)

	// ClientPackAppend appends a pack to a client.
	ClientPackAppend(string, string) error

	// ClientPackDelete remove a pack from a client.
	ClientPackDelete(string, string) error

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
}
