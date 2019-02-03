#!/env/bin/bash
set -exuo pipefail

if [ -z "$APP_ENVIRONMENT" ]; then
    echo "Need to set APP_ENVIRONMENT"
    exit 1
fi 

cd ${APP_ENVIRONMENT}

for APP_SERVICE_CONF in $(find ./ -name "*.json") ; do
    curl -X PUT http://consul-agent:8500/v1/kv/app/${APP_SERVICE_CONF} --data-binary @${APP_SERVICE_CONF} -H "content-type:application/json" --silent --output /dev/null
done

exit 0