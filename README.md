# Тестовое задание для Lamoda Tech
# Сервис для работы с товарами на складе

Для поднятия сервиса достаточно ввести команду:
```bash
  make up
```

Для запуска тестов:
```bash
  make test
```

Если не удалось поднять сервис, возможно стоит запустить команду под sudo.
Часть конфига описана в .env файле, часть - в config.yaml.
Реализованы только методы API, описанные в задании, так что используется маленькое допущение о том,
что некоторые таблицы изначально заполнены, так как методы для модификации этих таблиц просто не реализованы.

Postman collection: https://www.postman.com/plsmoment/workspace/lamoda-test/request/23506251-ad04c4f4-821f-44ee-95cb-269683017baf