/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-9
* Time: 下午6:10
* */
package easyutils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ProxySt struct {
	UserName    string
	Password    string
	ProxyServer string
	proxyUrl    *url.URL
}

var proxyD *ProxySt

// 初始化代理
// url,username,password
func InitProxy(arg ...string) (*ProxySt, error) {
	if proxyD != nil {
		return proxyD,nil
	}else{
		// 初始化
		proxyData := &ProxySt{}
		if len(arg) == 1 {
			parse, e := url.Parse("http://" + arg[0])
			if e != nil {
				return nil, e
			} else {
				proxyData.proxyUrl = parse
			}

			proxyD = proxyData

			return proxyData, nil
		} else if len(arg) == 3 {
			parse, e := url.Parse("http://" + arg[1] + ":" + arg[2] + "@" + arg[0])
			if e != nil {
				return nil, e
			} else {
				proxyData.proxyUrl = parse
			}

			proxyD = proxyData


			return proxyData, nil
		} else {

			proxyD = proxyData

			return proxyData, nil
		}
	}
}

var userAgentList = []string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

// 获取随机UserAgent
func (p *ProxySt)ReptileGetUserAgent() string {
	rand.Seed(time.Now().UnixNano())

	intn := rand.Intn(len(userAgentList))
	return userAgentList[intn]
}

var spiderAgent = []string{
	"Mozilla/5.0 (compatible; Baiduspider/2.0;+http://www.baidu.com/search/spider.html）",
	"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
	"Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
}

// 获取蜘蛛agent
func (p *ProxySt)ReptileGetSpiderAgent() string {
	rand.Seed(time.Now().UnixNano())

	intn := rand.Intn(len(spiderAgent))
	return spiderAgent[intn]
}

// 请求 假装成 蜘蛛
func (p *ProxySt) ReptileSpiderRequestFrom(targerUrl string, body io.Reader, cookies []*http.Cookie) (*http.Response, error) {
	targerUrl = strings.TrimSpace(targerUrl)
	var httpClient *http.Client
	if p != nil {
		httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(p.proxyUrl)}}
	} else {
		httpClient = &http.Client{}
	}
	if body != nil {
		request, e := http.NewRequest("POST", targerUrl, body)
		if e != nil {
			return nil, e
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Set("User-Agent", p.ReptileGetSpiderAgent())
		if cookies != nil {
			for _, v := range cookies {
				request.AddCookie(v)
			}
		}
		response, e := httpClient.Do(request)
		if e != nil {
			return nil, e
		}
		return response, e
	} else {
		request, e := http.NewRequest("GET", targerUrl, nil)
		if e != nil {
			return nil, e
		}
		request.Header.Set("User-Agent", p.ReptileGetSpiderAgent())
		if cookies != nil {
			for _, v := range cookies {
				request.AddCookie(v)
			}
		}
		response, e := httpClient.Do(request)
		if e != nil {
			return nil, e
		}
		return response, e
	}
}

// 请求 假装成 用户
func (p *ProxySt) ReptileUserRequestFrom(targerUrl string, body io.Reader, cookies []*http.Cookie) (*http.Response, error) {
	targerUrl = strings.TrimSpace(targerUrl)
	var httpClient *http.Client
	if p != nil {
		httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(p.proxyUrl)}}
	} else {
		httpClient = &http.Client{}
	}
	if body != nil {
		request, e := http.NewRequest("POST", targerUrl, body)
		if e != nil {
			return nil, e
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Set("User-Agent", p.ReptileGetUserAgent())
		if cookies != nil {
			for _, v := range cookies {
				request.AddCookie(v)
			}
		}
		response, e := httpClient.Do(request)
		if e != nil {
			return nil, e
		}
		return response, e
	} else {
		request, e := http.NewRequest("GET", targerUrl, nil)
		if e != nil {
			return nil, e
		}
		request.Header.Set("User-Agent", p.ReptileGetUserAgent())
		if cookies != nil {
			for _, v := range cookies {
				request.AddCookie(v)
			}
		}
		response, e := httpClient.Do(request)
		if e != nil {
			return nil, e
		}
		return response, e
	}
}

// 文件下载
func (p *ProxySt) ReptileDownloadSimple(targerUrl string, cookies []*http.Cookie) ([]byte, error) {
	targerUrl = strings.TrimSpace(targerUrl)

	var httpClient *http.Client
	if p != nil {
		httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(p.proxyUrl)}}
	} else {
		httpClient = &http.Client{}
	}

	request, e := http.NewRequest("GET", targerUrl, nil)
	if e != nil {
		return nil, e
	}
	request.Header.Set("User-Agent", p.ReptileGetUserAgent())
	if cookies != nil {
		for _, v := range cookies {
			request.AddCookie(v)
		}
	}
	response, e := httpClient.Do(request)
	if e != nil {
		return nil, e
	}

	bytes, e := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if e != nil {
		return nil, e
	}

	return bytes, e
}

// 文件下载并保存
// 目标地址,cookies,文件名称,新路径
func (p *ProxySt) ReptileDownloadAndSaveSimple(targerUrl string, cookies []*http.Cookie, name, path string) (string, error) {
	bytes, e := p.ReptileDownloadSimple(targerUrl, cookies)
	if e != nil {
		return "", e
	}

	s, e := FileSaveRenameSimple(name, bytes, path)
	if e != nil {
		return "", e
	}

	return s, e
}

// 验证代理是否可用
func (p *ProxySt) CheckProxy() error {
	fmt.Println(p.proxyUrl)
	if p.proxyUrl != nil {
		req, _ := http.NewRequest("GET", "https://www.baidu.com/", nil) //这里自己搭个web服务验证代理是否可用
		cli2 := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(p.proxyUrl),
			},
		}
		resp, err := cli2.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return errors.New("error")
		}
		return nil
	}
	return errors.New("is nol")

}
