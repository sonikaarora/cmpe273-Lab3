package main
import (
  "fmt"
  "consistentHashing/server"
)

func main() {
 go server.Server()

 var input int
 fmt.Scanln(&input)


}
