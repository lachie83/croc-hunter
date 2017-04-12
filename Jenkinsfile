#!/usr/bin/groovy

// load pipeline functions
// @Library('github.com/lachie83/jenkins-pipeline@master')

podTemplate(label: 'mypod', containers: [
    containerTemplate(name: 'jnlp', image: 'jenkinsci/jnlp-slave:2.62', args: '${computer.jnlpmac} ${computer.name}', workingDir: '/home/jenkins'),
    containerTemplate(name: 'docker', image: 'docker:1.13.1-dind',         privileged: true, command: 'cat', ttyEnabled: true),
    containerTemplate(name: 'golang', image: 'golang:1.7.5', command: 'cat', ttyEnabled: true),
    containerTemplate(name: 'helm', image: 'lachlanevenson/k8s-helm:v2.3.0', command: 'cat', ttyEnabled: true),
    containerTemplate(name: 'kubectl', image: 'lachlanevenson/k8s-kubectl:v1.4.8', command: 'cat', ttyEnabled: true)
]){

  def pipeline = new io.estrado.Pipeline()

  node ('mypod') {
    def goPath = "/go"
    def workDir = "${goPath}/src/github.com/lachie83/croc-hunter/"
    def pwd = pwd()
    def chart_dir = "${pwd}/charts/croc-hunter"


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

    // set additional git envvars for image tagging
    pipeline.gitEnvVars()

    // used to debug deployment setup
    env.DEBUG_DEPLOY = false

    // debugging helm deployments
    if (env.DEBUG_DEPLOY) {
      println "Runing helm tests"
      container('kubectl') {
      pipeline.kubectlTest()
      }
      container('helm') {
        pipeline.helmConfig()
      }
    }

    def acct = pipeline.getContainerRepoAcct(config)

    // tag image with version, and branch-commit_id
    def image_tags_map = pipeline.getContainerTags(config)

    // compile tag list
    def image_tags_list = pipeline.getMapValues(image_tags_map)

    stage ('test and compile') {

      container('golang') {
        sh "go test -v -race ./..."
        sh "make bootstrap build"
      }
    }

    stage ('test deployment') {

      // run go tests
      container('golang') {
      }

      container('helm') {

        // run helm chart linter
        pipeline.helmLint(chart_dir)

        // run dry-run helm chart installation
        pipeline.helmDeploy(
          dry_run       : true,
          name          : config.app.name,
          version_tag   : image_tags_list.get(0),
          chart_dir     : chart_dir,
          replicas      : config.app.replicas,
          cpu           : config.app.cpu,
          memory        : config.app.memory
        )

      }
    }

    stage ('publish container') {

      container('docker') {

        // perform docker login to quay as the docker-pipeline-plugin doesn't work with the next auth json format
        withCredentials([[$class          : 'UsernamePasswordMultiBinding', credentialsId: config.container_repo.jenkins_creds_id,
                        usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
          sh "docker login -e ${config.container_repo.dockeremail} -u ${env.USERNAME} -p ${env.PASSWORD} quay.io"
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

    }

    // deploy only the master branch
    if (env.BRANCH_NAME == 'master') {
      stage ('deploy to k8s') {

          container('helm') {

          // Deploy using Helm chart
          pipeline.helmDeploy(
            dry_run       : false,
            name          : config.app.name,
            version_tag   : image_tags_list.get(0),
            chart_dir     : chart_dir,
            replicas      : config.app.replicas,
            cpu           : config.app.cpu,
            memory        : config.app.memory
          )

        }
      }
    }
  }
}