package events

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ZoneCreatedEvent struct {
	Subject string          `json:"subject"`
	Data    ZoneCreatedData `json:"data"`
}

type ZoneCreatedData struct {
	TraceID             string             `json:"traceID"`
	ID                  primitive.ObjectID `json:"id,omitempty"`
	Name                string             `json:"name"`
	ZoneType            string             `json:"zone_type"`
	Email               string             `json:"email"`
	Mobile              int                `json:"mobile"`
	WhatsApp            int                `json:"whatsapp"`
	GSTNumber           string             `json:"gst_number"`
	PANNumber           string             `json:"pan_number"`
	IsBillingEnabled    bool               `json:"is_billing_enabled"`
	IsUserPortalEnabled bool               `json:"is_user_portal_enabled"`
	IsActive            bool               `bson:"is_active"`
	DocVersion          int                `bson:"doc_version"`
	PermanentAddress    Address            `json:"permanent_address"`
	CurrentAddress      Address            `json:"current_address"`
	BillingAddress      Address            `json:"billing_address"`
}

type Address struct {
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	District    string `json:"district"`
	Pincode     int    `json:"pincode"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	LandMark    string `json:"landmark"`
}
