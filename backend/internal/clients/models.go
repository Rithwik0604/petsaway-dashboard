package clients

type ClientDTO struct {
	Id                      *string  `json:"id"`
	ClientName              *string  `json:"client_name"`
	ClientPhone             *string  `json:"client_phone"`
	ImportExport            *string  `json:"import_export"`
	ImportFee               *float64 `json:"import_fee"`
	ExportFee               *float64 `json:"export_fee"`
	AfterHoursCharges       *float64 `json:"after_hours_charges"`
	PetName                 *string  `json:"pet_name"`
	Species                 *string  `json:"species"`
	Gender                  *string  `json:"gender"`
	Breed                   *string  `json:"breed"`
	DateOfBirth             *string  `json:"date_of_birth"`
	MicrochipNumber         *string  `json:"microchip_number"`
	Titre                   *string  `json:"titre"`
	LastRabiesDate          *string  `json:"last_rabies_date"`
	RabiesValidity          *string  `json:"rabies_validity"`
	DocumentationStatus     *bool    `json:"documentation_status"`
	RabiesVaccinationValid  *bool    `json:"rabies_vaccination_valid"`
	OtherVaccinesCompleted  *string  `json:"other_vaccines_completed"`
	HealthCertificateIssues *bool    `json:"health_certificate_issues"`
	ExportPermitApproved    *bool    `json:"export_permit_approved"`
	ImportPermitApproved    *bool    `json:"import_permit_approved"`
	AirlineApprovalReceived *bool    `json:"airline_approval_received"`
	CustomsClearanceDone    *bool    `json:"customs_clearance_done"`
	OriginCountry           *string  `json:"origin_country"`
	DestinationCountry      *string  `json:"destination_country"`
	ForwarderCharges        *float64 `json:"forwarder_charges"`
	Departure               *string  `json:"departure"`
	Airline                 *string  `json:"airline"`
	AirlineCharges          *float64 `json:"airline_charges"`
	CrateCost               *float64 `json:"crate_cost"`
	FlightNumber            *string  `json:"flight_number"`
	TypeOfTravel            *string  `json:"type_of_travel"`
	Etd                     *string  `json:"etd"`
	Eta                     *string  `json:"eta"`
	QuotedAmount            *float64 `json:"quoted_amount"`
	TotalCost               *float64 `json:"total_cost"`
	Profit                  *float64 `json:"profit"`
	AdvancedReceived        *float64 `json:"advanced_received"`
	BalancePending          *float64 `json:"balance_pending"`
	PaymentStatus           *string  `json:"payment_status"`
	Remarks                 *string  `json:"remarks"`
}
