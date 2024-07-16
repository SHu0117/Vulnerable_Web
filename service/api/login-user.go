package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SHu0117/WASA-Photo/service/api/reqcontext"
	"github.com/SHu0117/WASA-Photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.ExistUsername(user.Username, user.Password)
	if errors.Is(err, database.ErrDataDoesNotExist) {
		dbuser, err := rt.db.NewUser(user.UserToDatabase())
		if err != nil {
			ctx.Logger.WithError(err).Error("can't create the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.UserFromDatabase(dbuser)

		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(user)

	} else if errors.Is(err, database.ErrWrongPassword) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("qualcosa non va")
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else {
		dbuser, err := rt.db.GetUserID(user.Username)
		if err != nil {
			ctx.Logger.WithError(err).Error("qualcosa non va")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.UserFromDatabase(dbuser)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error while encoding data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}
