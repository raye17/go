package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"study/package/strings/vo"
)

var s = `grep: warning: stray \ before /
grep: warning: stray \ before /
+------------------------------------------------------------------------------------------------+
| npu-smi 25.0.rc1.1               Version: 25.0.rc1.1                                           |
+---------------------------+---------------+----------------------------------------------------+
| NPU   Name                | Health        | Power(W)    Temp(C)           Hugepages-Usage(page)|
| Chip                      | Bus-Id        | AICore(%)   Memory-Usage(MB)  HBM-Usage(MB)        |
+===========================+===============+====================================================+
| 0     910B4-1             | OK            | 93.6        32                0    / 0             |
| 0                         | 0000:C1:00.0  | 0           0    / 0          3381 / 65536         |
+===========================+===============+====================================================+
| 1     910B4-1             | OK            | 92.8        33                0    / 0             |
| 0                         | 0000:C2:00.0  | 0           0    / 0          3381 / 65536         |
+===========================+===============+====================================================+
| 2     910B4-1             | OK            | 88.1        32                0    / 0             |
| 0                         | 0000:81:00.0  | 0           0    / 0          3382 / 65536         |
+===========================+===============+====================================================+
| 3     910B4-1             | OK            | 89.2        31                0    / 0             |
| 0                         | 0000:82:00.0  | 0           0    / 0          3381 / 65536         |
+===========================+===============+====================================================+
| 4     910B4-1             | OK            | 91.7        34                0    / 0             |
| 0                         | 0000:01:00.0  | 0           0    / 0          3383 / 65536         |
+===========================+===============+====================================================+
| 5     910B4-1             | OK            | 92.4        37                0    / 0             |
| 0                         | 0000:02:00.0  | 0           0    / 0          3381 / 65536         |
+===========================+===============+====================================================+
| 6     910B4-1             | OK            | 92.2        37                0    / 0             |
| 0                         | 0000:41:00.0  | 0           0    / 0          46495/ 65536         |
+===========================+===============+====================================================+
| 7     910B4-1             | OK            | 88.4        35                0    / 0             |
| 0                         | 0000:42:00.0  | 0           0    / 0          46428/ 65536         |
+===========================+===============+====================================================+
+---------------------------+---------------+----------------------------------------------------+
| NPU     Chip              | Process id    | Process name             | Process memory(MB)      |
+===========================+===============+====================================================+
| No running processes found in NPU 0                                                            |
+===========================+===============+====================================================+
| No running processes found in NPU 1                                                            |
+===========================+===============+====================================================+
| No running processes found in NPU 2                                                            |
+===========================+===============+====================================================+
| No running processes found in NPU 3                                                            |
+===========================+===============+====================================================+
| No running processes found in NPU 4                                                            |
+===========================+===============+====================================================+
| No running processes found in NPU 5                                                            |
+===========================+===============+====================================================+
| 6       0                 | 3495003       | /usr/local/pyth          | 43087                   |
| 6       0                 | 3495004       | /usr/local/pyth          | 115                     |
+===========================+===============+====================================================+
| 7       0                 | 3495004       | /usr/local/pyth          | 43087                   |
+===========================+===============+====================================================+
`

