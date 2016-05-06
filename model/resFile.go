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
	Name     string `json:"-"`
	Uri      string `json:"uri"`
	Size     int64 `json:"size"`
	FileType string `json:"fileType"`
}

func (f *ResFile) AddAttribute() {
	file, err := os.Open(f.Name)
	if err != nil {
		return
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		f.Size = stat.Size()
	}
}


// Construct ResFile
func ResFileFromFileName(name string, uri string, fileType string) interface{} {
	var res ResFileWrapper
	resFile := ResFile{Name:name, Uri:uri, FileType:fileType}

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
