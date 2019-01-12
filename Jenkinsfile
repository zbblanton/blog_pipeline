pipeline {
    agent any
<<<<<<< HEAD
=======
    triggers {
        pollSCM('*/1 * * * *')
    }
    environment {
        OPENFAAS_PASS = credentials('openfaasPass')
    }
>>>>>>> c195a96... Added openfaas cred as Jenkins cred
    stages {
        stage('Clone Blog Repo') {
            steps {
                echo 'Im sure it worked.'
            }
        }
        stage('Copy Static Files') {
            steps {
                echo 'Im sure it worked.'
            }
        }
        stage('Build New Function Image') {
            steps {
                echo 'Im sure it worked.'
            }
        }
        stage('Push New Image') {
            steps {
                echo 'Im sure it worked.'
            }
        }
        stage('Deploy New Function') {
            steps {
<<<<<<< HEAD
<<<<<<< HEAD
                echo 'Im sure it worked.'
=======
                sh '''
                    curl -X PUT \
                    http://192.168.0.10:31112/system/functions \
                    -H 'authorization: Basic $OPENFAAS_PASS' \
                    -H 'cache-control: no-cache' \
                    -H 'content-type: application/json' \
                    -H 'postman-token: bbd2fd48-307d-efbb-f4b3-9b73692cead7' \
                    -d '{
                    	    "service": "blog",
                    	    "image": "localhost:5000/bt-blog",
                    	    "labels": {
                    	        "com.openfaas.scale.zero": "true"
                    	    }
                	    }'
                '''
=======
                script {
                    reqJSON = '''
                    {
                        "service": "blog",
                        "image": "localhost:5000/bt-blog",
                        "labels": {
                            "com.openfaas.scale.zero": "true"
                        }
                    }
                    '''
                    resp = sh(returnStdout: true, script: "curl -o /dev/null -s -w '%{http_code}' -X PUT http://192.168.0.10:31112/system/functions -H 'authorization: Basic " + env.OPENFAAS_PASS + "' -d '" + reqJSON + "'").trim()
                    if (resp != "200" && resp != "201" && resp != "202") {
                        error("Updating function failed. Status code: " + resp)
                    }
                    echo resp
                }
>>>>>>> 9f41f91... Added curl check for function update
            }
        }
        stage('Clean Workspacee') {
            steps {
                cleanWs()
>>>>>>> c195a96... Added openfaas cred as Jenkins cred
            }
        }
    }
}
