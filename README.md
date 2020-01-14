# arn.services

A utility service for Amazon Resource Names ([ARNs](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)), providing the functionality 
as listed below.

You can consume it as an online service via `arn.services` or run yourself as a
[SAM application](https://aws.amazon.com/serverless/sam/). For the former, I'm
[currently working on it](https://github.com/mhausenblas/arn.services/issues/1) and for the latter you need to: 1. clone this repo, 2. create an S3 bucket for the Lambda functions, 3. change the value of `ARNS_BUCKET` to your own S3 bucket in the [Makefile](https://github.com/mhausenblas/arn.services/blob/master/Makefile), and 4. execute `make deploy`.

## Compartmentalize

In order to compartmentalize an ARN, that is, to decompose it into its parts, use a `GET` on `explode/$ARN`, for example:

```sh
$ curl -s arn.services/explode/arn:aws:s3:us-west-2::abucket | jq .
{
  "Partition": "aws",
  "Service": "s3",
  "Region": "us-west-2",
  "AccountID": "",
  "Resource": "abucket"
}
```

## Generate

If you have (certain) components of an ARN and want to generate a fully qualified ARN, use a `POST` to `generate/`, for example:

```sh
$ curl -s -X POST \
       -H "Content-Type: application/json"
       -d '{"Service":"s3", "Resource":"somebucket/someobject"}' \
       arn.services/generate
arn:aws:s3:us-west-2::somebucket/someobject
```
