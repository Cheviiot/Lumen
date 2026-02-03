# Lumen - Торрент-стриминг приложение

![Lumen](Icons/Lumen(1).png)

Современное приложение для стриминга фильмов и сериалов через торренты.

## 🚀 Технологии

- **Backend**: Go + Wails 2
- **Frontend**: Vue 3 + Vite
- **Стилизация**: Tailwind CSS
- **База данных**: SQLite (локально) + Supabase (синхронизация)
- **API**: TMDB API
- **Торренты**: TorrServer + Jackett

## 📋 Требования

- Go 1.21+
- Node.js 18+
- npm или pnpm
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)
- TorrServer (опционально, для просмотра)
- Jackett (опционально, для поиска торрентов)

## 🛠️ Установка

### 1. Установите Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 2. Клонируйте репозиторий

```bash
git clone https://github.com/your-username/Lumen.git
cd Lumen
```

### 3. Установите зависимости

```bash
# Go зависимости
go mod tidy

# Frontend зависимости
cd frontend
npm install
cd ..
```

### 4. Запустите в режиме разработки

```bash
wails dev
```

### 5. Сборка продакшн версии

```bash
wails build
```

Готовое приложение будет в папке `build/bin/`

## ⚙️ Настройка сервисов

### TorrServer

1. Скачайте TorrServer: https://github.com/YouROK/TorrServer
2. Запустите: `./TorrServer`
3. По умолчанию работает на `http://localhost:8090`

### Jackett

1. Установите Jackett: https://github.com/Jackett/Jackett
2. Запустите и настройте индексаторы
3. По умолчанию работает на `http://localhost:9117`
4. Скопируйте API ключ из настроек Jackett

### Supabase (опционально)

1. Создайте проект на https://supabase.com
2. Создайте таблицы:

```sql
-- Favorites
create table favorites (
  id uuid default gen_random_uuid() primary key,
  user_id text not null,
  tmdb_id integer not null,
  media_type text not null,
  title text,
  poster text,
  year text,
  rating real,
  created_at timestamp with time zone default now()
);

-- History
create table history (
  id uuid default gen_random_uuid() primary key,
  user_id text not null,
  tmdb_id integer not null,
  media_type text not null,
  title text,
  poster text,
  episode text,
  progress real,
  duration integer,
  watched_at timestamp with time zone default now(),
  last_watched timestamp with time zone default now()
);

-- Bookmarks
create table bookmarks (
  id uuid default gen_random_uuid() primary key,
  user_id text not null,
  tmdb_id integer not null,
  media_type text not null,
  title text,
  poster text,
  year text,
  note text,
  created_at timestamp with time zone default now()
);
```

3. Скопируйте URL проекта и anon ключ в настройки приложения

## 📁 Структура проекта

```
Lumen/
├── main.go           # Точка входа Wails
├── app.go            # Управление окном
├── database.go       # Локальная SQLite база
├── tmdb.go           # TMDB API
├── torrserver.go     # TorrServer API
├── jackett.go        # Jackett API
├── supabase.go       # Supabase синхронизация
├── wails.json        # Конфигурация Wails
├── go.mod            # Go модули
├── Icons/            # Иконки приложения
└── frontend/         # Vue 3 приложение
    ├── src/
    │   ├── components/   # UI компоненты
    │   ├── views/        # Страницы
    │   ├── stores/       # Pinia сторы
    │   ├── router/       # Vue Router
    │   ├── App.vue       # Главный компонент
    │   └── main.js       # Точка входа
    ├── package.json
    └── vite.config.js
```

## 🎨 Возможности

- ✅ Кастомный Titlebar с поиском
- ✅ Сворачиваемый Sidebar с 7 вкладками
- ✅ Просмотр популярных фильмов и сериалов
- ✅ Поиск по TMDB
- ✅ Категория мультфильмов
- ✅ Избранное, История, Закладки
- ✅ Поиск торрентов через Jackett
- ✅ Стриминг через TorrServer
- ✅ Встроенный видеоплеер
- ✅ Локальное хранение данных
- ✅ Синхронизация с Supabase (опционально)

## 🔑 Горячие клавиши плеера

- `Пробел` / `K` - Пауза/Воспроизведение
- `←` / `→` - Перемотка на 10 секунд
- `↑` / `↓` - Громкость
- `M` - Выключить звук
- `F` - Полноэкранный режим
- `Escape` - Выход

## 📄 Лицензия

MIT License

## 🙏 Благодарности

- [TMDB](https://www.themoviedb.org/) за API
- [Wails](https://wails.io/) за фреймворк
- [TorrServer](https://github.com/YouROK/TorrServer) за стриминг
- [Jackett](https://github.com/Jackett/Jackett) за поиск