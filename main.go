package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret statis
var jwtSecret = []byte("ae7b72f896f54e649403ec1a53e6f1a4f7c9b334e1224b3bc9d2d5f2c0f739f1")

type UserData struct {
	Email      string `json:"email"`
	UserID     string `json:"user_id"`
	PrivateKey string `json:"private_key"`
}

type MyClaims struct {
	Data UserData `json:"data"`
	jwt.RegisteredClaims
}

// Generate JWT token
func GenerateJWT(data UserData) (string, error) {
	claims := MyClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "myapp",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)), // 10 menit
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(jwtSecret)
	return token.SignedString(jwtSecret)
}

// Verify dan cek expired
func VerifyJWT(tokenStr string) (*UserData, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// ✅ Validasi algoritma HARUS HS256 (HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token invalid: %v", err)
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token invalid or expired")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return &claims.Data, nil
}


func main() {
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		user := UserData{
			Email:      "vario_n_ramadhan_x@telkomsel.co.id",
			UserID:     "fb3fa8a8-5768-11ee-af6c-005056978071",
			PrivateKey: "ae7b72f896f54e649403ec1a53e6f1a4",
		}
		token, err := GenerateJWT(user)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, token)
	})

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.URL.Query().Get("token")
		if tokenStr == "" {
			http.Error(w, "Token missing in query param", http.StatusBadRequest)
			return
		}

		data, err := VerifyJWT(tokenStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "✅ Token Valid\nEmail: %s\nUserID: %s\nPrivateKey: %s\n",
			data.Email, data.UserID, data.PrivateKey)
	})

	fmt.Println("✅ Server running on :8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
