package curl

//curl 封装转发

import (
	"io/ioutil"
	"net/http"
	"strings"

	json "github.com/bitly/go-simplejson"
)

type ReturnJson struct {
	Code    int
	Message string
	Data    interface{}
}

func CurlGet(url string, header []string) (ret []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ret, err
	}
	//请求头的读取，对于报文的请求头，可能的格式是 [key:value,key:value,key:value],所以我们读取的方式是遍历header这个数组，对于每个string用:拆分，然后req.Header.Add（key，value）加进去
	for _, v := range header {
		t := strings.Split(v, ":")
		length := len(t)
		if length == 2 {
			req.Header.Add(t[0], t[1])
		} else if length == 1 {
			req.Header.Add(t[0], "")
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	ret, err = ioutil.ReadAll(resp.Body)

	return ret, err
}
//这里用的是simplejson，默认返回是一个json的格式数据，有code，message，data三个字段
func CurlGetReturnJson(url string, header []string) (r ReturnJson) {
	r = ReturnJson{}
	ret, err := CurlGet(url, header)
	if err == nil {
		data, err := json.NewJson(ret)
		if err == nil {
			r.Code, err = data.Get("code").Int()
			r.Message, err = data.Get("message").String()
			r.Data = data.Get("data").Interface()
		}
	}

	return r
}

func CurlPost(url string, header []string, data string) (ret []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return ret, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value; charset=utf-8")
	//请求头的读取，对于报文的请求头，可能的格式是 [key:value,key:value,key:value],所以我们读取的方式是遍历header这个数组，对于每个string用:拆分，然后req.Header.Add（key，value）加进去
	for _, v := range header {
		t := strings.Split(v, ":")
		length := len(t)
		if length == 2 {
			req.Header.Add(t[0], t[1])
		} else if length == 1 {
			req.Header.Add(t[0], "")
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	ret, err = ioutil.ReadAll(resp.Body)

	return ret, err
}
//这里用的是simplejson，默认返回是一个json的格式数据，有code，message，data三个字段
func CurlPostReturnJson(url string, header []string, data string) (r ReturnJson) {
	r = ReturnJson{}
	ret, err := CurlPost(url, header, data)
	if err == nil {
		data, err := json.NewJson(ret)
		if err == nil {
			r.Code, err = data.Get("code").Int()
			r.Message, err = data.Get("message").String()
			r.Data = data.Get("data").Interface()
		}

	}

	return r
}

func CurlPut(url string, header []string, data string) (ret []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, strings.NewReader(data))
	if err != nil {
		return ret, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value; charset=utf-8")
	for _, v := range header {
		t := strings.Split(v, ":")
		length := len(t)
		if length == 2 {
			req.Header.Add(t[0], t[1])
		} else if length == 1 {
			req.Header.Add(t[0], "")
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	ret, err = ioutil.ReadAll(resp.Body)

	return ret, err
}

func CurlPutReturnJson(url string, header []string, data string) (r ReturnJson) {
	r = ReturnJson{}
	ret, err := CurlPut(url, header, data)
	if err == nil {
		data, err := json.NewJson(ret)
		if err == nil {
			r.Code, err = data.Get("code").Int()
			r.Message, err = data.Get("message").String()
			r.Data = data.Get("data").Interface()
		}
	}

	return r
}

func CurlDelete(url string, header []string) (ret []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return ret, err
	}
	for _, v := range header {
		t := strings.Split(v, ":")
		length := len(t)
		if length == 2 {
			req.Header.Add(t[0], t[1])
		} else if length == 1 {
			req.Header.Add(t[0], "")
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	ret, err = ioutil.ReadAll(resp.Body)

	return ret, err
}

func CurlDeleteReturnJson(url string, header []string) (r ReturnJson) {
	r = ReturnJson{}
	ret, err := CurlDelete(url, header)
	if err == nil {
		data, err := json.NewJson(ret)
		if err == nil {
			r.Code, err = data.Get("code").Int()
			r.Message, err = data.Get("message").String()
			r.Data = data.Get("data").Interface()
		}
	}

	return r
}
