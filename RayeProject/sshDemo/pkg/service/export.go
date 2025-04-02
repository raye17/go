package service

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
)

// 导出user表数据到CSV
func ExportUserToCSV(db *sql.DB) error {
	// 获取总行数
	var totalRows int
	err := db.QueryRow("SELECT COUNT(*) FROM register_record").Scan(&totalRows)
	if err != nil {
		return fmt.Errorf("获取总行数失败: %w", err)
	}

	// 初始化进度条
	bar := progressbar.NewOptions(totalRows,
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionSetWidth(10),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
	)

	rows, err := db.Query("SELECT * FROM register_record")
	if err != nil {
		return fmt.Errorf("查询user表失败: %w", err)
	}
	defer rows.Close()

	// 创建CSV文件
	file, err := os.Create("user.csv")
	if err != nil {
		return fmt.Errorf("创建CSV文件失败: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("获取列名失败: %w", err)
	}
	writer.Write(columns)

	// 读取数据并写入CSV
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(valuePtrs...)
		if err != nil {
			return fmt.Errorf("读取数据失败: %w", err)
		}

		// 将值转换为字符串
		record := make([]string, len(values))
		for i, v := range values {
			if b, ok := v.([]byte); ok {
				record[i] = string(b)
			} else {
				record[i] = fmt.Sprintf("%v", v)
			}
		}
		writer.Write(record)
		bar.Add(1)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("读取行时出错: %w", err)
	}

	fmt.Println("成功将user表数据导出到user.csv")
	return nil
}
