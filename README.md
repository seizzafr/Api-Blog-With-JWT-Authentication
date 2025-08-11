Oke, ini isi **README.md**-nya saja, siap kamu copy-paste ke file `README.md` di repo-mu.

```md
# API-BLOG (JWT Authentication)

## Deskripsi
RESTful API dibangun dengan **Golang** yang menerapkan **JWT (JSON Web Token)** untuk autentikasi. Project ini menyediakan fitur register, login, logout (token blacklist), serta CRUD untuk `posts`, `categories`, dan `tags` dengan relasi many-to-many (`post_tag`). Dokumentasi Postman dan ERD disertakan di folder `docs/`.

## Fitur
- User: Register, Login, Logout
- JWT Authentication (dengan secret di `.env`)
- Token validation & blacklist (logout)
- CRUD Posts, Categories, Tags
- Many-to-Many relation: Posts ↔ Tags lewat `post_tag`
- Postman collection & ERD ada di `docs/`
- Menjalankan server dengan **air** (live reload)

## Tech Stack
- Golang
- GORM (ORM)
- MySQL / MariaDB
- JWT
- air (live-reload)

## Struktur Project
```

.
├─ docs/
│  ├─ API-BLOG.postman\_collection.json
│  ├─ ERD\_BLOG.drawio.png
│  └─ database/
│     └─ db\_blog.sql
├─ src/
├─ .env.example
├─ .gitignore
└─ README.md

````

## Setup
1. Clone repo:
```bash
git clone https://github.com/seizzafr/Api-Blog-With-JWT-Authentication.git
cd Api-Blog-With-JWT-Authentication 
````

2. Copy `.env.example` → `.env` dan isi variabelnya:

```env
DB_USER=your_db_user
DB_PASS=your_db_pass
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=db_blog
JWT_SECRET=your_jwt_secret
JWT_EXPIRES=24h
PORT=3000
```

3. Import database:

```bash
mysql -u root -p
CREATE DATABASE db_blog;
exit
mysql -u root -p db_blog < docs/database/db_blog.sql
```

4. Install dependencies:

```bash
go mod tidy
```

5. Jalankan server:

```bash
air
```

Atau:

```bash
go run ./cmd/main.go
```

## Dokumentasi

* Postman: `docs/API-BLOG.postman_collection.json`
* ERD: `docs/ERD_BLOG.drawio.png`

## License

MIT © 2025
Sheiza Fakhru Rasyid
```

Kamu mau aku tambahin sekalian?
```
