package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Data struct {
	psnInsuMgtEid string
	psnInsuRltsId string
}

type Ret struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Data    []Data
}

// queryPsnInsuByCertno 两定信息-人员参保
func queryPsnInsuByCertno(w http.ResponseWriter, r *http.Request) {
	// 判断是否multipart方式
	isMultipart := false
	fmt.Println(r.Header)
	for _, v := range r.Header["Content-Type"] {
		if strings.Index(v, "multipart/form-data") != -1 {
			isMultipart = true
		}
	}
	// 解析Body
	if isMultipart == true {
		r.ParseMultipartForm(128)
	} else {
		r.ParseForm()
	}

	// ret := r.PostForm
	// fmt.Println(ret)

	data := Data{psnInsuMgtEid: "a", psnInsuRltsId: "b"}
	ret := new(Ret)
	ret.Code = 0
	ret.Message = "成功"
	ret.Type = "success"

	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	// fmt.Println(ret)

	retJson, _ := json.Marshal(ret)
	// fmt.Println(retJson)
	io.WriteString(w, string(retJson))

}

func main() {
	http.HandleFunc("/api/publicQueryInsu/queryPsnInsuByCertno", queryPsnInsuByCertno)
	http.ListenAndServe("0.0.0.0:80", nil)
}
