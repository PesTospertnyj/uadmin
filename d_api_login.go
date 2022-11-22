package uadmin

import "net/http"

func dAPILoginHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	if s != nil {
		Logout(r)
	}

	// Get request variables
	username := r.FormValue("username")
	password := r.FormValue("password")
	otp := r.FormValue("otp")
	session := r.FormValue("session")

	optRequired := false
	if otp != "" {
		// Check if there is username and password or a session key
		if session != "" {
			w.WriteHeader(http.StatusAccepted)
			s = Login2FAKey(r, session, otp)
		} else {
			s = Login2FA(r, username, password, otp)
		}
	} else {
		s, optRequired = Login(r, username, password)
	}

	if optRequired {
		w.WriteHeader(http.StatusUnauthorized)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "OTP Required",
			"session": s.Key,
		})
		return
	}

	if s == nil {
		w.WriteHeader(http.StatusForbidden)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Invalid username or password",
		})
		return
	}

	jwt := SetSessionCookie(w, r, s)
	ReturnJSON(w, r, map[string]interface{}{
		"status":  "ok",
		"session": s.Key,
		"jwt":     jwt,
		"user": map[string]interface{}{
			"username":   s.User.Username,
			"first_name": s.User.FirstName,
			"last_name":  s.User.LastName,
			"group_id":   s.User.UserGroupID,
			"admin":      s.User.Admin,
		},
	})
}
