package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	*gorm.Model
	Name           string    `gorm:"column:name;unique_index;default:''"`
	Password       string    `gorm:"column:password;default:''"`
	Pin            string    `gorm:"column:pin;default:''"`
	Pic            string    `gorm:"column:pic"`
	Loggedin       int       `gorm:"column:loggedin;default:0"`
	Lastlogin      time.Time `gorm:"column:lastlogin"`
	Birthday       string    `gorm:"column:birthday;default:''"`
	Banned         int       `gorm:"column:banned;default:0"`
	Banreason      string    `gorm:"column:banreason;default:''"`
	Macs           string    `gorm:"column:macs;default:''"`
	NxCredit       int       `gorm:"column:nxCredit"`
	MaplePoint     int       `gorm:"column:maplePoint"`
	NxPrepaid      int       `gorm:"column:nxPrepaid"`
	Characterslots int       `gorm:"column:characterslots;default:3"`
	Gender         int       `gorm:"column:gender;default:10"`
	Tempban        time.Time `gorm:"column:tempban"`
	Greason        int       `gorm:"column:greason;default:0"`
	Tos            int       `gorm:"column:tos;default:0"`
	Sitelogged     string    `gorm:"column:sitelogged;default:''"`
	Webadmin       int       `gorm:"column:webadmin;default:0"`
	Nick           string    `gorm:"column:nick;default:''"`
	Mute           int       `gorm:"column:mute;default:0"`
	Email          string    `gorm:"column:email;default:''"`
	Ip             string    `gorm:"column:ip;default:''"`
	Rewardpoints   int       `gorm:"column:rewardpoints;default:0"`
	Votepoints     int       `gorm:"column:votepoints;default:0"`
	Hwid           string    `gorm:"column:hwid;default:''"`
	Language       int       `gorm:"column:language;default:2"`
}

type Alliance struct {
	*gorm.Model
	Name     string `gorm:"column:name;unique_index;default:''"`
	Capacity int    `gorm:"column:capacity;default:2"`
	Notice   string `gorm:"column:notice;default:''"`
	Rank1    string `gorm:"column:rank1;default:'Master'"`
	Rank2    string `gorm:"column:rank2;default:'Jr. Master'"`
	Rank3    string `gorm:"column:rank3;default:'Member'"`
	Rank4    string `gorm:"column:rank4;default:'Member'"`
	Rank5    string `gorm:"column:rank5;default:'Member'"`
}

type Allianceguilds struct {
	*gorm.Model
	Allianceid int `gorm:"column:allianceid;default:'-1'"`
	Guildid    int `gorm:"column:guildid;default:'-1'"`
}

type AreaInfo struct {
	*gorm.Model
	Charid int    `gorm:"column:charid"`
	Area   int    `gorm:"column:area"`
	Info   string `gorm:"column:info"`
}

type BosslogDaily struct {
	*gorm.Model
	Characterid int       `gorm:"column:characterid"`
	Bosstype    BoosType  `gorm:"column:bosstype"`
	Attempttime time.Time `gorm:"column:attempttime"`
}

type BosslogWeekly struct {
	*gorm.Model
	Characterid int       `gorm:"column:characterid"`
	Bosstype    BoosType  `gorm:"column:bosstype"`
	Attempttime time.Time `gorm:"column:attempttime"`
}

type BbsReplies struct {
	*gorm.Model
	Replyid   int    `gorm:"column:replyid;auto_increment"`
	Threadid  int    `gorm:"column:threadid"`
	Postercid int    `gorm:"column:postercid"`
	Timestamp int    `gorm:"column:timestamp"`
	Content   string `gorm:"column:content;default:''"`
}

type BbsThreads struct {
	*gorm.Model
	Threadid      int    `gorm:"column:threadid;auto_increment"`
	Postercid     int    `gorm:"column:postercid"`
	Name          string `gorm:"column:name;default:''"`
	Timestamp     int    `gorm:"column:timestamp"`
	Icon          int    `gorm:"column:icon"`
	Replycount    int    `gorm:"column:replycount;default:0"`
	Startpost     string `gorm:"column:startpost"`
	Guildid       int    `gorm:"column:guildid"`
	Localthreadid int    `gorm:"column:localthreadid"`
}

type Buddies struct {
	*gorm.Model
	Characterid int    `gorm:"column:characterid"`
	Buddyid     int    `gorm:"column:buddyid"`
	Pending     int    `gorm:"column:pending;default:0"`
	Group       string `gorm:"column:group;default:'0'"`
}

