package leap

const testVersion = 3

func IsLeapYear( year int) bool {
	// Write some code here to pass the test suite.
    if(year % 100 == 0 ){
        if(year % 400 == 0){
            return true
        }
    }else if(year % 4 == 0){
        return true
    }
    return false
}
