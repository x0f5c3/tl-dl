//go:build !linux

package products

import "path/filepath"

func (d *DownloadedPackage) Save(dir string) error {
	return d.Data.Save(filepath.Join(dir, d.Data.FileName))
}