type Characters struct {
	*gorm.Model
	Accountid            int       `gorm:"column:accountid;default:0"`
	World                int       `gorm:"column:world;default:0"`
	Name                 string    `gorm:"column:name;default:''"`
	Level                int       `gorm:"column:level;default:1"`
	Exp                  int       `gorm:"column:exp;default:0"`
	Gachaexp             int       `gorm:"column:gachaexp;default:0"`
	Str                  int       `gorm:"column:str;default:12"`
	Dex                  int       `gorm:"column:dex;default:5"`
	Luk                  int       `gorm:"column:luk;default:4"`
	Int                  int       `gorm:"column:int;default:4"`
	Hp                   int       `gorm:"column:hp;default:50"`
	Mp                   int       `gorm:"column:mp;default:5"`
	Maxhp                int       `gorm:"column:maxhp;default:50"`
	Maxmp                int       `gorm:"column:maxmp;default:5"`
	Meso                 int       `gorm:"column:meso;default:0"`
	HpMpUsed             int       `gorm:"column:hp_mp_used;default:0"`
	Job                  int       `gorm:"column:job;default:0"`
	Skincolor            int       `gorm:"column:skincolor;default:0"`
	Gender               int       `gorm:"column:gender;default:0"`
	Fame                 int       `gorm:"column:fame;type:int(11);default:0;NOT NULL" json:"fame"`
	Fquest               int       `gorm:"column:fquest;type:int(11);default:0;NOT NULL" json:"fquest"`
	Hair                 int       `gorm:"column:hair;type:int(11);default:0;NOT NULL" json:"hair"`
	Face                 int       `gorm:"column:face;type:int(11);default:0;NOT NULL" json:"face"`
	Ap                   int       `gorm:"column:ap;type:int(11);default:0;NOT NULL" json:"ap"`
	Sp                   string    `gorm:"column:sp;type:varchar(128);default:0,0,0,0,0,0,0,0,0,0;NOT NULL" json:"sp"`
	Map                  int       `gorm:"column:map;type:int(11);default:0;NOT NULL" json:"map"`
	Spawnpoint           int       `gorm:"column:spawnpoint;type:int(11);default:0;NOT NULL" json:"spawnpoint"`
	Gm                   int       `gorm:"column:gm;type:tinyint(1);default:0;NOT NULL" json:"gm"`
	Party                int       `gorm:"column:party;type:int(11);default:0;NOT NULL" json:"party"`
	BuddyCapacity        int       `gorm:"column:buddyCapacity;type:int(11);default:25;NOT NULL" json:"buddyCapacity"`
	Createdate           time.Time `gorm:"column:createdate;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"createdate"`
	Rank                 uint      `gorm:"column:rank;type:int(10) unsigned;default:1;NOT NULL" json:"rank"`
	RankMove             int       `gorm:"column:rankMove;type:int(11);default:0;NOT NULL" json:"rankMove"`
	JobRank              uint      `gorm:"column:jobRank;type:int(10) unsigned;default:1;NOT NULL" json:"jobRank"`
	JobRankMove          int       `gorm:"column:jobRankMove;type:int(11);default:0;NOT NULL" json:"jobRankMove"`
	Guildid              uint      `gorm:"column:guildid;type:int(10) unsigned;default:0;NOT NULL" json:"guildid"`
	Guildrank            uint      `gorm:"column:guildrank;type:int(10) unsigned;default:5;NOT NULL" json:"guildrank"`
	Messengerid          uint      `gorm:"column:messengerid;type:int(10) unsigned;default:0;NOT NULL" json:"messengerid"`
	Messengerposition    uint      `gorm:"column:messengerposition;type:int(10) unsigned;default:4;NOT NULL" json:"messengerposition"`
	Mountlevel           int       `gorm:"column:mountlevel;type:int(9);default:1;NOT NULL" json:"mountlevel"`
	Mountexp             int       `gorm:"column:mountexp;type:int(9);default:0;NOT NULL" json:"mountexp"`
	Mounttiredness       int       `gorm:"column:mounttiredness;type:int(9);default:0;NOT NULL" json:"mounttiredness"`
	Omokwins             int       `gorm:"column:omokwins;type:int(11);default:0;NOT NULL" json:"omokwins"`
	Omoklosses           int       `gorm:"column:omoklosses;type:int(11);default:0;NOT NULL" json:"omoklosses"`
	Omokties             int       `gorm:"column:omokties;type:int(11);default:0;NOT NULL" json:"omokties"`
	Matchcardwins        int       `gorm:"column:matchcardwins;type:int(11);default:0;NOT NULL" json:"matchcardwins"`
	Matchcardlosses      int       `gorm:"column:matchcardlosses;type:int(11);default:0;NOT NULL" json:"matchcardlosses"`
	Matchcardties        int       `gorm:"column:matchcardties;type:int(11);default:0;NOT NULL" json:"matchcardties"`
	MerchantMesos        int       `gorm:"column:MerchantMesos;type:int(11);default:0" json:"MerchantMesos"`
	HasMerchant          int       `gorm:"column:HasMerchant;type:tinyint(1);default:0" json:"HasMerchant"`
	Equipslots           int       `gorm:"column:equipslots;type:int(11);default:24;NOT NULL" json:"equipslots"`
	Useslots             int       `gorm:"column:useslots;type:int(11);default:24;NOT NULL" json:"useslots"`
	Setupslots           int       `gorm:"column:setupslots;type:int(11);default:24;NOT NULL" json:"setupslots"`
	Etcslots             int       `gorm:"column:etcslots;type:int(11);default:24;NOT NULL" json:"etcslots"`
	FamilyId             int       `gorm:"column:familyId;type:int(11);default:-1;NOT NULL" json:"familyId"`
	Monsterbookcover     int       `gorm:"column:monsterbookcover;type:int(11);default:0;NOT NULL" json:"monsterbookcover"`
	AllianceRank         int       `gorm:"column:allianceRank;type:int(10);default:5;NOT NULL" json:"allianceRank"`
	VanquisherStage      uint      `gorm:"column:vanquisherStage;type:int(11) unsigned;default:0;NOT NULL" json:"vanquisherStage"`
	AriantPoints         uint      `gorm:"column:ariantPoints;type:int(11) unsigned;default:0;NOT NULL" json:"ariantPoints"`
	DojoPoints           uint      `gorm:"column:dojoPoints;type:int(11) unsigned;default:0;NOT NULL" json:"dojoPoints"`
	LastDojoStage        uint      `gorm:"column:lastDojoStage;type:int(10) unsigned;default:0;NOT NULL" json:"lastDojoStage"`
	FinishedDojoTutorial uint      `gorm:"column:finishedDojoTutorial;type:tinyint(1) unsigned;default:0;NOT NULL" json:"finishedDojoTutorial"`
	VanquisherKills      uint      `gorm:"column:vanquisherKills;type:int(11) unsigned;default:0;NOT NULL" json:"vanquisherKills"`
	SummonValue          uint      `gorm:"column:summonValue;type:int(11) unsigned;default:0;NOT NULL" json:"summonValue"`
	PartnerId            int       `gorm:"column:partnerId;type:int(11);default:0;NOT NULL" json:"partnerId"`
	MarriageItemId       int       `gorm:"column:marriageItemId;type:int(11);default:0;NOT NULL" json:"marriageItemId"`
	Reborns              int       `gorm:"column:reborns;type:int(5);default:0;NOT NULL" json:"reborns"`
	PQPoints             int       `gorm:"column:PQPoints;type:int(11);default:0;NOT NULL" json:"PQPoints"`
	DataString           string    `gorm:"column:dataString;type:varchar(64);NOT NULL" json:"dataString"`
	LastLogoutTime       time.Time `gorm:"column:lastLogoutTime;type:timestamp;default:2015-01-01 05:00:00;NOT NULL" json:"lastLogoutTime"`
	LastExpGainTime      time.Time `gorm:"column:lastExpGainTime;type:timestamp;default:2015-01-01 05:00:00;NOT NULL" json:"lastExpGainTime"`
	PartySearch          int       `gorm:"column:partySearch;type:tinyint(1);default:1;NOT NULL" json:"partySearch"`
	Jailexpire           int64     `gorm:"column:jailexpire;type:bigint(20);default:0;NOT NULL" json:"jailexpire"`
}

type Cooldowns struct {
	*gorm.Model
	Charid    int    `gorm:"column:charid;type:int(11);NOT NULL" json:"charid"`
	SkillID   int    `gorm:"column:SkillID;type:int(11);NOT NULL" json:"SkillID"`
	Length    uint64 `gorm:"column:length;type:bigint(20) unsigned;NOT NULL" json:"length"`
	StartTime uint64 `gorm:"column:StartTime;type:bigint(20) unsigned;NOT NULL" json:"StartTime"`
}

type TempData struct {
	*gorm.Model
	Dropperid       int `gorm:"column:dropperid;type:int(11);primary_key" json:"dropperid"`
	Itemid          int `gorm:"column:itemid;type:int(11);default:0;NOT NULL" json:"itemid"`
	MinimumQuantity int `gorm:"column:minimum_quantity;type:int(11);default:1;NOT NULL" json:"minimum_quantity"`
	MaximumQuantity int `gorm:"column:maximum_quantity;type:int(11);default:1;NOT NULL" json:"maximum_quantity"`
	Questid         int `gorm:"column:questid;type:int(11);default:0;NOT NULL" json:"questid"`
	Chance          int `gorm:"column:chance;type:int(11);default:0;NOT NULL" json:"chance"`
}

type Skillmacros struct {
	*gorm.Model
	Characterid int    `gorm:"column:characterid;type:int(11);default:0;NOT NULL" json:"characterid"`
	Position    int    `gorm:"column:position;type:tinyint(1);default:0;NOT NULL" json:"position"`
	Skill1      int    `gorm:"column:skill1;type:int(11);default:0;NOT NULL" json:"skill1"`
	Skill2      int    `gorm:"column:skill2;type:int(11);default:0;NOT NULL" json:"skill2"`
	Skill3      int    `gorm:"column:skill3;type:int(11);default:0;NOT NULL" json:"skill3"`
	Name        string `gorm:"column:name;type:varchar(13)" json:"name"`
	Shout       int    `gorm:"column:shout;type:tinyint(1);default:0;NOT NULL" json:"shout"`
}

