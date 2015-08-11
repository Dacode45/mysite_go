package common

import (
  "log"
  "runtime"

  "github.com/golang/glog"
)

//Logs Errors in specific way
//0 = info
//1 = warning
//2 = error
//3 - fatal
const(
  Info = 0
  Warning = 1
  Error = 2
  Fatal = 3
)
func CheckError(err error, level int){
  if err != nil{
    var stack [4096]byte
    //formats the stack trace of gorutine into buf and returns the number of bytes writtent to
    runtime.Stack(stack[:], false)
    log.Printf("%q\n%s\n", err, stack[:])

    switch level {
    case Info:
      glog.Infoln("%q\n%s\n", err)
    case Warning:
      glog.Warning("%q\n%s\n", err)
    case Error:
      glog.Errorln("%q\n%s\n", err)
    case Fatal:
      glog.Fatalln("%q\n%s\n", err)
    }

    glog.Flush()
  }
}
