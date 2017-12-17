package interfaces

import "github.com/dtrinh100/Music-Playlist/src/api/usecases"

type DBRepo struct {
	dbHandlers map[string]DBHandler
	dbHandler  DBHandler
}

func (repo *DBRepo) getNextSequence(name string) (int, error) {
	update := usecases.M{"$inc": usecases.M{"seq": 1}}
	query := usecases.M{"_id": name}

	results := usecases.M{}

	_, modifyErr := repo.dbHandlers["DBCounterRepo"].FindAndModify(query, update, &results)

	return results["seq"].(int), modifyErr
}