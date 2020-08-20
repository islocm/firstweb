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

// Copyt qwe
type Copyt struct {
	Idt        string
	Fiot       string
	Kodt       string
	Birtht     string
	Relationt  string
	Jont       string
	Manzilt    string
	Raqamt     string
	Vaqtt      string
	Yashasht   string
	Foydat     string
	Hujjatt    string
	Izoht      string
	Timet      string
	Usert      string
	Idselyamit string
}

// Copys qwe
type Copys struct {
	Ids       string
	Fios      string
	Kods      string
	Births    string
	Relations string
	Jons      string
	Manzils   string
	Raqams    string
	Vaqts     string
	Yashashs  string
	Foydas    string
	Hujjats   string
	Izoh      string
	Times     string
	Users     string
}

func index(w http.ResponseWriter, r *http.Request) {

	switch {
	case rDelete.MatchString(r.URL.Path):
		delete(w, r)
	case rChange.MatchString(r.URL.Path):
		changego(w, r)
	case rClean.MatchString(r.URL.Path):
		cleango(w, r)
	case rTarkib.MatchString(r.URL.Path):
		tarkiblink(w, r)
	case rChop.MatchString(r.URL.Path):
		chop(w, r)
	case rAbc.MatchString(r.URL.Path):
		selyamilink(w, r)
	case rComchange.MatchString(r.URL.Path):
		comchange(w, r)
	case rCom.MatchString(r.URL.Path):
		compensation(w, r)
	default:
		tem, err := template.ParseFiles("template/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)

	}
}

func olmazor(w http.ResponseWriter, r *http.Request) {
	// var name string
	// row := db.QueryRow(`select datei from import order by datei desc limit 1;`)
	// row.Scan(&name)

	qwe := sessionManager.Keys(r.Context())
	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {
			tem, err := template.ParseFiles("template/Success.html")
			if err != nil {
				fmt.Println(err)
				return
			}

			if r.FormValue("exit") == "exit" {
				sessionManager.Remove(r.Context(), name)
				http.Redirect(w, r, "/", 308)
			}

			tem.Execute(w, nil)
		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)
		}
	} else {

		if r.FormValue("email") != "" && r.FormValue("psw") != "" {
			forval := r.FormValue("email")
			pswval := r.FormValue("psw")
			getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, forval)
			getpswval := fmt.Sprintf(`select password from users where password = '%s';`, pswval)
			sorov := db.QueryRow(getuserval)
			sorov1 := db.QueryRow(getpswval)
			var name string
			var psw string
			sorov.Scan(&name)
			sorov1.Scan(&psw)
			if name == r.FormValue("email") && psw == r.FormValue("psw") {
				if len(qwe) == 0 {
					sessionManager.Put(r.Context(), name, name)
					tem, err := template.ParseFiles("template/Success.html")
					if err != nil {
						fmt.Println(err)
						return
					}
					tem.Execute(w, nil)
				} else {
					tem, err := template.ParseFiles("template/error.html")
					if err != nil {
						fmt.Println(err)
						return
					}
					tem.Execute(w, nil)
				}
			} else {
				tem, err := template.ParseFiles("template/error.html")
				if err != nil {
					fmt.Println(err)
					return
				}
				tem.Execute(w, nil)
			}
		} else {
			tem, err := template.ParseFiles("template/Olmazor.html")
			if err != nil {
				fmt.Println(err)
				return
			}

			tem.Execute(w, nil)

		}
	}
}

// Rowchange asd
type Rowchange struct {
	mulk      string
	kod       string
	mahalla   string
	egalik    string
	pasport   string
	hujjat    string
	regkitob  string
	kitobbet  string
	gosraqam  string
	sananomer string
	miqdor    string
	xona      string
	sf        string
	sv        string
	po        string
	pj        string
	pp        string
	pzuo      string
	pzuz      string
	pzuzaxvat string
	pzupd     string
	pzupp     string
	npp       string
	npk       string
	spp       string
	spk       string
}

// Qow asf
type Qow struct {
	useri string
}

//Qow1 asd
type Qow1 struct {
	Useri string
}

// Selyami asd
type Selyami struct {
	Ids            string
	Tumans         string
	Mulks          string
	Kods           string
	Kompensatsiyas string
	Sostavs        string
	Huquqs         string
	Xonas          string
	Izoh           string
	Times          string
	Users          string
	Salom          []Selyami
}

func spisok(w http.ResponseWriter, r *http.Request) {
	msg := sessionManager.GetString(r.Context(), "message")
	getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
	sorov := db.QueryRow(getuserval)
	var name string
	fmt.Println(r.URL.Path)
	sorov.Scan(&name)

	if name == msg {

		tem, err := template.ParseFiles("template/development.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		Row := new(Selyami)

		rows, err := db.Query(`SELECT * FROM "Selyami";`)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			err = rows.Scan(&Row.Ids, &Row.Tumans, &Row.Mulks, &Row.Kods, &Row.Kompensatsiyas, &Row.Sostavs, &Row.Huquqs, &Row.Xonas, &Row.Izoh, &Row.Times, &Row.Users)
			if err != nil {
				fmt.Println(err)
				return
			}
			Row.Salom = append(Row.Salom, Selyami{Ids: Row.Ids, Tumans: Row.Tumans, Mulks: Row.Mulks, Kods: Row.Kods, Kompensatsiyas: Row.Kompensatsiyas, Sostavs: Row.Sostavs, Huquqs: Row.Huquqs, Xonas: Row.Xonas, Izoh: Row.Izoh, Times: Row.Times, Users: Row.Users})
		}
		fmt.Println(r.FormValue("kodi"))
		if r.FormValue("kodi") != "" {
			kodi := r.FormValue("kodi")
			quer := fmt.Sprintf(`SELECT mulk, kod, mahalla, egalik, pasport, hujjat, regkitob, kitobbet, gosraqam, sananomer, miqdor, xona, sf, sv, po, pj, pp, pzuo, pzuz, pzuzaxvat, pzupd, pzupp, npp, npk, spp, spk
				FROM kadastr WHERE kod='%s';`, kodi)
			fmt.Println(quer)
			result := db.QueryRow(quer)
			row := new(Rowchange)
			result.Scan(&row.mulk, &row.kod, &row.mahalla,
				&row.egalik,
				&row.pasport,
				&row.hujjat,
				&row.regkitob,
				&row.kitobbet,
				&row.gosraqam,
				&row.sananomer,
				&row.miqdor,
				&row.xona,
				&row.sf,
				&row.sv,
				&row.po,
				&row.pj,
				&row.pp,
				&row.pzuo,
				&row.pzuz,
				&row.pzuzaxvat,
				&row.pzupd,
				&row.pzupp,
				&row.npp,
				&row.npk,
				&row.spp,
				&row.spk)

			fmt.Println(row.mulk)
			repairnull := ""

			lastquer := fmt.Sprintf(`INSERT INTO import (mulk, kod, mahalla, egalik, pasport, hujjat, regkitob, kitobbet, gosraqam, sananomer, miqdor, xona, sf, sv, po, pj, pp, pzuo, pzuz, pzuzaxvat, pzupd, pzupp, npp, npk, spp, spk, useri)
				VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, row.mulk, row.kod, row.mahalla, row.egalik, row.pasport, row.hujjat, row.regkitob, row.kitobbet, row.gosraqam, row.sananomer, row.miqdor, row.xona, row.sf, row.sv, row.po, row.pj, row.pp, row.pzuo, row.pzuz, row.pzuzaxvat, row.pzupd, row.pzupp, row.npp, row.npk, row.spp, row.spk, name)

			lastquery := fmt.Sprintf(`insert into "Selyami" (kods, mulks, users, tumans, kompensatsiyas, sostavs, huquqs, xonas, izoh)
				values('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, row.kod, row.mulk, name, repairnull, repairnull, repairnull, repairnull, repairnull, repairnull)
			_, err = db.Exec(lastquer)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = db.Exec(lastquery)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
		// If else
		if r.FormValue("kods") != "" {
			valforupdate := r.FormValue("tumans")
			valforupdate1 := r.FormValue("kods")
			valforupdate2 := r.FormValue("kompensatsiyas")
			valforupdate3 := r.FormValue("sostavs")
			valforupdate4 := r.FormValue("huquqs")
			valforupdate5 := r.FormValue("xonas")
			valforupdate6 := r.FormValue("izoh")

			dbsorov := fmt.Sprintf(`UPDATE "Selyami"
					SET tumans = '%s', kompensatsiyas = '%s', sostavs = '%s', huquqs = '%s', xonas = '%s', izoh = '%s', users = '%s'
					WHERE kods = '%s';`, valforupdate, valforupdate2, valforupdate3, valforupdate4, valforupdate5, valforupdate6, name, valforupdate1)

			_, err = db.Exec(dbsorov)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		tem.Execute(w, Row.Salom)
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)

	}

}

func digits(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.String()))

}

func abc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Has abc"))
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
	qwe := sessionManager.Keys(r.Context())
	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), "islocm")
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

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

				defer os.Remove(header.Filename)
				defer f.Close()
				defer file.Close()

				for _, row := range rows {
					row[0] = strings.ReplaceAll(row[0], "‘", "")
					row[3] = strings.ReplaceAll(row[3], "‘", "")
					row[4] = strings.ReplaceAll(row[4], "‘", "")
					row[0] = strings.ReplaceAll(row[0], "'", "")
					row[3] = strings.ReplaceAll(row[3], "`", "")
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

			}
		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)
		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)
	}
}

