package main

import (
	"exqudens/model/user"
	"exqudens/util"
	"fmt"
	"github.com/golang/example/stringutil"
	"io/ioutil"
	"strconv"
)

var _ = stringutil.Reverse
var _ = ioutil.ReadFile
var _ = util.Count

func main() {
	/*bytes1, e := ioutil.ReadFile("/Users/exqudens/work/git/exqudens/go/go-1/src/config.yml")
	if e != nil {
		fmt.Println("Error: ", e)
	}*/
	fmt.Println("---")
	order1 := user.Order{1, "name_1"}
	order2 := order1
	update(&order2)
	order3 := order2
	update(&order3)
	order4 := order3
	fmt.Println(order1)
	fmt.Println(order2)
	fmt.Println(order3)
	fmt.Println(order4)
	orders := [] user.Order {}
	orders = append(orders, order1)
	orders = append(orders, order2)
	orders = append(orders, order3)
	orders = append(orders, order4)
	for i, o := range orders {
		fmt.Println(i, " --- ", o)
	}
	fmt.Println("---")
	str := order1.Name
	fmt.Println(str)
	str = stringutil.Reverse(str)
	fmt.Println(str)
	str = strconv.FormatInt(util.Count(str), 10)
	fmt.Println(str)
	fmt.Println("---")
}

func update(order *user.Order) {
	order.Id = order.Id + 1
	order.Name = "name_" + strconv.FormatInt(order.Id, 10)
}
