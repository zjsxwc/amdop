
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

6. beego如何通过反射读取配置

beego这里的配置就是 /home/wangchao/go/pkg/mod/github.com/astaxie/beego@v1.12.0/config.go 里结构体每个字段的名字=值

读取配置文件具体实现看     /home/wangchao/go/pkg/mod/github.com/astaxie/beego@v1.12.0/config.go:280

```
for _, i := range []interface{}{BConfig, &BConfig.Listen, &BConfig.WebConfig, &BConfig.Log, &BConfig.WebConfig.Session} {
   		assignSingleConfig(i, ac)
   	}


func assignSingleConfig(p interface{}, ac config.Configer) {
	pt := reflect.TypeOf(p)
	if pt.Kind() != reflect.Ptr {
		return
	}
	pt = pt.Elem()
	if pt.Kind() != reflect.Struct {
		return
	}
	pv := reflect.ValueOf(p).Elem()

	for i := 0; i < pt.NumField(); i++ {
		pf := pv.Field(i)
		if !pf.CanSet() {
			continue
		}
		name := pt.Field(i).Name
		switch pf.Kind() {
		case reflect.String:
			pf.SetString(ac.DefaultString(name, pf.String()))
		case reflect.Int, reflect.Int64:
			pf.SetInt(ac.DefaultInt64(name, pf.Int()))
		case reflect.Bool:
			pf.SetBool(ac.DefaultBool(name, pf.Bool()))
		case reflect.Struct:
		default:
			//do nothing here
		}
	}

}


// DefaultString returns the string value for a given key.
// if err != nil return defaultval
func (c *IniConfigContainer) DefaultString(key string, defaultval string) string {
	v := c.String(key)
	if v == "" {
		return defaultval
	}
	return v
}

// String returns the string value for a given key.
func (c *IniConfigContainer) String(key string) string {
	return c.getdata(key)
}



// section.key or key
func (c *IniConfigContainer) getdata(key string) string {
	if len(key) == 0 {
		return ""
	}
	c.RLock()
	defer c.RUnlock()

	var (
		section, k string
		sectionKey = strings.Split(strings.ToLower(key), "::") //《------------- struct的字段名变小写了
	)
	if len(sectionKey) >= 2 {
		section = sectionKey[0]
		k = sectionKey[1]
	} else {
		section = defaultSection
		k = sectionKey[0]
	}
	if v, ok := c.data[section]; ok {
		if vv, ok := v[k]; ok {
			return vv
		}
	}
	return ""
}

在解析ini  conf文件时都 不区分大小写，section全变成小写先
if bytes.HasPrefix(line, sectionStart) && bytes.HasSuffix(line, sectionEnd) {
			section = strings.ToLower(string(line[1 : len(line)-1])) // section name case insensitive
			if comment.Len() > 0 {
				cfg.sectionComment[section] = comment.String()
				comment.Reset()
			}
			if _, ok := cfg.data[section]; !ok {
				cfg.data[section] = make(map[string]string)
			}
			continue
		}

```


7.如何数据迁移

创建数据迁移代码
 1141  bee generate migration user
修改新创建的`database/migrations/20191115_151429_user.go` 写sql 
 1145  bee help migrate 
执行所有迁移 
 1146  bee migrate -conn="root:root@tcp(192.168.33.77:3306)/amdop"
回滚所有迁移
 1147  bee migrate reset -conn="root:root@tcp(192.168.33.77:3306)/amdop"
回滚一次最近的迁移
 1147  bee migrate rollback -conn="root:root@tcp(192.168.33.77:3306)/amdop"
全部撤销迁移后重新执行一次迁移
 1147  bee migrate refresh -conn="root:root@tcp(192.168.33.77:3306)/amdop"



8. 加接口获取当前员工权限

GET  /public/getaccesslist 

去掉第一个/ 以及全小写后比较map里是否有这个key
