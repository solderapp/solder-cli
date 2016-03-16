package solder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/vincent-petithory/dataurl"
)

// Message represents a standard response.
type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Build represents a build API response.
type Build struct {
	ID          int64         `json:"id"`
	Pack        PackName      `json:"pack"`
	PackID      int64         `json:"pack_id"`
	Minecraft   MinecraftName `json:"minecraft"`
	MinecraftID int64         `json:"minecraft_id"`
	Forge       ForgeName     `json:"forge"`
	ForgeID     int64         `json:"forge_id"`
	Slug        string        `json:"slug"`
	Name        string        `json:"name"`
	MinJava     string        `json:"min_java"`
	MinMemory   string        `json:"min_memory"`
	Published   bool          `json:"published"`
	Private     bool          `json:"private"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func (s *Build) String() string {
	return s.Name
}

// BuildName represents the mapped value of a simple build name.
type BuildName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *BuildName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Name string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode build: %v", err)
	}

	*s = BuildName(aux.Name)
	return nil
}

// Client represents a client API response.
type Client struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Value     string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Client) String() string {
	return s.Name
}

// ClientName represents the mapped value of a simple client name.
type ClientName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *ClientName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Name string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode client: %v", err)
	}

	*s = ClientName(aux.Name)
	return nil
}

// Forge represents a forge API response.
type Forge struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Version   string    `json:"version"`
	Minecraft string    `json:"minecraft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Forge) String() string {
	return s.Version
}

// ForgeName represents the mapped value of a simple forge name.
type ForgeName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *ForgeName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Version string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode forge: %v", err)
	}

	*s = ForgeName(aux.Version)
	return nil
}

// Key represents a key API response.
type Key struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Value     string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Key) String() string {
	return s.Name
}

// KeyName represents the mapped value of a simple key name.
type KeyName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *KeyName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Name string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode key: %v", err)
	}

	*s = KeyName(aux.Name)
	return nil
}

// Minecraft represents a minecraft API response.
type Minecraft struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Version   string    `json:"version"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Minecraft) String() string {
	return s.Version
}

// MinecraftName represents the mapped value of a simple Minecraft name.
type MinecraftName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *MinecraftName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Version string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode minecraft: %v", err)
	}

	*s = MinecraftName(aux.Version)
	return nil
}

// Mod represents a mod API response.
type Mod struct {
	ID          int64     `json:"id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Website     string    `json:"website"`
	Donate      string    `json:"donate"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *Mod) String() string {
	return s.Name
}

// ModName represents the mapped value of a simple mod name.
type ModName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *ModName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Name string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode mod: %v", err)
	}

	*s = ModName(aux.Name)
	return nil
}

// Pack represents a pack API response.
type Pack struct {
	ID            int64     `json:"id"`
	Slug          string    `json:"slug"`
	Name          string    `json:"name"`
	Icon          string    `json:"icon"`
	Logo          string    `json:"logo"`
	Background    string    `json:"background"`
	RecommendedID int64     `json:"recommended_id"`
	Recommended   BuildName `json:"recommended"`
	LatestID      int64     `json:"latest_id"`
	Latest        BuildName `json:"latest"`
	Website       string    `json:"website"`
	Published     bool      `json:"published"`
	Private       bool      `json:"private"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (s *Pack) String() string {
	return s.Name
}

