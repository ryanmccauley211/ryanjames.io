pipeline {
    agent { docker { image 'golang' } }
    environment {
        SMTP_EMAIL = credentials('smtp-email')
        SMTP_PASS = credentials('smtp-pass')
    }
    stages {
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go build'
            }
        }
        stage('Setup Environment') {
            steps {
                sh 'export email=$SMTP_EMAIL'
                sh 'export pass=$SMTP_PASS'
            }
        }
        stage('Run') {
            steps {
                echo 'Running ryanjames.io web server'
                sh 'go run *.go'
            }
        }
    }
}