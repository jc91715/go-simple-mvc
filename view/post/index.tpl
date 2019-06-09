{{define "index"}}
{{template "header"}}
 <h3 style="text-align:center">博客列表</h1>
<div style="text-align:center"> 
{{range .}}
 
  <a href="/posts/{{.Id}}" >{{.Title}}</a><br><br>
{{end}}
</div>
{{template "footer"}}
{{end}}