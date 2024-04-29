# my_local_communitate

### 项目说明

用于局域网加密传输文件工具，基于wails，使用Go+Vue+Element ui开发的桌面应用

- 支持在局域网传输任意文件，传输的文件为加密传输
  - 采用diffie-hellman进行密钥交换，双方得到对称密钥
  - 对称加密采用DES算法
  - 每次文件传输密钥会重新生成密钥
- LRU缓存本地密钥，最多不超过100个

### 如何使用

参考官网[安装 | Wails](https://wails.io/zh-Hans/docs/gettingstarted/installation)

在my_local_communitate目录下输入以下命令即可

```
wails build
```
在build/bin目录下生成相应程序
