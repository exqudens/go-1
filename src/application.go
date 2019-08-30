package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"io/ioutil"
	"strconv"
	"util/string"
)

var _ = stringutil.Reverse
var _ = ioutil.ReadFile
var _ = string.Count

func main() {
	/*bytes1, e := ioutil.ReadFile("/Users/exqudens/work/git/exqudens/go/go-1/src/config.yml")
	if e != nil {
		fmt.Println("Error: ", e)
	}*/
	fmt.Println("---")
	str := "abc"
	fmt.Println(str)
	str = stringutil.Reverse(str)
	fmt.Println(str)
	str = strconv.FormatInt(string.Count(str), 10)
	fmt.Println(str)
	fmt.Println("---")
}
