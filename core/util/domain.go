/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

//copied from http://git.oschina.net/Unknown/gowalker/blob/master/utils/path.go

import (
	"path"
	"regexp"
	"strings"
)

var validTLD = map[string]bool{
	// curl http://data.iana.org/TLD/tlds-alpha-by-domain.txt | sed  -e '/#/ d' -e 's/.*/"&": true,/' | tr [:upper:] [:lower:]
	"ac":                     true,
	"ad":                     true,
	"ae":                     true,
	"aero":                   true,
	"af":                     true,
	"ag":                     true,
	"ai":                     true,
	"al":                     true,
	"am":                     true,
	"an":                     true,
	"ao":                     true,
	"aq":                     true,
	"ar":                     true,
	"arpa":                   true,
	"as":                     true,
	"asia":                   true,
	"at":                     true,
	"au":                     true,
	"aw":                     true,
	"ax":                     true,
	"az":                     true,
	"ba":                     true,
	"bb":                     true,
	"bd":                     true,
	"be":                     true,
	"bf":                     true,
	"bg":                     true,
	"bh":                     true,
	"bi":                     true,
	"biz":                    true,
	"bj":                     true,
	"bm":                     true,
	"bn":                     true,
	"bo":                     true,
	"br":                     true,
	"bs":                     true,
	"bt":                     true,
	"bv":                     true,
	"bw":                     true,
	"by":                     true,
	"bz":                     true,
	"ca":                     true,
	"cat":                    true,
	"cc":                     true,
	"cd":                     true,
	"cf":                     true,
	"cg":                     true,
	"ch":                     true,
	"ci":                     true,
	"ck":                     true,
	"cl":                     true,
	"cm":                     true,
	"cn":                     true,
	"co":                     true,
	"com":                    true,
	"coop":                   true,
	"cr":                     true,
	"cu":                     true,
	"cv":                     true,
	"cw":                     true,
	"cx":                     true,
	"cy":                     true,
	"cz":                     true,
	"de":                     true,
	"dj":                     true,
	"dk":                     true,
	"dm":                     true,
	"do":                     true,
	"dz":                     true,
	"ec":                     true,
	"edu":                    true,
	"ee":                     true,
	"eg":                     true,
	"er":                     true,
	"es":                     true,
	"et":                     true,
	"eu":                     true,
	"fi":                     true,
	"fj":                     true,
	"fk":                     true,
	"fm":                     true,
	"fo":                     true,
	"fr":                     true,
	"ga":                     true,
	"gb":                     true,
	"gd":                     true,
	"ge":                     true,
	"gf":                     true,
	"gg":                     true,
	"gh":                     true,
	"gi":                     true,
	"gl":                     true,
	"gm":                     true,
	"gn":                     true,
	"gov":                    true,
	"gp":                     true,
	"gq":                     true,
	"gr":                     true,
	"gs":                     true,
	"gt":                     true,
	"gu":                     true,
	"gw":                     true,
	"gy":                     true,
	"hk":                     true,
	"hm":                     true,
	"hn":                     true,
	"hr":                     true,
	"ht":                     true,
	"hu":                     true,
	"id":                     true,
	"ie":                     true,
	"il":                     true,
	"im":                     true,
	"in":                     true,
	"info":                   true,
	"int":                    true,
	"io":                     true,
	"iq":                     true,
	"ir":                     true,
	"is":                     true,
	"it":                     true,
	"je":                     true,
	"jm":                     true,
	"jo":                     true,
	"jobs":                   true,
	"jp":                     true,
	"ke":                     true,
	"kg":                     true,
	"kh":                     true,
	"ki":                     true,
	"km":                     true,
	"kn":                     true,
	"kp":                     true,
	"kr":                     true,
	"kw":                     true,
	"ky":                     true,
	"kz":                     true,
	"la":                     true,
	"lb":                     true,
	"lc":                     true,
	"li":                     true,
	"lk":                     true,
	"lr":                     true,
	"ls":                     true,
	"lt":                     true,
	"lu":                     true,
	"lv":                     true,
	"ly":                     true,
	"ma":                     true,
	"mc":                     true,
	"md":                     true,
	"me":                     true,
	"mg":                     true,
	"mh":                     true,
	"mil":                    true,
	"mk":                     true,
	"ml":                     true,
	"mm":                     true,
	"mn":                     true,
	"mo":                     true,
	"mobi":                   true,
	"mp":                     true,
	"mq":                     true,
	"mr":                     true,
	"ms":                     true,
	"mt":                     true,
	"mu":                     true,
	"museum":                 true,
	"mv":                     true,
	"mw":                     true,
	"mx":                     true,
	"my":                     true,
	"mz":                     true,
	"na":                     true,
	"name":                   true,
	"nc":                     true,
	"ne":                     true,
	"net":                    true,
	"nf":                     true,
	"ng":                     true,
	"ni":                     true,
	"nl":                     true,
	"no":                     true,
	"np":                     true,
	"nr":                     true,
	"nu":                     true,
	"nz":                     true,
	"om":                     true,
	"org":                    true,
	"pa":                     true,
	"pe":                     true,
	"pf":                     true,
	"pg":                     true,
	"ph":                     true,
	"pk":                     true,
	"pl":                     true,
	"pm":                     true,
	"pn":                     true,
	"post":                   true,
	"pr":                     true,
	"pro":                    true,
	"ps":                     true,
	"pt":                     true,
	"pw":                     true,
	"py":                     true,
	"qa":                     true,
	"re":                     true,
	"ro":                     true,
	"rs":                     true,
	"ru":                     true,
	"rw":                     true,
	"sa":                     true,
	"sb":                     true,
	"sc":                     true,
	"sd":                     true,
	"se":                     true,
	"sg":                     true,
	"sh":                     true,
	"si":                     true,
	"sj":                     true,
	"sk":                     true,
	"sl":                     true,
	"sm":                     true,
	"sn":                     true,
	"so":                     true,
	"sr":                     true,
	"st":                     true,
	"su":                     true,
	"sv":                     true,
	"sx":                     true,
	"sy":                     true,
	"sz":                     true,
	"tc":                     true,
	"td":                     true,
	"tel":                    true,
	"tf":                     true,
	"tg":                     true,
	"th":                     true,
	"tj":                     true,
	"tk":                     true,
	"tl":                     true,
	"tm":                     true,
	"tn":                     true,
	"to":                     true,
	"tp":                     true,
	"tr":                     true,
	"travel":                 true,
	"tt":                     true,
	"tv":                     true,
	"tw":                     true,
	"tz":                     true,
	"ua":                     true,
	"ug":                     true,
	"uk":                     true,
	"us":                     true,
	"uy":                     true,
	"uz":                     true,
	"va":                     true,
	"vc":                     true,
	"ve":                     true,
	"vg":                     true,
	"vi":                     true,
	"vn":                     true,
	"vu":                     true,
	"wf":                     true,
	"ws":                     true,
	"xn--0zwm56d":            true,
	"xn--11b5bs3a9aj6g":      true,
	"xn--3e0b707e":           true,
	"xn--45brj9c":            true,
	"xn--80akhbyknj4f":       true,
	"xn--80ao21a":            true,
	"xn--90a3ac":             true,
	"xn--9t4b11yi5a":         true,
	"xn--clchc0ea0b2g2a9gcd": true,
	"xn--deba0ad":            true,
	"xn--fiqs8s":             true,
	"xn--fiqz9s":             true,
	"xn--fpcrj9c3d":          true,
	"xn--fzc2c9e2c":          true,
	"xn--g6w251d":            true,
	"xn--gecrj9c":            true,
	"xn--h2brj9c":            true,
	"xn--hgbk6aj7f53bba":     true,
	"xn--hlcj6aya9esc7a":     true,
	"xn--j1amh":              true,
	"xn--j6w193g":            true,
	"xn--jxalpdlp":           true,
	"xn--kgbechtv":           true,
	"xn--kprw13d":            true,
	"xn--kpry57d":            true,
	"xn--l1acc":              true,
	"xn--lgbbat1ad8j":        true,
	"xn--mgb9awbf":           true,
	"xn--mgbaam7a8h":         true,
	"xn--mgbayh7gpa":         true,
	"xn--mgbbh1a71e":         true,
	"xn--mgbc0a9azcg":        true,
	"xn--mgberp4a5d4ar":      true,
	"xn--mgbx4cd0ab":         true,
	"xn--o3cw4h":             true,
	"xn--ogbpf8fl":           true,
	"xn--p1ai":               true,
	"xn--pgbs0dh":            true,
	"xn--s9brj9c":            true,
	"xn--wgbh1c":             true,
	"xn--wgbl6a":             true,
	"xn--xkc2al3hye2a":       true,
	"xn--xkc2dl3a5ee0h":      true,
	"xn--yfro4i67o":          true,
	"xn--ygbi2ammx":          true,
	"xn--zckzah":             true,
	"xxx":                    true,
	"ye":                     true,
	"yt":                     true,
	"za":                     true,
	"zm":                     true,
	"zw":                     true,
}

