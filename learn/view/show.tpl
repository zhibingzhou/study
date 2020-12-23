<!doctype html>
<html>
  	<head>
    	<title>{{.title}}</title>
    	<meta http-equiv="content-type" content="text/html; charset=utf-8">
		
	</head>
<style>
table,tr,td{
	border:1px #000000 solid;
}
</style>
<body>
<p><b>爬虫数据</b></p>
<table>
<form  action="/all_data" method="GET">

<a>性别：</a>
<select name="gender">
<option value="man">男</option>
<option value="girl">女</option>
</select>

<a>  身材：</a>
<input list="browsers" name="body">
<datalist id="browsers">
{{range $k, $v := .body}}
<option value="{{$v.Body_title}}">{{$v.Body_title}}</option>
{{end}}
</datalist>


<a>  年龄以下：</a>
<input type="text" name="age" >

<a>  学历：</a>
<input list="browsers1" name="school">
<datalist id="browsers1">
{{range $k, $v := .school}}
<option value="{{$v.School_title}}">{{$v.School_title}}</option>
{{end}}
</datalist>



<a>  工作：</a>
<input list="browsers2" name="job">
<datalist id="browsers2">
{{range $k, $v := .job}}
<option value="{{$v.Job_title}}">{{$v.Job_title}}</option>
{{end}}
</datalist>


<a>  婚姻：</a>
<input list="browsers3" name="marry">
<datalist id="browsers3">
{{range $k, $v := .marry}}
<option value="{{$v.Marry_title}}">{{$v.Marry_title}}</option>
{{end}}
</datalist>

<a>  年收入（万）：</a>
<input list="browsers4" name="maxmoney">
<datalist id="browsers4">
<option value="10">10</option>
<option value="20">20</option>
<option value="30">30</option>
<option value="40">40</option>
<option value="50">50</option>
</datalist>

<br></br>
<a>  共{{.page}}（页） </a>

<a>  输入跳转页数</a>

<input type="text" name="page" >

<input type="submit">



</form>
<tr>
<td>相片</td>
<td>性别</td>
<td>年龄</td>
<td>身材</td>
<td>身高</td>
<td>体重</td>
<td>婚姻</td>
<td>学历</td>
<td>最大收入</td>
<td>最小收入</td>
<td>地址</td>
<td>备注</td>
<td>详情链接</td>
</tr>
	{{range $k, $v := .list}}
				<tr>
                <td><img src="{{$v.ImgUrl}}" width="80" height="80" /></td>
                <td>{{$v.Gender}}</td>
				<td>{{$v.Age}}</td>
				<td>{{$v.Body}}</td>
				<td>{{$v.Height}}cm</td>
				<td>{{$v.Weight}}kg</td>
				<td>{{$v.Marry}}</td>
				<td>{{$v.School}}</td>
				<td>{{$v.MoneyMax}}万</td>
				<td>{{$v.MoneyMin}}万</td>
				<td>{{$v.Address}}</td>
				<td>{{$v.Note}}</td>
				<td><a href="{{$v.GoUrl}}">点我跳转</a></td>
                </tr>
    {{end}}
 

</table>
<br></br>
</body>
</html>