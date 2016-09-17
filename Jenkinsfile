node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter"
  def home = env.BUILD_TAG


  stage ('preparation') {

  checkout scm

  sh "mkdir -p ${workDir}"
  sh "cp -R ${home}/croc-hunter ${workDir}"

  }

  stage ('compile') {

  sh "cd ${workDir} && make bootstrap build"
  sh "go env"

  }

  stage ('lint') {

  sh "/usr/local/linux-amd64/helm lint ${home}/croc-hunter/charts/croc-hunter"

  }

}
