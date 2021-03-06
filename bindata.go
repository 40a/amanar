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

var _amanar_config_schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x57\xdf\x6f\xdc\x36\x0c\x7e\xef\x5f\x41\xdc\x73\x71\x03\xfa\xd8\xb7\xa0\xc5\x86\x6c\x1d\xd0\x36\xc9\x80\x61\x28\x0e\x3a\x8b\x8e\xd9\xc9\x92\x43\x49\xb9\x1a\x43\xfe\xf7\x81\xb2\xe3\x1f\x77\xbe\x8b\x93\xcb\x92\x20\xcb\x43\x70\xb6\x28\x8a\x22\xf9\x7d\x1f\xfd\xcf\x1b\x00\x80\x05\x05\x2c\xfd\xe2\x3d\x34\x8f\xe9\x15\xe3\x55\x24\x46\xbd\x78\x0f\x7f\x75\x6f\xd3\x4a\xe6\x6c\x4e\x97\x91\xd5\xda\xa0\x5f\xbc\x1d\x2f\x5e\xab\x68\xc2\xaa\x52\xa1\x98\x5e\x61\x67\x70\xd1\x2d\x7c\xeb\x6d\x16\xa1\xae\x70\xf1\x1e\x16\x6e\xfd\x1d\xb3\x30\xd8\xbd\xa8\xd8\x55\xc8\x81\x70\x1c\xe2\xf6\x79\xdb\x6b\x69\x5d\xa3\xcf\x98\xaa\x40\xce\x8a\xf3\xf3\x02\x41\x8c\x81\xb1\x62\xf4\x68\x03\xd9\x4b\x08\x05\x82\x56\x41\xf9\xe0\x18\x81\x6c\x7a\xf1\x87\x78\x5e\x6e\xdd\x62\x14\xa9\x0f\x4c\xf6\x72\x31\x32\xb8\xd9\x7f\xeb\xb9\xf1\x89\xf1\x6e\x7c\x15\x72\x49\xde\x93\xb3\x1e\x42\xa1\x02\x28\x46\xf0\x2e\x5e\x16\x01\x82\xeb\x23\xee\x2f\x72\x74\xec\xe3\x42\x4f\x86\xbf\xb7\x68\x73\x8a\xd7\xd9\x90\x0d\x68\x0c\x7d\x5f\xa5\xd8\x5d\xe4\xec\x80\xf5\x54\xda\x4e\x8c\x71\x1b\x0f\xa7\x8d\x9f\x5f\xa1\xf7\x03\xd1\x23\x5b\x55\xa2\x07\x65\x35\x54\xca\xfb\x8d\x63\xed\x25\x67\x6b\x84\xac\x50\xf6\x12\xf5\x12\x7e\x77\x3e\x88\x6d\x1e\x0d\xe4\x8e\xe1\xa3\x0a\xea\x17\xa6\x2a\xed\x12\x77\x6b\xe5\xd1\xc3\x86\x42\x41\xb6\x3f\xe8\xc2\x04\x2a\x55\x98\xcc\x75\x7f\xbd\x1d\x74\x4d\x9a\xed\x45\xdc\x74\x06\xda\x98\x56\x31\x92\x3e\x70\xfa\x68\x43\x93\x93\x55\x4e\x06\x1b\xcc\x1c\xdc\xf6\xed\xb0\xd7\xbb\x8b\x3f\xb2\x9e\xd1\x08\x77\xc7\x3b\x67\x27\x1c\xc2\x7d\x70\xe3\x2e\x81\xb6\xdd\x40\xce\x58\xc2\x79\x41\xcd\x4f\x88\x3e\x2a\x63\x6a\x20\x0f\xd2\x3e\x4d\x13\x9c\x35\xc6\x4b\xe3\x32\x65\x96\x3f\x4a\x73\xa8\xee\x93\xb9\x9a\xc2\xdd\xd4\xdf\xcd\x8c\x8a\x06\x8e\x3e\xa0\x5e\xa9\xaa\x32\x94\x29\xb9\xeb\xbc\xdc\x4e\x65\xe8\x34\x4f\x1c\xf2\x1b\xd6\x59\xa1\xc8\x02\xda\xc0\x84\x1e\x6a\x17\x61\xa3\x6c\x22\x99\x8c\x51\x05\x04\x65\x18\x95\xae\x01\x7f\x90\x0f\x6f\x21\x48\xce\xc8\x43\xb4\x16\x33\xf4\x5e\x71\xbd\x84\x13\x30\xe4\x03\xb8\x1c\xf2\x68\x4c\xca\x7d\xc2\xdc\x30\xd6\xf4\x2c\xc8\x1d\x9f\x2c\x6c\x26\x37\x93\x1f\x8c\x3a\x66\x98\xd6\x6d\x2c\xd7\xc8\xe2\x71\x48\x85\x15\xbb\xb2\x0a\xbe\xad\x5c\x2a\xb1\x2f\x5c\x34\x1a\xd0\x6a\xa1\xf2\xa5\xaa\xaa\xd9\x55\x9a\x87\xd4\x71\x11\xee\x59\x58\x98\x57\xdc\x91\x6f\xc5\xac\xea\xc7\xe9\x99\x31\x6d\x1c\x03\xa7\x9e\x03\x2f\x4e\x3f\x26\xce\xbc\x95\x51\xf1\x3f\xea\x9b\x58\x69\xe1\x48\xf8\xd3\x45\xc8\x94\x85\x9c\xac\x6e\xda\x66\x2d\x5d\xa4\x4a\xb2\x43\x19\xde\x41\x59\x03\xce\xff\x0c\x6a\x07\x2d\xf6\xaf\x1e\x48\xf7\x9c\xca\xed\xd9\xde\x8b\x21\x47\xbb\xea\x24\xf8\x6e\x70\xef\xd1\xc4\x46\xe1\x12\xd6\xba\xda\xa8\x4c\x70\x2a\x70\xd6\x32\x61\x28\xb3\x2b\x6c\x1c\x2d\x8c\x0f\x7f\x7a\x89\x43\x7b\x4d\xec\x6c\x89\x36\xac\xae\x15\x93\xcc\x21\x73\x94\x6e\x37\x71\xab\xdc\x19\x8d\x3c\x35\x97\x4e\x7a\xe8\x50\x52\x38\x1f\x5e\xb2\x48\xde\x71\xd5\x87\xe2\xfb\x04\x34\x31\x66\xc1\x71\x2d\x5d\x10\x54\x03\x50\x65\xcc\xa1\x0e\x69\x67\xaa\x06\xd0\x32\x54\x5d\xb4\x1a\x2a\x40\x0e\xd8\xd0\x31\x69\x54\x3f\x71\xb4\x1f\xc6\xcd\x05\x5f\x77\xdd\x95\xaa\x06\x8b\xa8\x5b\xbf\xbe\x50\x8c\x1a\xd6\x98\xcb\x94\xbe\xc6\xcc\x95\x12\xd4\x35\x79\x5a\x9b\x76\x6c\x17\xf5\x4e\xf7\x7f\x56\x61\x9e\x6c\xdc\x63\xb8\x76\xe0\x10\x6e\x1d\xde\x7e\xa7\xec\xd4\x01\xa2\xd5\xc8\xb0\x29\x28\x2b\xc6\x9c\x9c\x39\x91\xe8\xe4\xf8\x19\xb3\x33\x06\xd7\x31\x69\xb9\x9d\xed\x47\xa3\x7d\x27\x45\x17\x5f\x3f\xc1\x86\x8c\x01\x67\x4d\x2d\x0d\xd4\xc8\x90\x06\x6a\x46\x1d\x39\x5e\xc6\x09\xb1\x6b\x73\x39\x99\xe7\x52\x85\xac\x10\x06\x95\xee\x6a\x12\xf1\x7f\xd1\xa2\xab\x88\x4c\x2e\xfa\x77\xc7\x7e\x99\xcd\x54\xa1\xf4\xdd\x9a\x78\x42\xc1\x97\xf6\x6c\x78\x07\x67\x5f\x3e\x51\xe8\xfb\xf8\xe9\x95\xa8\xcf\x83\xbf\x32\x14\xf0\xfe\x42\x92\xc6\xad\x17\x2c\x24\xd3\x37\x7c\x28\x38\x3f\xb7\x9f\x5a\x82\xa9\xad\xe2\x49\x71\x1b\x6a\x1a\x14\x38\xd5\xdd\x03\x05\x9f\xec\x9e\x95\xbc\x1f\x6f\x42\x8e\x96\xae\x22\x02\xa5\xf6\xce\x09\xb9\xa3\xa6\xee\xea\x5d\x56\x06\x53\xf2\x07\x65\x85\xad\x72\x17\xad\x96\x01\xd9\x38\xf7\xb7\x08\x5d\xcb\x51\xf7\x00\xc3\x71\xb9\x7a\x59\x54\xe4\xf1\x2a\xa2\x59\x55\xec\x9e\x88\x8b\xa4\x56\x67\xe9\x50\xf8\xcc\x0e\x2a\xf9\x9e\x7d\x86\x21\x78\x70\xef\x14\xc2\xeb\xa3\x9e\xe9\x1b\x3e\x06\xf5\x24\x77\x3d\xe1\x0c\xaa\xb9\x4d\x38\x70\x5e\x57\x94\xa5\x61\x55\xfe\xa3\x86\x9f\xd5\xb5\x63\x0a\xe8\x97\xc9\xcb\x6b\x67\xa4\x41\x6e\x1e\xc0\x49\x29\x45\xaf\x8e\x89\x76\xde\x8e\xdf\xf4\x4f\xcd\xaf\xf6\xb8\x6d\xd7\x37\x6f\xfe\x0d\x00\x00\xff\xff\x63\x45\x19\x06\x65\x19\x00\x00")

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

	info := bindataFileInfo{name: "amanar_config_schema.json", size: 6501, mode: os.FileMode(420), modTime: time.Unix(1502924339, 0)}
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

