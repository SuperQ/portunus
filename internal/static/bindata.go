// Code generated by go-bindata. DO NOT EDIT.
// sources:
// static/css/portunus.css

package static


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataCssPortunuscss = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x4d\x8f\xe3\x28\x13\xbe\xbf\xbf\x02\xcd\x68\xa4\x77\x5b\xb6\x45" +
	"\x3a\x9a\x68\x1a\x1f\xa3\xce\xaf\xd8\x0b\xb6\xcb\x31\x6a\x0c\x16\xe0\x7c\x2c\xca\x7f\x5f\x61\x70\x82\x63\xbb\x95" +
	"\xdb\x1e\xba\x63\xa0\x3e\x9f\x7a\x28\xea\x2d\x79\x23\x05\xd4\x52\x41\xf2\x46\x68\x6d\x40\xd9\x42\x5e\x52\xcd\xfe" +
	"\x61\xe2\x48\x0a\xa9\x2a\x50\x69\x21\x2f\xb7\xc6\xb4\xdc\xd6\x52\x18\x77\x06\x64\xb3\xeb\x2e\xf9\xb0\xac\x69\xcb" +
	"\xf8\x95\xa4\xb4\xeb\x38\xa4\xfa\xaa\x0d\xb4\x09\xf2\xbf\x69\xcf\x12\xa4\xa9\xd0\xa9\x06\xc5\xea\x9c\x33\x01\x69" +
	"\x03\xec\xd8\x18\xb2\xc9\x7e\xdf\x68\x42\x49\x2d\xcb\x5e\x27\x94\x34\xf2\x04\x2a\xa1\x84\x96\x86\x9d\xc0\x96\x92" +
	"\x4b\x45\x0a\xde\x43\x6e\xe0\x62\xd2\x0a\x4a\xa9\xa8\x61\x52\x10\x21\x05\x0c\xf1\x24\x85\xac\xae\xb6\xa5\xea\xc8" +
	"\x04\xc1\x79\x47\xab\xca\x45\x8d\x6f\x82\x9e\x6c\x41\xcb\xaf\xa3\x92\xbd\xa8\xc8\xcf\xfd\x7e\x9f\x8f\x06\x69\xf9" +
	"\xe5\xce\xd1\x9b\x0d\x81\xbc\x67\xbf\x15\xb4\x93\xd8\xfc\xd6\x20\xd6\x73\xdb\xd2\x4b\x7a\x66\x95\x69\xc8\x1f\x8c" +
	"\xbb\x4b\x3e\x3a\x44\xb4\x37\xf2\xe1\x15\xe1\xd1\x90\x36\xa9\x36\x57\x0e\x43\xa4\x79\xc5\x74\xc7\xe9\x95\xd4\x1c" +
	"\x2e\xc1\x24\xe2\xcc\x9e\x40\x19\x56\x52\x9e\x52\xce\x8e\x82\xb4\xac\xaa\x38\x3c\xce\x33\xdd\xd1\x12\x94\x75\x5a" +
	"\xc4\x39\x7a\x1c\xa1\x66\xf3\x48\xda\x7f\xa4\x6a\xc4\xd4\x45\x10\x95\x29\x4e\xc4\xa9\x52\x3b\x86\xc3\xc4\x90\x71" +
	"\xc1\x65\xf9\x15\xc0\x61\xa2\x01\xc5\xcc\x22\xe0\xb3\x3c\x63\xa3\x59\xd9\x2b\x05\xc2\x4c\x40\x3f\x1c\x0e\x4b\x32" +
	"\xa1\xc0\xc9\xd2\x91\xe7\xc0\xd2\xc9\xc0\x12\x3b\x09\xf3\xd6\x52\x26\xd6\x8a\xb3\x51\xd0\x0e\xf5\x41\xef\xe3\xd7" +
	"\x3c\x83\xac\x30\x62\x19\x8f\xbb\xa8\x07\x74\xb1\x56\x53\x36\xe7\xe7\x86\x19\x48\x87\xaa\x11\x21\xcf\x8a\x76\xbe" +
	"\x0e\x67\x2f\x20\xa4\x6a\x29\x8f\x4b\x33\x46\xe0\xfe\xd2\x4e\xb1\x96\xaa\xeb\x04\x40\x8c\x0f\xa1\x30\x83\xed\xdc" +
	"\xdf\x46\x82\x67\x5a\x23\xa4\xb3\x7d\x8f\xe7\x6c\xdb\x83\x19\xbb\xda\xed\x0e\xb7\x5a\xaa\x16\x55\xec\x94\xb9\x8f" +
	"\x54\xc9\x73\x72\xdf\x09\x18\xb9\x4d\x1b\x03\x8c\xe7\x4a\x88\xd3\x02\xf8\xa2\xea\xb0\x56\xf2\x9c\x0e\x22\x77\xe0" +
	"\x3d\xe2\x0f\x60\x70\xf6\xc7\x41\xb3\x62\x19\xe9\x8e\x0a\xbf\x07\x4a\x49\xf5\x82\xa7\x67\x15\x1b\xaa\x9b\x72\xa8" +
	"\x0d\xc1\x99\x63\x48\x00\x5a\x41\xb5\xe0\x98\x89\xae\x37\xdf\x3b\x3a\x51\xde\xc3\x53\x4a\x9e\x96\x1b\x8c\x7f\xad" +
	"\xd9\x24\x42\x9a\xff\x67\xcc\x75\x4a\xd7\x35\xfe\x7a\xc1\xc7\xb3\x8a\x9d\xf5\x9f\x40\x93\x4d\x77\x41\x5a\x72\x56" +
	"\xa1\xa1\xfd\x85\x56\xae\x68\xc5\x7a\x4d\xb6\x63\xf7\x8e\x10\x7f\xea\xce\x3b\xb7\x35\x59\xad\x66\x31\xd0\xcc\x06" +
	"\x0f\x1e\xc8\x9f\x1f\x1f\x1f\xab\xf2\x81\xae\x6b\xc7\x81\x9e\x13\x73\xdb\xed\x76\xcd\x5c\x5c\xd9\xa9\xd2\x1e\xe3" +
	"\x7c\xd2\x91\xf6\xfb\x17\x8c\x2c\xa7\x83\x97\xb8\x3e\x53\xfd\x3e\xb3\x58\x72\x31\xc9\xd8\xc9\xeb\x0c\x88\x53\xfc" +
	"\xfc\xfc\x7c\x98\xb8\xcb\xa0\xe1\x73\xb9\xd5\x4d\x1e\x90\x40\xa1\xff\x82\x53\x4f\xc1\xfa\x75\xd9\x40\xf9\x05\xd5" +
	"\xcb\x29\xfa\x75\x2f\x82\x5e\x18\x6d\x9e\x12\xcf\x4b\x29\x8c\x7b\x54\x7e\xfc\xfd\xbe\xdb\xe0\x1f\xaf\xc5\xf0\x8a" +
	"\xad\xcd\xba\xad\x51\x7d\xfa\xd2\x8e\x9d\x68\x44\x7f\xbb\x02\x09\xf5\x05\x9c\xbe\x7f\xab\x72\x61\xa6\x5a\x3f\x7f" +
	"\xa6\xe9\x4c\xc0\x5f\x81\x68\x06\x9b\xdf\xc6\xac\xe6\x54\x37\x4f\x2d\x6f\x1c\x4b\x06\xd6\xa0\x68\x24\xcb\xde\xfd" +
	"\x4e\x4c\xa7\x88\x38\xc1\x9a\xff\x3f\xb9\xcb\x13\xce\x2d\xdc\xe7\x89\xa2\xee\xcb\x12\xb4\x5e\x50\xc5\xfb\xa9\xea" +
	"\xfe\xb0\xbf\x19\x5a\x70\xb0\x8f\x2e\x1d\x25\xc9\x69\xa7\x81\x8c\x1f\xe3\x81\x7b\xdc\xfd\x80\x39\xa8\x22\xd3\x00" +
	"\xad\x90\xb9\xf7\x8a\x42\x1a\x23\xdb\xc8\xad\x9f\x36\xa7\xc2\xc8\x34\x36\xbe\x15\x43\xc9\x83\x88\x9b\x67\xbf\xb3" +
	"\xe7\x6e\xdd\x93\x2c\x32\x95\x5d\x18\x5d\xbd\x50\xe6\x0a\x2d\x85\x1e\x93\xfc\xe5\x47\x3b\x3f\xc5\x94\x20\x0c\xa8" +
	"\x85\xb9\xe5\xf6\xbf\x7f\x03\x00\x00\xff\xff\x3e\xb8\x31\xf3\x16\x0c\x00\x00")

func bindataCssPortunuscssBytes() ([]byte, error) {
	return bindataRead(
		_bindataCssPortunuscss,
		"css/portunus.css",
	)
}



func bindataCssPortunuscss() (*asset, error) {
	bytes, err := bindataCssPortunuscssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "css/portunus.css",
		size: 3094,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1556909438, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"css/portunus.css": bindataCssPortunuscss,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"css": {Func: nil, Children: map[string]*bintree{
		"portunus.css": {Func: bindataCssPortunuscss, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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