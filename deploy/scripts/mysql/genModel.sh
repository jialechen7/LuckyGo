#!/usr/bin/env bash

# ===========================
# 脚本说明：
# 用于根据数据库表生成 Go 语言 Model 文件
# 使用示例：
# ./genModel.sh usercenter user ./genModel
# ./genModel.sh usercenter user_auth ./genModel ./custom-goctl-template
# 生成的文件需要手动移动到目标服务的 model 目录并修改 package 名称 或 直接指定生成目录为目标服务的 model 目录
# ===========================

# 使脚本在遇到错误时退出
set -e

# 数据库配置
HOST="127.0.0.1"
PORT="3306"
USERNAME="root"
PASSWORD="root"

# 参数检查
if [[ $# -lt 2 ]]; then
  printf "Usage: %s <service-name> <table-name> [model-dir] [goctl-template-path]\n" "$(basename "$0")"
  exit 1
fi

# 获取脚本参数
SERVICE_NAME="$1"
TABLE_NAME="$2"
MODEL_DIR="${3:-./genModel}" # 如果没有提供 modelDir 参数，默认为 ./genModel
TEMPLATE_PATH="${4}"         # 第四个参数：goctl 模板路径

# 数据库名称
DB_NAME="go-lottery-${SERVICE_NAME}"

# 输出提示信息
printf "开始为数据库 '%s' 中的表 '%s' 生成模型...\n" "${DB_NAME}" "${TABLE_NAME}"

# 构造 goctl 命令
GOCTL_CMD="goctl model mysql datasource \
  -url=\"${USERNAME}:${PASSWORD}@tcp(${HOST}:${PORT})/${DB_NAME}\" \
  -table=\"${TABLE_NAME}\" \
  -dir=\"${MODEL_DIR}\" \
  -cache=true \
  --style=go_zero"

# 如果指定了模板路径，则添加 --home 参数
if [[ -n "${TEMPLATE_PATH}" ]]; then
  GOCTL_CMD="${GOCTL_CMD} --home=\"${TEMPLATE_PATH}\""
  printf "使用自定义 goctl 模板路径：%s\n" "${TEMPLATE_PATH}"
fi

# 执行命令
eval "${GOCTL_CMD}"

# 输出完成提示
printf "模型生成完成，文件保存在目录：%s\n" "${MODEL_DIR}"
