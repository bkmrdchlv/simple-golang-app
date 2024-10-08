#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
  sudo protoc --plugin="protoc-gen-go=${GOPATH}/bin/protoc-gen-go" --plugin="protoc-gen-go-grpc=${GOPATH}/bin/protoc-gen-go-grpc" -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done

for module in $(find $CURRENT_DIR/genproto/* -type d); do
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sudo sed -i "" -e "s/,omitempty//g" $module/*.go
    else
      sudo sed -i -e "s/,omitempty//g" $module/*.go
  fi
done;
