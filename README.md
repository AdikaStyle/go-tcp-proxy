# go-tcp-proxy

Sometime you need a tcp proxy, this will get you started fast.
Write a configuration file:
```json
{
  "routes": [
    {
      "name": "Google Redirect",
      "from": 1234,
      "to": "216.58.208.46:80"
    }
  ]
}
```

And start the proxy with it:
```
$> go-tcp-proxy path/to/config.json

2020/11/22 10:06:00 Loaded Rule: Google Redirect (1234 -> 216.58.208.46:80)
```

And that's about it, you've just created a proxy from port 1234 to 216.58.208.46:80.