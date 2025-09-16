pipeline {
    agent any
    tools {
        go 'go-1.25.0'
    }

    environment {
        APP_NAME = 'mallow-sale-api'
        RELEASE = '1.0.0'

        DOCKER_ACCOUNT = credentials('docker-hub-account')
        IMAGE_NAME = "${DOCKER_ACCOUNT_USR}/${APP_NAME}"
        IMAGE_TAG = "${RELEASE}-${BUILD_NUMBER}"

        JENKINS_HOST = 'host.docker.internal:8081'
        JENKINS_ACCOUNT = credentials('jenkins-account')
        CD_TRIGGER_TOKEN = credentials('cd-mls-api-trigger-token')

        // Go environment
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

        stage('Quality Checks') {
            parallel {
                stage('Unit Tests') {
                    steps {
                        sh '''
                            go test ./...
                        '''
                    }
                }
            }
        }

        stage('Build & Push Docker Image') {
            steps {
                script {
                    sh """
                        docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
                        docker tag ${IMAGE_NAME}:${IMAGE_TAG} ${IMAGE_NAME}:latest
                    """

                    withDockerRegistry(credentialsId: DOCKER_ACCOUNT, url: '') {
                        sh """
                            docker push ${IMAGE_NAME}:${IMAGE_TAG}
                            docker push ${IMAGE_NAME}:latest
                        """
                    }
                }
            }
        }

        stage('Trigger CD Pipeline') {
            steps {
                script {
                    echo 'Triggering CD pipeline...'

                    curlResponse = sh(
                        script: """
                            curl -w "%{http_code}" -o /tmp/cd_response.txt -s \
                                --user '${JENKINS_ACCOUNT_USR}:${JENKINS_ACCOUNT_PSW}' \
                                -X POST \
                                -H 'Cache-Control: no-cache' \
                                -H 'Content-Type: application/x-www-form-urlencoded' \
                                --data 'IMAGE_TAG=${IMAGE_TAG}' \
                                '${JENKINS_HOST}/job/mls-api-cd/buildWithParameters?token=${CD_TRIGGER_TOKEN}'
                        """,
                        returnStdout: true
                    ).trim()

                    if (curlResponse != '201' && curlResponse != '200') {
                        error "CD pipeline trigger failed with HTTP code: ${curlResponse}"
                    } else {
                        echo '✅ CD pipeline triggered successfully'
                    }
                }
            }
        }
    }

    post {
        always {
            script {
                sh """
                    docker rmi ${IMAGE_NAME}:${IMAGE_TAG} 2>/dev/null || echo 'Image ${IMAGE_NAME}:${IMAGE_TAG} not found'
                    docker rmi ${IMAGE_NAME}:latest 2>/dev/null || echo 'Image ${IMAGE_NAME}:latest not found'

                    # Clean up unused Docker resources
                    docker system prune -f 2>/dev/null || true
                """
            }
        }

        success {
            echo '✅ Pipeline completed successfully!'
        // Add notification here if needed
        }

        failure {
            echo '❌ Pipeline failed!'
        // Add notification here (Slack, email, etc.)
        }

        unstable {
            echo '⚠️ Pipeline completed with warnings'
        }
    }
}
