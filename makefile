.PHONY: deploy

compose:
	docker-compose rm -v -f
	docker-compose up --build

table:
	rm -rf TABLE.md && node table/index.js

dev:
	yarn start

build:
	docker build -t partner-matrix .

run:
	docker run -p 3000:80 partner-matrix

deploy: build
	docker tag partner-matrix 164382793440.dkr.ecr.us-west-2.amazonaws.com/partner-matrix:latest
	docker push 164382793440.dkr.ecr.us-west-2.amazonaws.com/partner-matrix:latest
	# aws cloudformation update-stack --stack-name partner-matrix --template-body file://deploy/cloudformation.yaml --parameters ParameterKey=VPC,ParameterValue=vpc-0470143a7c277ae00 ParameterKey=SubnetA,ParameterValue=subnet-0f5359d858ba4098d ParameterKey=SubnetB,ParameterValue=subnet-0548de7d8ab9a2c5e --capabilities CAPABILITY_NAMED_IAM
	aws ecs update-service --cluster partner-matrixCluster --service partner-matrix --force-new-deployment