type Skills struct {
	*gorm.Model
	Skillid     int   `gorm:"column:skillid;type:int(11);default:0;NOT NULL" json:"skillid"`
	Characterid int   `gorm:"column:characterid;type:int(11);default:0;NOT NULL" json:"characterid"`
	Skilllevel  int   `gorm:"column:skilllevel;type:int(11);default:0;NOT NULL" json:"skilllevel"`
	Masterlevel int   `gorm:"column:masterlevel;type:int(11);default:0;NOT NULL" json:"masterlevel"`
	Expiration  int64 `gorm:"column:expiration;type:bigint(20);default:-1;NOT NULL" json:"expiration"`
}

type Specialcashitems struct {
	*gorm.Model
	Sn       int `gorm:"column:sn;type:int(11);NOT NULL" json:"sn"`
	Modifier int `gorm:"column:modifier;type:int(11);comment:1024 is add/remove;NOT NULL" json:"modifier"`
	Info     int `gorm:"column:info;type:int(1);NOT NULL" json:"info"`
}

type Storages struct {
	*gorm.Model
	Storageid uint `gorm:"column:storageid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"storageid"`
	Accountid int  `gorm:"column:accountid;type:int(11);default:0;NOT NULL" json:"accountid"`
	World     int  `gorm:"column:world;type:int(2);NOT NULL" json:"world"`
	Slots     int  `gorm:"column:slots;type:int(11);default:0;NOT NULL" json:"slots"`
	Meso      int  `gorm:"column:meso;type:int(11);default:0;NOT NULL" json:"meso"`
}

type Trocklocations struct {
	*gorm.Model
	Trockid     int `gorm:"column:trockid;type:int(11);primary_key;AUTO_INCREMENT" json:"trockid"`
	Characterid int `gorm:"column:characterid;type:int(11);NOT NULL" json:"characterid"`
	Mapid       int `gorm:"column:mapid;type:int(11);NOT NULL" json:"mapid"`
	Vip         int `gorm:"column:vip;type:int(2);NOT NULL" json:"vip"`
}

type Wishlists struct {
	*gorm.Model
	Charid int `gorm:"column:charid;type:int(11);NOT NULL" json:"charid"`
	Sn     int `gorm:"column:sn;type:int(11);NOT NULL" json:"sn"`
}

type Worldtransfers struct {
	*gorm.Model
	Characterid    int          `gorm:"column:characterid;type:int(11);NOT NULL" json:"characterid"`
	From           int          `gorm:"column:from;type:tinyint(3);NOT NULL" json:"from"`
	To             int          `gorm:"column:to;type:tinyint(3);NOT NULL" json:"to"`
	RequestTime    time.Time    `gorm:"column:requestTime;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"requestTime"`
	CompletionTime sql.NullTime `gorm:"column:completionTime;type:timestamp" json:"completionTime"`
}

type DropDataGlobal struct {
	*gorm.Model
	Continent       int    `gorm:"column:continent;type:tinyint(1);default:-1;NOT NULL" json:"continent"`
	Itemid          int    `gorm:"column:itemid;type:int(11);default:0;NOT NULL" json:"itemid"`
	MinimumQuantity int    `gorm:"column:minimum_quantity;type:int(11);default:1;NOT NULL" json:"minimum_quantity"`
	MaximumQuantity int    `gorm:"column:maximum_quantity;type:int(11);default:1;NOT NULL" json:"maximum_quantity"`
	Questid         int    `gorm:"column:questid;type:int(11);default:0;NOT NULL" json:"questid"`
	Chance          int    `gorm:"column:chance;type:int(11);default:0;NOT NULL" json:"chance"`
	Comments        string `gorm:"column:comments;type:varchar(45)" json:"comments"`
}

func (m *DropDataGlobal) TableName() string {
	return "drop_data_global"
}

type Dueyitems struct {
	*gorm.Model
	PackageId       uint `gorm:"column:PackageId;type:int(10) unsigned;default:0;NOT NULL" json:"PackageId"`
	Inventoryitemid uint `gorm:"column:inventoryitemid;type:int(10) unsigned;default:0;NOT NULL" json:"inventoryitemid"`
}

type Dueypackages struct {
	*gorm.Model
	PackageId  uint           `gorm:"column:PackageId;type:int(10) unsigned;AUTO_INCREMENT;NOT NULL" json:"PackageId"`
	ReceiverId uint           `gorm:"column:ReceiverId;type:int(10) unsigned;NOT NULL" json:"ReceiverId"`
	SenderName string         `gorm:"column:SenderName;type:varchar(13);NOT NULL" json:"SenderName"`
	Mesos      uint           `gorm:"column:Mesos;type:int(10) unsigned;default:0" json:"Mesos"`
	TimeStamp  time.Time      `gorm:"column:TimeStamp;type:timestamp;default:2015-01-01 05:00:00;NOT NULL" json:"TimeStamp"`
	Message    sql.NullString `gorm:"column:Message;type:varchar(200)" json:"Message"`
	Checked    uint           `gorm:"column:Checked;type:tinyint(1) unsigned;default:1" json:"Checked"`
	Type       uint           `gorm:"column:Type;type:tinyint(1) unsigned;default:0" json:"Type"`
}

type Eventstats struct {
	*gorm.Model
	Characterid uint   `gorm:"column:characterid;type:int(11) unsigned;primary_key" json:"characterid"`
	Name        string `gorm:"column:name;type:varchar(11);default:0;comment:0;NOT NULL" json:"name"`
	Info        int    `gorm:"column:info;type:int(11);NOT NULL" json:"info"`
}

type Famelog struct {
	*gorm.Model
	Famelogid     int       `gorm:"column:famelogid;type:int(11);primary_key;AUTO_INCREMENT" json:"famelogid"`
	Characterid   int       `gorm:"column:characterid;type:int(11);default:0;NOT NULL" json:"characterid"`
	CharacteridTo int       `gorm:"column:characterid_to;type:int(11);default:0;NOT NULL" json:"characterid_to"`
	When          time.Time `gorm:"column:when;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"when"`
}

func (m *Famelog) TableName() string {
	return "famelog"
}

type FamilyCharacter struct {
	*gorm.Model

	Cid             int    `gorm:"column:cid;type:int(11);primary_key" json:"cid"`
	Familyid        int    `gorm:"column:familyid;type:int(11);NOT NULL" json:"familyid"`
	Seniorid        int    `gorm:"column:seniorid;type:int(11);NOT NULL" json:"seniorid"`
	Reputation      int    `gorm:"column:reputation;type:int(11);default:0;NOT NULL" json:"reputation"`
	Todaysrep       int    `gorm:"column:todaysrep;type:int(11);default:0;NOT NULL" json:"todaysrep"`
	Totalreputation int    `gorm:"column:totalreputation;type:int(11);default:0;NOT NULL" json:"totalreputation"`
	Reptosenior     int    `gorm:"column:reptosenior;type:int(11);default:0;NOT NULL" json:"reptosenior"`
	Precepts        string `gorm:"column:precepts;type:varchar(200)" json:"precepts"`
	Lastresettime   int64  `gorm:"column:lastresettime;type:bigint(20);default:0;NOT NULL" json:"lastresettime"`
}

