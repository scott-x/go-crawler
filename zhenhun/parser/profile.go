package parser

import (
	"github.com/scott-x/go-crawler/engine"
	"github.com/scott-x/go-crawler/model"
	"regexp"
)

var personalInfoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">(.+)</div>`)
var heartRe = regexp.MustCompile(`<span data-v-bff6f798="">(.+)</span>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {

	profile := model.Profile{}
	personalInfo := ExtractField(personalInfoRe, contents)
	heart := ExtractField(heartRe, contents)

	profile.Name = name
	profile.PersonalInfo = personalInfo
	profile.Heart = heart

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func ExtractField(re *regexp.Regexp, contents []byte) string {
	match := re.FindSubmatch(contents)
	if match != nil {
		return string(match[1])
	} else {
		return ""
	}
}
