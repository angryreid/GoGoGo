# Go Learning

## Notice

In this project, the `crawler` folder aims to fetch data form `www.zhenai.com` which has 
anti-reptile strategies to limit the operation of coders.
```text
2021/07/17 13:13:43 Engine: fetching url http://album.zhenai.com/u/1667348260
2021/07/17 13:13:44 Fetcher: error in fetching url http://album.zhenai.com/u/1667348260: http status code error: 202

2021/07/17 13:13:44 Engine: fetching url http://album.zhenai.com/u/1670931259
2021/07/17 13:13:44 Fetcher: error in fetching url http://album.zhenai.com/u/1670931259: http status code error: 202

2021/07/17 13:13:44 Engine: fetching url http://album.zhenai.com/u/1850286757
2021/07/17 13:13:44 Fetcher: error in fetching url http://album.zhenai.com/u/1850286757: http status code error: 202

2021/07/17 13:13:44 Engine: fetching url http://album.zhenai.com/u/1764840553
2021/07/17 13:13:44 Fetcher: error in fetching url http://album.zhenai.com/u/1764840553: http status code error: 202

2021/07/17 13:13:44 Engine: fetching url http://album.zhenai.com/u/1066492253
2021/07/17 13:13:44 Fetcher: error in fetching url http://album.zhenai.com/u/1066492253: http status code error: 202

```

But you can change to another website applying for crawler, good luck to you.

How to detect a site can be crawled ?

- execute below code and see if the http status is `202`

```go

func Fetch(url string) {
    const userAgent = "xxx" // find a user agent
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent) // to resolve 403

	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		//return nil, errors.New("http status error")
		return nil,
			fmt.Errorf("http status code error: %d\n", resp.StatusCode)
	}
	defer resp.Body.Close()
}
```

## tips

clone this repo & found can't push after you commited in you local cache.

plz using `git push --set-upstream origin master` to set the origin/head then push your code.
