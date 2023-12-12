package main


import (
    "encoding/json"
    "fmt"
    "os"
)

type Person struct {
    Name string `json:"n"`
    Age int     `json:"s"`
}

func main(){
    person := Person{Name: "Tribix", Age: 30}

    result, err := json.Marshal(person)

    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println(string(result))
    //Variable use to enconder of JSON
    err = json.NewEncoder(os.Stdout).Encode(person)

    if err != nil {
        fmt.Println(err)
    }
    
    JsonTrue := []byte(`{"n": "NewPerson", "a": 100}`)

    var newPerson Person

    err = json.Unmarshal(JsonTrue, &newPerson)

    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(newPerson)
}
