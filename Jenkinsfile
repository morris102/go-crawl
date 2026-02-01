pipeline {
    agent any

    environment {
        DOCKER_HUB_CREDENTIAL = credentials('docker-hub-credential')
        IMAGE_NAME = 'your-username/go-crawl'
        TAG = 'latest'
    }

    stages {
        stage('Prepare Environment') {
            steps {
                sh '''
                    #!/bin/bash
                    set -a
                    source .env
                    set +a
                    echo "DB_HOST=$DB_HOST"
                '''
            }
        }

        stage('Git Pull') {
            steps {
                checkout scm
                sh 'git checkout master'
                sh 'git pull origin master'
            }
        }

        stage('Docker Login') {
            steps {
                sh 'echo $DOCKER_HUB_CREDENTIAL_PSW | docker login -u $DOCKER_HUB_CREDENTIAL_USR --password-stdin'
                echo 'Docker login successful!'
            }
        }

        // stage('Build Docker Image') {
        //     steps {
        //         sh "docker build -t ${IMAGE_NAME}:${TAG} ."
        //     }
        // }

        // stage('Push to Docker Hub') {
        //     steps {
        //         sh "docker tag ${IMAGE_NAME}:${TAG} ${IMAGE_NAME}:${TAG}"
        //         sh "docker push ${IMAGE_NAME}:${TAG}"
        //     }
        // }
    }

    post {
        always {
            sh 'docker logout'
        }
    }
}
