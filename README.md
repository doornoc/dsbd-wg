# dsbd-wg
**wireguardのpeerをAdd,Delete,Getするアプリケーション**
### Add
**POST /api/v1/peer**
```json
{
  "public_key": "3wom4zptIT7Mc1UgEs5rmNbKwyYa/AXSlDoH/XitEFc=",
  "endpoint": "x.x.x.x:xxxx",
  "allowed_ips": [
    "x.x.x.x/32",
  ]
}
```

##### curlを使ったPOSTコマンド
```
curl -X POST -H "Content-Type: application/json" -d '{}' 
```


### Delete
**Delete /api/v1/peer**
```json
{
  "public_key": "3wom4zptIT7Mc1UgEs5rmNbKwyYa/AXSlDoH/XitEFc="
}
```

##### curlを使ったDELETEコマンド
```
curl -X DELETE -H "Content-Type: application/json" -d '{}' 
```


### Get
**GET /api/v1/peer**
