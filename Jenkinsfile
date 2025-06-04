pipeline {
    agent any
    tools {
        go 'go-1.24'
    }

    environment {
        GO111MODULE='on'
    }

    stages {
        stage('Test') {
            when {
                branch 'main'
            }
            steps {
                git 'https://github.com/hifat/mallow-sale-api.git'
                sh 'go test ./...'
            }
        }
    }
}