func (m *FamilyCharacter) TableName() string {
	return "family_character"
}

type FamilyEntitlement struct {
	*gorm.Model

	Charid        int   `gorm:"column:charid;type:int(11);NOT NULL" json:"charid"`
	Entitlementid int   `gorm:"column:entitlementid;type:int(11);NOT NULL" json:"entitlementid"`
	Timestamp     int64 `gorm:"column:timestamp;type:bigint(20);default:0;NOT NULL" json:"timestamp"`
}

func (m *FamilyEntitlement) TableName() string {
	return "family_entitlement"
}

type Fredstorage struct {
	*gorm.Model

	Cid       uint      `gorm:"column:cid;type:int(10) unsigned;NOT NULL" json:"cid"`
	Daynotes  uint      `gorm:"column:daynotes;type:int(4) unsigned;NOT NULL" json:"daynotes"`
	Timestamp time.Time `gorm:"column:timestamp;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"timestamp"`
}

func (m *Fredstorage) TableName() string {
	return "fredstorage"
}

type Gifts struct {
	*gorm.Model
	To      int    `gorm:"column:to;type:int(11);NOT NULL" json:"to"`
	From    string `gorm:"column:from;type:varchar(13);NOT NULL" json:"from"`
	Message string `gorm:"column:message;type:tinytext;NOT NULL" json:"message"`
	Sn      uint   `gorm:"column:sn;type:int(10) unsigned;NOT NULL" json:"sn"`
	Ringid  int    `gorm:"column:ringid;type:int(10);NOT NULL" json:"ringid"`
}

type Guilds struct {
	*gorm.Model

	Guildid     uint   `gorm:"column:guildid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"guildid"`
	Leader      uint   `gorm:"column:leader;type:int(10) unsigned;default:0;NOT NULL" json:"leader"`
	GP          uint   `gorm:"column:GP;type:int(10) unsigned;default:0;NOT NULL" json:"GP"`
	Logo        uint   `gorm:"column:logo;type:int(10) unsigned" json:"logo"`
	LogoColor   uint   `gorm:"column:logoColor;type:smallint(5) unsigned;default:0;NOT NULL" json:"logoColor"`
	Name        string `gorm:"column:name;type:varchar(45);NOT NULL" json:"name"`
	Rank1Title  string `gorm:"column:rank1title;type:varchar(45);default:Master;NOT NULL" json:"rank1title"`
	Rank2Title  string `gorm:"column:rank2title;type:varchar(45);default:Jr. Master;NOT NULL" json:"rank2title"`
	Rank3Title  string `gorm:"column:rank3title;type:varchar(45);default:Member;NOT NULL" json:"rank3title"`
	Rank4Title  string `gorm:"column:rank4title;type:varchar(45);default:Member;NOT NULL" json:"rank4title"`
	Rank5Title  string `gorm:"column:rank5title;type:varchar(45);default:Member;NOT NULL" json:"rank5title"`
	Capacity    uint   `gorm:"column:capacity;type:int(10) unsigned;default:10;NOT NULL" json:"capacity"`
	LogoBG      uint   `gorm:"column:logoBG;type:int(10) unsigned" json:"logoBG"`
	LogoBGColor uint   `gorm:"column:logoBGColor;type:smallint(5) unsigned;default:0;NOT NULL" json:"logoBGColor"`
	Notice      string `gorm:"column:notice;type:varchar(101)" json:"notice"`
	Signature   int    `gorm:"column:signature;type:int(11);default:0;NOT NULL" json:"signature"`
	AllianceId  uint   `gorm:"column:allianceId;type:int(11) unsigned;default:0;NOT NULL" json:"allianceId"`
}

type Hwidaccounts struct {
	*gorm.Model

	Accountid int       `gorm:"column:accountid;type:int(11);primary_key;default:0" json:"accountid"`
	Hwid      string    `gorm:"column:hwid;type:varchar(40);NOT NULL" json:"hwid"`
	Relevance int       `gorm:"column:relevance;type:tinyint(2);default:0;NOT NULL" json:"relevance"`
	Expiresat time.Time `gorm:"column:expiresat;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"expiresat"`
}

type Hwidbans struct {
	*gorm.Model

	Hwidbanid uint   `gorm:"column:hwidbanid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"hwidbanid"`
	Hwid      string `gorm:"column:hwid;type:varchar(30);NOT NULL" json:"hwid"`
}

type Inventoryequipment struct {
	*gorm.Model

	Inventoryequipmentid uint `gorm:"column:inventoryequipmentid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"inventoryequipmentid"`
	Inventoryitemid      uint `gorm:"column:inventoryitemid;type:int(10) unsigned;default:0;NOT NULL" json:"inventoryitemid"`
	Upgradeslots         int  `gorm:"column:upgradeslots;type:int(11);default:0;NOT NULL" json:"upgradeslots"`
	Level                int  `gorm:"column:level;type:int(11);default:0;NOT NULL" json:"level"`
	Str                  int  `gorm:"column:str;type:int(11);default:0;NOT NULL" json:"str"`
	Dex                  int  `gorm:"column:dex;type:int(11);default:0;NOT NULL" json:"dex"`
	Int                  int  `gorm:"column:int;type:int(11);default:0;NOT NULL" json:"int"`
	Luk                  int  `gorm:"column:luk;type:int(11);default:0;NOT NULL" json:"luk"`
	Hp                   int  `gorm:"column:hp;type:int(11);default:0;NOT NULL" json:"hp"`
	Mp                   int  `gorm:"column:mp;type:int(11);default:0;NOT NULL" json:"mp"`
	Watk                 int  `gorm:"column:watk;type:int(11);default:0;NOT NULL" json:"watk"`
	Matk                 int  `gorm:"column:matk;type:int(11);default:0;NOT NULL" json:"matk"`
	Wdef                 int  `gorm:"column:wdef;type:int(11);default:0;NOT NULL" json:"wdef"`
	Mdef                 int  `gorm:"column:mdef;type:int(11);default:0;NOT NULL" json:"mdef"`
	Acc                  int  `gorm:"column:acc;type:int(11);default:0;NOT NULL" json:"acc"`
	Avoid                int  `gorm:"column:avoid;type:int(11);default:0;NOT NULL" json:"avoid"`
	Hands                int  `gorm:"column:hands;type:int(11);default:0;NOT NULL" json:"hands"`
	Speed                int  `gorm:"column:speed;type:int(11);default:0;NOT NULL" json:"speed"`
	Jump                 int  `gorm:"column:jump;type:int(11);default:0;NOT NULL" json:"jump"`
	Locked               int  `gorm:"column:locked;type:int(11);default:0;NOT NULL" json:"locked"`
	Vicious              uint `gorm:"column:vicious;type:int(11) unsigned;default:0;NOT NULL" json:"vicious"`
	Itemlevel            int  `gorm:"column:itemlevel;type:int(11);default:1;NOT NULL" json:"itemlevel"`
	Itemexp              uint `gorm:"column:itemexp;type:int(11) unsigned;default:0;NOT NULL" json:"itemexp"`
	Ringid               int  `gorm:"column:ringid;type:int(11);default:-1;NOT NULL" json:"ringid"`
}

