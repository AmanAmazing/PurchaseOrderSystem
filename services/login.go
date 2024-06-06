package services

import (
	"PurchaseOrderSystem/models"
	"PurchaseOrderSystem/utils"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func PostLogin(username, password string) (string, string, error) {
	var user models.PostLogin
	err := db.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("invalid username or password")
		}
		return "", "", fmt.Errorf("failed to retrieve user: %v", err)
	}
	match := utils.CheckPasswordMatch(user.Password, password)
	if !match {
		return "", "", fmt.Errorf("invalid password")
	}
	// Query the user_roles table to get the user's role
	// FIX: Could do the following two queries in one I think
	var roleID int
	err = db.QueryRow(context.Background(), "SELECT role_id FROM user_roles WHERE user_id = $1", user.ID).Scan(&roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("user role not found")
		}
		return "", "", fmt.Errorf("failed to retrieve user role: %v", err)
	}

	// Query the roles table to get the role name
	err = db.QueryRow(context.Background(), "SELECT name FROM roles WHERE id = $1", roleID).Scan(&user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("role not found")
		}
		return "", "", fmt.Errorf("failed to retrieve role name: %v", err)
	}

	// Getting the user's department
	var departmentID int
	err = db.QueryRow(context.Background(), "SELECT department_id FROM user_departments WHERE user_id = $1", user.ID).Scan(&departmentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("user department not found")
		}
		return "", "", fmt.Errorf("failed to retrieve user department: %v", err)
	}
	// Query the departments table to get the department name
	err = db.QueryRow(context.Background(), "SELECT name FROM departments WHERE id = $1", departmentID).Scan(&user.Department)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", fmt.Errorf("department not found")
		}
		return "", "", fmt.Errorf("failed to retrieve department name: %v", err)
	}

	// Generate a JWT token
	claims := map[string]interface{}{
		"userId":     user.ID,
		"role":       user.Role,
		"department": user.Department,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Sign the token with a secret key
	tokenString, err := utils.GenerateJWT(claims)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate token: %v", err)
	}
	return tokenString, user.Role, nil

}
