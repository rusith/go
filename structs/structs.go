package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

type TestType struct {
	Name string `json:"name"`
	Id int `json:"id"`
}

func ConvertToJson() {
	testObj := &TestType{ Name: "Rusith" , Id: 15245 }
	j, err := json.MarshalIndent(*testObj, "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr,"%v\n", err)
		return
	}
	fmt.Printf("%s\n", j)

}

func ConvertFromJson() {
	j := `{                    
 		"name": "Rusith",   
 		"id": 15245         
	}`
	t := &TestType{}
	json.Unmarshal([]byte(j), t)
	fmt.Println(*t)
}