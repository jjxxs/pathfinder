package problem

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	TestProblemDir           = "../samples/"
	TestProblemFileGermany   = "../samples/germany10.json"
	TestProblemFileWorkpiece = "../samples/workpiece.json"
)

func TestProblemFileLoadGermany(t *testing.T) {
	// try to load problem
	problem, err := FromFile(TestProblemFileGermany)
	if err != nil {
		t.Fatalf("failed to load problem from file=%s, err=%s", TestProblemFileGermany, err)
	}

	// test info
	if problem.Info.Name != "Germany" ||
		problem.Info.Description != "Ten biggest german cities" ||
		problem.Info.Type != "geographic" {
		t.Fatalf("failed to load problem info")
	}

	// test points length
	if len(problem.Points) != 10 {
		t.Fatalf("failed to load problem-points, invalid length")
	}

	// verify points
	expectedPoints := []point{
		{X: 13.23, Y: 52.31, Name: "Berlin"},
		{X: 10.0, Y: 53.33, Name: "Hamburg"},
		{X: 11.34, Y: 48.8, Name: "Munich"},
		{X: 6.57, Y: 50.56, Name: "Cologne"},
		{X: 8.41, Y: 50.7, Name: "Frankfurt"},
		{X: 9.11, Y: 48.47, Name: "Stuttgart"},
		{X: 6.47, Y: 51.14, Name: "Düsseldorf"},
		{X: 7.28, Y: 51.31, Name: "Dortmund"},
		{X: 7.1, Y: 51.27, Name: "Essen"},
		{X: 12.23, Y: 51.2, Name: "Leipzig"},
	}

	for i, actualPoint := range problem.Points {
		expectedPoint := expectedPoints[i]
		if expectedPoint.X != actualPoint.X ||
			expectedPoint.Y != actualPoint.Y ||
			expectedPoint.Name != actualPoint.Name {
			t.Fatalf("failed to load problem-points, invalid point-details")
		}
	}

	// verify adjacency
	expectedAdjacency := [][]float64{
		{0, 375.94456892271836, 435.41080570770396, 764.9363495938279, 564.0500170213354, 620.698804527503, 762.5147542056219, 670.5879123295689, 691.0540137467287, 163.88474702971845},
		{375.94456892271836, 0, 516.9268227447614, 488.19591094433196, 338.5059712929509, 542.0111370296753, 460.5807783282825, 375.20510454953563, 394.05771472432156, 339.83756567837725},
		{435.41080570770396, 516.9268227447614, 0, 564.5107479834015, 386.59839473294056, 250.58002583843984, 599.3996606993541, 528.8019295041576, 543.8207798329414, 279.3554029050923},
		{764.9363495938279, 488.19591094433196, 564.5107479834015, 0, 205.17995727159678, 364.36784438241585, 65.03358325815319, 114.39657881554925, 98.06966348868784, 633.2638711583669},
		{564.0500170213354, 338.5059712929509, 386.59839473294056, 205.17995727159678, 0, 257.13401653023567, 221.10558969129497, 142.4881748560758, 158.62608010738126, 428.27052016343873},
		{620.698804527503, 542.0111370296753, 250.58002583843984, 364.36784438241585, 257.13401653023567, 0, 415.55075287357187, 372.9570270509856, 380.72522339447977, 457.52071180514633},
		{762.5147542056219, 460.5807783282825, 599.3996606993541, 65.03358325815319, 221.10558969129497, 415.55075287357187, 0, 92.00232196874512, 71.50827504313482, 640.5165769987474},
		{670.5879123295689, 375.20510454953563, 528.8019295041576, 114.39657881554925, 142.4881748560758, 372.9570270509856, 92.00232196874512, 0, 20.495772281405383, 550.5467849848291},
		{691.0540137467287, 394.05771472432156, 543.8207798329414, 98.06966348868784, 158.62608010738126, 380.72522339447977, 71.50827504313482, 20.495772281405383, 0, 570.4815418440626},
		{163.88474702971845, 339.83756567837725, 279.3554029050923, 633.2638711583669, 428.27052016343873, 457.52071180514633, 640.5165769987474, 550.5467849848291, 570.4815418440626, 0},
	}

	if !reflect.DeepEqual(expectedAdjacency, problem.Adjacency) {
		t.Fatalf("failed to load/calculate adjacency")
	}
}

