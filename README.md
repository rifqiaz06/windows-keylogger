# ⌨️ Windows Keylogger (Educational Purpose Only)

⚠️ **Disclaimer**:
Proyek ini hanya dibuat untuk **tujuan edukasi, pembelajaran, dan riset keamanan**.
**Jangan gunakan untuk aktivitas ilegal atau berbahaya** seperti mencuri data pribadi, password, atau informasi sensitif.
Penulis tidak bertanggung jawab atas penyalahgunaan dari kode ini.

---

## 📖 Overview

Keylogger sederhana berbasis Go untuk Windows yang menggunakan package [`go-hook`](https://github.com/moutend/go-hook).
Script ini akan menangkap event keyboard dan menyimpannya ke file `log.txt`.

Fungsinya lebih cocok untuk:

- Belajar bagaimana **event hook** bekerja di Windows
- Riset keamanan & forensik digital
- Debugging input aplikasi

---

## 🛠️ Installation

1. Pastikan sudah install **Go 1.16+**
   Download di: [https://go.dev/dl/](https://go.dev/dl/)

2. Clone repository:

   ```bash
   git clone https://github.com/username/windows-keylogger.git
   cd windows-keylogger
   ```

3. Install dependency:

   ```bash
   go get github.com/moutend/go-hook/pkg/keyboard
   go get github.com/moutend/go-hook/pkg/types
   ```

---

## 🚀 Usage

Build program:

```bash
go build -o keylogger.exe
```

Jalankan program:

```bash
keylogger.exe
```

Output di terminal:

```
start capturing keyboard input
```

Semua input keyboard akan disimpan ke file `log.txt` di direktori yang sama.

---

## 📝 Example Output

Isi `log.txt`:

```
A B C D 1 2 3 ENTER SPACE
```

---

## ⚠️ Limitations

- Hanya berjalan di **Windows** (karena menggunakan Windows API hook).
- Hanya menangkap event keyboard (tidak termasuk clipboard atau mouse).
- Log tidak terenkripsi → siapa pun yang buka `log.txt` bisa lihat hasilnya.
