#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0); pwd)
rm -rf "${SCRIPT_DIR}/../src/kafka_bridge"
docker run --rm -v "${SCRIPT_DIR}/../:/local" openapitools/openapi-generator-cli generate \
  -i /local/resources/openapi.json \
  -g go \
  --additional-properties="packageName=kafka_bridge,withGoMod=false,structPrefix=true" \
  -o /local/src/kafka_bridge \
  --git-user-id tihaya-anon \
  --git-repo-id tx_sys-event-event_repository/src/kafka_bridge