func hidedb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Database hidding query"))
	_, err := db.Exec(`CREATE TABLE "kadastr" (
		"mulk" varchar(255),
		"kod" varchar(255),
		"mahalla" varchar(255),
		"egalik" varchar(255),
		"pasport" varchar(255),
		"hujjat" varchar(255),
		"regkitob" varchar(50),
		"kitobbet" varchar(50),
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
		"pzuz" varchar(50),
		"pzuzaxvat" varchar(50),
		"pzupd" varchar(50),
		"pzupp" varchar(50),
		"npp" varchar(50),
		"npk" varchar(50),
		"spp" varchar(50),
		"spk" varchar(50)
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "import" (
		"id" serial,
		"qaror" varchar(50),
		"tuman" varchar(50),
		"mahalla" varchar(255),
		"kod" varchar(255),
		"nedvijimost" varchar(255),
		"pravoobladatel" varchar(255),
		"soprovoditelniy" varchar(255),
		"pzuo" varchar(50),
		"po" varchar(50),
		"pj" varchar(50),
		"xona" varchar(50),
		"datei" TIMESTAMP DEFAULT current_timestamp,
		"useri" varchar(50),
		CONSTRAINT "import_pk" PRIMARY KEY ("kod")
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "selyami" (
		"ids" serial,
		"fios" varchar(255),
		"kods" varchar(255) NOT NULL,
		"births" varchar(255),
		"relations" varchar(255),
		"jons" varchar(255),
		"manzils" varchar(255),
		"raqams" varchar(255),
		"vaqts" varchar(255),
		"yashashs" varchar(255),
		"foydas" varchar(255),
		"hujjats" varchar(255),
		"izoh" TEXT,
		"times" TIMESTAMP DEFAULT current_timestamp,
		"users" varchar(50) NOT NULL,
		CONSTRAINT "selyami_pk" PRIMARY KEY ("ids")
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "users" (
		"tel" varchar(14),
		"useru" varchar(50) NOT NULL,
		"password" varchar(30) NOT NULL,
		"dateu" TIMESTAMP NOT NULL DEFAULT current_timestamp,
		CONSTRAINT "users_pk" PRIMARY KEY ("useru")
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "compensation" (
		"idc" serial NOT NULL,
		"kodimport" varchar(255),
		"visionc" varchar(255),
		"kodc" varchar(255) NOT NULL,
		"manzilc" varchar(255),
		"maydonc" varchar(255),
		"xonac" varchar(255),
		"yermaydonc" varchar(255),
		"bozorc" varchar(255),
		"ijarac" varchar(255),
		"ijaramc" varchar(255),
		"protokolc" varchar(255),
		"orderc" varchar(255),
		"datac" TIMESTAMP DEFAULT current_timestamp,
		"userc" varchar(255),
		"idselyamic" int NOT NULL,
		CONSTRAINT "compensation_pk" PRIMARY KEY ("idc")
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "tarkib" (
		"idt" serial,
		"fiot" varchar(255),
		"kodt" varchar(255) NOT NULL,
		"birtht" varchar(255),
		"relationt" varchar(255),
		"jont" varchar(255),
		"manzilt" varchar(255),
		"raqamt" varchar(255),
		"vaqtt" varchar(255),
		"yashasht" varchar(255),
		"foydat" varchar(255),
		"hujjatt" varchar(255),
		"izoht" TEXT,
		"timet" TIMESTAMP DEFAULT current_timestamp,
		"usert" varchar(50) NOT NULL,
		"idselyamit" int NOT NULL
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "copys" (
		"statuss" varchar(50),
		"ids" varchar(255),
		"fios" varchar(255),
		"kods" varchar(255),
		"births" varchar(255),
		"relations" varchar(255),
		"jons" varchar(255),
		"manzils" varchar(255),
		"raqams" varchar(255),
		"vaqts" varchar(255),
		"yashashs" varchar(255),
		"foydas" varchar(255),
		"hujjats" varchar(255),
		"izoh" TEXT,
		"times" varchar(255),
		"users" varchar(50),
		"timesc" TIMESTAMP DEFAULT current_timestamp,
		"usersc" varchar(50)
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "copyt" (
		"statust" varchar(50),
		"idt" varchar(255),
		"fiot" varchar(255),
		"kodt" varchar(255),
		"birtht" varchar(255),
		"relationt" varchar(255),
		"jont" varchar(255),
		"manzilt" varchar(255),
		"raqamt" varchar(255),
		"vaqtt" varchar(255),
		"yashasht" varchar(255),
		"foydat" varchar(255),
		"hujjatt" varchar(255),
		"izoht" TEXT,
		"timet" varchar(255),
		"usert" varchar(50),
		"idselyamit" varchar(50),
		"timetc" TIMESTAMP DEFAULT current_timestamp,
		"usertc" varchar(50)
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "copyi" (
		"statusi" varchar(255),
		"id" varchar(50),
		"qaror" varchar(50),
		"tuman" varchar(50),
		"mahalla" varchar(255),
		"kod" varchar(255),
		"nedvijimost" varchar(255),
		"pravoobladatel" varchar(255),
		"soprovoditelniy" varchar(255),
		"pzuo" varchar(50),
		"po" varchar(50),
		"pj" varchar(50),
		"xona" varchar(50),
		"datei" varchar(255),
		"useri" varchar(50),
		"dateic" TIMESTAMP DEFAULT current_timestamp,
		"useric" varchar(255)
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	
	ALTER TABLE "import" ADD CONSTRAINT "import_fk0" FOREIGN KEY ("useri") REFERENCES "users"("useru");
	
	ALTER TABLE "selyami" ADD CONSTRAINT "selyami_fk0" FOREIGN KEY ("kods") REFERENCES "import"("kod");
	ALTER TABLE "selyami" ADD CONSTRAINT "selyami_fk1" FOREIGN KEY ("users") REFERENCES "users"("useru");
	
	
	ALTER TABLE "compensation" ADD CONSTRAINT "compensation_fk0" FOREIGN KEY ("kodimport") REFERENCES "import"("kod");
	ALTER TABLE "compensation" ADD CONSTRAINT "compensation_fk1" FOREIGN KEY ("userc") REFERENCES "users"("useru");
	ALTER TABLE "compensation" ADD CONSTRAINT "compensation_fk2" FOREIGN KEY ("idselyamic") REFERENCES "selyami"("ids");
	
	ALTER TABLE "tarkib" ADD CONSTRAINT "tarkib_fk0" FOREIGN KEY ("kodt") REFERENCES "import"("kod");
	ALTER TABLE "tarkib" ADD CONSTRAINT "tarkib_fk1" FOREIGN KEY ("usert") REFERENCES "users"("useru");
	ALTER TABLE "tarkib" ADD CONSTRAINT "tarkib_fk2" FOREIGN KEY ("idselyamit") REFERENCES "selyami"("ids");	
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

		}

		temp.Execute(w, Row.Salom)
	} else {
		temp.Execute(w, nil)
	}

}

// Import asd
type Import struct {
	ID              string
	Qaror           string
	Tuman           string
	Mahalla         string
	Kod             string
	Nedvijimost     string
	Pravoobladatel  string
	Soprovoditelniy string
	Pzuo            string
	Po              string
	Pj              string
	Xona            string
	Datei           string
	Useri           string
	Salom           []Import
}

// func importni(w http.ResponseWriter, r *http.Request) {
// 	temp, err := template.ParseFiles("template/import.html")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	getval := r.FormValue("getinfo")
// 	symval := "%"

// 	querylike := fmt.Sprintf(`SELECT *
// FROM
// 	import
// 	WHERE
// 	Egalik LIKE '%s%s%s' OR Mulk LIKE '%s%s%s'
// ORDER BY
// 		Egalik; `, symval, getval, symval, symval, getval, symval)
// 	if getval != "" {
// 		Row := new(Nomida)

// 		rows, err := db.Query(querylike)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		for rows.Next() {
// 			err = rows.Scan(&Row.Mulk, &Row.Kod, &Row.Mahalla, &Row.Egalik, &Row.Pasport, &Row.Hujjat, &Row.Regkitob, &Row.Kitobbet, &Row.Gosraqam, &Row.Sananomer, &Row.Miqdor, &Row.Xona, &Row.Sf, &Row.Sv, &Row.Po, &Row.Pj, &Row.Pp, &Row.Pzuo, &Row.Pzuz, &Row.Pzuzaxvat, &Row.Pzupd, &Row.Pzupp, &Row.Npp, &Row.Npk, &Row.Spp, &Row.Spk)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			Row.Salom = append(Row.Salom, Nomida{Mulk: Row.Mulk, Kod: Row.Kod, Mahalla: Row.Mahalla, Egalik: Row.Egalik, Pasport: Row.Pasport, Hujjat: Row.Hujjat, Regkitob: Row.Regkitob, Kitobbet: Row.Kitobbet, Gosraqam: Row.Gosraqam, Sananomer: Row.Sananomer, Miqdor: Row.Miqdor, Xona: Row.Xona, Sf: Row.Sf, Sv: Row.Sv, Po: Row.Po, Pj: Row.Pj, Pp: Row.Pp, Pzuo: Row.Pzuo, Pzuz: Row.Pzuz, Pzuzaxvat: Row.Pzuzaxvat, Pzupd: Row.Pzupd, Pzupp: Row.Pzupp, Npp: Row.Npp, Npk: Row.Npk, Spp: Row.Spp, Spk: Row.Spk})

// 			// kochir.Mulk = append(kochir.Mulk, Row.Mulk)
// 			// kochir.Kod = append(kochir.Kod, Row.Kod)
// 			// kochir.Mahalla = append(kochir.Mahalla, Row.Mahalla)
// 			// kochir.Egalik = append(kochir.Egalik, Row.Egalik)
// 			// kochir.Pasport = append(kochir.Pasport, Row.Pasport)
// 			// kochir.Hujjat = append(kochir.Hujjat, Row.Hujjat)
// 			// kochir.Regkitob = append(kochir.Regkitob, Row.Regkitob)
// 			// kochir.Kitobbet = append(kochir.Kitobbet, Row.Kitobbet)
// 			// kochir.Gosraqam = append(kochir.Gosraqam, Row.Gosraqam)
// 			// kochir.Sananomer = append(kochir.Sananomer, Row.Sananomer)
// 			// kochir.Miqdor = append(kochir.Miqdor, Row.Miqdor)
// 			// kochir.Xona = append(kochir.Xona, Row.Xona)
// 			// kochir.Sf = append(kochir.Sf, Row.Sf)
// 			// kochir.Sv = append(kochir.Sv, Row.Sv)
// 			// kochir.Po = append(kochir.Po, Row.Po)
// 			// kochir.Pj = append(kochir.Pj, Row.Pj)
// 			// kochir.Pp = append(kochir.Pp, Row.Pp)
// 			// kochir.Pzuo = append(kochir.Pzuo, Row.Pzuo)
// 			// kochir.Pzuz = append(kochir.Pzuz, Row.Pzuz)
// 			// kochir.Pzuzaxvat = append(kochir.Pzuzaxvat, Row.Pzuzaxvat)
// 			// kochir.Pzupd = append(kochir.Pzupd, Row.Pzupd)
// 			// kochir.Pzupp = append(kochir.Pzupp, Row.Pzupp)
// 			// kochir.Npp = append(kochir.Npp, Row.Npp)
// 			// kochir.Npk = append(kochir.Npk, Row.Npk)
// 			// kochir.Spp = append(kochir.Spp, Row.Spp)
// 			// kochir.Spk = append(kochir.Spk, Row.Spk)
// 		}

// 		temp.Execute(w, Row.Salom)
// 	} else {
// 		temp.Execute(w, nil)
// 	}

// }

func info(w http.ResponseWriter, r *http.Request) {

	qwe := sessionManager.Keys(r.Context())
	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			tem, err := template.ParseFiles("template/info.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			var checkkod string
			kodlist := fmt.Sprintf(`select kod from import where kod = '%s';`, r.FormValue("kod"))
			kodval := db.QueryRow(kodlist)
			kodval.Scan(&checkkod)
			var checkkod1 string
			kodlist1 := fmt.Sprintf(`select kod from import where kod = '%s';`, r.FormValue("checkgo"))
			kodval1 := db.QueryRow(kodlist1)
			kodval1.Scan(&checkkod1)

			if r.FormValue("check") != "on" && r.FormValue("kod") != checkkod {
				qaror := r.FormValue("qaror")
				tuman := r.FormValue("tuman")
				mahalla := r.FormValue("mahalla")
				kod := r.FormValue("kod")
				nedvijimost := r.FormValue("nedvijimost")
				pravoobladatel := r.FormValue("pravoobladatel")
				soprovoditelniy := r.FormValue("soprovoditelniy")
				pzuo := r.FormValue("pzuo")
				po := r.FormValue("po")
				pj := r.FormValue("pj")
				xona := r.FormValue("xona")
				nedvijimost = strings.ReplaceAll(nedvijimost, "'", "")
				pravoobladatel = strings.ReplaceAll(pravoobladatel, "'", "")
				pravoobladatel = strings.ReplaceAll(pravoobladatel, "‘", "")
				soprovoditelniy = strings.ReplaceAll(soprovoditelniy, "'", "")
				queryinsert := fmt.Sprintf(`insert into import (qaror, tuman, mahalla, kod, nedvijimost, pravoobladatel, soprovoditelniy, pzuo, po, pj, xona, useri)
				values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`, qaror, tuman, mahalla, kod, nedvijimost, pravoobladatel, soprovoditelniy, pzuo, po, pj, xona, name)

				_, err = db.Exec(queryinsert)
				if err != nil {
					fmt.Println(err)
					return
				}

			}

			if r.FormValue("check") == "on" && r.FormValue("checkgo") == checkkod1 {

				qaror := r.FormValue("qaror")
				tuman := r.FormValue("tuman")
				mahalla := r.FormValue("mahalla")
				kod := r.FormValue("kod")
				nedvijimost := r.FormValue("nedvijimost")
				pravoobladatel := r.FormValue("pravoobladatel")
				soprovoditelniy := r.FormValue("soprovoditelniy")
				pzuo := r.FormValue("pzuo")
				po := r.FormValue("po")
				pj := r.FormValue("pj")
				xona := r.FormValue("xona")
				check := r.FormValue("checkgo")

				queryinsert := fmt.Sprintf(`update import 
				SET qaror = '%s', tuman = '%s', mahalla = '%s', kod = '%s', nedvijimost = '%s', pravoobladatel = '%s', soprovoditelniy = '%s', pzuo = '%s', po = '%s', pj = '%s', xona = '%s', useri = '%s'
				WHERE kod = '%s';`, qaror, tuman, mahalla, kod, nedvijimost, pravoobladatel, soprovoditelniy, pzuo, po, pj, xona, name, check)

				_, err = db.Exec(queryinsert)
				if err != nil {
					fmt.Println(err)
					return
				}
				queryinsert1 := fmt.Sprintf(`update selyami set kods = '%s' where kods = '%s'`, kod, check)
				_, err = db.Exec(queryinsert1)
				if err != nil {
					fmt.Println(err)
					return
				}
				queryinsert2 := fmt.Sprintf(`update tarkib set kodt = '%s' where kodt = '%s'`, kod, check)
				_, err = db.Exec(queryinsert2)
				if err != nil {
					fmt.Println(err)
					return
				}
				queryinsert3 := fmt.Sprintf(`update compensation set kodc = '%s' where kodc = '%s'`, kod, check)
				_, err = db.Exec(queryinsert3)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			getval := r.FormValue("getinfo")
			symval := "%"

			querylike := fmt.Sprintf(`SELECT *
		FROM
		import
		WHERE
		kod LIKE '%s%s%s' OR nedvijimost LIKE '%s%s%s'
		ORDER BY 
		id; `, symval, getval, symval, symval, getval, symval)
			if getval != "" {
				Row := new(Import)

				rows, err := db.Query(querylike)
				if err != nil {
					fmt.Println(err)
				}
				for rows.Next() {
					err = rows.Scan(&Row.ID, &Row.Qaror, &Row.Tuman, &Row.Mahalla, &Row.Kod, &Row.Nedvijimost, &Row.Pravoobladatel, &Row.Soprovoditelniy, &Row.Pzuo, &Row.Po, &Row.Pj, &Row.Xona, &Row.Datei, &Row.Useri)
					if err != nil {
						fmt.Println(err)
						return
					}
					Row.Salom = append(Row.Salom, Import{ID: Row.ID, Qaror: Row.Qaror, Tuman: Row.Tuman, Mahalla: Row.Mahalla, Kod: Row.Kod, Nedvijimost: Row.Nedvijimost, Pravoobladatel: Row.Pravoobladatel, Soprovoditelniy: Row.Soprovoditelniy, Pzuo: Row.Pzuo, Po: Row.Po, Pj: Row.Pj, Xona: Row.Xona, Datei: Row.Datei, Useri: Row.Useri})

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

				tem.Execute(w, Row.Salom)
			} else {
				querylike := fmt.Sprintf(`select * from import order by id desc limit 20 ; `)

				Row := new(Import)

				rows, err := db.Query(querylike)
				if err != nil {
					fmt.Println(err)
				}
				for rows.Next() {
					err = rows.Scan(&Row.ID, &Row.Qaror, &Row.Tuman, &Row.Mahalla, &Row.Kod, &Row.Nedvijimost, &Row.Pravoobladatel, &Row.Soprovoditelniy, &Row.Pzuo, &Row.Po, &Row.Pj, &Row.Xona, &Row.Datei, &Row.Useri)
					if err != nil {
						fmt.Println(err)
						return
					}
					Row.Salom = append(Row.Salom, Import{ID: Row.ID, Qaror: Row.Qaror, Tuman: Row.Tuman, Mahalla: Row.Mahalla, Kod: Row.Kod, Nedvijimost: Row.Nedvijimost, Pravoobladatel: Row.Pravoobladatel, Soprovoditelniy: Row.Soprovoditelniy, Pzuo: Row.Pzuo, Po: Row.Po, Pj: Row.Pj, Xona: Row.Xona, Datei: Row.Datei, Useri: Row.Useri})

				}
				tem.Execute(w, Row.Salom)
			}

		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)

		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)
	}

}

func element(w http.ResponseWriter, r *http.Request) {

	qwe := sessionManager.Keys(r.Context())
	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), "islocm")
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			if r.Method == "GET" {
				tem, err := template.ParseFiles("template/element.html")
				if err != nil {
					fmt.Println(err)
					return
				}

				tem.Execute(w, nil)
			} else {
				defer http.Redirect(w, r, "/information", 307)

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

				exfile.GetRows(getsheet[0])
				defer os.Remove(header.Filename)
				defer f.Close()
				defer file.Close()
				rows, err := exfile.GetRows(getsheet[0])
				length := len(rows[0])

				if length == 11 {
					for _, row := range rows {
						row[3] = strings.ReplaceAll(row[3], "	", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						row[3] = strings.ReplaceAll(row[3], " ", "")
						var dbval []string
						getinfo, err := db.Query(`select kod from import`)
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

								if valtarget == row[3] {
									row[1] = strings.ReplaceAll(row[1], "'", "")
									row[2] = strings.ReplaceAll(row[2], "'", "")
									row[3] = strings.ReplaceAll(row[3], "'", "")
									row[4] = strings.ReplaceAll(row[4], "'", "")

									row[5] = strings.ReplaceAll(row[5], "‘", "")
									row[5] = strings.ReplaceAll(row[5], "'", "")
									row[6] = strings.ReplaceAll(row[6], "‘", "")
									row[6] = strings.ReplaceAll(row[6], "'", "")

									dbsorov := fmt.Sprintf(`UPDATE import
					SET qaror = '%s', tuman = '%s', mahalla= '%s', nedvijimost = '%s', pravoobladatel = '%s',
					soprovoditelniy = '%s', pzuo = '%s', po = '%s', pj = '%s', xona = '%s', useri = '%s'
					WHERE kod = '%s';`, row[0], row[1], row[2], row[4], row[5], row[6], row[7], row[8], row[9], row[10], name, valtarget)

									_, err = db.Exec(dbsorov)
									if err != nil {
										fmt.Println(err.Error())
										fmt.Println(row)
										return

									}

									break

								} else if len(dbval)-1 == lentarget {
									row[1] = strings.ReplaceAll(row[1], "'", "")
									row[2] = strings.ReplaceAll(row[2], "'", "")
									row[3] = strings.ReplaceAll(row[3], "'", "")
									row[4] = strings.ReplaceAll(row[4], "'", "")
									row[5] = strings.ReplaceAll(row[5], "‘", "")
									row[5] = strings.ReplaceAll(row[5], "'", "")
									row[6] = strings.ReplaceAll(row[6], "‘", "")
									row[6] = strings.ReplaceAll(row[6], "'", "")

									dbsorov1 := fmt.Sprintf(`INSERT INTO import (qaror, tuman, mahalla, kod, nedvijimost, pravoobladatel, soprovoditelniy, pzuo, po, pj, xona, useri)
							VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], name)

									_, err = db.Exec(dbsorov1)
									if err != nil {
										fmt.Println(err)
										fmt.Println(row)
										return

									}

								}

							}
						} else {
							row[1] = strings.ReplaceAll(row[1], "'", "")
							row[2] = strings.ReplaceAll(row[2], "'", "")
							row[3] = strings.ReplaceAll(row[3], "'", "")
							row[4] = strings.ReplaceAll(row[4], "'", "")
							row[5] = strings.ReplaceAll(row[5], "‘", "")
							row[5] = strings.ReplaceAll(row[5], "'", "")
							row[6] = strings.ReplaceAll(row[6], "‘", "")
							row[6] = strings.ReplaceAll(row[6], "'", "")
							dbsorov1 := fmt.Sprintf(`INSERT INTO import (qaror, tuman, mahalla, kod, nedvijimost, pravoobladatel, soprovoditelniy, pzuo, po, pj, xona, useri)
					VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], name)
							_, err = db.Exec(dbsorov1)
							if err != nil {
								fmt.Println(err)
								return

							}

						}

					}
				}

			}
		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}

			tem.Execute(w, nil)
		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		tem.Execute(w, nil)
	}
}

