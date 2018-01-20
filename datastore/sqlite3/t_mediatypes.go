// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sqlite3

// import (
// 	"bytes"
// 	"github.com/freetaxii/libstix2/datastore"
// )

// ----------------------------------------------------------------------
//
// Collection Table Private Functions and Methods
//
// ----------------------------------------------------------------------

/*
mediaTypeProperties  - This function will return the properties that make up the
media types table

media_type    = A media type
*/
func mediaTypeProperties() string {
	return `
	"row_id" INTEGER PRIMARY KEY,
	"media_type" TEXT NOT NULL
	`
}
