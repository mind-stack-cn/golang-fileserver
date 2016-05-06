/* 
 * Copyright 2015 JiaoHu. All rights reserved.
 * JiaoHu PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 * Created by tony on 15/12/17.
 */
package handle

import (
	"net/http"
	"fmt"
	"io"
	"strings"
	"path/filepath"
	"os"
	"log"
	"github.com/mind-stack-cn/golang-fileserver/model"
	"encoding/json"
	"github.com/satori/go.uuid"
)

// Handler Upload File Request
// Save It, Return saved file info
func FileUpload(dir string, w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var fileArray []interface{}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		randomUUId := uuid.NewV4()
		paSplit := strings.Split(randomUUId.String(), "-")
		// File Path
		relatedFileDir := paSplit[0] + "/" + paSplit[1] + "/" + paSplit[2] + "/" + paSplit[3] + "/"
		// File Name
		fileName := paSplit[4] + filepath.Ext(part.FileName())

		// Create File Dir if not
		var fileDir string
		if dir != "." {
			fileDir = dir + relatedFileDir
		}
		errPath := os.MkdirAll(fileDir, 0777)
		if errPath != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create File
		ff := fileDir + fileName
		dst, err := os.Create(ff)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Print("receive file successfully! fileName:" + part.FileName())

		// append resFile to response Array
		fileArray = append(fileArray,  model.ResFileFromFileName(dst.Name(), "/" + relatedFileDir + fileName, r.URL.Query().Get("fileType")))

		if len(fileArray) >= 10 {
			// 最多上传10个文件
			break;
		}
	}

	res := model.Response{Header:model.Header{Code:model.ServerSuccessCode, Description:model.ServerSuccessDesc}, Data:fileArray}

	// Generate Json
	jsonByte, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write Response
	fmt.Fprint(w, string(jsonByte))
	return
}
