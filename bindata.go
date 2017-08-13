// Code generated by go-bindata.
// sources:
// amanar_config_schema.json
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

var _amanar_config_schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x5d\x6e\xc2\x30\x0c\x7e\xef\x29\x22\x3f\xef\x04\x5c\x65\x42\x51\x68\x0d\x18\x85\xa4\x73\x1c\x24\x34\x71\xf7\xa9\x2d\xa3\x4d\x48\x81\xa1\x8e\x07\xda\xf8\xf7\xf3\xe7\x2f\xfd\xae\x94\x02\x12\x3c\x06\x58\xa9\xee\xa0\x14\x30\x7e\x45\x62\x6c\x60\xa5\x3e\x7b\x8b\x52\x50\x7b\xb7\xa5\x5d\x64\xb3\xb1\x18\xe0\xe3\xd7\x7c\x32\xd1\x8a\x6e\x8d\xec\x73\x1b\x7b\x8b\xd0\x9b\xd6\x83\x07\xe4\xdc\x22\xac\x14\xf8\xcd\x01\x6b\xb9\xc6\x43\xcb\xbe\x45\x16\xc2\x11\x40\x5a\x77\xb4\x4e\x6a\x04\x61\x72\x3b\xb8\x3a\x2e\xa5\xde\x7f\xca\x4b\xc7\x2b\xa6\x26\xb0\xe7\xa1\xf7\x1e\x72\x82\xd6\xd2\x41\x37\x46\x4c\xf0\x91\xeb\xbb\x98\x3b\xda\x27\x8e\xc2\x02\x26\xde\xae\xe6\xc6\x04\xd4\x31\x52\x33\xc1\x93\x04\x0c\x4d\xf5\x96\x2c\x0e\x34\x66\x61\xeb\x3c\x6f\x7e\xce\x67\xd3\x3e\xea\x5a\x8a\x9c\x5d\xc7\xf8\xbb\x14\xa6\x12\x8e\x41\xb0\xd1\xa6\x6d\x2d\xd5\x46\xc8\xbb\x32\x92\x07\xcc\xbe\xd8\xbe\x08\x60\x92\x67\x98\xcd\xf9\x35\xd4\xe9\xb2\xde\xa5\xa3\x7a\x74\xce\xda\xce\xa3\x4c\x02\x47\x8d\x72\x74\xfa\xa6\xff\x32\xad\xef\x4a\x15\xdd\x89\xd8\xbb\x23\x3a\xd1\x27\xc3\xd4\x5d\xaf\x92\x62\xef\x21\xe8\xad\xb7\x0d\x72\xfa\x69\x99\x64\xdc\x78\xdd\xfb\x20\xff\x2f\xee\x27\x00\x17\x94\x79\x91\xb2\x05\xeb\xa7\xc4\xbd\x7d\x45\x96\x92\x64\x95\xbf\x0d\xcf\xee\xbf\x2f\x92\xa7\x5e\xaa\x9f\x00\x00\x00\xff\xff\x91\x78\xb7\x12\xb2\x06\x00\x00")

func amanar_config_schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_amanar_config_schemaJson,
		"amanar_config_schema.json",
	)
}

func amanar_config_schemaJson() (*asset, error) {
	bytes, err := amanar_config_schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "amanar_config_schema.json", size: 1714, mode: os.FileMode(420), modTime: time.Unix(1502650368, 0)}
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
	"amanar_config_schema.json": amanar_config_schemaJson,
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
	"amanar_config_schema.json": &bintree{amanar_config_schemaJson, map[string]*bintree{}},
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
