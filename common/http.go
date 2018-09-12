package common

import (
	"crypto/tls"
	"crypto/x509"
	"image"
	"io/ioutil"
	"net"
	"net/http"
	c_url "net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type myRegexp struct {
	*regexp.Regexp
}

/**
* 写入cookie
 */

func InsertCookie(w http.ResponseWriter, doname string, key string, val string, exptime int) {
	name := key                                                     //cookie名称
	value := val                                                    //cookie的值
	path := "/"                                                     //作用目录
	expires := time.Now().Add(time.Second * time.Duration(exptime)) //失效时间，单位秒
	maxAge := exptime                                               //新的失效方式，单位秒
	secure := false                                                 //安全cookie
	httpOnly := true                                                //限制http访问

	computerCookie := &http.Cookie{Name: name, Value: value, Path: path, Domain: doname, Expires: expires, MaxAge: maxAge, Secure: secure, HttpOnly: httpOnly}
	http.SetCookie(w, computerCookie)
}

/**
* 读取cookie
 */
func ReadCookie(r *http.Request, key string) string {
	res := ""
	val, err := r.Cookie(key)
	if err == nil {
		res = val.Value
	}
	return res
}

/**
* 获得计算机标识=md5(注册IP+浏览器+操作系统os+当前时间戳)
* return string 获得计算机唯一标识
 */
func GetComputer(IP string, bb string, oo string) string {
	computer := GetMd5(IP + bb + oo + strconv.FormatInt(time.Now().Unix(), 10))
	return computer
}

/**
* 根据传入的用户访问头，判断浏览器类型
* @use_agent string  用户传递过来的header中的user-agent
* return string,string  返回浏览器类型,操作系统类型
 */

func GetBrowserOS(use_agent string) (string, string) {
	OS := "other"
	if strings.Index(use_agent, "Windows") > -1 || strings.Index(use_agent, "win") > -1 {
		var re = myRegexp{regexp.MustCompile(`[Windows|win] (.*?)[;|)| ].*?`)}
		res := re.FindStringSubmatch(use_agent)
		OS := "Windows"
		if len(res) > 1 {
			winVersion := ""
			if strings.Index(res[1], "95") > -1 {
				winVersion = "95"
			}
			if strings.Index(res[1], "98") > -1 {
				winVersion = "98"
			}
			if strings.Index(res[1], "4.9") > -1 {
				winVersion = "ME"
			}
			if strings.Index(res[1], "NT 4.0") > -1 {
				winVersion = "NT 4.0"
			}
			if strings.Index(res[1], "NT 5.0") > -1 {
				winVersion = "2000"
			}
			if strings.Index(res[1], "NT 5.1") > -1 {
				winVersion = "XP"
			}
			if strings.Index(res[1], "NT 5.2") > -1 {
				winVersion = "Server 2003"
			}
			if strings.Index(res[1], "NT 6.0") > -1 {
				winVersion = "Vista"
			}
			if strings.Index(res[1], "NT 6.1") > -1 {
				winVersion = "7"
			}
			if strings.Index(res[1], "NT 6.2") > -1 {
				winVersion = "8"
			}
			if strings.Index(res[1], "NT 6.3") > -1 {
				winVersion = "Server 2012"
			}
			if strings.Index(res[1], "NT 10.0") > -1 {
				winVersion = "10"
			}
			OS = OS + " " + winVersion
		}
	}
	if strings.Index(use_agent, "Mac") > -1 {
		OS = "Mac OS"
	}
	if strings.Index(use_agent, "Linux") > -1 {
		OS = "Linux"
	}
	if strings.Index(use_agent, "Unix") > -1 {
		OS = "Unix"
	}
	if strings.Index(use_agent, "Sun") > -1 {
		OS = "Sun OS"
	}
	if strings.Index(use_agent, "ibm") > -1 {
		OS = "Ibm"
	}
	if strings.Index(use_agent, "PowerPC") > -1 {
		OS = "PowerPC"
	}
	if strings.Index(use_agent, "AIX") > -1 {
		OS = "AIX"
	}
	if strings.Index(use_agent, "HPUX") > -1 {
		OS = "HPUX"
	}
	if strings.Index(use_agent, "BSD") > -1 {
		OS = "BSD"
	}

	if strings.Index(use_agent, "MSIE") > -1 {
		var re = myRegexp{regexp.MustCompile(`MSIE (.*?)[;|)| ].*?`)}
		res := re.FindStringSubmatch(use_agent)
		browser := "Internet Explorer"
		if len(res) > 1 {
			browser = browser + " " + res[1]
		}
		return browser, OS
	}
	if strings.Index(use_agent, "Firefox") > -1 {
		return "Firefox", OS
	}
	if strings.Index(use_agent, "Chrome") > -1 {
		return "Chrome", OS
	}
	if strings.Index(use_agent, "Opera") > -1 {
		return "Opera", OS
	}
	if strings.Index(use_agent, "Safari") > -1 {
		return "Safari", OS
	}
	return "other", OS
}

