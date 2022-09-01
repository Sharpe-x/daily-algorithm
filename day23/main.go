// 459.重复的子字符串

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern("ababa"))
	//fmt.Println(repeatedSubstringPattern("rstvvqxwwytlhtmddecbeluloucjbnjrltxkwykjehlpnioghuarrnitpgbxkjvenuphpmpszyospiyeitiyodcpnynckdqxozqmgghogviwkauvzwjvuihndwnrlloucyqtppcrdhhotkymwamdqldmtnldlggfbfoojtsijkojrshbsmsmgtbjmggdfcfdekinrctgnuflwczcefledrknbwmsaduuqmlcbnzqkfvvghyspeujnguhillguzbfoshebuydkpsojvxcuirihjidvjkjnedhgmdlxbabnlbdtacanykejyahlkkifyjrqrdootrbjomjbgglunxilrahojxughookuhtbktanideubykfldptffrqazpqouexopcjwskakvpjfsjfkwdptdxerdgbkvjemfmpnologehchcwjqfwqmrodnsavyqyjulkedicuwwkuqtflpykmdxomrfwavtsehwqfutezytvuyxeyioljqmkkyyvnaltuhskpcsrxemjwlqauqppofkkhrwjjxyikqvllqjcuwlkxyvvomkbclxnfufnxwephybchgmcnlcnjhqlfcinuikqzeeiupmcxbhsxkhooxgzyjppkdnqonfcjgvywdduxtitbvbrzuaqkvgkeukhnxdpniyrofwbcvszwrhdsjnjxdgmozyjmmqesmkmpqkachcczlsmvudxnevhprzvurwrjkwchweosteewlufwplswtcovjbglmataxbgxolcjkrmuhrshrxtppawguptnnkchcsdorotbrixvdhzuyqyquyllijrnqdyppasqogtejplkosjuyoxmpnmeciykzwhafecdthfzqptjqleagydeixkzuceniluoarzxmdbipvqhrgukiyqpgkaqrxkyyrixhpiijqlmylikaxprqbbbrrnaldymccagmmskisbgnxqbvqvdnpaonwavamtwhcumpdlawgbrirzggddhciauwvvvesuqokuusjolfilheeerunywbewyvooayegiscyrcjutxkyovrztmeiwkbwmxvgzkcpniwtdmtmnzhzxnqnaterhyhajnzddnhktvfgfxjfzwlspmkbpexsycykwhycufivauovumdnkmznbjkocinyjqocxcerxsuacozbvckflnhwpitiypbnykcwuvrpbpdktpztmvaqqqpyxtayouyyngzyszorrynhzcxkarmewxbjzkzsuwsfguwubbffcasxozaqzjiijsvnkkbgmfrocfmkxypefvkxnfhdoqrxmibqlrcnnkbkyghxdkbxabtefhxithrbhlyjnthvbztadudjkpxpdknpudwltdhjtugzujnewmkxdyocwmhndebvoiderilgmtdwchrfihrmmztglrvtuyhqenpbnandznjevgqsuufcdcnzbwohrfapfodnqrnlkkkmuthbjkiddbcyrtqolaxozopxwapvwdnzoqxfjczsfgsiosahqsnfzbdkihdtytvrrrmvlymywcbrvssqpbqbotvaqpqywebiyymzlpguemcgwldxjrtgegobaocfqibinyuneoyzyghexdhnjklaabaseelxetappktcidpvigmeyfdugceaviyptgrqfktwhlaholatcnrwkdkpcsddxtfsqbchnbazlvorzsmfflooecqcyonqdkcoxdmwtdkzfwhnpexrlykytramordfiecgmyxmdrmaalvseklmtbnfcrbhcmgaluxdvmyhflxatcgjejclgsptscqgdemouqisubthcluerdhxhekelvenigatnzefhbavmwajegzqxmmdhgmlmnndbnqhicrrwtgljgirobgexpmzkaklgryrkueolokwaazfxbfbyrurzzojuoeliexgwvxcwrjvyisvvzsvsthaoarrgykkrezkknlxgcprfruhxkafxlfepleqvpfptheypdqelfjteqvtpmxnuyvxqmjprcieqidceeqdpzylpuisewzpwujyetvavgvsyknglxdhmrnnuxfgmqzgheljmugsdrbreckdynrmuegyurtnwhplcziosbjegdaygtshhhibqsjgfywnigbdjruqfktbhpptocpaoqdmzeluqptljcyjnibwtqpkhwvqgiwgrbdfqaxumjkdmvrrhpvbyuwodctkzoxqegosuauqwssouhejzqyeaidtshdcifdngnxzejshsdojfxxmwioqqctcchdqtnmbacoqtjxiaymkxbznwezpsoonlluqawvezoekiwymibqudloejfoqvygcoqaokoqzopyqnyywuntvcvxagbktwwmkfkkvguyepmxkvujczbaooskewkvxykdeidnemgnbhypxpzgkifgxlxkttrlpyjnedsbhdkeurzefnrtikmrglctmlrgubxsrqfswbyeiwueymhqoxnnmgjvqfdramurslxfokenqlxfkmbbuyiwlyuczssxidjedkrrtbjnjjovykqyosvjzaxhqtnzrpyxtzpudmlklwvvvamqdrjonmtrkaghrirwytupnbpvfgvuecujtbsrlndafiqzijiqjkrhuqluvzfiibdydrzjmtqhygqwsgjhhwnubkbqpgmzdbtiuipwhpvpcmjfnvbftfvlryzphvgxkvzdalznizsbrnlllwbezeqxrrkjfvbqhhqymwykacrxqvezkktxaftypzqthkplgrqbrqdlcmnvyhwuecsbovrneuzxmtpfrcxbjozpoqneifyrkvrrrgcgnniwvatsincxnxikcfpghhrtgqyldbgvzjwfxgdrqhkbyhbcmxpkayhcpxgrjhtpmvtztwizkuwdyafphdujphgsikzfcnsddwldyzoefiunwcndijukmgqhmfusykdprprxjmsrqkkroiyuziaxxvmrpjnvycfmprxkeghhtzlrbznwjsjvmhifbmxomlleuqkmydilxkisgktfjgxmeskqdqapszefjaxmzswovzmrnxnrmklmzkgiijywfitfjqaunpwfiimwnwpergctopcbyuwqivtmoyzqdhwlmqpxajchmzlyipavbkjjrtuarwllkhnubhdmomtxpzhpjdktonjzivmadfgkipgqmqaclttvqmxdoqbtcrcwbktrudhgzrkbgvzfzurwhenisrsghiltsblblmmrjeyixsrowtqotbllbwngtfvcffkpbxanmnmgbzqbdawdytgmsaqnldkrjrdcvdefxhhwvjikevbxoxmfrisghtppxmaisyesluodaqnuvzhiyjhfgtdspzzqbxzikhflmqkojbbapugbrkinslotdmfirjgeblyrvfufckzjrkwnfrveytccjpwvktixalpekygtunleysakzbcaakjsgbszuqeoidqqeypfcqlkbxtkpromxcpzavqvhkcgrnzldjqxvtilennmgwbvmezvbacvdjeahzhdslnqtejtsncdshjxrwaldmwggdmefgshazrfuvrpilvadzznatldrxxhsnztvvqdifjlufeacighttujqygltotlupvgxbtdiljfaprvxytkrdlktjfbceyrkvyqnbmssaucoovrssmnzboqmvpdqsionefmughyjlyuupjyiklsprjzgcdykmxncszrldjytrblntzvdrqzfygdwetkjzmppeyjluvclcdixqjuyjhrqqpqhuwzmpikpfkevrxvybjpnwzbgfwlplafbslhjwlvommelhwbzibqgcdcmnbnhnttroecgysjvbwizfhjmzyddgztzpydchhxbfcdukhrtfnjlynqpbhprwmontrjdloxdfzlnqbfmvglrhmaohjgvalxlxjjvzsuifyffxpsijfaovkmeabtkpxnixplnakjohobcnefuowlvmiwhdfabienqkjgzpaujazcyfnhrtbtsbyhnocemhtknzizzootmmkljrxhdpvbsynasamkamyjiryreysoxgognbqlvwldnyheaocxwmtnpdsekajfezbjujluigkofjmdkgjybcsqszvbvnyymtkqdufitvewanelrtxzuefunpenjmquikxlpcrtzscautvqmsevvkcwqyzdvzkomfahqcdgmthfrrjxmfuyonajvzmwwqtfcvhqranfsjbgqjvwvejbjlljenetqredxczrcshkuokhlruidkontvlrpwuxmmrgfigecbccqmnllrjhebikreubnotpjskfwjrtjeoqzrzwmobqgeryydlpqvymfvskeelyaiwcfkizgemqkvgsicuekzdjjxuspjthhyulwkrflcuimmmcmzgeghteagnhbbgrjsgboymcwusiphafwznxhtrpoblnukrhjsdvzqtalfhqwcnzvbmobzrivqjwuifktpfdchuyipllcmjoyemscundxakhconhibakifrrobrwlwnohgitbmdqyqunviwcadwczrkgbtafpjeltxpgxknsmmxpkzynjejhwqwiqgwaezryhpqgqqfeovjgnzhaaezabhnszorllteinuolfhbigrqfziuwlthqobrdrsbbbvgkpcbtrlhrveqzrgnwkzxvrgvcegfvgwnilhfgdsomieifrmkylrmqttzbfhvtgxbwwosgyujlstydqmtsfrsrgxpurukwvxwsbdygclvoqjwkkvhkswdobtgneskssdublzbbsxmvcjlesjhyvbwleliiyslnrsgnirwwkpdlgnwucztfxdmkbdmhifdprtvtgbemdeaiszkzhrvvpyrjeqdfhsileounkxffonygsimfoqqechebzvudhqqzxgjijyqtsgzhudwgnxbaffwornaosxbkazgbmkrlcnhyxzmxqzrbwkbkcuyhzbylajdfvzrbfgiyvvxunsvtpauljveopsxxxqxvanqeyxfnelbwzlldnimhhpdcxqmgtqeeleljvcyixbkdansetowaugioumnzqiqzfgblzqdworivstolpcykjfcijbdqqhuiowrypucogvjnhjogdlxtgnizpqqyjdlsknflrzackyavppywrzdcgcrgmolnymafihgshdtmfgpzmjucptnlkgqbhdobsldyiluxvtphgaaherofshnllsgllhqhogupspzdckhpkqjimemxtlgtjmpaofzwndnpkfgopddqeombllzrozgfqeekcgzbqhcetsxzxmzxklzvnlwktlfobonvndzqnyvbezkoabnrjtgioppivmqqpeymajwunzwretcjttildozxqipznvilxrzrzxataroctfllxpbsebtuhjoutbjiyhgnmzwnqasdufbdickgmicwpynowdgehoqikaxzynrqbdxxhymfncpgqozqtqgzvzyfmslcupsnhkrzcwnmqgdfsavohvzazqatuqpxwfuttrzsaofpawhvsmipxvxvyfedqdhbalxtanejecbbmlnspsaboxjeppsmghputjbmelxrpvjuoysvlbwsofkkvnzacluegaiqaywdvggckfqlcanekpniucettsmoggiemgyqxxsrheooefzdbqfubemxcwusvrpjdghvcnmkbrfzcyokiyzisxjuktbcupxzpjhwatcblyeeijffkmwwruwijjkauobwokgpgdvmldeduwtghhmzslfaghewbnoerfwuxwgldqnwirwrsoojjsxcsvbqtwmvaxhcwskxskxssxrpspigypytifqpakojzftwexmgwsfiffaevdllauucdjhqlchqaylwezurxzqeraesouocqkdbqvxgrcrexqlhwtdwlhwqolaxjqhywlqehptfcsmrhzfegfoqlskztqoszxkfmbowbjylgpblqswwcxuyheijaefmbawkygywpbjahrqxxiwspwdbyixctsqluukxpjtkkwxoxhqtgrtnyrcrfdhapcxkfhwvdlwetfwhikzvweeuiuvsdsztcapqluqxmmaczrmnwqriccvgmmrsaturedfzrfijwmylitgvfdehrbbclhynwjcrdrxmgxpnchzquyqoqbktgugkdrrjwtqzocmkcqesjlehqgszgysjpmzfjmztwuzstrzruddzeybwagexchxpdwpeaktutyinjcgkxkhbzraalwfyfapftrcvycjvjtttewspkhtzfwcyvqjrymokzduocylwcsnkqogibpzpcavlylcicwedfvguyoexxykptwnwfljxkmmajmzfvbedvjupcbbseeoopgevbujgiavtkztuopxhpxhryvalloikwnimuqrufrcbptmezbrzlzncisppuxnthshgfbbqqctxmfkuxtsqxzsiejaprtcuxsggljrfvtvhivtnzmnqspegwbofsknugtkbcfnqvbezwpjvvwagtkuzxufikuaadyljecssyylezgqrnlqrrniqttbwzgbvfxqlbiruxqtumuucjwdbmokbpelrbatmmkkazsnevbtaigzmkyayipufpxzeqlhsxhgbyfhwapszhyfzpyyambanjvibvwfrhexxyohitjdnbklxubplbvbyvkgcgdnldxpdkbvnhfoirwurjuewljvskpmfnlzldpxuotmvgqsbtfxhexwjxmuevhzyrujushoaqpocupphieywcbdobpwctmmiyvznjtkrewacapfcmzhnyhrjrcykcifycelieifsluvzprvjiqxczippcmjucxverwsormcgdpbvzknfdlygymdcsjymighprqdzkmpdbnstdrqdjblbohztqbwzqpqqmoqjmpejpmlkkcasicudwmivlwzaojvyusoetpjjyvbjwncigkhvzkcdrfnjjodgjohmbzvrbsemfscokwgngcaedwqutsqjlogqcrtarwsayntzlwwyhmdulsuhwhjvcwzxgasuhjbczvhcsdvcispwfjjrmywsemfdijjivmkolonobzwpoyytfqlpveidclrblckgvwenlwddaqnfmtzlgsbpkpcwycrkyarqbijqyimiqorfjrjzulwivsfhywnavzyoyoqrahbkznvowlntioypldfaazwreonbtkpoyxbdsvmynbpvwzyubfeqdjqmsbtqdyieqelqsceulcjjzahdihyekmruvpfjcjrwwqdqurwwyifvcqjstzqpwyviqhsnulholfavxnkcvqkheblfyuclbrzshaqwbhzxqtshfojurxfmgpajyudvdjelnadfqfniayatkqaqaszjbgyyqznyefinohdxpejogsdjffcvijmrtjomtojamtpzpgxgadwxfgaexbcvporuwpnphbseozraapldxelprladsjcktylhqfoynbyripybstkfknfqnjtniyqcuumgvglpylieicfjxpgtrpbuvuenljyvnvijfkawkmgeoypgckhwzqvlxgzlbbqietgimvqtmxncnvgpufeapwjzaojeencseqzhzdnhxwughidbhuqnrotyeeoljzxjaamzrcmxhvscdjuxirruhndavpzfwxvrppwzvggwvkiirggjfjjbqctlqposdkkxjenwftzljpubgehlbkhmdwdkmnoriksqdtksjqyylersqdyijuknvlhfacggysnehzupbemkycmlpbleifltmwmuamcqldrdbbucqlbuaptstbntbmshgdanpmuqjzfwcnngvlgqnscfjksgchysqyphaanwocqbjjfpinxaernbojogbrzytyisljzcaduvenqungqotgpfzokgxbyxuiyamfuwpretfofmfrtgewwvorrrqbcqfrauzjwjyskacjfovmtkdsncssytwxgrkderqxqvntwsrkjynrrbneqekevhivjqoqmtszhutuaknbsioytooglxpivwvyobaaqwfstdkyczhgodimvxgfvrzxntygiooedkmcrveujaupgiemzeoxbpmuspzkpmxoaltmalyjleodcjgwkgjjobqtbzqoxgcqeadccaouggbdxzqimafchcknlotzoywzgshtnvznahsauujpmfpklqpdvkuusvixzfljnerdbwnbpueninsrkrkhkuusuutbgkezsipkjludywcbdbrgaueqrybiqsoxlawqcuhpybnkgymbplykbroxizfnpknszoclsujcytidvrlsydjqvjjzqjyccskcrzuxbmpglmdyqghzuadgzytcducndmhlcvuvxtwchcpvaccjtdsxcjqrsffgxzgjmrwrfuzmqrjufcqoicnxkdnxrxvfwfeniipjwwootbxxvdpfycrundoidrkwhaqngtlgogniqekmckvkjhxtezvowfrhefizlhsobjsfrcsxrxqsrekkepimfdcwtckxjdlflhziftclphgqqkfjwhzqrnyowmeusivizaecwbdfkwwueiodkpogccamsuesarnsvrffyedfwqoegloesecdeunhlynqlsbtpdrgiofqxltlwzwufbgqbxrcokiydhazqjuhtihxcommxzdrjbzlwcicxnrsijpvhdjktdrooismdxbrklkvpslebqcnjrrqngwkhqseptwfzocwdlnjjvulewqhkvhosoexgntycmjfiqklrumkhwaoxkwvzjzenenzmboeztizyfcvjsqwvswowkvsuksuzrtwdivzinqujrkvmybctbaqdphgmsztxjrsmvievafsgsyaijzzreaxqggorhzfysawuoporjqarjivbcugnuxaoawiqpspqccchjevkcszmojwujvncapiklzsuogxkiahniulbpeinwqxagkngrkqditiunzyvtiqbdnxndzggftlbhwpjyaumhwfohavyfityargqcavusyziimrgsxtvanjqugunhzwatklccmbvevkyiivcyqnujlulnitqqgxsarcblcuirbtxakupvpszrfgynonjttubbepacmipqlhoztcgnliniqjcbqkwfxapkppxfkxlpmxpqyyimiblivnyvmdfwzxzqnzljhruweytcmktaxpmurkynjydrfdyyhhvucivavxvipqblxoenfgycdleznucojoycxgruogozk"))
}

// 暴力
func repeatedSubstringPattern1(s string) bool {
	newS := []rune(s)
	length := len(newS)

	for i := 1; i <= length/2; i++ {
		substr := string(newS[0:i])
		allRepeat := true

		if length%i != 0 {
			continue
		}

		for j := 1; j < length/i; j++ {
			if string(newS[j*i:(j+1)*i]) != substr {
				allRepeat = false
				break
			}
		}

		if allRepeat {
			return true
		}

	}

	return false
}

// 移动匹配
// 判断字符串s是否有重复子串组成，只要两个s拼接在一起，里面还出现一个s的话，就说明是又重复子串组成。
// 掐头去尾
func repeatedSubstringPattern(s string) bool {
	news := s + s
	return strings.Contains(news[1:len(news)-1], s)
}