func TestProblemFileLoadWerkstueck(t *testing.T) {
	// try to load problem
	problem, err := FromFile(TestProblemFileWorkpiece)
	if err != nil {
		t.Fatalf("failed to load problem from file=%s, err=%s", TestProblemFileWorkpiece, err)
	}

	// test info
	if problem.Info.Name != "Workpiece" ||
		problem.Info.Description != "Technical drawing of a flat workpiece with prisms" ||
		problem.Info.Type != "cartesian" {
		t.Fatalf("failed to load problem info")
	}

	// test points length
	if len(problem.Points) != 30 {
		t.Fatalf("failed to load problem-points, invalid length")
	}

	// verify points
	expectedPoints := []point{
		{X: 230, Y: 138},
		{X: 195, Y: 197},
		{X: 157, Y: 198},
		{X: 157, Y: 298},
		{X: 187, Y: 328},
		{X: 157, Y: 357},
		{X: 157, Y: 550},
		{X: 218, Y: 611},
		{X: 309, Y: 611},
		{X: 357, Y: 611},
		{X: 514, Y: 611},
		{X: 278, Y: 555},
		{X: 389, Y: 555},
		{X: 513, Y: 537},
		{X: 559, Y: 537},
		{X: 559, Y: 138},
		{X: 309, Y: 207},
		{X: 350, Y: 274},
		{X: 270, Y: 274},
		{X: 328, Y: 432},
		{X: 328, Y: 450},
		{X: 308, Y: 450},
		{X: 275, Y: 475},
		{X: 239, Y: 441},
		{X: 276, Y: 406},
		{X: 309, Y: 433},
		{X: 456, Y: 312},
		{X: 456, Y: 235},
		{X: 417, Y: 273},
		{X: 494, Y: 273},
	}

	for i, actualPoint := range problem.Points {
		expectedPoint := expectedPoints[i]
		if expectedPoint.X != actualPoint.X ||
			expectedPoint.Y != actualPoint.Y ||
			expectedPoint.Name != actualPoint.Name {
			t.Fatalf("failed to load problem-points, invalid point-details")
		}
	}

	// verify adjacency
	expectedAdjacency := [][]float64{
		{0.00000000000000, 68.60029154456998, 94.49338601193207, 175.86642658563346, 194.80503073586164, 230.84626919229169, 418.41725585831188, 473.15219538748840, 479.55187414918942, 489.75299896988889, 551.71097505849923, 419.75349909202663, 446.28466251933867, 489.17277111466456, 517.14794788338861, 329.00000000000000, 104.89041900955492, 181.37254478007415, 141.76036117335480, 309.90321069650116, 327.02905069733481, 321.60223879817750, 339.99117635609309, 303.13363389765908, 271.91910561782890, 305.39482641328419, 285.22271999264012, 245.93698379869588, 230.63824487712353, 296.51475511346820},
		{68.60029154456998, 0.00000000000000, 38.01315561749642, 107.91200118615167, 131.24404748406687, 164.45060048537371, 355.03943442947292, 414.63839667835879, 429.40889604198935, 444.56720526822488, 522.64423846436875, 367.49557820469079, 407.18546143004664, 465.53624993119496, 498.09236091311419, 368.75059321986180, 114.43775600735974, 173.07223925286226, 107.48953437428223, 270.02592468131650, 285.82861998057507, 277.08843353702082, 289.28186946298587, 247.93547547698776, 224.14727301486406, 262.09158704544484, 285.21220170252184, 263.75177724519699, 234.64867355261143, 308.50769844527383},
		{94.49338601193207, 38.01315561749642, 0.00000000000000, 100.00000000000000, 133.41664064126334, 159.00000000000000, 352.00000000000000, 417.48053846856141, 440.08294672709144, 458.87797942372436, 545.91024903366667, 376.94827231332420, 425.76167042137553, 491.58620810596386, 525.85644428874309, 406.45294930655871, 152.26621424334422, 207.42468512691539, 136.18002790424151, 289.82235938588315, 304.54063768239536, 293.77712640707750, 301.08636634693374, 256.46247288833507, 239.63513932643517, 279.87318556803541, 319.99531246566721, 301.28060010561580, 270.60118255469615, 345.24484065659840},
		{175.86642658563346, 107.91200118615167, 100.00000000000000, 0.00000000000000, 42.42640687119285, 59.00000000000000, 252.00000000000000, 318.88869531546584, 347.95545691941663, 371.44178547923224, 474.78205526325445, 284.05985284795173, 346.22680427719632, 428.78549415762654, 467.68044645890427, 432.67077553262135, 177.15812146215595, 194.48650338776724, 115.52056094046635, 217.24870540465827, 228.79029699705362, 214.25452153921981, 212.72752525237536, 164.84234892769516, 160.70158679988197, 203.29535164385831, 299.32757975168278, 305.56505035753025, 261.19915773217951, 337.92602740836639},
		{194.80503073586164, 131.24404748406687, 133.41664064126334, 42.42640687119285, 0.00000000000000, 41.72529209005013, 224.01785643113362, 284.69281690973520, 308.17689725221129, 330.13482094441355, 432.45577808603736, 244.56083087853622, 303.86345617727710, 387.24281788046113, 426.69075452838206, 417.71282000915414, 171.82840277439584, 171.71196813268432, 99.02019995940222, 175.20559351801529, 186.45374761586316, 171.82840277439584, 171.32717239247253, 124.39051410778879, 118.34272263219231, 160.96272860510285, 269.47541631844638, 284.62255708218208, 236.48467180770933, 311.88780033851918},
		{230.84626919229169, 164.45060048537371, 159.00000000000000, 59.00000000000000, 41.72529209005013, 0.00000000000000, 193.00000000000000, 261.22212769977966, 296.00675667964066, 323.28934408668653, 438.13810608071969, 232.04525420701884, 305.00491799313664, 398.91853805006355, 440.45885165359090, 457.78269954204256, 213.55093069335942, 210.09045670853305, 140.20698984002189, 186.72439583514523, 194.65353837010002, 177.34147850968199, 166.87720036002523, 117.38824472663352, 128.69343417595164, 169.94116628998401, 302.36732627716242, 322.93188136199871, 273.23250172700904, 347.31109973624513},
		{418.41725585831188, 355.03943442947292, 352.00000000000000, 252.00000000000000, 224.01785643113362, 193.00000000000000, 0.00000000000000, 86.26702730475880, 163.78339354159201, 209.09567188251410, 362.17399133565624, 121.10326172320876, 232.05387305537479, 356.23728047468586, 402.21014407893790, 575.62835232465750, 375.17062784818324, 336.78628238097821, 298.23648334836565, 207.76188293332345, 198.09341230843594, 181.11046352985792, 139.81773850266640, 136.40014662748717, 186.80738743422327, 191.81501505356664, 382.15834414546021, 434.31094851500120, 379.90656746100086, 436.23158986941786},
		{473.15219538748840, 414.63839667835879, 417.48053846856141, 318.88869531546584, 284.69281690973520, 261.22212769977966, 86.26702730475880, 0.00000000000000, 91.00000000000000, 139.00000000000000, 296.00000000000000, 82.07313811473277, 179.93609976877903, 304.13977050034083, 348.93695705671536, 583.10376435073715, 414.12196271146985, 361.92955115602263, 340.98826959295826, 210.09759636892565, 194.98974332000131, 184.44782460088814, 147.46185947559457, 171.29214809792072, 213.04694318389082, 199.91248085099639, 382.15834414546021, 444.99438198700892, 392.23079940259663, 436.37140144606178},
		{479.55187414918942, 429.40889604198935, 440.08294672709144, 347.95545691941663, 308.17689725221129, 296.00675667964066, 163.78339354159201, 91.00000000000000, 0.00000000000000, 48.00000000000000, 205.00000000000000, 64.00781202322104, 97.65244492586962, 217.00691233230336, 260.72207424765554, 535.00373830469630, 404.00000000000000, 339.48490393535911, 339.24917096435183, 180.00555546982432, 162.11724152600181, 161.00310556011024, 140.18559127100045, 183.84776310850236, 207.63910999616618, 178.00000000000000, 333.18163214679169, 403.71400768365714, 354.83517300290288, 385.31675281513520},
		{489.75299896988889, 444.56720526822488, 458.87797942372436, 371.44178547923224, 330.13482094441355, 323.28934408668653, 209.09567188251410, 139.00000000000000, 48.00000000000000, 0.00000000000000, 157.00000000000000, 96.83491106000976, 64.49806198638839, 172.66151858477326, 215.12786895239771, 514.32771653878422, 406.84149247587811, 337.07269245668658, 348.04884714648892, 181.33394607739612, 163.59095329510126, 168.29141392239831, 158.80806024884254, 206.93960471596537, 220.42232191862965, 184.35834670553976, 314.96348994764458, 388.81486597093993, 343.28413887041154, 364.70947341685547},
		{551.71097505849923, 522.64423846436875, 545.91024903366667, 474.78205526325445, 432.45577808603736, 438.13810608071969, 362.17399133565624, 296.00000000000000, 205.00000000000000, 157.00000000000000, 0.00000000000000, 242.55308697272852, 136.97079980784227, 74.00675644831355, 86.60831368869850, 475.13577006998747, 453.03531871146646, 374.78660595063957, 416.05889006245258, 258.14143410154054, 246.00203251192866, 261.45171638373307, 274.98545416076104, 323.30326320654422, 314.11622052991788, 271.49401466699038, 304.57347225259127, 380.44710539048657, 351.64328516267733, 338.59119894055129},
		{419.75349909202663, 367.49557820469079, 376.94827231332420, 284.05985284795173, 244.56083087853622, 232.04525420701884, 121.10326172320876, 82.07313811473277, 64.00781202322104, 96.83491106000976, 242.55308697272852, 0.00000000000000, 111.00000000000000, 235.68835355188853, 281.57592226609148, 502.84192347098508, 349.37801877050020, 290.07757583101801, 281.11385593741193, 132.77424449041311, 116.29703349613007, 109.20164833920778, 80.05623023850174, 120.48651376813922, 149.01342221424215, 125.87692401707312, 301.21918929576844, 366.17482163578643, 314.39624679693617, 355.21824277477640},
		{446.28466251933867, 407.18546143004664, 425.76167042137553, 346.22680427719632, 303.86345617727710, 305.00491799313664, 232.05387305537479, 179.93609976877903, 97.65244492586962, 64.49806198638839, 136.97079980784227, 111.00000000000000, 0.00000000000000, 125.29964086141668, 170.95028517086482, 450.32099662351965, 357.07702250354896, 283.69349657685143, 305.15897496223175, 137.29530217745980, 121.43310915891102, 132.61221663180206, 139.26952286842948, 188.40382161729099, 187.00267377767625, 145.89036979869508, 252.06745129032427, 326.93883219954159, 283.38666164800350, 300.91360886473711},
		{489.17277111466456, 465.53624993119496, 491.58620810596386, 428.78549415762654, 387.24281788046113, 398.91853805006355, 356.23728047468586, 304.13977050034083, 217.00691233230336, 172.66151858477326, 74.00675644831355, 235.68835355188853, 125.29964086141668, 0.00000000000000, 46.00000000000000, 401.64287619725064, 387.96391584785306, 309.41557814693169, 358.07541105191797, 212.72047386182646, 204.43580899636933, 222.69710370815332, 245.94308284641795, 290.33084576048753, 270.79512551004314, 228.98034850178738, 232.10773360661639, 307.33206796558017, 280.91279785727102, 264.68282906150148},
		{517.14794788338861, 498.09236091311419, 525.85644428874309, 467.68044645890427, 426.69075452838206, 440.45885165359090, 402.21014407893790, 348.93695705671536, 260.72207424765554, 215.12786895239771, 86.60831368869850, 281.57592226609148, 170.95028517086482, 46.00000000000000, 0.00000000000000, 399.00000000000000, 414.00483088968900, 335.93154064481649, 390.75567814172581, 253.74396544548603, 246.84002916869053, 265.65014586858410, 290.68883707497264, 334.08980828513762, 311.84932259025351, 270.76927447552094, 247.45504642257754, 319.08149429260231, 299.76657585528108, 271.88416651213805},
		{329.00000000000000, 368.75059321986180, 406.45294930655871, 432.67077553262135, 417.71282000915414, 457.78269954204256, 575.62835232465750, 583.10376435073715, 535.00373830469630, 514.32771653878422, 475.13577006998747, 502.84192347098508, 450.32099662351965, 401.64287619725064, 399.00000000000000, 0.00000000000000, 259.34725755249468, 249.35316320431951, 319.40100187695089, 373.89437011006197, 388.20741878537046, 400.43101777959208, 440.70965498840616, 440.69150207372957, 389.76018267647606, 386.68462601970612, 202.20039564748632, 141.48498153514387, 195.93111034238538, 149.83324063771698},
		{104.89041900955492, 114.43775600735974, 152.26621424334422, 177.15812146215595, 171.82840277439584, 213.55093069335942, 375.17062784818324, 414.12196271146985, 404.00000000000000, 406.84149247587811, 453.03531871146646, 349.37801877050020, 357.07702250354896, 387.96391584785306, 414.00483088968900, 259.34725755249468, 0.00000000000000, 78.54934754662193, 77.52418977325722, 225.80079716422614, 243.74166652421167, 243.00205760445732, 270.14810752622350, 244.24577785501228, 201.71762441591463, 226.00000000000000, 180.64883060789515, 149.64290828502365, 126.57013865837392, 196.42046736529267},
		{181.37254478007415, 173.07223925286226, 207.42468512691539, 194.48650338776724, 171.71196813268432, 210.09045670853305, 336.78628238097821, 361.92955115602263, 339.48490393535911, 337.07269245668658, 374.78660595063957, 290.07757583101801, 283.69349657685143, 309.41557814693169, 335.93154064481649, 249.35316320431951, 78.54934754662193, 0.00000000000000, 80.00000000000000, 159.52429282087414, 177.36967046256808, 180.94197965093673, 214.53671014537349, 200.52431274037571, 151.32745950421557, 164.20109622045769, 112.60550608207397, 112.94689017409908, 67.00746227100382, 144.00347218036097},
		{141.76036117335480, 107.48953437428223, 136.18002790424151, 115.52056094046635, 99.02019995940222, 140.20698984002189, 298.23648334836565, 340.98826959295826, 339.24917096435183, 348.04884714648892, 416.05889006245258, 281.11385593741193, 305.15897496223175, 358.07541105191797, 390.75567814172581, 319.40100187695089, 77.52418977325722, 80.00000000000000, 0.00000000000000, 168.30923919975396, 185.31055015837603, 180.05554698481245, 201.06217943710845, 169.85287751463028, 132.13629327327143, 163.71316379570703, 189.84203960134857, 190.04473157654226, 147.00340132119393, 224.00223213173570},
		{309.90321069650116, 270.02592468131650, 289.82235938588315, 217.24870540465827, 175.20559351801529, 186.72439583514523, 207.76188293332345, 210.09759636892565, 180.00555546982432, 181.33394607739612, 258.14143410154054, 132.77424449041311, 137.29530217745980, 212.72047386182646, 253.74396544548603, 373.89437011006197, 225.80079716422614, 159.52429282087414, 168.30923919975396, 0.00000000000000, 18.00000000000000, 26.90724809414742, 68.24954212300622, 89.45389874119518, 58.13776741499453, 19.02629759044045, 175.45369759569047, 234.93190502781866, 182.21415971323415, 229.86300267768192},
		{327.02905069733481, 285.82861998057507, 304.54063768239536, 228.79029699705362, 186.45374761586316, 194.65353837010002, 198.09341230843594, 194.98974332000131, 162.11724152600181, 163.59095329510126, 246.00203251192866, 116.29703349613007, 121.43310915891102, 204.43580899636933, 246.84002916869053, 388.20741878537046, 243.74166652421167, 177.36967046256808, 185.31055015837603, 18.00000000000000, 0.00000000000000, 20.00000000000000, 58.60034129593445, 89.45389874119518, 68.11754546370561, 25.49509756796392, 188.22327167489146, 250.21790503479161, 198.11612756158948, 242.66231681083076},
		{321.60223879817750, 277.08843353702082, 293.77712640707750, 214.25452153921981, 171.82840277439584, 177.34147850968199, 181.11046352985792, 184.44782460088814, 161.00310556011024, 168.29141392239831, 261.45171638373307, 109.20164833920778, 132.61221663180206, 222.69710370815332, 265.65014586858410, 400.43101777959208, 243.00205760445732, 180.94197965093673, 180.05554698481245, 26.90724809414742, 20.00000000000000, 0.00000000000000, 41.40048308896890, 69.58448102845921, 54.40588203494178, 17.02938636592640, 202.35612172603032, 261.01532522057016, 207.87015177749788, 256.75864152935537},
		{339.99117635609309, 289.28186946298587, 301.08636634693374, 212.72752525237536, 171.32717239247253, 166.87720036002523, 139.81773850266640, 147.46185947559457, 140.18559127100045, 158.80806024884254, 274.98545416076104, 80.05623023850174, 139.26952286842948, 245.94308284641795, 290.68883707497264, 440.70965498840616, 270.14810752622350, 214.53671014537349, 201.06217943710845, 68.24954212300622, 58.60034129593445, 41.40048308896890, 0.00000000000000, 49.51767361255979, 69.00724599634447, 54.03702434442518, 243.57750306627253, 300.60106453570654, 246.91699009991191, 297.93455657241242},
		{303.13363389765908, 247.93547547698776, 256.46247288833507, 164.84234892769516, 124.39051410778879, 117.38824472663352, 136.40014662748717, 171.29214809792072, 183.84776310850236, 206.93960471596537, 323.30326320654422, 120.48651376813922, 188.40382161729099, 290.33084576048753, 334.08980828513762, 440.69150207372957, 244.24577785501228, 200.52431274037571, 169.85287751463028, 89.45389874119518, 89.45389874119518, 69.58448102845921, 49.51767361255979, 0.00000000000000, 50.93132631298737, 70.45565981523414, 252.44801445050030, 299.20728600754359, 244.76110802167898, 305.36699232235298},
		{271.91910561782890, 224.14727301486406, 239.63513932643517, 160.70158679988197, 118.34272263219231, 128.69343417595164, 186.80738743422327, 213.04694318389082, 207.63910999616618, 220.42232191862965, 314.11622052991788, 149.01342221424215, 187.00267377767625, 270.79512551004314, 311.84932259025351, 389.76018267647606, 201.71762441591463, 151.32745950421557, 132.13629327327143, 58.13776741499453, 68.11754546370561, 54.40588203494178, 69.00724599634447, 50.93132631298737, 0.00000000000000, 42.63801121065568, 203.06649157357302, 248.27605603440699, 193.82982226685345, 255.36836139193124},
		{305.39482641328419, 262.09158704544484, 279.87318556803541, 203.29535164385831, 160.96272860510285, 169.94116628998401, 191.81501505356664, 199.91248085099639, 178.00000000000000, 184.35834670553976, 271.49401466699038, 125.87692401707312, 145.89036979869508, 228.98034850178738, 270.76927447552094, 386.68462601970612, 226.00000000000000, 164.20109622045769, 163.71316379570703, 19.02629759044045, 25.49509756796392, 17.02938636592640, 54.03702434442518, 70.45565981523414, 42.63801121065568, 0.00000000000000, 190.39432764659770, 246.60291969074495, 193.03885619221847, 244.59149617270018},
		{285.22271999264012, 285.21220170252184, 319.99531246566721, 299.32757975168278, 269.47541631844638, 302.36732627716242, 382.15834414546021, 382.15834414546021, 333.18163214679169, 314.96348994764458, 304.57347225259127, 301.21918929576844, 252.06745129032427, 232.10773360661639, 247.45504642257754, 202.20039564748632, 180.64883060789515, 112.60550608207397, 189.84203960134857, 175.45369759569047, 188.22327167489146, 202.35612172603032, 243.57750306627253, 252.44801445050030, 203.06649157357302, 190.39432764659770, 0.00000000000000, 77.00000000000000, 55.15432893255070, 54.45181356024793},
		{245.93698379869588, 263.75177724519699, 301.28060010561580, 305.56505035753025, 284.62255708218208, 322.93188136199871, 434.31094851500120, 444.99438198700892, 403.71400768365714, 388.81486597093993, 380.44710539048657, 366.17482163578643, 326.93883219954159, 307.33206796558017, 319.08149429260231, 141.48498153514387, 149.64290828502365, 112.94689017409908, 190.04473157654226, 234.93190502781866, 250.21790503479161, 261.01532522057016, 300.60106453570654, 299.20728600754359, 248.27605603440699, 246.60291969074495, 77.00000000000000, 0.00000000000000, 54.45181356024793, 53.74011537017761},
		{230.63824487712353, 234.64867355261143, 270.60118255469615, 261.19915773217951, 236.48467180770933, 273.23250172700904, 379.90656746100086, 392.23079940259663, 354.83517300290288, 343.28413887041154, 351.64328516267733, 314.39624679693617, 283.38666164800350, 280.91279785727102, 299.76657585528108, 195.93111034238538, 126.57013865837392, 67.00746227100382, 147.00340132119393, 182.21415971323415, 198.11612756158948, 207.87015177749788, 246.91699009991191, 244.76110802167898, 193.82982226685345, 193.03885619221847, 55.15432893255070, 54.45181356024793, 0.00000000000000, 77.00000000000000},
		{296.51475511346820, 308.50769844527383, 345.24484065659840, 337.92602740836639, 311.88780033851918, 347.31109973624513, 436.23158986941786, 436.37140144606178, 385.31675281513520, 364.70947341685547, 338.59119894055129, 355.21824277477640, 300.91360886473711, 264.68282906150148, 271.88416651213805, 149.83324063771698, 196.42046736529267, 144.00347218036097, 224.00223213173570, 229.86300267768192, 242.66231681083076, 256.75864152935537, 297.93455657241242, 305.36699232235298, 255.36836139193124, 244.59149617270018, 54.45181356024793, 53.74011537017761, 77.00000000000000, 0.00000000000000},
	}

	for _, adjRow := range problem.Adjacency {
		str := "{"
		for i, ele := range adjRow {
			if i == len(adjRow)-1 {
				str += fmt.Sprintf("%.14f},", ele)
			} else {
				str += fmt.Sprintf("%.14f, ", ele)
			}
		}
		fmt.Println(str)
	}

	if !reflect.DeepEqual(expectedAdjacency, problem.Adjacency) {
		t.Fatalf("failed to load/calculate adjacency")
	}
}