/**
* 模拟get请求
* @url		string   需要抓取的url
* @charset	string	字符编码
 */
func HttpGetBody(url string, charset string) (string, int) {
	src := ""
	httpStart := true

	statusCode := 101

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		LogsWithcontent(url + err.Error())
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			httpStart = false
			//两次抓取都失败了，需要返回一个空
			return "", statusCode
		}
	}

	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {
			sHost = u.Host
		}

		if len(charset) > 0 {
			req.Header.Set("Content-Type", "charset="+charset)
		} else {
			req.Header.Set("Content-Type", "charset=UTF-8")
		}
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Host", sHost)
		req.Header.Set("Referer", sHost)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")
		tr := &http.Transport{DisableKeepAlives: true,
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*15) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
				return c, nil
			}}
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)

		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {

			statusCode = resp.StatusCode
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				LogsWithcontent(url + err.Error())
			} else {
				src = string(body)
			}
		}
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
	}
	return src, statusCode
}

/**
* 模拟get请求
* @url		string   需要抓取的url
* @charset	string	字符编码
* @return 返回图片
 */
func HttpGetBodyByImg(url string) (image.Image, int) {
	httpStart := true

	statusCode := 101

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		LogsWithcontent(url + err.Error())
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			httpStart = false
			//两次抓取都失败了，需要返回一个空
			return nil, statusCode
		}
	}

	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {
			sHost = u.Host
		}

		req.Header.Set("Content-Type", "charset=UTF-8")

		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Set("Connection", "	keep-alive")
		req.Header.Set("Host", sHost)
		req.Header.Set("Referer", sHost)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")
		tr := &http.Transport{DisableKeepAlives: true,
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*15) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
				return c, nil
			}}
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)

		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {

			statusCode = resp.StatusCode
			img, _, er := image.Decode(resp.Body)
			if er != nil {
				LogsWithcontent("Image Decode->" + url + err.Error())
			} else {
				return img, statusCode
			}
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
	}
	return nil, statusCode
}

/**
* 模拟POST请求
* @url		string	需要发送请求的url
* @param 	string	需要传递的参数,格式：key=val&key=val
* @gbk		bool 	如果需要抓取的网页是gbk的编码，则需要进行转码
 */
func HttpPostBody(url string, param string, gbk bool) (string, int) {
	src := ""
	statusCode := 101
	req, err := http.NewRequest("POST", url, strings.NewReader(param))

	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}

	//模拟一个host
	sHost := ""
	u, err := c_url.Parse(url)
	if err != nil {
		LogsWithcontent(url + err.Error())
	} else {
		sHost = u.Host
	}
	req.Header.Set("Host", sHost)
	req.Header.Set("Referer", sHost)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")

	if gbk == true {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=GBK")
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}
	tr := &http.Transport{DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*15) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		LogsWithcontent(url + err.Error())
	} else {

		statusCode = resp.StatusCode
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {
			src = string(contents)
		}
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	return src, statusCode
}

/**
* 模拟POST请求
* @url		string	需要发送请求的url
* @param 	string	需要传递的参数,格式：key=val&key=val
* @gbk		bool 	如果需要抓取的网页是gbk的编码，则需要进行转码
 */
