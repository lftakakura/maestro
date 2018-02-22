// Code generated by go-bindata.
// sources:
// migrations/0001-CreateSchedulerTable.sql
// migrations/0002-CreateUserTable.sql
// migrations/0003-CreateSchedulerVersionsTable.sql
// migrations/0004-AlterSchedulerTableAddVersion.sql
// DO NOT EDIT!

package migrations

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

var _migrations0001CreateschedulertableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xdf\x6e\xd3\x30\x14\xc6\xef\xf3\x14\x9f\x76\xd5\x4a\xa4\x19\x13\x13\x52\x87\x10\x65\x75\x87\x45\xea\x40\xeb\x68\x1d\x37\x91\xe7\x9c\x25\x11\x49\x1c\x6c\x67\xd5\x1e\x89\xd7\xe0\xc9\x50\xda\xad\xe3\x8f\x50\xb9\xf4\xf9\x7e\xbf\xe3\xa3\x73\xc2\x10\x8d\x22\xe7\xad\x09\xc2\x10\xa5\xf7\x9d\x9b\x46\x51\x51\xf9\xb2\xbf\x9d\x68\xd3\x44\xde\x74\x77\x96\xa8\x50\x0d\xb9\xe8\x19\x1d\xe8\xb8\xd2\xd4\x3a\xca\xd1\xb7\x39\x59\xf8\x92\xb0\xe4\x12\xf5\xbe\x3c\x7d\x6a\x38\x8d\xa2\xed\x76\x3b\x31\x1d\xb5\xce\xf4\x56\xd3\xc4\xd8\x22\x7a\xa4\x5c\xd4\x54\x3e\x7c\x7c\x0c\xc6\xa5\xe9\x1e\x6c\x55\x94\x1e\x3f\xbe\xe3\xec\xf4\xe5\x6b\x48\xd3\x61\x61\x89\x70\x35\xcc\x80\x37\xb7\x4a\x7f\xa5\x36\x7f\xe7\xef\x0a\x6d\x86\x19\xdf\x06\xc1\xe5\x8a\xcd\x24\x03\xdb\x48\x26\xd6\x3c\x11\xe0\x0b\x88\x44\x82\x6d\xf8\x5a\xae\x71\xd2\xf7\x55\x1e\x1a\xe7\xba\x93\x8b\x03\x2c\x67\xef\x63\x06\xa7\x4b\xca\xfb\x9a\xac\xc3\x28\x00\x80\x2a\x47\x9a\xf2\x39\x3e\xad\xf8\x72\xb6\xba\xc1\x47\x76\x83\x39\x5b\xcc\xd2\x58\x62\x68\x93\x15\xd4\x92\x55\x9e\xb2\xfb\x57\xa3\xf1\x8b\x9d\xd3\xaa\x86\x70\xaf\xac\x2e\x95\x1d\x9d\x9d\x9f\x8f\x77\x9f\x8b\x34\x8e\xf7\x79\x71\x24\x7f\x50\x4d\x0d\xc9\x36\xf2\x8f\xba\xf3\xca\xd3\x3f\x83\xac\x56\xce\x67\xba\x54\x6d\x41\x79\xa6\x3c\xb8\x90\xec\x8a\xad\x0e\xec\x61\xee\xd3\xbd\xb5\xe3\x9d\x56\x35\x65\xa6\xfb\x1f\x41\x5b\x52\x7e\xdf\xdb\x57\x0d\x39\xaf\x9a\x0e\xd7\x5c\x7e\x80\xe4\x4b\x86\x2f\x89\x60\x7f\xbb\x22\xb9\x7e\xda\x4b\xdf\xe5\xc7\xfd\x34\x8e\x83\xf1\xf3\x59\x52\xc1\x3f\xa7\x0c\x5c\xcc\xd9\xe6\x97\xeb\x64\xc3\x8e\xb3\xbe\xad\xbe\xf5\x84\x44\xfc\x76\xb7\x21\x1a\x5f\x04\x3f\x03\x00\x00\xff\xff\x05\x7a\x84\x4f\xcc\x02\x00\x00")

func migrations0001CreateschedulertableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0001CreateschedulertableSql,
		"migrations/0001-CreateSchedulerTable.sql",
	)
}

