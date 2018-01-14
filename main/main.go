package main

import (
	"encoding/json"
	"github.com/pingguoxueyuan/school_suggestion/logic"
	"net/http"
	"github.com/unrolled/render"
	"fmt"
	"time"
)

var (
	renderHtml *render.Render
)

func init() {
	option := render.Options{
		Directory: "views",
		Extensions: []string{".tmpl", ".html"},
	}
	renderHtml = render.New(option)
}

func handleIndex(w http.ResponseWriter, r* http.Request) {
	renderHtml.HTML(w, http.StatusOK, "index", nil)
}

func responseSuccess(w http.ResponseWriter, data interface{}) {
	m := make(map[string]interface{}, 16)
	m["code"] = 0
	m["message"] = "sucess"
	m["data"] = data

	result, err := json.Marshal(m)
	if err != nil {	
		return
	}

	w.Write(result)
}

func handleSearch(w http.ResponseWriter, r* http.Request) {
	r.ParseForm()
	keyword := r.FormValue("keyword")

	start := time.Now().UnixNano()
	//schools := logic.SearchV2(keyword, 16)
	schools := logic.Search(keyword, 16)
	end := time.Now().UnixNano()
	fmt.Printf("暴力:keyword:%s result:%d cost:%d us\n", keyword, len(schools), (end - start)/1000)
	responseSuccess(w, schools)
}

func main() {

	err := Init()
	if err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}

	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/school/search", handleSearch)
	err = http.ListenAndServe(":8080", nil)
	if err !=nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
}