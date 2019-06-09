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