node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter"
  def workSpace = "/home/jenkins/workspace"

  stage ('preparation') {

  checkout scm

  sh "mkdir -p ${workDir}"
  sh "cp -R ${workSpace}/croc-hunter ${workDir}"

  }

  stage ('compile') {

  sh "make -f ${workDir}/Makefile bootstrap build"
  sh "go env"

  }

  stage ('lint') {

  sh "/usr/local/linux-amd64/helm lint ${workSpace}/croc-hunter/charts/croc-hunter"

  }

}
