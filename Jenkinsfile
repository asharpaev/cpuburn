node {
    label 'slave'
    stage('Checkout'){
        checkout scm
    }
    stage('Check_Agent_Dependencies'){
        sh "if [ ! -f docker/docker ]; then curl https://download.docker.com/linux/static/stable/x86_64/docker-18.06.1-ce.tgz | tar xvz docker/docker; fi"
    }
    stage('Build') {
        sh "~/bin/docker build --no-cache -t registry:5000/cpuburn:latest ."
    }

}