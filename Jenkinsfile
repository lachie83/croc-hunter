node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter/"
  def pwd = pwd()
  def dockerEmail = "."
  def quay_creds_id = "quay_creds"

  stage ('preparation') {

  checkout scm

  sh "env"

  sh "mkdir -p ${workDir}"
  sh "cp -R ${pwd}/* ${workDir}"

  // read in required jenkins workflow config values
  def inputFile = readFile('Jenkinsfile.json')
  def config = new groovy.json.JsonSlurper().parseText(inputFile)
  println "workflow config ==> ${config}"

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

  withCredentials([[$class          : 'UsernamePasswordMultiBinding', credentialsId: quay_creds_id,
                      usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {

      sh "echo ${env.PASSWORD} | base64 --decode > ${pwd}/docker_pass"
      sh "docker login -e ${dockerEmail} -u ${env.USERNAME} -p `cat ${pwd}/docker_pass` quay.io"
      sh "cd ${pwd}"
      sh "make docker_build"
      sh "make docker_push"

      }
  }

  stage ('deploy') {

  def name = "croc-hunter"
  def replicas = "1"
  def cpu = "10m"
  def memory = "128Mi"

  // start kubectl proxy to enabled kube API access

  sh "kubectl proxy &"
  sh "kubectl --server=http://localhost:8001 get nodes"

  sh "/usr/local/linux-amd64/helm init"

  sh "/usr/local/linux-amd64/helm status croc-hunter || /usr/local/linux-amd64/helm install ${pwd}/charts/croc-hunter --name ${name} --set ImageTag=${env.BUILD_NUMBER},Replicas=${replicas},Cpu=${cpu},Memory=${memory} --namespace=${name}"

  sh "/usr/local/linux-amd64/helm upgrade croc-hunter ${pwd}/charts/croc-hunter --set ImageTag=${env.BUILD_NUMBER},Replicas=${replicas},Cpu=${cpu},Memory=${memory}"
  }
}