func (m *Inventoryequipment) TableName() string {
	return "inventoryequipment"
}

type Inventoryitems struct {
	*gorm.Model

	Inventoryitemid uint   `gorm:"column:inventoryitemid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"inventoryitemid"`
	Type            uint   `gorm:"column:type;type:tinyint(3) unsigned;NOT NULL" json:"type"`
	Characterid     int    `gorm:"column:characterid;type:int(11)" json:"characterid"`
	Accountid       int    `gorm:"column:accountid;type:int(11)" json:"accountid"`
	Itemid          int    `gorm:"column:itemid;type:int(11);default:0;NOT NULL" json:"itemid"`
	Inventorytype   int    `gorm:"column:inventorytype;type:int(11);default:0;NOT NULL" json:"inventorytype"`
	Position        int    `gorm:"column:position;type:int(11);default:0;NOT NULL" json:"position"`
	Quantity        int    `gorm:"column:quantity;type:int(11);default:0;NOT NULL" json:"quantity"`
	Owner           string `gorm:"column:owner;type:tinytext;NOT NULL" json:"owner"`
	Petid           int    `gorm:"column:petid;type:int(11);default:-1;NOT NULL" json:"petid"`
	Flag            int    `gorm:"column:flag;type:int(11);NOT NULL" json:"flag"`
	Expiration      int64  `gorm:"column:expiration;type:bigint(20);default:-1;NOT NULL" json:"expiration"`
	GiftFrom        string `gorm:"column:giftFrom;type:varchar(26);NOT NULL" json:"giftFrom"`
}

type Inventorymerchant struct {
	*gorm.Model

	Inventorymerchantid uint `gorm:"column:inventorymerchantid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"inventorymerchantid"`
	Inventoryitemid     uint `gorm:"column:inventoryitemid;type:int(10) unsigned;default:0;NOT NULL" json:"inventoryitemid"`
	Characterid         int  `gorm:"column:characterid;type:int(11)" json:"characterid"`
	Bundles             int  `gorm:"column:bundles;type:int(10);default:0;NOT NULL" json:"bundles"`
}

func (m *Inventorymerchant) TableName() string {
	return "inventorymerchant"
}

type Ipbans struct {
	*gorm.Model

	Ipbanid uint   `gorm:"column:ipbanid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"ipbanid"`
	Ip      string `gorm:"column:ip;type:varchar(40);NOT NULL" json:"ip"`
	Aid     string `gorm:"column:aid;type:varchar(40)" json:"aid"`
}

type Keymap struct {
	*gorm.Model

	Characterid int `gorm:"column:characterid;type:int(11);default:0;NOT NULL" json:"characterid"`
	Key         int `gorm:"column:key;type:int(11);default:0;NOT NULL" json:"key"`
	Type        int `gorm:"column:type;type:int(11);default:0;NOT NULL" json:"type"`
	Action      int `gorm:"column:action;type:int(11);default:0;NOT NULL" json:"action"`
}

func (m *Keymap) TableName() string {
	return "keymap"
}

type Macbans struct {
	*gorm.Model

	Macbanid uint   `gorm:"column:macbanid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"macbanid"`
	Mac      string `gorm:"column:mac;type:varchar(30);NOT NULL" json:"mac"`
	Aid      string `gorm:"column:aid;type:varchar(40)" json:"aid"`
}

type Macfilters struct {
	*gorm.Model

	Macfilterid uint   `gorm:"column:macfilterid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"macfilterid"`
	Filter      string `gorm:"column:filter;type:varchar(30);NOT NULL" json:"filter"`
}

type Makercreatedata struct {
	*gorm.Model

	Itemid        int  `gorm:"column:itemid;type:int(11);NOT NULL" json:"itemid"`
	ReqLevel      uint `gorm:"column:req_level;type:tinyint(3) unsigned;NOT NULL" json:"req_level"`
	ReqMakerLevel uint `gorm:"column:req_maker_level;type:tinyint(3) unsigned;NOT NULL" json:"req_maker_level"`
	ReqMeso       int  `gorm:"column:req_meso;type:int(11);NOT NULL" json:"req_meso"`
	ReqItem       int  `gorm:"column:req_item;type:int(11);NOT NULL" json:"req_item"`
	ReqEquip      int  `gorm:"column:req_equip;type:int(11);NOT NULL" json:"req_equip"`
	Catalyst      int  `gorm:"column:catalyst;type:int(11);NOT NULL" json:"catalyst"`
	Quantity      int  `gorm:"column:quantity;type:smallint(6);NOT NULL" json:"quantity"`
	Tuc           int  `gorm:"column:tuc;type:tinyint(3);NOT NULL" json:"tuc"`
}

type Makerrecipedata struct {
	*gorm.Model

	Itemid  int `gorm:"column:itemid;type:int(11);primary_key" json:"itemid"`
	ReqItem int `gorm:"column:req_item;type:int(11);NOT NULL" json:"req_item"`
	Count   int `gorm:"column:count;type:smallint(6);NOT NULL" json:"count"`
}

type Makerrewarddata struct {
	*gorm.Model

	Itemid   int  `gorm:"column:itemid;type:int(11);primary_key" json:"itemid"`
	Rewardid int  `gorm:"column:rewardid;type:int(11);NOT NULL" json:"rewardid"`
	Quantity int  `gorm:"column:quantity;type:smallint(6);NOT NULL" json:"quantity"`
	Prob     uint `gorm:"column:prob;type:tinyint(3) unsigned;default:100;NOT NULL" json:"prob"`
}

type Pets struct {
	*gorm.Model

	Petid     uint   `gorm:"column:petid;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"petid"`
	Name      string `gorm:"column:name;type:varchar(13)" json:"name"`
	Level     uint   `gorm:"column:level;type:int(10) unsigned;NOT NULL" json:"level"`
	Closeness uint   `gorm:"column:closeness;type:int(10) unsigned;NOT NULL" json:"closeness"`
	Fullness  uint   `gorm:"column:fullness;type:int(10) unsigned;NOT NULL" json:"fullness"`
	Summoned  int    `gorm:"column:summoned;type:tinyint(1);default:0;NOT NULL" json:"summoned"`
	Flag      uint   `gorm:"column:flag;type:int(10) unsigned;default:0;NOT NULL" json:"flag"`
}

type Petignores struct {
	*gorm.Model

	Petid  uint `gorm:"column:petid;type:int(11) unsigned;NOT NULL" json:"petid"`
	Itemid uint `gorm:"column:itemid;type:int(10) unsigned;NOT NULL" json:"itemid"`
}

type Playerdiseases struct {
	*gorm.Model

	Charid     int `gorm:"column:charid;type:int(11);NOT NULL" json:"charid"`
	Disease    int `gorm:"column:disease;type:int(11);NOT NULL" json:"disease"`
	Mobskillid int `gorm:"column:mobskillid;type:int(11);NOT NULL" json:"mobskillid"`
	Mobskilllv int `gorm:"column:mobskilllv;type:int(11);NOT NULL" json:"mobskilllv"`
	Length     int `gorm:"column:length;type:int(11);default:1;NOT NULL" json:"length"`
}

