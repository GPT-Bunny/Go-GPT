#!/bin/bash

# 初始化本地仓库
git init

# 将文件添加到暂存区
git add .

# 提交更改
git commit -m "提交说明"

# 连接到远程仓库
git remote add origin https://github.com/GPT-Bunny/Go-GPT.git

# 推送更改
git push -u origin main
