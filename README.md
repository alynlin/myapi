```bash
go mod init github.com/alynlin/myapi
cobra-cli init

export modeltmp=./modeltmp
openapi-generator-cli generate \
  -i ./designapi/user_openapi.yaml \
  -g go-gin-server \
  -o $modeltmp \
  --additional-properties=packageName=v1,apiPath=model/v1,enumClassPrefix=true,interfaceOnly=true \
  --git-repo-id myapi --git-user-id alynlin --git-host github.com
cp -R $modeltmp/model pkg
cp $modeltmp/api/openapi.yaml api/openapi.yaml
rm - rf $modeltmp
```