.PHONY: blockchain

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
# Currently, it generates smart contract for specific contract name.
# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

ROOT_PATH = $(shell pwd)
CONTRACT_PATH = ./contract
SCRIPTS_PATH = ./scripts

# Generate empty contract
init-contract:
	@echo "Initialize smart contract..."
	@mkdir contract
	@cd $(CONTRACT_PATH) && truffle init


# Generate smart contract with arg[1] name
add-contract:
	@echo "Generating smart contract..."
	@cd $(CONTRACT_PATH) && truffle create contract $(filter-out $@,$(MAKECMDGOALS))

# Compile smart contract, generate json file
compile-contract:
	@echo "Compiling smart contract..."
	@cd $(CONTRACT_PATH) && truffle compile

# Update golang package
update-package:
	@echo "Update golang package..."
	@cd $(SCRIPTS_PATH) && node extractabi.js
	@cd $(ROOT_PATH)
	@cd $(CONTRACT_PATH) && abigen --abi=build/contracts/BiddingWarsABI.json --pkg=biddingwars --out=../pkg/biddingwars/biddingwars.go
	@echo "Update golang package successfully!"
	@cd $(CONTRACT_PATH) && rm -f build/contracts/BiddingWarsABI.json
	@echo "Remove ABI json file successfully!"

# Test smart contract
test-contract:
	@echo "Testing smart contract..."
	@cd $(CONTRACT_PATH) && truffle test

# Deploy smart contract
deploy-contract:
	@echo "Deploying smart contract..."
	@cd $(CONTRACT_PATH) && truffle migrate
