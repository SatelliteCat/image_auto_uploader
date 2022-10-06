# Автозагрузка изображений из буфера обмена ОС

Изображения, попавшие в буфер обмена операционной системы, автоматически загружаются на сайт [imgur](https://imgur.com/upload).

Буфер проверяется каждые полсекунды.

Для сборки для windows использовать команду: `go build -ldflags "-s -w -v -H=windowsgui"`.

## Скриншоты

- macOS, `Ctrl+Shift+Cmd+4`
- Linux/Ubuntu, `Ctrl+Shift+PrintScreen`
- Windows, `Shift+Win+s`
