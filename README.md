# foodate


## Проект:
Веб-приложение для поиска пары для совместного посещения ресторанов. Пользователи могут создавать встречи, указывая ресторан, дату и время, а другие — присоединяться к ним.

## Инструменты:
1) Backend: GO
2) Frontend: JavaScript, HTML, CSS

## Инструкция по запуску:
Проще всего будет на линуксе
Очень попрошу использовать GoLand

1. скачать go https://go.dev/doc/install#
2. установить:
export PATH=$PATH:$HOME/go/bin
export GOPATH="$HOME/go"
скачать digen https://github.com/strider2038/digen

3. скачать goose https://github.com/pressly/goose

4. скачать Task: https://github.com/go-task/task/releases (инструкция: https://taskfile.dev/installation/)

5. скачать docker и docker engine(на сайте полноценный гайд по установке)

6. собрать docker-compose.yml при помощи docker-compose build.

если нет docker-compose- устанавливаем и его
docker-compose up запустит бд и админку("http://localhost:5050:80")
