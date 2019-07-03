package main

import (
	"net/http"
	"fmt"
)

func Fetch(url string) ([]byte, error){
   resp, err :=http.Get(url)
   if err!=nil {
   	return nil, err
   }   
   defer resp.Body.Close()

   if resp.StatusCode != resp.StatusOK{
   	return nil, fmt.Errorf("wrong status code %d", resp.StatusCode)
   }

   
}