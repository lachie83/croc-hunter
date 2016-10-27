node {
  def goPath = "/go"
  def workDir = "${goPath}/src/github.com/lachie83/croc-hunter/"
  def pwd = pwd()
  def chart_dir = "${pwd}/charts/croc-hunter"
  def dockerEmail = "."
  def quay_creds_id = "quay_creds"

  checkout scm

  // read in required jenkins workflow config values
  def inputFile = readFile('Jenkinsfile.json')
  def config = new groovy.json.JsonSlurperClassic().parseText(inputFile)
  println "pipeline config ==> ${config}"

  // continue only if pipeline enabled
  if (!config.pipeline.enabled) {
      println "pipeline disabled"
      return
  }

  // load pipeline class
  def pipeline = new io.estrado.Pipeline()

  // set additional git envvars for image tagging
  pipeline.gitEnvVars()

  // tag image with version, and branch-commit_id
  def acct = pipeline.getContainerRepoAcct(config)
  def image_tags_map = pipeline.getContainerTags(config)
  def image_tags_list = pipeline.getMapValues(image_tags_map)

  stage ('preparation') {

    sh "env | sort"

    sh "mkdir -p ${workDir}"
    sh "cp -R ${pwd}/* ${workDir}"

  }

  stage ('compile') {

    sh "cd ${workDir}"
    sh "make bootstrap build"
    sh "go test -v -race ./..."

  }

  stage ('lint') {

    pipeline.helmLint(chart_dir)

  }

  stage ('publish') {

    // perform docker login to quay as the docker-pipeline-plugin doesn't work with the next auth json format
    withCredentials([[$class          : 'UsernamePasswordMultiBinding', credentialsId: quay_creds_id,
                    usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
      sh "docker login -e ${dockerEmail} -u ${env.USERNAME} -p ${env.PASSWORD} quay.io"
    }

    // build and publish container
    pipeline.containerBuildPub(
        dockerfile: config.container_repo.dockerfile,
        host      : config.container_repo.host,
        acct      : acct,
        repo      : config.container_repo.repo,
        tags      : image_tags_list,
        auth_id   : config.container_repo.jenkins_creds_id
    )

  }

  // deploy only the master branch
  if (env.BRANCH_NAME == 'master') {
    stage ('deploy') {

      // start kubectl proxy to enable kube API access
      pipeline.kubectlProxy()

      // Deploy using Helm chart
      pipeline.helmDeploy(
        name          : config.app.name,
        build_number  : image_tags_list.get(0),
        chart_dir     : chart_dir,
        replicas      : config.app.replicas,
        cpu           : config.app.cpu,
        memory        : config.app.memory
      )

    }
  }
}
