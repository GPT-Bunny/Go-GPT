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

# 添加远程仓库，如果远程仓库origin不存在
git remote add origin $REMOTE_REPO_URL 2> /dev/null
if [ $? -ne 0 ]; then
  echo "远程仓库origin已经存在。"
fi

# 添加所有文件到暂存区
git add .

# 提交更改
git commit -m "$COMMIT_MESSAGE"

# 尝试拉取远程仓库的更改并合并
git fetch origin $BRANCH
if git merge-base --is-ancestor origin/$BRANCH HEAD; then
  echo "本地仓库是最新的，无需合并。"
else
  git merge origin/$BRANCH --allow-unrelated-histories -m "Merge remote-tracking branch 'origin/$BRANCH'"
  if [ $? -ne 0 ]; then
    echo "合并失败，请手动解决冲突后再次运行脚本。"
    exit 1
  fi
fi

# 推送到远程仓库
git push -u origin $BRANCH
if [ $? -eq 0 ]; then
  echo "代码已推送到远程仓库。"
else
  echo "推送失败。"
  exit 1
fi
