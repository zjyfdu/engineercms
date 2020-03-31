# 是什么
- 一个文档在线编辑平台，部署在局域网，可以用作组内日报/周报系统
- 其实是从[3xxx/engineercms](https://github.com/3xxx/engineercms) fork来的，只保留了文档编辑的功能

# 怎么用
- docker启动onlyoffice `docker run -i -t -d -p 9000:80 onlyoffice/documentserver`
- 修改conf/app.conf中的ip配置，注意不能写localhost，不然局域网的其他用户访问不到
```
documentserverip = http://192.168.1.107:9000
entryserverip = http://192.168.1.107:8088
```
- 解压main.rar，运行exe，浏览器端访问entryserverip

# 代码结构

- onlyoffice负责文档编辑的功能，onlyoffice的api可以参考[这里](https://www.jianshu.com/p/2d4f977ffeac)
- 本项目负责文档的上传、删除、用户管理等，采用beego MVC架构

# 相比[3xxx/engineercms](https://github.com/3xxx/engineercms)做了哪些修改

修改主要是为了更适合做日报系统
## 删除了文档编辑以外的功能
## 文档自动上传

- 日报中包含日期，如“总裁办公室20200331.xlsx”，会自动复制日期最新的一份日报，并对日报做清空后自动上传

- 那些内容清空？

  只保留第一个sheet，sheet命名为日期（MMdd）
  
  conf/app.conf中，设置那些行和列以内保留，≤rowkeep和colkeep的内容保留
```
#####日报清空，前多少行和列保留，固定只保留第一个sheet
rowkeep = 1
colkeep = 4
```

- 什么时候上传？

  每次打开文档列表也都判断是否有今天的日报，没有就自动上传

  节假日没有人访问，所以避免了逻辑判断是否上传

## 文档自动保存

- 原来文档不会自动保存吗？
  
  编辑后的文档都暂时存在onlyoffice的docker中
  
  只有在所有人关闭文档的情况下，才会将文档内容保存到本地

  做不到全部用户及时关闭文档，而如果docker崩了，文档内容就丢失了

- 如何解决？
  
  在没有用户使用的情况下强制保存。conf/app.conf中设置了定时保存，默认设置是晚上10点，每10分钟保存一次

```
#####文档保存本地
forcesavecron = "0 */10 22 * * *"
###              秒 分  时  日 月 周
```
  

## 登录权限
- 内网用户比较可靠，不做权限限制，不需要注册账户
- 可以不登录用ip作为用户名
- 可以“假登录”，直接登录，不检查用户是否存在，不检查密码