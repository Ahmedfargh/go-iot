package seeders

import (
	models "go-dashboard/internals/Models"
	crud_services "go-dashboard/internals/Services/Crud"
	"log"
)

func SeedSectors(sectorService *crud_services.SectorService) {
	sectors := []models.Sector{
		{Name: "Alpha", Description: "Primary manufacturing sector"},
		{Name: "Beta", Description: "Secondary storage and logistics sector"},
		{Name: "Gamma", Description: "Research and development sector"},
	}

	for _, s := range sectors {
		_, err := sectorService.Create(&s)
		if err != nil {
			log.Printf("Could not seed sector %s: %v\n", s.Name, err)
		} else {
			log.Printf("Sector %s seeded successfully.\n", s.Name)
		}
	}
}
