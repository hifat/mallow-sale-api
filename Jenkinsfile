pipeline {
    agent any
    tools {
        go 'go-1.25.0'
    }

    environment {
        APP_NAME = 'mallow-sale-api'
        RELEASE = '1.0.0'
        DOCKER_USER = 'butternoei008'
        DOCKER_PASS = 'docker-hub-account'
        IMAGE_NAME = "${DOCKER_USER}/${APP_NAME}"
        IMAGE_TAG = "${RELEASE}-${BUILD_NUMBER}"

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
                    withDockerRegistry(credentialsId: DOCKER_PASS, url: '') {
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
    }
}
