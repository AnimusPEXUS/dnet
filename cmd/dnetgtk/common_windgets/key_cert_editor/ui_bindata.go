// Code generated by go-bindata.
// sources:
// ui/key-cert-editor.glade
// DO NOT EDIT!

package key_cert_editor

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

var _uiKeyCertEditorGlade = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x5a\x4d\x6f\xe3\x36\x13\xbe\xfb\x57\xf0\xe5\xf5\x85\xf3\xb1\x45\x8b\x1e\x6c\x2d\xba\x0b\x64\xbb\xc8\xb6\x08\x90\xb4\x45\x4f\x02\x45\x8d\x65\xae\x29\x8e\x4a\x8e\xec\xa8\x45\xff\x7b\x21\xab\x46\xec\x98\xb2\xbe\x12\xc7\x49\x6f\x0e\xc4\xa1\x66\x9e\xe7\xe1\x70\x66\x94\xc9\xfb\xfb\x54\xb3\x25\x58\xa7\xd0\x4c\xf9\xe5\xd9\x05\x67\x60\x24\xc6\xca\x24\x53\xfe\xcb\xdd\xd5\xf8\x7b\xfe\x3e\x18\x4d\xfe\x37\x1e\xb3\x4f\x60\xc0\x0a\x82\x98\xad\x14\xcd\x59\xa2\x45\x0c\xec\x9b\xb3\x77\x17\x67\x17\x6c\x3c\x0e\x46\x13\x65\x08\xec\x4c\x48\x08\x46\x8c\x4d\x2c\xfc\x91\x2b\x0b\x8e\x69\x15\x4d\x79\x42\x8b\xff\xf3\x87\x17\x95\x66\xfc\x7c\xbd\x0e\xa3\xaf\x20\x89\x49\x2d\x9c\x9b\xf2\x4f\xb4\xf8\x09\x4c\xce\x99\x8a\xa7\x3c\x05\x93\x5f\xf2\x72\x15\x63\x93\xcc\x62\x06\x96\x0a\x66\x44\x0a\x53\xbe\x54\x4e\x45\x1a\x78\x70\x67\x73\x98\x9c\x6f\x9e\xfa\x17\x4b\x61\xc2\x19\xca\xdc\xf1\xe0\x4a\x68\xb7\xbf\x5e\xce\x95\x8e\xab\xdf\x3e\x97\x3e\xce\x41\xae\xfd\xfa\x4c\x90\x56\xbe\xc9\x54\x85\x12\xb3\x82\x6f\xac\x3a\xba\xd8\xc7\x4d\x9f\x8d\xc1\xd0\xcd\x71\x15\x0a\xad\xdb\xbf\x48\x8b\x08\x34\x67\x64\x85\x71\x5a\x90\x88\x34\x4c\x79\x01\x8e\x07\x3f\x68\x8d\x2b\xf6\x11\xb3\xa2\xcd\x3e\xb9\x83\x30\x37\x31\x58\xad\x4c\x6d\x9c\x93\xf3\x0a\xcf\x7f\xa1\x3e\xdf\xc2\xba\x27\xee\x65\xc0\x6f\x0c\xf7\xdb\x39\xae\xd8\x35\x78\x50\x1f\x00\xdf\x2d\x64\xc2\x0a\x42\xbb\x0b\xa1\x4b\x55\x48\x98\x24\x1a\x5c\xe8\x36\x4b\x8e\x8b\xe7\x80\xa0\x76\x63\x49\x55\xa8\x51\xc4\x2f\xad\x86\x7a\x62\xbf\xa0\x88\xd9\xcc\x62\xca\xae\x94\x86\xb3\xb3\x53\x38\x56\x7b\x10\x3a\xb1\x84\xd3\x85\xf0\x56\x2c\x81\x11\x9e\x10\x80\xfb\x07\xeb\xd5\x9e\x1e\x09\x96\x08\x51\x87\x0b\x28\x42\x65\x66\xf8\xd2\x3a\x78\xe2\xc4\xba\x89\x8f\x8d\xc7\x0b\x28\xc6\x65\x84\xa7\xa0\xa0\x7a\x1e\xb2\x3c\x7a\xfb\x54\x54\x41\xbe\x02\x36\xca\x1f\x6a\xa6\xa4\x20\x78\xe3\x94\x6c\x45\x7a\x24\x5e\xb6\x1f\xec\x73\xf2\x01\xef\x2b\x3a\x2c\x22\x1d\xa5\x09\x78\xb4\x1e\xad\x02\x43\x82\x14\x1a\x1e\x2c\x4b\x74\xa4\xd0\x4d\x46\x2e\x13\x52\x99\x84\x07\xdf\x76\xed\x31\xca\x78\x8f\x2d\xae\x1d\x97\xfc\x6e\x7d\xa9\xb4\x53\x12\xb1\x96\x51\x48\x45\xb6\x5d\x2b\xf4\x72\xb6\xaf\xc3\xdd\xc4\xfd\xd7\x35\x14\xec\xae\xc8\xe0\x6f\xff\x5e\xbb\xc2\xdc\xec\x2e\xe4\x42\x99\xe4\xf0\x3b\xe1\x3e\x13\x26\xee\xe8\xe8\x4c\x1d\x3a\xb3\x3e\x8b\x0c\x9d\xaa\xe4\x77\x51\x17\xc1\x9e\xbb\x3b\x79\xaf\x37\xc9\x60\xed\x4e\x4b\x70\x64\x96\x5f\x11\x33\x97\x7d\x99\x79\x1c\xe3\x7e\x7c\x9d\x63\xeb\x16\x57\x27\xb5\x3d\x8a\xa7\xf3\xed\xfa\x21\x27\x2a\xb7\x7e\x96\x04\xd7\xd6\xc4\x82\x04\xb5\x04\x17\xc6\x30\x13\xb9\xa6\x2e\xd8\x64\x79\xc6\x83\xf5\x40\xaa\x6f\x26\xfd\x9c\x8a\xe4\xa4\x32\xa7\x23\x94\x0b\x1e\x24\xb4\x18\x67\x16\x66\x60\xc1\x48\x70\x6d\x4f\xe4\xeb\x91\xb3\x87\xb1\x21\x72\xfe\x19\x09\x22\xc4\x45\x95\x2f\x4d\x74\x6c\x51\xb7\x90\xda\x63\x17\xc3\x59\x5e\x42\xf7\xac\xd2\xeb\x62\xb6\x2e\x61\x49\x44\x9d\x15\x5b\xda\x45\x68\x63\xb0\x4d\x96\x7b\x28\xf9\x91\xba\x95\x16\xb5\x86\xf8\x37\x65\xe2\x9d\xa1\xe2\x20\x98\x06\x40\xe5\x33\x9d\xbb\xb5\x97\x91\xb0\x61\x86\x5a\xc9\x82\x07\x42\xaf\x44\x51\x73\x56\x6b\x1c\xef\xb5\x89\x17\x46\x3f\x94\x77\x70\x4f\xbf\x2a\x58\x55\xa2\xa3\x95\x4f\x74\x83\x51\x1d\x88\xac\xcf\x1c\x62\xb5\xae\x1a\x1b\x14\x55\x67\x9e\xa2\xc1\xb2\xea\x6f\xe5\xbc\xaf\xb0\xd9\x3c\xf1\x2b\xd6\x6b\xe0\x5d\x5c\x71\xc5\xca\xfa\x7c\xca\x49\x44\x7b\xd0\xd7\xd5\x7d\xc7\x11\x7d\x23\xb8\xdd\x86\xe5\xa6\x7e\xab\x3a\x90\xfd\x15\xa4\xef\xd5\x24\xa2\xb0\xba\x6d\x1a\xbc\xf6\x94\x79\x0d\xfc\x34\x93\xb2\xdb\x08\x9e\x14\x25\xa9\xb0\x89\x32\xa1\x86\x19\xf1\xe0\xb2\xa6\x29\x39\x68\x6a\x55\x32\xef\x6b\x4b\x98\xf5\xb4\x8c\x90\x08\xd3\xae\xc6\xad\x06\x00\xcd\x0c\xfb\x59\xf6\x55\x83\x83\xb9\x1e\xca\xb7\xcf\xbe\x10\x5a\x25\xf5\x8d\x4e\xbd\xe1\x56\x79\x59\x37\x4f\x3a\x64\xae\x24\x9a\xd0\xa9\x3f\x81\x07\xdf\xf5\x4d\xaa\x87\xce\xbc\xef\x9d\x9b\x92\xb4\x0f\xce\xcd\x4d\x64\x9d\x65\x63\x9b\xff\x10\x68\x7d\x06\xf3\x5f\x1f\xdd\x34\x59\x77\x1d\xf8\x9c\x7e\x51\x4d\xd6\x5f\x0d\x84\xac\xac\x41\xd8\x8d\x55\x4b\x41\xc0\xae\xa1\x18\xf1\x87\x8f\xda\x9c\x09\x13\x33\xbe\xf9\xda\xca\x47\xca\xb0\xb2\xa9\x63\x22\xc2\x25\xb0\xd1\x08\x8d\x2e\xd8\xef\x98\x33\x37\x47\x1d\xb3\x85\xc1\x55\xf9\xa7\xdd\xde\xb0\xab\xb7\x5f\x73\x47\x6a\x56\xf0\x40\x82\x21\xb0\x7d\x0f\x60\xa3\x38\xde\xca\x29\x68\xc8\x34\xdd\x4f\xc1\xf0\x82\xa0\x95\x73\xdd\x8b\x81\xff\x42\xb1\xf6\xa3\x8a\x63\x78\xd6\x6a\xad\xa5\x72\x8e\x50\xe4\x35\x4e\x49\x1a\x88\x6f\x47\xfa\x89\x8e\xdb\xaf\x72\xad\xeb\xd3\x63\xff\x91\x6e\x2b\x9e\x9e\x6a\x12\x7e\x70\x1e\xf0\x92\x23\x93\xce\x33\x80\xc1\xfd\x7f\xdb\x7e\x65\xaf\xef\xcf\x84\xa5\x13\x9d\xa4\x74\xe8\xf5\x7b\xf7\xf9\xad\x5b\xf6\xfe\x47\xe2\x09\x3e\x3f\xbc\xe5\x4c\x74\x93\x47\x5a\x49\x76\x23\x2c\x3d\x75\x32\x6a\x71\xd7\x3c\x53\xfe\xea\x3f\x60\x6f\x3b\x7c\xee\x3b\x5f\x7f\xd7\x65\xbe\xfe\x10\xc6\xe4\x7c\xeb\x1f\x86\xff\x09\x00\x00\xff\xff\x9f\xfd\x1a\x6a\x89\x2c\x00\x00")

func uiKeyCertEditorGladeBytes() ([]byte, error) {
	return bindataRead(
		_uiKeyCertEditorGlade,
		"ui/key-cert-editor.glade",
	)
}

func uiKeyCertEditorGlade() (*asset, error) {
	bytes, err := uiKeyCertEditorGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/key-cert-editor.glade", size: 11401, mode: os.FileMode(420), modTime: time.Unix(1494318561, 0)}
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
	"ui/key-cert-editor.glade": uiKeyCertEditorGlade,
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
	"ui": &bintree{nil, map[string]*bintree{
		"key-cert-editor.glade": &bintree{uiKeyCertEditorGlade, map[string]*bintree{}},
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
