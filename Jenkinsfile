node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter/"
  def pwd = "/home/jenkins/workspace/croc-hunter/dev"

  stage ('preparation') {

  checkout scm

  sh "mkdir -p ${workDir}"
  sh "cp -R ${pwd}/* ${workDir}"

  }

  stage ('compile') {

  sh "cd ${workDir}
  sh "make bootstrap build"
  sh "go test -v -race ./..."

  }

  stage ('lint') {

  sh "/usr/local/linux-amd64/helm lint ${pwd}/charts/croc-hunter"

  }

}
