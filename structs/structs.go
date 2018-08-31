package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

type TestType struct {
	Name string `json:"name"` // Case insensitive, so theses tags are redundant
	Id int `json:"id"`
}


type  TestEmbed struct {
	TestType
}

func (t TestEmbed) GetName() string {
	fmt.Println("Calling non pointer")
	return t.Name
}

func (t *TestEmbed) GetNameOne() string {
	fmt.Println("Calling pointer")
	return t.Name;
}


func TestRun() {
	obj := TestEmbed{ TestType{ Name: "Rusith", Id: 300}}
	fu := obj.GetName

	p := &obj;
	p.GetName()
	p.GetNameOne()
	fmt.Println(fu())
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


type Point struct { X, Y float64 }

func (p *Point) Add(q Point) *Point { return &Point{ p.X + q.X , p.Y + q.Y } }
func (p *Point) Subs(q Point) *Point { return &Point{  p.X - q.X , p.Y - q.Y }}
