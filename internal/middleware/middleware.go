package middleware

// import (
// 	"net/http"
//
// 	"github.com/kelbwah/swiftlio/internal/auth"
// 	"github.com/kelbwah/swiftlio/internal/database"
// 	"github.com/kelbwah/swiftlio/utils"
// )
//
// type authedHandler func(http.ResponseWriter, *http.Request, database.User)
//
// func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		apiKey, err := auth.GetAPIKey(r.Header)
// 		if err != nil {
// 			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
// 			return
// 		}
//
// 		foundUser, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
// 		if err != nil {
// 			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 			return
// 		}
//
// 		handler(w, r, foundUser)
// 	}
// }
