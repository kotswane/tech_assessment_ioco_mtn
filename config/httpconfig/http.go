package httpconfig

import (
	 "fmt"
	 "os"
)

func Init() string{
	var ep = os.Getenv("ROBOTS_ENDPOINT")
	return ep
}

func checkErr(err error) {
  if err != nil {
    fmt.Print(err.Error())
  }
}