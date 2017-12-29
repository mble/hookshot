# Hookshot [![Build Status](https://travis-ci.org/mble/hookshot.svg?branch=master)](https://travis-ci.org/mble/hookshot) [![Go Report Card](https://goreportcard.com/badge/github.com/mble/hookshot)](https://goreportcard.com/report/github.com/mble/hookshot)

Hookshot is a *WIP* webhook handler for deploying Docker images.

Currently this is configured to authenticate against GitHub's `X-Hub-Signature` and to pull and deploy public images.

You'll need Go and Docker installed and running. `make docker` will fetch dependencies (with `dep`), compile the program
and build the docker image ready to be deployed.

## Usage

:warning: This is definitely not ready for production use.

You can run the container with the following after building with `make docker`:

```
docker run -d -p 2015:2015 -e PORT=2015 -e HUB_SECRET=secret -v /var/run/docker.sock:/var/run/docker.sock hookshot:latest
```

Unauthenticated requests return 404s:

```
$ curl -I localhost:2015
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Thu, 28 Dec 2017 20:09:22 GMT
Content-Length: 19
```

Successful authenticated requests return 200s:

```
curl -H "X-Hub-Signature: sha1=25af6174a0fcecc4d346680a72b7ce644b9a88e8" localhost:2015
{"version":"0.1.0","build":"666652c"}
```

## TODO

* Genericise authentication method, `X-Hookshot-Signature`?
* Allow for logging into private repositories and deploying private images
* Handle possible errors when starting containers
* Tests
