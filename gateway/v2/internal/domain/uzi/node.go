package domain

import "github.com/google/uuid"

type Node struct {
	Id       uuid.UUID
	Ai       bool
	UziID    uuid.UUID
	Tirads23 float64
	Tirads4  float64
	Tirads5  float64
}