func HttpPostBodyForOutApi(url string, param string) (string, int) {
	src := ""
	statusCode := 101

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		LogsWithcontent(url + err.Error())
		req, err = http.NewRequest("POST", url, nil)
		if err != nil {
			return "", statusCode
		}
	}

	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}

	//模拟一个host
	sHost := ""
	u, err := c_url.Parse(url)
	if err != nil {
		LogsWithcontent(url + err.Error())
	} else {
		sHost = u.Host
	}

	req.Header.Set("Host", sHost)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	buf := make([]byte, len(param))
	req.Body = ioutil.NopCloser(strings.NewReader(param))

	req.Body.Read(buf)

	tr := &http.Transport{DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*15) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		LogsWithcontent(url + err.Error())
	} else {

		statusCode = resp.StatusCode
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {
			src = string(contents)
		}
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	return src, statusCode
}

/**
* 模拟https的POST请求
* @url		string	需要发送请求的url
* @param 	string	需要传递的参数,格式：key=val&key=val
* @gbk		bool 	如果需要抓取的网页是gbk的编码，则需要进行转码
 */
func HttpsPostBody(url, param string, certPath []byte, keyPath []byte, headSet map[string]string, gbk bool) (string, int) {

	src := ""
	statusCode := 101
	//加载安全证书
	cert, err := tls.X509KeyPair(certPath, keyPath)

	if err != nil {
		LogsWithcontent("证书解密错误->" + err.Error())
		return "", statusCode
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(certPath)
	_tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	tr := &http.Transport{TLSClientConfig: _tlsConfig,
		DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*15) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}
	httpStart := true

	//模拟post请求
	req, err := http.NewRequest("POST", url, strings.NewReader(param))
	if err != nil {
		httpStart = false
		LogsWithcontent(url + err.Error())
		//两次连接都失败了，需要返回一个空
		return "", statusCode
	}

	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {
			sHost = u.Host
		}
		req.Header.Set("Host", sHost)
		req.Header.Set("Referer", sHost)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")
		for k, v := range headSet {
			req.Header.Add(k, v)
		}

		if gbk == true {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=GBK")
		} else {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		}

		resp, err := client.Do(req)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {

			statusCode = resp.StatusCode
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				LogsWithcontent(url + err.Error())
			} else {
				src = string(contents)
			}
		}
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
	}
	return src, statusCode
}

/**
* 模拟https的POST请求
* @url		string	需要发送请求的url
* @param 	string	需要传递的参数,格式：key=val&key=val
* @gbk		bool 	如果需要抓取的网页是gbk的编码，则需要进行转码
 */
func HttpsPostXmlNoCert(url, param string, gbk bool) (string, int) {

	src := ""
	statusCode := 101

	_tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	tr := &http.Transport{TLSClientConfig: _tlsConfig,
		DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*15) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}
	httpStart := true

	//模拟post请求
	req, err := http.NewRequest("POST", url, strings.NewReader(param))
	if err != nil {

		LogsWithcontent(url + err.Error())
		req, err = http.NewRequest("POST", url, strings.NewReader(param))
		if err != nil {
			httpStart = false
			//两次连接都失败了，需要返回一个空
			return "", statusCode
		}
	}

	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {
			sHost = u.Host
		}
		req.Header.Set("Host", sHost)
		req.Header.Set("Referer", sHost)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")

		if gbk == true {
			req.Header.Set("Content-Type", "text/xml; charset=GBK")
		} else {
			req.Header.Set("Content-Type", "text/xml; charset=UTF-8")
		}

		resp, err := client.Do(req)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {

			statusCode = resp.StatusCode
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				LogsWithcontent(url + err.Error())
			} else {
				src = string(contents)
			}
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
	}
	return src, statusCode
}

/**
* xml格式请求
* @param 	string 	url		请求的地址
* @param	string	params	带入的参数
* @param	string	key		head头健
* @param	string	val		head的值
 */
func HttpPostXml(url, params, key, val string, gbk bool) (string, int) {
	src := ""
	httpStart := true
	statusCode := 101
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {

		LogsWithcontent(url + err.Error())
		req, err = http.NewRequest("POST", url, strings.NewReader(params))
		if err != nil {
			httpStart = false
			//两次连接都失败了，需要返回一个空
			return "", statusCode
		}
	}

	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)

		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {
			sHost = u.Host
		}

		req.Header.Set("Host", sHost)
		req.Header.Set("Referer", sHost)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")

		if gbk == true {
			req.Header.Set("Content-Type", "text/xml; charset=GBK")
		} else {
			req.Header.Set("Content-Type", "text/xml; charset=UTF-8")
		}

		if len(key) > 1 {
			mid := key + ":" + val
			req.Header.Set("Headers", mid)
		}
		tr := &http.Transport{DisableKeepAlives: true,
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*15) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
				return c, nil
			}}
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {

			statusCode = resp.StatusCode
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				LogsWithcontent(url + err.Error())
			} else {
				src = string(contents)
			}
		}
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
	}
	return src, statusCode
}

