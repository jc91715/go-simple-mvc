package controller

import (
	"fmt"
	"model"

	"html/template"
	"net/http"

	"os"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type PostController struct {
	Controller
}

func (c *PostController) Index() {
	o := orm.NewOrm()
	fmt.Printf("Hello PostController Index")
	var posts []*model.RainlabBlogPosts

	num, err := o.QueryTable("rainlab_blog_posts").All(&posts)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	s1, _ := template.ParseFiles("view/layout/header.tpl", "view/post/index.tpl", "view/layout/footer.tpl")
	fmt.Println(posts)
	s1.ExecuteTemplate(os.Stdout, "index", posts)
	s1.ExecuteTemplate(c.Ct.ResponseWriter, "index", posts)
}
func (c *PostController) Show() {
	fmt.Println("\nHello PostController Show:%s", c.Ct.Params["post_id"])
	id, ok := c.Ct.Params["post_id"] /*如果确定是真实的,则存在,否则不存在 */
	if !ok {
		fmt.Println("不存在")
		http.NotFound(c.Ct.ResponseWriter, c.Ct.Request)
	} else {
		o := orm.NewOrm()
		intId, err := strconv.Atoi(id)
		CheckErr(err)
		post := model.RainlabBlogPosts{Id: intId}
		err = o.Read(&post)
		if err == orm.ErrNoRows {
			fmt.Println("查询不到")
		} else if err == orm.ErrMissPK {
			fmt.Println("找不到主键")
		} else {
			fmt.Println(post.Id, post.ContentHtml)
			s1, _ := template.ParseFiles("view/layout/header.tpl", "view/post/show.tpl", "view/layout/footer.tpl")

			// s1.ExecuteTemplate(os.Stdout, "show", nil)
			s1.ExecuteTemplate(c.Ct.ResponseWriter, "show", template.HTML(post.ContentHtml))
			// s1.Execute(c.Ct.ResponseWriter, post)
		}

		// s1.Execute(c.Ct.ResponseWriter, nil)
	}

}
