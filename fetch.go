package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main(){
  for _,url := range os.Args[1:]{
    resp,err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
    }
    b, err := ioutil.ReadAll(resp.Body)
    status := ioutil.ReadAll(resp.Status)
    resp.Body.Close()
    resp.Status.Close()
    if err != nil {
      fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
      os.Exit(1)
    }
    fmt.Printf("%s",b)
    fmt.Printf("%s",status)
  }
}