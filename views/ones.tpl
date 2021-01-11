<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>每日一图</title>
    <link rel="stylesheet" href="/static/css/onecss.css">
    <link rel="stylesheet" href="/static/css/pure-min.css">
</head>
<body>
    <div class="one-div-nav">
        <img src="http://image.wufazhuce.com/www-fp-logo.png">
    </div>
    {{ range .Ones }}
    <div class="one-div-connect">
        <div class="one-div-titulo">
            <P>{{.Titulo}}</P>
        </div>
        <div class="one-div-image">
            <img class="one-image" src={{ .Image}}>
        </div>
        <div class="one-div-leyenda">
            <P>{{.Leyenda}}</P>
        </div>
        <div class="one-div-title">
            <div class="one-div-cita">
                <P>{{.Cita}}</P>
            </div>
            <div class="one-div-pubdate">
                <P>{{.Pubdate}}</P>
            </div>
        </div>
    </div>
    {{ end }}
    <div class="one-buttons">
        <a class="one-button-previous pure-button-primary pure-button" href="/ones/pre/{{.FirstOne}}">上一页</a>
        <a class="one-button-next pure-button-primary pure-button" href="/ones/next/{{.LastOne}}">下一页</a>
    </div>

</body>
</html>
