##### 简单的ChatGPT聊天
- 不必放到海外服务器,使用的是本地代理。HTTP代理端口在completion文件的```getProxyHttpClient```方法中修改
- OpenAI的key需要自己去搞帐号,KEY在completion的```StreamCompletion```方法中替换自己的
- 前端只有一个index.html，是一个原生js编写的前端代码

##### MAC下交叉编译Linux: ```CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go```