package bob

import (
	"strings"
)

func Hey(remark string) string {
	var outputs  = [4]string{
		"Fine. Be that way!",
		"Whoa, chill out!",
		"Whatever.",
		"Sure.",
	}

	remark = strings.Replace(remark, " ", "", -1)
	remark = strings.Replace(remark, "\n", "", -1)
	remark = strings.Replace(remark, "\t", "", -1)
	remark = strings.Replace(remark, "\r", "", -1)
	if len(remark) == 0 {
		return outputs[0]
	}

	ascii := []rune(remark);

	var caps, small, nums, misc, status int = 0,0,0,0,2
	for i := 0; i<len(remark); i++ {
		if ( ascii[i] >= 97 && ascii[i] <= 122 ) {
			small += 1
		} else if(ascii[i] >= 48 && ascii[i] <= 57){
			nums += 1
		} else if(ascii[i] >= 65 && ascii[i] <= 90) {
			caps += 1
		} else{
			misc += 1
		}
	}

	if small == 0 && caps > 0{
		status = 1
	} else if remark[len(remark) - 1] == '?'{
		status = 3
	}

	return outputs[status]
}
