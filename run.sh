git clone https://github.com/briancabbott/bitbucket_client
openapi-generator-cli generate --input-spec ./swagger.json -g go --verbose --config ./client_api_config.yaml --output bitbucket_client