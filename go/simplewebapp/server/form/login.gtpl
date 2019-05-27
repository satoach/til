<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
<p>
<input type="hidden" name="token", value="{{.}}">
<input type="submit" value="ログイン">
ユーザー名:<input type="text" name="username">
パスワード:<input type="text" name="password">
</p>
<p>
<select name="fruit">
<option value="apple">apple</option>
<option value="pear">pear</option>
<option value="banana">banana</option>
</select>
</p>
<p>
<input type="radio" name="gender" value="1">男
<input type="radio" name="gender" value="2">女
</p>
</form>
</body>
</html>
