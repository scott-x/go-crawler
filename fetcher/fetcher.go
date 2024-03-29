package fetcher

import (
   "bufio"
   "fmt"
   "io"
   "io/ioutil"
   "log"
   "net/http"

   "golang.org/x/text/encoding"
   
   "golang.org/x/net/html/charset"
   "golang.org/x/text/encoding/unicode"
   "golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error){
   resp,err := http.Get(url)
   if err!=nil {
      return nil,err
   }

   defer resp.Body.Close()

   if resp.StatusCode!=http.StatusOK{
      return nil,fmt.Errorf("error: wrong status code %d",resp.StatusCode)
   }
   //解决乱码问题,gbk->utf8
   e := determinEcoding(resp.Body)
   utf8Reader :=transform.NewReader(resp.Body,e.NewDecoder())
   return ioutil.ReadAll(utf8Reader)
}

func determinEcoding(r io.Reader) encoding.Encoding{
   bytes,err :=bufio.NewReader(r).Peek(1024)
   if err!=nil{
      log.Printf("Fetcher error %v",err)
      return unicode.UTF8
   }
   e, _, _ :=charset.DetermineEncoding(bytes,"")
   return  e
}