var (
	validHost        = regexp.MustCompile(`^[-a-z0-9]+(?:\.[-a-z0-9]+)+$`)
	validPathElement = regexp.MustCompile(`^[-A-Za-z0-9~+][-A-Za-z0-9_.]*$`)
)

// IsValidRemotePath returns true if importPath is structurally valid for "go get".
func IsValidRemotePath(importPath string) bool {
	parts := strings.Split(importPath, "/")

	if len(parts) <= 1 {
		// Import path must contain at least one "/".
		return false
	}

	if !validTLD[path.Ext(parts[0])] {
		return false
	}

	if !validHost.MatchString(parts[0]) {
		return false
	}
	for _, part := range parts[1:] {
		if !validPathElement.MatchString(part) || part == "testdata" {
			return false
		}
	}

	return true
}

func joinPath(importPath string, num int) string {
	subdirs := strings.Split(importPath, "/")
	if len(subdirs) > num {
		return strings.Join(subdirs[:num], "/")
	}
	return importPath
}

var filterDirNames = []string{
	"static", "docs", "views", "js", "assets", "public", "img", "conf", "css"}

// FilterDirName guess the file or directory is or contains Go source files.
func FilterDirName(name string) bool {
	for _, v := range filterDirNames {
		if strings.Contains(strings.ToLower(name), v+"/") {
			return false
		}
	}
	return true
}
