# generateGoFromExcel
generateGoFromExcel
Parse excel file & generate golang code

excel or wps file (ItemModel.xlsm with sheel named ItemBase below)
-----------------------------------------------
SERVER	YES	YES	YES	YES	YES<br />
CLIENT	YES	YES	YES	YES	YES<br />
NAME	奖励ID	奖励类型	职业类型	概率	次数<br />
KEY	AwardID	AwardType	JobType	Ratio	Times<br />
TYPE	Uint32	Uint8	Uint8	Uint32	Other_;<br />
VALUE	100	0	0	0	0;0<br />
VALUE	101	0	0	0	0;0<br />
VALUE	102	0	0	0	0;0<br />
VALUE	111	0	0	0	0;0<br />
VALUE	112	0	0	0	0;0<br />
VALUE	113	0	0	0	0;0<br />

generate file(ItemBase.go)
----------------------------------------------
package cfg<br /><br />

var Sheet1Base_cfgMap map[int]*map[string]string = map[int]*map[string]string{ 
100:&map[string]string {
"AwardID":"100",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},<br />
101:&map[string]string {
"AwardID":"101",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},<br />
102:&map[string]string {
"AwardID":"102",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},<br />
111:&map[string]string {
"AwardID":"111",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},<br />
112:&map[string]string {
"AwardID":"112",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},<br />
113:&map[string]string {
"AwardID":"113",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},<br />
}
