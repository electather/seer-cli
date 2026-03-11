docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/open-api.yaml \
    -g go \
    -o /local/pkg/api \
    --additional-properties=packageName=api,isGoSubmodule=true,moduleName=seer-cli/pkg/api