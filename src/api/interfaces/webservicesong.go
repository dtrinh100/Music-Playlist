package interfaces

import (
	"net/http"
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"encoding/json"
	"github.com/dtrinh100/Music-Playlist/src/api/domain"
	gmux "github.com/gorilla/mux"
	"io"
	"strconv"
)

func (webhandler *WebserviceHandler) Songs(rw http.ResponseWriter, req *http.Request) {
	songs, allErr := webhandler.SongInteractor.All()

	if allErr != nil {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	webhandler.Responder.Success(rw, usecases.M{"songs": songs})
}

func (webhandler *WebserviceHandler) Song(rw http.ResponseWriter, req *http.Request) {
	songid, _ := strconv.Atoi(gmux.Vars(req)["id"])

	song, songErr := webhandler.SongInteractor.GetByID(songid)

	if songErr != nil {
		switch songErr.Status() {
		case usecases.UserFaultErr:
			webhandler.Responder.BadRequest(rw, songErr)
		default:
			webhandler.Responder.InternalServerError(rw)
		}
		return
	}

	webhandler.Responder.Success(rw, usecases.M{"song": song})
}

func (webhandler *WebserviceHandler) CreateSong(rw http.ResponseWriter, req *http.Request) {
	var newSong domain.Song

	if decodeErr := json.NewDecoder(req.Body).Decode(&newSong); decodeErr != nil {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	if createErr := webhandler.SongInteractor.Create(&newSong); createErr != nil {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	webhandler.Responder.NoContent(rw)
}

func (webhandler *WebserviceHandler) UpdateSong(rw http.ResponseWriter, req *http.Request) {
	songid, _ := strconv.Atoi(gmux.Vars(req)["id"])

	updateDict := usecases.M{}

	if decErr := json.NewDecoder(req.Body).Decode(&updateDict); decErr != nil && decErr != io.EOF {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	badRequestFn := func(msg string) {
		webhandler.Responder.BadRequest(rw, &usecases.FaultError{
			FaultEntity: usecases.UserFaultErr,
			Message:     msg})
	}

	if len(updateDict) == 0 {
		badRequestFn("empty updates")
		return
	}

	for key, value := range updateDict {
		val := value.(string)
		switch key {
		case "name":
			webhandler.SongInteractor.UpdateName(songid, val)
		case "country":
			webhandler.SongInteractor.UpdateCountry(songid, val)
		case "state":
			webhandler.SongInteractor.UpdateState(songid, val)
		default:
			badRequestFn("cannot update - " + key)
			return
		}
	}

	webhandler.Responder.NoContent(rw)
}

func (webhandler *WebserviceHandler) DeleteSong(rw http.ResponseWriter, req *http.Request) {
	songid, _ := strconv.Atoi(gmux.Vars(req)["id"])

	if deleteErr := webhandler.SongInteractor.Delete(songid); deleteErr != nil {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	webhandler.Responder.NoContent(rw)
}
