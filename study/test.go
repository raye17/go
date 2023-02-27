package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//viper.SetConfigFile("test.toml")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	glog.Errorf("failed to read config file")
	//	panic(err)
	//}
	//viper.Set("server.Name", "mux")
	//viper.Set("mysql.password", "raye001")
	//err = viper.WriteConfig()
	//if err != nil {
	//	panic(fmt.Errorf("fatal error config file:%v", err))
	//}
	//err = viper.ReadInConfig()
	//if err != nil {
	//	glog.Errorf("failed to read config file")
	//	panic(err)
	//}
	//fmt.Println(viper.Get("mysql.ip"))
	//viper.SetConfigFile("config")
	//viper.AddConfigPath(".")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	glog.Errorf("failed to read config")
	//	panic(err)
	//}
	//fmt.Println(viper.Get("Kind"))
	m := mux.NewRouter()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome")
	})
	http.ListenAndServe(":8888", m)
}
