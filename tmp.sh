rm -rf src/kafka_bridge
docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/resources/openapi.json \
    -g go \
    -o /local/src/kafka_bridge \
    --package-name kafka_bridge \
    --additional-properties="withGoMod=false,structPrefix=true" \
    --git-user-id tihaya-anon \
    --git-repo-id tx_sys-event-event_repository/src/kafka_bridge