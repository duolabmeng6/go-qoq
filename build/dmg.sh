#!/bin/bash

# 定义应用程序名称和版本号
APP_NAME="qoq"
# 环境变量 version
APP_VERSION="$version"

# 创建临时目录和镜像文件名
TEMP_DIR="./"
DMG_NAME="${APP_NAME}-${APP_VERSION}.dmg"

# 创建 Applications 文件夹的快捷方式
ln -s "/Applications" "${TEMP_DIR}/Applications"

# 创建空的镜像文件
hdiutil create -srcfolder "${TEMP_DIR}" -volname "${APP_NAME}" -format UDZO "${DMG_NAME}"

# 删除临时目录
rm -rf "${TEMP_DIR}"

echo "${DMG_NAME} 已经创建完成！"