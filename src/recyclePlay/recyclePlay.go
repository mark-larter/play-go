package main

import "fmt"
import "time"
import "math/rand"

func main() {
  materials := make(map[string]bool)
  materials["aluminum"] = true
  materials["cheese"] = false
  materials["meat"] = false
  materials["nickel"] = true
  materials["paper"] = true
  materials["plastic"] = true
  materials["steel"] = true
  materials["wheat"] = false
  items := []string{"aluminum", "cheese", "diamond", "meat", "nickel", "paper", "pickle", "plastic", "steel"}

  fmt.Println("Enter a 64-bit integer for random seed:")
  var seed int64
  _, err := fmt.Scanf("%d", &seed)
  if (err != nil) {
    seed = 77115
  }
  rand.Seed(seed)
  recycleItems(items, materials)

  fmt.Println("All items have been proccessed.")
}

func recycleItems(items []string, materials map[string]bool) {

  errChan := make(chan error)
  successChan := make(chan string)
  remainingItems := len(items)

  for _, item := range items {
    go recycleItem(item, materials, successChan, errChan)
  }

  for {
    select {
    case item := <-successChan:
      fmt.Printf("%s has been recycled.\n", item)
    case err := <-errChan:
      fmt.Println(err)
    }

    remainingItems--
    if remainingItems == 0 {
      break;
    }
  }
}

func recycleItem(item string, materials map[string]bool, successChan chan string, errChan chan error)  {
  time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
  recyclable, ok := materials[item] 
  if (!ok) {
    errChan <- fmt.Errorf("%s is an unknown material.", item)
  } else if (recyclable) {
    successChan <- item
  } else {
    errChan <- fmt.Errorf("%s is not recyclable.", item)
  }
}