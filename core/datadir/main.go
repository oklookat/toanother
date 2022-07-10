package datadir

import (
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const DEF_PERM = 0755

// TODO: разделить функции, добавить тесты, и так сделать со всем проектом.

// create [./data] dir if not exists.
func Init() (err error) {
	abs, err := GetFullPath(".")
	if err != nil {
		return
	}
	err = os.MkdirAll(abs, DEF_PERM)
	return
}

// if path absolute: do nothing
//
// if relative: returns absolute [./data] + [path arg].
func GetFullPath(path string) (absolute string, err error) {
	// check is absolute.
	if filepath.IsAbs(path) {
		// just clean.
		return filepath.Clean(path), nil
	}

	absolute = filepath.Join("./data", path)
	absolute, err = filepath.Abs(absolute)
	return
}

// (re)Write data to file (not binary).
//
// If file/dir not exists, creates it.
//
// path: abs/relative path to file.
func WriteFile(path string, data []byte) (err error) {
	var withoutFilename = filepath.Dir(path)
	if err := CreateDirIfNotExists(withoutFilename); err != nil {
		return err
	}
	abs, err := GetFullPath(path)
	if err != nil {
		return
	}
	return os.WriteFile(abs, data, DEF_PERM)
}

// Write JSON/YAML to file by struct.
//
// path: abs/relative path to file.
//
// If dir not exists - error.
//
// If file not exists - creates it.
func WriteFileStruct(path string, useYAML bool, structPointer interface{}) (err error) {
	if structPointer == nil {
		return errors.New("nil structPointer")
	}

	// open.
	file, err := OpenFile(path, os.O_WRONLY|os.O_CREATE)
	if err != nil {
		return err
	}
	defer file.Close()

	// clean.
	if err = file.Truncate(0); err != nil {
		return err
	}
	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return err
	}

	// write.
	if useYAML {
		err = yaml.NewEncoder(file).Encode(structPointer)
	} else {
		err = json.NewEncoder(file).Encode(structPointer)
	}

	return
}

// Open file and decode to struct.
//
// Path: abs/relative path to .json/.yml file.
func GetStructByFile(path string, useYAML bool, structPointer interface{}) (err error) {
	if structPointer == nil {
		return errors.New("nil structPointer")
	}

	// check.
	exists, err := IsFileExists(path)
	if err != nil {
		return
	}
	if !exists {
		err = errors.New(`file "` + path + `" not exists`)
		return
	}

	// open.
	file, err := OpenFile(path, os.O_RDONLY)
	if err != nil {
		return err
	}
	defer file.Close()

	if useYAML {
		err = yaml.NewDecoder(file).Decode(structPointer)
	} else {
		err = json.NewDecoder(file).Decode(structPointer)
	}

	return
}

// create temp file in OS temp dir.
func TempFile() (file *os.File, err error) {
	file, err = os.CreateTemp(os.TempDir(), "tan-*")
	return
}

// open file.
func OpenFile(path string, flag int) (file *os.File, err error) {
	abs, err := GetFullPath(path)
	if err != nil {
		return
	}
	return os.OpenFile(abs, flag, DEF_PERM)
}

// Create dir if not exists (recursive).
//
// Path: abs/relative path.
func CreateDirIfNotExists(path string) (err error) {
	abs, err := GetFullPath(path)
	if err != nil {
		return
	}
	_, err = os.Stat(abs)
	if os.IsNotExist(err) {
		return os.MkdirAll(abs, DEF_PERM)
	}
	return
}

// Check is file exists.
func IsFileExists(path string) (exists bool, err error) {
	abs, err := GetFullPath(path)
	if err != nil {
		return
	}
	stat, err := os.Stat(abs)
	if stat != nil && err == nil {
		exists = true
		return
	}
	if errors.Is(err, fs.ErrNotExist) {
		err = nil
		return
	}
	return
}
