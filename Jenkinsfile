pipeline {
    agent any
    tools {
        go 'go-1.24.4'
    }

    environment {
        GO111MODULE='on'
    }

    stages {
        stage('Checkout from SCM') {
            steps {
               git branch: 'main',
               credentialsId: 'github-credential',
               url: 'https://github.com/hifat/mallow-sale-api'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('SonarQube Analysis') {
            steps {
                script {
                    withSonarQubeEnv(credentialsId: 'sonarqube-token') {
                        sh '''
                            sonar-scanner \
                                -Dsonar.projectKey=mallow-sale-api \
                                -Dsonar.sources=. \
                        '''
                    }
                }
            }
        }
    }
}