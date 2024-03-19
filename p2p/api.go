package p2p

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHandlerFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			err = SendJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			fmt.Println(err)
		}
	}
}

func SendJSON(w http.ResponseWriter, status int, data any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

type APIServer struct {
	listenAddr string
	game       *Game
}

func NewAPIServer(addr string, game *Game) *APIServer {
	return &APIServer{
		listenAddr: addr,
		game:       game,
	}
}

func (s *APIServer) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/ready", makeHandlerFunc(s.handlePlayerReady))
	r.HandleFunc("/fold", makeHandlerFunc(s.handlePlayerFold))
	r.HandleFunc("/check", makeHandlerFunc(s.handlePlayerCheck))
	r.HandleFunc("/bet/{value}", makeHandlerFunc(s.handlePlayerBet))

	http.ListenAndServe(s.listenAddr, r)
}

func (s *APIServer) handlePlayerReady(w http.ResponseWriter, r *http.Request) error {
	s.game.SetReady()
	return SendJSON(w, http.StatusOK, map[string]string{
		"message": "player is ready",
	})
}

func (s *APIServer) handlePlayerFold(w http.ResponseWriter, r *http.Request) error {
	if err := s.game.TakeAction(PlayerActionFold, 0); err != nil {
		return SendJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return SendJSON(w, http.StatusOK, map[string]string{
		"message": "player is folding",
	})
}

func (s *APIServer) handlePlayerCheck(w http.ResponseWriter, r *http.Request) error {
	if err := s.game.TakeAction(PlayerActionCheck, 0); err != nil {
		return SendJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return SendJSON(w, http.StatusOK, map[string]string{
		"message": "player is checking",
	})
}

func (s *APIServer) handlePlayerBet(w http.ResponseWriter, r *http.Request) error {
	valueStr := mux.Vars(r)["value"]
	value, err := strconv.Atoi(valueStr)
	if err = s.game.TakeAction(PlayerActionBet, value); err != nil {
		return SendJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return SendJSON(w, http.StatusOK, map[string]string{
		"message": fmt.Sprintf("player's bet is %d", value),
	})
}
