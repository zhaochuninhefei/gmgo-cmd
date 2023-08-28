gitee.com/zhaochuninhefei/gmgo-cmd
====================

gmgo的命令行工具

# 编译
```sh
# cd 到当前目录后，执行编译脚本，并将编译后的可执行文件拷贝到${GOPATH}/bin目录下
# 默认${GOPATH}/bin已经加入本地环境变量PATH中
./build_copy_go_path.sh
```


# 查看gmgo生成的x509证书
目前支持sm2/ecdsa/ecdsa_ext三种签名算法。

```sh
# 查看sm2证书
gmgo-cmd x509 --text --in testdata/sm2_sign_cert.cer

# 查看ecdsa证书
gmgo-cmd x509 --text --in testdata/ecdsa_sign_cert.cer

# 查看ecdsaext证书
gmgo-cmd x509 --text --in testdata/ecdsaext_sign_cert.cer

```

# 生成口令
```sh
$ gmgo-cmd pwd -h
使用gmgo的口令生成器,支持大小写字母、数字和部分特殊符号(~!@#$%^&_-+=|:;)

Usage:
  gmgo-cmd pwd [flags]

Flags:
  -d, --display @                    显示口令，包含@时严格匹配口令键值，不包含`@`时作为口令键值的后缀查找，传入all时显示全部口令
  -h, --help                         help for pwd
  -k, --key 用户名@目标域名   口令保存键值，格式: 用户名@目标域名，如`testuser@test.com`
  -l, --length int                   口令长度(至少为4)
  -s, --strength int                 口令强度(1:大小写字母+数字, 2:大小写字母+数字+特殊符号, 默认:2)

```


# JetBrains support
Thanks to JetBrains for supporting open source projects.

<a href="https://jb.gg/OpenSourceSupport" target="_blank">https://jb.gg/OpenSourceSupport.</a>