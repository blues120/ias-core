package ibox

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"os/exec"
	"regexp"
	"strings"
)

func ParseDeviceConfigFile() map[string]string {
	// 获取设备信息
	// 获取设备信息
	cmd := fmt.Sprintf("cat /etc/ibox-release;" +
		"echo [hardware]" + "\n" +
		"echo CORES=$(grep 'processor' /proc/cpuinfo | sort -u | wc -l)-core" + "\n" +
		"echo MAC=$(cat /sys/class/net/`cat /etc/ibox-release |awk '/^ETH=/{print$0}'|awk -F\\\" '{print$2}'`/address)" + "\n" +
		"echo RamMemory=$(free -g | grep Mem | awk '{print $2+1}')" + "\n" +
		"echo RomMemory=$(df -h | awk '/^\\/dev/{print$2}' |grep G | awk -FG '{sum+=$1} END {print sum}')" + "\n")

	cmdExec := exec.Command("sh", "-c", cmd)
	cmdOutPut, err := cmdExec.Output()
	fmt.Println("cmdOutPut is ", string(cmdOutPut))
	if err != nil {
		log.Errorf("generateNodeLabels Error: %s %s", err, string(cmdOutPut))
		return nil
	}

	lines := strings.Split(string(cmdOutPut), "\n")
	pattern := `(\w+)\s*=\s*(?:"([^"]*)"|(\w+))`
	regex := regexp.MustCompile(pattern)

	attrMap := make(map[string]string)

	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		if len(match) < 4 {
			continue // 不匹配的行，跳过
		}
		key := match[1]
		if match[2] != "" {
			value := match[2]
			attrMap[key] = value
		} else if match[3] != "" {
			value := match[3]
			attrMap[key] = value
		}
	}
	log.Debugf("attrMap is %v\n", attrMap)

	device := map[string]string{
		"disk":     attrMap["RomMemory"],
		"mem":      attrMap["RamMemory"],
		"name":     attrMap["Device"],
		"cpu":      attrMap["CPU"] + "-" + attrMap["CORES"],
		"powerBy":  attrMap["PoweredBy"],
		"gpu":      attrMap["GPU"],
		"version":  attrMap["Version"],     // 系统版本号
		"model":    attrMap["DeviceModel"], // 设备型号
		"serialNo": attrMap["SerialNo"],    // 设备序列号
		"eth":      attrMap["ETH"],         // 设备型号
	}

	return device
}
