# notes-app-kode

Команда для сборки и запуска `docker-compose up`

В этом проекте подготовлены 2 пользователя. Их пары user:password :
`"boris": "verysecure"`
`"john": "notverysecure"`

Шаблоны запросов:

1) Создание новой заметки
   
   URL: localhost:4000/create
   Method: POST Body:
  `{
      "user":"boris",
      "password":"verysecure",
      "note":"Новая заметка"
  }`
Вывод: Заметка создана

2) Создание новой заметки
   
   URL: localhost:4000/create
  Method: POST
  Body:
  `{
      "user":"boris",
      "password":"verysecure",
      "note":"Новая замета"
  }`
  Вывод: `[{"code":1,"pos":6,"row":0,"col":6,"len":6,"word":"замека","s":["заметка","замка","заминка"]}]`

3) Создание новой заметки
   
   URL: localhost:4000/create
  Method: POST
  Body: 
  `{
      "user":"boris",
      "password":"NOTverysecure",
      "note":"Новая заметка"
  }`
  Вывод: `Forbidden`

4) Показать все заметки пользователя
   
   URL: localhost:4000/show
Method: GET
Body:
`{
    "user":"boris",
    "password":"verysecure"
}`
Вывод: *Список заметок*

5) Показать все заметки пользователя
   
   URL: localhost:4000/show
Method: GET
Body:
`{
    "user":"NOTboris",
    "password":"verysecure"
}`
Вывод: Forbidden
