node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter"


  stage ('preparation') {

  checkout scm

  mkdir -p "$(dirname ${workDir})"
  cp -R "${HOME}/croc-hunter" "${workDir}"

  }

  stage ('compile') {

  cd "${workDir}" && make bootstrap build
  go env
  }

  stage ('lint') {

  $/usr/local/linux-amd64/helm lint "${HOME}/croc-hunter/charts/croc-hunter"

  }

}
