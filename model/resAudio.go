/* 
 * Copyright 2015 JiaoHu. All rights reserved.
 * JiaoHu PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 * Created by tony on 15/12/18.
 */
package model
import (
	"strings"
	"os/exec"
	"strconv"
)

type ResAudio struct {
	ResFile
	Duration float64 `json:"duration"`
}

func (f *ResAudio) AddAttribute(){
	f.ResFile.AddAttribute()
	// Get File Duration
	duration, err := GetMediaDuration(f.Name)
	if err == nil {
		f.Duration = duration
	}
}

// Use Command Line "ffprobe" to Get media duration
func GetMediaDuration(filePath string) (float64, error) {
	out, err := exec.Command("sh", "-c", "ffprobe -i " + filePath + " -show_entries format=duration -v quiet -of csv=\"p=0\"").Output()
	if err == nil {
		durationStr := strings.Trim(string(out), "\n")
		duration, _ := strconv.ParseFloat(durationStr, 64)
		return duration, nil
	}else {
		return 0, err
	}
}
