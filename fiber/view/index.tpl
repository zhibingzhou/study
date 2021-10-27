<html>
<head>
  <title>{{ .title }}</title>
</head>
<body>
  <h1 style="color:red;">接口测试</h1>
  <span>host->{{ .host }}</span>


<p>测试删除记录接口<br/>POST<br/>/admin/del_history.do</p>
<form method="post" action="/admin/del_history.do">
<table border="1" cellspacing="0" cellpadding="0">
<tr>
<th>参数名</th>
<th>传值</th>
<th>备注说明</th>
</tr>
<tr align="center">
<td>password</td>
<td><input type="text" name="password"></td>
<td>加密的密钥</td>
</tr>
</tr>
<tr align="center">
<td>filename</td>
<td><input type="text" name="filename"></td>
<td>要删除的文件名</td>
</tr>
<tr align="center">
<td colspan="3"><input type="submit" value="提交"></td>
</tr>
</table>
</form>

</body>
</html>
