/* 
 * Copyright 2015 JiaoHu. All rights reserved.
 * JiaoHu PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 * Created by tony on 15/12/18.
 */
package model
import "github.com/disintegration/imaging"

type ResImage struct {
	ResFile
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (f *ResImage) AddAttribute(){
	f.ResFile.AddAttribute()
	// Image, Width & Height
	width, height, err := GetImageSize(f.AbsoluteFilePath)
	if err == nil {
		f.Width = width
		f.Height = height
	}
}

// Use imaging library to Get Image Size
func GetImageSize(name string) (int, int, error) {
	img, err := imaging.Open(name)
	if err == nil {
		srcBounds := img.Bounds()
		return srcBounds.Max.X, srcBounds.Max.Y, nil
	}else {
		return 0, 0, err
	}
}
