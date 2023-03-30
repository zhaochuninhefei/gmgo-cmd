gitee.com/zhaochuninhefei/gmgo-cmd
====================

gmgo的命令行工具

# 查看gmgo生成的x509证书
目前支持sm2/ecdsa/ecdsa_ext三种签名算法。

```sh
# cd 到当前目录
go build

# 查看sm2证书
./gmgo-cmd x509 --text --in testdata/sm2_sign_cert.cer

# 查看ecdsa证书
./gmgo-cmd x509 --text --in testdata/ecdsa_sign_cert.cer

# 查看ecdsaext证书
./gmgo-cmd x509 --text --in testdata/ecdsaext_sign_cert.cer

```


# JetBrains support
Thanks to JetBrains for supporting open source projects.

<a href="https://jb.gg/OpenSourceSupport" target="_blank">https://jb.gg/OpenSourceSupport.</a>