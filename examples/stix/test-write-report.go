// Copyright 2015-2018 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/freetaxii/libstix2/objects"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := objects.NewReport()

	r.AddLabel("Attack Report")
	r.SetName("Malware Foo Report 2016")
	r.SetDescription("This report gives us details about Malware Foo")
	r.SetPublished(time.Now())

	r.AddObject(stix.NewId("malware"))
	// r.AddObject(stix.NewId("campaign"))
	// r.AddObject(stix.NewId("sighting"))
	// r.AddObject(stix.NewId("sighting"))
	// r.AddObject(stix.NewId("threat-actor"))
	// r.AddObject(stix.NewId("threat-actor"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))
	// r.AddObject(stix.NewId("relationship"))

	// for j := 0; j <= 4; j++ {
	// 	r.AddObject(stix.NewId("indicator"))
	// }

	// Open connection to database
	filename := "/opt/go/src/github.com/freetaxii/libstix2/examples/db/freetaxii.sqlite"
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatalf("Unable to open file %s due to error %v", filename, err)
	}
	defer db.Close()

	r.AddToDatabase(db)

	var data []byte
	data, _ = json.MarshalIndent(r, "", "    ")

	fmt.Println(string(data))
}
