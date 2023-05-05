package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Entity represents an entity in the STEP file
type Person struct {
	ID    string `json:"id"`
	X     string `json:"x"`
	Y     string `json:"y"`
	Z     string `json:"z"`
	Types string `json:"type"`
}
type Profile struct {
	Name        string   `json:"name"`
	Note        string   `json:"note"`
	Customer    string   `json:"customer"`
	OrderNumber string   `json:"orderNumber"`
	Points      []Person `json:"points"`
}

func main() {

	db, err := sql.Open("sqlite3", "./pointexss.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cartesian_point (id INTEGER PRIMARY KEY, x TEXT,y TEXT,z TEXT,type TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("02052023.stp")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the lines of the file into a slice
	scanner := bufio.NewScanner(file)

	var line []string
	arr := []string{}
	cylindricalarr := []string{}
	oriented_edge := []string{}

	for scanner.Scan() {

		line := scanner.Text()
		//println(line)

		if strings.Contains(line, "CIRCLE") {
			fields := strings.Split(line, "CIRCLE('',#")
			if len(fields) > 1 {
				fieldx := strings.Split(fields[1], ",")
				fieldy := strings.Split(fieldx[1], ")")
				arr = append(arr, fieldx[0], fieldy[0])

				stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}
				_, err = stmt.Exec(fieldx[0], fieldy[0], 0, "CIRCLE")
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		if strings.Contains(line, "CYLINDRICAL_SURFACE") {
			fieldss := strings.Split(line, "CYLINDRICAL_SURFACE('',#")
			if len(fieldss) > 1 {
				fieldxx := strings.Split(fieldss[1], ",")
				fieldyy := strings.Split(fieldxx[1], ")")
				cylindricalarr = append(cylindricalarr, fieldxx[0], fieldyy[0])

				stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}
				_, err = stmt.Exec(fieldxx[0], fieldyy[0], 0, "CYLINDRICAL_SURFACE")
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		if strings.Contains(line, "ORIENTED_EDGE") {
			fieldsss := strings.Split(line, "'',*,*,#")
			if len(fieldsss) > 1 {
				fieldxxo := strings.Split(fieldsss[1], ",")
				fieldyyo := strings.Split(fieldxxo[1], ")")
				oriented_edge = append(oriented_edge, fieldxxo[0], fieldyyo[0])

				stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}
				_, err = stmt.Exec(fieldxxo[0], fieldyyo[0], 0, "ORIENTED_EDGE")
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		if strings.Contains(line, "EDGE_CURVE") {
			fieldssss := strings.Split(line, "'',")

			if len(fieldssss) > 1 {
				fieldxxoe := strings.Split(fieldssss[1], ",")
				fieldyyoe := strings.Split(fieldxxoe[1], ",")

				fieldyzoe := strings.Split(fieldxxoe[2], ",")

				stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}
				_, err = stmt.Exec(fieldxxoe[0], fieldyyoe[0], fieldyzoe[0], "EDGE_CURVE")
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		if strings.Contains(line, "CARTESIAN_POINT") {
			fieldssss := strings.Split(line, "CARTESIAN_POINT('',(")

			if len(fieldssss) > 1 {
				fieldxxoe := strings.Split(fieldssss[1], ",")

				fieldyyoe := strings.Split(fieldxxoe[1], ",")
				fruitsLength := len(fieldxxoe)
				if fruitsLength > 3 {
					fieldyzoe := strings.Split(fieldxxoe[2], ")")
					//fmt.Println(fieldyzoe)

					stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
					if err != nil {
						log.Fatal(err)
					}

					if fieldyzoe != nil {

						_, err = stmt.Exec(fieldxxoe[0], fieldyyoe[0], fieldyzoe[0], "CARTESIAN_POINT")
						if err != nil {
							log.Fatal(err)
						}
					} else {
						_, err = stmt.Exec(fieldxxoe[0], fieldyyoe[0], 0, "CARTESIAN_POINT")
						if err != nil {
							log.Fatal(err)
						}
					}
				} else {
					fieldyyoes := strings.Split(fieldyyoe[0], ")")
					//fmt.Println(fieldyzoe)

					stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
					if err != nil {
						log.Fatal(err)
					}

					_, err = stmt.Exec(fieldxxoe[0], fieldyyoes[0], 0, "CARTESIAN_POINT")
					if err != nil {
						log.Fatal(err)
					}

				}

			}
		}
		if strings.Contains(line, "DIRECTION") {
			fieldssss := strings.Split(line, "DIRECTION('',(")

			if len(fieldssss) > 1 {
				fieldxxoe := strings.Split(fieldssss[1], ",")

				fieldyyoe := strings.Split(fieldxxoe[1], ",")
				fruitsLength := len(fieldxxoe)
				if fruitsLength > 3 {
					fieldyzoe := strings.Split(fieldxxoe[2], ")")
					//fmt.Println(fieldyzoe)

					stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
					if err != nil {
						log.Fatal(err)
					}

					if fieldyzoe != nil {

						_, err = stmt.Exec(fieldxxoe[0], fieldyyoe[0], fieldyzoe[0], "DIRECTION")
						if err != nil {
							log.Fatal(err)
						}
					} else {
						_, err = stmt.Exec(fieldxxoe[0], fieldyyoe[0], 0, "DIRECTION")
						if err != nil {
							log.Fatal(err)
						}
					}
				} else {

					//fmt.Println(fieldyzoe)

					stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
					if err != nil {
						log.Fatal(err)
					}
					fieldyyoes := strings.Split(fieldyyoe[0], ")")
					_, err = stmt.Exec(fieldxxoe[0], fieldyyoes[0], 0, "DIRECTION")
					if err != nil {
						log.Fatal(err)
					}

				}

			}
		}
		if strings.Contains(line, "LINE") {
			fieldssss := strings.Split(line, "LINE('',")

			if len(fieldssss) > 1 {
				fieldxxoe := strings.Split(fieldssss[1], ",")

				fieldyyoe := strings.Split(fieldxxoe[1], ")")

				//fieldyzoe := strings.Split(fieldxxoe[2], ")")
				//fmt.Println(fieldyzoe)

				stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}

				_, err = stmt.Exec(fieldxxoe[0], fieldyyoe[0], 0, "LINE")
				if err != nil {
					log.Fatal(err)
				}

			}
		}

		if strings.Contains(line, "AXIS2_PLACEMENT_3D") {
			fieldssss := strings.Split(line, "AXIS2_PLACEMENT_3D('',")

			if len(fieldssss) > 1 {
				fieldxxoe := strings.Split(fieldssss[1], ",")

				fieldyyoe := strings.Split(fieldxxoe[1], ")")

				fieldyzoe := strings.Split(fieldxxoe[2], ")")
				//fmt.Println(fieldyzoe)

				stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}

				_, err = stmt.Exec(fieldxxoe[0], fieldyyoe[0], fieldyzoe[0], "AXIS2_PLACEMENT_3D")
				if err != nil {
					log.Fatal(err)
				}

			}
		}
		if strings.Contains(line, "MANIFOLD_SURFACE_SHAPE_REPRESENTATION") {
			fieldssss := strings.Split(line, "MANIFOLD_SURFACE_SHAPE_REPRESENTATION('',")

			if len(fieldssss) > 1 {
				fieldxxoe := strings.Split(fieldssss[1], ",")

				fieldyyoe := strings.Split(fieldxxoe[1], ")")

				//fieldyzoe := strings.Split(fieldxxoe[2], ")")
				//fmt.Println(fieldyzoe)

				stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}
				fieldxxoes := strings.Split(fieldxxoe[0], "(")
				fieldxxoess := strings.Split(fieldxxoes[1], ")")
				_, err = stmt.Exec(fieldxxoess[0], fieldyyoe[0], 0, "MANIFOLD_SURFACE_SHAPE_REPRESENTATION")
				if err != nil {
					log.Fatal(err)
				}

			}
		}
		if strings.Contains(line, "ADVANCED_FACE") {
			fieldssss := strings.Split(line, "ADVANCED_FACE('',(")

			if len(fieldssss) > 1 {
				fieldxxoe := strings.Split(fieldssss[1], ",")

				fieldyyoe := strings.Split(fieldxxoe[1], ")")

				fieldyzoe := strings.Split(fieldxxoe[2], ")")
				//fmt.Println(fieldyzoe)

				stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}
				fieldxxoes := strings.Split(fieldxxoe[0], ")")
				_, err = stmt.Exec(fieldxxoes[0], fieldyyoe[0], fieldyzoe[0], "ADVANCED_FACE")
				if err != nil {
					log.Fatal(err)
				}

			}
		}

		if strings.Contains(line, "FACE_BOUND") {
			fieldssss := strings.Split(line, "FACE_BOUND('',")

			if len(fieldssss) > 1 {
				fieldxxoe := strings.Split(fieldssss[1], ",")

				fieldyyoe := strings.Split(fieldxxoe[1], ")")

				//fieldyzoe := strings.Split(fieldxxoe[2], ")")
				//fmt.Println(fieldyzoe)

				stmt, err := db.Prepare("INSERT INTO cartesian_point (x,y,z,type) VALUES (?,?,?,?)")
				if err != nil {
					log.Fatal(err)
				}

				_, err = stmt.Exec(fieldxxoe[0], fieldyyoe[0], 0, "FACE_BOUND")
				if err != nil {
					log.Fatal(err)
				}

			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Println("Lines:", line)

	//convert array to json
	// Query the table
	rows, err := db.Query("SELECT * FROM cartesian_point")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	//add json into file
	file, err = os.OpenFile("pointexss.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var people []Person
	var id int
	for rows.Next() {
		var p Person
		err := rows.Scan(&id, &p.X, &p.Y, &p.Z, &p.Types)
		//err := rows.Scan(&id, &p.X, &p.Y)
		if err != nil {
			fmt.Println(err)
			return
		}
		p.ID = "P" + strconv.Itoa(id)
		people = append(people, p)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// data, err := json.Marshal(people)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//members := string(data)
	//fmt.Println(string(jsonResult))
	// data, err := json.Marshal(p)

	group := Profile{

		Name:        "developer",
		Note:        "Generate by testing algoritham",
		Customer:    "developer",
		OrderNumber: "123456",
		Points:      people,
	}
	datas, err := json.Marshal(group)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(datas)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}
