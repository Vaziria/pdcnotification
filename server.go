package pdcnotification

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vaziria/pdcnotification/repo"
	"github.com/vaziria/pdcnotification/services"
)

type ActionNotif string

const (
	SendAction     ActionNotif = "send_action"
	AddTokenAction             = "add_token"
)

type Payload struct {
	Action  ActionNotif `json:"action"`
	Email   string      `json:"email"`
	Message string      `json:"message"`
	Tokens  []string    `json:"tokens"`
}

type ResponseErrorCode string

const (
	UserNotFound   ResponseErrorCode = "user_not_found"
	UserExist                        = "user_exist"
	Success                          = "success"
	SendNotifError                   = "send_notif_error"
	AddTokenError                    = "add_token_error"
)

type Response struct {
	Errcode ResponseErrorCode `json:"errcode"`
	Message string            `json:"message"`
}

func (res Response) ReturnData(w io.Writer) {
	response, _ := json.Marshal(&res)
	fmt.Fprint(w, string(response))
}

func corsEnabled(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func Notification(w http.ResponseWriter, r *http.Request) {
	corsEnabled(w, r)

	payload := Payload{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	if payload.Action == SendAction {
		user, err := repo.FindByEmail(payload.Email)

		if err != nil {
			res := Response{
				UserNotFound,
				"user tidak ditemukan",
			}

			res.ReturnData(w)
			return
		}

		srverr := services.SendNotification(user, payload.Message, "")

		if srverr != nil {
			res := Response{
				SendNotifError,
				"notifikasi error",
			}

			res.ReturnData(w)
			return
		}

		res := Response{
			Success,
			"notifikasi berhasil dikirim",
		}

		res.ReturnData(w)
		return

	} else if payload.Action == AddTokenAction {
		_, errepo := repo.AddToken(payload.Email, payload.Tokens)

		if errepo != nil {
			res := Response{
				AddTokenError,
				"token error",
			}
			res.ReturnData(w)
			return
		}

		res := Response{
			Success,
			"token added",
		}
		res.ReturnData(w)
		return
	} else {
		fmt.Fprint(w, "no action not implemented")
	}

}
