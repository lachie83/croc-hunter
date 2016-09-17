node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter"


  stage ('preparation') {

  checkout scm

  sh "mkdir -p $(dirname ${workDir})"
  sh "cp -R ${HOME}/croc-hunter ${workDir}"

  }

  stage ('compile') {

  sh "cd ${workDir} && make bootstrap build"
  sh "go env"
  }

  stage ('lint') {

  sh "/usr/local/linux-amd64/helm lint ${HOME}/croc-hunter/charts/croc-hunter"

  }

}
