package clients

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"petsaway/internal/database"
	"reflect"
	"time"

	"github.com/google/uuid"
)

var (
	MailServer  string
	TargetEmail string
	Username    string
	SmtpServer  string
	SmtpPort    string
	SmtpUser    string
	SmtpPass    string
)

func GetAllClients() ([]database.Client, error) {
	clients, err := database.DB.GetAllClients(database.Qctx)
	if err != nil {
		return []database.Client{}, err
	}

	return clients, err
}

func AddClient(clientDetails ClientDTO) (database.Client, error) {
	var params database.InsertClientParams

	MapDTOToParams(&clientDetails, &params)

	params.ID = uuid.NewString()

	client, err := database.DB.InsertClient(database.Qctx, params)
	if err != nil {
		return database.Client{}, err
	}

	return client, nil
}

func GetClientById(id string) (database.Client, error) {
	client, err := database.DB.GetClientById(database.Qctx, id)
	return client, err
}

func UpdateClient(id string, updates ClientDTO) (database.Client, error) {
	// 1. Fetch the CURRENT data from the database
	currentClient, err := database.DB.GetClientById(database.Qctx, id)
	if err != nil {
		return database.Client{}, err
	}

	// 2. Initialize params with EVERY current value from the DB
	params := database.UpdateClientByIdParams{
		ID:                      id,
		ClientName:              currentClient.ClientName,
		ClientPhone:             currentClient.ClientPhone,
		ImportExport:            currentClient.ImportExport,
		ImportFee:               currentClient.ImportFee,
		ExportFee:               currentClient.ExportFee,
		AfterHoursCharges:       currentClient.AfterHoursCharges,
		PetName:                 currentClient.PetName,
		Species:                 currentClient.Species,
		Gender:                  currentClient.Gender,
		Breed:                   currentClient.Breed,
		DateOfBirth:             currentClient.DateOfBirth,
		MicrochipNumber:         currentClient.MicrochipNumber,
		MicrochipValidity:       currentClient.MicrochipValidity,
		Titre:                   currentClient.Titre,
		LastRabiesDate:          currentClient.LastRabiesDate,
		RabiesValidity:          currentClient.RabiesValidity,
		DocumentationStatus:     currentClient.DocumentationStatus,
		RabiesVaccinationValid:  currentClient.RabiesVaccinationValid,
		OtherVaccinesCompleted:  currentClient.OtherVaccinesCompleted,
		HealthCertificateIssues: currentClient.HealthCertificateIssues,
		ExportPermitApproved:    currentClient.ExportPermitApproved,
		ImportPermitApproved:    currentClient.ImportPermitApproved,
		AirlineApprovalReceived: currentClient.AirlineApprovalReceived,
		CustomsClearanceDone:    currentClient.CustomsClearanceDone,
		OriginCountry:           currentClient.OriginCountry,
		DestinationCountry:      currentClient.DestinationCountry,
		ForwarderCharges:        currentClient.ForwarderCharges,
		Departure:               currentClient.Departure,
		Airline:                 currentClient.Airline,
		AirlineCharges:          currentClient.AirlineCharges,
		CrateCost:               currentClient.CrateCost,
		FlightNumber:            currentClient.FlightNumber,
		TypeOfTravel:            currentClient.TypeOfTravel,
		Etd:                     currentClient.Etd,
		Eta:                     currentClient.Eta,
		QuotedAmount:            currentClient.QuotedAmount,
		TotalCost:               currentClient.TotalCost,
		Profit:                  currentClient.Profit,
		AdvancedReceived:        currentClient.AdvancedReceived,
		BalancePending:          currentClient.BalancePending,
		PaymentStatus:           currentClient.PaymentStatus,
		Remarks:                 currentClient.Remarks,
	}

	// 3. Overlay the updates from the DTO (Reflection only changes what was in JSON)
	MapDTOToParams(&updates, &params)

	// if params.QuotedAmount.Valid && params.QuotedAmount.Float64 != 0 && params.TotalCost.Valid && params.TotalCost.Float64 != 0 {
	// 	params.Profit.Float64 = params.QuotedAmount.Float64 - params.TotalCost.Float64
	// }

	// 4. Save the full record back (No data is lost because we filled 'params' with current data first)
	updatedClient, err := database.DB.UpdateClientById(database.Qctx, params)
	if err != nil {
		return database.Client{}, err
	}

	return updatedClient, nil
}

func DeleteClient(id string) error {
	err := database.DB.DeleteClientById(database.Qctx, id)
	return err
}

// EMAILS ----------------

func getExpiringMicrochip() ([]database.GetExpiringMicrochipRow, error) {
	clients, err := database.DB.GetExpiringMicrochip(database.Qctx)
	if err != nil {
		return []database.GetExpiringMicrochipRow{}, err
	}
	return clients, nil
}

func getExpiringRabies() ([]database.GetExpiringRabiesRow, error) {
	clients, err := database.DB.GetExpiringRabies(database.Qctx)
	if err != nil {
		return []database.GetExpiringRabiesRow{}, err
	}
	return clients, nil
}

func setMicrochipNotified(id string) error {
	err := database.DB.SetMicrochipNotified(database.Qctx, id)
	return err
}

func setRabiesNotified(id string) error {
	err := database.DB.SetRabiesNotified(database.Qctx, id)
	return err
}

