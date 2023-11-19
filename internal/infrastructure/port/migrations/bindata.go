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

var __00001_schema_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\xcb\x2f\x49\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\x0f\xd7\x09\x06\x12\x00\x00\x00")

func _00001_schema_down_sql() ([]byte, error) {
	return bindata_read(
		__00001_schema_down_sql,
		"00001_schema.down.sql",
	)
}

var __00001_schema_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xc8\xcb\x2f\x49\x2d\xe6\xd2\xe0\x52\x50\x50\x50\xc8\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\x54\x08\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\xd4\x01\xcb\x27\x96\x96\x64\xe4\x17\xc1\xd5\x98\x1a\x68\x2a\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\x40\x14\xe4\xa6\x16\x17\x27\xa6\xa7\xc2\x55\x18\x99\x9a\x21\x94\x70\x69\x5a\x73\x01\x02\x00\x00\xff\xff\xf2\x66\xb4\xb7\x7b\x00\x00\x00")

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
