package main

import "github.com/google/uuid"

type Link struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Url       string    `bun:",notnull"`
	Date      string    `bun:"type:timestamp,default:now()"`
	Visits    int       `bun:",notnull,default:0"`
	LastVisit string    `bun:"type:timestamp,default:now()"`
}
