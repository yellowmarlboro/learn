package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin/binding"
)

type UserAddIdRequest struct {
	Name string `json:"name" binding:"required"`
}

type UserAddIdResponse struct {
	UserAddIdRequest
	Id string `json:"id" binding:"required"`
	Age int   `json:"age" binding:"required,lt=1e00"`
}

func main()  {
	resp := `{"code":"000000","msg":"ok","data":{"name":"a","id":"91d747fc-e755-4979-b731-3480b6b2b6d6"}}`
	//resp := `{"name":"123","id":"91d747fc-e755-4979-b731-3480b6b2b6d6","age":120}`
	resp1 := []byte(resp)
	//ld1 := gojsonschema.NewStringLoader(resp)
	//st1 := UserAddIdResponse{Id: "1"}
	st1 := new(UserAddIdResponse)
	//ld2 := gojsonschema.NewGoLoader(st1)
	//result, err := gojsonschema.Validate(ld1, ld2)
	//fmt.Println(result, err)
	//proto.Merge()

	err := json.Unmarshal(resp1, st1)
	fmt.Println("1")
	bs := binding.JSON
	err = bs.BindBody(resp1, st1)
	fmt.Println(err)

	decoder := json.NewDecoder(bytes.NewReader(resp1))
	err = decoder.Decode(st1)
	fmt.Println(err)

	//r1 := new(UserAddIdResponse)
	//r2 := UserAddIdResponse{}
	//validate := validator.New()
	//err := json.Unmarshal(resp1, r1)
	//ret := json.Valid(resp1)
	//fmt.Println(ret)
	//err = json.Unmarshal(resp1, &r2)
	//fmt.Println(err)
}
