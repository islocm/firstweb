package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func index(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tem.Execute(w, nil)
	// 	rows := [][]string{{"qwe", "hayr", "tun"}, {"qwe1", "saqwe", "sadag"}, {"asdqq", "tekin", "arzon"}}

	// 	for _, row := range rows {
	// 		var dbval []string
	// 		getinfo, err := db.Query(`select qaror from qaror`)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			return
	// 		}
	// 		for getinfo.Next() {
	// 			var asval string
	// 			getinfo.Scan(&asval)
	// 			dbval = append(dbval, asval)

	// 		}

	// 		for lentarget, valtarget := range dbval {

	// 			if valtarget == row[0] {
	// 				dbsorov := fmt.Sprintf(`UPDATE qaror
	// 					SET adres = '%s', buz = '%s'
	// 					WHERE qaror = '%s';`, row[1], row[2], valtarget)

	// 				resultdb, err := db.Exec(dbsorov)
	// 				if err != nil {
	// 					fmt.Println(err.Error())
	// 					return

	// 				}
	// 				fmt.Println(resultdb)
	// 				break

	// 			} else if len(dbval)-1 == lentarget {
	// 				dbsorov1 := fmt.Sprintf(`INSERT INTO qaror (qaror, adres, buz)
	// 				VALUES('%s', '%s', '%s');`, row[0], row[1], row[2])
	// 				_, err = db.Exec(dbsorov1)
	// 				if err != nil {
	// 					fmt.Println(err)
	// 					return

	// 				}

	// 			}

	// 		}

	// 	}
}

