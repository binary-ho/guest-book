package register

/*
TODO : (0, 0) 부터 시작할지 (1, 1) 부터 시작할지?
(0, 0)은 기본값이므로 문제가 생길 수도 있음.
일단은 (0, 0)으로 한다. (1, 1)의 경우 값을 클라이언트로 내려줄 때 까다로욱 수도 있음
*/
type Colors map[ColorKey]LWWRegister

type ColorKey uint32
