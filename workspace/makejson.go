package main

import (
    "fmt"
    "encoding/json"
    "bufio"
    "os"
    "strings"
)

func main() {
    var my_map map[string]string
    var name string
    var address string
    
    consoleReader := bufio.NewReader(os.Stdin)
    
    my_map = make(map[string]string)
    
    fmt.Println("Enter your name:")
    name, _ = consoleReader.ReadString('\n')
    name = strings.Replace(name, "\n", "", 1)
    fmt.Println("Enter your adress:")
    address, _ = consoleReader.ReadString('\n')
    address = strings.Replace(address, "\n", "", 1)
    
    
    my_map[name] = address
    
    json_data, err_json := json.Marshal(my_map)
    
    if err_json != nil {
        fmt.Println("JSON build failed!")
    } else {
        fmt.Println("JSON data:")
        fmt.Println(json_data)
        fmt.Println("JSON string representation: ")
        fmt.Println(string(json_data))
    }
}