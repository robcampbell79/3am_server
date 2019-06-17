package getMovies

import(
	//"encoding/json"
	//"os"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Movies struct {
	MTitle string
	MPath string
	MGenre string
  }

func GetMovies(genre string) []Movies {
	db, err := sql.Open("mysql", "spooky:toor@/nightmare")
	if err != nil {
		fmt.Println("We got a nil here")
	} else {
		fmt.Println("We good")
	}

	defer db.Close()

	 selectAll, err := db.Query("SELECT title, path, genre FROM movies WHERE genre=?", genre)
	//insert, err := db.Query("INSERT INTO movies(title, path, supernatural_elements, genre) VALUES('Nosferatu', 'C:/Users/Robert/Videos/horror/Nosferatu_1922_Symphony_of_Horror_512kb', 'vampires', 'horror')")
	if err != nil {
		fmt.Println("I don't know what happened")
		panic(err.Error())
	}

	defer selectAll.Close()

	result := []Movies{}

	for selectAll.Next() {
		movie := Movies{}
		err := selectAll.Scan(&movie.MTitle, &movie.MPath, &movie.MGenre)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, movie)
	}

	// json.NewEncoder(os.Stdout).Encode(result)
	//json.Marshal(result)

	return result

}