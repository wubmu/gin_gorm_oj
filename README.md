
GOLANG下载网址： https://golang.google.cn/dl/

参考文档 Module：https://www.kancloud.cn/aceld/golang/1958311

GORM中文官网：https://gorm.io/zh_CN/docs/

GIN中文官网：https://gin-gonic.com/zh-cn/docs/

### 整合Swagger
参考文档： https://github.com/swaggo/gin-swagger 接口访问地址：http://localhost:8080/swagger/index.html
```go
// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem-list [get]
```

## 安装JWT
```shell
go get github.com/dgrijalva/jwt-go
```