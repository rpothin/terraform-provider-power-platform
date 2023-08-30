build:
	go mod tidy
	go build -o ./bin/

install:
	go mod tidy
	go install

userdocs:
	tfplugindocs generate --provider-name powerplatform --rendered-provider-name "Power Platform"

servedocs:
	tfplugindocs generate --provider-name powerplatform --rendered-provider-name "Power Platform"
	mkdocs serve

unittest:
	export TF_ACC=0
	go install
	go clean -testcache
	go test -v ./... -run "^TestUnit"

acctest: 
	export TF_ACC=1
	go install
	go clean -testcache
	go test -v ./... -run "^TestAcc"

test:
	export TF_ACC=1
	go install
	go clean -testcache
	go test -v ./...

deps:
	go mod tidy

mocks:
	mockgen -destination=/workspaces/terraform-provider-power-platform/internal/mocks/client_mocks_bapi.go -package=powerplatform_mocks github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/bapi ApiClientInterface
  