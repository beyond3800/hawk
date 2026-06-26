# Environment Variables

Hawk uses a `.env` file to configure the application.

---

## Application

```env
APP_NAME=Hawk
APP_ENV=development
APP_DEBUG=true
```

---

## Database

```env
DB_ENABLED=true

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=hawk
DB_USERNAME=root
DB_PASSWORD=
```

Disable the database:

```env
DB_ENABLED=false
```

---

## Redis

```env
REDIS_ENABLED=true

REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=
```

Disable Redis:

```env
REDIS_ENABLED=false
```

---

## Notes

* Empty values disable optional services.
* Database and Redis connections are only initialized when enabled.
* Changes to the `.env` file require restarting the application.
