pipeline {
    agent any
    tools {
        go 'go-1.25.0'
    }

    environment {
        APP_NAME = 'mallow-sale-api'
        RELEASE = '1.0.0'
        DOCKER_USER = 'butternoei008'
        DOCKER_ACCOUNT = 'docker-hub-account'
        IMAGE_NAME = "${DOCKER_USER}/${APP_NAME}"
        IMAGE_TAG = "${RELEASE}-${BUILD_NUMBER}"
        JENKINS_HOST = 'host.docker.internal:8081'
        JENKINS_API_TOKEN = credentials('jenkins-api-token')

        GO111MODULE = 'on'
    }

    stages {
        stage('Clean up workspace') {
            steps {
                cleanWs()
            }
        }

        stage('Checkout from SCM') {
            steps {
                git branch: 'main',
                credentialsId: 'github-credentials',
                url: 'https://github.com/hifat/mallow-sale-api'
            }
        }

        stage('Unit Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build & Push to registry') {
            steps {
                script {
                    withDockerRegistry(credentialsId: DOCKER_ACCOUNT, url: '') {
                        dockerImage = docker.build("${IMAGE_NAME}")
                        dockerImage.push("${IMAGE_TAG}")
                        dockerImage.push('latest')
                    }
                }
            }
        }

        stage('Cleanup Artifacts') {
            steps {
                script {
                    sh "docker rmi ${IMAGE_NAME}:${IMAGE_TAG}"
                    sh "docker rmi ${IMAGE_NAME}:latest"
                }
            }
        }

        stage('Trigger CD Pipeline') {
            steps {
                script {
                    sh """
                        curl -v -k --user butter:${JENKINS_API_TOKEN} \
                            -X POST \
                            -H 'cache-control: no-cache' \
                            -H 'content-type: application/x-www-form-urlencoded' \
                            --data 'IMAGE_TAG=${IMAGE_TAG}' \
                            '${JENKINS_HOST}/job/mls-api-cd/buildWithParameters?token=jenkins-mls-token'
                    """
                }
            }
        }
    }
}
