#!/bin/bash

# key 为项目名称，value 为yapi token
declare -A swagMap
# 由于基础信息生成默认main.go文件，所以需要指定一个接口目录下的一个go文件。任意文件即可
declare -A fileMap
swagMap["auth"]="27845ae85ebd2974b61478705b575a7d48ad5f538c27c6b99777b15456782d74"
fileMap["auth"]="auth_controller.go"

swagMap["operation"]="db575e62667e710556a92e5b81eaf4c3509102d627eaeaa1105f430b557438b5"
fileMap["operation"]="operation_controller.go"

url='http://47.96.85.188:3000'
for key in ${!swagMap[@]}; do
  if [[ $# == 1 && $1 != $key ]]; then
    continue
  fi

  swag init --parseDependency --parseDepth=6 --dir ./internal/service/$key --output ./docs/$key -g ${fileMap[$key]}
  json=`cat ./docs/$key/swagger.json | sed -n -e 'H;${x;s/    //g;p;}' | sed -n -e 'H;${x;s/\n//g;p;}'`

  curl -H "Content-type: application/x-www-form-urlencoded" \
   -H "Expect: 100-continue" \
   -X POST \
   -d "json=${json}" \
   -d 'merge=merge' \
   -d 'type=swagger' \
   -d "token=${swagMap[$key]}" \
   "${url}/api/open/import_data"
done
