# Demo Walkthrough

## Acknowledgements 
Continuation of the awesome work by everett-toews.
* https://gist.github.com/everett-toews/ed56adcfd525ce65b178d2e5a5eb06aa

## Watch Demo

https://www.youtube.com/watch?v=eMOzF_xAm7w

# Prerequisites
kubectl access to a Kubernetes 1.4+ cluster

# Install Helm (Mac OS)

```
brew install kubernetes-helm
helm init
helm repo update
```

## Fork repo
``` 
https://github.com/lachie83/croc-hunter#fork-destination-box
```

## Install Kube Lego chart
```
helm install stable/kube-lego --set config.LEGO_EMAIL=<valid-email>,config.LEGO_URL=https://acme-v01.api.letsencrypt.org/directory
```

## Install Nginx ingress chart
```
helm install stable/nginx-ingress

Follow the notes from helm status to determine the external IP of the nginx-ingress service
```

## Add a DNS entry with your provider and point it do the external IP
```
blah.test.com in A <nginx ingress svc external-IP>

or *.test.com in A <nginx ingress svc external-IP>

```


## Update jenkins.values.yaml
```
Find and replace `jenkins.acs.az.estrado.io` with the DNS name provisioned above

helm --namespace jenkins --name jenkins -f ./jenkins-values.yaml install stable/jenkins

watch kubectl get svc --namespace jenkins # wait for external ip
export JENKINS_IP=$(kubectl get svc jenkins-jenkins --namespace jenkins --template "{{ range (index .status.loadBalancer.ingress 0) }}{{.}}{{ end }}")
export JENKINS_URL=http://${JENKINS_IP}:8080

kubectl get pods --namespace jenkins # wait for running
open ${JENKINS_URL}/login

printf $(kubectl get secret --namespace jenkins jenkins-jenkins -o jsonpath="{.data.jenkins-admin-password}" | base64 --decode) | pbcopy
```

## Add credentials for private container registry (optional)
```
kubectl create secret docker-registry croc-hunter-secrets --docker-server=$DOCKER_SERVER --docker-username=$DOCKER_USERNAME --docker-password=$DOCKER_PASSWORD --docker-email=$DOCKER_EMAIL --namespace=croc-hunter
```
Reference to the secret name must also be added to the chart values.yaml or set on install.

## Login and configure Jenkins and setup pipeline
```
# username: admin
# password: <paste>

If you're not using quay you can configure this to alternate locations in Jenkinsfile.json
# Credentials > Jenkins > Global credentials > Add Credentials
#   Username: lachie83
#   Password: ***
#   ID: quay_creds
#   Description: https://quay.io/user/lachie83

# Open Blue Ocean
# Create a new Pipeline
# Where do you store your code?
#   GitHub
# Connect to Github
#   Create an access key here
#     Token description: kubernetes-jenkins
#   Generate token > Copy Token > Paste back in Jenkins  
# Which organization does the repository belong to?
#   lachie83
# Create a single Pipeline or discover all Pipelines?
#   New pipeline
# Choose a repository
#   croc-hunter
# Create Pipeline
```

## Watch Jenkins build agents run
```
kubectl get pods --namespace jenkins
```

## Update Org to build PRs
```
# Classic Jenkins
# lachie83 (GitHub org)
# Configure
# Advanced
#   Build origin PRs (merged with base branch)
# Save
```


## Setup Webhook in Github
``` 
printf ${JENKINS_URL}/github-webhook/ | pbcopy

# https://github.com/lachie83/croc-hunter/settings/hooks
# Add webhook
#   Payload URL: <paste>
# Which events would you like to trigger this webhook?
#   Send me everything.
# Add webhook
```

## Update croc-hunter ingress records
```
Update croc-hunter.acs.az.estrado.io in charts/croc-hunter/values.yaml

Configured DNS A record to point to the Nginx Ingress IP
Once master branch is pushed it should be available at that name
```


## Pushing Game update
```
git checkout dev
sed -i "" "s/game\.js/game2\.js/g" croc-hunter.go
git commit -am "Game 2"
git push
```

### Building and releasing
```
open ${JENKINS_URL}/blue/organizations/jenkins/lachie83%2Fcroc-hunter/activity/

# dev branch builds

open https://github.com/lachie83/croc-hunter

# PR from dev to master
# PR builds
# merge the PR
# master builds and deploys new version
```
