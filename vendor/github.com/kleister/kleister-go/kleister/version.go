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
	Builds    []*Build    `json:"builds,omitempty"`
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
