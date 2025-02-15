package hole

import (
	"math/rand"
	"strings"
)

func proximityGrid() []Scorecard {
	args := []string{
		"---------\n------0--\n---------\n---------\n---------\n---------\n---------\n---------\n---------",
		"----####-\n----#-0#-\n----####-\n---------\n---------\n---------\n---------\n-----0---\n---------",
		"-#-------\n-#-------\n-#-----0-\n-#-------\n-#-------\n-#-------\n-#-------\n-#-------\n0#-------",
		"-#-------\n-#-------\n-#-----0-\n-#-------\n-#-------\n-#-------\n-########\n-#-------\n-#-------",
		"---------\n---------\n---------\n---------\n----0----\n---------\n---------\n---------\n---------",
		"-0-0-0-0-\n---------\n-0-----0-\n---------\n-0-----0-\n---------\n-0-----0-\n---------\n-0-0-0-0-",
		"---------\n---------\n---------\n---###---\n---#-#---\n---###---\n---------\n---------\n--------0",
		"--------0\n-########\n---------\n########-\n---------\n-########\n---------\n########-\n---------",
		"---------\n-#######-\n-#-----#-\n-#-###-#-\n-#---#-#-\n-#####-#-\n-------#-\n########-\n0--------",
		"---------\n-#######-\n---------\n-#######-\n-------0-\n--#------\n---#-----\n----#----\n-----#---",
		"---------\n--------#\n--------0\n---------\n---------\n---------\n---------\n-----#-##\n---------",
		"---------\n---------\n---------\n---------\n---------\n---------\n---------\n---------\n---------",
		"#--------\n-#-------\n--#------\n---#-----\n----#----\n-----#---\n------#--\n-0-----#-\n--------#",
		"----#--#-\n---#--#-0\n--#--#---\n-#--#----\n#--#-----\n--#------\n-#-------\n---------\n---------",
		"---------\n---------\n---------\n---------\n---------\n---------\n---------\n-------#-\n--------0",
		"---------\n----#----\n---------\n---------\n---------\n---------\n---------\n---------\n000000000",
		"--------0\n---------\n-0#------\n---------\n---------\n---------\n---------\n-----0---\n---------",
		"-------#0\n##-##--#-\n0#-##--#-\n-##-#-#--\n----#--#-\n----##---\n-----####\n---------\n-0-------",
		"#----#---\n--##-#-#-\n-#---#-#-\n#--##--#-\n--#---#--\n-#--##--#\n-#-#---#0\n-#-#-##--\n---#----#",
	}

	outs := []string{
		"765432123\n654321012\n765432123\n876543234\n987654345\nA98765456\nBA9876567\nCBA987678\nDCBA98789",
		"CBA9####A\nBA98#10#9\nA987####8\n987654567\n876543456\n765432345\n654321234\n543210123\n654321234",
		"8#7654323\n7#6543212\n6#5432101\n5#6543212\n4#7654323\n3#8765434\n2#9876545\n1#A987656\n0#BA98767",
		"-#7654323\n-#6543212\n-#5432101\n-#6543212\n-#7654323\n-#8765434\n-########\n-#-------\n-#-------",
		"876545678\n765434567\n654323456\n543212345\n432101234\n543212345\n654323456\n765434567\n876545678",
		"101010101\n212121212\n101232101\n212343212\n101232101\n212343212\n101232101\n212121212\n101010101",
		"GFEDCBA98\nFEDCBA987\nEDCBA9876\nDCB###765\nCBA#-#654\nBA9###543\nA98765432\n987654321\n876543210",
		"876543210\n9########\nABCDEFGHI\n########J\nSRQPONMLK\nT########\nUVWXYZabc\n########d\nmlkjihgfe",
		"ONMLKJIHG\nP#######F\nQ#ihgfe#E\nR#j###d#D\nS#klm#c#C\nT#####b#B\nUVWXYZa#A\n########9\n012345678",
		"BCBA98765\nA#######4\n9A9876543\n8#######2\n765432101\n87#543212\n989#54323\nA9AB#5434\nBABCD#545",
		"A98765434\n98765432#\n876543210\n987654321\nA98765432\nBA9876543\nCBA987654\nDCBA9#7##\nEDCBA989A",
		"---------\n---------\n---------\n---------\n---------\n---------\n---------\n---------\n---------",
		"#--------\n7#-------\n65#------\n545#-----\n4345#----\n32345#---\n212345#--\n1012345#-\n21234567#",
		"----#QR#1\n---#OP#10\n--#MN#321\n-#KL#5432\n#IJ#76543\nGH#987654\nF#BA98765\nEDCBA9876\nFEDCBA987",
		"GFEDCBA98\nFEDCBA987\nEDCBA9876\nDCBA98765\nCBA987654\nBA9876543\nA98765432\n9876543#1\n876543210",
		"888898888\n7777#7777\n666666666\n555555555\n444444444\n333333333\n222222222\n111111111\n000000000",
		"323443210\n212344321\n10#455432\n212344543\n323443454\n434432345\n544321234\n543210123\n654321234",
		"IHGFEDE#0\n##H##CD#1\n0#I##BC#2\n1##6#A#43\n2345#98#4\n3345##765\n32345####\n212345678\n101234567",
		"#nmlk#KJI\npo##j#L#H\nq#ghi#M#G\n#ef##ON#F\ncd#RQP#DE\nb#TS##BC#\na#U#89A#0\nZ#V#7##21\nYXW#6543#",
	}

	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
		outs[i], outs[j] = outs[j], outs[i]
	})

	return []Scorecard{{Args: args, Answer: strings.Join(outs, "\n\n")}}
}