type Playernpcs struct {
	*gorm.Model

	Name         string `gorm:"column:name;type:varchar(13);NOT NULL" json:"name"`
	Hair         int    `gorm:"column:hair;type:int(11);NOT NULL" json:"hair"`
	Face         int    `gorm:"column:face;type:int(11);NOT NULL" json:"face"`
	Skin         int    `gorm:"column:skin;type:int(11);NOT NULL" json:"skin"`
	Gender       int    `gorm:"column:gender;type:int(11);default:0;NOT NULL" json:"gender"`
	X            int    `gorm:"column:x;type:int(11);NOT NULL" json:"x"`
	Cy           int    `gorm:"column:cy;type:int(11);default:0;NOT NULL" json:"cy"`
	World        int    `gorm:"column:world;type:int(11);default:0;NOT NULL" json:"world"`
	Map          int    `gorm:"column:map;type:int(11);default:0;NOT NULL" json:"map"`
	Dir          int    `gorm:"column:dir;type:int(11);default:0;NOT NULL" json:"dir"`
	Scriptid     uint   `gorm:"column:scriptid;type:int(10) unsigned;default:0;NOT NULL" json:"scriptid"`
	Fh           int    `gorm:"column:fh;type:int(11);default:0;NOT NULL" json:"fh"`
	Rx0          int    `gorm:"column:rx0;type:int(11);default:0;NOT NULL" json:"rx0"`
	Rx1          int    `gorm:"column:rx1;type:int(11);default:0;NOT NULL" json:"rx1"`
	Worldrank    int    `gorm:"column:worldrank;type:int(11);default:0;NOT NULL" json:"worldrank"`
	Overallrank  int    `gorm:"column:overallrank;type:int(11);default:0;NOT NULL" json:"overallrank"`
	Worldjobrank int    `gorm:"column:worldjobrank;type:int(11);default:0;NOT NULL" json:"worldjobrank"`
	Job          int    `gorm:"column:job;type:int(11);default:0;NOT NULL" json:"job"`
}

type PlayernpcsEquip struct {
	*gorm.Model

	Npcid    int `gorm:"column:npcid;type:int(11);default:0;NOT NULL" json:"npcid"`
	Equipid  int `gorm:"column:equipid;type:int(11);NOT NULL" json:"equipid"`
	Type     int `gorm:"column:type;type:int(11);default:0;NOT NULL" json:"type"`
	Equippos int `gorm:"column:equippos;type:int(11);NOT NULL" json:"equippos"`
}

func (m *PlayernpcsEquip) TableName() string {
	return "playernpcs_equip"
}

type PlayernpcsField struct {
	*gorm.Model

	World  int `gorm:"column:world;type:int(11);NOT NULL" json:"world"`
	Map    int `gorm:"column:map;type:int(11);NOT NULL" json:"map"`
	Step   int `gorm:"column:step;type:tinyint(1);default:0;NOT NULL" json:"step"`
	Podium int `gorm:"column:podium;type:smallint(8);default:0;NOT NULL" json:"podium"`
}

func (m *PlayernpcsField) TableName() string {
	return "playernpcs_field"
}

type Plife struct {
	*gorm.Model

	World   int    `gorm:"column:world;type:int(11);default:-1;NOT NULL" json:"world"`
	Map     int    `gorm:"column:map;type:int(11);default:0;NOT NULL" json:"map"`
	Life    int    `gorm:"column:life;type:int(11);default:0;NOT NULL" json:"life"`
	Type    string `gorm:"column:type;type:varchar(1);default:n;NOT NULL" json:"type"`
	Cy      int    `gorm:"column:cy;type:int(11);default:0;NOT NULL" json:"cy"`
	F       int    `gorm:"column:f;type:int(11);default:0;NOT NULL" json:"f"`
	Fh      int    `gorm:"column:fh;type:int(11);default:0;NOT NULL" json:"fh"`
	Rx0     int    `gorm:"column:rx0;type:int(11);default:0;NOT NULL" json:"rx0"`
	Rx1     int    `gorm:"column:rx1;type:int(11);default:0;NOT NULL" json:"rx1"`
	X       int    `gorm:"column:x;type:int(11);default:0;NOT NULL" json:"x"`
	Y       int    `gorm:"column:y;type:int(11);default:0;NOT NULL" json:"y"`
	Hide    int    `gorm:"column:hide;type:int(11);default:0;NOT NULL" json:"hide"`
	Mobtime int    `gorm:"column:mobtime;type:int(11);default:0;NOT NULL" json:"mobtime"`
	Team    int    `gorm:"column:team;type:int(11);default:0;NOT NULL" json:"team"`
}

func (m *Plife) TableName() string {
	return "plife"
}

type Questactions struct {
	*gorm.Model

	Questactionid uint   `gorm:"column:questactionid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"questactionid"`
	Questid       int    `gorm:"column:questid;type:int(11);default:0;NOT NULL" json:"questid"`
	Status        int    `gorm:"column:status;type:int(11);default:0;NOT NULL" json:"status"`
	Data          string `gorm:"column:data;type:blob;NOT NULL" json:"data"`
}

type Questprogress struct {
	*gorm.Model

	Characterid   int    `gorm:"column:characterid;type:int(11);NOT NULL" json:"characterid"`
	Queststatusid uint   `gorm:"column:queststatusid;type:int(10) unsigned;default:0;NOT NULL" json:"queststatusid"`
	Progressid    int    `gorm:"column:progressid;type:int(11);default:0;NOT NULL" json:"progressid"`
	Progress      string `gorm:"column:progress;type:varchar(15);NOT NULL" json:"progress"`
}

func (m *Questprogress) TableName() string {
	return "questprogress"
}

type Questrequirements struct {
	*gorm.Model

	Questrequirementid uint   `gorm:"column:questrequirementid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"questrequirementid"`
	Questid            int    `gorm:"column:questid;type:int(11);default:0;NOT NULL" json:"questid"`
	Status             int    `gorm:"column:status;type:int(11);default:0;NOT NULL" json:"status"`
	Data               string `gorm:"column:data;type:blob;NOT NULL" json:"data"`
}

type Queststatus struct {
	*gorm.Model

	Queststatusid uint  `gorm:"column:queststatusid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"queststatusid"`
	Characterid   int   `gorm:"column:characterid;type:int(11);default:0;NOT NULL" json:"characterid"`
	Quest         int   `gorm:"column:quest;type:int(11);default:0;NOT NULL" json:"quest"`
	Status        int   `gorm:"column:status;type:int(11);default:0;NOT NULL" json:"status"`
	Time          int   `gorm:"column:time;type:int(11);default:0;NOT NULL" json:"time"`
	Expires       int64 `gorm:"column:expires;type:bigint(20);default:0;NOT NULL" json:"expires"`
	Forfeited     int   `gorm:"column:forfeited;type:int(11);default:0;NOT NULL" json:"forfeited"`
	Completed     int   `gorm:"column:completed;type:int(11);default:0;NOT NULL" json:"completed"`
	Info          int   `gorm:"column:info;type:tinyint(3);default:0;NOT NULL" json:"info"`
}

func (m *Queststatus) TableName() string {
	return "queststatus"
}

type Quickslotkeymapped struct {
	*gorm.Model

	Accountid int   `gorm:"column:accountid;type:int(11);primary_key" json:"accountid"`
	Keymap    int64 `gorm:"column:keymap;type:bigint(20);default:0;NOT NULL" json:"keymap"`
}

func (m *Quickslotkeymapped) TableName() string {
	return "quickslotkeymapped"
}

