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

var _templatesClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\xd1\x6a\xeb\x30\x0c\x7d\xcf\x57\xe8\x16\xee\x25\x29\x25\x7d\x0f\xf4\xe5\xae\x65\x14\x46\x29\x83\x7d\x80\x9b\x28\xa9\x59\xaa\x04\x5b\x69\x19\xa5\xff\x3e\xc9\xce\xd6\xac\xcc\x4f\xf6\x91\x7c\xce\xf1\x91\x7b\x53\xbe\x9b\x06\xe1\x7a\x85\x7c\x3f\xee\x6f\xb7\x24\xb1\xa7\xbe\x73\x0c\x69\x02\xb2\xb4\xb8\x0d\x80\xd7\x62\x96\x24\xcb\x65\x00\xd7\xe8\x4b\x67\x7b\xb6\x1d\x69\x81\x3f\xfa\xc8\xb4\x33\x27\xa5\x01\xcf\x6e\x28\x19\xae\x81\xe5\xbf\xf1\xf8\xf6\xfa\x02\xf3\xc1\xb5\xb9\x6c\x02\xf8\xd4\x5a\x24\x06\x60\xdb\xf8\x23\x73\x9f\x47\x20\xb9\x05\x8d\x1d\x5e\xa6\x74\xa5\x43\xc3\xe8\xc1\x00\xe1\xe5\x87\xd0\xe0\x2d\x35\xc0\x47\x84\x40\xb2\xc6\xda\x0c\x2d\x8f\xe4\x86\xaa\x50\x6a\xec\x19\x09\x0e\xd1\x46\xae\xfc\xdb\x3a\x14\x14\x02\xb5\x56\x1a\x02\xea\x18\x0e\x08\xbd\x71\x1e\x2b\xb9\x0b\xe8\x5c\xe7\xc0\x7a\x70\xc8\x83\x23\xac\xf2\xa4\x1e\xa8\x7c\x30\x97\x8e\xbc\xfa\x66\xf1\x92\x41\x3a\x9f\x54\x17\x91\x25\x1b\xa3\x18\xc2\x19\x8a\x15\x68\x16\x7b\x95\xfa\xba\x9f\x85\x06\x5b\x87\x86\x3f\x2b\x20\xdb\x8e\x97\x74\x45\x0b\x0a\x2e\xa0\x3e\x71\xbe\x51\xd6\x3a\x9d\x59\x3a\x9b\xd6\x56\xf7\x97\xd4\x62\xf9\x21\xa4\x02\xfe\xfa\x59\x10\x8e\x1a\x12\xf1\x84\xf2\xdf\xa4\xf3\xae\x37\xce\xac\x10\xc7\xdf\x58\x4c\xb5\xf8\x25\xe9\xd8\x23\x8f\x15\x7f\x32\x41\x61\x74\x86\xe4\x43\xe5\x1b\xaa\xfa\xce\x52\xf8\x3e\xaa\xf3\x8c\x84\x4e\x46\x19\x8f\x28\xf3\x91\xdf\xf3\x19\x00\x00\xff\xff\x48\x4c\x78\x42\x8b\x02\x00\x00")

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

	info := bindataFileInfo{name: "templates/client.tmpl", size: 651, mode: os.FileMode(436), modTime: time.Unix(1449238052, 0)}
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
