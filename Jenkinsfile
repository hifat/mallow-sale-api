pipeline {
    agent any
    tools {
        go 'go-1.24'
    }

    environment {
        GO111MODULE='on'
        BRANCH_NAME='main'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: "*/${BRANCH_NAME}"]],
                    userRemoteConfigs: [[
                        url: 'https://github.com/hifat/mallow-sale-api.git'
                    ]]
                ])
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}