package main

type valBirth struct {
	digits int
	min    int
	max    int
}
type valIssue struct {
	digits int
	min    int
	max    int
}
type valExp struct {
	digits int
	min    int
	max    int
}
type valHeight struct {
	heightOpts []heightOpt
}
type heightOpt struct {
	prefix string
	min    int
	max    int
}
type valHairColor struct {
	prefix  string
	digits  string
	letters string
}

type valEyeColor []string

type valPassportId int
