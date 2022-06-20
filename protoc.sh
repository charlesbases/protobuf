#!/usr/bin/env bash

dir="google/protobuf/extend/"
files=`ls ${dir} | grep '.proto'`

for i in ${files}; do
  prefix=`echo ${i} | cut -d "." -f1`

  protoc -I=. --gogo_out=${GOPATH}/src ${dir}/${i}
done
