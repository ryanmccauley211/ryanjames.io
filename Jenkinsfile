pipeline {
    agent {
        docker {
            image 'golang'
            args '-e XDG_CACHE_HOME=/tmp/.cache'
        }
    }
    environment {
        REGISTRY = 'ryanmccauley211/ryanjames'
        SMTP_EMAIL = credentials('smtp-email')
        SMTP_PASS = credentials('smtp-pass')
    }
    stages {
        stage('Environment Setup') {
            steps {
                sh 'go version'
                sh 'export email=$SMTP_EMAIL'
                sh 'export pass=$SMTP_PASS'
            }
        }
        stage('Compile') {
            steps {
                echo 'Installing go dependencies'
                sh 'go get -d -v ./...'
                sh 'go build'
            }
        }
        stage('Build') {
            steps {
                script {
                    docker.build REGISTRY + ":$BUILD_NUMBER"
                }
            }
        }
    }
}