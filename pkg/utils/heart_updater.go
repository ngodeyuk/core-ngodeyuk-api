package utils

import (
	"log"
	"time"

	"ngodeyuk-core/internal/domain/repositories"
)

func StartHeartUpdater(repository repositories.UserRepository, updateInterval time.Duration) {
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				UpdateAllUserHearts(repository)
			}
		}
	}()
}

func UpdateAllUserHearts(repository repositories.UserRepository) {
	users, err := repository.FindAll()
	if err != nil {
		log.Printf("error retrieving users: %v", err)
		return
	}

	for _, user := range users {
		if user.Heart < 5 {
			user.Heart++
			user.LastHeartTime = time.Now()
			if err := repository.Update(&user); err == nil {
				log.Printf("updated user %s: Heart = %d", user.Username, user.Heart)
			} else {
				log.Printf("error updating user %s: Heart = %d", user.Username, err)
			}
		}
	}
}
