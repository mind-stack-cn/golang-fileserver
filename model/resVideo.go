/* 
 * Copyright 2015 JiaoHu. All rights reserved.
 * JiaoHu PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 * Created by tony on 15/12/18.
 */
package model
import (
	"os/exec"
	"strings"
	"path"
	"log"
	"fmt"
	"bytes"
)

const thumbExt = ".jpg"

type ResVideo struct {
	ResAudio
	Thumbnail interface{} `json:"thumbnail"`
}

func (f *ResVideo) AddAttribute() {
	// add audio attribute
	f.ResAudio.AddAttribute()

	// Thumbnail path
	thumbNailPath := strings.TrimSuffix(f.Name, path.Ext(f.Name)) + thumbExt
	thumbNailUri := strings.TrimSuffix(f.Uri, path.Ext(f.Uri)) + thumbExt
	// Generate default thumbnail
	if err := GetVideoThumbnail(f.Name, thumbNailPath); err == nil {
		f.Thumbnail = ResFileFromFileName(thumbNailPath, thumbNailUri, FileTypeImage)
	}else{
		log.Print("GetVideoThumbnail Error: ", err.Error())
	}
}

// Use Command Line "ffmpeg" to Get media duration
func GetVideoThumbnail(filePath string, thumbNailPath string) error {
	cmd:= exec.Command("/bin/sh", "-c", "ffmpeg -i " + filePath + " -ss 00:00:01.000 -vframes 1 " + thumbNailPath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	fmt.Println("Result: " + out.String())
	return nil
}
