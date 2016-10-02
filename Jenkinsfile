node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter/"
  def pwd = pwd()
  def dockerEmail = "."
  def quay_creds_id = "quay_creds"

  stage ('preparation') {

  checkout scm

  sh "env | sort"

  sh "mkdir -p ${workDir}"
  sh "cp -R ${pwd}/* ${workDir}"

  // read in required jenkins workflow config values
  def inputFile = readFile('Jenkinsfile.json')
  def config = new groovy.json.JsonSlurperClassic().parseText(inputFile)
  println "pipeline config ==> ${config}"

  // continue only if pipeline enabled
  if (!config.pipeline.enabled) {
      println "pipeline disabled"
      return
  }

  // load pipeline library
  dir('lib/jenkins-pipeline') {
      git branch: config.pipeline.library.branch,
              url: 'https://github.com/lachie83/jenkins-pipeline.git'
  }

  // load quay library module
  def quay = load 'lib/jenkins-pipeline/quay.groovy'

  }

  stage ('compile') {

  sh "cd ${workDir}"
  sh "make bootstrap build"
  sh "go test -v -race ./..."

  }

  stage ('lint') {

  sh "/usr/local/linux-amd64/helm lint ${pwd}/charts/croc-hunter"

  }

  stage ('publish') {

      quay.login(quay_creds_id)
      sh "cd ${pwd}"
      sh "make docker_build"
      sh "make docker_push"

  }

  stage ('deploy') {

  //def name = "croc-hunter"
  //def replicas = "1"
  //def cpu = "10m"
  //def memory = "128Mi"

  // start kubectl proxy to enable kube API access

  sh "kubectl proxy &"
  sh "kubectl --server=http://localhost:8001 get nodes"

  sh "/usr/local/linux-amd64/helm init"

  sh "/usr/local/linux-amd64/helm status croc-hunter || /usr/local/linux-amd64/helm install ${pwd}/charts/croc-hunter --name config.app.name --set ImageTag=${env.BUILD_NUMBER},Replicas=config.app.replicas,Cpu=config.app.cpu,Memory=config.app.memory --namespace=config.app.name"

  sh "/usr/local/linux-amd64/helm upgrade croc-hunter ${pwd}/charts/croc-hunter --set ImageTag=${env.BUILD_NUMBER},Replicas=config.app.replicas,Cpu=config.app.cpu,Memory=config.app.memory"
  
  }
}