// Selyamigo qwe
type Selyamigo struct {
	Ids       string
	Fios      string
	Kods      string
	Births    string
	Relations string
	Jons      string
	Manzils   string
	Raqams    string
	Vaqts     string
	Yashashs  string
	Foydas    string
	Hujjats   string
	Izoh      string
	Times     string
	Users     string
	Salom     []Selyamigo
}

// Importdatabase qqmore
type Importdatabase struct {
	Kodi     string
	Manzili  string
	Umumiyi  string
	Yashashi string
	Hujjati  string
}

func selyamilink(w http.ResponseWriter, r *http.Request) {
	qwe := sessionManager.Keys(r.Context())
	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {
			tem, err := template.ParseFiles("template/selyami.html")
			if err != nil {
				fmt.Println(err)
				return
			}

			urlcode := r.URL.Path
			urlcode = strings.ReplaceAll(urlcode, "/selyami/", "")
			var checkkod string
			kodlist := fmt.Sprintf(`select kod from import where kod = '%s';`, urlcode)
			kodval := db.QueryRow(kodlist)
			kodval.Scan(&checkkod)
			if r.FormValue("jons") != "" && urlcode == checkkod {

				fios := r.FormValue("fios")
				births := r.FormValue("births")
				relations := r.FormValue("relations")
				jons := r.FormValue("jons")
				manzils := r.FormValue("manzils")
				raqams := r.FormValue("raqams")
				vaqts := r.FormValue("vaqts")
				yashashs := r.FormValue("yashashs")
				foydas := r.FormValue("foydas")
				hujjats := r.FormValue("hujjats")
				izoh := r.FormValue("izoh")
				fios = strings.ReplaceAll(fios, "'", "")
				manzils = strings.ReplaceAll(manzils, "'", "")
				fios = strings.ReplaceAll(fios, "‘", "")
				selyamiquery := fmt.Sprintf(`insert into selyami (fios, kods, births, relations, jons, manzils, raqams, vaqts, yashashs, foydas, hujjats, izoh, users)
			values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, fios, urlcode, births, relations, jons, manzils, raqams, vaqts, yashashs, foydas, hujjats, izoh, name)
				_, err = db.Exec(selyamiquery)
				if err != nil {
					fmt.Println(err)
					return
				}

				var num int

				idselyami := db.QueryRow(`select ids from selyami order by ids desc;`)
				idselyami.Scan(&num)

				selyamilinkquery := fmt.Sprintf(`insert into tarkib (fiot, kodt, birtht, relationt, jont, manzilt, raqamt, vaqtt, yashasht, foydat, hujjatt, izoht, usert, idselyamit)
			values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%d');`, fios, urlcode, births, relations, jons, manzils, raqams, vaqts, yashashs, foydas, hujjats, izoh, name, num)

				_, err = db.Exec(selyamilinkquery)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			if r.FormValue("update") == "12311231345345" {

				newrow, err := db.Query(`select kod, nedvijimost, po, pj, soprovoditelniy from import`)
				if err != nil {
					fmt.Println(err)
					return
				}
				for newrow.Next() {
					irow := new(Importdatabase)
					newrow.Scan(&irow.Kodi, &irow.Manzili, &irow.Umumiyi, &irow.Yashashi, &irow.Hujjati)

					queryinsert2 := fmt.Sprintf(`UPDATE selyami 
				SET manzils = '%s', yashashs = '%s',
				foydas = '%s', hujjats = '%s'
				WHERE kods = '%s';`, irow.Manzili, irow.Umumiyi, irow.Yashashi, irow.Hujjati, irow.Kodi)
					_, err = db.Exec(queryinsert2)
					if err != nil {
						fmt.Println(err)
						return
					}

					queryinsert3 := fmt.Sprintf(`UPDATE tarkib 
				SET manzilt = '%s', yashasht = '%s',
				foydat = '%s', hujjatt = '%s'
				WHERE kodt = '%s' AND relationt = 'О/Б';`, irow.Manzili, irow.Umumiyi, irow.Yashashi, irow.Hujjati, irow.Kodi)

					_, err = db.Exec(queryinsert3)
					if err != nil {
						fmt.Println(err)
						return
					}

				}

			}

			querylike := fmt.Sprintf(`select * from selyami where kods = '%s'; `, urlcode)

			Row := new(Selyamigo)

			rows, err := db.Query(querylike)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				err = rows.Scan(&Row.Ids, &Row.Fios, &Row.Kods, &Row.Births, &Row.Relations, &Row.Jons, &Row.Manzils, &Row.Raqams, &Row.Vaqts, &Row.Yashashs, &Row.Foydas, &Row.Hujjats, &Row.Izoh, &Row.Times, &Row.Users)
				if err != nil {
					fmt.Println(err)
					return
				}
				Row.Salom = append(Row.Salom, Selyamigo{Ids: Row.Ids, Fios: Row.Fios, Kods: Row.Kods, Births: Row.Births, Relations: Row.Relations, Jons: Row.Jons, Manzils: Row.Manzils, Raqams: Row.Raqams, Vaqts: Row.Vaqts, Yashashs: Row.Yashashs, Foydas: Row.Foydas, Hujjats: Row.Hujjats, Izoh: Row.Izoh, Times: Row.Times, Users: Row.Users})

			}
			tem.Execute(w, Row.Salom)

		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)
		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)
	}
}

