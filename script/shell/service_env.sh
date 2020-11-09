#!/usr/bin/env bash
# 设置微服务的环境
echo "开始配置服务环境"

echo "配置sftp环境"
if [ ! -d /home/ftp ];then
  mkdir -p /home/ftp/sftp
  usermod -d /home/ftp/sftp sftp
  chown root:root /home/ftp
  chmod 755 /home/ftp/sftp
  mkdir /home/ftp/sftp/upload
  chmod 755 /home/ftp/sftp/upload
  echo "创建sftp路径完毕"

  groupadd sftp
  echo "创建sftp用户组"
  useradd -g sftp -s /bin/false sftp
  echo "创建sftp用户"
  echo "请自行修改密码"

  echo "开始配置sftp路径权限"
  chown root:sftp /home/ftp/sftp
  chown sftp:sftp /home/ftp/sftp/upload
  echo "权限配置完毕"
  echo "sshd Chroot路径配置为/home/ftp/%u"
  echo "在配置完成sshd后重启sshd生效"

else
  echo "sftp路径已存在"
fi

echo "配置apps路径"
if [ ! -d /home/apps ];then
  mkdir /home/apps
  chmod 666 -R /home/apps
  echo "创建apps路径完毕"
else
  echo "apps路径已存在"
fi

echo "配置nginx缓存"
if [ ! -d /home/nginx ];then
  mkdir -p /home/nginx/cache
  chmod 666 -R /home/nginx/cache
  mkdir -p /home/nginx/img
  chmod 666 -R /home/nginx/img
  echo "创建nginx缓存路径完毕"
else
  echo "nginx缓存路径已存在"
fi

echo "配置mgek file服务"
if [ ! -d /home/file ];then
  mkdir -p /home/file/dl
  chmod 644 -R /home/file/dl
  mkdir -p /home/file/images
  chmod 644 -R /home/file/images
  echo "创建mgek file路径完毕"
else
  echo "mgek file路径已存在"
fi

echo "检查python3环境"
python3 -V
if [ $? != 0 ] && [ ! -d /usr/local/python3 ];then
  echo "不存在python3环境"
  version=https://www.python.org/ftp/python/3.7.6/Python-3.7.6.tgz
  echo "开始下载python3, 版本$version"
  cd /home && mkdir tmp && wget $version
  tar -xzvf Python-3.7.6.tgz
  echo "解压完毕, 开始编译python3"
  echo "默认安装路径/usr/local"
  mkdir /usr/local/python3
  cd Python-3.7.6 && ./configure --prefix=/usr/local/python3 && make && make install
  echo "如果出现错误, 请检查python3依赖包是否安装"
  echo "创建软链接"
  ln -s /usr/local/python3/bin/python3 /usr/bin/python3
  ln -s /usr/local/python3/bin/pip3 /usr/bin/pip3
  echo "python3安装完毕"
else
  echo "python3环境已满足"
fi