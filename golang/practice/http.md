+ **实现**
```golang
  resp,err := http.Get(url)
  resp,err := http.Post(url, "image/jpeg", &buf)
  resp,err := http.PoatForm(url, url.Values{"key": {"Value"}, "id": {"123"}})
```

+ **关闭客户端响应体**
```goalng
  defer resp.Body.Close()
  body,err := ioutil.ReadAll(resp.Body)
```
