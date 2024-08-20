#!/bin/bash
CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/genprotos
for x in $(find ${CURRENT_DIR}/ReserveDesk_protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/ReserveDesk_protos -I /usr/local/go --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done
