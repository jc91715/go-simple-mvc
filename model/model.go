package model

import (
	// "encoding/json"

	"github.com/astaxie/beego/orm"
)

func init() {

	orm.RegisterModel(new(RainlabBlogPosts))
}
