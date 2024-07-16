# Banking App

Bu proje, bir bankacılık uygulamasında para transferi işlemlerini gerçekleştirmek için geliştirilmiştir. Uygulama, kullanıcı hesaplarının oluşturulmasını, bakiye kontrollerini ve para transferlerini içermektedir.

## Özellikler

- Kullanıcı hesapları oluşturma
- Kullanıcılar arasında para transferi
- Hesap bakiyesi kontrolü
- Kimlik doğrulama ve yetkilendirme
- Veritabanı yönetimi (MySQL)

## Kurulum

### Gereksinimler

- Go 1.11 veya üzeri
- MySQL veritabanı
- Docker (isteğe bağlı)

### Adımlar

1. Bu repository'yi klonlayın:
    ```bash
    git clone https://github.com/Saudade01/banking-app.git
    cd banking-app
    ```

2. Bağımlılıkları yükleyin:
    ```bash
    go mod tidy
    ```

3. `.env` dosyasını oluşturun ve veritabanı bağlantı bilgilerinizi girin:
    ```env
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_HOST=localhost:3306
    DB_NAME=banking
    ```

4. Veritabanını oluşturun ve gerekli tabloları ekleyin:
    ```sql
    CREATE DATABASE banking;

    USE banking;

    CREATE TABLE users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL
    );

    CREATE TABLE accounts (
        id INT AUTO_INCREMENT PRIMARY KEY,
        owner VARCHAR(255) NOT NULL,
        balance DECIMAL(10, 2) NOT NULL,
        currency VARCHAR(10) NOT NULL
    );

    CREATE TABLE transfers (
        id INT AUTO_INCREMENT PRIMARY KEY,
        from_account_id INT NOT NULL,
        to_account_id INT NOT NULL,
        amount DECIMAL(10, 2) NOT NULL,
        FOREIGN KEY (from_account_id) REFERENCES accounts(id),
        FOREIGN KEY (to_account_id) REFERENCES accounts(id)
    );
    ```

5. Uygulamayı çalıştırın:
    ```bash
    go run main.go
    ```

## API Kullanımı

### Kayıt Olma

- **Endpoint:** `POST /register`
- **Body:**
    ```json
    {
        "username": "testuser",
        "password": "password123"
    }
    ```

### Giriş Yapma

- **Endpoint:** `POST /login`
- **Body:**
    ```json
    {
        "username": "testuser",
        "password": "password123"
    }
    ```

### Hesap Oluşturma

- **Endpoint:** `POST /api/accounts`
- **Headers:**
    ```http
    Authorization: Bearer <token>
    Content-Type: application/json
    ```
- **Body:**
    ```json
    {
        "owner": "testuser",
        "balance": 1000.0,
        "currency": "USD"
    }
    ```

### Hesap Bilgilerini Alma

- **Endpoint:** `GET /api/accounts/{id}`
- **Headers:**
    ```http
    Authorization: Bearer <token>
    ```

### Para Transferi

- **Endpoint:** `POST /api/transfers`
- **Headers:**
    ```http
    Authorization: Bearer <token>
    Content-Type: application/json
    ```
- **Body:**
    ```json
    {
        "from_account_id": 1,
        "to_account_id": 2,
        "amount": 100.0
    }
    ```

### Transferleri Listeleme

- **Endpoint:** `GET /api/transfers`
- **Headers:**
    ```http
    Authorization: Bearer <token>
    ```
    

## Katkıda Bulunma

Katkıda bulunmak isterseniz, lütfen bir `issue` açın veya bir `pull request` gönderin.


