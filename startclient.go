package main
import (
  "fmt"
  "consistentHashing/client"
)

func main() {
 go client.Client()

 var input int
 fmt.Scanln(&input)


}
