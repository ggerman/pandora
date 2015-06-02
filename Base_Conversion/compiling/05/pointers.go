package main

import "fmt"

// STRUCTS0 OMIT

type person struct {
  name string
  age int
  address string
}

type node struct {
  info person
  back, next *node
}

// STRUCTS0 OMIT

// ADDNODE1 OMIT

func add(pointer *node, data person) *node {
  var cell node
  cell.info = data
  cell.back, cell.next = pointer, nil
  newer := &cell
  if pointer == nil {
    pointer = newer
  } else {
    pointer.next = newer
    pointer = newer
  }
  return pointer
}

// ADDNODE1 OMIT

// DUMP OMIT

func dump(pointer *node) {
  if pointer != nil {
    fmt.Printf("Content: %p \n", *pointer)
    fmt.Println("--------------------")
    dump(pointer.back)
  }
}

// DUMP OMIT

// APPLICATION OMIT

func main() {
    var info person
    info.name, info.age, info.address = "Germán Alberto Giménez Silva", 38, "Nazario Medrano 394" 
    last := add(nil, info)
    info.name, info.age, info.address = "Juan Francisco Giménez Silva", 28, "Don Bosco 239" 
    last = add(last, info)
    info.name, info.age, info.address = "Ernestina Giménez Franco", 1, "Nazario Medrano 394"
    last = add(last, info)
    info.name, info.age, info.address = "María Jimena Franco", 34, "Nazario Medrano 394"
    last = add(last, info)

    dump(last)

}

// FINISH OMIT
