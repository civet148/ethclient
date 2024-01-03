#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
GO_ETHEREUM=/tmp/go-ethereum
CONTRACTS_DIR=contracts
CONTRACTS_ABI=abis
CONTRACTS_OUT=${CONTRACTS_DIR}/${CONTRACTS_ABI}

# --- ERC721 contract arguments ---
ERC721_PKG=erc721
ERC721_OUT=${CONTRACTS_DIR}/${ERC721_PKG}

build:
	solc --base-path $(CONTRACTS_DIR) --include-path $(CONTRACTS_DIR)/@openzeppelin --bin --abi --overwrite -o ${CONTRACTS_OUT} ${CONTRACTS_DIR}/@openzeppelin/contracts/token/ERC721/ERC721.sol && \
    mkdir -p ${ERC721_OUT} && abigen --abi ${CONTRACTS_OUT}/ERC721.abi --bin ${CONTRACTS_OUT}/ERC721.bin --pkg ${ERC721_PKG} --out ${ERC721_OUT}/erc721.go

install: openzeppelin solc abigen

solc:
	wget -O solc https://github.com/ethereum/solidity/releases/download/v0.8.17/solc-static-linux && chmod +x solc && mv solc $(GOPATH)/bin

abigen:
	- rm -rf ${GO_ETHEREUM}
	git clone -b v1.13.5 https://github.com/ethereum/go-ethereum.git ${GO_ETHEREUM} && cd ${GO_ETHEREUM} && go install -mod=readonly ./cmd/abigen

# Install openzeppelin solidity contracts
openzeppelin:
	@echo "Importing openzeppelin contracts..."
	@mkdir -p $(CONTRACTS_DIR) && cd $(CONTRACTS_DIR) && npm install && mv node_modules/@openzeppelin . && rm -rf node_modules

clean:
	rm -rf $(CLEAN) $(BINS)
.PHONY: clean
