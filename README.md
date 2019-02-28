# alfred-keepassxc

> alfred-keepassxc 功能只有一个 **`查询 Title, 复制密码`**

![demo.png](images/demo.png)

## Todo

+ [x] `golang 1.12 下编译通过`
+ [x] 基础功能实现
+ [ ] 支持 key 认证
+ [ ] 支持模糊搜索

## Usage


在 `Alfred Workflows` 打开 `KeepassXC`,  双击 `Script Filter` 修改对应变量

`shell` 参考 [kpa.sh](kpa.sh)
![](images/usage1.jpg)

### Changelog

**v1.0.2**
1. 使用 ` ` 空格表示 `.*` 通配符，模糊搜索更方便。 `kpa www com` = `kpa www.*com`

**v1.0.1**
1. 支持正则匹配搜索 `kpa www.*com`
