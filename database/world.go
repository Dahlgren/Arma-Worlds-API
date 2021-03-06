package database

import (
	"context"
	"fmt"
	"os"
)

// World database model
type World struct {
	ID              string `json:"id" db:"id"`
	Name            string `json:"name" db:"name"`
	Title           string `json:"title" db:"title"`
	Size            int    `json:"size" db:"size"`
	SteamWorkshopID *int   `json:"steamWorkshopId" db:"steam_workshop_id"`
}

// FetchWorlds loads all worlds from database
func (db *Database) FetchWorlds(ctx context.Context) (*[]World, error) {
	var query = `
		SELECT
			worlds.id,
			worlds.name,
			worlds.title,
			worlds.size,
			worlds.steam_workshop_id
		FROM
			worlds
		ORDER BY
			worlds.title
	`
	rows, err := db.client.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var worlds []World

	for rows.Next() {
		var world World
		err = rows.Scan(&world.ID, &world.Name, &world.Title, &world.Size, &world.SteamWorkshopID)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading user from database: %q\n", err)
		} else {
			worlds = append(worlds, world)
		}
	}

	rows.Close()

	return &worlds, nil
}
