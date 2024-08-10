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
