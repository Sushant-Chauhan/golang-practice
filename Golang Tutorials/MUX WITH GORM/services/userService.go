package services

import (
	"BankingApp/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser creates a new user
func CreateUser(username, password, firstName, lastName string, isAdmin bool) error {
	tx := models.DB.Begin() // Start a transaction

	// Check if the user already exists
	var existingUser models.User
	if err := tx.Where("username = ?", username).First(&existingUser).Error; err == nil {
		tx.Rollback() // Rollback the transaction if user already exists
		return errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Create new user object
	user := models.User{
		Username:  username,
		Password:  string(hashedPassword),
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   isAdmin,
		IsActive:  true,
	}

	// Insert into the database
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit() // Commit the transaction
	return nil
}

// GetAllUsers returns all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUserByID updates a user by ID
func UpdateUserByID(id int, firstName, lastName, username, password string, isAdmin bool) error {
	tx := models.DB.Begin() // Start a transaction

	var user models.User
	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		return errors.New("user not found")
	}

	// Update the fields if provided
	if firstName != "" {
		user.FirstName = firstName
	}
	if lastName != "" {
		user.LastName = lastName
	}
	if username != "" && username != user.Username {
		// Check if the new username already exists
		var existingUser models.User
		if err := tx.Where("username = ?", username).First(&existingUser).Error; err == nil {
			tx.Rollback()
			return errors.New("username already exists")
		}
		user.Username = username
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			tx.Rollback()
			return err
		}
		user.Password = string(hashedPassword)
	}
	user.IsAdmin = isAdmin

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit() // Commit the transaction
	return nil
}

// DeleteUserByID deletes a user by ID
func DeleteUserByID(id int) error {
	tx := models.DB.Begin() // Start a transaction

	if err := tx.Delete(&models.User{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit() // Commit the transaction
	return nil
}
