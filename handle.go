package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func index(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tem.Execute(w, nil)

}

func olmazor(w http.ResponseWriter, r *http.Request) {
	x := r.FormValue("name")
	y := r.FormValue("qwe")
	fmt.Println(x)
	fmt.Println(y)
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

	tem, err := template.ParseFiles("template/excel.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	tem.Execute(w, nil)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.ParseMultipartForm(10 >> 20)
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

	_, err = excelize.OpenFile(header.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
	f.Close()
	err = os.Remove(header.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// rows, err := f.GetRows("TDSheet")

	// for _, row := range rows {
	// 	var a []string
	// 	// qq = fmt.Printf(" (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)", row)
	// 	a = row

	// 	qq := fmt.Sprintf("INSERT INTO kadastr (islocm,islocm1,islocm2,islocm3,islocm4,islocm5,islocm6,islocm7,islocm8,islocm9,islocm10,islocm12,islocm13,islocm14,islocm15,islocm16,islocm17,islocm18,islocm19,islocm20,islocm21,islocm22,islocm23,islocm24,islocm25) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');", a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], a[16], a[17], a[18], a[19], a[20], a[21], a[22], a[23], a[24])

	// 	fmt.Println(qq)
	// 	q, err := db.Exec(qq)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(q)
	// 	break
	// }
	// fmt.Println(f)

}

func hidedb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Database hidding query"))
	_, err := db.Exec(`CREATE TABLE "qaror" (
		"qaror" character varying NOT NULL,
		CONSTRAINT "qaror_pk" PRIMARY KEY ("qaror")
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "malumot" (
		"Data" DATE NOT NULL,
		"resheniya" character varying NOT NULL,
		"vladelets" character varying NOT NULL,
		"adres" character varying NOT NULL,
		CONSTRAINT "malumot_pk" PRIMARY KEY ("adres")
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	CREATE TABLE "kompensatsiya" (
		"adresg" character varying NOT NULL,
		"kompensatsiya" character varying NOT NULL,
		"xona" character varying NOT NULL
	) WITH (
	  OIDS=FALSE
	);
	
	
	
	
	ALTER TABLE "malumot" ADD CONSTRAINT "malumot_fk0" FOREIGN KEY ("resheniya") REFERENCES "qaror"("qaror");
	
	ALTER TABLE "kompensatsiya" ADD CONSTRAINT "kompensatsiya_fk0" FOREIGN KEY ("adresg") REFERENCES "malumot"("adres");
	`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code executed ...")

}
