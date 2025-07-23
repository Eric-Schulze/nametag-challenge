build-remote-world-pop:
	@mkdir -p ./world-pop/temp
	@ORIGINAL_VERSION=$$(world-pop -v | cut -d' ' -f3) ; \
	go build -C world-pop -ldflags "-X main.version=$$ORIGINAL_VERSION" -o ./temp/
	@curl  --request POST \
		-H "Content-Type: application/octet-stream" \
		--data-binary @./world-pop/temp/world-pop \
		http://localhost:4040/updater/upload
	@rm -rf ./world-pop/temp/

install-local-world-pop:
	@ORIGINAL_VERSION=$$(world-pop -v | cut -d' ' -f3) ; \
	go install -C world-pop -ldflags "-X main.version=$$ORIGINAL_VERSION" 

reset-versions:
	@go install -C world-pop -ldflags "-X main.version=1.0.0" 
	@mkdir -p ./world-pop/temp
	@ORIGINAL_VERSION=$$(world-pop -v | cut -d' ' -f3) ; \
	go build -C world-pop -ldflags "-X main.version=$$ORIGINAL_VERSION" -o ./temp/
	@curl  --request POST \
		-H "Content-Type: application/octet-stream" \
		--data-binary @./world-pop/temp/world-pop \
		http://localhost:4040/updater/upload
	@rm -rf ./world-pop/temp/


increase-remote-version:
	@mkdir -p ./world-pop/temp
	@ORIGINAL_VERSION=$$(world-pop -v | cut -d' ' -f3) ; \
	ORIGINAL_PATCH_VERSION=$$(echo "$$ORIGINAL_VERSION" | cut -d'.' -f3) ; \
	NEW_PATCH_VERSION=$$(($$ORIGINAL_PATCH_VERSION+1)) ; \
	NEW_VERSION=$$(echo "$$ORIGINAL_VERSION" | cut -d'.' -f1-2)".$$NEW_PATCH_VERSION" ; \
	go build -C world-pop -ldflags "-X main.version=$$NEW_VERSION" -o ./temp/
	@curl  --request POST \
		-H "Content-Type: application/octet-stream" \
		--data-binary @./world-pop/temp/world-pop \
		http://localhost:4040/updater/upload
	@rm -rf ./world-pop/temp/

toggle-logging:
	@ORIGINAL_VALUE=$$(cat ./world-pop/internal/init/config.yaml | grep enable_logging | cut -d':' -f2 | sed 's/[[:space:]]//g') ; \
	if [ "$$ORIGINAL_VALUE" == true ]; then \
		NEW_VALUE=false ; \
	else \
		NEW_VALUE=true ; \
	fi ; \
	sed -r -i.bak "s/enable_logging:( |\t)(true|false)/enable_logging: $$NEW_VALUE/g" ./world-pop/internal/init/config.yaml
	@rm -f ./world-pop/internal/init/config.yaml.bak
	@ORIGINAL_VERSION=$$(world-pop -v | cut -d' ' -f3) ; \
	go install -C world-pop -ldflags "-X main.version=$$ORIGINAL_VERSION" 

toggle-auto-update:
	@ORIGINAL_VALUE=$$(cat ./world-pop/internal/init/config.yaml | grep auto_update | cut -d':' -f2 | cut -d'#' -f1 | sed 's/[[:space:]]//g') ; \
	if [ "$$ORIGINAL_VALUE" == true ]; then \
		NEW_VALUE=false ; \
	else \
		NEW_VALUE=true ; \
	fi ; \
	echo "$$NEW_VALUE" ; \
	sed -r -i.bak "s/auto_update:( |\t)(true|false)/auto_update: $$NEW_VALUE/g" ./world-pop/internal/init/config.yaml
	@rm -f ./world-pop/internal/init/config.yaml.bak
	@ORIGINAL_VERSION=$$(world-pop -v | cut -d' ' -f3) ; \
	go install -C world-pop -ldflags "-X main.version=$$ORIGINAL_VERSION" 

