package workshop4

func Isfriend(arg1 int, arg2 int) bool {
	var arg1bolenlertoplami int
	var arg2bolenlertoplami int

	for i := 1; i < arg1; i++ {
		if arg1%i == 0 {
			arg1bolenlertoplami = arg1bolenlertoplami + i
		}
	}
	for i := 1; i < arg2; i++ {
		if arg2%i == 0 {
			arg2bolenlertoplami = arg2bolenlertoplami + i
		}
	}

	if arg1bolenlertoplami == arg2 && arg2bolenlertoplami == arg1 {
		return true
	} else {
		return false
	}

}
