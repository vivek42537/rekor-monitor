//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mirroring

import (
	// "bufio"
	// "encoding/json"
	"fmt"
	"strconv"
	// "os"
	"testing"

	// "github.com/sigstore/rekor/pkg/client"
	// "github.com/spf13/viper"
	// "github.com/go-openapi/loads/fmts"
	// "github.com/mattn/go-sqlite3"
	"database/sql"
)

func TestGetLatestIndex(t *testing.T) {
	database, _ := sql.Open("sqlite3", "./test.db") //open database
	rows, err := getLatestX(database, 3)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	// if(id != 1999){
	// 	t.Errorf("Expected Result 1999, instead retrieved %d", id)
	// 	t.Errorf("payload is: %s", stringPay)
	// }
	var id int
	var payload string
	for rows.Next() {
		rows.Scan(&id, &payload)
		fmt.Println(strconv.Itoa(id) + ": " + payload)
	}

}

func TestInsert(t *testing.T) {
	database, _ := sql.Open("sqlite3", "./test.db")
	d := Data{
		ID:      8,
		Payload: "test4",
	}
	rows, err := Insert(database, d)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	if rows == -1 {
		t.Errorf("Expected to get a row insert but instead recieved error")
	}
}
