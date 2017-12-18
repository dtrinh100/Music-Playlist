package infrastructure

import (
	"github.com/dtrinh100/Music-Playlist/src/api/interfaces"
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"github.com/dtrinh100/Music-Playlist/src/api/domain"
	"gopkg.in/mgo.v2"
	"time"
)

func newMongoHandler(session *mgo.Session, dbName, dbTableName string) *MongoHandler {
	mongoHandler := new(MongoHandler)
	mongoHandler.session = session
	mongoHandler.dbName = dbName
	mongoHandler.dbTableName = dbTableName

	return mongoHandler
}

func newMongoSession(addrs, un, pw string) *mgo.Session {
	session, sessionErr := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{addrs},
		Username: un,
		Password: pw,
		Timeout:  60 * time.Second,
	})

	if sessionErr != nil {
		panic(sessionErr)
		return nil
	}
	return session
}

func GetAndInitDBHandlersForDBName(dbName string) map[string]interfaces.DBHandler {
	session := newMongoSession("MPDatabase", "", "")
	dbUserHandler := newMongoHandler(session, dbName, "usertable")
	dbSongHandler := newMongoHandler(session, dbName, "songtable")
	dbCounterHandler := newMongoHandler(session, dbName, "countertable")

	songSeq := interfaces.Counter{"songid", 0}
	userSeq := interfaces.Counter{"userid", 0}
	dbCounterHandler.Create(songSeq)
	dbCounterHandler.Create(userSeq)

	if ensureErr := dbUserHandler.EnsureIndex("username"); ensureErr != nil {
		panic(ensureErr)
	}

	handlers := make(map[string]interfaces.DBHandler)
	handlers["DBUserRepo"] = dbUserHandler
	handlers["DBSongRepo"] = dbSongHandler
	handlers["DBCounterRepo"] = dbCounterHandler

	return handlers
}

func GetAndInitUserInteractor(logger *Logger, handlers map[string]interfaces.DBHandler) *usecases.UserInteractor {
	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDBUserRepo(handlers)
	userInteractor.Logger = logger

	return userInteractor
}

func GetAndInitSongInteractor(logger *Logger, handlers map[string]interfaces.DBHandler) *usecases.SongInteractor {
	songInteractor := new(usecases.SongInteractor)
	songInteractor.SongRepository = interfaces.NewDBSongRepo(handlers)
	songInteractor.Logger = logger

	initSongs(songInteractor)

	return songInteractor
}

