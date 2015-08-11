package main

import (
  "fmt"
  "flag"
  "net/http"
  "time"

  "github.com/Dacode45/mysite_go/app/common"
  "github.com/Dacode45/mysite_go/app/home"
  "github.com/Dacode45/mysite_go/app/user"

  "github.com/golang/glog"
  "github.com/gorilla/mux"
)

func main(){
  //glog flags
  flag.Parse()
  defer glog.Flush()

  router := mux.NewRouter()
  http.Handle("/", httpInterceptor(router))

  router.HandleFunc("/", home.GetHomePage).Methods("GET")
  router.HandleFunc("/user{_:/?}", user.GetHomePage).Methods("GET")

  router.HandleFunc("/user/view/{id:[0-9]+}", user.GetViewPage).Methods("GET")
  router.HandleFunc("/user/{id:[0-9]+}", user.GetViewPage).Methods("GET")

  fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
  http.Handle("/static/", fileServer)

  err := http.ListenAndServe(":8080", nil)

  if err != nil{
    fmt.Println(err)
  }


}

//Function to time the handling of request, and log each one.
//Only times Okay Status
func httpInterceptor(router http.Handler) http.Handler{
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request){
    startTime := time.Now()

    router.ServeHTTP(w, req)

    finishTime := time.Now()
    elapsedTime := finishTime.Sub(startTime)

    switch req.Method{
    case "GET":
      //TODO sometimes you don't want StatusOK, but we will for Now
      common.LogAccess(w, req, elapsedTime)
    case "POST":
      //check http.StatusCreated
    }
  })
}