func migrations0001CreateschedulertableSql() (*asset, error) {
	bytes, err := migrations0001CreateschedulertableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0001-CreateSchedulerTable.sql", size: 716, mode: os.FileMode(420), modTime: time.Unix(1495042893, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations0002CreateusertableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8e\xd1\x6a\xdb\x30\x18\x46\xef\xfd\x14\xdf\x5d\x63\x98\xa3\xad\x50\x06\x59\x29\xf3\x3c\x75\x35\x75\x9c\x2d\xc8\x94\xee\xc6\xa8\xf2\x1f\xdb\xa4\xb6\x84\x24\x2f\xf3\x23\xed\x35\xf6\x64\xc3\x4d\x4b\x1b\x46\xd8\x7a\x29\x71\xbe\xf3\x9f\x28\x42\x27\xc9\x79\xab\x83\x28\x42\xe3\xbd\x71\x0b\xc6\xea\xd6\x37\xc3\xdd\x5c\xe9\x8e\x79\x6d\x36\x96\xa8\x96\x1d\x39\xf6\x8c\x4e\x74\xd6\x2a\xea\x1d\x55\x18\xfa\x8a\x2c\x7c\x43\x58\xa6\x02\xf7\xfb\xef\xc5\x93\x70\xc1\xd8\x6e\xb7\x9b\x6b\x43\xbd\xd3\x83\x55\x34\xd7\xb6\x66\x8f\x94\x63\x5d\xeb\xa3\xc7\xc7\xb4\x48\xb4\x19\x6d\x5b\x37\x1e\xbf\x7f\xe1\xf4\xed\xbb\xf7\x10\xda\xe0\xd2\x12\xe1\xcb\xd4\x80\xf3\x3b\xa9\xb6\xd4\x57\x1f\xfd\xa6\x56\x7a\x6a\xbc\x08\x82\x64\xcd\x63\xc1\x21\xe2\x4f\x19\xc7\xe0\xc8\x3a\xcc\x02\x00\xd8\xd2\x58\x4a\xa5\xc8\xb9\xd2\xeb\x2d\xf5\xf8\x21\xad\x6a\xa4\x9d\x9d\x9e\x9d\x85\xf8\xba\x4e\x97\xf1\xfa\x16\xd7\xfc\x16\xc9\x15\x4f\xae\x31\xfb\x6b\x70\x7e\x81\x93\x93\xf0\xcd\x83\xed\xb8\xa9\xc8\xd3\x6f\x05\x47\xbe\x12\xc8\x8b\x2c\x7b\xb2\x1d\x35\x59\xda\x58\x72\xcd\x2b\x54\x87\x8b\x97\x2e\xfa\x69\x5a\x3b\xc2\xb7\x1d\x39\x2f\x3b\x83\x9b\x54\x5c\x41\xa4\x4b\x8e\xef\xab\xfc\xd9\xb4\xa7\x1f\xf6\xa5\x1f\x0d\x1d\x9e\x3d\xa4\xa8\x93\xed\xfd\x7f\x75\xed\xc9\x97\x3d\xca\x92\xf4\x54\x95\xd2\xff\xbb\x09\x9f\xf9\x65\x5c\x64\x02\xf9\xea\x66\x16\x06\xe1\x87\xe0\x4f\x00\x00\x00\xff\xff\x16\x41\x45\x18\x92\x02\x00\x00")

func migrations0002CreateusertableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0002CreateusertableSql,
		"migrations/0002-CreateUserTable.sql",
	)
}

func migrations0002CreateusertableSql() (*asset, error) {
	bytes, err := migrations0002CreateusertableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0002-CreateUserTable.sql", size: 658, mode: os.FileMode(420), modTime: time.Unix(1496674916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations0003CreateschedulerversionstableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xd1\x6e\x9b\x40\x10\x45\xdf\xf9\x8a\x51\x9e\x8c\x14\x42\x1b\x35\x52\x65\x57\x55\xa9\x19\xa7\xab\xe2\x25\x85\x45\x71\xfa\x82\x36\x30\x01\x54\xc3\xd2\xdd\xc5\x56\x3e\xa9\xbf\xd1\x2f\xab\xb0\x4d\x1c\x45\x7d\x9c\xd9\x7b\xef\x9e\xb9\x9e\x07\xad\x24\x63\xb5\x72\x3c\x0f\x6a\x6b\x7b\x33\xf7\xfd\xaa\xb1\xf5\xf0\x78\x55\xa8\xd6\xb7\xaa\x7f\xd2\x44\x95\x6c\xc9\xf8\x67\xe9\xa8\x8e\x9a\x82\x3a\x43\x25\x0c\x5d\x49\x1a\x6c\x4d\xb0\x66\x02\xb6\xc7\xf5\x7c\x0a\x9c\xfb\xfe\x7e\xbf\xbf\x52\x3d\x75\x46\x0d\xba\xa0\x2b\xa5\x2b\xff\xa4\x32\x7e\xdb\x58\xef\x34\x8c\x8e\xa5\xea\x9f\x75\x53\xd5\x16\xfe\xfe\x81\xeb\x77\xef\x3f\x82\x50\x3d\xac\x34\x11\xdc\x8e\x0c\xf0\xe9\x51\x16\xbf\xa8\x2b\xbf\xd8\xa7\xaa\x50\x23\xe3\x67\xc7\x59\x26\x18\x08\x04\xdc\x08\xe4\x29\x8b\x39\xb0\x15\xf0\x58\x00\x6e\x58\x2a\x52\xb8\x18\x86\xa6\xf4\x94\x31\xfd\xc5\xe2\x45\x2c\x82\xaf\x11\x82\x29\x6a\x2a\x87\x2d\xe9\x7c\x47\xda\x34\xaa\x33\x30\x73\x00\x00\x9a\x12\xb2\x8c\x85\x70\x97\xb0\x75\x90\x3c\xc0\x77\x7c\x80\x10\x57\x41\x16\x09\x18\xe3\xf2\x8a\x3a\xd2\xd2\x52\xbe\xfb\x30\x73\x2f\x0f\x9e\x4e\xb6\x04\x3b\xa9\x8b\x5a\xea\xd9\xf5\xcd\x8d\x7b\x80\xe0\x59\x14\x41\x82\x2b\x4c\x90\x2f\x31\x3d\x7f\x69\x66\xa3\xc1\x85\x98\x43\x88\x11\x0a\x84\x65\x90\x2e\x83\x10\x8f\x69\x27\x20\x60\x5c\xe0\x2d\x26\x2f\x59\xc7\xd7\x67\xd9\x6e\x41\xe0\x46\xbc\xd9\x17\x9a\xa4\xa5\x32\x97\x16\x04\x5b\x63\x2a\x82\xf5\x1d\xdc\x33\xf1\xed\x30\xc2\xcf\x98\xe3\x99\x6a\x3a\x88\xc7\xf7\x33\xd7\x71\xcf\xe5\x64\x9c\xfd\xc8\x10\x18\x0f\x71\xf3\x0a\x78\x2a\x29\x1f\xba\xe6\xf7\x40\x23\xf9\xff\x0a\x1c\xcf\xba\x9c\xf8\xdd\xc5\x14\xfa\x26\x2d\x3f\xd4\xf5\x3a\xe2\x64\x75\x17\xce\xbf\x00\x00\x00\xff\xff\x58\xbc\x08\xe9\x98\x02\x00\x00")

