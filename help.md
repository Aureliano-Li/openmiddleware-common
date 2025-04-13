
git add .


git remote add origin git@gitlab.com:yourusername/yourproject.git


git commit -m "Add .gitignore and ignore vendor directory"


# 仅从暂存区移除 vendor 目录
git rm -r --cached vendor

# 提交更改
git commit -m "Remove vendor directory from staging area"

# 添加 .gitignore 规则
echo "vendor/" >> .gitignore

# 添加 .gitignore 文件到暂存区并提交
git add .gitignore

git commit -m "Add vendor directory to .gitignore"

# 推送更改到远程仓库
git push origin master


重置git认证信息
# 若需要重新输入认证信息，可先清除旧的认证信息
git config --global --unset credential.helper

# 再次设置凭证存储
git config --global credential.helper store

# 推送代码
git push -u origin master

