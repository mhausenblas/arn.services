# arn.services

A utility service for Amazon Resource Names ([ARNs](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)), providing the functionality 
as listed below. You can consume it as an online service via `https://arn.services` or run it yourself as a
[SAM application](https://aws.amazon.com/serverless/sam/), see details below in the [run it yourself](#run-it-yourself) section.

## Use it

ARN Services has two features: you can [break up an ARN into its components](#compartmentalize)
and you can [generate an ARN from its components](#generate). There are at the moment no
semantic checks done, however for the `generate/` endpoint some defaulting is provided.

### Compartmentalize

In order to break up an ARN into its components (or: compartmentalize), use an
HTTP `GET` on `explode/$ARN`, for example:

```sh
$ curl -s https://arn.services/explode/arn:aws:s3:us-west-2::abucket | jq .
{
  "Partition": "aws",
  "Service": "s3",
  "Region": "us-west-2",
  "AccountID": "",
  "Resource": "abucket"
}
```

### Generate

To generate a fully qualified ARN from (certain) components, use an HTTP `POST` 
to `generate/`, for example:

```sh
$ curl -s -X POST \
       -H "Content-Type: application/json" \
       -d '{"Service":"s3", "Resource":"somebucket/someobject"}' \
       https://arn.services/generate
arn:aws:s3:us-west-2::somebucket/someobject
```

## Run it yourself

If you want to run ARN services yourself, you need to:

1. clone this repo (using `git clone https://github.com/mhausenblas/arn.services.git` for example)
1. create an S3 bucket for the Lambda functions and change the value of `ARNS_BUCKET` to your own S3 bucket in the [Makefile](https://github.com/mhausenblas/arn.services/blob/master/Makefile)
1. execute `make deploy`
1. OPTIONALLY: set up [API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
   and Route 53 for a custom domain and run `make up ARNS_ACM_CERT=$ARNS_ACM_CERT` with the ARN of the ACM cert (I did email validate for my Namecheap domain).