func olmazor(w http.ResponseWriter, r *http.Request) {

	tem, err := template.ParseFiles("template/Olmazor.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	tem.Execute(w, nil)

}
func datab(w http.ResponseWriter, r *http.Request) {
	tabledb, err := db.Prepare(`SELECT table_name
	FROM information_schema.tables
	WHERE table_type='BASE TABLE'
	AND table_schema='public';`)
	if err != nil {
		fmt.Println(err)
		return
	}
	var asd Name
	rows, _ := tabledb.Query()
	for rows.Next() {
		var qwe string
		rows.Scan(&qwe)

		asd = Name{[]string{qwe}}
		fmt.Println(asd)
	}

	// data := ViewData{
	// 	Tableg: []string{tableqq},
	// }
	x := r.FormValue("create")
	y := r.FormValue("alter")
	z := r.FormValue("column")
	q := r.FormValue("type")
	tem, err := template.ParseFiles("template/database.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	tem.Execute(w, asd)
	if x != "" {
		foo := fmt.Sprintf(`CREATE TABLE %s (anything varchar(50));`, x)
		_, err = db.Exec(foo)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if y != "" {
		altq := fmt.Sprintf(`ALTER TABLE %s ADD COLUMN %s %s;`, y, z, q)

		_, err = db.Exec(altq)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}

func wrexcel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tem, err := template.ParseFiles("template/excel.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		tem.Execute(w, nil)
	} else {

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		r.Header.Add("Content-Type", writer.FormDataContentType())
		r.ParseMultipartForm(32 << 20)
		file, header, err := r.FormFile("exwork")
		if err != nil {
			fmt.Println(err)
			return
		}

		f, err := os.OpenFile("./"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		io.Copy(f, file)

		exfile, err := excelize.OpenFile(header.Filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		getsheet := exfile.GetSheetList()

		rows, err := exfile.GetRows(getsheet[0])
		fmt.Println(len(rows))
		for _, row := range rows {
			var dbval []string
			getinfo, err := db.Query(`select kod from kadastr`)
			if err != nil {
				fmt.Println(err)
				return
			}
			for getinfo.Next() {
				var asval string
				getinfo.Scan(&asval)
				dbval = append(dbval, asval)

			}

			if dbval != nil {
				for lentarget, valtarget := range dbval {

					if valtarget == row[1] {
						row[3] = strings.ReplaceAll(row[3], "'", "")
						row[4] = strings.ReplaceAll(row[4], "'", "")

						dbsorov := fmt.Sprintf(`UPDATE kadastr
					SET mulk = '%s', mahalla = '%s', egalik= '%s', pasport = '%s', hujjat = '%s',
					regkitob = '%s', kitobbet = '%s', gosraqam = '%s', sananomer = '%s', miqdor = '%s', xona = '%s', sf = '%s',
					sv = '%s', po = '%s', pj = '%s', pp = '%s', pzuo = '%s', pzuz = '%s',
					pzuzaxvat = '%s', pzupd = '%s', pzupp = '%s', npp = '%s',
					npk = '%s', spp = '%s', spk = '%s'
					WHERE kod = '%s';`, row[0], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14], row[15], row[16], row[17], row[18], row[19], row[20], row[21], row[22], row[23], row[24], row[25], valtarget)

						_, err = db.Exec(dbsorov)
						if err != nil {
							fmt.Println(err.Error())
							return

						}

						break

					} else if len(dbval)-1 == lentarget {
						row[3] = strings.ReplaceAll(row[3], "'", "")
						row[4] = strings.ReplaceAll(row[4], "'", "")

						dbsorov1 := fmt.Sprintf(`INSERT INTO kadastr (mulk, kod, mahalla, egalik, pasport, hujjat, regkitob, kitobbet, gosraqam, sananomer, miqdor, xona, sf, sv, po, pj, pp, pzuo, pzuz, pzuzaxvat, pzupd, pzupp, npp, npk, spp, spk)
				VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14], row[15], row[16], row[17], row[18], row[19], row[20], row[21], row[22], row[23], row[24], row[25])

						_, err = db.Exec(dbsorov1)
						if err != nil {
							fmt.Println(err)
							return

						}

					}

				}
			} else {
				dbsorov1 := fmt.Sprintf(`INSERT INTO kadastr (mulk, kod, mahalla, egalik, pasport, hujjat, regkitob, kitobbet, gosraqam, sananomer, miqdor, xona, sf, sv, po, pj, pp, pzuo, pzuz, pzuzaxvat, pzupd, pzupp, npp, npk, spp, spk)
				VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14], row[15], row[16], row[17], row[18], row[19], row[20], row[21], row[22], row[23], row[24], row[25])
				_, err = db.Exec(dbsorov1)
				if err != nil {
					fmt.Println(err)
					return

				}

			}

		}
		file.Close()
		f.Close()
		err = os.Remove(header.Filename)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}

func hidedb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Database hidding query"))
	_, err := db.Exec(`CREATE TABLE "kadastr" (
		"mulk" varchar(255) NOT NULL,
		"kod" varchar(255) NOT NULL,
		"mahalla" varchar(255) NOT NULL,
		"egalik" varchar(255) NOT NULL,
		"pasport" varchar(255),
		"hujjat" varchar(255),
		"regkitob" varchar(50),
		"regkitob стр." varchar(50),
		"gosraqam" varchar(50),
		"sananomer" varchar(50),
		"miqdor" varchar(50),
		"xona" varchar(50),
		"sf" varchar(50),
		"sv" varchar(50),
		"po" varchar(50),
		"pj" varchar(50),
		"pp" varchar(50),
		"pzuo" varchar(50),
		"pzuzая" varchar(50),
		"pzuzaxvat" varchar(50),
		"pzupdором" varchar(50),
		"pzuppстройкой" varchar(50),
		"npp" varchar(50),
		npk" varchar(50),
		"spp" varchar(50),
		"spk" varchar(50)
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code executed ...")

}

// Nomida nima bor
type Nomida struct {
	Mulk      string
	Kod       string
	Mahalla   string
	Egalik    string
	Pasport   string
	Hujjat    string
	Regkitob  string
	Kitobbet  string
	Gosraqam  string
	Sananomer string
	Miqdor    string
	Xona      string
	Sf        string
	Sv        string
	Po        string
	Pj        string
	Pp        string
	Pzuo      string
	Pzuz      string
	Pzuzaxvat string
	Pzupd     string
	Pzupp     string
	Npp       string
	Npk       string
	Spp       string
	Spk       string
	Salom     []Nomida
}

func otiochsin(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("template/otiochsin.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	getval := r.FormValue("getinfo")
	symval := "%"

	querylike := fmt.Sprintf(`SELECT *
FROM
	kadastr
	WHERE
	Egalik LIKE '%s%s%s' OR Mulk LIKE '%s%s%s'
ORDER BY 
		Egalik; `, symval, getval, symval, symval, getval, symval)
	if getval != "" {
		Row := new(Nomida)

		rows, err := db.Query(querylike)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			err = rows.Scan(&Row.Mulk, &Row.Kod, &Row.Mahalla, &Row.Egalik, &Row.Pasport, &Row.Hujjat, &Row.Regkitob, &Row.Kitobbet, &Row.Gosraqam, &Row.Sananomer, &Row.Miqdor, &Row.Xona, &Row.Sf, &Row.Sv, &Row.Po, &Row.Pj, &Row.Pp, &Row.Pzuo, &Row.Pzuz, &Row.Pzuzaxvat, &Row.Pzupd, &Row.Pzupp, &Row.Npp, &Row.Npk, &Row.Spp, &Row.Spk)
			if err != nil {
				fmt.Println(err)
				return
			}
			Row.Salom = append(Row.Salom, Nomida{Mulk: Row.Mulk, Kod: Row.Kod, Mahalla: Row.Mahalla, Egalik: Row.Egalik, Pasport: Row.Pasport, Hujjat: Row.Hujjat, Regkitob: Row.Regkitob, Kitobbet: Row.Kitobbet, Gosraqam: Row.Gosraqam, Sananomer: Row.Sananomer, Miqdor: Row.Miqdor, Xona: Row.Xona, Sf: Row.Sf, Sv: Row.Sv, Po: Row.Po, Pj: Row.Pj, Pp: Row.Pp, Pzuo: Row.Pzuo, Pzuz: Row.Pzuz, Pzuzaxvat: Row.Pzuzaxvat, Pzupd: Row.Pzupd, Pzupp: Row.Pzupp, Npp: Row.Npp, Npk: Row.Npk, Spp: Row.Spp, Spk: Row.Spk})

			// kochir.Mulk = append(kochir.Mulk, Row.Mulk)
			// kochir.Kod = append(kochir.Kod, Row.Kod)
			// kochir.Mahalla = append(kochir.Mahalla, Row.Mahalla)
			// kochir.Egalik = append(kochir.Egalik, Row.Egalik)
			// kochir.Pasport = append(kochir.Pasport, Row.Pasport)
			// kochir.Hujjat = append(kochir.Hujjat, Row.Hujjat)
			// kochir.Regkitob = append(kochir.Regkitob, Row.Regkitob)
			// kochir.Kitobbet = append(kochir.Kitobbet, Row.Kitobbet)
			// kochir.Gosraqam = append(kochir.Gosraqam, Row.Gosraqam)
			// kochir.Sananomer = append(kochir.Sananomer, Row.Sananomer)
			// kochir.Miqdor = append(kochir.Miqdor, Row.Miqdor)
			// kochir.Xona = append(kochir.Xona, Row.Xona)
			// kochir.Sf = append(kochir.Sf, Row.Sf)
			// kochir.Sv = append(kochir.Sv, Row.Sv)
			// kochir.Po = append(kochir.Po, Row.Po)
			// kochir.Pj = append(kochir.Pj, Row.Pj)
			// kochir.Pp = append(kochir.Pp, Row.Pp)
			// kochir.Pzuo = append(kochir.Pzuo, Row.Pzuo)
			// kochir.Pzuz = append(kochir.Pzuz, Row.Pzuz)
			// kochir.Pzuzaxvat = append(kochir.Pzuzaxvat, Row.Pzuzaxvat)
			// kochir.Pzupd = append(kochir.Pzupd, Row.Pzupd)
			// kochir.Pzupp = append(kochir.Pzupp, Row.Pzupp)
			// kochir.Npp = append(kochir.Npp, Row.Npp)
			// kochir.Npk = append(kochir.Npk, Row.Npk)
			// kochir.Spp = append(kochir.Spp, Row.Spp)
			// kochir.Spk = append(kochir.Spk, Row.Spk)
		}

		temp.Execute(w, Row.Salom)
	} else {
		temp.Execute(w, nil)
	}

}

