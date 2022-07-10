# toanother
Манипуляции с музыкальными стриминговыми сервисами.




# Навигация
[Сначала установите](#сначала-установите)

[Яндекс.Музыка](#яндексмузыка)

[Спотифай](#спотифай)

[Для разработчиков: зависимости](#для-разработчиков-зависимости)

[Для разработчиков: прочее](#для-разработчиков-прочее)




# Сначала установите

## Windows 10 и ниже
- [WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2)

## macOS (не знаю)

## Linux
- libgtk3 
- libwebkit
- Но это не точно




# Яндекс.Музыка
## Что можно
- Импортировать лайкнутые треки в Спотифай.
- Получить список лайкнутых плейлистов, артистов, альбомов.
- Посмотреть треки в плейлистах.

## Что нужно
- **Если вы работаете в Яндексе: попросить разработчиков сделать нормальное публичное API как в Спотифае.**
- [Сделать фонотеку публичной](https://music.yandex.ru/settings/other). Еще для каждого плейлиста есть своя настройка приватности, ну там думаю разберетесь.
- Не дрючить программу частым нажатием кнопки `Загрузить`, а то вылезет Яндексовская капча, и все сломается. Придется ждать полчаса-час.

## В программе
1. Перейти в раздел Яндекса.
2. Указать свой логин в настройках.
3. Применить настройки.
4. В нужном разделе нажать кнопку `Загрузить`.
5. Далее зависит от ваших нужд. Например, для импорта в Спотифай, надо перейти в раздел `Спотифай`. 




# Спотифай
## Что можно
- Импортировать лайкнутые треки из Яндекс.Музыки.

## Что нужно
1. Включить VPN если вы из России. Еще поменять страну аккаунта на страну VPN'а.
2. [Создать приложение](https://developer.spotify.com/dashboard/applications). Название и описание не имеют значения.
3. Тыкнуть на созданное приложение. Откроется меню вашего приложения.
4. Скопируйте куда-нибудь `Client ID` и `Client Secret`.
5. Нажмите на кнопку `EDIT SETTINGS` (справа вверху).
6. В поле `Redirect URIs` вставьте
`http://localhost:8080/spotify/callback` и потом нажмите кнопку `ADD`.
7. Пролистайте чуть вниз и нажмите кнопку `Save`.

## В программе
1. Перейти в раздел Спотифая
2. Указать `Client ID` и `Client Secret`.
3. Применить настройки.
4. Чтоб убедиться что все работает можете нажать кнопку `Пинг`.




# Для разработчиков: зависимости

## Linux
- libgtk3
- libwebkit

## macOS
```shell
xcode-select --install
```

## Windows 10 и ниже:
- [WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2)

## Все системы
- GCC
- Wails:
```shell
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```
- `cd` в этот проект
- запустить: `wails dev`. Для брейкпоинтов запускать через `Run and Debug` (VS Code)
- билд: `wails build`



# Для разработчиков: прочее
## Технологии
- Wails (связывает Go и фронт на JS). Используется бета.
- Vite, Svelte (TS), SASS на фронте.

## Обновить Wails
1. `wails update -pre`
2. Изменить версию в go.mod
3. `go mod tidy`