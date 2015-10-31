// Code generated by go-bindata.
// sources:
// templates/client.tmpl
// DO NOT EDIT!

package main

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\xc1\x6a\xeb\x30\x10\xbc\xfb\x2b\xf6\x05\xde\xc3\x0e\xc1\xb9\x1b\x72\x79\x4d\x28\x81\x12\x42\xa1\x1f\xa0\xd8\x6b\x47\xd4\x59\x0b\x69\x95\x50\x42\xfe\xbd\x5a\xc9\x6d\xdc\x50\x9d\xa4\xd9\xd5\xcc\x68\xb4\x46\xd5\xef\xaa\x43\xb8\x5e\xa1\xdc\x8f\xfb\xdb\x2d\xcb\xf4\xc9\x0c\x96\x21\xcf\x20\x2c\x29\x6e\x23\xe0\xa4\x58\x64\xd9\x72\x19\xc1\x35\xba\xda\x6a\xc3\x7a\x20\x29\xf0\x87\x49\x4c\x3b\x75\x12\x1a\x70\x6c\x7d\xcd\x70\x8d\x2c\xff\x95\xc3\xb7\xd7\x17\x98\x7b\xdb\x97\x61\x13\xc1\xa7\x5e\x23\x31\x00\xeb\xce\x1d\x99\x4d\x99\x80\xec\x16\x35\x76\x78\x99\xd2\xd5\x16\x15\xa3\x03\x05\x84\x97\x1f\x42\xde\x69\xea\x80\x8f\x08\x91\x64\x8d\xad\xf2\x3d\x8f\xe4\x8a\x1a\x21\x93\x6a\xa7\xcf\x48\x70\x48\x4e\x4a\xd8\xb6\x11\x95\x33\x88\xb5\x5a\x11\xd0\xc0\x70\x40\x30\xca\x3a\x6c\xc2\x5d\x40\x6b\x07\x0b\xda\x81\x45\xf6\x96\xb0\x29\xb3\xd6\x53\xfd\x60\x2e\x1f\x49\xe5\xcd\xc1\x4b\x01\xf9\x7c\x52\x5d\x24\x96\x62\x8c\xc2\xc7\x33\x54\x2b\x90\x2c\xf6\x22\xf5\x75\xbf\x88\x0d\xba\x8d\x0d\x7f\x56\x40\xba\x1f\x2f\xc9\x4a\x16\x04\x5c\x40\x7b\xe2\x72\x23\xac\x6d\x3e\xd3\x74\x56\xbd\x6e\xee\x2f\x69\x83\xe5\x87\x90\x2a\xf8\xeb\x66\x51\x38\x69\x84\x88\x27\x94\xff\x26\x9d\x77\xbd\xf1\xcf\xaa\xe0\xf8\x1b\x4b\xa9\x56\xbf\x24\x9d\x7a\xc2\x63\x83\xbf\xf0\x83\x81\xd1\x2a\x0a\x03\x55\x6e\xa8\x31\x83\xa6\x38\x3e\xa2\xf3\x8c\x84\x36\x7c\x65\x3a\x22\x35\x32\x3d\x9f\x01\x00\x00\xff\xff\xe9\x4a\x5d\x30\x8b\x02\x00\x00")

func templatesClientTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesClientTmpl,
		"templates/client.tmpl",
	)
}

func templatesClientTmpl() (*asset, error) {
	bytes, err := templatesClientTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/client.tmpl", size: 651, mode: os.FileMode(436), modTime: time.Unix(1446314732, 0)}
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
	"templates/client.tmpl": templatesClientTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"client.tmpl": &bintree{templatesClientTmpl, map[string]*bintree{}},
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

