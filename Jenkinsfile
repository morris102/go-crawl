pipeline {
    agent any

    stages {
        stage('Load .env') {
            steps {
                sh '''
                  set -a
                  source .env
                  set +a
                  echo "DB_HOST=$DB_HOST"
                '''
            }
        }
    }
}
