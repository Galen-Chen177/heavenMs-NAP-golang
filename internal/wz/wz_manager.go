package wz

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// WzDir 根目录
	WzDir *os.File

	// WzBaseDir 基础文件夹
	WzBaseDir *os.File

	// WzCharacterDir 角色文件夹
	WzCharacterDir *os.File

	// WzEffectDir 特效文件夹
	WzEffectDir *os.File

	// WzEtcDir 装饰文件夹
	WzEtcDir *os.File

	// WzItemDir 物品文件夹
	WzItemDir *os.File

	// WzMapDir 地图文件夹
	WzMapDir *os.File

	// WzMobDir Mod文件夹(模型?)
	WzMobDir *os.File

	// WzMorphDir 变形文件夹(什么鬼?)
	WzMorphDir *os.File

	// WzNpcDir NPC文件夹
	WzNpcDir *os.File

	// WzQuestDir 任务文件夹?
	WzQuestDir *os.File

	// WzReactorDir 反应堆文件夹(什么东西?)
	WzReactorDir *os.File

	// WzSkillDir 技能文件夹
	WzSkillDir *os.File

	// WzSoundDir 声音文件夹
	WzSoundDir *os.File

	// WzStringDir 文本文件夹
	WzStringDir *os.File

	// WzTamingMobDir 驯服模型文件夹(宠物?)
	WzTamingMobDir *os.File

	// WzUiDir UI画面文件夹
	WzUiDir *os.File
)

func InitWzManage(wzPath string) {
	logrus.Info("[wz] 开始初始化")
	handleErr := func(err error) {
		if err != nil {
			logrus.Fatalf("wz 文件夹打开失败")
			return
		}
	}

	WzDir, err := os.Open(wzPath)
	handleErr(err)
	WzBaseDir, err = ChildFile(WzDir, "Base.wz")
	handleErr(err)
	WzCharacterDir, err = ChildFile(WzDir, "Character.wz")
	handleErr(err)
	WzEffectDir, err = ChildFile(WzDir, "Effect.wz")
	handleErr(err)
	WzEtcDir, err = ChildFile(WzDir, "Etc.wz")
	handleErr(err)
	WzItemDir, err = ChildFile(WzDir, "Item.wz")
	handleErr(err)
	WzMapDir, err = ChildFile(WzDir, "Map.wz")
	handleErr(err)
	WzMobDir, err = ChildFile(WzDir, "Mob.wz")
	handleErr(err)
	WzMorphDir, err = ChildFile(WzDir, "Morph.wz")
	handleErr(err)
	WzNpcDir, err = ChildFile(WzDir, "Npc.wz")
	handleErr(err)
	WzQuestDir, err = ChildFile(WzDir, "Quest.wz")
	handleErr(err)
	WzReactorDir, err = ChildFile(WzDir, "Reactor.wz")
	handleErr(err)
	WzSkillDir, err = ChildFile(WzDir, "Skill.wz")
	handleErr(err)
	WzSoundDir, err = ChildFile(WzDir, "Sound.wz")
	handleErr(err)
	WzStringDir, err = ChildFile(WzDir, "String.wz")
	handleErr(err)
	WzTamingMobDir, err = ChildFile(WzDir, "TamingMob.wz")
	handleErr(err)
	WzUiDir, err = ChildFile(WzDir, "UI.wz")
	handleErr(err)
}

// ChildFile 获取子文件
func ChildFile(parent *os.File, name string) (*os.File, error) {
	return os.Open(fmt.Sprintf("%s%s%s", parent.Name(), string(os.PathSeparator), name))
}