/**
* @param 	string 	url		请求的地址
* @param	string	params	带入的参数
* @param	map	header		head头
* @param	string	method		http方法
 */
func HttpContentByHeader(url, params string, header map[string]string, method string) (string, int) {
	src := ""
	httpStart := true
	statusCode := 101
	req, err := http.NewRequest(method, url, strings.NewReader(params))
	if err != nil {

		LogsWithcontent(url + err.Error())
		req, err = http.NewRequest(method, url, strings.NewReader(params))
		if err != nil {
			httpStart = false
			//两次连接都失败了，需要返回一个空
			return "", statusCode
		}
	}

	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)

		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {
			sHost = u.Host
		}

		req.Header.Set("Host", sHost)
		req.Header.Set("Referer", sHost)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")

		req.Header.Set("Content-Type", "application/xml; charset=UTF-8")

		if len(header) > 0 {
			//mid := ""
			for key, val := range header {
				//mid = key + ":" + val
				req.Header.Set(key, val)
			}
		}
		tr := &http.Transport{DisableKeepAlives: true,
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*15) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
				return c, nil
			}}
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {

			statusCode = resp.StatusCode
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				LogsWithcontent(url + err.Error())
			} else {
				src = string(contents)
			}
		}
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
	}
	return src, statusCode
}

/**
* 推送post请求
 */
func PushHttpPost(url, params string) (string, int) {
	//post请求
	src := ""
	httpStart := true
	statusCode := 101
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {

		LogsWithcontent(url + err.Error())
		req, err = http.NewRequest("POST", url, strings.NewReader(params))
		if err != nil {
			httpStart = false
			//两次连接都失败了，需要返回一个空
			return "", statusCode
		}
	}

	if req != nil && req.Body != nil {
		defer req.Body.Close()
	}

	//只有连接成功后，才会写入头的读取字节流
	if httpStart == true {
		//模拟一个host
		sHost := ""
		u, err := c_url.Parse(url)
		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {
			sHost = u.Host
		}

		req.Header.Set("Host", sHost)
		req.Header.Set("Referer", sHost)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

		tr := &http.Transport{DisableKeepAlives: true,
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*15) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(30 * time.Second)) //设置发送接收数据超时
				return c, nil
			}}
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)

		if err != nil {
			LogsWithcontent(url + err.Error())
		} else {

			statusCode = resp.StatusCode
			contents, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				LogsWithcontent(url + err.Error())
			} else {
				src = string(contents)
			}
		}
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
	}
	return src, statusCode
}

/**
* 获取访问网站的客户端类型
* @userAgent string 浏览器中的user-Agent
* @return string 客户端类型
* pc:电脑端  android:安卓手机  iphone:iphone，ipad等移动端
 */
func GetWebVisitType(userAgent string) string {
	res := "pc"
	agent := strings.ToLower(userAgent)

	if len(agent) < 1 {
		return res
	}

	if strings.Contains(agent, "android") {
		res = "android"
	}

	if strings.Contains(agent, "iphone") || strings.Contains(agent, "ipod") || strings.Contains(agent, "ipad") {
		res = "ios"
	}

	//	mobile_map := map[string]string{
	//		"up.browser":    "up.browser",
	//		"up.link":       "up.link",
	//		"mmp":           "mmp",
	//		"symbian":       "symbian",
	//		"smartphone":    "smartphone",
	//		"midp":          "midp",
	//		"wap":           "wap",
	//		"phone":         "phone",
	//		"iphone":        "iphone",
	//		"ipad":          "iphone",
	//		"ipod":          "iphone",
	//		"android":       "android",
	//		"xoom":          "xoom",
	//		"operamini":     "operamini",
	//		"windows phone": "wp",
	//	}

	//	if len(mobile_map[agent]) > 1 {
	//		res = mobile_map[agent]
	//	}

	return res
}