// Userinfo qq
type Userinfo struct {
	Tel      string
	Useru    string
	Password string
	Vaqt     string
	Salom    []Userinfo
}

func islocm(w http.ResponseWriter, r *http.Request) {

	tem, err := template.ParseFiles("template/islocm.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.FormValue("password") != "" {
		tel := r.FormValue("tel")
		login := r.FormValue("login")
		pass := r.FormValue("password")
		query := fmt.Sprintf(`insert into users (tel, useru, password) values ('%s', '%s', '%s');`, tel, login, pass)
		_, err = db.Exec(query)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	queryuser, err := db.Query(`select * from users`)
	if err != nil {
		fmt.Println(err)
		return
	}
	Row := new(Userinfo)
	for queryuser.Next() {
		queryuser.Scan(&Row.Tel, &Row.Useru, &Row.Password, &Row.Vaqt)
		Row.Salom = append(Row.Salom, Userinfo{Tel: Row.Tel, Useru: Row.Useru, Password: Row.Password, Vaqt: Row.Vaqt})
	}
	tem.Execute(w, Row.Salom)
}

//Tarkib aasd
type Tarkib struct {
	Idt        string
	Fiot       string
	Kodt       string
	Birtht     string
	Relationt  string
	Jont       string
	Manzilt    string
	Raqamt     string
	Vaqtt      string
	Yashasht   string
	Foydat     string
	Hujjatt    string
	Izoht      string
	Timet      string
	Usert      string
	Idselyamit int
	Salom      []Tarkib
}

func tarkiblink(w http.ResponseWriter, r *http.Request) {
	qwe := sessionManager.Keys(r.Context())
	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			tem, err := template.ParseFiles("template/tarkib.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			var checkkod string
			urlcode := r.URL.Path
			urlcode = strings.ReplaceAll(urlcode, "/selyami/", "")

			listurl := strings.Split(urlcode, "*")

			kodlist := fmt.Sprintf(`select kod from import where kod = '%s';`, listurl[0])
			kodval := db.QueryRow(kodlist)
			kodval.Scan(&checkkod)

			if r.FormValue("fiot") != "" && listurl[0] == checkkod {

				fiot := r.FormValue("fiot")
				birtht := r.FormValue("birtht")
				relationt := r.FormValue("relationt")
				jont := r.FormValue("jont")
				manzilt := r.FormValue("manzilt")
				raqamt := r.FormValue("raqamt")
				vaqtt := r.FormValue("vaqtt")
				yashasht := r.FormValue("yashasht")
				foydat := r.FormValue("	foydat")
				hujjatt := r.FormValue("hujjatt")
				izoht := r.FormValue("izoht")
				fiot = strings.ReplaceAll(fiot, "'", "")
				fiot = strings.ReplaceAll(fiot, "‘", "")
				manzilt = strings.ReplaceAll(manzilt, "'", "")

				queryinsert := fmt.Sprintf(`insert into tarkib (fiot, kodt, birtht, relationt, jont, manzilt, raqamt, vaqtt, yashasht, foydat, hujjatt, izoht, usert, idselyamit)
			values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`, fiot, listurl[0], birtht, relationt, jont, manzilt, raqamt, vaqtt, yashasht, foydat, hujjatt, izoht, name, listurl[2])

				_, err = db.Exec(queryinsert)
				if err != nil {
					fmt.Println(err)
					return
				}

			}

			querylike := fmt.Sprintf(`SELECT *
		FROM
		tarkib
		WHERE
		idselyamit = '%s';`, listurl[2])

			Row := new(Tarkib)

			rows, err := db.Query(querylike)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				err = rows.Scan(&Row.Idt, &Row.Fiot, &Row.Kodt, &Row.Birtht, &Row.Relationt, &Row.Jont, &Row.Manzilt, &Row.Raqamt, &Row.Vaqtt, &Row.Yashasht, &Row.Foydat, &Row.Hujjatt, &Row.Izoht, &Row.Timet, &Row.Usert, &Row.Idselyamit)
				if err != nil {
					fmt.Println(err)
					return
				}
				Row.Salom = append(Row.Salom, Tarkib{Idt: Row.Idt, Fiot: Row.Fiot, Kodt: Row.Kodt, Birtht: Row.Birtht, Relationt: Row.Relationt, Jont: Row.Jont, Manzilt: Row.Manzilt, Raqamt: Row.Raqamt, Vaqtt: Row.Vaqtt, Yashasht: Row.Yashasht, Foydat: Row.Foydat, Hujjatt: Row.Hujjatt, Izoht: Row.Izoht, Timet: Row.Timet, Usert: Row.Usert, Idselyamit: Row.Idselyamit})

			}

			tem.Execute(w, Row.Salom)

		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)

		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)

	}

}

