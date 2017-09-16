package main

import (
	"database/sql"
	"io"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error

	if db, err = sql.Open("postgres", ""); err != nil {
		panic(err)
	}
}

func addSolution(userID int, hole, lang, code string) {
	// Update the code if it's the same characters or less, but only update
	// the submitted time if the solution is shorter. This avoids a user
	// moving down the leaderboard by matching their personal best.
	if _, err := db.Exec(`
	    INSERT INTO solutions
	         VALUES (NOW(), $1, $2, $3, $4)
	    ON CONFLICT ON CONSTRAINT solutions_pkey
	  DO UPDATE SET submitted = CASE
	                    WHEN LENGTH($4) < LENGTH(solutions.code)
	                    THEN NOW()
	                    ELSE solutions.submitted
	                END,
	                code = CASE
	                    WHEN LENGTH($4) > LENGTH(solutions.code)
	                    THEN solutions.code
	                    ELSE $4
	                END
	`, userID, hole, lang, code); err != nil {
		panic(err)
	}
}

func getUser(login string) bool {
	var one int

	if err := db.QueryRow(
		"SELECT 1 FROM users WHERE login = $1", login,
	).Scan(&one); err != nil && err != sql.ErrNoRows {
		panic(err)
	} else {
		return err != sql.ErrNoRows
	}
}

func getUserSolutions(userID int, hole string) map[string]string {
	rows, err := db.Query(
		`SELECT code, lang
		   FROM solutions
		  WHERE user_id = $1 AND hole = $2`,
		userID, hole,
	)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	solutions := make(map[string]string)

	for rows.Next() {
		var code, lang string

		if err := rows.Scan(&code, &lang); err != nil {
			panic(err)
		}

		solutions[lang] = code
	}

	return solutions
}

func addUser(id int, login string) {
	if _, err := db.Exec(
		`INSERT INTO users VALUES($1, $2)
		 ON CONFLICT(id) DO UPDATE SET login = $2`,
		id, login,
	); err != nil {
		panic(err)
	}
}

func printLeaderboards(w io.WriteCloser) {
	rows, err := db.Query(
		`SELECT hole,
		        CONCAT(
		            '<tr><td>',
		            place,
		            '<td class=',
		            lang,
		            '>',
		            strokes,
		            '<td><img src="//avatars.githubusercontent.com/',
		            login,
		            '?size=26"><a href="u/',
		            login,
		            '">',
		            login,
		            '</a>'
		        )
		   FROM leaderboard
		   JOIN users ON user_id = id
		  WHERE place < 4`,
	)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	prevHole := ""
	w.Write([]byte("<article id=home>"))

	for rows.Next() {
		var hole, row string

		if err := rows.Scan(&hole, &row); err != nil {
			panic(err)
		}

		if hole != prevHole {
			if prevHole != "" {
				w.Write([]byte("</table></div>"))
			}
			w.Write([]byte(intros[hole]))
			prevHole = hole
		}

		w.Write([]byte(row))
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