type Reactordrops struct {
	*gorm.Model

	Reactordropid uint `gorm:"column:reactordropid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"reactordropid"`
	Reactorid     int  `gorm:"column:reactorid;type:int(11);NOT NULL" json:"reactorid"`
	Itemid        int  `gorm:"column:itemid;type:int(11);NOT NULL" json:"itemid"`
	Chance        int  `gorm:"column:chance;type:int(11);NOT NULL" json:"chance"`
	Questid       int  `gorm:"column:questid;type:int(5);default:-1;NOT NULL" json:"questid"`
}

type MtsCart struct {
	*gorm.Model

	Cid    int `gorm:"column:cid;type:int(11);NOT NULL" json:"cid"`
	Itemid int `gorm:"column:itemid;type:int(11);NOT NULL" json:"itemid"`
}

func (m *MtsCart) TableName() string {
	return "mts_cart"
}

type MtsItems struct {
	*gorm.Model

	Tab          int    `gorm:"column:tab;type:int(11);default:0;NOT NULL" json:"tab"`
	Type         int    `gorm:"column:type;type:int(11);default:0;NOT NULL" json:"type"`
	Itemid       uint   `gorm:"column:itemid;type:int(10) unsigned;default:0;NOT NULL" json:"itemid"`
	Quantity     int    `gorm:"column:quantity;type:int(11);default:1;NOT NULL" json:"quantity"`
	Seller       int    `gorm:"column:seller;type:int(11);default:0;NOT NULL" json:"seller"`
	Price        int    `gorm:"column:price;type:int(11);default:0;NOT NULL" json:"price"`
	BidIncre     int    `gorm:"column:bid_incre;type:int(11);default:0" json:"bid_incre"`
	BuyNow       int    `gorm:"column:buy_now;type:int(11);default:0" json:"buy_now"`
	Position     int    `gorm:"column:position;type:int(11);default:0" json:"position"`
	Upgradeslots int    `gorm:"column:upgradeslots;type:int(11);default:0" json:"upgradeslots"`
	Level        int    `gorm:"column:level;type:int(11);default:0" json:"level"`
	Itemlevel    int    `gorm:"column:itemlevel;type:int(11);default:1;NOT NULL" json:"itemlevel"`
	Itemexp      uint   `gorm:"column:itemexp;type:int(11) unsigned;default:0;NOT NULL" json:"itemexp"`
	Ringid       int    `gorm:"column:ringid;type:int(11);default:-1;NOT NULL" json:"ringid"`
	Str          int    `gorm:"column:str;type:int(11);default:0" json:"str"`
	Dex          int    `gorm:"column:dex;type:int(11);default:0" json:"dex"`
	Int          int    `gorm:"column:int;type:int(11);default:0" json:"int"`
	Luk          int    `gorm:"column:luk;type:int(11);default:0" json:"luk"`
	Hp           int    `gorm:"column:hp;type:int(11);default:0" json:"hp"`
	Mp           int    `gorm:"column:mp;type:int(11);default:0" json:"mp"`
	Watk         int    `gorm:"column:watk;type:int(11);default:0" json:"watk"`
	Matk         int    `gorm:"column:matk;type:int(11);default:0" json:"matk"`
	Wdef         int    `gorm:"column:wdef;type:int(11);default:0" json:"wdef"`
	Mdef         int    `gorm:"column:mdef;type:int(11);default:0" json:"mdef"`
	Acc          int    `gorm:"column:acc;type:int(11);default:0" json:"acc"`
	Avoid        int    `gorm:"column:avoid;type:int(11);default:0" json:"avoid"`
	Hands        int    `gorm:"column:hands;type:int(11);default:0" json:"hands"`
	Speed        int    `gorm:"column:speed;type:int(11);default:0" json:"speed"`
	Jump         int    `gorm:"column:jump;type:int(11);default:0" json:"jump"`
	Locked       int    `gorm:"column:locked;type:int(11);default:0" json:"locked"`
	Isequip      int    `gorm:"column:isequip;type:int(1);default:0" json:"isequip"`
	Owner        string `gorm:"column:owner;type:varchar(16)" json:"owner"`
	Sellername   string `gorm:"column:sellername;type:varchar(16);NOT NULL" json:"sellername"`
	SellEnds     string `gorm:"column:sell_ends;type:varchar(16);NOT NULL" json:"sell_ends"`
	Transfer     int    `gorm:"column:transfer;type:int(2);default:0" json:"transfer"`
	Vicious      uint   `gorm:"column:vicious;type:int(2) unsigned;default:0;NOT NULL" json:"vicious"`
	Flag         uint   `gorm:"column:flag;type:int(2) unsigned;default:0;NOT NULL" json:"flag"`
	Expiration   int64  `gorm:"column:expiration;type:bigint(20);default:-1;NOT NULL" json:"expiration"`
	GiftFrom     string `gorm:"column:giftFrom;type:varchar(26);NOT NULL" json:"giftFrom"`
}

type Namechanges struct {
	*gorm.Model

	Characterid    int          `gorm:"column:characterid;type:int(11);NOT NULL" json:"characterid"`
	Old            string       `gorm:"column:old;type:varchar(13);NOT NULL" json:"old"`
	New            string       `gorm:"column:new;type:varchar(13);NOT NULL" json:"new"`
	RequestTime    time.Time    `gorm:"column:requestTime;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"requestTime"`
	CompletionTime sql.NullTime `gorm:"column:completionTime;type:timestamp" json:"completionTime"`
}

type Newyear struct {
	*gorm.Model

	Senderid        int    `gorm:"column:senderid;type:int(10);default:-1;NOT NULL" json:"senderid"`
	Sendername      string `gorm:"column:sendername;type:varchar(13)" json:"sendername"`
	Receiverid      int    `gorm:"column:receiverid;type:int(10);default:-1;NOT NULL" json:"receiverid"`
	Receivername    string `gorm:"column:receivername;type:varchar(13)" json:"receivername"`
	Message         string `gorm:"column:message;type:varchar(120)" json:"message"`
	Senderdiscard   int    `gorm:"column:senderdiscard;type:tinyint(1);default:0;NOT NULL" json:"senderdiscard"`
	Receiverdiscard int    `gorm:"column:receiverdiscard;type:tinyint(1);default:0;NOT NULL" json:"receiverdiscard"`
	Received        int    `gorm:"column:received;type:tinyint(1);default:0;NOT NULL" json:"received"`
	Timesent        uint64 `gorm:"column:timesent;type:bigint(20) unsigned;NOT NULL" json:"timesent"`
	Timereceived    uint64 `gorm:"column:timereceived;type:bigint(20) unsigned;NOT NULL" json:"timereceived"`
}

func (m *Newyear) TableName() string {
	return "newyear"
}

type Notes struct {
	*gorm.Model

	To        string `gorm:"column:to;type:varchar(13);NOT NULL" json:"to"`
	From      string `gorm:"column:from;type:varchar(13);NOT NULL" json:"from"`
	Message   string `gorm:"column:message;type:text;NOT NULL" json:"message"`
	Timestamp uint64 `gorm:"column:timestamp;type:bigint(20) unsigned;NOT NULL" json:"timestamp"`
	Fame      int    `gorm:"column:fame;type:int(11);default:0;NOT NULL" json:"fame"`
	Deleted   int    `gorm:"column:deleted;type:int(2);default:0;NOT NULL" json:"deleted"`
}

