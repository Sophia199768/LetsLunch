# foodate

## Проект:
Веб-приложение для поиска пары для совместного посещения ресторанов. Пользователи могут создавать встречи, указывая ресторан, дату и время, а другие - присоединяться к ним.

Проблематика:
Более 26% российских первокурсников испытывают трудности общении со сверстниками, а учитывая огромную загруженность в университете с такой проблемой сталкивается любой студент

## Инструменты:
1) Backend: GO, Docker
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

## Приложение:
1. Окно регистрации
<img width="865" height="620" alt="Регистрация" src="https://github.com/user-attachments/assets/80cb87df-37ed-41b4-81a9-ca436bb62e18" />

2. Вход

<img width="518" height="604" alt="Вход" src="https://github.com/user-attachments/assets/413f1bf6-a77b-4f5b-96b8-eca77d508f39" />

4. Главная страница, на которой есть заведения которые можно выбрать для встречи

<img width="874" height="354" alt="Главная страница" src="https://github.com/user-attachments/assets/58772f45-b82b-4976-b5df-dcd2a1701d0e" />

5. Окно для создания встречи

<img width="841" height="467" alt="Создание встречи" src="https://github.com/user-attachments/assets/a165e8c4-ec95-4044-9bca-4ab353314394" />

   

