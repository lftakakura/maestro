// maestro
// https://github.com/topfreegames/maestro
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package models

import (
	"time"

	"github.com/topfreegames/extensions/interfaces"
	pg "gopkg.in/pg.v5"
)

// Room is the struct that defines a room in maestro
type Room struct {
	ID         string      `db:"id"`
	ConfigID   string      `db:"config_id"`
	Status     string      `db:"status"`
	LastPingAt pg.NullTime `db:"last_ping_at"`
}

// RoomStatus is the struct that defines a room status
type RoomStatus struct {
	Count  int
	Status string
}

// RoomsStatusCount is the struct that defines the rooms status status count
type RoomsStatusCount struct {
	Creating    int
	Occupied    int
	Ready       int
	Terminating int
	Total       int
}

// NewRoom is the room constructor
func NewRoom(id, configID string) *Room {
	return &Room{
		ID:       id,
		ConfigID: configID,
		Status:   "creating",
	}
}

// Create creates a room in the database
func (r *Room) Create(db interfaces.DB) error {
	_, err := db.Query(r, `
		INSERT INTO rooms (id, config_id, status) VALUES (?id, ?config_id, ?status)
		RETURNING id
	`, r)
	return err
}

// SetStatus updates the status of a given room in the database
func (r *Room) SetStatus(db interfaces.DB, status string) error {
	_, err := db.Query(r, `
		UPDATE rooms SET status = ?, last_ping_at = ? WHERE id = ?
	`, status, time.Now(), r.ID)
	return err
}

// Ping updates the last_ping_at field of a given room in the database
func (r *Room) Ping(db interfaces.DB) error {
	_, err := db.Query(r, `UPDATE rooms SET last_ping_at = ? WHERE id = ?`, time.Now(), r.ID)
	return err
}

// GetRoomsCountByStatus returns the count of rooms for each status
func GetRoomsCountByStatus(db interfaces.DB, configID string) (*RoomsStatusCount, error) {
	roomStatuses := []*RoomStatus{}
	_, err := db.Query(
		roomStatuses,
		`SELECT COUNT(*) as count, status FROM rooms WHERE config_id = ? GROUP BY status`,
		configID,
	)
	if err != nil {
		return nil, err
	}
	countByStatus := &RoomsStatusCount{}
	for _, rs := range roomStatuses {
		switch status := rs.Status; status {
		case "creating":
			countByStatus.Creating = rs.Count
			countByStatus.Total += rs.Count
		case "ready":
			countByStatus.Ready = rs.Count
			countByStatus.Total += rs.Count
		case "occupied":
			countByStatus.Occupied = rs.Count
			countByStatus.Total += rs.Count
		case "terminating":
			countByStatus.Terminating = rs.Count
			countByStatus.Total += rs.Count
		default:
			countByStatus.Total += rs.Count
		}
	}
	return countByStatus, nil
}
