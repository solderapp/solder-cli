package mocks

import "github.com/solderapp/solder-cli/solder"
import "github.com/stretchr/testify/mock"

// API describes a Solder API client.
type API struct {
	mock.Mock
}

// ProfileGet provides a mock function with given fields:
func (_m *API) ProfileGet() (*solder.Profile, error) {
	ret := _m.Called()

	var r0 *solder.Profile
	if rf, ok := ret.Get(0).(func() *solder.Profile); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProfilePatch provides a mock function with given fields: _a0
func (_m *API) ProfilePatch(_a0 *solder.Profile) (*solder.Profile, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Profile
	if rf, ok := ret.Get(0).(func(*solder.Profile) *solder.Profile); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.Profile) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForgeList provides a mock function with given fields:
func (_m *API) ForgeList() ([]*solder.Forge, error) {
	ret := _m.Called()

	var r0 []*solder.Forge
	if rf, ok := ret.Get(0).(func() []*solder.Forge); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Forge)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForgeGet provides a mock function with given fields: _a0
func (_m *API) ForgeGet(_a0 string) (*solder.Forge, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Forge
	if rf, ok := ret.Get(0).(func(string) *solder.Forge); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Forge)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForgeRefresh provides a mock function with given fields:
func (_m *API) ForgeRefresh() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForgeBuildList provides a mock function with given fields: _a0
func (_m *API) ForgeBuildList(_a0 string) ([]*solder.Build, error) {
	ret := _m.Called(_a0)

	var r0 []*solder.Build
	if rf, ok := ret.Get(0).(func(string) []*solder.Build); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForgeBuildAppend provides a mock function with given fields: _a0, _a1
func (_m *API) ForgeBuildAppend(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForgeBuildDelete provides a mock function with given fields: _a0, _a1
func (_m *API) ForgeBuildDelete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MinecraftList provides a mock function with given fields:
func (_m *API) MinecraftList() ([]*solder.Minecraft, error) {
	ret := _m.Called()

	var r0 []*solder.Minecraft
	if rf, ok := ret.Get(0).(func() []*solder.Minecraft); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Minecraft)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MinecraftGet provides a mock function with given fields: _a0
func (_m *API) MinecraftGet(_a0 string) (*solder.Minecraft, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Minecraft
	if rf, ok := ret.Get(0).(func(string) *solder.Minecraft); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Minecraft)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MinecraftRefresh provides a mock function with given fields:
func (_m *API) MinecraftRefresh() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MinecraftBuildList provides a mock function with given fields: _a0
func (_m *API) MinecraftBuildList(_a0 string) ([]*solder.Build, error) {
	ret := _m.Called(_a0)

	var r0 []*solder.Build
	if rf, ok := ret.Get(0).(func(string) []*solder.Build); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MinecraftBuildAppend provides a mock function with given fields: _a0, _a1
func (_m *API) MinecraftBuildAppend(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MinecraftBuildDelete provides a mock function with given fields: _a0, _a1
func (_m *API) MinecraftBuildDelete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PackList provides a mock function with given fields:
func (_m *API) PackList() ([]*solder.Pack, error) {
	ret := _m.Called()

	var r0 []*solder.Pack
	if rf, ok := ret.Get(0).(func() []*solder.Pack); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Pack)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PackGet provides a mock function with given fields: _a0
func (_m *API) PackGet(_a0 string) (*solder.Pack, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Pack
	if rf, ok := ret.Get(0).(func(string) *solder.Pack); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Pack)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PackPost provides a mock function with given fields: _a0
func (_m *API) PackPost(_a0 *solder.Pack) (*solder.Pack, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Pack
	if rf, ok := ret.Get(0).(func(*solder.Pack) *solder.Pack); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Pack)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.Pack) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PackPatch provides a mock function with given fields: _a0
func (_m *API) PackPatch(_a0 *solder.Pack) (*solder.Pack, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Pack
	if rf, ok := ret.Get(0).(func(*solder.Pack) *solder.Pack); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Pack)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.Pack) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PackDelete provides a mock function with given fields: _a0
func (_m *API) PackDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PackClientList provides a mock function with given fields: _a0
func (_m *API) PackClientList(_a0 string) ([]*solder.Client, error) {
	ret := _m.Called(_a0)

	var r0 []*solder.Client
	if rf, ok := ret.Get(0).(func(string) []*solder.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PackClientAppend provides a mock function with given fields: _a0, _a1
func (_m *API) PackClientAppend(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PackClientDelete provides a mock function with given fields: _a0, _a1
func (_m *API) PackClientDelete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BuildList provides a mock function with given fields: _a0
func (_m *API) BuildList(_a0 string) ([]*solder.Build, error) {
	ret := _m.Called(_a0)

	var r0 []*solder.Build
	if rf, ok := ret.Get(0).(func(string) []*solder.Build); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildGet provides a mock function with given fields: _a0, _a1
func (_m *API) BuildGet(_a0 string, _a1 string) (*solder.Build, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *solder.Build
	if rf, ok := ret.Get(0).(func(string, string) *solder.Build); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildPost provides a mock function with given fields: _a0, _a1
func (_m *API) BuildPost(_a0 string, _a1 *solder.Build) (*solder.Build, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *solder.Build
	if rf, ok := ret.Get(0).(func(string, *solder.Build) *solder.Build); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *solder.Build) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildPatch provides a mock function with given fields: _a0, _a1
func (_m *API) BuildPatch(_a0 string, _a1 *solder.Build) (*solder.Build, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *solder.Build
	if rf, ok := ret.Get(0).(func(string, *solder.Build) *solder.Build); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *solder.Build) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildDelete provides a mock function with given fields: _a0, _a1
func (_m *API) BuildDelete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BuildVersionList provides a mock function with given fields: _a0, _a1
func (_m *API) BuildVersionList(_a0 string, _a1 string) ([]*solder.Version, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*solder.Version
	if rf, ok := ret.Get(0).(func(string, string) []*solder.Version); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildVersionAppend provides a mock function with given fields: _a0, _a1, _a2
func (_m *API) BuildVersionAppend(_a0 string, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BuildVersionDelete provides a mock function with given fields: _a0, _a1, _a2
func (_m *API) BuildVersionDelete(_a0 string, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ModList provides a mock function with given fields:
func (_m *API) ModList() ([]*solder.Mod, error) {
	ret := _m.Called()

	var r0 []*solder.Mod
	if rf, ok := ret.Get(0).(func() []*solder.Mod); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Mod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModGet provides a mock function with given fields: _a0
func (_m *API) ModGet(_a0 string) (*solder.Mod, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Mod
	if rf, ok := ret.Get(0).(func(string) *solder.Mod); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Mod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModPost provides a mock function with given fields: _a0
func (_m *API) ModPost(_a0 *solder.Mod) (*solder.Mod, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Mod
	if rf, ok := ret.Get(0).(func(*solder.Mod) *solder.Mod); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Mod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.Mod) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModPatch provides a mock function with given fields: _a0
func (_m *API) ModPatch(_a0 *solder.Mod) (*solder.Mod, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Mod
	if rf, ok := ret.Get(0).(func(*solder.Mod) *solder.Mod); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Mod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.Mod) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModDelete provides a mock function with given fields: _a0
func (_m *API) ModDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ModUserList provides a mock function with given fields: _a0
func (_m *API) ModUserList(_a0 string) ([]*solder.User, error) {
	ret := _m.Called(_a0)

	var r0 []*solder.User
	if rf, ok := ret.Get(0).(func(string) []*solder.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModUserAppend provides a mock function with given fields: _a0, _a1
func (_m *API) ModUserAppend(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ModUserDelete provides a mock function with given fields: _a0, _a1
func (_m *API) ModUserDelete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VersionList provides a mock function with given fields: _a0
func (_m *API) VersionList(_a0 string) ([]*solder.Version, error) {
	ret := _m.Called(_a0)

	var r0 []*solder.Version
	if rf, ok := ret.Get(0).(func(string) []*solder.Version); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VersionGet provides a mock function with given fields: _a0, _a1
func (_m *API) VersionGet(_a0 string, _a1 string) (*solder.Version, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *solder.Version
	if rf, ok := ret.Get(0).(func(string, string) *solder.Version); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VersionPost provides a mock function with given fields: _a0, _a1
func (_m *API) VersionPost(_a0 string, _a1 *solder.Version) (*solder.Version, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *solder.Version
	if rf, ok := ret.Get(0).(func(string, *solder.Version) *solder.Version); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *solder.Version) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VersionPatch provides a mock function with given fields: _a0, _a1
func (_m *API) VersionPatch(_a0 string, _a1 *solder.Version) (*solder.Version, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *solder.Version
	if rf, ok := ret.Get(0).(func(string, *solder.Version) *solder.Version); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *solder.Version) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VersionDelete provides a mock function with given fields: _a0, _a1
func (_m *API) VersionDelete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VersionBuildList provides a mock function with given fields: _a0, _a1
func (_m *API) VersionBuildList(_a0 string, _a1 string) ([]*solder.Build, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*solder.Build
	if rf, ok := ret.Get(0).(func(string, string) []*solder.Build); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VersionBuildAppend provides a mock function with given fields: _a0, _a1, _a2
func (_m *API) VersionBuildAppend(_a0 string, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VersionBuildDelete provides a mock function with given fields: _a0, _a1, _a2
func (_m *API) VersionBuildDelete(_a0 string, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientList provides a mock function with given fields:
func (_m *API) ClientList() ([]*solder.Client, error) {
	ret := _m.Called()

	var r0 []*solder.Client
	if rf, ok := ret.Get(0).(func() []*solder.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientGet provides a mock function with given fields: _a0
func (_m *API) ClientGet(_a0 string) (*solder.Client, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Client
	if rf, ok := ret.Get(0).(func(string) *solder.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientPost provides a mock function with given fields: _a0
func (_m *API) ClientPost(_a0 *solder.Client) (*solder.Client, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Client
	if rf, ok := ret.Get(0).(func(*solder.Client) *solder.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientPatch provides a mock function with given fields: _a0
func (_m *API) ClientPatch(_a0 *solder.Client) (*solder.Client, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Client
	if rf, ok := ret.Get(0).(func(*solder.Client) *solder.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientDelete provides a mock function with given fields: _a0
func (_m *API) ClientDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientPackList provides a mock function with given fields: _a0
func (_m *API) ClientPackList(_a0 string) ([]*solder.Pack, error) {
	ret := _m.Called(_a0)

	var r0 []*solder.Pack
	if rf, ok := ret.Get(0).(func(string) []*solder.Pack); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Pack)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientPackAppend provides a mock function with given fields: _a0, _a1
func (_m *API) ClientPackAppend(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientPackDelete provides a mock function with given fields: _a0, _a1
func (_m *API) ClientPackDelete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserList provides a mock function with given fields:
func (_m *API) UserList() ([]*solder.User, error) {
	ret := _m.Called()

	var r0 []*solder.User
	if rf, ok := ret.Get(0).(func() []*solder.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserGet provides a mock function with given fields: _a0
func (_m *API) UserGet(_a0 string) (*solder.User, error) {
	ret := _m.Called(_a0)

	var r0 *solder.User
	if rf, ok := ret.Get(0).(func(string) *solder.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserPost provides a mock function with given fields: _a0
func (_m *API) UserPost(_a0 *solder.User) (*solder.User, error) {
	ret := _m.Called(_a0)

	var r0 *solder.User
	if rf, ok := ret.Get(0).(func(*solder.User) *solder.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserPatch provides a mock function with given fields: _a0
func (_m *API) UserPatch(_a0 *solder.User) (*solder.User, error) {
	ret := _m.Called(_a0)

	var r0 *solder.User
	if rf, ok := ret.Get(0).(func(*solder.User) *solder.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserDelete provides a mock function with given fields: _a0
func (_m *API) UserDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserModList provides a mock function with given fields: _a0
func (_m *API) UserModList(_a0 string) ([]*solder.Mod, error) {
	ret := _m.Called(_a0)

	var r0 []*solder.Mod
	if rf, ok := ret.Get(0).(func(string) []*solder.Mod); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Mod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserModAppend provides a mock function with given fields: _a0, _a1
func (_m *API) UserModAppend(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserModDelete provides a mock function with given fields: _a0, _a1
func (_m *API) UserModDelete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KeyList provides a mock function with given fields:
func (_m *API) KeyList() ([]*solder.Key, error) {
	ret := _m.Called()

	var r0 []*solder.Key
	if rf, ok := ret.Get(0).(func() []*solder.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*solder.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeyGet provides a mock function with given fields: _a0
func (_m *API) KeyGet(_a0 string) (*solder.Key, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Key
	if rf, ok := ret.Get(0).(func(string) *solder.Key); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeyPost provides a mock function with given fields: _a0
func (_m *API) KeyPost(_a0 *solder.Key) (*solder.Key, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Key
	if rf, ok := ret.Get(0).(func(*solder.Key) *solder.Key); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.Key) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeyPatch provides a mock function with given fields: _a0
func (_m *API) KeyPatch(_a0 *solder.Key) (*solder.Key, error) {
	ret := _m.Called(_a0)

	var r0 *solder.Key
	if rf, ok := ret.Get(0).(func(*solder.Key) *solder.Key); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*solder.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*solder.Key) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeyDelete provides a mock function with given fields: _a0
func (_m *API) KeyDelete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
