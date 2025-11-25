module vanity-worker

go 1.22

toolchain go1.23.4

require (
	github.com/aws/aws-sdk-go-v2 v1.39.2
	github.com/aws/aws-sdk-go-v2/config v1.31.12
	github.com/aws/aws-sdk-go-v2/credentials v1.18.16
	github.com/aws/aws-sdk-go-v2/service/s3 v1.88.4
	github.com/gagliardetto/solana-go v1.14.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
)

require (
	filippo.io/edwards25519 v1.0.0-rc.1
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.1
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.9
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.9
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.9
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.9
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.1
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.0
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.9
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.9
	github.com/aws/aws-sdk-go-v2/service/sso v1.29.6
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.35.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.38.6
	github.com/aws/smithy-go v1.23.0
	github.com/blendle/zapdriver v1.3.1
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.9.0
	github.com/gagliardetto/binary v0.8.0
	github.com/gagliardetto/treeout v0.1.4
	github.com/json-iterator/go v1.1.12
	github.com/klauspost/compress v1.13.6
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/mattn/go-colorable v0.1.4
	github.com/mattn/go-isatty v0.0.11
	github.com/mitchellh/go-testing-interface v1.14.1
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.2
	github.com/mostynb/zstdpool-freelist v0.0.0-20201229113212-927304c0c3b1
	github.com/mr-tron/base58 v1.2.0
	github.com/streamingfast/logging v0.0.0-20230608130331-f22c91403091
	go.mongodb.org/mongo-driver v1.12.2
	go.uber.org/atomic v1.7.0
	go.uber.org/multierr v1.6.0
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
)
