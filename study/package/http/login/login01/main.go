package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Sum struct {
	Number []int `json:"number"`
}

func Compare(data User) Respond {
	var result Respond
	if data.Token != "TOKEN" {
		result.Code = 404
	} else if data.Username == "admin" && data.Password == "123456" {
		result.Code = 200
		result.Session = "success"
	} else {
		result.Code = 400
		result.Information = "wrong username or password"
	}
	return result
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Respond struct {
	Code        int32  `json:"code,omitempty"`
	Information string `json:"info,omitempty"`
	Session     string `json:"session,omitempty"`
	Result      int32  `json:"result,omitempty"`
}

func login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	con, _ := ioutil.ReadAll(r.Body)
	var data User
	err := json.Unmarshal(con, &data)
	if err != nil {
		return
	}
	con01, _ := json.Marshal(Compare(data))
	w.Write(con01)
}
func Get_sum(writer http.ResponseWriter, request *http.Request) {
	sum := 0
	defer request.Body.Close()
	con, _ := ioutil.ReadAll(request.Body)
	var data Sum
	json.Unmarshal(con, &data)
	for i := range data.Number {
		sum = sum + data.Number[i]
	}
	var result Respond
	result.Result = int32(sum)
	con, _ = json.Marshal(result)
	writer.Write(con)
}
func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/sum", Get_sum)
	http.ListenAndServe("localhost:9999", nil)

}
