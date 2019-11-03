// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ab_pats.tsv
// ac_pats.tsv

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

var _bindataAbpatstsv = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xbb\x4e\x23\x41\x10\x45\xe3\xd9\xef\xd8\xcd\x1c\x4c\x3d\xbc\x8f" +
	"\xd1\x6a\x93\x25\x80\x04\x3e\x00\x59\x16\x48\x0e\x31\x04\xb6\xc4\xe7\x23\xdc\x08\xf7\xa9\xae\x1e\x1b\x6b\x24\xcb" +
	"\xb7\x1f\xa7\x6f\xd5\xed\x99\xf5\xf5\x34\x7c\xbf\x7d\x78\xda\x0c\x3f\x6e\xb6\x2f\xfb\xdd\xbd\x4e\xe3\x62\x5c\xfd" +
	"\xd5\x69\xb9\x58\xfe\xab\x07\x65\x55\x2b\x85\x32\x28\xaf\x94\x2c\x46\x28\x81\x52\x28\x83\xaa\x29\x0a\x8a\x82\xa2" +
	"\xa0\x28\x28\x0a\x8a\x81\x62\xa0\x18\x28\x06\x8a\x81\xe2\xa0\x38\x28\x0e\x8a\x83\xe2\x07\xca\xdd\x7e\x97\x74\xb9" +
	"\x1a\x15\x2e\x52\x4a\xa3\x04\xb1\xb4\xba\x92\x42\xa9\x94\x46\x09\x94\x12\xa5\x44\x29\x51\x4a\x94\x12\x65\x44\x19" +
	"\x51\x46\x94\x11\x65\x44\x39\x51\x4e\x94\x13\xe5\x44\x95\xee\xff\x7f\xde\xee\x36\xaf\x4d\xfb\xeb\x61\x09\xcb\x34" +
	"\x68\x0b\x9a\xd8\x92\x41\xad\x25\x68\x0d\xda\x82\x26\x4f\x03\x4f\x03\x4f\x03\x4f\x03\x4f\x03\xcf\x02\xcf\x02\xcf" +
	"\x02\xcf\x02\xcf\x02\xcf\x03\xcf\x03\xcf\x03\xcf\x03\xef\x3d\x96\x6f\xeb\xab\x69\x78\x1c\x87\x71\x90\x61\xc4\x23" +
	"\x1f\xbf\x72\xf8\x17\x67\x47\x8c\x0b\xf6\x64\xb4\xe3\xa8\xa4\x67\xa4\x7b\x8a\xb9\xbe\xb1\x68\x23\x3f\x52\x3e\xd7" +
	"\xe6\xe6\x5a\xf3\x3c\x87\xb6\x85\xe6\x14\xdd\xe8\x55\x2f\xcd\x8a\x4b\xcb\x39\x9d\x8e\x1c\xcd\x59\x1a\x5b\x16\x62" +
	"\x3f\x94\xde\xc1\x73\xa5\xe5\x9d\x0b\xb1\x3a\x36\xf6\x6f\x8e\x74\x66\xe7\x7a\xda\x2f\x2e\x76\xae\x79\x8a\xb9\x65" +
	"\x37\x4e\x69\x0e\xa0\xa9\xb9\x00\x5b\x23\xa7\xee\x76\x98\x2b\xe6\x7e\x5e\x14\xd2\x9c\x1d\x49\x7a\xd3\xb2\xcf\xea" +
	"\xdc\xaf\xf4\x86\x65\xc1\x9c\xd7\x87\x5e\x37\xf3\xe2\xba\x3b\x8a\xb9\xdf\x89\x91\x3e\x28\x0f\xed\xdc\x72\xf2\x75" +
	"\x33\xe6\xfe\xa4\x21\x45\xa8\x34\x5f\xa3\xde\x2b\x90\xbd\x38\x5f\xbb\x73\xe5\x53\xf2\x16\x00\x00\xff\xff\x3d\xdf" +
	"\xd4\xdd\xa6\x0a\x00\x00")

func bindataAbpatstsvBytes() ([]byte, error) {
	return bindataRead(
		_bindataAbpatstsv,
		"ab_pats.tsv",
	)
}



