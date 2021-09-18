把下面json保存为“users.json”，然后写一个main.go 从文件读取并解析成user。
```
{
  "users": [
    {
      "name": "Elliot",
      "type": "Reader",
      "age": 23,
      "social": {
        "facebook": "https://facebook.com",
        "twitter": "https://twitter.com"
      }
    },
    {
      "name": "Fraser",
      "type": "Author",
      "age": 17,
      "social": {
        "facebook": "https://facebook.com",
        "twitter": "https://twitter.com"
      }
    }
  ]
}
```

执行结果如下
`$ go run main.go`
```shell
Successfully opened users.json
User Type: Reader
User Age: 23
User Name: Elliot
Facebook Url: https://facebook.com
User Type: Author
User Age: 17
User Name: Fraser
Facebook Url: https://facebook.com
{
  [
    {
      Elliot Reader 23 {https://facebook.com https://twitter.com}
    } 
    {
      Fraser Author 17 {https://facebook.com https://twitter.com}
    }
  ]
}
```

