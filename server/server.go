package server

import (
	"net/http"
	"time"

	"github.com/timmilesdw/slack-to-telegram/slack"
	"github.com/timmilesdw/slack-to-telegram/telegram"
	"github.com/timmilesdw/slack-to-telegram/template"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	address  string
	router   *mux.Router
	template *template.Template
	telegram *telegram.TelegramBot
	server   *http.Server
}

func NewHttpServer(
	address string,
	template *template.Template,
	telegram *telegram.TelegramBot,
) *HttpServer {
	r := mux.NewRouter()

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		// Handler:      r,
	}

	return &HttpServer{
		router:   r,
		address:  address,
		template: template,
		telegram: telegram,
		server:   srv,
	}
}

func (h *HttpServer) SetupRoutes() {
	r := h.router
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		form := r.Form

		payload, ok := form["payload"]
		if !ok {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Bad request!"))
			return
		}
		sm, err := slack.Parse(payload[0])

		if err != nil {
			log.Errorf("can't parse slack message body, err: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Bad request!"))
			return
		}

		str, err := h.template.RenderMessage(sm)
		if err != nil {
			log.Errorf("can't template slack message, err: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Bad request!"))
			return
		}
		channel := ""
		if sm.Channel != "" {
			channel = sm.Channel[1:]
		}
		if err := h.telegram.SendMesage(str, channel); err != nil {
			log.Errorf("can't send telegram message, err: %s", err)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Internal server error!"))
			return
		}

		rw.Write([]byte("Hello!"))
	})

	h.server.Handler = r
}

func (h *HttpServer) Listen() error {
	return h.server.ListenAndServe()
}