func migrations0003CreateschedulerversionstableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0003CreateschedulerversionstableSql,
		"migrations/0003-CreateSchedulerVersionsTable.sql",
	)
}

func migrations0003CreateschedulerversionstableSql() (*asset, error) {
	bytes, err := migrations0003CreateschedulerversionstableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0003-CreateSchedulerVersionsTable.sql", size: 664, mode: os.FileMode(420), modTime: time.Unix(1519473040, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations0004AlterschedulertableaddversionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x41\x4e\xc3\x30\x10\x45\xf7\x3e\xc5\x5c\xc0\x35\xb0\x42\x01\x21\x02\x0d\x55\xa5\xc0\xa2\xf2\x05\x52\x67\x6a\x5b\x34\x1e\x6b\x66\x42\xc4\x91\xb8\x06\x27\x43\x11\x45\x2c\xff\xd7\xfb\x5f\xcf\x5a\x98\x06\x14\x65\x32\xd6\x42\x52\xad\xd2\x38\x17\xb3\xa6\xf9\xb8\x09\x34\x39\xa5\x7a\x62\xc4\x38\x4c\x28\xee\x1f\x5d\xe9\x3e\x07\x2c\x82\x23\xcc\x65\x44\x06\x4d\x08\xaf\x7b\x0f\xe7\xdf\xba\xf9\x3b\x6c\x9c\x5b\x96\x65\x43\x15\x8b\xd0\xcc\x01\x37\xc4\xd1\x5d\x28\x71\x53\x56\x7b\x09\xeb\xe2\x99\xea\x27\xe7\x98\x14\xbe\xbf\xe0\xe6\xea\xfa\x16\x3c\x55\x78\x61\x44\xd8\xad\x0e\x70\x7f\x1c\xc2\x3b\x96\xf1\x51\x4f\x31\xd0\xea\xf8\x60\x4c\xdb\xfb\xee\x00\xbe\x7d\xea\x3b\x90\x90\x70\x9c\xcf\xc8\x62\xda\xed\x16\x3e\x90\x25\x53\x81\xfd\x9b\xef\x76\xdd\xe1\xce\xfc\x04\x00\x00\xff\xff\xe9\xe1\x9a\x95\xf3\x00\x00\x00")

func migrations0004AlterschedulertableaddversionSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0004AlterschedulertableaddversionSql,
		"migrations/0004-AlterSchedulerTableAddVersion.sql",
	)
}

func migrations0004AlterschedulertableaddversionSql() (*asset, error) {
	bytes, err := migrations0004AlterschedulertableaddversionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0004-AlterSchedulerTableAddVersion.sql", size: 243, mode: os.FileMode(420), modTime: time.Unix(1519473040, 0)}
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
	"migrations/0001-CreateSchedulerTable.sql": migrations0001CreateschedulertableSql,
	"migrations/0002-CreateUserTable.sql": migrations0002CreateusertableSql,
	"migrations/0003-CreateSchedulerVersionsTable.sql": migrations0003CreateschedulerversionstableSql,
	"migrations/0004-AlterSchedulerTableAddVersion.sql": migrations0004AlterschedulertableaddversionSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"0001-CreateSchedulerTable.sql": &bintree{migrations0001CreateschedulertableSql, map[string]*bintree{}},
		"0002-CreateUserTable.sql": &bintree{migrations0002CreateusertableSql, map[string]*bintree{}},
		"0003-CreateSchedulerVersionsTable.sql": &bintree{migrations0003CreateschedulerversionstableSql, map[string]*bintree{}},
		"0004-AlterSchedulerTableAddVersion.sql": &bintree{migrations0004AlterschedulertableaddversionSql, map[string]*bintree{}},
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

