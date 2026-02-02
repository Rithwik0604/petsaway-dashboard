package clients

import (
	"database/sql"
	"petsaway/internal/database"
	"reflect"
	"time"

	"github.com/google/uuid"
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
	var params database.UpdateClientByIdParams

	MapDTOToParams(&updates, &params)
	params.ID = id

	client, err := database.DB.UpdateClientById(database.Qctx, params)
	if err != nil {
		return database.Client{}, err
	}

	return client, nil
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
				// 1. Check if target is NullTime
				if targetType == reflect.TypeOf(sql.NullTime{}) {
					if v == "" {
						continue
					} // Skip empty date strings
					t, err := time.Parse("2006-01-02", v) // Standard HTML date format
					if err == nil {
						paramField.Set(reflect.ValueOf(sql.NullTime{Time: t, Valid: true}))
					}
				} else if targetType == reflect.TypeOf(sql.NullString{}) {
					paramField.Set(reflect.ValueOf(sql.NullString{String: v, Valid: true}))
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
