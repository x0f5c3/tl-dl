package products

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/pterm/pterm"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func (d *DownloadedPackage) Unpack(dir string) error {
	g, err := gzip.NewReader(bytes.NewBuffer(*d.Data.Data))
	if err != nil {
		return err
	}
	r := tar.NewReader(g)
	for {
		f, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		t := f.FileInfo()
		if t.IsDir() {
			continue
		}
		p, err := NewDefaultProgReader(r, int(f.Size)).Start()
		if err != nil {
			return err
		}
		b, err := ioutil.ReadAll(p)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(dir, strings.Split(f.Name, "/")[1]), b, 0777)
		if err != nil {
			return err
		}
		err = p.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}

type ProgReader struct {
	r io.Reader
	p *pterm.ProgressbarPrinter
}

func NewDefaultProgReader(r io.Reader, total int) *ProgReader {
	return &ProgReader{
		r: r,
		p: pterm.DefaultProgressbar.WithTotal(total).WithTitle("Unpacking"),
	}
}

func (p2 *ProgReader) Start() (*ProgReader, error) {
	var err error
	p2.p, err = p2.p.Start()
	if err != nil {
		return nil, err
	}
	return p2, nil
}

func (p2 *ProgReader) Stop() error {
	_, err := p2.p.Stop()
	return err
}

func (p2 ProgReader) Read(p []byte) (n int, err error) {
	n, err = p2.r.Read(p)
	if err != nil {
		return
	}
	p2.p.Add(n)
	return
}
