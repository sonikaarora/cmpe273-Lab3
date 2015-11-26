package server

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

type Keyvalue struct {
  Key string
  Value string
}

type ResponseStruct struct {
	Response [] Keyvalue
}

var keyValueCacheServer1 map[string]string = make(map[string]string)
var keyValueCacheServer2 map[string]string = make(map[string]string)
var keyValueCacheServer3 map[string]string = make(map[string]string)

func createServerInstance(port string, mux http.Handler) {

   http.ListenAndServe("localhost:"+port, mux)
}

func Server() {
  fmt.Println("starting server..........")

  mux1 := httprouter.New()
  mux1.PUT("/keys/:key/:value", updateKeyValueHandlerServer1)
  mux1.GET("/keys/:key_id",getKeyHandlerServer1)
  mux1.GET("/keys",getAllKeysHandlerServer1)

  mux2 := httprouter.New()
  mux2.PUT("/keys/:key/:value", updateKeyValueHandlerServer2)
  mux2.GET("/keys/:key_id",getKeyHandlerServer2)
  mux2.GET("/keys",getAllKeysHandlerServer2)

  mux3 := httprouter.New()
  mux3.PUT("/keys/:key/:value", updateKeyValueHandlerServer3)
  mux3.GET("/keys/:key_id",getKeyHandlerServer3)
  mux3.GET("/keys",getAllKeysHandlerServer3)

  go createServerInstance("3000",mux1)
  go createServerInstance("3001",mux2)
  go createServerInstance("3002",mux3)
}
