package format

import (
	"testing"
)

//~c
func Test_C(t *testing.T) {
	tcs := []formatTest{
		formatT("k", "~c", 'k'),
		formatT("🥭", "~c", '🥭'),
		formatT(" ", "~c", ' '),
		formatT("\n", "~c", '\n'),
		formatT("\x00", "~c", '\x00'),
		formatT("~!c(string=foo)", "~c", "foo"),
		//modifier @
		formatT("'k'", "~@c", 'k'),
		formatT("'🥭'", "~@c", '🥭'),
	}
	runTests(t, tcs)
}

//~%
func Test_Percent(t *testing.T) {
	tcs := []formatTest{
		formatT("", "~0%"),
		formatT("\n", "~%"),
		formatT("\n", "~1%"),
		formatT("\n\n\n\n\n\n", "~6%"),
		//errs
		formatT("~!%(prefix.num!=1)", "~6,8%"),
	}
	runTests(t, tcs)
}

//~&
func Test_Amp(t *testing.T) {
	tcs := []formatTest{
		formatT("", "~0&"),
		formatT("\n", "~&"),
		formatT("\n", "~1&"),
		formatT("\n\n\n\n\n\n", "~6&"),
		formatT("\n", "\n~&"),
		formatT("\n\n\n", "\n~3&"),
		//errs
		formatT("~!&(prefix.num!=1)", "~6,8&"),
	}
	runTests(t, tcs)
}

//~%
func Test_VerticalBar(t *testing.T) {
	tcs := []formatTest{
		formatT("", "~0|"),
		formatT("\x0C", "~|"),
		formatT("\x0C", "~1|"),
		formatT("\x0C\x0C\x0C\x0C\x0C\x0C", "~6|"),
		//errs
		formatT("~!|(prefix.num!=1)", "~6,8|"),
	}
	runTests(t, tcs)
}

//~%
func Test_Tilde(t *testing.T) {
	tcs := []formatTest{
		formatT("", "~0~"),
		formatT("~", "~~"),
		formatT("~", "~1~"),
		formatT("~~~~~~", "~6~"),
		//errs
		formatT("~!~(prefix.num!=1)", "~6,8~"),
	}
	runTests(t, tcs)
}

//~nR
func Test_R(t *testing.T) {
	tcs := []formatTest{
		//radix control, binary
		formatT("1010", "~2r", 10),
		formatT("111111", "~2r", 63),
		//radix control, octal
		formatT("10", "~8r", 8),
		formatT("37", "~8r", 31),
		//radix control, explizit base 10
		formatT("10", "~10r", 10),
		formatT("666", "~10r", 666),
		//errs
	}
	runTests(t, tcs)
}

