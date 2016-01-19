# ss3

A readonly command line tool to search S3 buckets, written in go.

## Authentication

ss3 uses credentials stored in `~/.aws/credentials`

## Usage

```
$ ss3
=> List all buckets

$ ss3 <bucket>
=> List all objects in <bucket>

$ ss3 -match "string" <bucket>
=> Print all matches in the bucket
```

## Quick Commands + Examples

```
$ ss3 -match "production/postgres/" "database-backups"
=> get most recent production copy

$ ss3 "database-backups" | grep "production/postgres/" | sort | tail -1 | cut -d " " -f 1
=> same

$ ss3 -match "production/postgres/" "database-backups" | xargs -I {} aws s3 cp s3://database-backups/{} .
=> start downloading most recent prod copy, with shared credentials
```
