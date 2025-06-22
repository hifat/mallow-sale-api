pipeline {
    agent any
    tools {
        go 'go-1.24.4'
        'hudson.plugins.sonar.SonarRunnerInstallation' 'sonarqube-scanner-latest'
    }

    environment {
        APP_NAME='mallow-sale-api'
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
                            ${SONAR_SCANNER_HOME}/bin/sonar-scanner \
                                -Dsonar.projectKey=mallow-sale-api \
                                -Dsonar.sources=. \
                                -Dsonar.exclusions=**/*_test.go,**/vendor/**,**/mock/**
                        '''
                    }
                }
            }
        }
    }
}
