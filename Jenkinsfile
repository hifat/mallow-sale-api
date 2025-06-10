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
        // stage('Build') {
        //     steps {
        //         sh 'go build -o app'
        //     }
        // }
        // stage('Push to Registry') {
        //     steps {
        //         script {
        //             docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
        //                 def app = docker.build("butter48/mallow-sale-api:${env.BUILD_NUMBER}")
        //                 app.push()
        //             }
        //         }
        //     }
        // }
    }
}