type Nxcode struct {
	*gorm.Model

	Code       string `gorm:"column:code;type:varchar(17);unique;NOT NULL" json:"code"`
	Retriever  string `gorm:"column:retriever;type:varchar(13)" json:"retriever"`
	Expiration uint64 `gorm:"column:expiration;type:bigint(20) unsigned;default:0;NOT NULL" json:"expiration"`
}

func (m *Nxcode) TableName() string {
	return "nxcode"
}

type NxcodeItems struct {
	*gorm.Model

	Codeid   int `gorm:"column:codeid;type:int(11);NOT NULL" json:"codeid"`
	Type     int `gorm:"column:type;type:int(11);default:5;NOT NULL" json:"type"`
	Item     int `gorm:"column:item;type:int(11);default:4000000;NOT NULL" json:"item"`
	Quantity int `gorm:"column:quantity;type:int(11);default:1;NOT NULL" json:"quantity"`
}

type Nxcoupons struct {
	*gorm.Model

	Couponid  int `gorm:"column:couponid;type:int(11);default:0;NOT NULL" json:"couponid"`
	Rate      int `gorm:"column:rate;type:int(11);default:0;NOT NULL" json:"rate"`
	Activeday int `gorm:"column:activeday;type:int(11);default:0;NOT NULL" json:"activeday"`
	Starthour int `gorm:"column:starthour;type:int(11);default:0;NOT NULL" json:"starthour"`
	Endhour   int `gorm:"column:endhour;type:int(11);default:0;NOT NULL" json:"endhour"`
}

type Marriages struct {
	*gorm.Model

	Marriageid uint `gorm:"column:marriageid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"marriageid"`
	Husbandid  uint `gorm:"column:husbandid;type:int(10) unsigned;default:0;NOT NULL" json:"husbandid"`
	Wifeid     uint `gorm:"column:wifeid;type:int(10) unsigned;default:0;NOT NULL" json:"wifeid"`
}

type Medalmaps struct {
	*gorm.Model

	Characterid   int  `gorm:"column:characterid;type:int(11);NOT NULL" json:"characterid"`
	Queststatusid uint `gorm:"column:queststatusid;type:int(11) unsigned;NOT NULL" json:"queststatusid"`
	Mapid         int  `gorm:"column:mapid;type:int(11);NOT NULL" json:"mapid"`
}

type Monsterbook struct {
	*gorm.Model

	Charid uint `gorm:"column:charid;type:int(11) unsigned;NOT NULL" json:"charid"`
	Cardid int  `gorm:"column:cardid;type:int(11);NOT NULL" json:"cardid"`
	Level  int  `gorm:"column:level;type:int(1);default:1" json:"level"`
}

func (m *Monsterbook) TableName() string {
	return "monsterbook"
}

type Monstercarddata struct {
	*gorm.Model

	Cardid int `gorm:"column:cardid;type:int(11);default:0;NOT NULL" json:"cardid"`
	Mobid  int `gorm:"column:mobid;type:int(11);default:0;NOT NULL" json:"mobid"`
}

type Makerreagentdata struct {
	*gorm.Model

	Itemid int    `gorm:"column:itemid;type:int(11);primary_key" json:"itemid"`
	Stat   string `gorm:"column:stat;type:varchar(20);NOT NULL" json:"stat"`
	Value  int    `gorm:"column:value;type:smallint(6);NOT NULL" json:"value"`
}

type Shops struct {
	*gorm.Model

	Shopid uint `gorm:"column:shopid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"shopid"`
	Npcid  int  `gorm:"column:npcid;type:int(11);default:0;NOT NULL" json:"npcid"`
}

type Reports struct {
	*gorm.Model

	Reporttime  time.Time `gorm:"column:reporttime;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"reporttime"`
	Reporterid  int       `gorm:"column:reporterid;type:int(11);NOT NULL" json:"reporterid"`
	Victimid    int       `gorm:"column:victimid;type:int(11);NOT NULL" json:"victimid"`
	Reason      int       `gorm:"column:reason;type:tinyint(4);NOT NULL" json:"reason"`
	Chatlog     string    `gorm:"column:chatlog;type:text;NOT NULL" json:"chatlog"`
	Description string    `gorm:"column:description;type:text;NOT NULL" json:"description"`
}

type Responses struct {
	*gorm.Model

	Chat     string `gorm:"column:chat;type:text" json:"chat"`
	Response string `gorm:"column:response;type:text" json:"response"`
	Id       uint   `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
}

type Rings struct {
	*gorm.Model

	PartnerRingId int    `gorm:"column:partnerRingId;type:int(11);default:0;NOT NULL" json:"partnerRingId"`
	PartnerChrId  int    `gorm:"column:partnerChrId;type:int(11);default:0;NOT NULL" json:"partnerChrId"`
	Itemid        int    `gorm:"column:itemid;type:int(11);default:0;NOT NULL" json:"itemid"`
	Partnername   string `gorm:"column:partnername;type:varchar(255);NOT NULL" json:"partnername"`
}

type Savedlocations struct {
	*gorm.Model

	Characterid  int    `gorm:"column:characterid;type:int(11);NOT NULL" json:"characterid"`
	Locationtype string `gorm:"column:locationtype;type:enum('FREE_MARKET','WORLDTOUR','FLORINA','INTRO','SUNDAY_MARKET','MIRROR','EVENT','BOSSPQ','HAPPYVILLE','DEVELOPER','MONSTER_CARNIVAL');NOT NULL" json:"locationtype"`
	Map          int    `gorm:"column:map;type:int(11);NOT NULL" json:"map"`
	Portal       int    `gorm:"column:portal;type:int(11);NOT NULL" json:"portal"`
}

type ServerQueue struct {
	*gorm.Model

	Accountid   int       `gorm:"column:accountid;type:int(11);default:0;NOT NULL" json:"accountid"`
	Characterid int       `gorm:"column:characterid;type:int(11);default:0;NOT NULL" json:"characterid"`
	Type        int       `gorm:"column:type;type:tinyint(2);default:0;NOT NULL" json:"type"`
	Value       int       `gorm:"column:value;type:int(10);default:0;NOT NULL" json:"value"`
	Message     string    `gorm:"column:message;type:varchar(128);NOT NULL" json:"message"`
	CreateTime  time.Time `gorm:"column:createTime;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"createTime"`
}

func (m *ServerQueue) TableName() string {
	return "server_queue"
}

type Shopitems struct {
	*gorm.Model

	Shopitemid uint `gorm:"column:shopitemid;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"shopitemid"`
	Shopid     uint `gorm:"column:shopid;type:int(10) unsigned;NOT NULL" json:"shopid"`
	Itemid     int  `gorm:"column:itemid;type:int(11);NOT NULL" json:"itemid"`
	Price      int  `gorm:"column:price;type:int(11);NOT NULL" json:"price"`
	Pitch      int  `gorm:"column:pitch;type:int(11);default:0;NOT NULL" json:"pitch"`
	Position   int  `gorm:"column:position;type:int(11);comment:sort is an arbitrary field designed to give leeway when modifying shops. The lowest number is 104 and it increments by 4 for each item to allow decent space for swapping/inserting/removing items.;NOT NULL" json:"position"`
}
