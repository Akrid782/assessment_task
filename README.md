# Анализ процессов системы

## Требования

Необходимо написать программу, которая берёт ID процессов в системе и генерирует JSON файл следующей структуры

```json
[
    {
        "Process Name": "foo",
        "Process ID": 41,
        "Is Prime": true
    },
    {
        "Process Name": "bar",
        "Process ID": 42,
        "Is Prime": false
    }
]
```

Is Prime - результат проверки ID процесса как [натурального числа на простоту](https://ru.wikipedia.org/wiki/%D0%9F%D1%80%D0%BE%D1%81%D1%82%D0%BE%D0%B5_%D1%87%D0%B8%D1%81%D0%BB%D0%BE)

Process Name - название процесса в системе

Process ID - ID процесса в системе

В программе должна быть предусмотрена возможность через переменную окружения (ENV) задать с помощью регулярного
выражения (RegExp)  задать фильтр имени процессов, по которым будет выдан результат проверки “на простоту”.

В программе должен быть предусмотрен режим проверки уже сгенерированного файла “на ошибки”,
например есть такой файл
```json
[
    {
        "Process Name": "winamp.exe",
        "Process ID": 32,
        "Is Prime": true
    },
    {
        "Process Name": "antivirus_popova.bin",
        "Process ID": 64,
        "Is Prime": true
    }
]
```

В таком случае, ваша программа, в этом “дополнительном режиме”
(техническая реализация переключения режима - произвольная) должна будет
вывести следующее сообщение (формат произвольный)

```
Для winamp.exe процесса найдена ошибка.
Число(PID) 32 на самом деле составное, а в файле написано простое.

Для antivirus_popova.bin процесса найдена ошибка.
Число(PID) 64 на самом деле составное, а в файле написано простое.

В файле найдены ошибки: 2 штуки
```

## Инструкция

### Сборка и запуск проекта

Клонируем репозиторий для запуска проекта

```bash
git clone git@github.com:Akrid782/assessment_task.git
```

После того как мы склонировали репозиторий, заходим в директорию проекта и запускаем
первую команду для сборки проекта

```bash
docker-compose build
```

После этого можно запустить контейнер и посмотреть его логи.

Запуск:
```bash
docker-compose up -d
```

Просмотр логов:
```bash
docker container logs assessment_task
```

### Конфигурация

По умолчанию приложение анализирует процессы системы и генерирует файл
processes.json в соотвествии требованием.

```json
[
    {
        "Process Name": "gsd-a11y-settings",
        "Process ID": 8675,
        "Is Prime": false
    },
    {
        "Process Name": "gsd-color",
        "Process ID": 8676,
        "Is Prime": false
    },
    {
        "Process Name": "gsd-datetime",
        "Process ID": 8677,
        "Is Prime": true
    }
]
```

Чтобы изменить работу приложения, необходимо зайти в docker-compose.yml и изменить параметр
environment. Что предоставляет environment:
1. PROCESS_REGEXP_FILTER - Регулярка для фильтра процессов системы по их имени
2. MOD - Режимы работы SYSTEM_ANALYSIS или FILE_ANALYSIS. Где SYSTEM_ANALYSIS это режим анализа процессов системы,
а FILE_ANALYSIS это анализ загруженнего файла
3. PATH_FILE_ANALYSIS - путь до файла куда он был загружен в систему. Чтобы загрузить файл достаточно прописать volumes.

Пример:
```yml
      volumes:
          ...
          - {путь до файла в вашей системе}:/usr/local/src/assets/upload/example.json
      environment:
          PATH_FILE_ANALYSIS: /assets/upload/example.json
```
