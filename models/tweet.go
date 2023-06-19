package models

import (
	"time"
)

type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}

/*Grabo Tweet es el formo o estructura que tendra el tweet*/
type GraboTweet struct {
	UserId  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitemp"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