func selyamiexcel(w http.ResponseWriter, r *http.Request) {

	msg := sessionManager.GetString(r.Context(), "islocm")
	getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
	sorov := db.QueryRow(getuserval)
	var name string

	sorov.Scan(&name)

	if name == msg && msg != "" {

		if r.Method == "GET" {
			tem, err := template.ParseFiles("template/element.html")
			if err != nil {
				fmt.Println(err)
				return
			}

			tem.Execute(w, nil)
		} else {
			defer http.Redirect(w, r, "/information", 307)
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
			exfile.GetRows(getsheet[0])

			defer os.Remove(header.Filename)
			defer f.Close()
			defer file.Close()

			rows, err := exfile.GetRows(getsheet[0])
			length := len(rows[0])

			if length == 12 {
				for _, row := range rows {
					row[0] = strings.ReplaceAll(row[0], "‘", "")
					row[0] = strings.ReplaceAll(row[0], "'", "")
					row[5] = strings.ReplaceAll(row[5], "‘", "")

					row[1] = strings.ReplaceAll(row[1], "	", "")
					row[1] = strings.ReplaceAll(row[1], " ", "")
					row[1] = strings.ReplaceAll(row[1], " ", "")
					row[1] = strings.ReplaceAll(row[1], " ", "")
					row[1] = strings.ReplaceAll(row[1], " ", "")
					row[1] = strings.ReplaceAll(row[1], " ", "")
					row[1] = strings.ReplaceAll(row[1], " ", "")
					row[1] = strings.ReplaceAll(row[1], " ", "")
					row[1] = strings.ReplaceAll(row[1], " ", "")
					row[3] = strings.ReplaceAll(row[3], "	", "")
					row[3] = strings.ReplaceAll(row[3], " ", "")
					row[3] = strings.ReplaceAll(row[3], " ", "")
					row[3] = strings.ReplaceAll(row[3], " ", "")
					row[3] = strings.ReplaceAll(row[3], " ", "")
					row[3] = strings.ReplaceAll(row[3], " ", "")
					row[3] = strings.ReplaceAll(row[3], " ", "")
					row[3] = strings.ReplaceAll(row[3], " ", "")
					row[3] = strings.ReplaceAll(row[3], " ", "")
					row[3] = strings.ReplaceAll(row[3], " ", "")
					row[3] = strings.ReplaceAll(row[3], "б", "Б")
					row[3] = strings.ReplaceAll(row[3], "о", "О")
					row[5] = strings.ReplaceAll(row[5], "'", "")
					row[11] = strings.ReplaceAll(row[11], "'", "")

					selyamiquery := fmt.Sprintf(`insert into selyami (fios, kods, births, relations, jons, manzils, raqams, vaqts, yashashs, foydas, hujjats, izoh, users)
					values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], name)
					_, err = db.Exec(selyamiquery)
					if err != nil {
						fmt.Println(err)
						fmt.Println(row)
						return
					}

					var num int

					idselyami := db.QueryRow(`select ids from selyami order by ids desc;`)
					idselyami.Scan(&num)

					selyamilinkquery := fmt.Sprintf(`insert into tarkib (fiot, kodt, birtht, relationt, jont, manzilt, raqamt, vaqtt, yashasht, foydat, hujjatt, izoht, usert, idselyamit)
					values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%d');`, row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], name, num)

					_, err = db.Exec(selyamilinkquery)
					if err != nil {
						fmt.Println(err)
						fmt.Println(row)
						return
					}

				}
			}

		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		tem.Execute(w, nil)
	}
}

func changego(w http.ResponseWriter, r *http.Request) {
	qwe := sessionManager.Keys(r.Context())
	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			tem, err := template.ParseFiles("template/change.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			var checkkod string
			urlcode := r.URL.Path
			urlcode = strings.ReplaceAll(urlcode, "/selyami/", "")
			listurl := strings.Split(urlcode, "*")
			kodlist := fmt.Sprintf(`select kod from import where kod = '%s';`, listurl[0])
			kodval := db.QueryRow(kodlist)
			kodval.Scan(&checkkod)

			if r.FormValue("fiot") != "" && listurl[0] == checkkod {

				fiot := r.FormValue("fiot")
				birtht := r.FormValue("birtht")
				relationt := r.FormValue("relationt")
				jont := r.FormValue("jont")
				manzilt := r.FormValue("manzilt")
				raqamt := r.FormValue("raqamt")
				vaqtt := r.FormValue("vaqtt")
				yashasht := r.FormValue("yashasht")
				foydat := r.FormValue("	foydat")
				hujjatt := r.FormValue("hujjatt")
				izoht := r.FormValue("izoht")
				fiot = strings.ReplaceAll(fiot, "'", "")
				manzilt = strings.ReplaceAll(manzilt, "'", "")
				if r.FormValue("relationt") == "О/Б" {
					queryinsert := fmt.Sprintf(`UPDATE tarkib 
			SET fiot = '%s', kodt = '%s', birtht = '%s', relationt = '%s', jont = '%s', manzilt = '%s', raqamt = '%s', vaqtt = '%s', yashasht = '%s',
			foydat = '%s', hujjatt = '%s', izoht = '%s', usert = '%s'
			WHERE idt = '%s';`, fiot, listurl[0], birtht, relationt, jont, manzilt, raqamt, vaqtt, yashasht, foydat, hujjatt, izoht, name, listurl[4])

					_, err = db.Exec(queryinsert)
					if err != nil {
						fmt.Println(err)
						return
					}
					queryinsert2 := fmt.Sprintf(`UPDATE selyami 
			SET fios = '%s', kods = '%s', births = '%s', relations = '%s', jons = '%s', manzils = '%s', raqams = '%s', vaqts = '%s', yashashs = '%s',
			foydas = '%s', hujjats = '%s', izoh = '%s', users = '%s'
			WHERE ids = '%s';`, fiot, listurl[0], birtht, relationt, jont, manzilt, raqamt, vaqtt, yashasht, foydat, hujjatt, izoht, name, listurl[2])
					_, err = db.Exec(queryinsert2)
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					queryinsert3 := fmt.Sprintf(`UPDATE tarkib 
			SET fiot = '%s', kodt = '%s', birtht = '%s', relationt = '%s', jont = '%s', manzilt = '%s', raqamt = '%s', vaqtt = '%s', yashasht = '%s',
			foydat = '%s', hujjatt = '%s', izoht = '%s', usert = '%s'
			WHERE idt = '%s';`, fiot, listurl[0], birtht, relationt, jont, manzilt, raqamt, vaqtt, yashasht, foydat, hujjatt, izoht, name, listurl[4])

					_, err = db.Exec(queryinsert3)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}

			querylike := fmt.Sprintf(`SELECT *
		FROM
		tarkib
		WHERE
		idt = '%s';`, listurl[4])

			Row := new(Tarkib)

			rows, err := db.Query(querylike)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				err = rows.Scan(&Row.Idt, &Row.Fiot, &Row.Kodt, &Row.Birtht, &Row.Relationt, &Row.Jont, &Row.Manzilt, &Row.Raqamt, &Row.Vaqtt, &Row.Yashasht, &Row.Foydat, &Row.Hujjatt, &Row.Izoht, &Row.Timet, &Row.Usert, &Row.Idselyamit)
				if err != nil {
					fmt.Println(err)
					return
				}
				Row.Salom = append(Row.Salom, Tarkib{Idt: Row.Idt, Fiot: Row.Fiot, Kodt: Row.Kodt, Birtht: Row.Birtht, Relationt: Row.Relationt, Jont: Row.Jont, Manzilt: Row.Manzilt, Raqamt: Row.Raqamt, Vaqtt: Row.Vaqtt, Yashasht: Row.Yashasht, Foydat: Row.Foydat, Hujjatt: Row.Hujjatt, Izoht: Row.Izoht, Timet: Row.Timet, Usert: Row.Usert, Idselyamit: Row.Idselyamit})

			}

			tem.Execute(w, Row.Salom)

		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)

		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)

	}

}

