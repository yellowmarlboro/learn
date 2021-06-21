package main

import (
	"bytes"
	"encoding/json"
	"io"

	//"encoding/json"
	"fmt"
	//"io"

	"github.com/go-playground/validator/v10"
)

var EnableDecoderUseNumber = false

// EnableDecoderDisallowUnknownFields is used to call the DisallowUnknownFields method
// on the JSON Decoder instance. DisallowUnknownFields causes the Decoder to
// return an error when the destination is a struct and the input contains object
// keys which do not match any non-ignored, exported fields in the destination.
var EnableDecoderDisallowUnknownFields = false

type A struct {
	X, Y int
}

func (a *A) Tes() int {
	return 100
}

type AI interface {
	Tes() int
}

type B struct {
	Z, T int
}

func (b *B) Ig() int {
	return 200
}

type BI interface {
	Ig() int
}

type CI interface {
	AI
	BI
}

type C struct {
	A
	B
	O int
}

type D struct {
	G CI
}

type Inner struct {
	A  *float64 `json:"a" validate:"required"`
	AA *int `json:"aa" validate:"required"`
}

type Outer struct {
	Inner []Inner `json:"inner" validate:"required,dive"`
}

type OOuter struct {
	Outer Outer `json:"outer" validate:"required"`
}


func BindJsonBody(body []byte, obj interface{}) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

func Validate(obj interface{}) error {
	return validate(obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}

func validate(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		return err
	}
	return nil
}

//type AppInfo struct {
//	AppID  string `json:"AppId" validate:"required"`
//	AppVer string `json:"AppVer" validate:"required"`
//}
//
//type BaseInfo struct {
//	BizId  int `json:"bizid" validate:"required"`
//	UserId int `json:"userid" validate:"required"`
//}
//
//type FeeRequest struct {
//	Data BaseInfo `json:"Data" validate:"required,dive"`
//	Info AppInfo  `json:"Info" validate:"required"`
//}

//
//var Validator StructValidator = &defaultValidator{}
//
//func validate(obj interface{}) error {
//	if Validator == nil {
//		return nil
//	}
//	return Validator.ValidateStruct(obj)
//}
//
//func decodeJSON(r io.Reader, obj interface{}) error {
//	decoder := json.NewDecoder(r)
//	if EnableDecoderUseNumber {
//		decoder.UseNumber()
//	}
//	if EnableDecoderDisallowUnknownFields {
//		decoder.DisallowUnknownFields()
//	}
//	if err := decoder.Decode(obj); err != nil {
//		return err
//	}
//	return validate(obj)
//}

func main() {
	//c := C{
	//	//A: A{},
	//	//B: B{},
	//	O: 300,
	//}
	//
	//d := D{G: &c}
	//a := A{1, 2}
	//b := B{3, 4}
	//var c C
	//c.B = new(B)
	//c.AA.X = 123
	//c.AA.Y = 456
	//c.B.Z = 789
	//c.B.T = 000
	//outName := jsoniter.Config{TagKey: "outname"}.Froze()
	//
	//v := new(User)
	//s := `{"user_id": 111}`
	//s1 := []byte(s)
	//err := binding.JSON.BindBody(s1, v)
	//fmt.Println(err)
	//fmt.Printf("%+v", v)
	//ret, _ := outName.Marshal(v)
	//fmt.Printf("%s", ret)
	o := new(OOuter)
	s := `{"outer": {"inner": [{"a": 1.22222222222222, "aa": 0}]}}`

	body := []byte(s)
	//decoder := json.NewDecoder(bytes.NewReader(body))
	//err := decoder.Decode(o)
	//err := binding.JSON.BindBody(body, o)
	//s := new(FeeRequest)
	//s.Info.AppID=""
	//validate := validator.New()
	//err = validate.Struct(*o)
	err := BindJsonBody(body, o)
	if err != nil {
		fmt.Println(err)
	}
	r, err := json.Marshal(o)
	//cal := *o.Outer.Inner[0].A / 2
	fmt.Printf("data is %s", r)
	fmt.Printf("data is %s", *o.Outer.Inner[0].A)
}
