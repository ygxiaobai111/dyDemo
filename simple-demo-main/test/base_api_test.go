package test

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

// 这是一个示例的Go代码片段

func TestFeed(t *testing.T) {
	// 创建一个期望对象
	e := newExpect(t)

	// 发送GET请求，获取/feed/接口的响应
	feedResp := e.GET("/douyin/feed/").Expect().Status(http.StatusOK).JSON().Object()
	feedResp.Value("status_code").Number().Equal(0)
	feedResp.Value("video_list").Array().Length().Gt(0)

	// 遍历视频列表
	for _, element := range feedResp.Value("video_list").Array().Iter() {
		// 获取视频对象
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}
}

func TestUserAction(t *testing.T) {
	// 创建一个期望对象
	e := newExpect(t)

	// 生成一个随机的注册值
	rand.Seed(time.Now().UnixNano())
	registerValue := fmt.Sprintf("douyin%d", rand.Intn(65536))

	// 发送POST请求，进行用户注册
	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", registerValue).WithQuery("password", registerValue).
		WithFormField("username", registerValue).WithFormField("password", registerValue).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	registerResp.Value("status_code").Number().Equal(0)
	registerResp.Value("user_id").Number().Gt(0)
	registerResp.Value("token").String().Length().Gt(0)

	// 发送POST请求，进行用户登录
	loginResp := e.POST("/douyin/user/login/").
		WithQuery("username", registerValue).WithQuery("password", registerValue).
		WithFormField("username", registerValue).WithFormField("password", registerValue).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	loginResp.Value("status_code").Number().Equal(0)
	loginResp.Value("user_id").Number().Gt(0)
	loginResp.Value("token").String().Length().Gt(0)

	// 获取登录响应中的token
	token := loginResp.Value("token").String().Raw()

	
