<h1 align="center">⚽ Foot-Star</h1>

<p align="center">
  <a href="https://golang.org/"><img alt="Go Version" src="https://img.shields.io/badge/Go-1.16+-blue"></a>
  <a href="LICENSE"><img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-green.svg"></a>
  <a href="https://github.com/S-n-g-1/Foot-Star/issues"><img alt="Issues" src="https://img.shields.io/github/issues/S-n-g-1/Foot-Star"></a>
  <a href="https://github.com/S-n-g-1/Foot-Star/stargazers"><img alt="Stars" src="https://img.shields.io/github/stars/S-n-g-1/Foot-Star?style=social"></a>
</p>

<p align="center"><i>Simulasi sepak bola ringan berbasis teks, dibuat dengan Go.</i></p>

---

## ✨ Fitur Utama

<ul>
  <li>✅ Dibuat 100% dengan <b>Go</b>, tanpa dependensi eksternal</li>
  <li>🎮 Simulasi karier sepak bola berbasis teks</li>
  <li>🖥️ Antarmuka CLI yang sederhana dan intuitif</li>
  <li>⚡ Performa cepat, ringan, dan hemat sumber daya</li>
  <li>💼 Struktur kode modular dan mudah dikembangkan</li>
</ul>

---

## 📦 Instalasi

Pastikan <b>Go 1.16+</b> sudah terinstal di sistem Anda.

### Clone & Build Manual
```bash

 git clone https://github.com/S-n-g-1/Foot-Star.git
 
 cd Foot-Star
 go build -o cmd/main.go
 ./cmd/main.go
```


📁 Struktur Proyek

<pre>
Foot-Sta
.
├── ai
│   ├── match_ai.go
│   ├── scouting_ai.go
│   ├── tactic_ai.go
│   └── transfer_ai.go
├── assets
│   ├── data
│   │   ├── json
│   │   │   └── player.json
│   │   └── sqlite
│   │       └── Foot-Star.db
│   └── font
│       └── PressStart2P-Regular.ttf
├── cmd
│   └── main.go
├── config
│   └── default_settings.go
├── controllers
│   ├── akademi_controller.go
│   ├── club_controller.go
│   ├── event_controller.go
│   ├── formasi_controller.go
│   ├── history_controller.go
│   ├── jadwal_controller.go
│   ├── manager_controller.go
│   ├── match_controller.go
│   ├── notification_controller.go
│   ├── pengelolaan_controller.go
│   ├── player_controller.go
│   ├── savegame_controller.go
│   ├── stadion_controller.go
│   ├── taktik_controller.go
│   └── transfer_controller.go
├── docs
│   ├── API_REFERENCE.md
│   ├── DESIGN.md
│   └── GAME_FLOW.md
├── go.mod
├── go.sum
├── locales
│   ├── en.json
│   └── id.json
├── models
│   ├── club.go
│   ├── commentary.go
│   ├── history.go
│   ├── league.go
│   ├── manager.go
│   ├── match_event.go
│   ├── match.go
│   ├── player.go
│   ├── schedule.go
│   ├── tactic.go
│   ├── transfer.go
│   ├── type.go
│   └── youth_player.go
├── notifications
│   ├── builder.go
│   ├── center.go
│   ├── handler.go
│   ├── README.md
│   └── types.go
├── README.md
├── save
│   ├── auto_save.json
│   ├── save1.json
│   └── save2.json
├── services
│   ├── match_service.go
│   ├── scouting_service.go
│   ├── tactic_service.go
│   ├── training_service.go
│   └── transfer_service.go
├── tests
│   ├── match_test.go
│   ├── player_test.go
│   └── transfer_test.go
├── utils
│   └── logic.go
└── views
    ├── lineup.go
    ├── main_menu.go
    ├── pengelolaan.go
    └── scene
        ├── akademi
        │   └── akademi.go
        ├── dialog
        │   └── dialog.go
        ├── formasi
        │   └── formasi.go
        ├── schedule
        │   └── jadwal_pertandingan.go
        ├── stadion
        │   ├── match_day.go
        │   ├── stadion.go
        │   └── upgrade.go
        ├── training
        │   └── latihan.go
        └── transfer
            ├── daftar_pemain.go
            └── detail_pemain.go
           # Lisensi MIT
</pre>

---

‍💻 Kontribusi

Kontribusi sangat terbuka dan disambut! 💡

<b>Cara Berkontribusi:</b>

<ol>
  <li>Fork repositori ini</li>
  <li>Buat branch fitur baru:<br><code>git checkout -b fitur/nama-fitur</code></li>
  <li>Lakukan perubahan & commit:<br><code>git commit -m "Menambahkan fitur X"</code></li>
  <li>Push ke branch Anda:<br><code>git push origin fitur/nama-fitur</code></li>
  <li>Buat Pull Request ke branch <code>main</code></li>
</ol>
---

🛣️ Roadmap

<ul>
  <li>[ ] Mode karier pemain</li>
  <li>[ ] Statistik & performa klub</li>
  <li>[ ] Leaderboard lokal</li>
  <li>[ ] Porting ke web atau terminal UI</li>
  <li>[ ] Dukungan ekspor hasil pertandingan</li>
</ul>
---

📄 Lisensi

Proyek ini menggunakan <a href="LICENSE">MIT License</a> — bebas digunakan, dimodifikasi, dan disebarluaskan dengan tetap mencantumkan kredit.


---

🙋 Kontak & Dukungan

Ajukan bug, saran, atau pertanyaan melalui <a href="https://github.com/S-n-g-1/Foot-Star/issues">GitHub Issues</a>

Pemilik proyek: <a href="https://github.com/S-n-g-1">S-n-g-1</a>



---

<p align="center"><i>Dibuat dengan semangat open-source 💙 oleh developer Go, <b>S-n-g-1</b></i></p>