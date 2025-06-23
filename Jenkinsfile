pipeline {
    agent any
    tools {
        go 'go-1.24.4'
        'hudson.plugins.sonar.SonarRunnerInstallation' 'sonarqube-scanner-latest'
    }

    environment {
        APP_NAME='mallow-sale-api'
        RELEASE = "1.0.0"
        DOCKER_USER = "butternoei008"
        DOCKER_PASS = 'docker-hub-account'
        IMAGE_NAME = "${DOCKER_USER}" + "/" + "${APP_NAME}"
        IMAGE_TAG = "${RELEASE}-${BUILD_NUMBER}"

        GO111MODULE='on'
        SONAR_SCANNER_HOME = tool 'sonarqube-scanner-latest'
    }

    stages {
        stage('Checkout from SCM') {
            steps {
               git branch: 'main',
               credentialsId: 'github-credential',
               url: 'https://github.com/hifat/mallow-sale-api'
            }
        }
        stage('Unit Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('SonarQube Analysis') {
            steps {
                script {
                    withSonarQubeEnv(credentialsId: 'sonarqube-token') {
                        sh '''
                            ${SONAR_SCANNER_HOME}/bin/sonar-scanner \
                                -Dsonar.projectKey=mallow-sale-api \
                                -Dsonar.sources=. \
                                -Dsonar.exclusions=**/*_test.go,**/vendor/**,**/mock/**
                        '''
                    }
                    // waitForQualityGate abortPipeline: false
                }
            }
        }
        stage('Build & Push to registry') {
            steps {
                script {
                    docker.withRegistry('', DOCKER_PASS) {
                        docker_image = docker.build "${IMAGE_NAME}"
                    }

                    docker.withRegistry('', DOCKER_PASS) {
                        docker_image.push("${IMAGE_TAG}")
                        docker_image.push('latest')
                    }
                }
            }
        }
    }
}
