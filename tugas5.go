package main

import "fmt"
import "math/rand"

type pelajar struct {
 nama string
 umur int
}

func main(){
var n1,n2,n3 pelajar;
n1.nama = "Ahmad"
n1.umur = rand.Intn(100)
n2.nama = "Abi"
n2.umur = rand.Intn(100)
n3.nama = "Zaka"
n3.umur = rand.Intn(100)
n1.namapelajar()
n2.namapelajar()
n3.namapelajar()
}



func (s pelajar)namapelajar()  {
  fmt.Println(s.nama, s.umur)

}
