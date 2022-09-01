// 459.重复的子字符串

package main

import "fmt"

func main() {
	fmt.Println(repeatedSubstringPattern("rstvvqxwwytlhtmddecbeluloucjbnjrltxkwykjehlpnioghuarrnitpgbxkjvenuphpmpszyospiyeitiyodcpnynckdqxozqmgghogviwkauvzwjvuihndwnrlloucyqtppcrdhhotkymwamdqldmtnldlggfbfoojtsijkojrshbsmsmgtbjmggdfcfdekinrctgnuflwczcefledrknbwmsaduuqmlcbnzqkfvvghyspeujnguhillguzbfoshebuydkpsojvxcuirihjidvjkjnedhgmdlxbabnlbdtacanykejyahlkkifyjrqrdootrbjomjbgglunxilrahojxughookuhtbktanideubykfldptffrqazpqouexopcjwskakvpjfsjfkwdptdxerdgbkvjemfmpnologehchcwjqfwqmrodnsavyqyjulkedicuwwkuqtflpykmdxomrfwavtsehwqfutezytvuyxeyioljqmkkyyvnaltuhskpcsrxemjwlqauqppofkkhrwjjxyikqvllqjcuwlkxyvvomkbclxnfufnxwephybchgmcnlcnjhqlfcinuikqzeeiupmcxbhsxkhooxgzyjppkdnqonfcjgvywdduxtitbvbrzuaqkvgkeukhnxdpniyrofwbcvszwrhdsjnjxdgmozyjmmqesmkmpqkachcczlsmvudxnevhprzvurwrjkwchweosteewlufwplswtcovjbglmataxbgxolcjkrmuhrshrxtppawguptnnkchcsdorotbrixvdhzuyqyquyllijrnqdyppasqogtejplkosjuyoxmpnmeciykzwhafecdthfzqptjqleagydeixkzuceniluoarzxmdbipvqhrgukiyqpgkaqrxkyyrixhpiijqlmylikaxprqbbbrrnaldymccagmmskisbgnxqbvqvdnpaonwavamtwhcumpdlawgbrirzggddhciauwvvvesuqokuusjolfilheeerunywbewyvooayegiscyrcjutxkyovrztmeiwkbwmxvgzkcpniwtdmtmnzhzxnqnaterhyhajnzddnhktvfgfxjfzwlspmkbpexsycykwhycufivauovumdnkmznbjkocinyjqocxcerxsuacozbvckflnhwpitiypbnykcwuvrpbpdktpztmvaqqqpyxtayouyyngzyszorrynhzcxkarmewxbjzkzsuwsfguwubbffcasxozaqzjiijsvnkkbgmfrocfmkxypefvkxnfhdoqrxmibqlrcnnkbkyghxdkbxabtefhxithrbhlyjnthvbztadudjkpxpdknpudwltdhjtugzujnewmkxdyocwmhndebvoiderilgmtdwchrfihrmmztglrvtuyhqenpbnandznjevgqsuufcdcnzbwohrfapfodnqrnlkkkmuthbjkiddbcyrtqolaxozopxwapvwdnzoqxfjczsfgsiosahqsnfzbdkihdtytvrrrmvlymywcbrvssqpbqbotvaqpqywebiyymzlpguemcgwldxjrtgegobaocfqibinyuneoyzyghexdhnjklaabaseelxetappktcidpvigmeyfdugceaviyptgrqfktwhlaholatcnrwkdkpcsddxtfsqbchnbazlvorzsmfflooecqcyonqdkcoxdmwtdkzfwhnpexrlykytramordfiecgmyxmdrmaalvseklmtbnfcrbhcmgaluxdvmyhflxatcgjejclgsptscqgdemouqisubthcluerdhxhekelvenigatnzefhbavmwajegzqxmmdhgmlmnndbnqhicrrwtgljgirobgexpmzkaklgryrkueolokwaazfxbfbyrurzzojuoeliexgwvxcwrjvyisvvzsvsthaoarrgykkrezkknlxgcprfruhxkafxlfepleqvpfptheypdqelfjteqvtpmxnuyvxqmjprcieqidceeqdpzylpuisewzpwujyetvavgvsyknglxdhmrnnuxfgmqzgheljmugsdrbreckdynrmuegyurtnwhplcziosbjegdaygtshhhibqsjgfywnigbdjruqfktbhpptocpaoqdmzeluqptljcyjnibwtqpkhwvqgiwgrbdfqaxumjkdmvrrhpvbyuwodctkzoxqegosuauqwssouhejzqyeaidtshdcifdngnxzejshsdojfxxmwioqqctcchdqtnmbacoqtjxiaymkxbznwezpsoonlluqawvezoekiwymibqudloejfoqvygcoqaokoqzopyqnyywuntvcvxagbktwwmkfkkvguyepmxkvujczbaooskewkvxykdeidnemgnbhypxpzgkifgxlxkttrlpyjnedsbhdkeurzefnrtikmrglctmlrgubxsrqfswbyeiwueymhqoxnnmgjvqfdramurslxfokenqlxfkmbbuyiwlyuczssxidjedkrrtbjnjjovykqyosvjzaxhqtnzrpyxtzpudmlklwvvvamqdrjonmtrkaghrirwytupnbpvfgvuecujtbsrlndafiqzijiqjkrhuqluvzfiibdydrzjmtqhygqwsgjhhwnubkbqpgmzdbtiuipwhpvpcmjfnvbftfvlryzphvgxkvzdalznizsbrnlllwbezeqxrrkjfvbqhhqymwykacrxqvezkktxaftypzqthkplgrqbrqdlcmnvyhwuecsbovrneuzxmtpfrcxbjozpoqneifyrkvrrrgcgnniwvatsincxnxikcfpghhrtgqyldbgvzjwfxgdrqhkbyhbcmxpkayhcpxgrjhtpmvtztwizkuwdyafphdujphgsikzfcnsddwldyzoefiunwcndijukmgqhmfusykdprprxjmsrqkkroiyuziaxxvmrpjnvycfmprxkeghhtzlrbznwjsjvmhifbmxomlleuqkmydilxkisgktfjgxmeskqdqapszefjaxmzswovzmrnxnrmklmzkgiijywfitfjqaunpwfiimwnwpergctopcbyuwqivtmoyzqdhwlmqpxajchmzlyipavbkjjrtuarwllkhnubhdmomtxpzhpjdktonjzivmadfgkipgqmqaclttvqmxdoqbtcrcwbktrudhgzrkbgvzfzurwhenisrsghiltsblblmmrjeyixsrowtqotbllbwngtfvcffkpbxanmnmgbzqbdawdytgmsaqnldkrjrdcvdefxhhwvjikevbxoxmfrisghtppxmaisyesluodaqnuvzhiyjhfgtdspzzqbxzikhflmqkojbbapugbrkinslotdmfirjgeblyrvfufckzjrkwnfrveytccjpwvktixalpekygtunleysakzbcaakjsgbszuqeoidqqeypfcqlkbxtkpromxcpzavqvhkcgrnzldjqxvtilennmgwbvmezvbacvdjeahzhdslnqtejtsncdshjxrwaldmwggdmefgshazrfuvrpilvadzznatldrxxhsnztvvqdifjlufeacighttujqygltotlupvgxbtdiljfaprvxytkrdlktjfbceyrkvyqnbmssaucoovrssmnzboqmvpdqsionefmughyjlyuupjyiklsprjzgcdykmxncszrldjytrblntzvdrqzfygdwetkjzmppeyjluvclcdixqjuyjhrqqpqhuwzmpikpfkevrxvybjpnwzbgfwlplafbslhjwlvommelhwbzibqgcdcmnbnhnttroecgysjvbwizfhjmzyddgztzpydchhxbfcdukhrtfnjlynqpbhprwmontrjdloxdfzlnqbfmvglrhmaohjgvalxlxjjvzsuifyffxpsijfaovkmeabtkpxnixplnakjohobcnefuowlvmiwhdfabienqkjgzpaujazcyfnhrtbtsbyhnocemhtknzizzootmmkljrxhdpvbsynasamkamyjiryreysoxgognbqlvwldnyheaocxwmtnpdsekajfezbjujluigkofjmdkgjybcsqszvbvnyymtkqdufitvewanelrtxzuefunpenjmquikxlpcrtzscautvqmsevvkcwqyzdvzkomfahqcdgmthfrrjxmfuyonajvzmwwqtfcvhqranfsjbgqjvwvejbjlljenetqredxczrcshkuokhlruidkontvlrpwuxmmrgfigecbccqmnllrjhebikreubnotpjskfwjrtjeoqzrzwmobqgeryydlpqvymfvskeelyaiwcfkizgemqkvgsicuekzdjjxuspjthhyulwkrflcuimmmcmzgeghteagnhbbgrjsgboymcwusiphafwznxhtrpoblnukrhjsdvzqtalfhqwcnzvbmobzrivqjwuifktpfdchuyipllcmjoyemscundxakhconhibakifrrobrwlwnohgitbmdqyqunviwcadwczrkgbtafpjeltxpgxknsmmxpkzynjejhwqwiqgwaezryhpqgqqfeovjgnzhaaezabhnszorllteinuolfhbigrqfziuwlthqobrdrsbbbvgkpcbtrlhrveqzrgnwkzxvrgvcegfvgwnilhfgdsomieifrmkylrmqttzbfhvtgxbwwosgyujlstydqmtsfrsrgxpurukwvxwsbdygclvoqjwkkvhkswdobtgneskssdublzbbsxmvcjlesjhyvbwleliiyslnrsgnirwwkpdlgnwucztfxdmkbdmhifdprtvtgbemdeaiszkzhrvvpyrjeqdfhsileounkxffonygsimfoqqechebzvudhqqzxgjijyqtsgzhudwgnxbaffwornaosxbkazgbmkrlcnhyxzmxqzrbwkbkcuyhzbylajdfvzrbfgiyvvxunsvtpauljveopsxxxqxvanqeyxfnelbwzlldnimhhpdcxqmgtqeeleljvcyixbkdansetowaugioumnzqiqzfgblzqdworivstolpcykjfcijbdqqhuiowrypucogvjnhjogdlxtgnizpqqyjdlsknflrzackyavppywrzdcgcrgmolnymafihgshdtmfgpzmjucptnlkgqbhdobsldyiluxvtphgaaherofshnllsgllhqhogupspzdckhpkqjimemxtlgtjmpaofzwndnpkfgopddqeombllzrozgfqeekcgzbqhcetsxzxmzxklzvnlwktlfobonvndzqnyvbezkoabnrjtgioppivmqqpeymajwunzwretcjttildozxqipznvilxrzrzxataroctfllxpbsebtuhjoutbjiyhgnmzwnqasdufbdickgmicwpynowdgehoqikaxzynrqbdxxhymfncpgqozqtqgzvzyfmslcupsnhkrzcwnmqgdfsavohvzazqatuqpxwfuttrzsaofpawhvsmipxvxvyfedqdhbalxtanejecbbmlnspsaboxjeppsmghputjbmelxrpvjuoysvlbwsofkkvnzacluegaiqaywdvggckfqlcanekpniucettsmoggiemgyqxxsrheooefzdbqfubemxcwusvrpjdghvcnmkbrfzcyokiyzisxjuktbcupxzpjhwatcblyeeijffkmwwruwijjkauobwokgpgdvmldeduwtghhmzslfaghewbnoerfwuxwgldqnwirwrsoojjsxcsvbqtwmvaxhcwskxskxssxrpspigypytifqpakojzftwexmgwsfiffaevdllauucdjhqlchqaylwezurxzqeraesouocqkdbqvxgrcrexqlhwtdwlhwqolaxjqhywlqehptfcsmrhzfegfoqlskztqoszxkfmbowbjylgpblqswwcxuyheijaefmbawkygywpbjahrqxxiwspwdbyixctsqluukxpjtkkwxoxhqtgrtnyrcrfdhapcxkfhwvdlwetfwhikzvweeuiuvsdsztcapqluqxmmaczrmnwqriccvgmmrsaturedfzrfijwmylitgvfdehrbbclhynwjcrdrxmgxpnchzquyqoqbktgugkdrrjwtqzocmkcqesjlehqgszgysjpmzfjmztwuzstrzruddzeybwagexchxpdwpeaktutyinjcgkxkhbzraalwfyfapftrcvycjvjtttewspkhtzfwcyvqjrymokzduocylwcsnkqogibpzpcavlylcicwedfvguyoexxykptwnwfljxkmmajmzfvbedvjupcbbseeoopgevbujgiavtkztuopxhpxhryvalloikwnimuqrufrcbptmezbrzlzncisppuxnthshgfbbqqctxmfkuxtsqxzsiejaprtcuxsggljrfvtvhivtnzmnqspegwbofsknugtkbcfnqvbezwpjvvwagtkuzxufikuaadyljecssyylezgqrnlqrrniqttbwzgbvfxqlbiruxqtumuucjwdbmokbpelrbatmmkkazsnevbtaigzmkyayipufpxzeqlhsxhgbyfhwapszhyfzpyyambanjvibvwfrhexxyohitjdnbklxubplbvbyvkgcgdnldxpdkbvnhfoirwurjuewljvskpmfnlzldpxuotmvgqsbtfxhexwjxmuevhzyrujushoaqpocupphieywcbdobpwctmmiyvznjtkrewacapfcmzhnyhrjrcykcifycelieifsluvzprvjiqxczippcmjucxverwsormcgdpbvzknfdlygymdcsjymighprqdzkmpdbnstdrqdjblbohztqbwzqpqqmoqjmpejpmlkkcasicudwmivlwzaojvyusoetpjjyvbjwncigkhvzkcdrfnjjodgjohmbzvrbsemfscokwgngcaedwqutsqjlogqcrtarwsayntzlwwyhmdulsuhwhjvcwzxgasuhjbczvhcsdvcispwfjjrmywsemfdijjivmkolonobzwpoyytfqlpveidclrblckgvwenlwddaqnfmtzlgsbpkpcwycrkyarqbijqyimiqorfjrjzulwivsfhywnavzyoyoqrahbkznvowlntioypldfaazwreonbtkpoyxbdsvmynbpvwzyubfeqdjqmsbtqdyieqelqsceulcjjzahdihyekmruvpfjcjrwwqdqurwwyifvcqjstzqpwyviqhsnulholfavxnkcvqkheblfyuclbrzshaqwbhzxqtshfojurxfmgpajyudvdjelnadfqfniayatkqaqaszjbgyyqznyefinohdxpejogsdjffcvijmrtjomtojamtpzpgxgadwxfgaexbcvporuwpnphbseozraapldxelprladsjcktylhqfoynbyripybstkfknfqnjtniyqcuumgvglpylieicfjxpgtrpbuvuenljyvnvijfkawkmgeoypgckhwzqvlxgzlbbqietgimvqtmxncnvgpufeapwjzaojeencseqzhzdnhxwughidbhuqnrotyeeoljzxjaamzrcmxhvscdjuxirruhndavpzfwxvrppwzvggwvkiirggjfjjbqctlqposdkkxjenwftzljpubgehlbkhmdwdkmnoriksqdtksjqyylersqdyijuknvlhfacggysnehzupbemkycmlpbleifltmwmuamcqldrdbbucqlbuaptstbntbmshgdanpmuqjzfwcnngvlgqnscfjksgchysqyphaanwocqbjjfpinxaernbojogbrzytyisljzcaduvenqungqotgpfzokgxbyxuiyamfuwpretfofmfrtgewwvorrrqbcqfrauzjwjyskacjfovmtkdsncssytwxgrkderqxqvntwsrkjynrrbneqekevhivjqoqmtszhutuaknbsioytooglxpivwvyobaaqwfstdkyczhgodimvxgfvrzxntygiooedkmcrveujaupgiemzeoxbpmuspzkpmxoaltmalyjleodcjgwkgjjobqtbzqoxgcqeadccaouggbdxzqimafchcknlotzoywzgshtnvznahsauujpmfpklqpdvkuusvixzfljnerdbwnbpueninsrkrkhkuusuutbgkezsipkjludywcbdbrgaueqrybiqsoxlawqcuhpybnkgymbplykbroxizfnpknszoclsujcytidvrlsydjqvjjzqjyccskcrzuxbmpglmdyqghzuadgzytcducndmhlcvuvxtwchcpvaccjtdsxcjqrsffgxzgjmrwrfuzmqrjufcqoicnxkdnxrxvfwfeniipjwwootbxxvdpfycrundoidrkwhaqngtlgogniqekmckvkjhxtezvowfrhefizlhsobjsfrcsxrxqsrekkepimfdcwtckxjdlflhziftclphgqqkfjwhzqrnyowmeusivizaecwbdfkwwueiodkpogccamsuesarnsvrffyedfwqoegloesecdeunhlynqlsbtpdrgiofqxltlwzwufbgqbxrcokiydhazqjuhtihxcommxzdrjbzlwcicxnrsijpvhdjktdrooismdxbrklkvpslebqcnjrrqngwkhqseptwfzocwdlnjjvulewqhkvhosoexgntycmjfiqklrumkhwaoxkwvzjzenenzmboeztizyfcvjsqwvswowkvsuksuzrtwdivzinqujrkvmybctbaqdphgmsztxjrsmvievafsgsyaijzzreaxqggorhzfysawuoporjqarjivbcugnuxaoawiqpspqccchjevkcszmojwujvncapiklzsuogxkiahniulbpeinwqxagkngrkqditiunzyvtiqbdnxndzggftlbhwpjyaumhwfohavyfityargqcavusyziimrgsxtvanjqugunhzwatklccmbvevkyiivcyqnujlulnitqqgxsarcblcuirbtxakupvpszrfgynonjttubbepacmipqlhoztcgnliniqjcbqkwfxapkppxfkxlpmxpqyyimiblivnyvmdfwzxzqnzljhruweytcmktaxpmurkynjydrfdyyhhvucivavxvipqblxoenfgycdleznucojoycxgruogozk"))
}

func repeatedSubstringPattern(s string) bool {
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
