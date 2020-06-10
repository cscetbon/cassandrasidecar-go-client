#! /bin/bash

export wv_go_path="module github.com/cscetbon/cassandrasidecar-go-client"
export workspace_dir=$PWD
export wv_tmp_dir=${workspace_dir}/pkg

echo Generating Cassandra Sidecar Go client

rm -fr ${workspace_dir}/pkg && mkdir -p ${workspace_dir}/pkg/cassandrasidecar
cd resources/client_gen
./generate_api_client.sh
cd -

echo Cleaning Cassandra Sidecar Go client
mv ${wv_tmp_dir}/cassandrasidecar/go-client/* ${wv_tmp_dir}/cassandrasidecar && rm -rf ${wv_tmp_dir}/cassandrasidecar/go-client
rm -f ${wv_tmp_dir}/cassandrasidecar/go.sum
rm -f ${wv_tmp_dir}/cassandrasidecar/git_push.sh

echo Moving documentations files to root dirs
rm -fr docs && mv ${wv_tmp_dir}/cassandrasidecar/docs .
mv ${wv_tmp_dir}/cassandrasidecar/README.md .
mv ${wv_tmp_dir}/cassandrasidecar/go.mod .
sed -i.bak "1s#.*#$wv_go_path#" go.mod && echo "go 1.14" >> go.mod
