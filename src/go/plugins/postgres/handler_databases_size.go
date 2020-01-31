/* /*
** Zabbix
** Copyright (C) 2001-2019 Zabbix SIA
**
** This program is free software; you can redistribute it and/or modify
** it under the terms of the GNU General Public License as published by
** the Free Software Foundation; either version 2 of the License, or
** (at your option) any later version.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program; if not, write to the Free Software
** Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
**/

package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	keyPostgresDatabasesSize = "pgsql.db.size"
)

// databasesSizeHandler gets info about count and size of archive files and returns JSON if all is OK or nil otherwise.
func (p *Plugin) databasesSizeHandler(conn *postgresConn, params []string) (interface{}, error) {
	var countSize int64
	var key, formatedLine string

	if len(params) > 0 {
		key = params[0]
		formatedLine = fmt.Sprintf(`select pg_database_size(datname::text) from pg_catalog.pg_database where  datistemplate = false and datname = '%v';`, key)
	} else {
		return nil, errorEmptyParam
	}

	err := conn.postgresPool.QueryRow(context.Background(), formatedLine).Scan(&countSize)
	if err != nil {
		if err == sql.ErrNoRows {
			p.Errf(err.Error())
			return nil, errorEmptyResult
		}
		p.Errf(err.Error())
		return nil, errorCannotFetchData
	}
	return countSize, nil
}
