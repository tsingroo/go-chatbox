package gpt3

import (
	"net/http/httputil"
	"net/url"
)

func GetStreamCompletionReverseProxy() *httputil.ReverseProxy {
	// 定义转发的目标地址
	targetURL, _ := url.Parse("http://example.com/forward")

	// 创建一个 ReverseProxy 实例
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return proxy
}
