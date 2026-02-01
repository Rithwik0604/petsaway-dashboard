package clients

import "petsaway/internal/database"

func GetAllClients() ([]database.Client, error) {
	clients, err := database.DB.GetAllClients(database.Qctx)
	if err != nil {
		return []database.Client{}, err
	}

	return clients, err
}
