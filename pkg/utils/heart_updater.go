package utils

import (
	"log"
	"time"

	"ngodeyuk-core/internal/domain/repositories"
)

// fungsi untuk memulai update otomatis pada heart di user
func StartHeartUpdater(repository repositories.UserRepository, updateInterval time.Duration) {
	// untuk mengatur interval update
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				// untuk mengupdate semua heart yang ada pada user
				UpdateAllUserHearts(repository)
			}
		}
	}()
}

// fungsi untuk mengupdate heart pada semua user
func UpdateAllUserHearts(repository repositories.UserRepository) {
	// untuk mengambil semua data user dari repository
	users, err := repository.FindAll()
	if err != nil {
		log.Printf("error retrieving users: %v", err)
		return
	}

	// untuk mengupdate user yang memiliki heart kurang < 5
	for _, user := range users {
		if user.Heart < 5 {
			user.Heart++
			user.LastHeartTime = time.Now()
			// memperbarui semua pengguna yang ada direpository
			if err := repository.Update(&user); err == nil {
				log.Printf("updated user %s: Heart = %d", user.Username, user.Heart)
			} else {
				log.Printf("error updating user %s: Heart = %d", user.Username, err)
			}
		}
	}
}
