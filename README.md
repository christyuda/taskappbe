TaskApp Backend merupakan implementasi bagian belakang (backend) sebuah aplikasi menggunakan bahasa pemrograman Golang. Backend ini dirancang untuk mengelola berbagai operasi terkait tugas (tasks) dalam aplikasi. Dalam backend ini, terdapat sejumlah endpoint yang memungkinkan penggunaan berbagai fitur, seperti:

1. **GET Tasks**: Menggunakan endpoint ini, pengguna dapat mengambil daftar tugas yang ada dalam sistem.

2. **POST Tasks**: Endpoint ini memungkinkan pengguna untuk membuat tugas baru dengan mengirimkan data yang diperlukan.

3. **PUT Tasks**: Pengguna dapat memperbarui informasi tugas yang sudah ada menggunakan endpoint ini.

4. **DELETE Tasks**: Endpoint ini digunakan untuk menghapus tugas yang sudah ada dari sistem.

Selain operasi dasar pada tugas, backend ini juga mendukung fitur tambahan:

- **POST Subtasks**: Pengguna dapat membuat subtask yang terkait dengan tugas utama. Subtask merupakan bagian lebih rinci dari sebuah tugas utama.

- **POST Attachment**: Endpoint ini memungkinkan pengguna untuk melampirkan berkas atau file terkait dengan tugas tertentu.

- **DELETE Subtasks**: Pengguna juga dapat menghapus subtask yang tidak lagi diperlukan.

- **DELETE Attachment**: Fitur ini memungkinkan pengguna untuk menghapus berkas terlampir dari sebuah tugas.

Dengan adanya backend TaskApp ini, pengguna dapat melakukan manajemen tugas secara efisien, melacak status tugas, serta mengelola detail-detail terkait tugas seperti subtask dan lampiran. Semua operasi tersebut dapat diakses melalui berbagai endpoint yang disediakan, memungkinkan integrasi yang lebih baik dengan aplikasi frontend dan memberikan pengalaman pengguna yang lebih baik dalam mengelola tugas sehari-hari.