// Compensation qweasd qwads
type Compensation struct {
	Idc        string
	Kodimport  string
	Visionc    string
	Kodc       string
	Manzilc    string
	Maydonc    string
	Xonac      string
	Yermaydonc string
	Bozorc     string
	Ijarac     string
	Ijaramc    string
	Protokolc  string
	Orderc     string
	Datac      string
	Userc      string
	Idselyamic string
	Salom      []Compensation
}

func compensation(w http.ResponseWriter, r *http.Request) {

	qwe := sessionManager.Keys(r.Context())

	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			tem, err := template.ParseFiles("template/compensation.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			var checkkod string
			urlcode := r.URL.Path
			urlcode = strings.ReplaceAll(urlcode, "/compensation/", "")
			listurl := strings.Split(urlcode, "*")
			kodlist := fmt.Sprintf(`select kod from import where kod = '%s';`, listurl[0])
			kodval := db.QueryRow(kodlist)
			kodval.Scan(&checkkod)

			if r.FormValue("manzilc") != "" && listurl[0] == checkkod {

				visionc := r.FormValue("visionc")
				kodc := r.FormValue("kodc")
				manzilc := r.FormValue("manzilc")
				maydonc := r.FormValue("maydonc")
				xonac := r.FormValue("xonac")
				yermaydonc := r.FormValue("yermaydonc")
				bozorc := r.FormValue("bozorc")
				ijarac := r.FormValue("ijarac")
				ijaramc := r.FormValue("ijaramc")
				protokolc := r.FormValue("protokolc")
				orderc := r.FormValue("orderc")
				manzilc = strings.ReplaceAll(manzilc, "'", "")

				queryinsert := fmt.Sprintf(`insert into compensation (kodimport, visionc, kodc, manzilc, maydonc, xonac, yermaydonc, bozorc, ijarac, ijaramc, protokolc, orderc, userc, idselyamic)
			values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`, listurl[0], visionc, kodc, manzilc, maydonc, xonac, yermaydonc, bozorc, ijarac, ijaramc, protokolc, orderc, name, listurl[1])

				_, err = db.Exec(queryinsert)
				if err != nil {
					fmt.Println(err)
					return
				}

			}

			querylike := fmt.Sprintf(`SELECT *
		FROM
		compensation
		WHERE
		idselyamic = '%s';`, listurl[1])

			Row := new(Compensation)

			rows, err := db.Query(querylike)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				err = rows.Scan(&Row.Idc, &Row.Kodimport, &Row.Visionc, &Row.Kodc, &Row.Manzilc, &Row.Maydonc, &Row.Xonac, &Row.Yermaydonc, &Row.Bozorc, &Row.Ijarac, &Row.Ijaramc, &Row.Protokolc, &Row.Orderc, &Row.Datac, &Row.Userc, &Row.Idselyamic)
				if err != nil {
					fmt.Println(err)
					return
				}
				Row.Salom = append(Row.Salom, Compensation{Idc: Row.Idc, Kodimport: Row.Kodimport, Visionc: Row.Visionc, Kodc: Row.Kodc, Manzilc: Row.Manzilc, Maydonc: Row.Maydonc, Xonac: Row.Xonac, Yermaydonc: Row.Yermaydonc, Bozorc: Row.Bozorc, Ijarac: Row.Ijarac, Ijaramc: Row.Ijaramc, Protokolc: Row.Protokolc, Orderc: Row.Orderc, Datac: Row.Datac, Userc: Row.Userc, Idselyamic: Row.Idselyamic})

			}

			tem.Execute(w, Row.Salom)

		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)

		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)

	}

}

