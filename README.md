## 说明

gin的基础框架


## 部署与管理

- 本地编译

```
go get github.com/13sai/gin-frame
```

配置参考：
```
name: demo2-gin-frame
db:
  name: blog
  host: 127.0.0.1:3306
  username: root
  password: 111111
  charset: utf8mb4
  MaxIdleConns: 10
  MaxOpenConns: 2
  log: true
redis:
  host: 127.0.0.1
  port: 6379
  auth:
mode: debug
addr: :8080
max_ping_count: 4
log:
  level: debug
  file_format: "%Y%m%d"
  max_save_days: 30
  file_type: one #one, level
```

main.go可参照app/app.go