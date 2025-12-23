package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GoogleTokenInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	HD            string `json:"hd"`
}

func VerifyGoogleToken(ctx context.Context, idToken string) (*GoogleTokenInfo, error) {
	url := fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", idToken)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error verificando token con Google: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("token inválido: %s", string(body))
	}

	var tokenInfo GoogleTokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta de Google: %w", err)
	}

	if !tokenInfo.EmailVerified {
		return nil, fmt.Errorf("el email no está verificado en Google")
	}

	allowedDomain := os.Getenv("GOOGLE_ALLOWED_DOMAIN") // ej: "unal.edu.co"
	if allowedDomain != "" && tokenInfo.HD != allowedDomain {
		return nil, fmt.Errorf("el dominio %s no está permitido (se requiere %s)", tokenInfo.HD, allowedDomain)
	}

	return &tokenInfo, nil
}
