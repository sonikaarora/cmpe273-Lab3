package server

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "encoding/json"
)

func updateKeyValueHandlerServer1(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
  var keyvalue Keyvalue
  keyvalue.Key = p.ByName("key")
  keyvalue.Value = p.ByName("value")
  keyValueCacheServer1[keyvalue.Key] = keyvalue.Value
  rw.WriteHeader(200) // Status code for success
  fmt.Fprint(rw)
}
func updateKeyValueHandlerServer2(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
  var keyvalue Keyvalue
  keyvalue.Key = p.ByName("key")
  keyvalue.Value = p.ByName("value")
  keyValueCacheServer2[keyvalue.Key] = keyvalue.Value
  rw.WriteHeader(200) // Status code for success
  fmt.Fprint(rw)
}
func updateKeyValueHandlerServer3(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
  var keyvalue Keyvalue
  keyvalue.Key = p.ByName("key")
  keyvalue.Value = p.ByName("value")
  keyValueCacheServer3[keyvalue.Key] = keyvalue.Value
  rw.WriteHeader(200) // Status code for success
  fmt.Fprint(rw)
}

func getKeyHandlerServer1(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {

 keyId := p.ByName("key_id")
 var keyvalue Keyvalue
 value := keyValueCacheServer1[keyId]
 keyvalue.Key = keyId
 keyvalue.Value = value
 jsonResponse, _ := json.Marshal(&keyvalue)
 rw.Header().Set("Content-Type", "application/json")
 rw.WriteHeader(200) // Status code for success
 fmt.Fprintf(rw, "%s", jsonResponse)
}

func getKeyHandlerServer2(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
 keyId := p.ByName("key_id")
 var keyvalue Keyvalue
 value := keyValueCacheServer2[keyId]
 keyvalue.Key = keyId
 keyvalue.Value = value
 jsonResponse, _ := json.Marshal(&keyvalue)
 rw.Header().Set("Content-Type", "application/json")
 rw.WriteHeader(200) // Status code for success
 fmt.Fprintf(rw, "%s", jsonResponse)
}

func getKeyHandlerServer3(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
 keyId := p.ByName("key_id")
 var keyvalue Keyvalue
 value := keyValueCacheServer3[keyId]
 keyvalue.Key = keyId
 keyvalue.Value = value
 jsonResponse, _ := json.Marshal(&keyvalue)
 rw.Header().Set("Content-Type", "application/json")
 rw.WriteHeader(200) // Status code for success
 fmt.Fprintf(rw, "%s", jsonResponse)
}

func getAllKeysHandlerServer1(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
var Response1 ResponseStruct
 for key,value := range keyValueCacheServer1 {
   temp := Keyvalue{}
   temp.Key = key
   temp.Value = value
   Response1.Response = append(Response1.Response,temp)
 }
 jsonResponse, _ := json.MarshalIndent(Response1, "", "\t")
 rw.Header().Set("Content-Type", "application/json")
 rw.WriteHeader(200)
 fmt.Fprintf(rw, "The JSON response received is as follows %s", jsonResponse)

}

func getAllKeysHandlerServer2(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
var Response2 ResponseStruct
  for key,value := range keyValueCacheServer2 {
    temp := Keyvalue{}
    temp.Key = key
    temp.Value = value
    Response2.Response = append(Response2.Response,temp)
  }
  jsonResponse, _ := json.MarshalIndent(Response2, "", "\t")
  rw.Header().Set("Content-Type", "application/json")
  rw.WriteHeader(200)
  fmt.Fprintf(rw, "The JSON response received is as follows %s", jsonResponse)
}

func getAllKeysHandlerServer3(rw http.ResponseWriter, request *http.Request, p httprouter.Params) {
var Response3 ResponseStruct
  for key,value := range keyValueCacheServer3 {
    temp := Keyvalue{}
    temp.Key = key
    temp.Value = value
    Response3.Response = append(Response3.Response,temp)
  }
  jsonResponse, _ := json.MarshalIndent(Response3, "", "\t")
  rw.Header().Set("Content-Type", "application/json")
  rw.WriteHeader(200)
  fmt.Fprintf(rw, "The JSON response received is as follows %s", jsonResponse)
}
