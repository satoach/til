<html>
<head>
<title>アップロード</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
<input type="hidden" name="token" value="{{.}}" />
<input type="submit" value="upload" />
<input type="file" name="uploadfile" />
</form>
</body>
</html>
