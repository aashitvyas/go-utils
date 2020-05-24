all: 
	@sh validate_template.sh;zip main.zip main.yaml;aws s3 sync . s3://ecs-cf-ca-central-1/ --delete

