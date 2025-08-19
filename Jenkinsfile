pipeline {
    agent any
    tools {
        go 'go-1.25.0'
    }

    environment {
        GO111MODULE='on'
    }

    stages {
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
    }
}
