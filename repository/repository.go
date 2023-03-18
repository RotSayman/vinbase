package repository

import (
	"fmt"
	"strings"

	"github.com/RotSayman/vinbase/models"
	"github.com/RotSayman/vinbase/storage"
	"gorm.io/gorm"
)

var r *Repository

type Repository struct {
	DB *gorm.DB
}

func GetRepo() *Repository {
	r = &Repository{
		DB: storage.Init(),
	}
	return r
}

func (r *Repository) GetAllVins() []models.Vin {
	var vins []models.Vin
	r.DB.Select("name", "vin_id").Find(&vins)
	return vins
}

func (r *Repository) AddVin(name, vin string) {
	r.DB.Create(&models.Vin{
		Name:  name,
		VinID: vin,
	})
}

func (r *Repository) GetVin(id string) models.Vin {
	var vin models.Vin
	r.DB.First(&vin, "vin_id", id)
	return vin
}

func (r *Repository) UpdateVin(id, name, vin string) {
	v := r.GetVin(id)
	r.DB.Model(&v).Updates(models.Vin{Name: name, VinID: vin})
}

func (r *Repository) DeleteVin(id string) {
	var vin models.Vin
	r.DB.First(&vin, "vin_id", id)
	r.DB.Delete(&vin)

}

func (r *Repository) FindByName(name string) []models.Vin {
	var vins []models.Vin
	fmt.Println(strings.ToLower("%" + name + "%"))
	r.DB.Select("name, vin_id").Where("name LIKE ? collate nocase", "%"+name+"%").Find(&vins)
	return vins
}
