// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 11, 21, 4, 4, 42, 173662100, time.UTC),
		},
		"/doc.go": &vfsgen۰FileInfo{
			name:    "doc.go",
			modTime: time.Date(2018, 11, 21, 3, 21, 26, 156175500, time.UTC),
			content: []byte("\x2f\x2f\x67\x6f\x3a\x67\x65\x6e\x65\x72\x61\x74\x65\x20\x76\x66\x73\x67\x65\x6e\x64\x65\x76\x20\x2d\x73\x6f\x75\x72\x63\x65\x3d\x22\x67\x69\x74\x68\x75\x62\x2e\x63\x6f\x6d\x2f\x72\x61\x73\x61\x2f\x73\x68\x6f\x72\x74\x6d\x65\x2f\x74\x65\x6d\x70\x6c\x61\x74\x65\x22\x2e\x41\x73\x73\x65\x74\x73\x0a\x0a\x70\x61\x63\x6b\x61\x67\x65\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x0a"),
		},
		"/health.html": &vfsgen۰FileInfo{
			name:    "health.html",
			modTime: time.Date(2018, 11, 4, 17, 55, 4, 333016300, time.UTC),
			content: []byte("\x4f\x4b"),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2018, 11, 17, 19, 55, 39, 937570200, time.UTC),
			uncompressedSize: 5276,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x6b\x53\xe3\x38\xd6\xfe\xce\xaf\x38\xed\xf7\xad\x6a\x68\xd6\x36\x21\xd0\x30\x3b\x31\xb5\x5c\x02\x4d\xa0\x69\x3a\x24\x0d\xec\xd4\x7c\x90\xed\x63\x5b\x89\x2c\x19\x49\x76\xc8\x74\xf1\xdf\xb7\x64\xe7\xe2\x84\x40\x5f\x66\x66\xab\x36\x5f\x6c\xc9\xe7\xfa\xe8\x3c\xd2\x51\x5a\x6f\x4e\x3e\x1d\xf7\xee\xaf\xdb\x90\xe8\x94\x1d\xac\xb5\xcc\x03\x18\xe1\xb1\x67\x21\xb7\x0e\xd6\x00\x5a\x09\x92\xd0\xbc\x00\xb4\x52\xd4\x04\x82\x84\x48\x85\xda\xb3\x72\x1d\xd9\xfb\x56\xfd\x53\xa2\x75\x66\xe3\x43\x4e\x0b\xcf\xba\xb3\xfb\x87\xf6\xb1\x48\x33\xa2\xa9\xcf\xd0\x82\x40\x70\x8d\x5c\x7b\xd6\x79\xdb\xc3\x30\xc6\x05\x4d\x4e\x52\xf4\xac\x82\xe2\x28\x13\x52\xd7\x84\x47\x34\xd4\x89\x17\x62\x41\x03\xb4\xcb\xc1\x3f\x80\x72\xaa\x29\x61\xb6\x0a\x08\x43\xaf\x31\x35\xf4\xc6\xb6\xa1\x97\x20\x10\x5f\x14\x08\x4d\x28\x0d\x6b\x12\x2b\x78\x97\xe6\x4a\xbf\x83\x40\xa4\x08\x11\x95\x4a\x03\xe5\xa0\x13\x04\x93\xdb\xaf\x40\xf8\x18\x84\x4e\x50\x96\xe3\xa9\x6f\x30\x4a\x95\xce\x3b\x12\x69\x94\xef\x8c\x8a\xc2\xca\xa4\x6d\x3f\x0f\x3f\x44\x15\x48\x9a\x69\x2a\x78\x2d\x83\x7b\xd4\x70\xc8\x2b\xfb\xfd\xee\x25\xdc\x24\x42\x6a\xe4\x94\xc7\x70\x83\xd2\xe4\x65\xa2\x39\x13\x06\xf5\x15\x98\x90\x5c\x27\x42\xd6\xec\x1d\xf2\x70\x0c\x77\x74\x86\x1f\xa3\x7c\x08\x12\x99\x67\xd1\xc0\x38\x4e\x24\x46\x9e\x15\x91\xc2\x0c\x1d\x1a\x08\xeb\x60\xad\x92\xd4\x54\x33\x3c\xf8\xfa\xd5\xe9\x99\x97\xa7\xa7\x96\x5b\xcd\xac\xcd\xf1\x3b\x12\x42\x2b\x2d\x49\x06\x81\x90\x08\xc7\x37\x37\xf3\x4c\xe7\x8e\x94\x1e\x33\x54\x09\xa2\x9e\xba\x33\x2b\xaf\xfe\xe9\xba\x41\xc8\x07\xca\x09\x98\xc8\xc3\x88\x11\x89\x4e\x20\x52\x97\x0c\xc8\xa3\xcb\xa8\xaf\x5c\x3d\xa2\x5a\xa3\xb4\xfd\xa9\x1b\xb7\xe9\x34\x9d\xf7\x6e\xa0\x94\x3b\x9b\x73\x52\xca\x9d\x40\x29\x0b\x28\xd7\x18\x4b\xaa\xc7\x9e\xa5\x12\xb2\xbd\xfb\xde\xde\x53\xbb\xf9\xc9\xd9\x6d\xf3\xf0\xc3\xc3\xe8\xfd\xa3\xee\xa4\x57\x57\x5a\x6e\x7e\x3a\xea\x76\xfa\x2c\x1e\x5e\x75\xda\x62\x6f\xff\x7a\xc7\xdf\x1a\x77\x47\x9e\x05\x81\x14\x4a\x09\x49\x63\xca\x3d\x8b\x70\xc1\xc7\xa9\xc8\x95\x05\x6e\x3d\xe5\xf3\x76\x63\x0b\xa6\x95\x07\x09\x09\x86\x10\x09\x09\x37\xb9\x8c\x48\x80\x6e\x88\x6a\xa8\x45\x06\xb7\x94\x87\x62\xa4\x60\x1f\xfc\x3c\x5e\x02\x65\x11\x83\x94\x3c\x06\x21\x77\x66\xf9\x98\x81\x81\xc1\x24\x49\xb1\xb1\x65\x4f\x9d\xd9\x7e\x1e\xdb\x23\x21\x87\x44\x8a\x9c\x87\x55\xce\xcb\x00\xd7\x43\x3d\xce\x95\x16\x29\x54\x5f\xcb\x28\x75\x42\x15\x68\x4c\x33\x46\x34\xae\x8c\xca\x55\x9a\x68\x1a\x94\xde\x95\x29\xbd\x14\xbf\xed\xa8\x63\x4a\xdf\xd8\x0f\xd1\xcf\xe3\xd8\xd4\x6a\x96\xcb\x4c\x28\x54\x0e\x9c\x08\xfe\x56\x03\x09\x74\x4e\x18\x1b\x43\x20\xb2\xf1\x84\x19\xdb\xc0\x28\x47\xf5\x66\x1e\xc8\x1b\xdb\xfe\x8d\x46\xc0\x34\x9c\xb7\xe1\x97\xdf\x0f\x5a\x15\x43\x40\xc9\x60\x1e\xda\xc0\xe0\xb2\x6f\x4b\x54\x99\xe0\x8a\x16\x68\x47\x94\xa1\x3d\x22\xd2\xb0\xc4\x19\x28\xeb\xa0\xe5\x56\x8a\x07\xad\x37\xbf\x21\x0f\x69\xf4\xfb\xcc\x45\xdd\x62\xad\x0c\x9d\x81\x0a\x91\xd1\x42\x3a\x1c\xb5\xcb\xb3\x74\x5e\x60\xb6\x88\xa2\x80\xf0\x82\xa8\x7f\x35\x9c\x2d\x67\xcb\xa5\x68\x63\x9a\x33\x62\x98\x6b\xa7\x22\x44\x55\x77\xbe\xa2\x0c\xc7\xed\xcf\xd9\x85\xff\x41\x9d\x7e\xde\x1a\x37\xf7\xb6\x1a\x4c\xec\xfa\x18\xe7\xb4\x93\xf7\xce\x4e\x4f\x36\xaf\xfc\xe6\xd9\xc3\xf6\xb0\xb9\x39\xfa\x24\x5e\x2e\xc3\x79\x52\x35\xe4\x3f\xf4\x3e\x5e\xee\x82\x4a\x68\x0a\x84\x87\xd0\x2d\x31\x09\x9d\x41\xb5\xdc\xe7\xed\x7d\x50\x79\x56\x56\xaa\x88\x26\xc2\xc8\x30\x45\xae\x55\xa9\x90\x62\x48\x09\x3c\xe4\x28\x29\xaa\x97\xd7\xa1\x9c\x7d\x19\xbc\xd7\x38\x6c\x8e\x88\x5d\x95\xd0\xc2\x6d\x3a\x7b\xce\xf6\x7c\x5c\xf2\x76\x35\x5e\x3b\x9f\xe4\xf9\xf1\xc9\xe0\xe8\x3e\xba\xc0\xa8\xed\x7f\xe9\xed\x8d\xda\xbd\xee\xe5\xd5\xe9\x30\x7f\xd8\xe9\x75\xbe\xec\xde\x5e\x9e\x15\x83\x87\xec\xec\x70\xf8\x5d\x78\xfd\x7c\xf4\x72\x06\xa8\xdb\x70\x76\x9c\xed\xd9\xc4\xcb\xc1\xc7\xef\xe9\x61\x54\xfc\x3b\xdb\xe4\x27\x9f\xb7\x7b\x61\xaf\xeb\x7e\xf9\x72\xd1\x89\x9a\xfe\x99\x14\x3b\xb9\xbf\x1b\x15\xb7\x37\xb7\x5f\xba\x74\xfb\xaa\xfd\xdd\xc1\x2f\x94\xf1\xab\x75\xfc\x5a\x32\x03\xb3\xce\x63\xb7\xe1\x34\x1a\x4e\x73\x32\x7a\x25\x11\xa9\xae\xfb\x67\x61\xff\xfa\xe8\x2e\x26\xac\x38\x1f\xec\xdc\x5f\x74\xa4\xec\xb3\xf4\xf2\xce\xff\xe4\xbf\x3f\xce\xf6\x82\xf0\x91\x37\x1e\xb0\x1f\x7c\x7f\x22\x3f\x1d\xb6\xf3\x20\x03\x11\xa2\xdb\x70\xb6\x16\x67\x5e\xc9\xe0\x97\x8f\x7f\x8c\x2e\xb6\x87\x9d\x8b\xa3\x54\x75\x4e\xc3\x20\xb8\x13\xe7\x27\x27\x5a\xf9\xb7\xa7\xc9\xbe\x3f\xb8\xbf\x70\x8f\xf7\xfa\x83\xa3\xc6\xe1\x78\xeb\xe7\x32\xa8\x6d\x49\xd3\xcd\x72\xb0\xcc\xd4\x96\x5b\xb5\x44\xe6\xd5\x17\xe1\x78\xba\x80\x21\x2d\x20\x60\x44\x29\xcf\x32\x87\x35\xa1\x1c\xe5\x74\x57\x9d\xb0\xfb\xa6\xb4\x0e\x9c\x14\x3e\x91\x33\x76\x02\xb4\x38\x99\xe9\x4e\x3e\x56\x0f\x3b\xc4\x88\xe4\x4c\x5b\x53\xc9\x17\xfc\xd8\x11\xcb\x69\x58\x93\x5a\x94\x9b\x18\x33\x71\x97\x31\x41\xed\xd7\x22\x4b\x52\xbe\x24\x3c\x9c\x1e\xec\xae\xb5\xd0\x30\x90\x05\x0f\x6e\x48\x8b\x03\x93\x98\xeb\x70\x52\xd8\x81\x60\x8c\x64\x0a\x6b\x89\x2d\x0a\x2d\x85\x5b\x07\xc0\xe5\xa4\x58\xc4\xea\x23\xa1\xdc\xb4\x5f\x99\xe0\xa6\x1b\x33\xdb\x1f\x81\x4c\xd2\x94\xc8\x31\xa4\x44\x0e\x51\x9b\x73\x29\x45\xa5\x48\x8c\x20\x24\x04\x84\x31\xd0\xc2\x1c\x4d\x54\xf0\xba\xf5\x1a\x14\x83\x3c\xf5\x85\x96\x82\x3f\x07\x4b\x0a\x86\x9e\x25\xc5\x68\x09\xa1\x45\x03\x94\x67\xb9\xb6\x63\x29\xf2\xec\x99\x5c\x29\x5b\x0a\x80\x1e\x67\xe8\x59\x1a\x1f\x4d\x37\x5b\x69\x46\x42\xa6\xb6\xc1\x40\x0a\x06\xe6\x8b\x1d\x20\xd7\x28\x2d\xa0\xa1\x67\x31\xc1\xe3\x7e\xf7\xd2\x82\x8c\x91\x00\x13\xc1\x42\x94\x9e\x75\x29\x78\x6c\xda\x46\x07\xa6\xe4\x1a\x8d\x46\x4e\x2c\x44\xcc\x2a\x5a\xf1\x60\x79\x41\x5f\x0d\xd9\xf6\x35\x5f\x29\x5f\xea\xf8\xb9\xd6\x82\x4f\x62\x57\xb9\x9f\xd2\x79\xf4\xbe\xe6\xe0\x6b\x6e\xab\x3c\x08\x50\xa9\x12\xe5\x02\xab\xd8\x4b\xae\x1c\x95\xca\x16\x08\x1e\x30\x1a\x0c\x3d\xab\xec\x74\xd7\xdf\x4e\x12\x7b\xbb\x61\x1d\x94\x33\x2d\xb7\x72\xb3\x32\xea\xb2\x5a\x96\xd1\x5f\x9e\x7c\x3e\x61\x92\x9d\x05\xd2\xef\x5e\x1e\x31\xc2\x87\x97\x94\xa3\xe1\xee\x2a\xe1\x49\x52\x2f\x2e\xf6\xcc\x18\x72\x0c\xcb\x65\x99\x68\x24\x3b\xd5\xd2\x4d\x2b\x31\x10\xcc\x7e\x54\xf6\xce\xc2\x82\x3e\x77\xba\xda\xf0\xe7\xae\x35\xa7\xf2\x9f\xb0\xd3\xef\x5e\x5e\x4b\x34\x7d\xe5\x8f\xda\x5b\x98\x98\x0c\xfe\x97\x58\x33\x5d\xf0\x25\xda\xdc\x4c\x91\xa9\xb8\xf3\xf5\xab\x73\x33\x11\x7c\x7a\xfa\xef\xd2\x65\x52\x27\x55\xb4\xf8\x98\x11\x1e\x3e\x23\x4a\xbb\x9c\x5e\x7f\x3b\x4d\xc6\x50\xa5\x9a\xfb\x3b\xb9\x52\x05\x53\x16\xcf\x5f\x42\x97\x9a\xbd\x1f\x62\x0b\xd4\x15\xbf\x51\xf1\x53\xd1\x3f\x49\x9c\x9a\xc7\x9f\xe0\x4d\x6d\x50\xbd\x56\x07\x96\x3b\x3b\xde\x60\xde\xd5\xad\xb8\x54\x77\x48\x41\x6e\xca\x76\xa2\x14\xf1\x7e\xf8\xb7\xd0\xd6\xc3\xb5\xa9\xfb\x10\x88\x2e\xff\xd1\x40\x1e\x9a\x9b\x81\x79\x0d\x45\x90\x9b\x8b\x01\x28\x51\x8e\x33\x12\xa3\x02\x26\x48\x08\x11\x51\x1a\xe7\x0d\xc8\xcf\x34\x71\x2f\x5d\xe5\x07\xcb\x37\xf9\xd5\x9d\xdc\xc5\x1d\xdf\xcd\xf2\x8f\xc5\xe3\xf1\x68\x33\x3c\xbc\xff\x83\xe7\x7c\x33\x94\x1f\xc3\xb3\xc6\xf9\x29\x6b\x92\xf8\x62\x2b\x73\xb3\x87\xde\x2f\x17\x87\xdf\x77\x83\x82\xbf\xe6\x3a\xbf\x0a\x88\x97\xee\xf3\x83\x6f\x5c\xe7\x07\x2f\xb5\x9a\x07\x6b\xff\xbf\x6e\xfd\xdf\xf4\xbc\xdf\x70\x86\x38\xce\x24\x2a\xb5\x1e\xe5\xbc\x6c\x5d\xd6\xb1\x40\xae\x37\xe0\xeb\x1a\xd0\x08\xaa\x91\x33\x4a\x68\x90\x80\xe7\x41\xa3\x59\x7e\x81\x6a\x3a\x93\xe5\xf3\xa4\xea\x13\xd7\x37\x7e\x5d\x03\x30\xe6\xeb\x47\xf2\x86\x53\x6e\x34\xe5\xc7\xa7\xb5\xa7\x8d\x5f\xd7\x66\x12\x7f\x63\x04\x0b\x9b\xdd\xf3\x10\x2a\xfa\xcc\xe1\x69\xb9\x55\x33\xdd\x72\xab\xff\x22\xff\x13\x00\x00\xff\xff\x4c\x0b\x1f\x23\x9c\x14\x00\x00"),
		},
		"/template.go": &vfsgen۰CompressedFileInfo{
			name:             "template.go",
			modTime:          time.Date(2018, 11, 21, 3, 21, 26, 154175400, time.UTC),
			uncompressedSize: 312,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\xc1\x4a\xc4\x30\x10\x86\xcf\x99\xa7\x18\x73\x6a\xb0\x24\x77\xa1\x07\x61\x59\xf0\xa4\x07\x5f\x20\xdb\xce\xa6\xc1\x34\x09\x93\xe9\x82\xc8\xbe\xbb\xb4\xbb\xa2\x9e\x7e\x42\xbe\xff\xe7\x63\x9c\xc3\xc7\xd3\x1a\xd3\x84\x13\x5d\x00\xaa\x1f\x3f\x7c\x20\x14\x5a\x6a\xf2\x42\x00\x71\xa9\x85\x05\x3b\x50\x3a\x14\xb7\xa3\x1a\x94\x4e\x25\x6c\x91\x49\xdc\x2c\x52\x35\x18\x80\xf3\x9a\x47\xbc\xf1\x6f\x5e\xe6\xf7\x72\x88\xdc\xfd\xbe\xb1\x09\xc7\x1c\xcc\x3d\xf1\x0b\x54\xed\x91\x98\xf1\x69\xc0\x7d\xd8\xbe\xec\xf0\x9f\x4e\x8f\x5a\xf7\xf7\xcf\x63\xcc\xd3\x6b\x4e\x9f\x06\x54\x3c\xef\xbd\x87\x01\x73\x4c\xdb\x90\x4a\x25\xd8\xa3\x17\x9f\x52\xee\x88\xd9\x80\xba\x82\x62\x92\x95\x33\x56\x7b\x88\x0c\x57\x80\x8b\x67\x7c\x6e\x8d\xa4\xe1\x80\x9b\xb6\xfd\x6f\x78\x33\xd6\x21\xca\xbc\x9e\xec\x58\x16\xc7\xbe\x79\xd7\xe6\xc2\xb2\x90\xfb\x39\x8a\x36\x06\xbe\x03\x00\x00\xff\xff\x2b\x11\xe0\x47\x38\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/doc.go"].(os.FileInfo),
		fs["/health.html"].(os.FileInfo),
		fs["/index.html"].(os.FileInfo),
		fs["/template.go"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
