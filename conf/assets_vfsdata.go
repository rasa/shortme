// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package conf

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
			modTime: time.Date(2018, 11, 29, 6, 10, 8, 938975300, time.UTC),
		},
		"/config.json": &vfsgen۰CompressedFileInfo{
			name:             "config.json",
			modTime:          time.Date(2018, 11, 29, 5, 54, 50, 516444500, time.UTC),
			uncompressedSize: 1892,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\xd9\x6e\xeb\x36\x10\x7d\x37\xe0\x7f\x18\xb0\x2f\x2d\x60\x38\x72\x9c\x55\x4f\x6d\x9a\x74\x4d\xd2\x36\x4b\x57\x14\x02\x2d\x8e\x2d\xc6\x14\xa9\x90\x94\x2d\xa7\xc8\x6b\x7f\xa0\x7f\xd8\x2f\xb9\x18\xc9\x92\xbc\xc4\x46\x62\x03\x16\xc5\x33\x33\xe7\x78\xe6\x88\xfa\xa7\xdb\x01\x60\x89\xf7\x19\x0b\xa1\xbc\x01\x60\x9f\xc1\xb5\x74\x1e\x35\x70\x21\x2c\x3a\xc7\x42\x60\xff\xff\xfb\x1f\xeb\x2d\x71\x55\xa2\xb4\x1b\xf4\xcb\x6f\x38\x0c\x86\x01\x23\xf4\xb5\x47\xf5\x1c\x3e\xe7\xa8\x63\x8c\x46\x3c\x9e\xa2\x16\x14\x9a\x2e\xdc\xb3\xaa\x4a\xac\xe2\x62\x82\x76\x85\x5a\x48\xba\x63\xfd\x86\x6b\xc6\x55\x8e\xd1\xd6\xb6\x5b\xe8\x38\x9a\x5b\xe9\x91\xd4\x79\x9b\x63\x8d\x4c\x71\x11\x69\x9e\x22\xc5\xbb\xc4\x58\x9f\x62\x93\x35\xe2\x5a\xcc\xa5\xf0\x09\x0b\x61\x10\x04\x41\xa5\x78\x5d\x92\x18\x91\x9c\xba\x11\x37\xa4\x1a\x6a\x10\x26\xa8\xd1\x72\x6f\x2c\x5c\xde\xdf\x36\x6d\x59\x4a\x77\x65\x4b\xea\xd8\xb0\x5e\x7c\xe9\xe3\xec\xf3\xc1\xe1\x69\xd9\xaa\x41\x38\x1c\x06\x27\x5f\x1c\xd4\x20\xeb\x75\xd6\xa9\x62\xa3\x35\xc6\x5e\x1a\x0d\x99\x31\x0a\x52\x5e\x80\x14\x0a\x57\x80\x0d\xde\x94\x17\x11\x45\x44\x14\x41\xdd\x38\x7a\x4f\x4d\x93\xa1\xde\x5f\x93\x22\xda\x9a\xdb\xb3\xb5\x28\xa4\x6b\x7b\x45\x5e\x61\x21\xd4\x1f\x16\x9e\x0c\x4f\xcf\xeb\x7a\x19\x77\x6e\x6e\xac\xa8\x03\x58\x03\x18\xa3\x22\x27\x5f\x70\x89\x0c\x82\xde\xd6\x18\xab\x8c\x7a\x96\xad\x10\xda\x58\x8e\xab\xbb\x31\x2f\x82\xc0\xa1\x9d\xc9\x18\xc1\x22\x17\x20\x46\x6b\x23\x5b\x26\x10\x14\xd5\x83\xab\x08\x22\x1b\x36\xab\xb7\x47\xd7\x9a\x6a\x2f\x6d\xe9\xce\x1d\xbc\x25\xb6\x49\x3c\x6f\x88\xe7\x1f\x24\xfe\x88\x69\xba\x3b\x5c\x73\xf6\xbe\xb2\xbb\x7c\xd3\xdd\x61\x9c\xb3\x95\x67\x2c\x36\x69\x5a\xa6\xb4\xf3\xaa\x5a\x96\x5b\xe5\xc0\x27\xdc\xc3\x5c\x2a\x05\x23\x84\xb1\x54\x1e\x2d\x0a\xf0\x06\x72\x87\x9b\x34\x23\xc5\xe3\x69\x54\x39\x80\x92\x59\x08\x7f\xb1\x19\x5a\x47\x92\x7a\xc0\x12\xe4\xca\x27\xb4\x2a\x63\x68\x81\x45\xc6\xb5\xa0\x55\xec\x1c\x5d\x9e\xca\xdf\x71\x1e\x4f\xcb\x38\x9f\x67\x52\xb0\xbf\x57\xbb\x70\xc1\x1d\x82\xf3\x56\xea\x09\x89\x28\xc5\x2c\x4f\x00\x6c\x95\x6f\x69\xe3\x0e\xa3\x2a\x8b\xa0\x4b\x37\xfc\xf1\xfc\xcf\xdb\xd9\x6f\xe9\x77\x31\x9f\xda\x81\xf9\x59\x17\xc9\xd1\x73\x76\x73\xf5\xf2\xd5\x02\xcf\xe6\xbf\x1f\x7f\x2f\x7e\x38\xbc\xfe\x26\x7f\x7a\x9c\xf8\xaf\x83\x53\xf5\xd3\xc3\xe8\xe4\xdb\x3f\x2e\x7e\xf9\x55\xde\x8f\xef\xd6\x86\x7d\x5f\x53\x36\x1e\x13\x26\xe5\x52\x03\x3d\x28\x7d\x78\x48\xa4\x03\xe9\x1a\xa9\x55\x0f\x5b\xa1\xa0\x8c\xc9\xfa\x9b\x72\xab\x12\xcd\x91\xb9\x62\x3a\x3a\xd3\xf7\xd3\xbb\x38\xc1\x94\x87\x40\xef\x0f\x30\xb6\xbc\xba\x2d\x86\x2a\x8a\x76\xcb\xf7\xcc\x6a\xc9\x07\xe9\x15\x82\x19\x83\xd4\x02\x8b\x7e\xe2\x53\xb5\x95\xee\x29\x86\x36\x4b\xfe\x9b\x75\xff\x5f\x15\x3c\xcd\x14\xc2\xe3\xdd\x35\x48\xbd\xaf\x4c\x63\x96\xed\x7f\x79\x70\x59\x47\xb5\x0f\x00\x2f\x64\x9a\xa7\xa0\x50\x4f\x7c\x42\x0a\xd7\xf2\xdf\x2c\x1c\xa5\xbc\x60\x21\x1c\x1e\x1f\xd7\xa5\x5e\xbb\x9d\xd7\x6e\xe7\x53\x00\x00\x00\xff\xff\xe4\x7e\xcd\x22\x64\x07\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/config.json"].(os.FileInfo),
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