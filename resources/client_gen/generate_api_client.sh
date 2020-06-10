#!/usr/bin/env bash

# Instructions
# You will need Java installed on the local machine

# A new client is generated from the Swagger Def, and then manually merged into
# NiPyApi so that changes may be assessed, tests written, and trouble avoided

# Params
echo Exporting Params
export wv_client_name=${wv_client_name:-cassandrasidecar}

export wv_codegen_filename=${wv_codegen_filename:-swagger-codegen-cli-3.0.20.jar}
export wv_tmp_dir=${wv_tmp_dir:-${HOME}/Projects/tmp}
export wv_client_dir=${wv_tmp_dir}/${wv_client_name}

export wv_openapi_yaml_url=https://raw.githubusercontent.com/instaclustr/cassandra-sidecar/master/spec.yaml
export wv_generator_url=http://api.openapi-generator.tech/api/gen

echo Downloading go client for ${wv_openapi_yaml_url} from ${wv_generator_url}

code=$(curl -X POST "${wv_generator_url}/clients/go" \
  -H "Accept: */*" \
  -H "Content-Type: application/json" \
  -d "{ \"openAPIUrl\": \""${wv_openapi_yaml_url}"\", \"options\": { \"packageName\": \""${wv_client_name}"\", \"packageVersion\": \"1.0.0\" }}" | jq --raw-output '.code')


echo Downloading generated code

mkdir tmp
wget "${wv_generator_url}/download/${code}" -P tmp
rm -fr ${wv_client_dir}
unzip "tmp/${code}" -d ${wv_client_dir}
rm -fr tmp


