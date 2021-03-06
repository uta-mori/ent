// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x51\x4f\xdb\x30\x14\x85\x9f\xe3\x5f\x71\x16\x31\x91\xb0\xe2\x02\x6f\x9b\xd4\x07\x04\x9d\xd4\x69\x83\x49\x45\xda\x03\x43\xc8\x75\x6e\x5a\x8b\xd4\xce\xae\x5d\xb4\xca\xca\x7f\x9f\xec\xb4\x6c\x7b\x4b\x7c\xbe\x7b\xce\xb9\x76\x8c\xd3\x33\x71\xe3\xfa\x3d\x9b\xf5\x26\xe0\xea\xe2\xf2\xe3\x79\xcf\xe4\xc9\x06\x7c\x56\x9a\x56\xce\xbd\x60\x61\xb5\xc4\x75\xd7\x21\x43\x1e\x49\xe7\x57\x6a\xa4\x78\xd8\x18\x0f\xef\x76\xac\x09\xda\x35\x04\xe3\xd1\x19\x4d\xd6\x53\x83\x9d\x6d\x88\x11\x36\x84\xeb\x5e\xe9\x0d\xe1\x4a\x5e\x1c\x55\xb4\x6e\x67\x1b\x61\x6c\xd6\xbf\x2e\x6e\xe6\x77\xcb\x39\x5a\xd3\x11\x0e\x67\xec\x5c\x40\x63\x98\x74\x70\xbc\x87\x6b\x11\xfe\x09\x0b\x4c\x24\xc5\xd9\x74\x18\x84\x88\x11\x0d\xb5\xc6\x12\xca\xad\x32\xb6\xc4\x30\x88\xe9\x14\x37\xa9\xcf\x9a\x2c\xb1\x0a\xd4\x60\xb5\xc7\x29\xd9\xa0\xdf\x8e\x4e\x25\x6e\xef\x71\x77\xff\x80\xf9\xed\xe2\x41\x8a\x5e\xe9\x17\xb5\x26\x24\x0f\x21\xcc\xb6\x77\x1c\x50\x89\xa2\x74\xbe\x14\x45\xb9\xda\x07\x4a\x1f\x31\x22\xd0\xb6\xef\x54\x20\x94\x23\xe5\x73\xa4\x28\xc8\x06\xaf\x37\xb4\x55\x88\x11\x3d\x1b\x1b\x5a\x94\xef\x7f\x95\x90\xdf\x0f\xde\xc3\x20\x6a\x21\x5e\x15\x63\x04\x3d\x66\x78\x7c\x22\x1b\xe4\xc2\x06\xe2\x56\x69\x8a\x29\xe2\x1c\xac\xec\x9a\x70\xf2\x3c\xc1\x89\x55\x5b\xc2\xa7\x19\xe4\x9d\xda\x92\x4f\x1e\xc5\xdf\x28\x99\xe0\xb7\x2c\x1f\x87\xf2\x30\x30\x0c\x93\xd1\x89\x6c\x93\x66\x06\x21\xda\x9d\xd5\x79\xbd\xaa\x46\x14\x45\xaa\xd1\x19\x4b\x1e\x8f\x4f\x8f\x4f\x69\x3f\x51\xb4\x8e\xf1\x3c\x39\xb4\x4b\xa1\x63\x8f\x63\xdb\x28\x8a\x62\x35\x01\x31\x27\xed\x9b\x62\xbf\x51\xdd\x32\x8b\xd5\xc8\xd4\xa2\x28\x4c\x9b\x89\x77\x33\x58\xd3\xe5\x99\xa2\x55\xa6\xab\x88\x39\xc9\xa9\xff\x98\x3b\x83\xea\x7b\xb2\x4d\x95\x7f\x27\x58\xd5\x22\xa9\xce\xcb\x65\x68\xdc\x2e\xc8\x1f\x6c\x02\x55\xf9\xea\xe5\x17\x67\xec\x11\x1c\xeb\x56\xe5\x4f\x5b\xd6\x75\xfd\xb6\xdb\x31\x25\xc5\x3b\xce\x4b\x8e\x5e\xc4\x3c\x7a\x2d\x03\x1b\xbb\x4e\x8c\x9c\x27\xa6\xaa\x3f\x64\x93\x0c\xce\x7f\x9b\x50\x5d\x66\xbb\xff\x5e\x79\xdc\x6c\x7c\xe4\x18\x8f\x17\xfa\x27\x00\x00\xff\xff\x54\xe7\x81\x8f\x3b\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 827, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\xdd\x6f\xdb\x46\x12\x7f\x96\xfe\x8a\xa9\x81\x06\x64\xa0\xd2\xbd\xa2\x28\xee\x94\xd3\x01\x45\x9a\xa0\xbe\x5e\xdc\xa0\x49\xfa\x62\x18\x2e\x4d\x0e\xe5\x8d\xc9\x5d\x75\xb9\xf2\x47\x5d\xff\xef\x87\x9d\xd9\x25\x97\x14\x29\x29\xfe\xc8\x4b\xc4\xd9\x99\xd9\x99\x1f\xe7\x6b\x97\x3e\x3c\x84\xd7\x6a\x75\xab\xc5\xf2\xc2\xc0\x77\xdf\xfe\xe3\x5f\xdf\xac\x34\xd6\x28\x0d\xbc\x4d\x33\x3c\x57\xea\x12\x8e\x64\x96\xc0\x8f\x65\x09\xc4\x54\x83\x5d\xd7\x57\x98\x27\xd3\xc3\x43\xf8\x78\x21\x6a\xa8\xd5\x5a\x67\x08\x99\xca\x11\x44\x0d\xa5\xc8\x50\xd6\x98\xc3\x5a\xe6\xa8\xc1\x5c\x20\xfc\xb8\x4a\xb3\x0b\x84\xef\x92\x6f\xfd\x2a\x14\x6a\x2d\x73\xab\x42\x48\x62\xf9\xdf\xd1\xeb\x37\xc7\x1f\xde\x40\x21\x4a\xf4\x34\xad\x94\x81\x5c\x68\xcc\x8c\xd2\xb7\xa0\x0a\x30\xc1\x7e\x46\x23\x26\xd3\xe9\x2a\xcd\x2e\xd3\x25\x42\xa9\xd2\x7c\x3a\x15\xd5\x4a\x69\x03\xd1\x74\x72\x80\x32\x53\xb9\x90\xcb\xc3\xcf\xb5\x92\x07\xd3\xc9\x41\x51\x19\xfb\x9f\xc6\xa2\xc4\xcc\x1c\x4c\x89\xc7\x2c\x55\x22\xd4\x21\x4a\x5a\x0b\x9f\x0f\xeb\xec\x02\xab\x74\x84\x7c\x88\xf9\x12\xc7\xd6\x0a\x81\x65\x3e\xb6\x28\x64\x8e\x37\x07\xd3\x78\x6a\xbd\xff\x40\x34\xd0\xe8\x70\xaf\x21\x95\x80\xd2\x24\x6e\xc1\x5c\xa4\x06\xae\xd3\x9a\xdc\xc3\x1c\x0a\xad\x2a\x48\x21\x53\xd5\xaa\x14\x16\xe3\x1a\x35\x38\x08\x92\xa9\xb9\x5d\xa1\x57\x59\x1b\xbd\xce\x0c\xdc\x4d\x27\xc7\x69\x85\xe0\xfe\xd5\x46\x0b\xb9\x84\xfe\xbf\x3f\x2c\x46\xf3\x03\x99\x56\x38\x53\x95\x30\x58\xad\xcc\xed\xc1\x1f\xd3\xc9\x6b\x25\x0b\xe1\xf8\xad\x59\xe1\x73\x57\x36\xa3\x95\xae\xf4\x9b\x7c\x89\xb5\x63\x3b\x39\x7d\x69\x1f\x47\x76\xb6\x68\xd6\x5d\xe1\xb7\x16\xc4\xba\x11\xa6\xc7\x61\x61\x82\xbb\x27\x7d\x64\x51\x76\x9b\x9f\x9c\xbe\xa4\xc7\x61\x69\xc1\x9c\x5d\xf1\x9f\x95\xba\x0c\x2c\x7f\xaf\x6a\x61\x84\x92\x03\xe2\x17\x96\xb3\x2b\xfc\x5e\x95\x22\xbb\xdd\x47\x78\x45\x9c\x5d\xe9\x1f\xa5\x54\x26\xb5\x02\x35\x54\xe9\xea\x84\x5f\xd9\xa9\x90\x06\x75\x91\x66\x78\x77\xef\xa5\xd3\x96\xb3\xa3\xe2\x9e\x42\xab\xd9\x36\xc7\x3a\xd3\xe2\x1c\x6b\x48\x61\xe5\x89\x2e\xc5\x38\x26\x5d\xe4\x34\x12\x6d\xec\x04\xb8\x09\x69\x00\x0e\x0f\x81\x49\x4e\x9e\xa0\x3f\xb4\x18\x40\x29\x6a\x93\x4c\x27\xef\xc4\x0d\xe6\x47\xe4\xec\xb9\x52\xa5\x93\x10\x59\x6a\xb0\x06\x51\x04\xbb\x82\x3a\xff\x8c\x19\x87\x77\x65\xa5\xbe\x11\x92\x15\x08\xe9\x37\xe1\x2d\x89\x04\x22\xdc\xb8\x22\x12\xef\xc9\xfe\x72\x80\x6c\x66\x12\xd3\x1f\x90\x48\x2c\x38\x9c\x47\xa3\x99\x34\x9e\x4a\x47\xb2\x50\x2d\xdb\x4b\x42\x2e\xf9\x78\xbb\xc2\xce\x82\x13\xb7\x06\x74\xc5\x3f\xa6\xe1\x66\x3b\x76\x37\x69\x2f\x13\x3f\x88\xbf\x02\xdb\x5f\x0a\x69\x7e\xf8\x7e\x54\xba\x16\x7f\xf5\x36\x7f\x23\xd7\x55\xdd\xb0\x9d\x9c\x32\x28\x77\x70\x3c\x83\xdf\xbd\x2d\x4d\x58\xa2\x65\xee\xca\x7f\x92\xe2\xcf\x75\x63\x00\xc5\xc5\xc0\x3f\x27\xbf\x26\xe6\xae\x82\x63\x51\x96\xe9\x79\x89\x7b\x29\x90\x8e\xb9\xab\xe2\xd7\x95\x8d\xed\xb4\xdc\x4b\x85\x72\xcc\x5d\x15\x3f\x61\x91\xae\x4b\xb3\x9f\x1b\x39\x33\x0f\x6a\xf8\x3d\x2d\x2d\x1c\x61\x4e\x8f\x6b\x38\xbb\xb2\xdc\x83\x7a\x7e\x11\xd2\xd6\x44\xd7\xd2\x12\xf7\x38\xa6\xe7\x52\xc8\xbc\xf7\x5e\x56\x79\x6a\xd0\xbb\xb5\xeb\xbd\x10\xf3\xd9\xa0\x5f\x47\x55\xb5\x36\xcd\x0b\xda\xa1\x48\x78\xe6\xae\x8e\xdf\xd3\x52\xe4\xa9\x51\x9a\x22\x8d\x72\x7f\x5c\xc7\x55\xc3\xdc\x0b\x74\xa3\x74\xba\xc4\x5f\x90\xea\xef\x8e\x34\xa9\x99\xf9\xec\x12\x6f\xfb\x15\x3c\x2c\xd9\x83\x15\x3c\x2c\xe2\xbc\xda\x33\x04\xa5\x25\x5f\xed\x85\x48\xed\x99\x7b\x3a\xa8\x4e\xda\x1a\x61\x79\x83\x66\xd0\xf1\xcb\xeb\x20\xe6\xb3\xcd\xca\x11\x36\x14\x18\x6b\x29\xbb\x7a\xca\xe4\xb5\xaa\x2a\x6c\xde\xc9\x0e\x60\x33\x66\x1e\xe8\x4a\x34\x03\x6c\x16\x69\x22\x3f\xa0\x46\x93\xdc\xd3\x8c\x3a\x1e\xe6\xdd\xb2\xdb\x8b\xf3\x0e\xd9\x7e\x65\x0e\xe7\x9a\xed\xa2\xd4\x31\xba\xc2\xbf\x61\xd1\xb8\xbc\x5d\x58\x63\x71\xb6\xe9\xf3\x6f\x58\x34\x8c\x83\xe3\x59\x28\x3f\x5e\xd2\x47\xa2\x7b\x4b\x3d\x3f\x92\x57\xa8\xeb\xad\xb9\xd1\x8c\x67\xc4\xd9\xb7\xfb\xcf\xb5\xd0\x98\xef\x16\xd7\x8e\x73\xbc\x4a\xbc\xb4\xb3\x67\xd2\xad\x1b\x7b\x94\x88\xa7\x1a\xd3\x78\xd2\xd9\xcc\x08\xa6\x3f\x20\x25\x58\xb0\xcd\x89\xe0\x45\x35\x50\x6d\x79\x33\xc1\xcc\x7e\xe2\xeb\xcc\x5e\x43\x7a\x9f\x7b\x68\x2a\x0f\x50\x6e\xc2\x75\x07\xd0\x8c\xd2\x31\x5e\x53\x78\x66\x1a\x69\x90\x4c\xa5\x47\xc4\x1a\xc5\xb0\xd0\x2f\x1e\x76\x57\x46\xe9\x64\x5a\xac\x65\xe6\x25\x23\xcc\xdd\x9b\xfe\xa9\xe1\x88\x5d\xcc\xdf\x4d\x27\x12\x61\xbe\x80\x17\xf6\xf1\x6e\x3a\xb1\xf9\x3c\x6f\x22\x09\xf3\xe4\x63\xba\x9c\x59\xf2\xed\x0a\xe7\x21\xd9\x16\x82\xe9\x84\xca\x4e\x48\xb7\xcf\x96\x4e\xe8\xcc\x5b\x3a\x3d\xdb\x05\x7e\x27\xf3\x66\x81\x9f\xed\x8a\x4b\x8c\xb9\x5f\x71\xcf\x76\xc9\x07\xfd\xdc\x2d\xf9\x67\x5e\x2b\x5a\x23\x68\xad\xf0\x46\xb4\x98\xcf\x69\xa9\x7d\xb6\xab\x41\x18\xcf\xa1\x4a\x2f\x31\x1a\x0e\xe6\x78\x36\x9d\xdc\x4f\x27\x85\xd2\x70\x36\x83\xd4\x58\xb8\x74\x2a\x97\x68\x55\x86\xb9\x60\xe1\x93\x98\xa4\x79\xde\x52\xa3\xd4\xc4\x24\x2e\x0a\x3b\xab\x58\x59\xb6\xf1\x15\x3d\x7e\xb5\x00\x29\x4a\x2f\x69\x6b\xd2\xa2\x79\x6d\x1a\x8b\x98\xe9\x41\xec\x2c\x80\xf9\x02\x1a\xa9\xd7\x68\xd6\x5a\x82\xc4\x36\x6a\xb8\xbc\x36\x61\xd3\xe4\x11\x91\x29\x6c\xf8\xe7\x50\xdc\x90\x6c\x54\xe4\x7e\x58\x0f\x23\x27\xe2\x03\xe9\x0c\x50\x6b\xfb\x7c\x47\xce\x15\x79\xf2\x46\xeb\xd0\x21\x6f\x92\x28\x67\x50\x54\xc6\x2e\x2b\x5d\x44\x9c\x1d\xf0\xf5\x9f\x73\xf8\xfa\xea\x60\x66\x05\xe9\x7d\x39\x0d\x8c\x56\x4d\x48\xbd\xa0\x8d\xee\xfa\x61\x06\x8d\x0c\x45\x4d\xa1\xba\x2b\x96\x32\xeb\x47\x32\xad\xb8\x58\xa6\x91\x7e\x1e\x2e\x10\x65\x23\x3a\x69\xa9\x8d\x4f\x3f\x88\xcf\x5b\x1b\xfc\xb4\x3d\x9d\x34\x33\x76\xbb\xea\x29\x76\xd5\xcd\x99\xf3\x56\xaf\x9f\x3c\x19\x30\xda\x3b\x9c\x48\xe7\xb4\x77\x67\x46\x6d\x39\x9b\x91\x73\xde\xf8\xdc\xcc\x95\xfd\xb0\xa7\xe5\x6e\xe0\xb7\xd3\x26\xad\x97\x28\xa3\x22\x4f\x5a\x6a\x4c\x4a\xfc\x5c\xd6\xec\xd1\x50\x68\xb9\x99\xcf\x9a\x3d\x1a\xca\x46\x72\xc1\xae\xf4\xf2\x23\x56\x80\x8f\xa3\x8c\xe6\x5e\xb1\x99\x7b\x75\x31\x9e\x7b\x75\x41\x71\x01\x8b\xdd\xf1\x59\x89\xba\xb6\xf5\x99\x5a\x8a\xb0\x42\x76\x7b\x1f\xb5\x07\x33\xab\xcb\x46\x5f\xab\xdb\x9e\x30\xe7\x0b\xa0\xa3\xa5\x85\xd2\x1e\x39\xe3\x57\x4c\xff\x6a\x01\xdf\x7a\xeb\xe8\x28\xba\x80\x17\x76\x21\x30\xcc\xbf\x60\xc7\x15\x1e\x70\x16\xcd\x01\xc7\x02\xfb\x6b\x11\xb5\x91\x13\xd3\x99\x27\x62\x2b\x6c\x37\xe5\x0b\x06\x77\x46\x01\x3a\x39\x41\x96\x4a\x38\x47\xa0\x0b\x41\xcc\xc1\x28\xe2\x59\xa2\x44\x9d\x52\xc2\x5b\xc9\xb7\x4a\x03\xde\xa4\xd5\xaa\xc4\x19\x48\x65\x20\x05\x5b\x07\x68\xec\x2f\xc5\x25\x82\x11\x15\x26\xc7\xea\x3a\x21\x8b\xcf\x28\xf3\xad\xc3\xb6\x7d\x25\xef\x52\x5d\x5f\xa4\x65\x68\xd9\x2b\x62\x08\xa0\x6e\xbd\xe2\xe3\xdf\x22\xc8\x80\xb0\x7c\xd5\xc5\xcc\xca\xb4\x35\x8c\x3b\xfa\x66\xeb\xe3\x0b\x11\x2a\x62\xfc\x73\xa8\x88\x91\x70\x24\xf2\x1b\x7b\xea\xcf\xf1\xa6\xdb\xff\x58\xf5\x5d\xb3\xf7\x0b\x22\x58\x6b\x69\x0e\x70\xa1\x28\xf2\x1b\x9a\xd0\xeb\xa6\xa9\xb9\x15\xbb\xc0\xcf\xfd\xba\x61\x57\xda\xaa\x11\x26\xa3\x5d\xe9\xa4\xe2\xbd\xf3\xd4\x61\xe8\xee\x2f\xf9\x6d\xd1\x9b\x0a\xee\x43\x9b\x94\xb1\xbf\x14\xa4\xf0\xdf\x0f\xbf\x1e\x5b\x61\x1a\x94\xdc\x8b\xce\x91\x5f\x34\xb1\x58\x05\x1f\x3a\xf7\x4d\xfc\x9f\x43\xa8\xb3\x69\x54\xfb\xbd\xed\xfc\xe5\x76\x8a\x21\x3a\x87\x93\xd3\xf3\x5b\x83\xfc\xce\xdb\x8a\x5f\x53\x7d\x66\xd9\x3b\x4a\x60\x59\x08\x5f\x6f\xdd\xd5\x1a\xd3\xa2\x78\x63\x4e\x10\x92\xaf\xb6\xa3\x5e\x70\xb3\x5c\x1c\x53\x7a\xb1\xdc\x17\x36\x69\x51\xf8\xd8\xac\x13\x1b\x2a\x74\x7d\xe6\xf5\x72\x58\xee\xd1\xa1\x1c\x16\xd4\xa2\xae\x6d\xc2\xbb\x0e\x85\xbe\x3d\xb5\xd7\xd7\xc1\x84\x0b\xea\x0a\xb5\x16\x39\x36\x57\x7a\xe1\x6a\x32\x58\xc5\x1c\x52\x81\x97\x51\xcc\x19\x33\x5e\xca\x3a\x0e\x72\x08\x3e\xbd\x87\x3c\xf0\x36\x7b\xa5\x05\x52\x16\xf8\x8d\x1a\x43\x9e\x62\x2f\x87\x0b\x86\x83\x95\x9d\xc4\x19\x07\x9e\xca\x17\x90\xae\x56\x28\xf3\xc8\x11\x66\xed\x74\x1b\xa4\x75\x14\xc7\x0e\x26\x77\xb3\x1d\x3a\xe0\xee\xc5\x9f\xd3\x05\x5b\x6b\x1a\x27\x9c\x0d\xce\x0d\x7f\x2b\x1f\x38\x72\xe4\x8d\x0c\x6b\xd5\xa0\x37\xbd\x97\x4e\x57\xf4\x4f\xff\xce\xfb\xdb\xf0\x65\xfe\xd3\xef\xe3\x04\x3b\xdd\xa3\x8e\x5d\x29\xfc\x24\xab\x4e\x31\xe4\x8a\x56\x73\xdf\x12\x57\x28\xe1\x7c\x5d\x14\xa8\x81\x6a\xa0\x6b\x07\xfe\x2e\x9f\xea\x5a\x4f\x43\x74\xbe\x2e\x5c\x11\xb3\xb3\x2b\x13\x67\x63\xa5\xac\x03\x03\x59\xd8\xa8\xb3\x8a\x66\x50\x6f\x07\x02\xb5\x0e\x03\xa2\x08\x52\xdd\xb5\x0b\x12\x69\xf7\x28\x12\xd7\xb1\xeb\x68\x53\xf3\xa6\x6a\xab\x3b\xe8\x97\x61\xbb\x6c\xea\x1d\xfd\xaa\xdd\x77\x02\xa3\xfc\x37\x07\x3e\x29\x86\xf5\xdd\x01\x16\xd5\xe0\x60\x89\xa1\x5f\x34\xfb\x0d\x81\x60\xb3\xb6\x91\xf6\x4e\x7e\x75\x6a\xed\x96\xec\x0a\x21\x12\x33\xa8\x82\x94\x61\x93\xe9\x34\x94\x56\x6e\xa6\x1a\x6e\x15\xd5\x4d\xd3\x26\xa6\x93\x89\x3b\x70\x87\xd6\xb8\xc2\x58\xdd\xc4\x2d\xdc\x03\xc8\x76\x07\x3f\xbb\x7b\x13\xb7\x32\x88\x5a\x6b\x2f\x19\xfc\xb9\xf3\x4e\x8b\xf6\x8d\x4e\xec\xec\xe2\xf6\x6f\x0f\x50\xdd\x6c\xb6\x6c\x03\xa6\x7c\xa9\x2d\x64\x8c\x9d\xa9\x9a\xcb\xd9\x05\xbc\xf0\xbf\x59\x23\x95\x13\xd7\x6f\x3f\xcf\x88\xe4\xbe\x4e\x11\xd1\x68\x1e\x4e\x26\xc1\x27\xa7\x39\x88\x59\xab\xdc\x07\x6b\x50\xae\xdc\xb4\x03\x75\xe1\x01\x19\x6b\x12\x4f\x0d\xfa\x58\x73\x78\x50\x77\x20\xad\xdb\xfa\xc3\x33\x58\x3f\xda\x17\x1e\xd3\x18\x68\x03\xfe\x06\x1b\xba\xc1\xcd\xe1\xc9\xe3\xbe\xb5\x9f\xb6\xf4\xd6\xf3\xd7\xe2\xc0\xf6\x9f\xd9\xa0\x27\x8c\x47\x6f\x86\xfb\x62\x1c\xfa\xea\x3a\xd4\x53\x3a\x2b\x0a\xe0\x8d\x3a\x8a\xea\xc4\x7d\xd9\x0e\x3c\x7d\xef\xec\xe9\xb9\xfa\xc5\x7e\x0d\x8c\x85\xd5\xcd\xc0\x48\x38\x3c\x13\x76\x1b\x42\xb7\x1b\xb8\x1c\xe6\x76\xc0\x07\xd8\x07\xb4\x83\xce\x88\x39\xda\x0f\xc6\x4b\xf0\x17\x77\x84\xe1\x02\xbb\x5f\x7d\x1d\x0f\x82\xa6\x7d\x8e\x56\x4e\xff\x7a\x88\x67\x57\x01\xdc\xc0\x7c\x10\xbb\x70\x52\x1b\x85\x6e\x2c\x87\xbf\x10\xb8\xa1\x0c\xdd\x37\x41\x9b\xfc\xe4\xd8\x6c\x62\xb8\x48\x4b\xbe\x81\xbd\xdf\xdb\xe5\xce\xd4\x38\xea\xf3\x78\x32\xef\xef\xf5\x60\xaa\xee\x97\xa9\xfb\xb9\xd3\x4b\x37\xb9\x79\x5c\xa3\xcc\xcc\xd6\x5a\xcf\x40\x5d\xf2\xe4\x1c\x24\xee\x49\x2a\xdd\x8c\x72\x4a\xd6\x7e\xa5\x2e\x9d\x8d\xc3\x4c\xd6\x66\xd9\xf8\xe9\x7d\xac\xbc\x6e\xbb\x4f\xe2\xf0\x49\xde\xa1\x5e\xa2\x8e\x5f\xc1\x6e\x9d\x15\x33\x47\xa9\x8c\xdd\x95\x03\x7b\x8a\xfc\x11\xe0\x41\x7e\xe2\x3e\x7e\x8e\x31\x3d\xc6\xcf\x2d\x3a\xc7\xfc\x2c\x80\x6f\xac\x1f\xe4\x68\xb1\x8f\xa3\x63\x4c\x8f\x71\x74\x8b\xce\xdd\x8e\xb6\x67\x8a\x36\xe5\xac\xbd\xed\x65\xe3\xdf\x7f\xdb\xa7\x23\x59\xa8\xe4\x78\x5d\xa1\x16\x59\x14\x5b\x62\xef\xfe\xb1\xbd\x80\x7c\x6b\xb7\xe8\x1e\x77\xc8\x21\xd9\x22\x15\xde\xf1\x25\x51\x51\xaa\xd4\xfc\xf0\x7d\xdc\x41\x6a\xa0\x23\xaf\x25\xde\xac\x30\x33\x98\xf7\x2e\x2f\xe9\x02\xb6\xb9\x7b\x9d\xf3\xe5\x6b\x78\xf7\x5a\x5f\x0b\x93\x5d\x80\xe1\xdd\xc9\x17\x7b\x3c\x78\x45\xef\x30\xad\x11\x0c\xfc\x67\x01\xe1\x9f\x1b\x99\x7f\xc2\x8b\x17\x60\xe0\xdf\x3d\xf2\x0f\xdf\xcf\x09\xf2\xde\x2d\x25\xdf\xe8\x5a\x94\x87\xd4\x7d\x12\xc3\xfa\x3e\x89\x51\x85\xeb\x56\xe3\x50\xd3\x6e\xbb\x26\x5c\xeb\x74\x55\x87\x7f\xa8\xe6\xe8\xa9\xcc\xf9\x98\xe4\x09\x15\x9a\x0b\x95\xc3\xb5\x30\x17\xa0\x31\x53\x57\x7c\x36\x46\x59\xaf\x35\x82\x54\xb0\x4a\xa5\xc8\x6a\x10\x12\xdc\x41\x56\xc8\xa5\x6b\xf5\x41\x97\x2e\xf2\xe0\x6f\x72\xc0\x11\x63\x38\x39\x6d\xff\x90\xec\x3e\x86\xc8\x35\xe4\x80\xdc\xbf\x19\xcc\xd1\x9e\xce\xad\x7a\x37\xb7\x88\x02\xae\xa8\x37\xb1\x71\xf6\x98\x7b\xd5\x69\xd0\x74\x59\xdc\x09\x89\xaf\x3f\x7a\xef\xd8\xf8\xe6\x53\xd1\x0c\xae\xe8\x04\x54\xf8\xe6\x4c\x51\x48\x33\x90\x3d\x08\xfa\xe8\x72\xdf\x1b\xeb\x28\x9e\xf5\xd0\xe5\xf3\xc2\x06\xb8\x4c\x7e\x2c\x94\xe1\x15\x59\x88\x26\xd3\x3d\x98\xf4\xe1\xd5\x62\xc9\x07\x99\x96\xf8\x1c\x48\x76\xfc\xeb\x80\xc9\x40\xa2\x3b\x3f\x0d\xe2\x18\x0a\x6f\x42\xe9\x0f\x2e\x1b\x60\xfa\x85\xc7\xc2\xd9\xbd\xb0\x0b\x01\xf5\x2b\x1e\x52\xbe\xcb\xb7\x98\xfa\xc3\x55\x40\x7f\x46\x58\xbd\xa7\x03\xc0\x8a\xe6\x58\xb7\x0d\xda\xc6\x91\x3e\xb8\x7c\x91\xb3\x01\x2d\x93\x1f\x0b\xec\xb6\x0b\x9e\x88\x0f\x48\x8c\xdf\xbb\xf6\x92\xe7\x59\xf0\x63\x77\x06\xd0\x63\x23\xb6\x63\xc7\x5e\x6c\x20\xc7\x03\xef\x06\x72\x4c\x7e\x2c\x72\x9d\x79\x3e\x08\x48\xa6\xfb\x70\xb4\x4f\x14\x8d\x3c\x88\xb7\xc4\x67\x84\x92\xfd\x1b\x80\xf2\xc2\x1d\x00\xb6\x41\xe9\xcc\xef\x43\xe9\x26\xe9\x0d\x2c\x1d\xfd\xb1\x60\x76\x4f\x0a\x01\x9a\x6e\x21\xa6\xd8\x74\x9b\x59\x38\xdd\xb4\xdf\x52\x9f\x11\x4f\xb7\xed\x00\xa0\x2b\x7f\xbe\xd8\x86\xa8\x77\x61\xd6\x39\x5c\x34\xb7\x99\xa6\xf3\x5d\x37\xee\x3c\xd1\x69\x5a\x69\x30\xee\x03\x6f\x38\x84\xbd\x37\x34\xca\x4d\x0c\x2c\xc0\x24\x6f\x4a\xac\xa2\xce\x28\x61\xa6\xf7\xd3\xff\x07\x00\x00\xff\xff\x50\x9b\x50\x07\xda\x32\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 13018, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
