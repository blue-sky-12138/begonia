// Time : 2020/9/27 17:35
// Author : Kieran

// bench
package bench

import (
	"fmt"
	"github.com/hamba/avro"
	"testing"
)

// avro_test.go something

//func BenchmarkLinkedinEncode(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		linkedinEncode()
//	}
//}
//
//func BenchmarkLinkedinDecode(b *testing.B){
//	b.ResetTimer()
//	for i:=0;i<b.N;i++{
//		linkedinDecode()
//	}
//}

//func BenchmarkHambaEncode(b *testing.B) {
//	hambaEncode()
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		hambaEncode()
//	}
//}

//func BenchmarkHambaDecode(b *testing.B){
//	hambaDecode()
//	b.ResetTimer()
//	for i:=0;i<b.N;i++{
//		hambaDecode()
//	}
//}

func TestLinkedinEncode(t *testing.T) {
	linkedinEncode()
}

func TestLinkedinDecode(T *testing.T) {
	linkedinDecode()
}

func TestHambaEncode(t *testing.T) {
	hambaEncode()
}

func TestHambaDecode(t *testing.T) {
	hambaDecode()
}

func Test(t *testing.T) {
	sIn := `{"type":"string","name":"serviceName"}`
	c, _ := avro.Parse(sIn)
	res, _ := avro.Marshal(c, "hello")
	fmt.Println(res)

	sOut := `{"type":"array","name":"test","items":{
	"type":"record",
    "name":"test1",
    "fields":[
	{"type":"string","name":"fun"},
    {"type":"string","name":"test1"}
]
}}`

	type test struct {
		Fun   string `avro:"fun"`
		Test1 string `avro:"test1"`
	}

	var tt []test
	tt = []test{{Fun: "s1", Test1: "t1"}, {Fun: "s2", Test1: "t2"}}

	o, err := avro.Parse(sOut)
	if err != nil {
		panic(err)
	}
	b, err := avro.Marshal(o, tt)
	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	var ttt []test
	err = avro.Unmarshal(o, b, &ttt)
	if err != nil {
		panic(err)
	}

	fmt.Println(ttt)
	//var f =  func(v interface{}){
	//	fmt.Println(reflect2.TypeOf(v).Kind())
	//}
	//f(map[string]interface{}{})

}

type rFun struct {
	Name      string `avro:"name"`
	InSchema  string `avro:"inSchema"`
	OutSchema string `avro:"outSchema"`
}

func TestByte(t *testing.T) {
	//tmp:=make([]bool,256)
	//str:="qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_-1234567890"
	//for i:=0;i<len(str);i++{
	//	tmp[str[i]]=true
	//}
	//
	//for i:=0;i<256;i++{
	//	if !tmp[i]{
	//		fmt.Printf("%d ",i)
	//	}
	//}
	//str:=string([]byte{0x00})
	//fmt.Printf(str)

	//fmt.Printf("%b\n", 0xfa)
	//fmt.Printf("%b\n", 15)
	//fmt.Printf("%b\n", 10)
	//fmt.Printf("%b\n", 15<<4)

	typCode := 1 // 0 ~ 1

	dispatchCode := 4 // 0 ~ 7

	version := 8 // 0 ~ 15

	opcode := ((typCode<<3)|dispatchCode)<<4 | version
	fmt.Printf("opcode: %08b %d\n", opcode, opcode)

	version = opcode & 0b00001111
	fmt.Printf("versionCode:%04b %d\n", version, version)

	dispatchCode = opcode >> 4 & 0b0111
	fmt.Printf("dispatchCode:%03b %d\n", dispatchCode, dispatchCode)
	//
	typCode = opcode >> 7
	fmt.Printf("typCode:%01b %d\n", typCode, typCode)

	//l:=500
	//fmt.Printf("%b\n",l)
	//
	//l2:=l
	//for l2>255{
	//	l2=l2>>1
	//}
	//fmt.Printf("%b\n",l2)
}

func TestPPT(t *testing.T) {

	schema := avro.MustParse(`
{
	"namespace": "example.avro",
	"type": "record",
	"name": "User",
	"fields": [
		 {"name": "name", "type": "string"},
		 {"name": "age",  "type": "int"}
	]
}`)

	res, err := avro.Marshal(schema, People{
		Name: "kieran",
		Age:  18,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("0x%x\n", res)

	type RemoteFun struct {
		Name      string `avro:"name"`
		InSchema  string `avro:"inSchema"`
		OutSchema string `avro:"outSchema"`
	}

}

type People struct {
	Name string `avro:"name"`
	Age  int    `avro:"age"`
}

type Service int

func (*Service) AddAge(p People) People {
	p.Age += 1
	return p
}