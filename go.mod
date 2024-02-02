module github.com/giaoduc/giaoduc_api

go 1.16

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20190108154635-47c0da630f72

require (
	cloud.google.com/go v0.38.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/aws/aws-sdk-go v1.37.20
	github.com/bwmarrin/snowflake v0.3.0
	github.com/cenkalti/backoff v2.1.1+incompatible
	github.com/certifi/gocertifi v0.0.0-20190506164543-d2eda7129713 // indirect
	github.com/confluentinc/confluent-kafka-go v1.7.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gammazero/workerpool v0.0.0-20200311205957-7b00833861c6
	github.com/getsentry/raven-go v0.2.0
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.7.2
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-resty/resty/v2 v2.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang-migrate/migrate/v4 v4.3.1
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.1
	github.com/h2non/filetype v1.0.8
	github.com/imdario/mergo v0.3.8
	github.com/jinzhu/gorm v1.9.8
	github.com/jordic/lti v0.0.0-20160211051708-2c756eacbab9
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.1.1
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/minio/minio-go/v6 v6.0.31
	github.com/mitchellh/mapstructure v1.1.2
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/slack-go/slack v0.10.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/ugorji/go v1.2.6 // indirect
	github.com/xkeyideal/captcha v0.0.0-20171218103243-508b0d532b36
	github.com/xuri/excelize/v2 v2.4.1
	go.opencensus.io v0.22.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.1.4 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.18.3
)