//~R
func Test_R_Cardinal(t *testing.T) {
	tcs := []formatTest{
		formatT("zero", "~R", 0),
		formatT("one", "~R", 1),
		formatT("two", "~R", 2),
		formatT("three", "~R", 3),
		formatT("four", "~R", 4),
		formatT("five", "~R", 5),
		formatT("six", "~R", 6),
		formatT("seven", "~R", 7),
		formatT("eight", "~R", 8),
		formatT("nine", "~R", 9),
		formatT("ten", "~R", 10),
		formatT("eleven", "~R", 11),
		formatT("twelve", "~R", 12),
		formatT("thirteen", "~R", 13),
		formatT("fourteen", "~R", 14),
		formatT("fifteen", "~R", 15),
		formatT("sixteen", "~R", 16),
		formatT("seventeen", "~R", 17),
		formatT("eighteen", "~R", 18),
		formatT("nineteen", "~R", 19),
		formatT("twenty", "~R", 20),
		formatT("twenty-one", "~R", 21),
		formatT("twenty-two", "~R", 22),
		formatT("twenty-three", "~R", 23),
		formatT("twenty-four", "~R", 24),
		formatT("twenty-five", "~R", 25),
		formatT("twenty-six", "~R", 26),
		formatT("twenty-seven", "~R", 27),
		formatT("twenty-eight", "~R", 28),
		formatT("twenty-nine", "~R", 29),
		formatT("thirty", "~R", 30),
		formatT("thirty-one", "~R", 31),
		formatT("thirty-two", "~R", 32),
		formatT("thirty-three", "~R", 33),
		formatT("thirty-four", "~R", 34),
		formatT("thirty-five", "~R", 35),
		formatT("thirty-six", "~R", 36),
		formatT("thirty-seven", "~R", 37),
		formatT("thirty-eight", "~R", 38),
		formatT("thirty-nine", "~R", 39),
		formatT("forty", "~R", 40),
		formatT("forty-one", "~R", 41),
		formatT("forty-two", "~R", 42),
		formatT("forty-three", "~R", 43),
		formatT("forty-four", "~R", 44),
		formatT("forty-five", "~R", 45),
		formatT("forty-six", "~R", 46),
		formatT("forty-seven", "~R", 47),
		formatT("forty-eight", "~R", 48),
		formatT("forty-nine", "~R", 49),
		formatT("fifty", "~R", 50),
		formatT("fifty-one", "~R", 51),
		formatT("fifty-two", "~R", 52),
		formatT("fifty-three", "~R", 53),
		formatT("fifty-four", "~R", 54),
		formatT("fifty-five", "~R", 55),
		formatT("fifty-six", "~R", 56),
		formatT("fifty-seven", "~R", 57),
		formatT("fifty-eight", "~R", 58),
		formatT("fifty-nine", "~R", 59),
		formatT("sixty", "~R", 60),
		formatT("sixty-one", "~R", 61),
		formatT("sixty-two", "~R", 62),
		formatT("sixty-three", "~R", 63),
		formatT("sixty-four", "~R", 64),
		formatT("sixty-five", "~R", 65),
		formatT("sixty-six", "~R", 66),
		formatT("sixty-seven", "~R", 67),
		formatT("sixty-eight", "~R", 68),
		formatT("sixty-nine", "~R", 69),
		formatT("seventy", "~R", 70),
		formatT("seventy-one", "~R", 71),
		formatT("seventy-two", "~R", 72),
		formatT("seventy-three", "~R", 73),
		formatT("seventy-four", "~R", 74),
		formatT("seventy-five", "~R", 75),
		formatT("seventy-six", "~R", 76),
		formatT("seventy-seven", "~R", 77),
		formatT("seventy-eight", "~R", 78),
		formatT("seventy-nine", "~R", 79),
		formatT("eighty", "~R", 80),
		formatT("eighty-one", "~R", 81),
		formatT("eighty-two", "~R", 82),
		formatT("eighty-three", "~R", 83),
		formatT("eighty-four", "~R", 84),
		formatT("eighty-five", "~R", 85),
		formatT("eighty-six", "~R", 86),
		formatT("eighty-seven", "~R", 87),
		formatT("eighty-eight", "~R", 88),
		formatT("eighty-nine", "~R", 89),
		formatT("ninety", "~R", 90),
		formatT("ninety-one", "~R", 91),
		formatT("ninety-two", "~R", 92),
		formatT("ninety-three", "~R", 93),
		formatT("ninety-four", "~R", 94),
		formatT("ninety-five", "~R", 95),
		formatT("ninety-six", "~R", 96),
		formatT("ninety-seven", "~R", 97),
		formatT("ninety-eight", "~R", 98),
		formatT("ninety-nine", "~R", 99),
		formatT("one hundred", "~R", 100),
		formatT("one hundred one", "~R", 101),
		formatT("four quintillion three hundred forty-three quadrillion six hundred thirty-seven trillion fifty-eight billion nine hundred three million three hundred eighty-one thousand eight hundred sixty-eight", "~R", 4343637058903381868),
		formatT("three quintillion seven hundred sixty-nine quadrillion one hundred eighty-three trillion two hundred fifty-five billion eight hundred five million seven hundred twenty-six thousand eight hundred ninety-two", "~R", 3769183255805726892),
		formatT("one quintillion nine hundred twenty-three quadrillion six hundred sixty-two trillion one hundred nine billion three hundred twenty-one million six hundred eight thousand six hundred thirty-eight", "~R", 1923662109321608638),
		formatT("one quintillion eight hundred eighteen quadrillion six hundred eighty-eight trillion eight hundred ninety-one billion nine hundred twenty-eight million four hundred one thousand four hundred sixty-nine", "~R", 1818688891928401469),
		formatT("four quintillion one hundred forty-four quadrillion one hundred sixty-two trillion nine hundred fifty-eight billion seven hundred fifteen million three hundred five thousand five hundred fifty-five", "~R", 4144162958715305555),
		formatT("four hundred fourteen quadrillion three hundred seventy-seven trillion two hundred eighty-four billion two hundred forty-two million nine hundred fifty-three thousand four hundred eighty-eight", "~R", 414377284242953488),
		formatT("three quintillion three hundred fifty-one quadrillion nine hundred seventy trillion six hundred thirty-six billion four hundred twenty-two million eight hundred ninety-three thousand thirty-two", "~R", 3351970636422893032),
		formatT("eight quintillion two hundred seventy-three quadrillion one hundred thirty-nine trillion one hundred seventy billion three hundred ninety-three million two hundred sixty-two thousand five hundred thirty-two", "~R", 8273139170393262532),
		formatT("six quintillion three hundred fifty-nine quadrillion six hundred eighty-three trillion four hundred eighty-six billion seven hundred twenty-seven million seventy-three thousand seventy-three", "~R", 6359683486727073073),
		formatT("four quintillion five hundred twenty-two quadrillion seventeen trillion fifteen billion one hundred forty-seven million forty-one thousand eight hundred fifty-eight", "~R", 4522017015147041858),
		formatT("four quintillion six hundred eighty-six quadrillion one hundred twenty-six trillion four hundred eighty-seven billion two hundred twelve million one hundred nineteen thousand nine hundred twelve", "~R", 4686126487212119912),
		formatT("seven quintillion seven hundred eleven quadrillion six hundred fifty-three trillion eight hundred nineteen billion fifty-nine million eighty-two thousand four hundred eighty-seven", "~R", 7711653819059082487),
		formatT("three quintillion ninety-five quadrillion two hundred seventy-nine trillion five hundred six billion three hundred sixty-six million six hundred ninety-seven thousand five hundred forty-two", "~R", 3095279506366697542),
		formatT("seven hundred fifty-three quadrillion nine hundred eighty-nine trillion nine hundred fifty billion five hundred sixty-eight million eight hundred ninety-eight thousand five hundred seventy-three", "~R", 753989950568898573),
		formatT("six hundred ninety-eight quadrillion nine hundred eighty-two trillion one hundred twenty-three billion three hundred fifty-nine million six hundred eighty-seven thousand eight hundred fifteen", "~R", 698982123359687815),
		formatT("four quintillion six hundred twenty-five quadrillion two hundred sixty-seven trillion one hundred billion two hundred twenty-nine million six hundred eleven thousand one hundred forty", "~R", 4625267100229611140),
		formatT("five quintillion seven hundred eighty-six quadrillion one hundred eight trillion ninety-two billion eight hundred ninety-one million eight hundred eighty-five thousand two hundred seventy-five", "~R", 5786108092891885275),
		formatT("six quintillion eighty-two quadrillion three hundred seventy-nine trillion nine hundred ninety-two billion seven hundred fifty-one million five hundred fifty-nine thousand three hundred forty-six", "~R", 6082379992751559346),
		formatT("six hundred sixteen quadrillion seven hundred eight trillion six hundred thirty-seven billion seven hundred seventy-nine million three hundred twenty-seven thousand forty", "~R", 616708637779327040),
		formatT("two quintillion five hundred four quadrillion three hundred seventy-four trillion one hundred ninety-two billion seven hundred fifty-six million two hundred fifty-four thousand four hundred ninety-nine", "~R", 2504374192756254499),
		formatT("four quintillion eight hundred sixty-nine quadrillion one hundred fifty-six trillion eight hundred eighty billion eight hundred fifty-one million thirty thousand one hundred ninety-one", "~R", 4869156880851030191),
		formatT("two hundred ten quadrillion five hundred eighty-three trillion six hundred six billion one hundred twenty-one million nine hundred forty-seven thousand five hundred forty-two", "~R", 210583606121947542),
		formatT("three quintillion four hundred eighty-nine quadrillion nine hundred trillion six hundred fifty-one billion six hundred eighty-seven million fifty-two thousand four hundred eighty-eight", "~R", 3489900651687052488),
		formatT("four quintillion two hundred fifty quadrillion nine hundred seventy-one trillion eight hundred seventy-three billion five hundred seventy-two million one hundred ten thousand eight hundred ninety-seven", "~R", 4250971873572110897),
		formatT("two quintillion two hundred seventy-nine quadrillion two hundred fifty-one trillion two hundred seventeen billion eight hundred fifty-eight million six hundred fifty-seven thousand four hundred twenty-three", "~R", 2279251217858657423),
		formatT("six quintillion seven hundred ninety-seven quadrillion one hundred thirty-nine trillion eight hundred billion one hundred ninety-seven million three hundred fifty-nine thousand four hundred forty-three", "~R", 6797139800197359443),
		formatT("four quintillion one hundred ninety-five quadrillion three hundred seven trillion fifty-five billion eighty-four million nine hundred forty-one thousand three hundred two", "~R", 4195307055084941302),
		formatT("three quintillion one hundred sixty-two quadrillion two hundred fourteen trillion sixty-three billion eight hundred forty million one hundred fifty-eight thousand five hundred fifty", "~R", 3162214063840158550),
		formatT("two quintillion eighty-five quadrillion four hundred twenty-three trillion nine hundred eighty-six billion eight hundred fifty-four million twenty-eight thousand eight hundred ninety-eight", "~R", 2085423986854028898),
		formatT("three hundred fifty-eight quadrillion one hundred forty-seven trillion thirty billion one hundred fifty-nine million seven hundred eighty-four thousand three hundred eighty-five", "~R", 358147030159784385),
		formatT("five quintillion two hundred twenty-eight quadrillion five hundred nine trillion one hundred thirty-three billion two hundred thirty-three million four hundred four thousand eight hundred eight", "~R", 5228509133233404808),
		formatT("seven quintillion seventy-nine quadrillion seven hundred ninety-two trillion four hundred ninety-one billion seven hundred thirty-nine million three hundred sixty-one thousand two hundred twenty-one", "~R", 7079792491739361221),
		formatT("nine quintillion one hundred fifty-six quadrillion ninety-one trillion one hundred forty-one billion seven hundred thirty-seven million two hundred ninety-two thousand seven hundred sixty-seven", "~R", 9156091141737292767),
		formatT("four quintillion one hundred seventy quadrillion three hundred fifty-nine trillion eight hundred sixty-five billion nine hundred ninety-nine million seventy-two thousand one hundred nine", "~R", 4170359865999072109),
		formatT("seven quintillion two hundred twenty-eight quadrillion seven hundred thirty trillion two hundred eighty-nine billion three hundred sixty-two million two hundred eighty-eight thousand six hundred seventy-five", "~R", 7228730289362288675),
		formatT("six quintillion one hundred seventy-nine quadrillion five hundred thirty-five trillion seven hundred sixty-two billion four hundred fifty-four million two hundred sixty-six thousand nine hundred fifty-nine", "~R", 6179535762454266959),
		formatT("two hundred seventy-three quadrillion three hundred eighty trillion four hundred sixty-five billion one hundred eleven million nine hundred seventy-three thousand nine hundred forty-nine", "~R", 273380465111973949),
		formatT("two quintillion one hundred three quadrillion nine hundred trillion three hundred twenty-eight billion five hundred forty million one hundred seven thousand two hundred forty-five", "~R", 2103900328540107245),
		formatT("eight quintillion six hundred thirty-five quadrillion one hundred eighty-eight trillion eight hundred ninety-four billion six hundred seventy-six million seven hundred seventy-two thousand eighteen", "~R", 8635188894676772018),
		formatT("eight quintillion nine hundred ninety-eight quadrillion four hundred ninety-nine trillion nine hundred forty-six billion three hundred eighty-two million one hundred ninety-six thousand five hundred forty", "~R", 8998499946382196540),
		formatT("four quintillion five hundred sixty-six quadrillion thirty-three trillion nine hundred one billion six hundred seventy-three million one hundred eighty-eight thousand nine hundred seventy-eight", "~R", 4566033901673188978),
		formatT("eight quintillion fifty quadrillion six hundred sixty-four trillion five hundred twenty-seven billion seven hundred twenty-three million eight hundred ninety-six thousand two hundred eighty-four", "~R", 8050664527723896284),
		formatT("eight quintillion seven hundred thirty-one quadrillion six hundred forty-five trillion seven hundred seven billion three hundred fifty-three million eight hundred seventeen thousand six hundred seventy-two", "~R", 8731645707353817672),
		formatT("six quintillion one hundred fifty-five quadrillion three hundred twenty-one trillion nine hundred twenty-three billion eight hundred eighty-two million five hundred fifty-seven thousand sixty-six", "~R", 6155321923882557066),
		formatT("one quintillion seven hundred ninety-three quadrillion one hundred eighty-five trillion six hundred fifty-seven billion eight hundred twenty-five million nine hundred eighty-nine thousand eight hundred eighty-five", "~R", 1793185657825989885),
		formatT("eight quintillion seven hundred sixteen quadrillion seven hundred twenty-six trillion one hundred fifty-one billion six hundred sixty-nine million six hundred thirty-two thousand six hundred fifty-one", "~R", 8716726151669632651),
		formatT("eight quintillion three hundred sixty-nine quadrillion four hundred ninety-one trillion nine hundred seventy-five billion three hundred sixteen million six hundred seventy-four thousand six hundred twelve", "~R", 8369491975316674612),
		formatT("five quintillion five hundred eleven quadrillion two hundred ninety trillion six hundred fifty-three billion two hundred fifty-two million eight hundred sixty-six thousand two hundred ninety-five", "~R", 5511290653252866295),
		formatT("eight quintillion one hundred fifty-nine quadrillion four hundred sixty trillion eighteen billion six hundred ninety-three million eight hundred sixty-one thousand six hundred thirty-three", "~R", 8159460018693861633),
		formatT("eight quintillion eight hundred nine quadrillion two hundred eighty-seven trillion one hundred thirty billion ninety-one million two hundred six thousand one hundred ninety-five", "~R", 8809287130091206195),
		formatT("four quintillion four hundred ninety quadrillion twenty-four trillion eight hundred five billion six hundred eighty-four million seven hundred twenty-nine thousand six hundred eighty-nine", "~R", 4490024805684729689),
		formatT("six hundred twenty-six quadrillion nine hundred forty-nine trillion three hundred ninety-five billion seven hundred thirty million four hundred seventy-four thousand two hundred sixty-one", "~R", 626949395730474261),
		formatT("two quintillion four hundred fifteen quadrillion fifty-two trillion four hundred sixty-six billion four hundred eighty-nine million six hundred seven thousand six hundred ninety-three", "~R", 2415052466489607693),
		formatT("five quintillion six hundred seventy-six quadrillion thirty-one trillion seven hundred sixty-four billion three hundred ten million five hundred twenty thousand seven hundred forty-seven", "~R", 5676031764310520747),
		formatT("eight quintillion six hundred forty-five quadrillion seven hundred thirty-six trillion sixty-five billion seven hundred seventy-six million four hundred fourteen thousand seven hundred fifty-two", "~R", 8645736065776414752),
		formatT("one quintillion eight hundred fifty-one quadrillion nine hundred eighty-six trillion eighty-nine billion nine hundred thirty-seven million two hundred eleven thousand two hundred twenty-one", "~R", 1851986089937211221),
		formatT("eight quintillion ninety-seven quadrillion two hundred thirty-three trillion one hundred thirty-nine billion six hundred sixty-two million seven hundred thirty thousand four hundred twenty-three", "~R", 8097233139662730423),
		formatT("five quintillion one hundred eleven quadrillion eight hundred seventy trillion eighty-six billion four hundred ninety-nine million eight hundred nineteen thousand six hundred eighty-five", "~R", 5111870086499819685),
		formatT("one quintillion eight hundred seventy-four quadrillion three hundred forty-one trillion nine hundred fifty-two billion five hundred twenty-two million one hundred fourteen thousand nine hundred twenty-eight", "~R", 1874341952522114928),
		formatT("three quintillion one hundred eighty-seven quadrillion three hundred eighty-six trillion five hundred ninety-five billion nine hundred thirty-nine million two hundred fifty-four thousand three hundred seven", "~R", 3187386595939254307),
		formatT("five quintillion six hundred six quadrillion seventy-nine trillion seven hundred sixty-five billion five hundred ninety-nine million eight hundred twenty-six thousand four hundred fifty-five", "~R", 5606079765599826455),
		formatT("two quintillion four hundred sixty-seven quadrillion six hundred thirteen trillion eighty-three billion two hundred eighty-four million three hundred forty-seven thousand three hundred forty-one", "~R", 2467613083284347341),
		formatT("six quintillion three hundred ninety quadrillion twelve trillion four hundred ninety-seven billion seven hundred thirty-five million five hundred seventy-five thousand five hundred thirty-three", "~R", 6390012497735575533),
		formatT("five hundred sixty-three quadrillion three hundred seventy-one trillion thirty-seven billion seven hundred fifty-seven million two hundred seventy thousand five hundred fifty-three", "~R", 563371037757270553),
		formatT("four quintillion nine hundred twenty quadrillion eight hundred thirty-five trillion nine hundred sixty billion one hundred twenty-six million nine hundred ninety-two thousand one hundred sixty-five", "~R", 4920835960126992165),
		formatT("two quintillion three hundred nineteen quadrillion three hundred twenty-eight trillion thirty billion five hundred thirty-eight million two hundred seventy-seven thousand eighty-four", "~R", 2319328030538277084),
		formatT("five quintillion six hundred forty-nine quadrillion six hundred ninety trillion three hundred eighty-three billion three hundred sixty-five million eight hundred twenty-eight thousand four hundred ninety", "~R", 5649690383365828490),
		formatT("three hundred seventy quadrillion nine hundred forty trillion two hundred fifty-eight billion six hundred thirty-six million four hundred thirty-four thousand eight hundred seventy-three", "~R", 370940258636434873),
		formatT("four quintillion six hundred eighty quadrillion five hundred thirty-four trillion three hundred fifty billion one hundred eighty-three million eight hundred ninety-five thousand seven hundred eighty-three", "~R", 4680534350183895783),
		formatT("eight quintillion six hundred eighty-six quadrillion two hundred forty-seven trillion five hundred twenty-nine billion seventeen million eight hundred forty-two thousand six hundred fifty-one", "~R", 8686247529017842651),
		formatT("two quintillion seven hundred forty quadrillion eight hundred fifty-five trillion eight hundred fifty-nine billion seven hundred seventy-one million six hundred sixty-seven thousand one hundred sixty-six", "~R", 2740855859771667166),
		formatT("five quintillion six hundred ninety-five quadrillion nine hundred fifty-five trillion one hundred ninety-eight billion nine hundred eighty-three million sixty-six thousand seventy-five", "~R", 5695955198983066075),
		formatT("seven quintillion seven hundred seventy-four quadrillion five hundred fifty-six trillion three hundred forty billion two hundred eighteen million one hundred forty-nine thousand seven hundred seventy-six", "~R", 7774556340218149776),
		formatT("seven quintillion six hundred sixty quadrillion two hundred thirty trillion seven hundred fifty-four billion two hundred seventy-five million three hundred seventy-two thousand nine hundred ninety", "~R", 7660230754275372990),
		formatT("six quintillion sixty-one quadrillion nine hundred twenty-eight trillion one hundred sixty-five billion three hundred thirty-two million two hundred thousand three hundred twenty-five", "~R", 6061928165332200325),
		formatT("seven quintillion eight hundred sixty-one quadrillion four hundred four trillion four hundred sixty billion one hundred forty-eight million three hundred fifty-six thousand two hundred fifty-six", "~R", 7861404460148356256),
		formatT("six quintillion eight hundred ten quadrillion sixty-four trillion nine hundred thirty billion nine hundred twenty-two million five hundred forty thousand six hundred eighty-nine", "~R", 6810064930922540689),
		formatT("three quintillion fifty-four quadrillion three hundred seventy-five trillion four hundred fourteen billion nine hundred fourteen million five hundred forty-nine thousand three hundred seventy-four", "~R", 3054375414914549374),
		formatT("six quintillion two hundred forty-seven quadrillion three hundred thirty-eight trillion three hundred seventy-three billion four hundred thirteen million seven hundred ninety-nine thousand five hundred seventy-three", "~R", 6247338373413799573),
		formatT("eight quintillion fifty-four quadrillion four hundred thirty trillion nine hundred forty-seven billion one hundred five million one hundred nineteen thousand two hundred seventy-nine", "~R", 8054430947105119279),
		formatT("four quintillion eighty-two quadrillion ninety-two trillion four hundred twenty billion three hundred twenty-three million five hundred seventy-eight thousand eight hundred eighty-nine", "~R", 4082092420323578889),
		formatT("five quintillion five hundred twenty-seven quadrillion two hundred thirteen trillion nine hundred eighty-five billion three hundred thirty-nine million five hundred ninety-three thousand six hundred seventy-nine", "~R", 5527213985339593679),
		formatT("three quintillion six hundred sixty-three quadrillion eighty-five trillion seventy-six billion sixty-three million five hundred thirty-four thousand seven hundred sixty-one", "~R", 3663085076063534761),
		formatT("two quintillion six hundred sixty-six quadrillion six hundred twenty-eight trillion four hundred ninety-six billion four hundred eighteen million eight hundred eighty-three thousand two hundred fifty-five", "~R", 2666628496418883255),
		formatT("three quintillion four hundred fifty-five quadrillion six hundred forty-one trillion eight hundred billion seven hundred twenty-five million five hundred seventy-eight thousand four hundred sixty-six", "~R", 3455641800725578466),
		formatT("one quintillion eight hundred twenty-seven quadrillion eight hundred four trillion four hundred twenty-five billion three hundred seventeen million four hundred fifty thousand seven hundred eighty-four", "~R", 1827804425317450784),
		formatT("nine quintillion ninety-two quadrillion four hundred thirty-four trillion thirty billion five hundred forty-one million six hundred fifty-nine thousand two hundred thirty-six", "~R", 9092434030541659236),
		formatT("five quintillion fourteen quadrillion nineteen trillion one hundred twenty-six billion two hundred twenty-six million two hundred eighty-seven thousand six hundred fifty-eight", "~R", 5014019126226287658),
		formatT("eight quintillion five hundred seventy-five quadrillion three hundred fifty-five trillion three hundred thirty-seven billion one hundred eighty-seven million six hundred four thousand five hundred sixty-two", "~R", 8575355337187604562),
		formatT("nine quintillion sixty-five quadrillion five hundred sixty-five trillion two hundred sixty-five billion seven hundred fifty-five million four hundred thirty-seven thousand eight hundred seventy-two", "~R", 9065565265755437872),
		formatT("eight quintillion two hundred ninety-four quadrillion nine hundred two trillion three hundred fourteen billion two hundred twenty-five million two hundred ten thousand eight hundred twenty-three", "~R", 8294902314225210823),
		formatT("two quintillion seven hundred eighty-five quadrillion six hundred trillion eight hundred twenty-six billion two hundred sixty-six million sixty-three thousand forty-nine", "~R", 2785600826266063049),
		formatT("six quintillion one hundred fifty-five quadrillion sixty-eight trillion five hundred sixty-nine billion seven hundred million twenty-eight thousand seven hundred six", "~R", 6155068569700028706),
		formatT("seven quintillion thirty-four quadrillion four hundred sixty-six trillion four hundred five billion five million one hundred fifty-nine thousand two hundred thirty-two", "~R", 7034466405005159232),
		formatT("eight quintillion seven hundred eighty-five quadrillion eight hundred eighty trillion one hundred seventy-one billion five hundred ninety-nine million six hundred forty-eight thousand seven hundred twenty", "~R", 8785880171599648720),
		formatT("one hundred eight trillion eight hundred four billion two hundred seventy-four million four hundred six thousand three hundred fifty-one", "~R", 108804274406351),
		formatT("three quintillion seven hundred ninety-five quadrillion five hundred forty-seven trillion seven hundred sixty-six billion four hundred eighty-two million ninety-nine thousand eight hundred fifty-six", "~R", 3795547766482099856),
		formatT("four quintillion nine hundred eighty-two quadrillion nine hundred thirty-one trillion three hundred thirty-eight billion five hundred twenty-four million one hundred forty-six thousand six hundred fifty-eight", "~R", 4982931338524146658),
		formatT("five quintillion seven hundred fifty-five quadrillion six hundred seventy-six trillion five hundred ninety-four billion nine hundred fifty-one million one hundred eleven thousand nine hundred twenty-six", "~R", 5755676594951111926),
		formatT("two quintillion nine hundred seventy quadrillion three hundred ninety-five trillion forty-five billion eight hundred fifty-six million thirty-three thousand four hundred twenty-five", "~R", 2970395045856033425),
		formatT("nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven", "~R", 9223372036854775807),
	}
	runTests(t, tcs)

}

func Test_A(t *testing.T) {
	tcs := []formatTest{
		formatT("Hello", "~a", "Hello"),
	}
	runTests(t, tcs)
}

//helper

func runTests(t *testing.T, tests []formatTest) {
	for _, tc := range tests {
		t.Run(tc.format, func(t *testing.T) {
			result := Sformat(tc.format, tc.args...)
			if result != tc.expected {
				t.Errorf("expected `%s`, got `%s`", tc.expected, result)
			}
		})
	}
}

type repeatDef struct {
	index   int
	linksTo int
}

type formatTest struct {
	format   string
	args     []interface{}
	expected string
}

func formatT(expected string, format string, args ...interface{}) formatTest {
	return formatTest{format, args, expected}
}
