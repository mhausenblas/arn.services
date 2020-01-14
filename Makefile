ARNS_BUCKET?=arnservices-code
ARNS_STACK?=arnservices
# APIGATEWAY_ENDPOINT:=$(shell aws cloudformation describe-stacks --stack-name ${ARNS_STACK} --query "Stacks[0].Outputs[?OutputKey=='ASEndpoint'].OutputValue" --output text)
EXPLODE_ENDPOINT:=${APIGATEWAY_ENDPOINT}/explode
GENERATE_WEBHOOK_ENDPOINT:=${APIGATEWAY_ENDPOINT}/generate

.PHONY: build buildexplode buildgenerate up deploy destroy status

build: buildexplode buildgenerate

buildexplode:
	GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/explode ./explode

buildgenerate:
	GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/generate ./generate

up: 
	sam package --template-file template.yaml --output-template-file current-stack.yaml --s3-bucket ${ARNS_BUCKET}
	sam deploy --template-file current-stack.yaml --stack-name arnservices --capabilities CAPABILITY_IAM

deploy: build up

destroy:
	aws cloudformation delete-stack --stack-name ${ARNS_STACK}

status:
	aws cloudformation describe-stacks --stack-name ${ARNS_STACK}