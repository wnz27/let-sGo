/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/8/30 01:09 8月
 **/
package main

func reallyLongCalculation(done <-chan interface{}, value interface{}) interface{} {
	intermediateResult := longCalculation(value)
	select {
	case <-done:
		return nil
	default:
	}
	return longCalculattion(intermediateResult)
}

func reallyLongCalculation2(done <-chan interface{}, value interface{}) interface{} {
	intermediateResult := longCalculation(done, value)
	return longCalculattion(done, intermediateResult)
}




func main() {
	var value interface{}

	select {
	case <-done:
		return
	case value = <- valueStream:
	}

	result := reallyLongCalculation(value)

	select {
	case <-done:
		return
	case resultStream <-result:
	}
}