func (s *Pack) DownloadIcon(path string) error {
	tmpfile, err := ioutil.TempFile("", "icon")

	if err != nil {
		return fmt.Errorf("Failed to create a temporary icon")
	}

	defer os.Remove(tmpfile.Name())

	resp, err := http.Get(path)

	if err != nil {
		return fmt.Errorf("Failed to download the icon")
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if _, err = tmpfile.WriteString(buf.String()); err != nil {
		return fmt.Errorf("Failed to copy the icon content")
	}

	return s.EncodeIcon(tmpfile.Name())
}

func (s *Pack) EncodeIcon(path string) error {
	file, err := ioutil.ReadFile(
		path,
	)

	if err != nil {
		return fmt.Errorf("Failed to read icon")
	}

	mimeType := http.DetectContentType(
		file,
	)

	data := dataurl.New(
		file,
		mimeType,
	)

	s.Icon = data.String()
	return nil
}

func (s *Pack) DownloadLogo(path string) error {
	tmpfile, err := ioutil.TempFile("", "logo")

	if err != nil {
		return fmt.Errorf("Failed to create a temporary logo")
	}

	defer os.Remove(tmpfile.Name())

	resp, err := http.Get(path)

	if err != nil {
		return fmt.Errorf("Failed to download the logo")
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if _, err = tmpfile.WriteString(buf.String()); err != nil {
		return fmt.Errorf("Failed to copy the logo content")
	}

	return s.EncodeLogo(tmpfile.Name())
}

func (s *Pack) EncodeLogo(path string) error {
	file, err := ioutil.ReadFile(
		path,
	)

	if err != nil {
		return fmt.Errorf("Failed to read logo")
	}

	mimeType := http.DetectContentType(
		file,
	)

	data := dataurl.New(
		file,
		mimeType,
	)

	s.Logo = data.String()
	return nil
}

func (s *Pack) DownloadBackground(path string) error {
	tmpfile, err := ioutil.TempFile("", "background")

	if err != nil {
		return fmt.Errorf("Failed to create a temporary background")
	}

	defer os.Remove(tmpfile.Name())

	resp, err := http.Get(path)

	if err != nil {
		return fmt.Errorf("Failed to download the background")
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if _, err = tmpfile.WriteString(buf.String()); err != nil {
		return fmt.Errorf("Failed to copy the background content")
	}

	return s.EncodeBackground(tmpfile.Name())
}

func (s *Pack) EncodeBackground(path string) error {
	file, err := ioutil.ReadFile(
		path,
	)

	if err != nil {
		return fmt.Errorf("Failed to read background")
	}

	mimeType := http.DetectContentType(
		file,
	)

	data := dataurl.New(
		file,
		mimeType,
	)

	s.Background = data.String()
	return nil
}

// PackName represents the mapped value of a simple pack name.
type PackName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *PackName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Name string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode pack: %v", err)
	}

	*s = PackName(aux.Name)
	return nil
}

// User represents a user API response.
type User struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *User) String() string {
	return s.Username
}

// UserName represents the mapped value of a simple user username.
type UserName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *UserName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Username string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode user: %v", err)
	}

	*s = UserName(aux.Username)
	return nil
}

// Version represents a version API response.
type Version struct {
	ID        int64     `json:"id"`
	Mod       ModName   `json:"mod"`
	ModID     int64     `json:"mod_id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	File      string    `json:"file"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Version) String() string {
	return s.Name
}

func (s *Version) DownloadFile(path string) error {
	tmpfile, err := ioutil.TempFile("", "version")

	if err != nil {
		return fmt.Errorf("Failed to create a temporary file")
	}

	defer os.Remove(tmpfile.Name())

	resp, err := http.Get(path)

	if err != nil {
		return fmt.Errorf("Failed to download the file")
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if _, err = tmpfile.WriteString(buf.String()); err != nil {
		return fmt.Errorf("Failed to copy the file content")
	}

	return s.EncodeFile(tmpfile.Name())
}

func (s *Version) EncodeFile(path string) error {
	file, err := ioutil.ReadFile(
		path,
	)

	if err != nil {
		return fmt.Errorf("Failed to read file")
	}

	mimeType := http.DetectContentType(
		file,
	)

	data := dataurl.New(
		file,
		mimeType,
	)

	s.File = data.String()
	return nil
}

// VersionName represents the mapped value of a simple version name.
type VersionName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *VersionName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Name string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode version: %v", err)
	}

	*s = VersionName(aux.Name)
	return nil
}

// Profile represents a profile API response.
type Profile struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Profile) String() string {
	return s.Username
}

// ProfileName represents the mapped value of a simple profile name.
type ProfileName string

// UnmarshalJSON just maps the nested struct into a simple string.
func (s *ProfileName) UnmarshalJSON(data []byte) error {
	r := bytes.NewReader(data)

	var aux struct {
		Username string
	}

	if err := json.NewDecoder(r).Decode(&aux); err != nil {
		return fmt.Errorf("Failed to decode profile: %v", err)
	}

	*s = ProfileName(aux.Username)
	return nil
}
