Flash Sale Service

Bu proje, Go (Golang), Fiber, PostgreSQL ve Docker kullanılarak geliştirilmiş, yüksek trafik altında çalışan bir stok yönetim servisidir.

Özellikle E-Ticaret sitelerindeki "Black Friday" gibi yoğun dönemlerde ortaya çıkan Race Condition problemini simüle eder ve Database Locking  yöntemiyle çözer.

 Kullanılan Teknolojiler

Go: Backend servisi
Fiber: Hızlı HTTP Web Framework
GORM: Veritabanı yönetimi 
PostgreSQL: Veritabanı
Docker & Docker Compose: Konteynerizasyon

Senaryo:
Depoya 100 adet iPhone 15 eklenir.
"attacker" scripti ile sisteme aynı anda 1000 adet satın alma isteği gönderilir.
Race Condition Koruması: SELECT .. FOR UPDATE kilitleme mekanizması sayesinde stok hatasız bir şekilde düşülür ve eksiye düşmesi engellenir.
