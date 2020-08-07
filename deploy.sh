#!/bin/bash
# Deployment script
# There are three parameters
# $1: your project dir name
# $2: the compiled executable file name
# $3: the environment your app run on, Mainly used to load configuration files according to the environment

# usage: sh /path_to_gaea/deploy.sh gaea gaea $env


#Configure according to actual situation
supervisorConfDir="/etc/supervisor"
appDir="/home/www"

projectname=$1
servicename=$2
environment=$3

if [ x"$projectname" = x ]; then
    echo "projectname param err"
    exit 1
fi

if [ x"$servicename" = x ]; then
    echo "service param err"
    exit 1
fi
if [ x"$environment" = x ]; then
    echo "not found env param"
else
    env_file="${appDir}/$1/conf/conf_${environment}.ini"
    if [ ! -f "$env_file" ]; then
        echo "$env_file config not found"
        exit 1
    fi
    conf_file="${appDir}/$1/conf/conf.ini"
    cp -f ${env_file} ${conf_file}
fi


function runserver(){
    project=$1
    service=$2
    useSignal=$3
    supervisorini="${supervisorConfDir}/${service}.ini"
    projectini="${appDir}/${project}/conf/${service}.ini"

    if [ ! -f "$projectini" ]; then
        echo "$projectini config not found"
        return
    fi

    if [ ! -f "$supervisorini" ]; then
        cp -f ${projectini} ${supervisorini}
        supervisorctl update
        return
    else
        checksum=`md5sum "${supervisorini}" | cut -d " " -f1`
        checksum1=`md5sum "${projectini}" | cut -d " " -f1`
        echo ${checksum}
        echo ${checksum1}
        if [ "$checksum" = "$checksum1" ]; then
            echo "ini not change"
        else
            echo "ini change"
            cp -f ${projectini} ${supervisorini}
            supervisorctl update

            #restart server when conf.ini is changed
            supervisorctl restart ${service}
            return
        fi
    fi

    serviceStatus=`supervisorctl  status ${service}  | awk  '{print $2}'`

    if [ "$useSignal" == "0" ]; then
        supervisorctl restart ${service}
    else
        if [ "$serviceStatus" !=  "RUNNING"  ]; then
            supervisorctl restart ${service}
        else
            supervisorctl  signal SIGUSR2 ${service}
        fi
    fi

}

runserver ${projectname} ${servicename} 1