func main() {
	res, _ := parsAscendFromNpuSmi(s)
	f, err := os.Create("result.json")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(res)
	defer f.Close()
	_, err = io.Copy(f, &buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("result.json")

}
func parsAscendFromNpuSmi(output string) ([]vo.GPUDeviceInfo, error) {
	//fmt.Printf("%q\n", output)
	output = strings.ReplaceAll(output, "\\n", "\n")
	lines := strings.Split(strings.TrimSpace(output), "\n")
	devices := make(map[int]vo.GPUDeviceInfo)

	var (
		inDeviceTable  = false
		inProcessTable = false
		headerLine     string
		blockLines     []string
		occupied       = make(map[int]bool)
	)
	//var currentIndex = -1
	flushBlock := func() {
		if len(blockLines) != 2 {
			blockLines = nil
			return
		}
		line1 := strings.TrimSpace(blockLines[0])
		line2 := strings.TrimSpace(blockLines[1])
		//解析第一行
		fields1 := parseTableLine(line1)
		if len(fields1) < 3 {
			blockLines = nil
			return
		}
		idName := strings.Fields(fields1[0])
		if len(idName) < 2 {
			blockLines = nil
			return
		}
		index, err := strconv.Atoi(idName[0])
		if err != nil {
			blockLines = nil
			return
		}
		name := idName[1]
		health := fields1[1]
		powerTemp := strings.Fields(fields1[2])
		if len(powerTemp) < 2 {
			blockLines = nil
			return
		}
		power, _ := strconv.ParseFloat(powerTemp[0], 64)
		temp, _ := strconv.ParseFloat(powerTemp[1], 64)
		dev := vo.GPUDeviceInfo{
			UUID:        string(vo.GPUVendorAscend) + "-" + strconv.Itoa(index),
			Index:       index,
			DeviceIndex: index,
			Vendor:      vo.GPUVendorAscend,
			Type:        vo.GPUBackendCANN,
			Name:        name,
			Temperature: temp,
			Health:      health,
			Power:       &vo.GPUPower{Draw: power},
			Core:        &vo.GPUCoreInfo{},
			Memory:      &vo.MemoryInfo{},
			Network:     &vo.GPUNetworkInfo{},
		}
		//解析第二行
		fields2 := parseTableLine(line2)
		if len(fields2) >= 3 {
			//AICore(%)
			aiCoreStr := fields2[2]
			if v, err := strconv.ParseFloat(aiCoreStr, 64); err == nil {
				dev.Core.UtilizationRate = v
			}
		}
		//解析HBM
		if headerLine != "" {
			if used, total, usage, ok := parseHBM(headerLine, line2); ok {
				dev.Memory.Used = used * 1024 * 1024   // HBM MB
				dev.Memory.Total = total * 1024 * 1024 // HBM MB
				dev.Memory.UtilizationRate = usage
			}
		}
		devices[index] = dev
		blockLines = nil
	}
	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}
		//设备表
		if strings.HasPrefix(line, "| NPU   Name") {
			inDeviceTable = true
			inProcessTable = false
			continue
		}
		//process table
		if strings.Contains(line, "Process id") {
			inDeviceTable = false
			inProcessTable = true
			continue
		}
		//HBM部分
		if inDeviceTable && strings.Contains(line, "HBM-Usage") {
			headerLine = line
			continue
		}
		if strings.HasPrefix(line, "+") {
			if inDeviceTable {
				flushBlock()
			}
			continue
		}
		//设备表处理
		if inDeviceTable && strings.HasPrefix(line, "|") {
			blockLines = append(blockLines, line)
			if len(blockLines) == 2 {
				flushBlock()
			}
			continue
		}
		if inProcessTable && strings.HasPrefix(line, "|") {
			//明确空闲的NPU
			if strings.Contains(line, "No running processes found in NPU") {
				id := extractNpuID(line)
				if id >= 0 {
					occupied[id] = false
				}
				continue
			}
			//fields := parseTableLine(line)
			//if len(fields) >= 3 {
			//	ids := strings.Fields(fields[0])
			//	if npuID, err := strconv.Atoi(ids[0]); err == nil {
			//		occupied[npuID] = true
			//	}
			//}
		}
	}
	// ---------------- 最终处理占用状态 ----------------
	var result []vo.GPUDeviceInfo
	for idx, dev := range devices {
		if _, ok := occupied[idx]; !ok {
			dev.InUse = true
		}
		result = append(result, dev)
	}

	// 可选：按 Index 排序
	sort.Slice(result, func(i, j int) bool { return result[i].Index < result[j].Index })

	return result, nil
}
func parseTableLine(line string) []string {
	line = strings.Trim(line, "|")
	parts := strings.Split(line, "|")

	var fields []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			fields = append(fields, p)
		}
	}
	return fields
}
func parseHBM(headerLine, dataLine string) (used, total int64, usage float64, ok bool) {
	idx := strings.Index(headerLine, "HBM-Usage(MB)")
	if idx == -1 {
		return
	}

	if len(dataLine) <= idx {
		return
	}

	segment := dataLine[idx:]

	re := regexp.MustCompile(`(\d+)\s*/\s*(\d+)`)
	match := re.FindStringSubmatch(segment)
	if len(match) != 3 {
		return
	}

	used, _ = strconv.ParseInt(match[1], 10, 64)
	total, _ = strconv.ParseInt(match[2], 10, 64)
	if total > 0 {
		usage = float64(used) / float64(total) * 100
	}
	ok = true
	return
}
func extractNpuID(line string) int {
	idx := strings.LastIndex(line, "NPU")
	if idx == -1 {
		return -1
	}
	idStr := strings.TrimSpace(line[idx+3:])
	idStr = strings.TrimRight(idStr, "|")
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1
	}
	return id
}
