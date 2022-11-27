# Varint的Go实现 



# 一、这是什么

可变长int的类型，类似于数据库中的varchar类型。

# 什么样子的数据适合使用varint 

值的分布范围很大，并且绝大多数的值都分布在较小的那边，对于这种类型的树，直接使用偏移收缩不合适，直接使用较小的数据类型不合适大的放不下了，那么就是用varint类型，使得绝大多数数据都能消耗较小的空间，并且部分有需要的仍然使用较大的空间。

## 应用场景

- 无符号证书持久化存储时节省空间
- 内存计算节省空间，支持更大规模的计算量 

# 二、安装

```bash
go get -u github.com/compression-algorithm-research-lab/go-varint
```

# 三、Example 

TODO 2022-11-27 21:57:06 













