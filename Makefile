APP_NAME=KetQuaXoSo

APP_ROOT=$(PWD)
APP_SRC=$(APP_ROOT)/main.go
APP_OUT=$(APP_ROOT)/out

.PHONY: complete
complete: android windows linux

.PHONY: android
android: clean out
	fyne package --os android --name="KetQuaXoSo" --id=com.dopaemon.ketquaxoso --icon $(APP_ROOT)/assets/Icon.png --release && \
	mv KetQuaXoSo.apk $(APP_OUT)/KetQuaXoSo-android-universal.apk

.PHONY: ios
ios: clean out

.PHONY: windows
windows: clean windows-amd64

.PHONY: windows-amd64
windows-amd64: clean out
	CC=x86_64-w64-mingw32-gcc fyne package --os windows --icon $(APP_ROOT)/assets/Icon.png --name="KetQuaXoSo" --id=com.dopaemon.ketquaxoso --release && \
	mv KetQuaXoSo.exe $(APP_OUT)/KetQuaXoSo-windows-amd64.exe

.PHONY: linux
linux: clean linux-arm64 linux-amd64

.PHONY: linux-arm64
linux-arm64: clean out
	CC=/usr/bin/aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -v && \
	mv KetQuaXoSo $(APP_OUT)/KetQuaXoSo-linux-aarch64

.PHONY: linux-amd64
linux-amd64: clean out
	CC=/usr/bin/x86_64-linux-gnu-gcc fyne package --os linux --icon $(APP_ROOT)/assets/Icon.png --name="KetQuaXoSo" --id=com.dopaemon.ketquaxoso --release && \
	mv KetQuaXoSo.tar.xz $(APP_OUT)/KetQuaXoSo-linux-amd64.tar.xz && \
	CC=/usr/bin/x86_64-linux-gnu-gcc GOOS=linux GOARCH=amd64 go build -v && \
	mv KetQuaXoSo $(APP_OUT)/KetQuaXoSo-linux-amd64

.PHONY: out
out:
	mkdir -p $(APP_OUT)

.PHONY: clean
clean:
	rm -rf $(APP_OUT)
