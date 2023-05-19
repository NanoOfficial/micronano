//
//
// @filename: common/file.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package common

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/NanoOfficial/micronano/logger"
)

type Type string

const (
	JSON Type = "JSON"
	YML  Type = "YML"
)

type File struct {
	file     *os.File
	content  []byte
	fileType Type
}

func openFile(filePath string, fileType Type) (*File, error) {
	log.New("open file: "+filePath, log.TypeDebug)

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	content, _ := ioutil.ReadAll(file)

	return &File{
		file:     file,
		content:  content,
		fileType: fileType,
	}, nil

}

func (f *File) Close() {
	if f.file != nil {
		_ = f.file.Close()
	}
}

func (f *File) Parse(out interface{}) error {
	if f.fileType == JSON {
		err := f.parseJson(out)

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (f *File) parseJson(out interface{}) error {
	errUnmarshal := json.Unmarshal(f.content, out)
	if errUnmarshal != nil {
		return errUnmarshal
	}

	return nil
}
