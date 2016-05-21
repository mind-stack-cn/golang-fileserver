/* 
 * Copyright 2015 JiaoHu. All rights reserved.
 * JiaoHu PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 * Created by tony on 15/12/17.
 */
package handle

import (
	"log"
	"github.com/disintegration/imaging"
	"os"
	"strings"
	"path"
	"strconv"
	"regexp"
	"net/http"
	"path/filepath"
)

// Handler File Download Request
// Crop for Image if Query Parameter contains width or height
func FileDownload(dir string, w http.ResponseWriter, r *http.Request) {
	var forTest bool
	urlPath := r.URL.Path
	if urlPath == "/" {
		urlPath = urlPath + "/index.html"
		// Redirect test dir
		dir, _ = filepath.Abs(".")
		dir = dir + "/"
		forTest = true
	} else if strings.HasPrefix(urlPath, "/test") {
		if strings.Compare(urlPath, "/test") == 0 {
			urlPath = urlPath + "/index.html"
		}
		if strings.Compare(urlPath, "/test/") == 0 {
			urlPath = urlPath + "index.html"
		}
		// Redirect test dir
		dir, _ = filepath.Abs(".")
		dir = dir + "/"
		forTest = true
	}

	fileAbsPath := path.Clean(dir + urlPath)

	if !forTest {
		r.Header.Del("If-Modified-Since")
		width, height := CheckThumbParameters(r.URL.Query().Get("width"), r.URL.Query().Get("height"))
		fileAbsPath = GenerateThumbIfNeed(fileAbsPath, width, height)
	}

	log.Printf("downloading file %s", path.Clean(fileAbsPath))
	http.ServeFile(w, r, fileAbsPath)
}

// Check If Query Parameter contains Thumb parameter
// return width,height in Query Parameters, otherwise return -1,-1
func CheckThumbParameters(width string, height string) (int, int) {
	matched, err := regexp.MatchString("^[1-9][0-9]*$", width)
	if (err == nil) && matched {
		matched, err := regexp.MatchString("^[1-9][0-9]*$", height)
		if (err == nil) && matched {
			iWidth, _ := strconv.Atoi(width)
			iHeight, _ := strconv.Atoi(height)
			return iWidth, iHeight
		}
	}
	return -1, -1
}

// Generate Thumbnail if not exist
func GenerateThumbIfNeed(fileAbsPath string, width int, height int) string {
	if (width == -1) || (height == -1) {
		return fileAbsPath
	}

	extension := path.Ext(fileAbsPath)
	name := strings.TrimSuffix(fileAbsPath, extension)
	thumbNailPath := name + "_" + strconv.Itoa(width) + "_" + strconv.Itoa(height) + extension
	// Generate Thumbnail If not exist
	if _, err := os.Stat(thumbNailPath); os.IsNotExist(err) {
		// path/to/whatever exists
		log.Print("generate thumbnail:" + fileAbsPath)
		img, err := imaging.Open(fileAbsPath)
		if err == nil {
			dstImage := imaging.Thumbnail(img, width, height, imaging.CatmullRom)

			// save the combined image to file
			err := imaging.Save(dstImage, thumbNailPath)
			if err == nil {
				return thumbNailPath
			} else {
				log.Print("imaging.Save error:")
			}
		} else {
			log.Print("imaging.Open error:")
		}
	} else {
		log.Print("thubnail has already exist")
		return thumbNailPath
	}
	return fileAbsPath
}

