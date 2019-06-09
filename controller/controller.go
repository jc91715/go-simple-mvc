package controller

import (
	"contex"
	"fmt"
)

type Controller struct {
	Ct        *contex.Context
	ChildName string
}

func (c *Controller) Init(ct *contex.Context, cn string) {

	c.ChildName = cn
	c.Ct = ct

	fmt.Println("\n---------")
	fmt.Println("\nhello Init")

}

func (c *Controller) Prepare() {
	fmt.Println("\nhello Prepare")
}

func (c *Controller) Finish() {
	fmt.Println("\nhello Finish")
	fmt.Println("\n---------")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
