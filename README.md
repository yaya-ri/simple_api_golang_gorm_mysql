
## Dokumentasi

Dalam API yang saya buat terdiri dari 2 endpoint GET dan 2 endpoint POST sesuai deskripsi.
    
## EndPoint utama

    - localhost:3000/classes (GET)  //menampilkan semua class
    - localhost:3000/students/{class_id} (GET) //menampilkan class beserta student disetiap classnya dengan parameter id kelas
    - localhost:3000/class (POST) //menambahkan data pada tabel classes
    - localhost:3000/student (POST) //menambahkan data pada tabel students

## Library/Package
pada pengerjaan API ini saya menggunakan beberapa library/package tambahan berupa:

    - "github.com/gin-gonic/gin"  
    - "github.com/go-sql-driver/mysql"
    - "github.com/jinzhu/gorm"
    - "github.com/jinzhu/gorm/dialects/mysql"

## **NOTE
masih banyak error dalam pengerjaanya. karna memang ada beberapa kendala teknis dan dalam tahap belajar juga. terimakasih
