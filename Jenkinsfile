node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter"
  def pwd = env.PWD

  stage ('preparation') {

  checkout scm

  sh "mkdir -p ${workDir}"
  sh "cp -R ${pwd}/croc-hunter ${workDir}"

  }

  stage ('compile') {

  sh "cd ${workDir} && make bootstrap build"
  sh "go env"

  }

  stage ('lint') {

  sh "/usr/local/linux-amd64/helm lint ${pwd}/croc-hunter/charts/croc-hunter"

  }

}
