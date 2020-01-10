# arn.services

A utility service for Amazon Resource Names ([ARNs](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)), providing the following:

## Compartmentalize

In order to compartmentalize an ARN, that is, to decompose it into its parts, use a `GET` on `explode/$ARN`, for example:

```sh
$ curl arn.services/explode/arn:aws:s3:us-west-2::abucket
{
  "partition": "aws",
  "service": "s3",
  "region": "us-west-2",
  "resource-id": "abucket"
}
```

## Generate

If you have (certain) components of an ARN and want to generate a fully qualified ARN, use a `POST` to `generate/`, for example:

```sh
$ curl -X POST \
       -H "Content-Type: application/json"
       -d '{"service":"s3", "resource-id":"somebucket/someobject"}' \
       arn.services/generate
arn:aws:s3:us-west-2::somebucket/someobject
```
