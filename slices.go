//learning maps and slice in Golang
package main

import (
  "fmt"
)

func main() {
  secretGarden := map[string]int{
    "potato":  12,
    "tomato":  5,
    "carrot":  9,
    "corn": 3,
  }	
  secretGarden["corn"] = 30
  secretGarden["onions"] = 5
  fmt.Println(secretGarden)

  //another for loop but better and easier

  for i, v := range secretGarden {
    fmt.Println(i, v)


    }
  }


