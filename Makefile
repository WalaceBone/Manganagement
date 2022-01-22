.PHONY: build

s3bucket		=	manganagement
stage			=	live
prefix_bucket	=	dev
stack_name		=	manganagement
region			=	eu-west-3

all: build deploy

build:
	sam build

deploy:
	aws s3 cp spec/api-spec.yaml s3://$(s3bucket)/spec/api-spec.yaml
	sam package --s3-bucket $(s3bucket) --output-template-file packaged.yaml
	sam deploy packaged.yaml \
	--capabilities CAPABILITY_IAM \
	--parameter-overrides StageName=$(stage) BucketPrefix=$(prefix_bucket) CicdBucket=$(s3bucket) \
	--s3-bucket $(s3bucket) --stack-name $(stack_name)-$(stage) \
	--region $(region)

invoke:
	sam local invoke

deployg:
	sam deploy --guided

test:
	sam local start-api

clean: clean-stack
	rm -rf .aws-sam
	rm -rf package.yaml

clean-stack:
	aws cloudformation delete-stack --stack-name $(stack_name)-$(stage)