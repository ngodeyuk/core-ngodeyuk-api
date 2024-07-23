#!/bin/bash

# Konfigurasi
MIGRATE_PATH="./migrations"  # Path ke folder migrasi
DATABASE_URL="postgres://upii:secret1212@localhost:5432/ngodeyuk?sslmode=disable"  # URL database

# Fungsi untuk memeriksa apakah go-migrate sudah terinstal
check_migrate_installed() {
  if command -v migrate >/dev/null 2>&1; then
    echo "go-migrate sudah terinstal."
    return 0
  else
    echo "go-migrate tidak terinstal."
    return 1
  fi
}

# Fungsi untuk menginstal go-migrate
install_migrate() {
  echo "Menginstal go-migrate..."
  # Unduh dan pasang go-migrate
  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz -o migrate.tar.gz
  tar -xzf migrate.tar.gz
  sudo mv migrate /usr/local/bin/
  rm migrate.tar.gz
  echo "go-migrate berhasil diinstal."
}

# Fungsi untuk menerapkan migrasi
apply_migrations() {
  echo "Menerapkan migrasi..."
  migrate -path "$MIGRATE_PATH" -database "$DATABASE_URL" up
}

# Fungsi untuk membatalkan migrasi
rollback_migrations() {
  echo "Membatalkan migrasi..."
  migrate -path "$MIGRATE_PATH" -database "$DATABASE_URL" down
}

# Tanyakan kepada pengguna apakah mereka ingin menginstal go-migrate
echo "Apakah Anda sudah menginstal go-migrate? (y/n)"
read -p "Masukkan pilihan: " install_choice

if [[ "$install_choice" == "y" || "$install_choice" == "Y" ]]; then
  if ! check_migrate_installed; then
    install_migrate
  fi
fi

# Pilihan menu
echo "Pilih opsi:"
echo "1) Terapkan migrasi"
echo "2) Batalkan migrasi"
read -p "Masukkan pilihan (1/2): " choice

case $choice in
  1)
    apply_migrations
    ;;
  2)
    rollback_migrations
    ;;
  *)
    echo "Pilihan tidak valid"
    ;;
esac
