gen-docs:
	SWAGGER_GENERATE_EXTENSION=false swagger generate spec --scan-models --output=./docs/swagger.yaml