#!/bin/bash

# 配置
REMOTE_REPO_URL="https://github.com/GPT-Bunny/Go.git" # 替换为您的远程仓库URL
BRANCH="main" # 或者如果您的默认分支是main，则使用main
COMMIT_MESSAGE="Initial commit" # 提交信息

# 检查Git是否已安装
if ! [ -x "$(command -v git)" ]; then
  echo "错误：Git没有安装。" >&2
  exit 1
fi

# 初始化本地仓库
git init

# 添加远程仓库
git remote add origin $REMOTE_REPO_URL

# 添加所有文件到暂存区
git add .

# 提交更改
git commit -m "$COMMIT_MESSAGE"

# 拉取远程仓库的更改并合并/重置（如果需要）
# git pull origin $BRANCH --allow-unrelated-histories

# 推送到远程仓库
git push -u origin $BRANCH

echo "代码已推送到远程仓库。"
