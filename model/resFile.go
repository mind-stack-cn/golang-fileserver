/* 
 * Copyright 2015 JiaoHu. All rights reserved.
 * JiaoHu PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 * Created by tony on 15/12/18.
 */
package model
import (
	"os"
)

type ResFileWrapper interface {
	// Add File Attribute
	AddAttribute()
}

type ResFile struct {
	AbsoluteFilePath string `json:"-"`
	Uri              string `json:"uri"`
	Size             int64 `json:"size"`
	FileType         string `json:"fileType"`
}

func (f *ResFile) AddAttribute() {
	file, err := os.Open(f.AbsoluteFilePath)
	if err != nil {
		return
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		f.Size = stat.Size()
	}
}


// Construct ResFile
func ResFileFromFileName(absoluteFilePath string, uri string, fileType string) interface{} {
	var res ResFileWrapper
	resFile := ResFile{AbsoluteFilePath:absoluteFilePath, Uri:uri, FileType:fileType}

	switch fileType {
	case FileTypeImage:
		res = &ResImage{ResFile:resFile}
	case FileTypeAudio:
		res = &ResAudio{ResFile:resFile}
	case FileTypeVideo:
		res = &ResVideo{ResAudio:ResAudio{ResFile:resFile}}
	default:
		res = &resFile
	}
	res.AddAttribute()
	return res
}
