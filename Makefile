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

build-contract: check setup-submodules
	cd sub_modules/feralfile-exhibition-smart-contract && \
	npm install && \
	truffle compile && \
	jq -r ".bytecode" build/contracts/FeralfileExhibitionV2.json > ./build/FeralfileExhibitionV2.bin && \
	jq -r ".abi" build/contracts/FeralfileExhibitionV2.json > ./build/FeralfileExhibitionV2.abi && \
	jq -r ".bytecode" build/contracts/FeralfileExhibitionV3.json > ./build/FeralfileExhibitionV3.bin && \
	jq -r ".abi" build/contracts/FeralfileExhibitionV3.json > ./build/FeralfileExhibitionV3.abi && \
	jq -r ".bytecode" build/contracts/FeralfileExhibitionV4.json > ./build/FeralfileExhibitionV4.bin && \
	jq -r ".abi" build/contracts/FeralfileExhibitionV4.json > ./build/FeralfileExhibitionV4.abi && \
	jq -r ".bytecode" build/contracts/FeralfileEnglishAuction.json > ./build/FeralfileEnglishAuction.bin && \
	jq -r ".abi" build/contracts/FeralfileEnglishAuction.json > ./build/FeralfileEnglishAuction.abi && \
	jq -r ".bytecode" build/contracts/FeralFileAirdropV1.json > ./build/FeralFileAirdropV1.bin && \
	jq -r ".abi" build/contracts/FeralFileAirdropV1.json > ./build/FeralFileAirdropV1.abi && \
	cd -

build: build-contract
	abigen --abi sub_modules/feralfile-exhibition-smart-contract/build/FeralfileExhibitionV2.abi --bin sub_modules/feralfile-exhibition-smart-contract/build/FeralfileExhibitionV2.bin --pkg feralfilev2 -type FeralfileExhibitionV2 --out contracts/feralfile-exhibition-v2/abi.go
	abigen --abi sub_modules/feralfile-exhibition-smart-contract/build/FeralfileExhibitionV3.abi --bin sub_modules/feralfile-exhibition-smart-contract/build/FeralfileExhibitionV3.bin --pkg feralfilev3 -type FeralfileExhibitionV3 --out contracts/feralfile-exhibition-v3/abi.go
	abigen --abi sub_modules/feralfile-exhibition-smart-contract/build/FeralfileExhibitionV4.abi --bin sub_modules/feralfile-exhibition-smart-contract/build/FeralfileExhibitionV4.bin --pkg feralfilev4 -type FeralfileExhibitionV4 --out contracts/feralfile-exhibition-v4/abi.go
	abigen --abi sub_modules/feralfile-exhibition-smart-contract/build/FeralfileEnglishAuction.abi --bin sub_modules/feralfile-exhibition-smart-contract/build/FeralfileEnglishAuction.bin --pkg english_auction -type FeralfileEnglishAuction --out contracts/feralfile-english-auction/abi.go
	abigen --abi sub_modules/feralfile-exhibition-smart-contract/build/FeralFileAirdropV1.abi --bin sub_modules/feralfile-exhibition-smart-contract/build/FeralFileAirdropV1.bin --pkg airdropv1 -type FeralFileAirdropV1 --out contracts/feralfile-airdrop-v1/abi.go
