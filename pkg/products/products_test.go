package products

import (
	"crypto/sha256"
	"fmt"
	"github.com/x0f5c3/manic-go/pkg/downloader"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestDownload(t *testing.T) {
	box, err := GetToolbox()
	if err != nil {
		t.Error(err)
	}
	rel, err := box.LatestRelease()
	if err != nil {
		t.Error(err)
	}
	o, err := rel.GetOS()
	if err != nil {
		t.Error(err)
	}
	shaResp, err := http.Get(o.ChecksumLink)
	if err != nil {
		t.Error(err)
	}
	shaB, err := ioutil.ReadAll(shaResp.Body)
	if err != nil {
		t.Error(err)
	}
	shaString := strings.TrimSpace(strings.Split(string(shaB), " ")[0])
	size := int(o.Size)
	dl, err := downloader.New(o.Link, shaString, http.DefaultClient, &size)
	if err != nil {
		t.Error(err)
	}
	err = dl.Download(5, 4, true)
	if err != nil {
		t.Error(err)
	}
	err = dl.GetFilename()
	if err != nil {
		t.Error(err)
	}
	saveFile, err := ioutil.TempFile("", dl.FileName)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err := saveFile.Close()
		if err != nil {
			t.Error(err)
		}
		err = os.Remove(saveFile.Name())
		if err != nil {
			t.Error(err)
		}
	}()
	err = dl.Save(saveFile.Name())
	if err != nil {
		t.Error(err)
	}
	rb, err := ioutil.ReadAll(saveFile)
	if err != nil {
		t.Error(err)
	}
	sum := fmt.Sprintf("%x", sha256.Sum256(rb))
	if strings.ToLower(sum) != strings.ToLower(shaString) {
		t.Errorf("Expected %s\nGot %s", shaString, sum)
	}

}
