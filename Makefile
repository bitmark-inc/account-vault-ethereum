# SPDX-License-Identifier: ISC
# Copyright (c) 2019-2023 Bitmark Inc.
# Use of this source code is governed by an ISC
# license that can be found in the LICENSE file.

.PHONY:

truffle-check: ; @which truffle > /dev/null
abigen-check: ; @which abigen > /dev/null
jq-check: ; @which jq >/dev/null
npm-check: ;@which npm >/dev/null

check: truffle-check abigen-check jq-check npm-check

clean:
	rm -rf sub_modules/feralfile-exhibition-smart-contract/build
	rm -rf **/abi.go

setup-submodules:
	cp sub_modules/feralfile-exhibition-smart-contract/.secret.json.example sub_modules/feralfile-exhibition-smart-contract/.secret.json

check-build-contract: check
ifndef contract_name 
	$(error contract_name is undefined)
endif
ifndef output
	$(error output is undefined)
endif
ifndef pkg
	$(error pkg is undefined)
endif

build-contract: check-build-contract setup-submodules
	pushd sub_modules/feralfile-exhibition-smart-contract && \
	npm install && \
	truffle compile && \
	jq -r ".bytecode" build/contracts/$(contract_name).json > ./build/$(contract_name).bin && \
	jq -r ".abi" build/contracts/$(contract_name).json > ./build/$(contract_name).abi && \
	popd && \
	abigen --abi sub_modules/feralfile-exhibition-smart-contract/build/$(contract_name).abi --bin sub_modules/feralfile-exhibition-smart-contract/build/$(contract_name).bin --pkg $(pkg) -type $(contract_name) --out $(output)/abi.go  
