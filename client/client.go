package client

import (
  "fmt"
  "net/http"
  "hash/crc32"
  "strings"
  "sort"
  "bytes"
  "io/ioutil"

)

type Pair struct {
  Key string
  Value int
}

type Response struct {
  Key string
  Value string
}


var serverNameMapping map[int]string
var serverNameHashingMapping map[string]int
var sortedResult PairList

func init() {
  serverNameMapping = make(map[int]string)
  serverNameMapping[0]="3000"
  serverNameMapping[1]="3001"
  serverNameMapping[2]="3002"
  sortedResult = hashServer()
}

func Client() {
  var inputArr[10] string
  inputArr[0]= "1/a"
  inputArr[1]= "2/b"
  inputArr[2]= "3/c"
  inputArr[3]= "4/d"
  inputArr[4]= "5/e"
  inputArr[5]= "6/f"
  inputArr[6]= "7/g"
  inputArr[7]= "8/h"
  inputArr[8]= "9/i"
  inputArr[9]= "10/j"
  client := &http.Client{}
//PUT request from user
  for i:=0; i<len(inputArr); i++ {
    arr := strings.Split(inputArr[i],"/")
    key := arr[0]
    value := arr[1]
    server := hashing(key)
    url := "http://localhost:"+server+"/keys/"+key+"/"+value
    jsonStr := []byte(`{}`)
    resp, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))

    resp.Header.Set("Content-Type", "application/json")
    resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    response, err := client.Do(resp)
      if err!=nil{
        panic(err)
      }
    fmt.Println(response.StatusCode)
    defer response.Body.Close()
 }

 //GET request from user
 var inputGetArr[10] string
 inputGetArr[0]= "1"
 inputGetArr[1]= "2"
 inputGetArr[2]= "3"
 inputGetArr[3]= "4"
 inputGetArr[4]= "5"
 inputGetArr[5]= "6"
 inputGetArr[6]= "7"
 inputGetArr[7]= "8"
 inputGetArr[8]= "9"
 inputGetArr[9]= "10"

 for i:=0; i<len(inputGetArr); i++ {
   key := inputGetArr[i]
   server := hashing(key)
   url := "http://localhost:"+server+"/keys/"+key
   resp, _ := http.Get(url)
   defer resp.Body.Close()
   contents,_ := ioutil.ReadAll(resp.Body)
   fmt.Printf("%s\n", string(contents))
 }
}

func hashing(keyData string) (port string) {
  hash := crc32.ChecksumIEEE
  keyHash := int(hash([]byte(keyData)))
  var serverPort string

  max := sortedResult[0].Value
  min := sortedResult[1].Value
  mid := sortedResult[2].Value

  if keyHash > max {
    serverPort = sortedResult[2].Key // key stored at min position
  }
  if keyHash > mid && keyHash < max || keyHash == max { // key stored at max position
    serverPort = sortedResult[0].Key
  }
  if keyHash > min && keyHash < mid || keyHash == mid{ // key stored at mid position
  serverPort = sortedResult[1].Key
  }
  if keyHash < min || keyHash == min {
    serverPort = sortedResult[2].Key
  }
 return serverPort
}

func hashServer() (PairList){
  serverNameHashingMapping = make(map[string]int)
  for _,value := range serverNameMapping {
    hash := crc32.ChecksumIEEE
    serverHash := int(hash([]byte(value)))
    serverNameHashingMapping[value] = serverHash

  }
  result := sortByValue(serverNameHashingMapping)
  return result
}

func sortByValue(mapToBeSorted map[string]int) PairList{
  pl := make(PairList, len(mapToBeSorted))
  i := 0
  for k, v := range mapToBeSorted {
    pl[i] = Pair{k, v}
    i++
  }
  sort.Sort(sort.Reverse(pl))
  return pl
}


type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
