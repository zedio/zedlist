// Copyright 2015-2017 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migration

import (
	"testing"
)

func TestMigration(t *testing.T) {
	DropTables()
	MigrateTables()
}
