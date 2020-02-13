package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
	"log"
	"os"
)

var listall = `query{
 	listAnimals{
		items{
			id
			Name
			Genus
		}
	}
}`

var getHendrix = `query{
  getAnimals(id:"e8005160-47fd-4400-8cc5-1c580ffd09ad") {
    id
    Name
    Genus
  }
}`

var rosie = `mutation{
  createAnimals(input: {Name:"Rosie",Genus:"Dog"}){
    id
    Name
    Genus
  }
}`

/*
	Name
	Genus

*/
/***********************************************************************************************************************
	Types and Interfaces
***********************************************************************************************************************/

type Animals struct {
	Id    string `json:"id"`
	Name  string `json:"Name"`
	Genus string `json:"Genus"`
}
type data struct {
	Items []Animals `json:"items"`
}

type Allanimals struct {
	listAnimals map[string]data
}

/***********************************************************************************************************************
	Constants
***********************************************************************************************************************/
var _EndPoint string
var _API_KEY string

/***********************************************************************************************************************
	Init
***********************************************************************************************************************/
func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Errorf("%s", e)
	}
	_EndPoint = os.Getenv("END_POINT")
	_API_KEY = os.Getenv("API_KEY")
}

/***********************************************************************************************************************
	Set Request Headers
***********************************************************************************************************************/
func setRequestHeaders(r *graphql.Request) {
	r.Header.Set("Cache-Control", "no-cache")

	r.Header.Set("Content-Type", "application/json")
	r.Header.Add("x-api-key", _API_KEY)
}

/***********************************************************************************************************************
	Hendrix
***********************************************************************************************************************/
func Hendrix(c *graphql.Client) {

	req := graphql.NewRequest(getHendrix)
	setRequestHeaders(req)

	ctx := context.Background()

	var MyAnimals map[string]interface{}
	//	var MyAnimals 	Allanimals
	if err := c.Run(ctx, req, &MyAnimals); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MyAnimals is of Type : %T \n", MyAnimals)

	//	v := MyAnimals["getAnimals"].(map[string]interface{})
	//	fmt.Printf("v is of Type : %T \n",v)

	//	for key, val := range v {
	//		fmt.Printf("Key : %s  - Value : %s\n",key,val)
	//	}
}

/***********************************************************************************************************************
	Coerse the data
***********************************************************************************************************************/
func getData([]interface{}) {

}

/***********************************************************************************************************************
	All Animals
***********************************************************************************************************************/
func AllMyAnimals(c *graphql.Client) {
	req := graphql.NewRequest(listall)
	setRequestHeaders(req)

	ctx := context.Background()

	var MyAnimals map[string]interface{}
	if err := c.Run(ctx, req, &MyAnimals); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MyAnimals is of Type : %T \n", MyAnimals)

	v := MyAnimals["listAnimals"].(map[string]interface{})
	fmt.Printf("v is of Type : %T \n", v)

	var items []interface{}
	items = v["items"].([]interface{})

	var thisItem map[string]interface{}

	for i := range items {

		thisItem = items[i].(map[string]interface{})
		fmt.Printf("Index = %d : id : %s Name : %s Genus : %s value =%T\n",
			i,
			thisItem["id"],
			thisItem["Name"],
			thisItem["Genus"],
			thisItem)

	}
}

/***********************************************************************************************************************
	Mutation
***********************************************************************************************************************/
func mutation(c *graphql.Client) {
	/*
	   	req := graphql.NewRequest(`
	          query ($key: String!) {
	              items (id:$key) {
	                  field1
	                  field2
	                  field3
	              }
	          }
	      `)

	      // set any variables
	      req.Var("key", "value")
	*/
	req := graphql.NewRequest(rosie)
	setRequestHeaders(req)

	ctx := context.Background()

	var MyAnimals map[string]interface{}
	if err := c.Run(ctx, req, &MyAnimals); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MyAnimals is of Type : %T \n", MyAnimals)

	v := MyAnimals["listAnimals"].(map[string]interface{})
	fmt.Printf("v is of Type : %T \n", v)

}

/***********************************************************************************************************************
	Main
***********************************************************************************************************************/
func main() {

	client := graphql.NewClient(_EndPoint)

	Hendrix(client)
	AllMyAnimals(client)

	mutation(client)
	fmt.Println("Back - what has happened ")

}
