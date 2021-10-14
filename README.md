# go-vue-api

# 项目初始化

在项目文件目录中生成go.mod文件

```
go mod init go-vue-api
```

![image-20211013153902210](README.assets/image-20211013153902210.png)

安装Gin

```
go get -u github.com/gin-gonic/gin
```

-u标识安装最新版本



安装配置仓库

```
go get gopkg.in/ini.v1
```

![image-20211013160852709](README.assets/image-20211013160852709.png)

安装日志仓库

```
go get -u github.com/sirupsen/logrus
```



安装gorm包

安装gorm MYSQL驱动

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

安装验证器

```
go get github.com/go-playground/validator/v10
```

安装in18国际化

```
go get github.com/go-playground/universal-translator
go get github.com/go-playground/locales
```

