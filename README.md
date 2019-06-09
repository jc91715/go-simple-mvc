# go-simple-mvc

[demo](http://go-simple-mvc.jc91715.top:9090/)

从这个文档[build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)学习到很多。

## 安装依赖

```
go get github.com/astaxie/beego
go get github.com/go-sql-driver/mysql

```

## 运行

```
go run main.go
```

访问 localhost:9090

## 使用说明

## 1 添加路由
在`routes/web.go`中

```
a.AddRoute("/posts/:post_id([0-9]+)", map[string]string{
		"GET": "Show",//对应PostController的Show方法
	}, &controller.PostController{})
```
## 2 创建控制器

在`controller`下创建`PostController.go`

```
package controller

//导入要用到的包

type PostController struct {
	Controller
}
func (c *PostController) Show() {//添加方法

}
```
## 3 模型
使用的[beego-orm](https://beego.me/docs/mvc/model/overview.md)

```
package model

type RainlabBlogPosts struct {
	Id          int
	Title       string
	ContentHtml string
}

```
使用
```
    o := orm.NewOrm()
		
    post := model.RainlabBlogPosts{Id: id}
```
## 4 视图

在 `view`文件夹下

抽离出`header.tpl`和`footer.tpl`
```
{{define "header"}}
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>go-simple-mvc</title>
    <style>
      body{
        margin:30px;
     
      }
    </style>
</head>
<body>
{{end}}
```
```
{{define "footer"}}
</body>
</html>
{{end}}
```
例子 `post/show.tpl`

```
{{define "show"}}
{{template "header"}}
<div style="text-align:center"> 
  <a href="/" >回首页</a>
</div>
<div style="width:80%;margin-left:10%"> 
  {{.}}
</div>

{{template "footer"}}
{{end}}

```

控制器调用视图


```
  s1, _ := template.ParseFiles("view/layout/header.tpl", "view/post/show.tpl", "view/layout/footer.tpl")//装载视图

  s1.ExecuteTemplate(c.Ct.ResponseWriter, "show", "string")

```
