package solder

import (
	"bytes"
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

// Attachment represents a attachment API response.
type Attachment struct {
	URL       string    `json:"url,omitempty"`
	MD5       string    `json:"md5,omitempty"`
	Upload    string    `json:"upload,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Attachment) String() string {
	return s.URL
}

// Build represents a build API response.
type Build struct {
	ID          int64      `json:"id"`
	Pack        *Pack      `json:"pack,omitempty"`
	PackID      int64      `json:"pack_id"`
	Minecraft   *Minecraft `json:"minecraft,omitempty"`
	MinecraftID int64      `json:"minecraft_id"`
	Forge       *Forge     `json:"forge,omitempty"`
	ForgeID     int64      `json:"forge_id"`
	Slug        string     `json:"slug"`
	Name        string     `json:"name"`
	MinJava     string     `json:"min_java"`
	MinMemory   string     `json:"min_memory"`
	Published   bool       `json:"published"`
	Private     bool       `json:"private"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (s *Build) String() string {
	return s.Name
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

// Pack represents a pack API response.
type Pack struct {
	ID            int64       `json:"id"`
	Slug          string      `json:"slug"`
	Name          string      `json:"name"`
	Icon          *Attachment `json:"icon,omitempty"`
	Logo          *Attachment `json:"logo,omitempty"`
	Background    *Attachment `json:"background,omitempty"`
	RecommendedID int64       `json:"recommended_id"`
	Recommended   *Build      `json:"recommended,omitempty"`
	LatestID      int64       `json:"latest_id"`
	Latest        *Build      `json:"latest,omitempty"`
	Website       string      `json:"website"`
	Published     bool        `json:"published"`
	Private       bool        `json:"private"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

func (s *Pack) String() string {
	return s.Name
}

// DownloadIcon is responsible for downloading an icon from remote.
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

// EncodeIcon is responsible for encoding an icon to a dataurl.
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

	s.Icon = &Attachment{
		Upload: data.String(),
	}

	return nil
}

// DownloadLogo is responsible for downloading a logo from remote.
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

// EncodeLogo is responsible for encoding a logo to a dataurl.
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

	s.Logo = &Attachment{
		Upload: data.String(),
	}

	return nil
}

// DownloadBackground is responsible for downloading a background from remote.
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

// EncodeBackground is responsible for encoding a background to a dataurl.
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

	s.Background = &Attachment{
		Upload: data.String(),
	}

	return nil
}

// User represents a user API response.
type User struct {
	ID         int64     `json:"id"`
	Slug       string    `json:"slug"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Packs      []*Pack   `json:"packs,omitempty"`
	Permission struct {
		DisplayUsers   bool `json:"display_users"`
		ChangeUsers    bool `json:"change_users"`
		DeleteUsers    bool `json:"delete_users"`
		DisplayKeys    bool `json:"display_keys"`
		ChangeKeys     bool `json:"change_keys"`
		DeleteKeys     bool `json:"delete_keys"`
		DisplayClients bool `json:"display_clients"`
		ChangeClients  bool `json:"change_clients"`
		DeleteClients  bool `json:"delete_clients"`
		DisplayPacks   bool `json:"display_packs"`
		ChangePacks    bool `json:"change_packs"`
		DeletePacks    bool `json:"delete_packs"`
		DisplayMods    bool `json:"display_mods"`
		ChangeMods     bool `json:"change_mods"`
		DeleteMods     bool `json:"delete_mods"`
	}
}

func (s *User) String() string {
	return s.Username
}

// Version represents a version API response.
type Version struct {
	ID        int64       `json:"id"`
	Mod       *Mod        `json:"mod,omitempty"`
	ModID     int64       `json:"mod_id"`
	Slug      string      `json:"slug"`
	Name      string      `json:"name"`
	File      *Attachment `json:"file"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (s *Version) String() string {
	return s.Name
}

// DownloadFile is responsible for downloading a file from remote.
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

// EncodeFile is responsible for encoding a file to a dataurl.
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

	s.File = &Attachment{
		Upload: data.String(),
	}

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