func zaprost(w http.ResponseWriter, r *http.Request) {
	qwe := sessionManager.Keys(r.Context())
	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			tem, err := template.ParseFiles("template/zapros.html")
			if err != nil {
				fmt.Println(err)
				return
			}

			querylike := fmt.Sprintf(`SELECT *
		FROM
		tarkib
		order by idselyamit,
		idt asc;`)

			Row := new(Tarkib)

			rows, err := db.Query(querylike)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				err = rows.Scan(&Row.Idt, &Row.Fiot, &Row.Kodt, &Row.Birtht, &Row.Relationt, &Row.Jont, &Row.Manzilt, &Row.Raqamt, &Row.Vaqtt, &Row.Yashasht, &Row.Foydat, &Row.Hujjatt, &Row.Izoht, &Row.Timet, &Row.Usert, &Row.Idselyamit)
				if err != nil {
					fmt.Println(err)
					return
				}
				Row.Salom = append(Row.Salom, Tarkib{Idt: Row.Idt, Fiot: Row.Fiot, Kodt: Row.Kodt, Birtht: Row.Birtht, Relationt: Row.Relationt, Jont: Row.Jont, Manzilt: Row.Manzilt, Raqamt: Row.Raqamt, Vaqtt: Row.Vaqtt, Yashasht: Row.Yashasht, Foydat: Row.Foydat, Hujjatt: Row.Hujjatt, Izoht: Row.Izoht, Timet: Row.Timet, Usert: Row.Usert, Idselyamit: Row.Idselyamit})

			}

			tem.Execute(w, Row.Salom)

		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)

		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)

	}

}

func comchange(w http.ResponseWriter, r *http.Request) {

	qwe := sessionManager.Keys(r.Context())

	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			tem, err := template.ParseFiles("template/comchange.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			var checkkod string
			urlcode := r.URL.Path
			urlcode = strings.ReplaceAll(urlcode, "/compensation/", "")

			listurl := strings.Split(urlcode, "*")
			kodlist := fmt.Sprintf(`select kod from import where kod = '%s';`, listurl[0])
			kodval := db.QueryRow(kodlist)
			kodval.Scan(&checkkod)

			if r.FormValue("manzilc") != "" && listurl[0] == checkkod {

				visionc := r.FormValue("visionc")
				kodc := r.FormValue("kodc")
				manzilc := r.FormValue("manzilc")
				maydonc := r.FormValue("maydonc")
				xonac := r.FormValue("xonac")
				yermaydonc := r.FormValue("yermaydonc")
				bozorc := r.FormValue("bozorc")
				ijarac := r.FormValue("ijarac")
				ijaramc := r.FormValue("ijaramc")
				protokolc := r.FormValue("protokolc")
				orderc := r.FormValue("orderc")

				queryinsert := fmt.Sprintf(`UPDATE compensation 
				SET kodimport = '%s', visionc = '%s', kodc = '%s', manzilc = '%s', maydonc = '%s', xonac = '%s', yermaydonc = '%s', bozorc = '%s', ijarac = '%s', ijaramc = '%s', protokolc = '%s', orderc = '%s', userc = '%s'
				WHERE idc = '%s';`, listurl[0], visionc, kodc, manzilc, maydonc, xonac, yermaydonc, bozorc, ijarac, ijaramc, protokolc, orderc, name, listurl[3])

				_, err = db.Exec(queryinsert)
				if err != nil {
					fmt.Println(err)
					return
				}

			}

			querylike := fmt.Sprintf(`SELECT *
		FROM
		compensation
		WHERE
		idselyamic = '%s';`, listurl[1])

			Row := new(Compensation)

			rows, err := db.Query(querylike)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				err = rows.Scan(&Row.Idc, &Row.Kodimport, &Row.Visionc, &Row.Kodc, &Row.Manzilc, &Row.Maydonc, &Row.Xonac, &Row.Yermaydonc, &Row.Bozorc, &Row.Ijarac, &Row.Ijaramc, &Row.Protokolc, &Row.Orderc, &Row.Datac, &Row.Userc, &Row.Idselyamic)
				if err != nil {
					fmt.Println(err)
					return
				}
				Row.Salom = append(Row.Salom, Compensation{Idc: Row.Idc, Kodimport: Row.Kodimport, Visionc: Row.Visionc, Kodc: Row.Kodc, Manzilc: Row.Manzilc, Maydonc: Row.Maydonc, Xonac: Row.Xonac, Yermaydonc: Row.Yermaydonc, Bozorc: Row.Bozorc, Ijarac: Row.Ijarac, Ijaramc: Row.Ijaramc, Protokolc: Row.Protokolc, Orderc: Row.Orderc, Datac: Row.Datac, Userc: Row.Userc, Idselyamic: Row.Idselyamic})

			}

			tem.Execute(w, Row.Salom)

		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)

		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)

	}

}

// Copyi qwe
type Copyi struct {
	ID              string
	Qaror           string
	Tuman           string
	Mahalla         string
	Kod             string
	Nedvijimost     string
	Pravoobladatel  string
	Soprovoditelniy string
	Pzuo            string
	Po              string
	Pj              string
	Xona            string
	Datei           string
	Useri           string
}

