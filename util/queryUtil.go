package util

// ex) str = "'a_value', 'b_value', 'c_value' ..."
func StringSerialize(list []string) string {
	var str string
	for i := 0; i < len(list); i++ {
		str += "'"
		str += list[i]
		str += "'"
		if (i < len(list) - 1) {
			str += ", "
		}
	}

	return str
}

// ex) str = "WHERE 1 = 1 AND a_column = a_value AND b_column = b_value ..."
func commonUtil(maps map[string]string) string {
	var str string

	str += " WHERE 1 = 1"
	for index, value := range maps {
		str += " AND " + index + " = " + value
	}

	return str
}

func DynamicInsertUtil(table string, maps map[string]string) string {
	var columns []string
	var values []string

	var index = 0;
	for column, value := range maps {
		columns[index] = column
		values[index] = value
		index++
	}

	query := "INSERT INTO "
	query += table + "("
	query += util.StringSerialize(columns)
	query += ") VALUES ("
	query += util.StringSerialize(values)
	query += ");"
	
	return query;
}
