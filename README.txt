https://cloud.mail.ru/public/Gqjn/xEBuEQz82
План разработки приложения file-watcher на Go:

	I. Базовая функциональность:

		1. Выбор библиотеки для отслеживания файлов:

		Изучить возможности стандартной библиотеки os/fsnotify:

			•Преимущества: встроенная, простая в использовании.

			•Недостатки: может быть менее эффективна для большого количества файлов.

		Рассмотреть сторонние библиотеки, например, fsnotify/fsnotify или go-watcher/watcher:

			•Преимущества: оптимизированы для производительности, могут иметь дополнительные функции.

			•Недостатки: требуется установка, могут иметь зависимости.

		2. Создание основной логики приложения:

		Реализовать функцию watchDir(dirPath string), которая:

			•Принимает путь к директории для отслеживания.

			•Инициализирует наблюдатель за файловой системой.

			•Запускает бесконечный цикл, который ожидает событий от наблюдателя.

			•Обрабатывает события:

		При изменении файла:

			•Выводит в консоль информацию об изменении (имя файла, тип события).

			•Вызывает функцию handleFileChange(filePath string).

		3. Реализация функции обработки изменений:

		Создать функцию handleFileChange(filePath string), которая:

			Принимает путь к измененному файлу.

			Выполняет заданные действия:

			•Например, копирует файл в другую директорию.

			•Запускает скрипт.

			•Отправляет уведомление.

		4. Запуск приложения из командной строки:

		Использовать пакет flag для обработки аргументов командной строки:

			•Путь к отслеживаемой директории.

			•Опциональные параметры, например, интервал проверки изменений.

		Запустить приложение с тестовой директорией.

	II. Запуск как сервиса (демона):

		1. Выбор способа запуска:

			•Использовать сторонний менеджер процессов, например, systemd (Linux) или launchd (macOS).

			Реализовать собственную логику демонизации:

			•Создать дочерний процесс.

			•Отключить терминал и перенаправить потоки ввода/вывода.

		2. Настройка логирования:

			•Использовать пакет log для записи информации о работе приложения в файл.

			•Добавить возможность настройки уровня логирования.

	III. GUI клиент:

		1. Выбор библиотеки для создания GUI:

			•Изучить возможности fyne, walk, qt, Electron и выбрать подходящую.

		2. Разработка интерфейса:

			•Окно для выбора директории для отслеживания.

			•Список отслеживаемых файлов с информацией об изменениях.

			•Кнопки для запуска/остановки отслеживания.

			•Возможность настройки действий при изменении файла.

		3. Взаимодействие с сервисом:

			Выбрать способ коммуникации:

			•TCP сокеты.

			•Unix domain socket.

			•Файлы конфигурации.

			Реализовать обмен сообщениями между клиентом и сервисом.

	IV. Дополнительные возможности:

		•Обработка ошибок: добавить обработку ошибок на всех этапах работы приложения.

		•Тестирование: написать unit-тесты для основных функций.

		•Документация: создать документацию с описанием работы приложения и инструкцией по использованию.