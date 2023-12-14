package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __00001_schema_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x2e\x4d\x2a\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x2b\xb6\xe6\xc2\x21\x13\x5f\x90\x93\x88\x26\x9d\x5c\x5a\x5c\x92\x9f\x9b\x5a\x54\x6c\xcd\x05\x08\x00\x00\xff\xff\x20\x49\x04\x93\x4f\x00\x00\x00")

func _00001_schema_down_sql() ([]byte, error) {
	return bindata_read(
		__00001_schema_down_sql,
		"00001_schema.down.sql",
	)
}

var __00001_schema_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x4d\x4b\x03\x31\x10\x86\xef\xf9\x15\x73\x4c\xc0\x8b\x05\x45\xe8\x29\xae\x53\x5d\xdc\xa6\x32\x46\xa1\xa7\x25\xb6\x11\x02\xbb\xdb\x25\x1f\xe0\xcf\x97\x5a\x57\x89\x11\x71\xef\xcf\xcc\xfb\xbc\x33\x15\xa1\xd4\x08\x5a\x5e\x37\x08\xbb\x14\xe2\xa1\xb7\x3e\x30\xce\x00\x00\x52\x72\x7b\x78\x96\x54\xdd\x49\xe2\x8b\x8b\x4b\x01\x0f\x54\xaf\x25\x6d\xe1\x1e\xb7\x67\x1f\xc8\xab\xf3\x21\xb6\x83\xe9\xed\x17\x78\xbe\xb8\x12\xa0\x36\x1a\xd4\x53\xd3\x9c\xa8\xce\xfc\x03\xb2\xbd\x71\x5d\x1e\x37\x01\x4c\x2c\x19\xcb\x54\x43\x7a\x09\x3b\xef\xc6\xe8\x0e\x43\x3b\x76\x66\x98\xe1\x9c\x89\xcc\xc8\x99\x11\x61\xdf\x46\xe7\x6d\xbb\x37\xd1\xc2\x8d\xd4\xa8\xeb\x35\xfe\xa8\x3b\x1d\xfb\x77\x93\x13\x53\xb4\xfc\x0b\x5e\x6d\x08\xeb\x5b\x75\x94\xe0\xd3\x72\x01\x84\x2b\x24\x54\x15\x3e\x7e\xbf\x97\x1f\x1b\x88\x62\xa8\x48\xcb\xa6\xcb\x8b\x7f\xae\x11\x4b\xf6\x1e\x00\x00\xff\xff\xa8\x05\x63\xc8\x46\x02\x00\x00")

func _00001_schema_up_sql() ([]byte, error) {
	return bindata_read(
		__00001_schema_up_sql,
		"00001_schema.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"00001_schema.down.sql": _00001_schema_down_sql,
	"00001_schema.up.sql": _00001_schema_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"00001_schema.down.sql": &_bintree_t{_00001_schema_down_sql, map[string]*_bintree_t{
	}},
	"00001_schema.up.sql": &_bintree_t{_00001_schema_up_sql, map[string]*_bintree_t{
	}},
}}
