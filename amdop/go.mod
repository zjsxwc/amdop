module amdop

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/beego/admin v0.0.0-20171002082758-83609ddd6f2d
	github.com/liudng/godump v0.0.0-20150708094948-5c7e73aafb21
	github.com/smartystreets/goconvey v1.6.4
)

replace github.com/beego/admin => ../localpkg/admin
