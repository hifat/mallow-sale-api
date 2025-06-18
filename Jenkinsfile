pipeline {
    agent any
    tools {
        go 'go-1.24.4'
        sonarScanner 'sonarqube-scanner-latest'  // Add this line
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
                            ${SCANNER_HOME}/bin/sonar-scanner \
                                -Dsonar.projectKey=mallow-sale-api \
                                -Dsonar.sources=. \
                                -Dsonar.host.url=http://192.168.1.11:9000 \
                                -Dsonar.go.coverage.reportPaths=coverage.out \
                                -Dsonar.go.tests.reportPaths=test-report.json \
                                -Dsonar.exclusions=**/*_test.go,**/vendor/**,**/mock/**
                        '''
                    }
                }
            }
        }
    }
}