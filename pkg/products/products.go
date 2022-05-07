package products

import (
	"encoding/json"
	"errors"
	"github.com/Masterminds/semver"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/x0f5c3/manic-go/pkg/downloader"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

type Products []Product

type Product struct {
	Distributions struct {
		Linux struct {
			Extension string `json:"extension"`
			Name      string `json:"name"`
		} `json:"linux"`
		Mac struct {
			Extension string `json:"extension"`
			Name      string `json:"name"`
		} `json:"mac"`
		MacM1 struct {
			Extension string `json:"extension"`
			Name      string `json:"name"`
		} `json:"macM1"`
		MacUpdate struct {
			Extension string `json:"extension"`
			Name      string `json:"name"`
		} `json:"macUpdate"`
		Windows struct {
			Extension string `json:"extension"`
			Name      string `json:"name"`
		} `json:"windows"`
	} `json:"distributions"`
	Link     string    `json:"link"`
	Name     string    `json:"name"`
	Releases []Release `json:"releases"`
}

type Release struct {
	Build     string `json:"build"`
	Date      string `json:"date"`
	Downloads struct {
		Linux     OS `json:"linux"`
		Mac       OS `json:"mac"`
		MacM1     OS `json:"macM1"`
		MacUpdate OS `json:"macUpdate"`
		Windows   OS `json:"windows"`
	} `json:"downloads"`
	LicenseRequired        interface{} `json:"licenseRequired"`
	MajorVersion           string      `json:"majorVersion"`
	NotesLink              string      `json:"notesLink"`
	Patches                struct{}    `json:"patches"`
	PrintableReleaseType   interface{} `json:"printableReleaseType"`
	Type                   string      `json:"type"`
	UninstallFeedbackLinks interface{} `json:"uninstallFeedbackLinks"`
	Version                string      `json:"version"`
	Whatsnew               string      `json:"whatsnew"`
}

type DownloadedPackage struct {
	Build     *semver.Version
	Date      string
	NotesLink string
	Whatsnew  string
	Data      *downloader.File
	OS
}

type OS struct {
	ChecksumLink string `json:"checksumLink"`
	Link         string `json:"link"`
	Size         int64  `json:"size"`
}

func (o OS) String() string {
	res := cfmt.Sprintf("{{ChecksumLink: %s}}::green\n", o.ChecksumLink)
	res += cfmt.Sprintf("{{Link: %s}}::green\n", o.Link)
	res += cfmt.Sprintf("{{Size: %d}}::green", o.Size)
	return res
}

func GetProducts() (Products, error) {
	resp, err := http.Get("https://data.services.jetbrains.com/products?code=TBA&release.type=eap%2Crc%2Crelease&fields=distributions%2Clink%2Cname%2Creleases")
	if err != nil {
		return nil, err
	}
	var res Products
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetToolbox() (Product, error) {
	prods, err := GetProducts()
	if err != nil {
		return Product{}, err
	}
	if len(prods) == 0 {
		return Product{}, errors.New("no toolbox found")
	}
	first := prods[0]
	return first, nil
}

func (p Product) LatestRelease() (Release, error) {
	if len(p.Releases) == 0 {
		return Release{}, errors.New("no releases found")
	}
	return p.Releases[0], nil
}

func (r Release) GetOS() (OS, error) {
	switch runtime.GOOS {
	case "windows":
		return r.Downloads.Windows, nil
	case "linux":
		return r.Downloads.Linux, nil
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			return r.Downloads.Mac, nil
		case "arm64":
			return r.Downloads.MacM1, nil
		}
	}
	return OS{}, errors.New("unknown OS")
}

func (r Release) Download() (*DownloadedPackage, error) {
	o, err := r.GetOS()
	shaResp, err := http.Get(o.ChecksumLink)
	if err != nil {
		return nil, err
	}
	shaB, err := ioutil.ReadAll(shaResp.Body)
	if err != nil {
		return nil, err
	}
	shaString := strings.TrimSpace(strings.Split(string(shaB), " ")[0])
	size := int(o.Size)
	dl, err := downloader.New(o.Link, shaString, http.DefaultClient, &size)
	if err != nil {
		return nil, err
	}
	err = dl.Download(5, 4, true)
	if err != nil {
		return nil, err
	}
	err = dl.GetFilename()
	if err != nil {
		return nil, err
	}
	vers, err := semver.NewVersion(r.Build)
	if err != nil {
		return nil, err
	}
	return &DownloadedPackage{
		Build:     vers,
		Date:      r.Date,
		NotesLink: r.NotesLink,
		Whatsnew:  r.Whatsnew,
		Data:      dl,
		OS:        o,
	}, nil
}

func DownloadNative() (*DownloadedPackage, error) {
	prod, err := GetToolbox()
	if err != nil {
		return nil, err
	}
	latest, err := prod.LatestRelease()
	if err != nil {
		return nil, err
	}
	return latest.Download()
}

func (d *DownloadedPackage) String() string {
	res := ""
	res += cfmt.Sprintf("{{Build: %s}}::green\n", d.Build)
	res += cfmt.Sprintf("{{Date: %s}}::green\n", d.Date)
	res += cfmt.Sprintf("{{NotesLink: %s}}::green\n", d.NotesLink)
	res += cfmt.Sprintf("{{Whatsnew: %s}}::green\n", d.Whatsnew)
	res += cfmt.Sprintf("{{OS: %s}}::green", d.OS.String())
	return res
}
