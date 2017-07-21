// Copyright 2018 Mobile Data Books, LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/tecbot/gorocksdb"
	"path/filepath"
	"strings"
	"time"
)

//////////////////////////////////////////////////////
func main() {
	t0 := time.Now()
	t1 := time.Now()

	homePathDB := "/vagrant/"

	MDBrocksDBPath := homePathDB + "TC.RocksDB"

	dbRocksDB_NAICSCodes_Name := "NAICSCodes"

	//////////////////////////////////////////////////////////
	dbMainName := dbDir(MDBrocksDBPath, dbRocksDB_NAICSCodes_Name)
	options := gorocksdb.NewDefaultOptions()
	options.SetTargetFileSizeBase(67108864) //64MB
	//dbMain, _ := gorocksdb.OpenDb(options, dbMainName)
	dbMain, _ := gorocksdb.OpenDbForReadOnly(options, dbMainName, false)
	defer options.Destroy()
	defer dbMain.Close()
	//////////////////////////////////////////////////////////
	var pnum uint64
	pnum = 10
	fmt.Printf("\n%s|%d|%s|%s|%s|%s", "", pnum, " ", " ", " ", " ")

	Query := ""
	DBRocksKeys_List(
		Query,
		pnum,
		dbMain)

	//fmt.Printf("\ndbLevelMDB_PropertyAddressKeys=\n%s\n", dbLevelMDB_PropertyAddressKeys)
	//fmt.Printf("\n%s|%d|%s|\n%s|%s|%s", "", TotalCnt, "", "", " ", " ")
	t1 = time.Now()
	tStr1 := fmt.Sprintf("\nBEGIN=%v", t0)
	fmt.Printf("\n%s|%s|%s|%s|%s|%s", "", tStr1, " ", " ", " ", " ")
	tStr1 = fmt.Sprintf("\nEND=%v", t1)
	fmt.Printf("\n%s|%s|%s|%s|%s|%s", "", tStr1, " ", " ", " ", " ")
	tStr1 = fmt.Sprintf("\nIt took %v to run", t1.Sub(t0))
	fmt.Printf("\n%s|%s|%s|%s|%s|%s", "", tStr1, " ", " ", " ", " ")

}

////////////////////////////////////////////////////////////////////////////////////////////
func dbDir(home_path string, dbname string) string {
	//bottom := fmt.Sprintf("levigo-test-%d", rand.Int())
	//path := filepath.Join(os.TempDir(), bottom)
	path := filepath.Join(home_path, dbname)
	//deleteDBDirectory(path)
	//fmt.Printf("dbDir path: %v\n", path)
	return path
}

func DBRocksKeys_List(
	Query string,
	nPoints uint64,
	dbrocksdb *gorocksdb.DB) {

	routine := "DBRocksKeys_List"

	fmt.Printf("\n%s|%s|%s|%s|%s|%s", routine, "", "BEGIN", "<--------------------------------------------", "", "")
	fmt.Printf("\n%s|%s|%s|%s|%s|%s", routine, "", "Query", Query, "", "")
	///////////////////////////////////////
	ro := gorocksdb.NewDefaultReadOptions()
	//ro.SetFillCache(false)
	iter := dbrocksdb.NewIterator(ro)
	iter.SeekToFirst()
	defer ro.Destroy()
	defer iter.Close()
	///////////////////////////////////////
	if strings.TrimSpace(Query) == "" {

	} else {
		//fmt.Printf("\n%s|%s|%s|%s|%s|%s", routine, "iter.Seek(Query)):", Query, "", "", "")
		iter.Seek([]byte(Query))
	}
	//

	fmt.Printf("\n%s|%s|%v|%s|%s|%s", routine, "iter.Valid()", iter.Valid(), "", "", "")

	if !iter.Valid() {
		fmt.Errorf("Read iterator should be valid after seeking to first record")
		fmt.Printf("\n%s|%s|%s|%s|%s|%s", routine, "ERROR=", "iter.SeekToFirst()", "", "", "")
		fmt.Printf("\n%s|%s|%s|%s|%s|%s", routine, "ERROR=", "Read iterator should be valid after seeking to first record", "", "", "")

	} else {

		fmt.Printf("\n%s|%s|%s|%s|%s|%s", routine, "iter.SeekToFirst()", "OK", "", "", "")
		fmt.Printf("\n%s|%s|%d|%s|%s|%s", routine, "iter nPoints=", nPoints, "", "", "")

		//rec_ref := ""

		var i uint64
		i = 0
		for iter = iter; iter.Valid(); iter.Next() {
			if strings.TrimSpace(Query) == "" {

			} else {
				if strings.Contains(string(iter.Key().Data()), Query) {

				} else {
					break
				}

			}
			//if i > 100-1 {
			//	break
			//}
			if nPoints > 0 {
				if i > nPoints-1 {
					break
				}
			}
			//fmt.Printf("%00005d:%s=%s\n", i, iter.Key(), iter.Value())

			//sdelimited = fmt.Sprintf("%s", string(iter.Value()))
			// quick check
			//fmt.Printf("\n%s|%s|%s|%s|%s|%s", iter.Key(), sdelimited, "", "", "", "<->")
			sl0 := strings.Split(string(iter.Key().Data()), "|")
			fmt.Printf("\n%00005d[%d]:%s|%s|%s|%s|%s|%s==>", i+1, len(sl0), string(iter.Key().Data()), "", "", "", "", "")

			sl1 := strings.Split(string(iter.Value().Data()), "\t")
			fmt.Printf("\n%00005d[%d]:%s|%s|%s|%s|%s|%s", i+1, len(sl1), string(iter.Value().Data()), "", "", "", "", "")
			i++
		}
		fmt.Printf("\n\n%s|%s|%d|%s|%d|%s", routine, "nPoints=", nPoints, "i=", i, "")
	}
	fmt.Printf("\n%s|%s|%s|%s|%s|%s", routine, "", "END", "<--------------------------------------------", "", "")

	/////////////////////////////////////////////////

}
