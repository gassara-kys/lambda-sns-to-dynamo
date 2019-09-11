# lambda-sns-to-dynamo

## Pre-Require
- Create SNS topic.
- Create DynamoDB Table.(ex. `sns_alert` table)

- Set the required IAMRole
  - Read to the SNS topic
  - Write to the DynamoDB Table resource

## Run Local
```bash
# make sure that the required IAM is set in advance
$ make
```

## Lambda zip file
```bash
$ make zip
```