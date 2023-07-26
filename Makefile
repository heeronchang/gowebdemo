# note: call scripts from /scripts

GO = go 
BUILD_FLAGS = -ldflags "-s -w"

BASE_OUTPUT_DIR = ./bin

######## appone beign ########
APP_ONE_BINARY_NAME = appone
APP_ONE_DIR = app_one

appone:
	mkdir -p $(BASE_OUTPUT_DIR)/$(APP_ONE_DIR)
	mkdir -p $(BASE_OUTPUT_DIR)/$(APP_ONE_DIR)/configs
	$(GO) build $(BUILD_FLAGS) -o $(BASE_OUTPUT_DIR)/$(APP_ONE_DIR)/$(APP_ONE_BINARY_NAME) ./cmd/app_one/main.go
	cp ./configs/appone/config.yaml $(BASE_OUTPUT_DIR)/$(APP_ONE_DIR)/configs

runappone: 
	$(BASE_OUTPUT_DIR)/$(APP_ONE_DIR)/appone

cleanappone:
	rm -rf $(BASE_OUTPUT_DIR)/$(APP_ONE_DIR)
	
######## appone end ########

clean:
	rm -rf $(BASE_OUTPUT_DIR)/*