func bindataAbpatstsv() (*asset, error) {
	bytes, err := bindataAbpatstsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "ab_pats.tsv",
		size: 2726,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1572767188, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataAcpatstsv = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xcd\x4a\x03\x41\x10\x84\xcf\x9b\xe7\xd0\x5b\x0e\xd3\x3f\xf1\x67" +
	"\x11\x2f\x7a\xd0\x8b\x3e\x80\x84\x20\x92\xa3\xd1\x43\x02\x3e\xbe\x98\x11\x9c\xaa\xe9\xde\xc4\x10\x08\x5b\xb3\x33" +
	"\x5f\x57\xd7\x34\xbb\x7a\x18\x87\xb3\xa7\xd7\xf7\xf5\x70\xfe\xb8\xf9\xdc\x6d\x5f\x74\x2c\xf3\xb2\xbc\xd1\x71\x31" +
	"\x5f\xdc\xb6\x8b\xb2\x6c\x95\x82\x32\x50\xde\x28\x99\x17\x50\x02\x4a\x41\x19\xa8\x96\xa2\x40\x51\xa0\x28\x50\x14" +
	"\x28\x0a\x14\x03\x8a\x01\xc5\x80\x62\x40\x31\xa0\x38\x50\x1c\x28\x0e\x14\x07\x8a\xef\x29\xcf\xbb\x6d\x90\x72\xb3" +
	"\x2a\xb8\x49\x51\x1a\x4a\x20\xd6\xa8\x1b\x29\x28\x15\xa5\xa1\x04\x94\x22\x4a\x11\xa5\x88\x52\x44\x29\xa2\x0c\x51" +
	"\x86\x28\x43\x94\x21\xca\x10\xe5\x88\x72\x44\x39\xa2\x1c\x51\x35\xfd\xbb\x8f\xcd\x76\xfd\xd5\xc5\xdf\x2e\x0b\x6d" +
	"\x53\xd2\x46\x1a\xb1\xf5\x0e\x5a\x2d\xa4\x95\xb4\x91\x46\x9e\x12\x4f\x89\xa7\xc4\x53\xe2\x29\xf1\x8c\x78\x46\x3c" +
	"\x23\x9e\x11\xcf\x88\xe7\xc4\x73\xe2\x39\xf1\x9c\x78\x3f\xd7\x32\x5b\xdd\x8f\xc3\x5b\x19\xca\x20\x43\x81\x9f\xfc" +
	"\xfe\xcb\xfe\x89\xdf\xb6\xbb\x24\x39\x97\x9f\x6a\xeb\x49\x5e\xa5\x9a\xcb\x8d\xe5\x85\xe4\xa8\x67\xe9\x8c\x70\x8d" +
	"\x92\xd7\xa8\xe6\x94\x60\x3d\xaa\xed\x71\x7a\x35\x32\x92\xdb\x12\x4a\x30\x30\x67\xc1\x05\x30\x3c\xb3\xf6\xff\x76" +
	"\xf8\xbc\x4c\x27\xe7\x5d\x0f\x53\xc5\xf2\x52\xf1\x0c\x4e\x8d\x45\x94\x9c\xa0\xb9\xc5\x49\x13\xd4\x27\x97\x5f\x72" +
	"\x46\x16\xba\xe4\x20\xb9\x8b\xa0\x93\xb8\xcf\x3c\xd3\xcc\xd2\x21\x62\x3c\x73\xf2\x67\xee\x32\x9c\xb0\xbe\x7c\x64" +
	"\x36\xb3\xc4\x69\x4e\x99\x93\xb8\x85\x6a\xee\xea\xe4\x1c\xa2\x91\x8f\xbf\x7b\xfc\xee\xe8\x4f\xc9\x75\x18\x3c\xf6" +
	"\x9d\x0f\xf6\xa1\x19\xed\x0d\xf7\xd7\x1a\xee\x9d\x7d\x07\x00\x00\xff\xff\x46\x59\x13\xf9\xa6\x0a\x00\x00")

func bindataAcpatstsvBytes() ([]byte, error) {
	return bindataRead(
		_bindataAcpatstsv,
		"ac_pats.tsv",
	)
}



func bindataAcpatstsv() (*asset, error) {
	bytes, err := bindataAcpatstsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "ac_pats.tsv",
		size: 2726,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1572767225, 0),
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
	"ab_pats.tsv": bindataAbpatstsv,
	"ac_pats.tsv": bindataAcpatstsv,
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
	"ab_pats.tsv": {Func: bindataAbpatstsv, Children: map[string]*bintree{}},
	"ac_pats.tsv": {Func: bindataAcpatstsv, Children: map[string]*bintree{}},
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
