# 构建 go

## 构建所有的 go 项目

```shell
 make go.build
```

## 构建指定的 go 项目

```shell
 make go.build BINS=krm-apiserver
```

## 构建所有的平台 go 项目
```shell
make go.build.multiarch
```