module amdop

go 1.12

require (
	github.com/astaxie/beego v1.12.2
	github.com/beego/admin v0.0.0-20171002082758-83609ddd6f2d
	github.com/smartystreets/goconvey v1.6.4
	github.com/syyongx/php2go v0.9.4
	onerequest v0.0.0-00010101000000-000000000000
)

replace github.com/beego/admin => ../localpkg/admin

replace onerequest => ../localpkg/onerequest
