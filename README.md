# naruto-go 

## 路由

## 数据库表

#### sqlc 命令

```bash
sqlc generate
```

#### go-migrate 命令

安装 go-migrate
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

```bash
# 创建 migrate 文件
migrate create -ext sql -dir config/migrations -seq create_users_table
# 升级
migrate -database "postgres://naruto:123456@192.168.110.109:5432/naruto_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" up

```



