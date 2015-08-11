package common

import (
  "fmt"
  "time"
  "strings"
  "strconv"
  "net/http"

  "github.com/golang/glog"
)

const(
  logFile = "access_log.txt"
)

type accessLog struct {
  ip, method, uri, protocol, host string
  finishTime time.Time
  elapsedTime time.Duration
}
//Writes the duration of each request to the log
//Writes IP, Method, URI, Protocol, Host, ElapsedTime
func LogAccess(w http.ResponseWriter, req *http.Request, finishTime time.Time, duration time.Duration){
  clientIP := req.RemoteAddr

  //don't need port
  if colon := strings.LastIndex(clientIP, ":"); colon != -1{
    clientIP = clientIP[:colon]
  }

  record := &accessLog{
    ip: clientIP,
    method: req.Method,
    uri: req.RequestURI,
    protocol: req.Proto,
    host: req.Host,
    finishTime: finishTime,
    elapsedTime: duration,
  }

  writeAccessLog(record)
}

func writeAccessLog(r *accessLog){
  logRecod := fmt.Sprintf("(%v) [%v] %v %v: %v, host: %v (load time: %v seconds)",r.finishTime, r.ip, r.protocol, r.method, r.uri, r.host, strconv.FormatFloat(r.elapsedTime.Seconds(), 'f', 5, 64) )
  glog.Infoln(logRecod)
  glog.Flush()
}
