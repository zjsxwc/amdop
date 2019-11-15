
`/tool/bee` 是 beego 的 bee tool 工具,
需要把这个可执行文件放到 `/usr/local/bin/`目录下，
使用方式参考`https://beego.me/docs/install/bee.md`

创建本项目：
```bash
wangchao@wangchao-PC:~/go/src/amdop$ pwd
/home/wangchao/go/src/amdop
wangchao@wangchao-PC:~/go/src/amdop$ bee new src
wangchao@wangchao-PC:~/go/src/amdop$ cd src
```

新创建项目后需要：
1. 把beego admin 的配置文件拷贝到我们这些项目的app.conf里
```bash
vim $GOPATH/src/github.com/beego/admin/conf/app.conf
vim ./conf/app.conf
```
2. 把beego admin的静态资源文件拷贝到我们这个项目里
```bash
wangchao@wangchao-PC:~/go/src/amdop/src$ pwd
/home/wangchao/go/src/amdop/src
wangchao@wangchao-PC:~/go/src/amdop/src$ ls
conf  controllers  main.go  models  routers  src  static  tests  views
wangchao@wangchao-PC:~/go/src/amdop/src$ cp -Rf  $GOPATH/src/github.com/beego/admin/static ./
wangchao@wangchao-PC:~/go/src/amdop/src$ cp -Rf  $GOPATH/src/github.com/beego/admin/views ./
```

3. 编译本项目

```bash
wangchao@wangchao-PC:~/go/src/amdop/src$ pwd
/home/wangchao/go/src/amdop/src
wangchao@wangchao-PC:~/go/src/amdop/src$ go build
```

4. 同步数据库，先到mysql里创建一个叫做`amdop`的数据库
```sql
CREATE DATABASE IF NOT EXISTS amdop DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
```
虽然直接也会创建数据库 但默认数据库字符集不是mb字符集

```bash
wangchao@wangchao-PC:~/go/src/amdop/src$ pwd
/home/wangchao/go/src/amdop/src
wangchao@wangchao-PC:~/go/src/amdop/src$ ./src  -syncdb
```

5. 运行看看
```bash
wangchao@wangchao-PC:~/go/src/amdop/src$ pwd
/home/wangchao/go/src/amdop/src
wangchao@wangchao-PC:~/go/src/amdop/src$ ./src

浏览器访问 http://localhost:8080/
```


beego 编译产物部署时必须要在同目录下有`conf/app.conf`与`static/*`与`views/*`这三样


在 beego admin 的 RBAC 中
node 是每个可以操作的接口资源，比如`UpdateUser`这个node表示更新用户操作
group 是对 node 的分组，node 与 group 是多对多关系
user 就是普通用户
role 是指 user 与 node 的对应关系，user 与 node 是多对多关系，group 只是为了方便 role 分配 node 而出现的


在 beego admin 的 RBAC 中，
RBAC 默认的三层 node 节点“爷父子”的 name 用斜杠分割后就组合成了 http 请求地址，
权限也是靠这种方式处理来鉴定的