func SetupEnvVars() {
	MailServer = os.Getenv("MAILSERVER")
	TargetEmail = os.Getenv("TARGETEMAIL")
	Username = os.Getenv("USERNAME")
	SmtpServer = os.Getenv("SMTPHOST")
	SmtpPort = os.Getenv("SMTPPORT")
	SmtpUser = os.Getenv("SMTPUSER")
	SmtpPass = os.Getenv("SMTPPASS")
}

func RunExpirationCheck() {
	log.Println("Running expiration check...")

	// get microchip expirations
	microchipExpirations, err := getExpiringMicrochip()
	if err != nil {
		log.Println("Error getting microchip expirations:", err.Error())
		microchipExpirations = nil
	}

	// get rabies expirations
	rabiesExpirations, err := getExpiringRabies()
	if err != nil {
		log.Println("Error getting rabies expirations:", err.Error())
		rabiesExpirations = nil
	}

	if microchipExpirations != nil {
		petEntry := []PetEntry{}
		length := len(microchipExpirations)
		for _, v := range microchipExpirations {
			petEntry = append(petEntry, PetEntry{
				ID:      v.ID,
				PetName: v.ClientName.String,
				Phone:   v.ClientPhone.String,
				Date:    v.MicrochipValidity.Time.String(),
			})
		}
		err := sendGroupedEmail(TargetEmail, EmailData{
			Username:   Username,
			ExpiryType: "Microchip",
			Items:      petEntry,
		})
		if err != nil {
			log.Println("Error sending microchip email: ", err.Error())
		} else {
			log.Printf("Sent %d Microchip Emails\n", length)
			for _, v := range microchipExpirations {
				if err := setMicrochipNotified(v.ID); err != nil {
					log.Println("Error setting microchip validity notified. Id: ", v.ID)
				}
			}
		}
	}

	if rabiesExpirations != nil {
		petEntry := []PetEntry{}
		length := len(rabiesExpirations)
		for _, v := range rabiesExpirations {
			petEntry = append(petEntry, PetEntry{
				ID:      v.ID,
				PetName: v.ClientName.String,
				Phone:   v.ClientPhone.String,
				Date:    v.RabiesValidity.Time.String(),
			})
		}
		err := sendGroupedEmail(TargetEmail, EmailData{
			Username:   Username,
			ExpiryType: "Rabies",
			Items:      petEntry,
		})
		if err != nil {
			log.Println("Error sending rabies email: ", err.Error())
		} else {
			log.Printf("Sent %d Rabies Emails\n", length)

			for _, v := range rabiesExpirations {
				if err := setRabiesNotified(v.ID); err != nil {
					log.Println("Error setting rabies validity notified. Id: ", v.ID)
				}
			}
		}
	}

	if microchipExpirations == nil && rabiesExpirations == nil {
		log.Println("No expiration emails to send")
	}
}

// helpers

func MapDTOToParams(dto any, params any) {
	dtoVal := reflect.ValueOf(dto).Elem()
	paramsVal := reflect.ValueOf(params).Elem()

	for i := 0; i < dtoVal.NumField(); i++ {
		typeName := dtoVal.Type().Field(i).Name
		dtoField := dtoVal.Field(i)

		paramField := paramsVal.FieldByName(typeName)
		if !paramField.IsValid() || !paramField.CanSet() {
			continue
		}

		if dtoField.Kind() == reflect.Pointer && !dtoField.IsNil() {
			val := dtoField.Elem()
			targetType := paramField.Type()

			switch v := val.Interface().(type) {
			case string:
				if targetType == reflect.TypeOf(sql.NullTime{}) {
					if v == "" {
						paramField.Set(reflect.ValueOf(sql.NullTime{Valid: false}))
						continue
					}

					var t time.Time
					var err error

					// Try parsing as Datetime-local (ISO-ish) first
					t, err = time.Parse("2006-01-02T15:04", v)
					if err != nil {
						// Fallback to standard Date format if the first fails
						t, err = time.Parse("2006-01-02", v)
					}

					if err == nil {
						paramField.Set(reflect.ValueOf(sql.NullTime{Time: t, Valid: true}))
					} else {
						// Log the error so you can see why parsing failed
						fmt.Printf("Failed to parse date %s: %v\n", v, err)
					}
				} else if targetType == reflect.TypeOf(sql.NullString{}) {
					if v == "" {
						paramField.Set(reflect.ValueOf(sql.NullString{Valid: false}))
					} else {
						paramField.Set(reflect.ValueOf(sql.NullString{String: v, Valid: true}))
					}
				} else if targetType.Kind() == reflect.Interface {
					// Handles ETD/ETA being generated as interface{}
					if v == "" {
						paramField.Set(reflect.Zero(targetType)) // Sets it to nil
					} else {
						paramField.Set(reflect.ValueOf(v)) // Sets it to the string "17:12"
					}
				}

			case float64:
				if targetType == reflect.TypeOf(sql.NullFloat64{}) {
					paramField.Set(reflect.ValueOf(sql.NullFloat64{Float64: v, Valid: true}))
				}

			case bool:
				if targetType == reflect.TypeOf(sql.NullInt64{}) {
					var i int64 = 0
					if v {
						i = 1
					}
					paramField.Set(reflect.ValueOf(sql.NullInt64{Int64: i, Valid: true}))
				}
			}
		}
	}
}
