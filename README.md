# generateGoFromExcel
generateGoFromExcel
Parse excel file & generate golang code

excel or wps file (ItemModel.xlsm with sheel named ItemBase below)
-----------------------------------------------
SERVER	YES	YES	YES	YES	YES
CLIENT	YES	YES	YES	YES	YES
NAME	奖励ID	奖励类型	职业类型	概率	次数
KEY	AwardID	AwardType	JobType	Ratio	Times
TYPE	Uint32	Uint8	Uint8	Uint32	Other_;
VALUE	100	0	0	0	0;0
VALUE	101	0	0	0	0;0
VALUE	102	0	0	0	0;0
VALUE	111	0	0	0	0;0
VALUE	112	0	0	0	0;0
VALUE	113	0	0	0	0;0

generate file(ItemBase.go)
----------------------------------------------
package cfg

var Sheet1Base_cfgMap map[int]*map[string]string = map[int]*map[string]string{ 
100:&map[string]string {
"AwardID":"100",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},
101:&map[string]string {
"AwardID":"101",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},
102:&map[string]string {
"AwardID":"102",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},
111:&map[string]string {
"AwardID":"111",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},
112:&map[string]string {
"AwardID":"112",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},
113:&map[string]string {
"AwardID":"113",
"AwardType":"0",
"JobType":"0",
"Ratio":"0",
"Times":"0;0",
},
}
