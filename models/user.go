package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id				int 		`form:"id" json:"id" validate:"required"` 
	Name			string	`form:"name" json:"name" validate:"required"` 
	Email			string	`form:"email" json:"email" validate:"required"` 
	Username	string	`form:"username" json:"username" validate:"required"` 
	Password 	string 	`form:"password" json:"password" validate:"required"`
}

//CRUD
// POST /user/create
func CreateUser(db *gorm.DB, newUser *User) (err error) {
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

// func ReadProducts(db *gorm.DB, products *[]Product)(err error)  {
// 	err = db.Find(products).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func ReadUserByUsername(db *gorm.DB, user *User, username string)(err error)  {
	err = db.Where("username=?", username).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

// func UpdateProduct(db *gorm.DB, product *Product)(err error)  {
// 	db.Save(product)
// 	return nil
// }

// func DeleteProductById(db *gorm.DB, product *Product, id int)(err error)  {
// 	db.Where("id=?", id).Delete(product)
// 	return nil
// }