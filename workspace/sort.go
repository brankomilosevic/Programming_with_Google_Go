package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
  "sync"
)

func slice_sort(wg *sync.WaitGroup, is []int) {
  fmt.Println(is)
  sort.Ints(is)
  wg.Done()
}

func main() {
  var numbers []int
  var wg sync.WaitGroup
  var sorted []int

  fmt.Println("Enter integer list for sorting:")
  buff := bufio.NewReader(os.Stdin)
  input_str, _, _ := buff.ReadLine()
  tokens := strings.Split(string(input_str), " ")

  for _, str := range tokens {
    n, _ := strconv.Atoi(str)
    numbers = append(numbers, n)
  }

  size := len(numbers) / 4

  // not so elegant ...
  wg.Add(1)
  go slice_sort(&wg, numbers[0*size:1*size])
  wg.Add(1)
  go slice_sort(&wg, numbers[1*size:2*size])
  wg.Add(1)
  go slice_sort(&wg, numbers[2*size:3*size])
  wg.Add(1)
  go slice_sort(&wg, numbers[3*size:])

  wg.Wait()

  slice_1 := numbers[0*size : 1*size]
  slice_2 := numbers[1*size : 2*size]
  slice_3 := numbers[2*size : 3*size]
  slice_4 := numbers[3*size:]

  // definitelly not so elegant, but I am focused on concurrency for the moment
  // it should be some kind of merge procedure
  for k := 0; k < len(numbers); k++ {
    i := 0
    j := 0
    if len(slice_1) != 0 {
      i = slice_1[0]
      j = 1
    }
    if len(slice_2) != 0 {
      if j == 0 {
        i = slice_2[0]
        j = 2
      } else {
        if slice_2[0] < i {
          i = slice_2[0]
          j = 2
        }
      }
    }
    if len(slice_3) != 0 {
      if j == 0 {
        i = slice_3[0]
        j = 3
      } else {
        if slice_3[0] < i {
          i = slice_3[0]
          j = 3
        }
      }
    }
    if len(slice_4) != 0 {
      if j == 0 {
        i = slice_4[0]
        j = 4
      } else {
        if slice_4[0] < i {
          i = slice_4[0]
          j = 4
        }
      }
    }

    sorted = append(sorted, i)

    switch j {
    case 1:
      slice_1 = slice_1[1:]
    case 2:
      slice_2 = slice_2[1:]
    case 3:
      slice_3 = slice_3[1:]
    case 4:
      slice_4 = slice_4[1:]
    }
  }

  fmt.Println(sorted)
}