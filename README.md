使用beego admin


cd ~/amdop/amdop
GO111MODULE=on GOPROXY=https://goproxy.io  go mod tidy
GO111MODULE=on go build

./amdop -syncdb
./amdop




如何把三方包拿到本地改，先复制的localpkg目录下，然后进入三方包，

执行 `GO111MODULE=on go mod init github.com/beego/admin`

注意这里必须是 `github.com/beego/admin` 也就是三方包自己已经对外使用的名称

然后在三方包目录下执行 `GO111MODULE=on GOPROXY=https://goproxy.io  go mod tidy`

在当下项目里的`go.mod` 里添加一句`replace github.com/beego/admin => ../localpkg/admin`

执行 `GO111MODULE=on GOPROXY=https://goproxy.io  go mod tidy`与
`GO111MODULE=on go build`