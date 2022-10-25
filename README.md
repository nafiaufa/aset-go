## Dependency in this Go program

```sh
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/gorilla/mux v1.8.0
	gorm.io/driver/postgres v1.4.5
	gorm.io/gorm v1.24.1-0.20221019064659-5dd2bb482755
```

## Installation

- `git clone https://github.com/nafiaufa/aset-go.git`
- `cd aset-go`
- Edit models.setup to Set your database connection details
- `go run main.go`

## API Endpoints

| Methods  | Endpoints             | Description                                                           |
| :------- | :-------------------  | :-------------------------------------------------------------------- |
| `POST`   | /login                | login account must given `username` & `password` to body request      |
| `POST`   | /register             | Register must given `nama_lengkap`,`username`,`password` to body      |                                                   |
| `GET`    | /api/asets            | Get All aset                                                          |
| `POST`   | /api/aset             | Create aset                                                           |
| `GET`    | /api/aset/:id         | Get aset By Id                                                        |
| `PUT`    | /api/aset/:id         | Edit aset By Id                                                       |
| `DELETE` | /api/aset             | Delete aset By Id                                                     |
| `POST`   | /api/file             | Upload Image invoice                                                  |
| `GET`    | /logout               | Logout account                                                        |
