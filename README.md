# heavenMs-NAP-golang

## 开发规范
1. 自行添加githook,每次有提交前，会自动检查能否编译
    ```shell
    # 在项目目录下,执行以下命令
    cd .git/hooks
    vim pre-commit
    # 将以下内容复制进去
    
    #/bin/bash
    set -ex
    go mod tidy
    go build -o GitHookPreCommitCheck ./main.go
    rm GitHookPreCommitCheck
    ```

## 依赖文件

### [WZ文件](https://github.com/Galen-Chen177/heavenMs-NAP-golang/releases/download/wz/wz_root.zip)

## 客户端下载
083cn初始英文端：https://pan.baidu.com/s/1GAgyysoRqKsfv-ODvnGkfA 提取码：ysn1

083cn初始英文端：https://pan.baidu.com/s/17z3pBKu3jz5AOJmB9eFzqQ 提取码：f5zc

70%汉化小体积端: https://pan.baidu.com/s/1qhEgKyvz7HqWPdSfK8tkVA 提取码: ruct ——一个热心网友提供的客户端