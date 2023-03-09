#!/bin/bash

# key 为项目名称，value 为yapi token
declare -A swagMap
# 由于基础信息生成默认main.go文件，所以需要指定一个接口目录下的一个go文件。任意文件即可
declare -A fileMap
swagMap["auth"]="0f9195fe3a9d44dea349f27067b1348b817ea331db259427095860da3150e7cb"
fileMap["auth"]="auth_controller.go"

swagMap["operation"]="0b50b0b4d03f06bee06cdda91309e4f2d94352dce0dc2ea8d7b3d34d3f52e458"
fileMap["auth"]="operation_controller.go"

url='http://yapi.smart-xwork.cn/'
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
