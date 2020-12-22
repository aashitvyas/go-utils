import boto3
import json

ec2client = boto3.client('ec2')
response = ec2client.describe_security_groups()
for r in response.get('SecurityGroups'):
    for grouperule in r:
        print(grouperule)
    # print(r['IpPermissions'])
    break