func initSongs(songInteractor *usecases.SongInteractor) {
	// If there aren't any songs in the database, add some.
	if songs, _ := songInteractor.All(); len(songs) == 0 {
		songs := []domain.Song{
			domain.Song{
				Name: "Take On Me",
				Artist: "A-Ha",
				Description: "\"Take On Me\" is a song by Norwegian synthpop band A-ha. The self-composed original version was produced by Tony Mansfield, and remixed by John Ratcliff. The second version was produced by Alan Tarney for the group's debut studio album Hunting High and Low (1985). The song combines synthpop with a varied instrumentation that includes acoustic guitars, keyboards and drums.",
				AudioPath: "http://dora-robo.com/muzyka/70's-80's-90's%20/A-Ha%20-%20Take%20On%20Me%20(Extended%20Version).mp3",
				ImgURL: "https://upload.wikimedia.org/wikipedia/en/thumb/d/d5/A-ha_take_on_me-1stcover.jpg/220px-A-ha_take_on_me-1stcover.jpg",
				AltText: "Album Cover Art",
			},
			domain.Song{
				Name: "Hotel California",
				Artist: "Eagles",
				Description: "\"Hotel California\" is the title track from the Eagles' album of the same name and was released as a single in February 1977.[2] Writing credits for the song are shared by Don Felder (music), Don Henley, and Glenn Frey (lyrics). The Eagles' original recording of the song features Henley singing the lead vocals and concludes with an extended section of electric guitar interplay between Felder and Joe Walsh.",
				AudioPath: "http://dora-robo.com/muzyka/70's-80's-90's%20/Hotel%20California.mp3",
				ImgURL: "https://upload.wikimedia.org/wikipedia/en/thumb/3/33/Eagles-usa-hotel-california-asylum.jpg/220px-Eagles-usa-hotel-california-asylum.jpg",
				AltText: "Album Cover Art",
			},
			domain.Song{
				Name: "Ice, Ice Baby",
				Artist: "Vanilla Ice",
				Description: "\"Ice Ice Baby\" is a hip hop song written by American rapper Vanilla Ice and DJ Earthquake. It was based on the bassline of \"Under Pressure\" by Queen and David Bowie, who did not initially receive songwriting credit or royalties until after it had become a hit. Originally released on Vanilla Ice's 1989 debut album Hooked and later on his 1990 national debut To the Extreme, it is his best known song. It has appeared in remixed form on Platinum Underground and Vanilla Ice Is Back! A live version appears on the album Extremely Live, while a nu metal version appears on the album Hard to Swallow, under the title \"Too Cold\".",
				AudioPath: "http://dora-robo.com/muzyka/70's-80's-90's%20/Ice,%20Ice%20Baby.mp3",
				ImgURL: "https://upload.wikimedia.org/wikipedia/en/thumb/3/3f/Vanilla_Ice_-_Ice_Ice_Baby.jpg/220px-Vanilla_Ice_-_Ice_Ice_Baby.jpg",
				AltText: "Album Cover Art",
			},
			domain.Song{
				Name: "I Love Rock 'N Roll",
				Artist: "Joan Jett",
				Description: "\"I Love Rock 'n' Roll\" is a rock song written in 1975 by Alan Merrill of the Arrows, who recorded the first released version.[1] The song was later made famous by Joan Jett & the Blackhearts in 1982.[2] Alan Merrill still plays the song live in Europe, Japan and most often in his home town New York City.",
				AudioPath: "http://dora-robo.com/muzyka/70's-80's-90's%20/Joan%20Jett%20-%20I%20Love%20Rock%20'N%20Roll%20%5bApril%201982%5d.mp3",
				ImgURL: "https://upload.wikimedia.org/wikipedia/en/thumb/8/84/Arrows_I_Love_Rock_n_Roll.jpg/220px-Arrows_I_Love_Rock_n_Roll.jpg",
				AltText: "Album Cover Art",
			},
			domain.Song{
				Name: "U Can't Touch This",
				Artist: "MC Hammer",
				Description: "\"U Can't Touch This\" is a song co-written, produced and performed by MC Hammer from his 1990 album Please Hammer, Don't Hurt 'Em. The track is considered to be Hammer's signature song and is his most successful single.",
				AudioPath: "http://dora-robo.com/muzyka/70's-80's-90's%20/U%20Can't%20Touch%20This.mp3",
				ImgURL: "https://upload.wikimedia.org/wikipedia/en/thumb/d/d0/Hammer_Touch.jpg/220px-Hammer_Touch.jpg",
				AltText: "Album Cover Art",
			},
			domain.Song{
				Name: "Video Killed The Radio Star",
				Artist: "Bruce Woolley",
				Description: "\"Video Killed the Radio Star\" is a song written by Trevor Horn, Geoff Downes and Bruce Woolley in 1978. It was first recorded by Bruce Woolley and The Camera Club (with Thomas Dolby on keyboards) for their album English Garden, and later by British group the Buggles, consisting of Horn and Downes. The track was recorded and mixed in 1979, released as their debut single on 7 September 1979 by Island Records, and included on their first album The Age of Plastic. The backing track was recorded at Virgin's Town House in West London, and mixing and vocal recording would later take place at Sarm East Studios.",
				AudioPath: "http://dora-robo.com/muzyka/70's-80's-90's%20/Video%20Killed%20The%20Radio%20Star.mp3",
				ImgURL: "https://upload.wikimedia.org/wikipedia/en/thumb/a/a8/Video_Killed_the_Radio_Star_Bruce_Woolley.jpg/220px-Video_Killed_the_Radio_Star_Bruce_Woolley.jpg",
				AltText: "Album Cover Art",
			},
			domain.Song{
				Name: "Two Princes",
				Artist: "Spin Doctors",
				Description: "\"Two Princes\" is a song by the New York City-based band Spin Doctors. Released as a single in 1993, it reached No. 7 in the United States, No. 2 in Canada, and No. 3 in the United Kingdom. It was the band's highest-charting single internationally. It earned them a Grammy Award nomination for Best Rock Performance by a Duo or Group.[2]",
				AudioPath: "http://dora-robo.com/muzyka/70's-80's-90's%20/Two%20Princes.mp3",
				ImgURL: "https://upload.wikimedia.org/wikipedia/en/thumb/d/dd/TwoPrinces.jpg/220px-TwoPrinces.jpg",
				AltText: "Album Cover Art",
			},
		};

		for _, song := range songs {
			songInteractor.Create(&song)
		}
	}
}