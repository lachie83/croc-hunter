# croc-hunter

[![](https://images.microbadger.com/badges/image/lachlanevenson/croc-hunter.svg)](http://microbadger.com/images/lachlanevenson/croc-hunter "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/lachlanevenson/croc-hunter.svg)](http://microbadger.com/images/lachlanevenson/croc-hunter "Get your own version badge on microbadger.com")
[![](https://images.microbadger.com/badges/commit/lachlanevenson/croc-hunter.svg)](http://microbadger.com/images/lachlanevenson/croc-hunter "Get your own commit badge on microbadger.com")

[![CircleCI](https://circleci.com/gh/lachie83/croc-hunter.svg?style=svg)](https://circleci.com/gh/lachie83/croc-hunter)


# Deploy using Deis Workflow
## Dockerfile
Add Dockerfile to your repo
```
cd <repo-path>
deis create
git push deis master
```
## Docker images
```
deis create --no-remote
deis pull <repo>/<image-name> -a <app-name>
```
