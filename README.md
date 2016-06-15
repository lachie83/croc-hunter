# croc-hunter

# Deploy using Deis Workflow
## Buildpacks
See Heroku documentation for repo requisites -- https://github.com/heroku/go-getting-started
```
cd <repo-path>
deis create
git push deis master
```
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
