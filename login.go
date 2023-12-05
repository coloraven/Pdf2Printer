package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
)

func login(client *req.Client) {
	// 第一步: 访问登录页面
	loginPageURL := "http://www.baidu.com/cas/login?service=http%3A%2F%2Fwww.baidu.com%3A8070%2Fbase%2Fadmin%2Flogin%2Fcas"
	resp, err := client.R().Get(loginPageURL)
	if err != nil {
		log.Fatal(err)
	}
	// 提取 execute 字段的值
	// 解析 HTML 文档
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 使用 CSS 选择器查找特定的 input 标签并获取其 value 值
	var executionValue string
	doc.Find("input[name='execute'][type='hidden']").Each(func(i int, s *goquery.Selection) {
		executionValue, exists := s.Attr("value")
		if exists {
			fmt.Println("Found executionValue:", executionValue)
		}
	})

	// 从响应中提取 cookies
	cookies := resp.Cookies()

	// 第二步: 添加额外的 cookies 并提交登录表单
	loginFormData := url.Values{
		"execution": {executionValue},
		"_eventId":  {"submit"},
		"loginType": {"1"},
		"username":  {"administrator"},
		"password":  {"pwd123123"},
	}
	loginActionURL := "http://www.baidu.com/cas/login?service=http%3A%2F%2Fwww.baidu.com%3A8070%2Fbase%2Fadmin%2Flogin%2Fcas"

	// 添加 logintype 和 username 到 cookies
	cookies = append(cookies, &http.Cookie{Name: "logintype", Value: "1"})
	cookies = append(cookies, &http.Cookie{Name: "username", Value: "administrator"})

	resp, err = client.R().
		SetCookies(cookies...).
		SetFormDataFromValues(loginFormData).
		Post(loginActionURL)
	if err != nil {
		log.Fatal(err)
	}

	// 第三步: 处理重定向
	cookies = cookies[1:] // 去掉第一个 cookie——sessionid，因为它是登录成功后的 cookie
	redirectedURL := resp.Header.Get("Location")
	resp, err = client.R().SetCookies(cookies...).Get(redirectedURL)
	if err != nil {
		log.Fatal(err)
	}
	cookies = resp.Cookies() // 其中有新的sessionid
	// 第四步: 访问最终 URL
	// 添加 logintype 和 username 到 cookies
	cookies = append(cookies, &http.Cookie{Name: "logintype", Value: "1"})
	cookies = append(cookies, &http.Cookie{Name: "username", Value: "administrator"})
	finalURL := "http://www.baidu.com/base/"
	resp, err = client.R().SetCookies(cookies...).Get(finalURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("最终响应:", resp)
}
