package kleister

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/vincent-petithory/dataurl"
)

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
	Clients       []*Client   `json:"clients,omitempty"`
	Users         []*User     `json:"users,omitempty"`
	UserPacks     []*UserPack `json:"user_packs,omitempty"`
	Teams         []*Team     `json:"teams,omitempty"`
	TeamPacks     []*TeamPack `json:"team_packs,omitempty"`
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