func delete(w http.ResponseWriter, r *http.Request) {

	qwe := sessionManager.Keys(r.Context())
	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			urlcode := r.URL.Path
			urlcode = strings.ReplaceAll(urlcode, "/delete/", "")

			vicopy := fmt.Sprintf(`select * from import where kod = '%s'`, urlcode)

			vicopyq := db.QueryRow(vicopy)

			goes := new(Copyi)

			vicopyq.Scan(&goes.ID, &goes.Qaror, &goes.Tuman, &goes.Mahalla, &goes.Kod, &goes.Nedvijimost, &goes.Pravoobladatel, &goes.Soprovoditelniy, &goes.Pzuo, &goes.Po, &goes.Pj, &goes.Xona, &goes.Datei, &goes.Useri)
			status := "Deleted"
			queryinsert := fmt.Sprintf(`insert into copyi (statusi, id, qaror, tuman, mahalla, kod, nedvijimost, pravoobladatel, soprovoditelniy, pzuo, po, pj, xona, datei, useri, useric)
				values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`, status, goes.ID, goes.Qaror, goes.Tuman, goes.Mahalla, goes.Kod, goes.Nedvijimost, goes.Pravoobladatel, goes.Soprovoditelniy, goes.Pzuo, goes.Po, goes.Pj, goes.Xona, goes.Datei, goes.Useri, name)

			_, err := db.Exec(queryinsert)
			if err != nil {
				fmt.Println(err)
				return
			}

			kodlist := fmt.Sprintf(`delete from import where kod = '%s';`, urlcode)
			_, err = db.Exec(kodlist)
			if err != nil {
				fmt.Println(err)
				return
			}
			http.Redirect(w, r, "/information", 307)
		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)
		}
	} else {
		tem, err := template.ParseFiles("template/error.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tem.Execute(w, nil)
	}
}

func cleango(w http.ResponseWriter, r *http.Request) {

	qwe := sessionManager.Keys(r.Context())

	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			var checkkod string
			var fiosgo string
			var fiotgo string
			urlcode := r.URL.Path
			urlcode = strings.ReplaceAll(urlcode, "/selyami/", "")
			listurl := strings.Split(urlcode, "*")
			kodlist := fmt.Sprintf(`select kod from import where kod = '%s';`, listurl[0])
			kodval := db.QueryRow(kodlist)
			kodval.Scan(&checkkod)

			fiotch := fmt.Sprintf(`select fiot from tarkib where kodt = '%s' AND idt = '%s';`, listurl[0], listurl[4])
			fiotchq := db.QueryRow(fiotch)
			fiotchq.Scan(&fiotgo)

			fiosch := fmt.Sprintf(`select fios from selyami where kods = '%s' AND ids = '%s';`, listurl[0], listurl[2])
			fioschq := db.QueryRow(fiosch)
			fioschq.Scan(&fiosgo)
			status := "Deleted"

			if fiotgo != fiosgo && listurl[0] == checkkod {
				redir := fmt.Sprintf(`/selyami/%s*tarkib*%s`, listurl[0], listurl[2])
				defer http.Redirect(w, r, redir, 307)

				hit := fmt.Sprintf(`select * from tarkib where idt = '%s'`, listurl[4])
				hitq := db.QueryRow(hit)

				action := new(Copyt)
				err := hitq.Scan(&action.Idt, &action.Fiot, &action.Kodt, &action.Birtht, &action.Relationt, &action.Jont, &action.Manzilt, &action.Raqamt, &action.Vaqtt, &action.Yashasht, &action.Foydat, &action.Hujjatt, &action.Izoht, &action.Timet, &action.Usert, &action.Idselyamit)
				if err != nil {
					fmt.Println(err)
					return
				}

				selyamilinkquery := fmt.Sprintf(`insert into copyt (statust, idt, fiot, kodt, birtht, relationt, jont, manzilt, raqamt, vaqtt, yashasht, foydat, hujjatt, izoht, timet, usert, idselyamit, usertc)
					values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, status, action.Idt, action.Fiot, action.Kodt, action.Birtht, action.Relationt, action.Jont, action.Manzilt, action.Raqamt, action.Vaqtt, action.Yashasht, action.Foydat, action.Hujjatt, action.Izoht, action.Timet, action.Usert, action.Idselyamit, name)

				_, err = db.Exec(selyamilinkquery)
				if err != nil {
					fmt.Println(err)

					return
				}

				cleant := fmt.Sprintf(`delete from tarkib where idt = '%s';`, listurl[4])
				_, err = db.Exec(cleant)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				tem, err := template.ParseFiles("template/error.html")
				if err != nil {
					fmt.Println(err)
					return
				}
				tem.Execute(w, nil)

			}
		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)

		}

	}
}

func chop(w http.ResponseWriter, r *http.Request) {

	qwe := sessionManager.Keys(r.Context())

	if len(qwe) == 1 {
		msg := sessionManager.GetString(r.Context(), qwe[0])
		getuserval := fmt.Sprintf(`select useru from users where useru = '%s';`, msg)
		sorov := db.QueryRow(getuserval)
		var name string

		sorov.Scan(&name)

		if name == msg && msg != "" {

			var checkkod string

			urlcode := r.URL.Path
			urlcode = strings.ReplaceAll(urlcode, "/selyami/", "")
			listurl := strings.Split(urlcode, "*")
			kodlist := fmt.Sprintf(`select kod from import where kod = '%s';`, listurl[0])
			kodval := db.QueryRow(kodlist)
			kodval.Scan(&checkkod)

			if listurl[0] == checkkod {
				redir := fmt.Sprintf(`/selyami/%s`, listurl[0])
				defer http.Redirect(w, r, redir, 307)
				his := fmt.Sprintf(`select * from selyami where ids = '%s'`, listurl[2])
				hisq := db.QueryRow(his)
				move := new(Copys)
				hisq.Scan(&move.Ids, &move.Fios, &move.Kods, &move.Births, &move.Relations, &move.Jons, &move.Manzils, &move.Raqams, &move.Vaqts, &move.Yashashs, &move.Foydas, &move.Hujjats, &move.Izoh, &move.Times, &move.Users)
				status := "Deleted"
				selyamiquery := fmt.Sprintf(`insert into copys (statuss, ids, fios, kods, births, relations, jons, manzils, raqams, vaqts, yashashs, foydas, hujjats, izoh, times, users, usersc)
					values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s', '%s','%s', '%s');`, status, move.Ids, move.Fios, move.Kods, move.Births, move.Relations, move.Jons, move.Manzils, move.Raqams, move.Vaqts, move.Yashashs, move.Foydas, move.Hujjats, move.Izoh, move.Times, move.Users, name)
				_, err := db.Exec(selyamiquery)
				if err != nil {
					fmt.Println(err)

					return
				}

				hit := fmt.Sprintf(`select * from tarkib where idselyamit = '%s'`, listurl[2])
				hitq, err := db.Query(hit)
				if err != nil {
					fmt.Println(err)
					return
				}
				for hitq.Next() {
					action := new(Copyt)
					hitq.Scan(&action.Idt, &action.Fiot, &action.Kodt, &action.Birtht, &action.Relationt, &action.Jont, &action.Manzilt, &action.Raqamt, &action.Vaqtt, &action.Yashasht, &action.Foydat, &action.Hujjatt, &action.Izoht, &action.Timet, &action.Usert, &action.Idselyamit)

					selyamilinkquery := fmt.Sprintf(`insert into copyt (statust, idt, fiot, kodt, birtht, relationt, jont, manzilt, raqamt, vaqtt, yashasht, foydat, hujjatt, izoht, timet, usert, idselyamit, usertc)
					values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');`, status, action.Idt, action.Fiot, action.Kodt, action.Birtht, action.Relationt, action.Jont, action.Manzilt, action.Raqamt, action.Vaqtt, action.Yashasht, action.Foydat, action.Hujjatt, action.Izoht, action.Timet, action.Usert, action.Idselyamit, name)

					_, err = db.Exec(selyamilinkquery)
					if err != nil {
						fmt.Println(err)

						return
					}
				}

				cleans := fmt.Sprintf(`delete from tarkib where idselyamit = '%s';`, listurl[2])
				_, err = db.Exec(cleans)
				if err != nil {
					fmt.Println(err)
					return
				}
				cleant := fmt.Sprintf(`delete from selyami where ids = '%s';`, listurl[2])
				_, err = db.Exec(cleant)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				tem, err := template.ParseFiles("template/error.html")
				if err != nil {
					fmt.Println(err)
					return
				}
				tem.Execute(w, nil)

			}
		} else {
			tem, err := template.ParseFiles("template/error.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			tem.Execute(w, nil)

		}

	}
}
