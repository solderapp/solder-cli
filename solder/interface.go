package solder

// API describes a Solder API client.
type API interface {
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
	ForgeBuildList(string) ([]*Build, error)

	// ForgeBuildAppend appends a Forge version to a build.
	ForgeBuildAppend(string, string) error

	// ForgeBuildAppend remove a Forge version from a build.
	ForgeBuildDelete(string, string) error

	// MinecraftList returns a list of all Minecraft versions.
	MinecraftList() ([]*Minecraft, error)

	// MinecraftGet returns a Minecraft.
	MinecraftGet(string) (*Minecraft, error)

	// MinecraftRefresh refreshs the available Minecraft versions.
	MinecraftRefresh() error

	// MinecraftBuildList returns a list of related builds for a Minecraft version.
	MinecraftBuildList(string) ([]*Build, error)

	// MinecraftBuildAppend appends a Minecraft version to a build.
	MinecraftBuildAppend(string, string) error

	// MinecraftBuildAppend remove a Minecraft version from a build.
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

	// PackClientDelete remove a client from a pack.
	PackClientDelete(string, string) error

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
	BuildVersionList(string, string) ([]*Version, error)

	// BuildVersionAppend appends a version to a build.
	BuildVersionAppend(string, string, string) error

	// BuildVersionDelete remove a version from a build.
	BuildVersionDelete(string, string, string) error

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
	ModUserList(string) ([]*User, error)

	// ModUserAppend appends a user to a mod.
	ModUserAppend(string, string) error

	// ModUserDelete remove a user from a mod.
	ModUserDelete(string, string) error

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
	VersionBuildList(string, string) ([]*Build, error)

	// VersionBuildAppend appends a build to a version.
	VersionBuildAppend(string, string, string) error

	// VersionBuildDelete remove a build from a version.
	VersionBuildDelete(string, string, string) error

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
	UserModList(string) ([]*Mod, error)

	// UserModAppend appends a mod to a user.
	UserModAppend(string, string) error

	// UserModDelete remove a mod from a user.
	UserModDelete(string, string) error

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
