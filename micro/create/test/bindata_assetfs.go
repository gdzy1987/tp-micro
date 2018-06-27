// Code generated by go-bindata.
// sources:
// __tp-micro__tpl__.go
// DO NOT EDIT!

package test

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

var ___tpMicro__tpl__Go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x92\xcf\x6e\xdb\x30\x0c\xc6\xcf\xd2\x53\x10\x3d\x25\x39\xcc\xe9\x56\xf4\x10\x0c\x03\xbc\x76\x58\x0a\x24\x4d\xb6\xb4\x87\x9d\x5c\xcd\x66\x14\x6d\x96\xe4\x4a\xcc\xb0\xa2\xc8\xbb\x0f\x92\x9c\xfa\xcf\x72\x30\x14\x91\xfc\xf1\xe3\x27\x66\x19\x34\xa2\xfc\x2d\x24\x42\x51\x3c\x6c\x57\x45\x01\xca\x03\x1d\x10\x1a\x67\x7f\x61\x49\x40\xa8\x9b\x5a\x10\xf2\x51\x1e\xe7\x59\x06\x45\x91\x6f\xef\x8a\x62\xfb\xb8\x0a\x95\x0e\xa5\xf2\x84\x0e\xc2\x7f\x70\xf6\x48\xe8\x16\x21\x0d\xb2\x83\xd5\x98\x4e\x5a\xd0\x21\xab\xd4\x1f\x55\x21\xa7\x97\x06\x47\x0c\x65\x08\xdd\x5e\x94\x08\xaf\x9c\x2d\xad\xc6\xc9\xcc\x93\x3b\x96\xf4\x7a\x9a\xc2\x2c\x5c\x7c\x47\x7f\xac\x89\xb3\xb5\xa0\x03\x3f\x0d\x65\xec\x96\x43\x19\xbb\xe5\x50\x86\x27\x41\xa3\xae\xb1\x64\xd0\x75\x47\x82\x26\xb3\xf0\xcd\x9d\x9c\xbe\xb5\x58\xff\xd8\x7d\x5b\x15\xc5\x7a\x73\xfb\x25\x08\x2d\x1d\x0a\x42\xd0\x2f\xfe\xb9\x06\x6d\x2b\xac\xcf\xe0\x51\x62\x92\x1f\xc0\x8f\x1e\x5d\x87\xdb\xdc\x7f\xdd\xfc\x8f\xb3\x46\xda\xea\xe7\x08\x38\x4c\xed\x80\xb7\xd1\xc6\xdc\xc9\x96\x1a\x2c\x81\xd2\x1a\x72\xb6\xae\xd1\xa5\xfa\x78\x39\x18\x30\xcb\x20\x55\xc2\x41\x98\x2a\x24\xb6\xa4\xc9\xec\x8d\x38\x85\xf6\xdc\xda\x9d\x1a\x74\xfe\x43\x78\x51\x70\x29\x18\xfb\xf4\x62\x9d\xc2\x1b\x6b\x08\x4d\xbc\x51\x46\x42\x96\x01\xe1\xdf\x48\x8b\x35\x93\x9e\x98\xdc\x49\x48\x7b\x01\xa2\x51\x20\x9c\xec\x0d\xd8\x43\x86\x8a\x94\x67\x2a\xce\x58\x0e\xfb\xda\x0a\xba\xbe\xea\x22\xde\x3a\xce\xd8\xe7\x73\x00\x9e\x1a\xe1\x84\x5e\x5c\x7c\x74\xc2\x48\x5c\xc0\xfc\xdd\xfc\x72\x71\x39\x0f\xbf\x4f\x17\x4f\x9c\x9d\x7a\x2a\xda\x09\x7a\x42\xda\x19\xd9\x20\x3e\x94\xf3\x7c\xb4\xa4\xd0\x10\x67\xec\xa6\x93\x73\xe2\xd3\x68\x5a\xbb\x49\x10\x96\xef\xec\x78\x1c\x2f\x5a\xd0\x45\xcf\xc4\x07\x1f\x9e\xeb\xfa\x2a\xba\xa5\x34\x7a\x12\xba\xf1\xed\x0b\x84\x1d\x82\x63\xf8\x28\xb3\xb7\x09\x11\xef\xba\xfa\xbb\x0a\x20\x11\x38\xbb\x17\x1a\x5b\xf3\x39\xcb\x25\xc6\xc0\x87\xf7\xfc\xc4\xff\x05\x00\x00\xff\xff\xaf\x58\xc0\xae\xfa\x03\x00\x00")

func __tpMicro__tpl__GoBytes() ([]byte, error) {
	return bindataRead(
		___tpMicro__tpl__Go,
		"__tp-micro__tpl__.go",
	)
}

func __tpMicro__tpl__Go() (*asset, error) {
	bytes, err := __tpMicro__tpl__GoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "__tp-micro__tpl__.go", size: 1018, mode: os.FileMode(420), modTime: time.Unix(1530009776, 0)}
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
	"__tp-micro__tpl__.go": __tpMicro__tpl__Go,
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
	"__tp-micro__tpl__.go": &bintree{__tpMicro__tpl__Go, map[string]*bintree{}},
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