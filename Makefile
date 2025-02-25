APP_NAME = Marisalt
APP_BUNDLE = $(APP_NAME).app
APP_EXECUTABLE = $(APP_NAME)
APP_IDENTIFIER = com.muddxyii.marisalt
ASSETS_DIR = assets
BUILD_DIR = builds

.PHONY: build-mac clean

build-mac:
	@echo "Building macOS application..."
	@mkdir -p $(BUILD_DIR)
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_EXECUTABLE)
	mkdir -p $(BUILD_DIR)/$(APP_BUNDLE)/Contents/{MacOS,Resources}
	mv $(BUILD_DIR)/$(APP_EXECUTABLE) $(BUILD_DIR)/$(APP_BUNDLE)/Contents/MacOS/
	cp -r $(ASSETS_DIR) $(BUILD_DIR)/$(APP_BUNDLE)/Contents/Resources
	echo '<?xml version="1.0" encoding="UTF-8"?>\
	<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">\
	<plist version="1.0">\
	<dict>\
		<key>CFBundleExecutable</key>\
		<string>$(APP_NAME)</string>\
		<key>CFBundleIdentifier</key>\
		<string>$(APP_IDENTIFIER)</string>\
		<key>CFBundleName</key>\
		<string>$(APP_NAME)</string>\
		<key>CFBundlePackageType</key>\
		<string>APPL</string>\
		<key>NSHighResolutionCapable</key>\
		<true/>\
	</dict>\
	</plist>' > $(BUILD_DIR)/$(APP_BUNDLE)/Contents/Info.plist
	@echo "Done! Application bundle created at ./$(BUILD_DIR)/$(APP_BUNDLE)"

clean:
	rm -rf $(BUILD_DIR)
