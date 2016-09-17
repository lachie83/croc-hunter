node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter/"
  def pwd = "/home/jenkins/workspace/croc-hunter/dev"
  def dockerEmail = "."
  def quay_creds_id = "quay_creds"

  stage ('preparation') {

  checkout scm

  sh "mkdir -p ${workDir}"
  sh "cp -R ${pwd}/* ${workDir}"

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
      sh "make docker_publish"

      }
  }

}
