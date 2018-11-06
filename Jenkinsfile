node {
    label 'slave'
    stage('Checkout'){
        checkout scm
    }
    stage('Build') {

        sh "k8sversion=v1.8.11; curl -LO https://storage.googleapis.com/kubernetes-release/release/$k8sversion/bin/linux/amd64/kubectl"
        sh "chmod +x ./kubectl"
        sh "./kubectl get pods"
    }

}