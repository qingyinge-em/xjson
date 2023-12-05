// go test
// go test -v -run Test1
package xjson

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

type testJson struct {
	A int    `json:"a"`
	B string `json:"b"`
	C struct {
		Ca int `json:"ca"`
		Cb int `json:"cb"`
	} `json:"c"`
	D []string `json:"d"`
	E []struct {
		E1 int    `json:"e1"`
		E2 string `json:"e2"`
	} `json:"e"`
}

type cJson struct {
	Ca int `json:"ca"`
	Cb int `json:"cb"`
}

func Test1(t *testing.T) {
	j, _ := NewFromFile("test.json")
	fmt.Println(j.IndentString())
	fmt.Println(j.CompactString())

	fmt.Println("========get========")
	fmt.Printf("a=%d\n", j.GetInt("a"))
	fmt.Printf("b=%s\n", j.GetString("b"))
	fmt.Printf("c.ca=%d\n", j.GetInt("c.ca"))
	fmt.Printf("c.cb=%d\n", j.GetInt("c.cb"))
	dsize := j.GetArraySize("d")
	fmt.Printf("d array size=%d\n", dsize)
	for i := 0; i < dsize; i++ {
		fmt.Printf("d.%d=%s\n", i, j.GetString("d."+cast.ToString(i)))
	}

	fmt.Println("======Unmarshal==========")

	var tj testJson
	j.Unmarshal("", &tj)
	fmt.Printf("tj: %v\n", tj)
	var cj cJson
	j.Unmarshal("c", &cj)
	fmt.Printf("cj: %v\n", cj)

	fmt.Println("======obj array==========")

	esize := j.GetArraySize("e")
	fmt.Printf("e array size=%d\n", esize)
	for i := 0; i < esize; i++ {
		fmt.Printf("e.%d.e1=%d\n", i, j.GetInt("e."+cast.ToString(i)+".e1"))
		fmt.Printf("e.%d.e2=%s\n", i, j.GetString("e."+cast.ToString(i)+".e2"))

		var e struct {
			E1 int    `json:"e1"`
			E2 string `json:"e2"`
		}
		j.Unmarshal("e."+cast.ToString(i), &e)
		fmt.Printf("e Unmarshal: %v\n", e)
	}
}

func Test2(t *testing.T) {
	j, _ := NewFromFile("test.json")
	fmt.Println(j.IndentString())
	fmt.Println(j.CompactString())

	fmt.Println("======set==========")
	j.Set("a", 2)
	j.Set("c.ca", 333)
	j.Set("d.1", "X2")
	j.Set("e.0.e2", "E2222")
	fmt.Println(j.IndentString())

	fmt.Println("======delete==========")
	j.Delete("a")
	j.Delete("c.cb")
	j.Delete("e.0")
	fmt.Println(j.IndentString())
}
