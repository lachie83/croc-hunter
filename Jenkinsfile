node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter"


  stage ('preparation') {

  checkout scm

  mkdir -p "$(dirname ${WORKDIR})"
  cp -R "${HOME}/croc-hunter" "${WORKDIR}"

  }

  stage ('compile') {

  cd "${WORKDIR}" && make bootstrap build
  go env
  }

  stage ('lint') {

  $/usr/local/linux-amd64/helm lint "${HOME}/croc-hunter/charts/croc-hunter"

  }

}
