# Go Auth System (JWT & Bitmask Permissions)

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)
![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)

Bu proje, Go (Golang) kullanılarak geliştirilmiş, güvenli, genişletilebilir ve profesyonel bir kimlik doğrulama (Authentication) ve yetkilendirme (Authorization) sistemidir.

Özellikle Discord gibi platformlarda kullanılan **Bitmask (Bit Düzeyinde)** yetkilendirme yapısını simüle eder. Bu sayede tek bir tamsayı (integer) içinde birden fazla yetki saklanabilir ve yönetilebilir.

## Özellikler

*   **JWT (JSON Web Token)** tabanlı kimlik doğrulama.
*   **Bcrypt** ile güvenli şifre saklama.
*   **Bitmask Yetkilendirme Sistemi**: Esnek ve performanslı yetki kontrolü.
*   **Middleware** yapısı ile korumalı rotalar.
*   **Temiz Mimari**: Models, Handlers, Utils ayrımı.
*   **Birim Testleri**: Kapsamlı test senaryoları.

## Kurulum

1.  Projeyi klonlayın veya indirin.
2.  Gerekli Go modüllerini yükleyin:
    ```bash
    go mod tidy
    ```
3.  Sunucuyu başlatın:
    ```bash
    go run main.go
    ```

## Yetki Sistemi (Permissions)

Bu sistemde yetkiler 2'nin kuvvetleri (bitler) olarak tanımlanmıştır. İstediğiniz yetkileri toplayarak kullanıcıya atayabilirsiniz.

| Yetki Adı | Değer (Decimal) | Bit Değeri | Açıklama |
| :--- | :--- | :--- | :--- |
| `PermRead` | 1 | `0001` | Okuma yetkisi |
| `PermWrite` | 2 | `0010` | Yazma yetkisi |
| `PermDelete` | 4 | `0100` | Silme yetkisi |
| `PermAdmin` | 8 | `1000` | Yönetici yetkisi |

**Örnek Kombinasyonlar:**
*   **Sadece Okuma**: `1`
*   **Okuma + Yazma**: `1 + 2 = 3`
*   **Tam Yetki (Admin)**: `1 + 2 + 4 + 8 = 15`

## API Kullanımı

### 1. Kayıt Ol (Register)

Yeni bir kullanıcı oluşturur. `permissions` alanı ile yetkileri belirlenir.

*   **URL**: `/register`
*   **Method**: `POST`
*   **Body**:
    ```json
    {
        "username": "kullanici",
        "password": "sifre123",
        "permissions": 3  // (Okuma + Yazma)
    }
    ```

### 2. Giriş Yap (Login)

Giriş yapar ve JWT token döner.

*   **URL**: `/login`
*   **Method**: `POST`
*   **Body**:
    ```json
    {
        "username": "kullanici",
        "password": "sifre123"
    }
    ```
*   **Response**:
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIs..."
    }
    ```

### 3. Korumalı Rotalara Erişim

Aşağıdaki rotalara erişmek için `Authorization` header'ında token gönderilmelidir.

**Header:**
`Authorization: Bearer <TOKEN>`

| Rota | Method | Gerekli Yetki | Açıklama |
| :--- | :--- | :--- | :--- |
| `/home` | GET | - | Giriş yapmış herkes erişebilir. |
| `/users` | GET | `PermRead (1)` | Tüm kullanıcıları listeler. |
| `/admin` | GET | `PermAdmin (8)` | Sadece yöneticiler erişebilir. |

## Test Etme

Projedeki birim testlerini çalıştırmak için:

```bash
go test ./... -v
```

## Proje Yapısı

*   `main.go`: Sunucu ayarları ve rotalar.
*   `models/`: Veri yapıları ve veritabanı simülasyonu.
*   `handlers/`: HTTP isteklerini işleyen fonksiyonlar.
*   `utils/`: Yardımcı araçlar (JWT vb.).
