package kleister

// ForgeBuildParams is used to assign builds to a forge.
type ForgeBuildParams struct {
	Forge string `json:"forge"`
	Pack  string `json:"pack"`
	Build string `json:"build"`
}

// MinecraftBuildParams is used to assign builds to a minecraft.
type MinecraftBuildParams struct {
	Minecraft string `json:"minecraft"`
	Pack      string `json:"pack"`
	Build     string `json:"build"`
}

// PackClientParams is used to assign clients to a pack.
type PackClientParams struct {
	Pack   string `json:"pack"`
	Client string `json:"client"`
}

// ClientPackParams is used to assign packs to a client.
type ClientPackParams struct {
	Client string `json:"client"`
	Pack   string `json:"pack"`
}

// PackUserParams is used to assign users to a pack.
type PackUserParams struct {
	Pack string `json:"pack"`
	User string `json:"user"`
}

// UserPackParams is used to assign packs to a user.
type UserPackParams struct {
	User string `json:"user"`
	Pack string `json:"pack"`
}

// BuildVersionParams is used to assign versions to a build.
type BuildVersionParams struct {
	Pack    string `json:"pack"`
	Build   string `json:"build"`
	Mod     string `json:"mod"`
	Version string `json:"version"`
}

// VersionBuildParams is used to assign builds to a version.
type VersionBuildParams struct {
	Mod     string `json:"mod"`
	Version string `json:"version"`
	Pack    string `json:"pack"`
	Build   string `json:"build"`
}

// ModUserParams is used to assign users to a mod.
type ModUserParams struct {
	Mod  string `json:"mod"`
	User string `json:"user"`
}

// UserModParams is used to assign mods to a user.
type UserModParams struct {
	User string `json:"user"`
	Mod  string `json:"mod"`
}

// TeamUserParams is used to assign users to a team.
type TeamUserParams struct {
	Team string `json:"team"`
	User string `json:"user"`
}

// UserTeamParams is used to assign teams to a user.
type UserTeamParams struct {
	User string `json:"user"`
	Team string `json:"team"`
}

// PackTeamParams is used to assign teams to a pack.
type PackTeamParams struct {
	Pack string `json:"pack"`
	Team string `json:"team"`
}

// TeamPackParams is used to assign packs to a team.
type TeamPackParams struct {
	Team string `json:"team"`
	Pack string `json:"pack"`
}

// ModTeamParams is used to assign teams to a mod.
type ModTeamParams struct {
	Mod  string `json:"mod"`
	Team string `json:"team"`
}

// TeamModParams is used to assign mods to a team.
type TeamModParams struct {
	Team string `json:"team"`
	Mod  string `json:"mod"`
}
